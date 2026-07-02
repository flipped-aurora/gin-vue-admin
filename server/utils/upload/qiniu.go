package upload

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type Qiniu struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [ccfish86](https://github.com/ccfish86)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@object: *Qiniu
//@function: UploadFile
//@description: 上传文件
//@param: file *multipart.FileHeader
//@return: string, string, error

func (*Qiniu) UploadFile(ctx context.Context, file *multipart.FileHeader) (string, string, error) {
	putPolicy := storage.PutPolicy{Scope: global.GVA_CONFIG.Qiniu.Bucket}
	mac := qbox.NewMac(global.GVA_CONFIG.Qiniu.AccessKey, global.GVA_CONFIG.Qiniu.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := qiniuConfig()
	formUploader := storage.NewFormUploader(cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{Params: map[string]string{"x:name": "github logo"}}

	f, openError := file.Open()
	if openError != nil {
		logger.WithCtx(ctx).Mod("upload").Err(openError).Error("function file.Open() failed")

		return "", "", errors.New("function file.Open() failed, err:" + openError.Error())
	}
	defer f.Close()                                                  // 创建文件 defer 关闭
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename) // 文件名格式 自己可以改 建议保证唯一性
	putErr := formUploader.Put(context.Background(), &ret, upToken, fileKey, f, file.Size, &putExtra)
	if putErr != nil {
		logger.WithCtx(ctx).Mod("upload").Err(putErr).Error("function formUploader.Put() failed")
		return "", "", errors.New("function formUploader.Put() failed, err:" + putErr.Error())
	}
	return global.GVA_CONFIG.Qiniu.ImgPath + "/" + ret.Key, ret.Key, nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [ccfish86](https://github.com/ccfish86)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@object: *Qiniu
//@function: DeleteFile
//@description: 删除文件
//@param: key string
//@return: error

func (*Qiniu) DeleteFile(ctx context.Context, key string) error {
	bucketManager := newBucketManager()
	if err := bucketManager.Delete(global.GVA_CONFIG.Qiniu.Bucket, key); err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("function bucketManager.Delete() failed")
		return errors.New("function bucketManager.Delete() failed, err:" + err.Error())
	}
	return nil
}

// newBucketManager 创建七牛 BucketManager 实例。
func newBucketManager() *storage.BucketManager {
	mac := qbox.NewMac(global.GVA_CONFIG.Qiniu.AccessKey, global.GVA_CONFIG.Qiniu.SecretKey)
	cfg := qiniuConfig()
	return storage.NewBucketManager(mac, cfg)
}

// Exists 检查对象是否存在，七牛 code 612（或 "no such file"）视为不存在。
func (*Qiniu) Exists(ctx context.Context, key string) (bool, error) {
	bucketManager := newBucketManager()
	if _, err := bucketManager.Stat(global.GVA_CONFIG.Qiniu.Bucket, key); err != nil {
		msg := err.Error()
		// 七牛 612 表示资源不存在
		if strings.Contains(msg, "612") || strings.Contains(msg, "no such file") || strings.Contains(msg, "no such file or directory") {
			return false, nil
		}
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("function bucketManager.Stat() failed")
		return false, errors.New("function bucketManager.Stat() failed, err:" + err.Error())
	}
	return true, nil
}

// DeleteFiles 批量删除对象，通过通用 Batch 接口，逐项检查 code（200/204 成功）。
func (*Qiniu) DeleteFiles(ctx context.Context, keys []string) ([]DeleteFailure, error) {
	if len(keys) == 0 {
		return nil, nil
	}

	bucket := global.GVA_CONFIG.Qiniu.Bucket
	operations := make([]string, 0, len(keys))
	for _, key := range keys {
		operations = append(operations, storage.URIDelete(bucket, key))
	}

	bucketManager := newBucketManager()
	ret, err := bucketManager.Batch(operations)
	if err != nil {
		// 部分 key 失败时 Batch 仍可能返回结果，这里仅当完全没有结果时视为致命错误。
		if len(ret) == 0 {
			logger.WithCtx(ctx).Mod("upload").Err(err).Error("function bucketManager.Batch() failed")
			return nil, errors.New("function bucketManager.Batch() failed, err:" + err.Error())
		}
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("function bucketManager.Batch() partial failure")
	}

	failed := make([]DeleteFailure, 0)
	for i, op := range ret {
		// 200 / 204 视为删除成功
		if op.Code != http.StatusOK && op.Code != http.StatusNoContent {
			var key string
			if i < len(keys) {
				key = keys[i]
			}
			failed = append(failed, DeleteFailure{Key: key, Err: fmt.Errorf("delete failed, code: %d", op.Code)})
		}
	}
	return failed, nil
}

// ListFiles 按前缀列举对象，cursor 映射为七牛的 marker。
func (*Qiniu) ListFiles(ctx context.Context, prefix, cursor string, limit int) ([]FileInfo, string, bool, error) {
	if limit <= 0 {
		limit = 100
	}

	bucketManager := newBucketManager()
	entries, _, nextMarker, hasNext, err := bucketManager.ListFiles(global.GVA_CONFIG.Qiniu.Bucket, prefix, "", cursor, limit)
	if err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("function bucketManager.ListFiles() failed")
		return nil, "", false, errors.New("function bucketManager.ListFiles() failed, err:" + err.Error())
	}

	files := make([]FileInfo, 0, len(entries))
	for _, entry := range entries {
		// PutTime 单位为 100 纳秒，除以 1e7 得到秒级 Unix 时间戳
		files = append(files, FileInfo{
			Key:          entry.Key,
			Size:         entry.Fsize,
			LastModified: time.Unix(0, entry.PutTime*100),
			ContentType:  entry.MimeType,
		})
	}

	nextCursor := ""
	if hasNext {
		nextCursor = nextMarker
	}
	return files, nextCursor, hasNext, nil
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@object: *Qiniu
//@function: qiniuConfig
//@description: 根据配置文件进行返回七牛云的配置
//@return: *storage.Config

func qiniuConfig() *storage.Config {
	cfg := storage.Config{
		UseHTTPS:      global.GVA_CONFIG.Qiniu.UseHTTPS,
		UseCdnDomains: global.GVA_CONFIG.Qiniu.UseCdnDomains,
	}
	switch global.GVA_CONFIG.Qiniu.Zone { // 根据配置文件进行初始化空间对应的机房
	case "ZoneHuadong":
		cfg.Zone = &storage.ZoneHuadong
	case "ZoneHuabei":
		cfg.Zone = &storage.ZoneHuabei
	case "ZoneHuanan":
		cfg.Zone = &storage.ZoneHuanan
	case "ZoneBeimei":
		cfg.Zone = &storage.ZoneBeimei
	case "ZoneXinjiapo":
		cfg.Zone = &storage.ZoneXinjiapo
	}
	return &cfg
}
