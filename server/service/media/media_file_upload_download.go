package media

import (
	"context"
	"errors"
	"mime/multipart"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/media"
	"github.com/flipped-aurora/gin-vue-admin/server/model/media/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/upload"
	"gorm.io/gorm"
)

type FileUploadAndDownloadService struct{}

var FileUploadAndDownloadServiceApp = new(FileUploadAndDownloadService)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Upload
//@description: 创建文件上传记录
//@param: file model.FileUploadAndDownload
//@return: error

func (e *FileUploadAndDownloadService) Upload(ctx context.Context, file media.FileUploadAndDownload) error {
	return global.GVA_DB.WithContext(ctx).Create(&file).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: FindFile
//@description: 查询文件记录
//@param: id uint
//@return: model.FileUploadAndDownload, error

func (e *FileUploadAndDownloadService) FindFile(ctx context.Context, id uint) (media.FileUploadAndDownload, error) {
	var file media.FileUploadAndDownload
	err := global.GVA_DB.WithContext(ctx).Where("id = ?", id).First(&file).Error
	return file, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFile
//@description: 删除文件记录
//@param: file model.FileUploadAndDownload
//@return: err error

func (e *FileUploadAndDownloadService) DeleteFile(ctx context.Context, file media.FileUploadAndDownload) (err error) {
	var fileFromDb media.FileUploadAndDownload
	fileFromDb, err = e.FindFile(ctx, file.ID)
	if err != nil {
		return
	}
	// 引用计数：同一 key 可能被秒传复制出多条记录，仅当最后一条时才真正删 OSS 对象，
	// 避免删一个副本导致其他副本文件丢失。
	var count int64
	if err = global.GVA_DB.WithContext(ctx).Model(&media.FileUploadAndDownload{}).Where("key = ?", fileFromDb.Key).Count(&count).Error; err != nil {
		return err
	}
	if count <= 1 {
		oss := upload.NewOss()
		if err = oss.DeleteFile(ctx, fileFromDb.Key); err != nil {
			return errors.New("文件删除失败")
		}
	}
	err = global.GVA_DB.WithContext(ctx).Where("id = ?", file.ID).Unscoped().Delete(&file).Error
	return err
}

// DeleteFiles 批量删除文件记录。逐条复用 DeleteFile（含引用计数逻辑），
// 任一失败收集到 failedIDs 返回，不中断其余删除。
func (e *FileUploadAndDownloadService) DeleteFiles(ctx context.Context, ids []uint) (failedIDs []uint, err error) {
	for _, id := range ids {
		if delErr := e.DeleteFile(ctx, media.FileUploadAndDownload{GVA_MODEL: global.GVA_MODEL{ID: id}}); delErr != nil {
			failedIDs = append(failedIDs, id)
		}
	}
	return failedIDs, nil
}

// GetFileDetail 获取单个文件详情
func (e *FileUploadAndDownloadService) GetFileDetail(ctx context.Context, id uint) (media.FileUploadAndDownload, error) {
	return e.FindFile(ctx, id)
}

// ListOssFiles 直接列举 OSS 存储桶内的真实文件（不查 DB）
func (e *FileUploadAndDownloadService) ListOssFiles(ctx context.Context, prefix, cursor string, limit int) (files []upload.FileInfo, nextCursor string, hasMore bool, err error) {
	oss := upload.NewOss()
	return oss.ListFiles(ctx, prefix, cursor, limit)
}

// EditFileName 编辑文件名或者备注
func (e *FileUploadAndDownloadService) EditFileName(ctx context.Context, file media.FileUploadAndDownload) (err error) {
	var fileFromDb media.FileUploadAndDownload
	return global.GVA_DB.WithContext(ctx).Where("id = ?", file.ID).First(&fileFromDb).Update("name", file.Name).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFileRecordInfoList
//@description: 分页获取数据
//@param: info request.AttachmentCategorySearch
//@return: list interface{}, total int64, err error

func (e *FileUploadAndDownloadService) GetFileRecordInfoList(ctx context.Context, info request.AttachmentCategorySearch) (list []media.FileUploadAndDownload, total int64, err error) {
	limit, offset := info.LimitOffset()
	db := global.GVA_DB.WithContext(ctx).Model(&media.FileUploadAndDownload{})

	if len(info.Keyword) > 0 {
		db = db.Where("name LIKE ?", "%"+info.Keyword+"%")
	}

	if info.ClassId > 0 {
		db = db.Where("class_id = ?", info.ClassId)
	}

	if info.Tag != "" {
		db = db.Where("tag = ?", info.Tag)
	}

	if info.UserID > 0 {
		db = db.Where("user_id = ?", info.UserID)
	}

	if !info.StartCreatedAt.IsZero() {
		db = db.Where("created_at >= ?", info.StartCreatedAt)
	}
	if !info.EndCreatedAt.IsZero() {
		db = db.Where("created_at < ?", info.EndCreatedAt)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	// 排序：允许指定字段，默认 id desc；白名单防注入
	orderKey := info.OrderKey
	allowedOrder := map[string]bool{"id": true, "size": true, "created_at": true, "name": true}
	if !allowedOrder[orderKey] {
		orderKey = "id"
	}
	order := orderKey
	if info.Desc {
		order += " desc"
	}
	err = db.Limit(limit).Offset(offset).Order(order).Find(&list).Error
	return list, total, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UploadFile
//@description: 根据配置文件判断是文件上传到本地或者七牛云
//@param: header *multipart.FileHeader, noSave string
//@return: file model.FileUploadAndDownload, err error

func (e *FileUploadAndDownloadService) UploadFile(ctx context.Context, header *multipart.FileHeader, noSave string, classId int, userID uint) (file media.FileUploadAndDownload, err error) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(ctx, header)
	if uploadErr != nil {
		return file, uploadErr
	}
	s := strings.Split(header.Filename, ".")
	// 计算文件 MD5（读取文件内容），用于去重校验
	md5sum, _ := utils.MD5VFile(header)
	f := media.FileUploadAndDownload{
		Url:    filePath,
		Name:   header.Filename,
		ClassId: classId,
		Tag:    s[len(s)-1],
		Key:    key,
		Size:   header.Size,
		Mime:   utils.DetectMIME(header),
		Md5:    md5sum,
		UserID: userID,
	}
	if noSave == "0" {
		// 检查是否已存在相同key的记录
		var existingFile media.FileUploadAndDownload
		err = global.GVA_DB.WithContext(ctx).Where(&media.FileUploadAndDownload{Key: key}).First(&existingFile).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return f, e.Upload(ctx, f)
		}
		return f, err
	}
	return f, nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ImportURL
//@description: 导入URL
//@param: file model.FileUploadAndDownload
//@return: error

func (e *FileUploadAndDownloadService) ImportURL(ctx context.Context, file *[]media.FileUploadAndDownload) error {
	return global.GVA_DB.WithContext(ctx).Create(&file).Error
}
