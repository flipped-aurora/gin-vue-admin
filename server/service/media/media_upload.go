package media

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/media"
	"github.com/flipped-aurora/gin-vue-admin/server/model/media/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/media/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/upload"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MediaUploadService struct{}

func (s *MediaUploadService) Init(userID uint, req request.UploadInitReq) (response.UploadInitResp, error) {
	var resp response.UploadInitResp

	// 1) 秒传:同用户已完成同 hash
	var done media.MediaUpload
	err := global.GVA_DB.Where("user_id = ? AND file_hash = ? AND status = ?", userID, req.FileHash, media.UploadStatusCompleted).First(&done).Error
	if err == nil && done.MediaID != 0 {
		var m media.FileUploadAndDownload
		if e := global.GVA_DB.First(&m, done.MediaID).Error; e == nil {
			// 为当前用户复制一条媒体库记录,共享 storage key
			nm := media.FileUploadAndDownload{Name: req.FileName, Url: m.Url, Tag: m.Tag, Key: m.Key}
			if e = global.GVA_DB.Create(&nm).Error; e == nil {
				resp.Instant = true
				resp.Media = &nm
				return resp, nil
			}
		}
	}

	// 2) 取/建进行中会话
	var up media.MediaUpload
	err = global.GVA_DB.Where("user_id = ? AND file_hash = ? AND status = ?", userID, req.FileHash, media.UploadStatusUploading).First(&up).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		up = media.MediaUpload{
			UserID: userID, FileName: req.FileName, FileHash: req.FileHash,
			FileSize: req.FileSize, ChunkSize: req.ChunkSize, ChunkTotal: req.ChunkTotal,
			Status: media.UploadStatusUploading,
		}
		if err = global.GVA_DB.Create(&up).Error; err != nil {
			return resp, err
		}
	} else if err != nil {
		return resp, err
	}

	// 3) 已收分片索引
	var idx []int
	global.GVA_DB.Model(&media.MediaUploadChunk{}).Where("upload_id = ?", up.ID).
		Order("chunk_index").Pluck("chunk_index", &idx)
	sort.Ints(idx)
	resp.UploadID = up.ID
	resp.UploadedChunks = idx
	return resp, nil
}

func (s *MediaUploadService) SaveChunk(userID, uploadID uint, index int, chunkHash string, data []byte) error {
	var up media.MediaUpload
	if err := global.GVA_DB.First(&up, uploadID).Error; err != nil {
		return errors.New("上传会话不存在")
	}
	if up.UserID != userID {
		return errors.New("无权操作该上传")
	}
	if up.Status != media.UploadStatusUploading {
		return errors.New("上传会话状态不允许收片")
	}
	sum := md5.Sum(data)
	if hex.EncodeToString(sum[:]) != chunkHash {
		return fmt.Errorf("分片 %d 校验失败", index)
	}
	if _, err := upload.SaveChunkFile(uploadID, index, data); err != nil {
		return err
	}
	rec := media.MediaUploadChunk{UploadID: uploadID, ChunkIndex: index, ChunkHash: chunkHash, Size: int64(len(data))}
	return global.GVA_DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "upload_id"}, {Name: "chunk_index"}},
		DoNothing: true,
	}).Create(&rec).Error
}

func (s *MediaUploadService) Complete(userID, uploadID uint) (media.FileUploadAndDownload, error) {
	var m media.FileUploadAndDownload
	var up media.MediaUpload
	if err := global.GVA_DB.First(&up, uploadID).Error; err != nil {
		return m, errors.New("上传会话不存在")
	}
	if up.UserID != userID {
		return m, errors.New("无权操作该上传")
	}
	// 原子抢占:仅 uploading -> merging 的赢家继续
	res := global.GVA_DB.Model(&media.MediaUpload{}).
		Where("id = ? AND status = ?", uploadID, media.UploadStatusUploading).
		Update("status", media.UploadStatusMerging)
	if res.Error != nil {
		return m, res.Error
	}
	if res.RowsAffected == 0 {
		return m, errors.New("上传不在可合并状态(可能已在合并或已完成)")
	}

	fail := func(e error) (media.FileUploadAndDownload, error) {
		global.GVA_DB.Model(&media.MediaUpload{}).Where("id = ?", uploadID).Update("status", media.UploadStatusFailed)
		return m, e
	}

	// 校验分片齐全
	var cnt int64
	global.GVA_DB.Model(&media.MediaUploadChunk{}).Where("upload_id = ?", uploadID).Count(&cnt)
	if int(cnt) != up.ChunkTotal {
		return fail(fmt.Errorf("分片不全: %d/%d", cnt, up.ChunkTotal))
	}

	// 合并到临时成品
	merged := filepath.Join(upload.ChunkDir(uploadID), "merged.bin")
	gotMd5, err := upload.MergeChunks(uploadID, up.ChunkTotal, merged)
	if err != nil {
		return fail(err)
	}
	if gotMd5 != up.FileHash {
		return fail(errors.New("整文件校验失败"))
	}

	// 经 OSS 接口推到配置存储
	fh, cleanup, err := upload.BuildFileHeader(merged, "file", up.FileName)
	if err != nil {
		return fail(err)
	}
	defer cleanup()
	oss := upload.NewOss()
	url, key, err := oss.UploadFile(fh)
	if err != nil {
		return fail(err)
	}

	// 登记媒体库
	ext := ""
	if i := strings.LastIndex(up.FileName, "."); i >= 0 {
		ext = strings.TrimPrefix(up.FileName[i:], ".")
	}
	m = media.FileUploadAndDownload{Name: up.FileName, Url: url, Tag: ext, Key: key}
	if err = global.GVA_DB.Create(&m).Error; err != nil {
		return fail(err)
	}

	// 回填会话并清理
	global.GVA_DB.Model(&media.MediaUpload{}).Where("id = ?", uploadID).
		Updates(map[string]interface{}{"status": media.UploadStatusCompleted, "media_id": m.ID, "storage_key": key})
	global.GVA_DB.Where("upload_id = ?", uploadID).Delete(&media.MediaUploadChunk{})
	_ = upload.RemoveUploadDir(uploadID)
	return m, nil
}

func (s *MediaUploadService) Cancel(userID, uploadID uint) error {
	var up media.MediaUpload
	if err := global.GVA_DB.First(&up, uploadID).Error; err != nil {
		return errors.New("上传会话不存在")
	}
	if up.UserID != userID {
		return errors.New("无权操作该上传")
	}
	global.GVA_DB.Where("upload_id = ?", uploadID).Delete(&media.MediaUploadChunk{})
	_ = upload.RemoveUploadDir(uploadID)
	return global.GVA_DB.Delete(&media.MediaUpload{}, uploadID).Error
}

func (s *MediaUploadService) CleanupStale(ttlHours int) error {
	if ttlHours <= 0 {
		ttlHours = 24
	}
	deadline := time.Now().Add(-time.Duration(ttlHours) * time.Hour)
	var stale []media.MediaUpload
	global.GVA_DB.Where("status = ? AND updated_at < ?", media.UploadStatusUploading, deadline).Find(&stale)
	for _, up := range stale {
		global.GVA_DB.Where("upload_id = ?", up.ID).Delete(&media.MediaUploadChunk{})
		_ = upload.RemoveUploadDir(up.ID)
		global.GVA_DB.Delete(&media.MediaUpload{}, up.ID)
	}
	return nil
}
