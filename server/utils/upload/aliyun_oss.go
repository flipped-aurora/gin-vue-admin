package upload

import (
	"context"
	"errors"
	"mime/multipart"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
)

type AliyunOSS struct{}

func (*AliyunOSS) UploadFile(ctx context.Context, file *multipart.FileHeader) (string, string, error) {
	bucket, err := NewBucket()
	if err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("function AliyunOSS.NewBucket() Failed")
		return "", "", errors.New("function AliyunOSS.NewBucket() Failed, err:" + err.Error())
	}

	// 读取本地文件。
	f, openError := file.Open()
	if openError != nil {
		logger.WithCtx(ctx).Mod("upload").Err(openError).Error("function file.Open() Failed")
		return "", "", errors.New("function file.Open() Failed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭
	// 上传阿里云路径 文件名格式 自己可以改 建议保证唯一性
	// yunFileTmpPath := filepath.Join("uploads", time.Now().Format("2006-01-02")) + "/" + file.Filename
	yunFileTmpPath := global.GVA_CONFIG.AliyunOSS.BasePath + "/" + "uploads" + "/" + time.Now().Format("2006-01-02") + "/" + file.Filename

	// 上传文件流。
	err = bucket.PutObject(yunFileTmpPath, f)
	if err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("function formUploader.Put() Failed")
		return "", "", errors.New("function formUploader.Put() Failed, err:" + err.Error())
	}

	return global.GVA_CONFIG.AliyunOSS.BucketUrl + "/" + yunFileTmpPath, yunFileTmpPath, nil
}

func (*AliyunOSS) DeleteFile(ctx context.Context, key string) error {
	bucket, err := NewBucket()
	if err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("function AliyunOSS.NewBucket() Failed")
		return errors.New("function AliyunOSS.NewBucket() Failed, err:" + err.Error())
	}

	// 删除单个文件。objectName表示删除OSS文件时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	// 如需删除文件夹，请将objectName设置为对应的文件夹名称。如果文件夹非空，则需要将文件夹下的所有object删除后才能删除该文件夹。
	err = bucket.DeleteObject(key)
	if err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("function bucketManager.Delete() failed")
		return errors.New("function bucketManager.Delete() failed, err:" + err.Error())
	}

	return nil
}

func NewBucket() (*oss.Bucket, error) {
	// 创建OSSClient实例。
	client, err := oss.New(global.GVA_CONFIG.AliyunOSS.Endpoint, global.GVA_CONFIG.AliyunOSS.AccessKeyId, global.GVA_CONFIG.AliyunOSS.AccessKeySecret)
	if err != nil {
		return nil, err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(global.GVA_CONFIG.AliyunOSS.BucketName)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}

// Exists 检查对象是否存在，"不存在"统一降级为 (false, nil)。
func (*AliyunOSS) Exists(ctx context.Context, key string) (bool, error) {
	bucket, err := NewBucket()
	if err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("function AliyunOSS.NewBucket() Failed")
		return false, errors.New("function AliyunOSS.NewBucket() Failed, err:" + err.Error())
	}

	exist, err := bucket.IsObjectExist(key)
	if err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("function bucket.IsObjectExist() failed")
		return false, errors.New("function bucket.IsObjectExist() failed, err:" + err.Error())
	}
	return exist, nil
}

// DeleteFiles 批量删除对象，对比入参与已删除列表，未删的收集为 DeleteFailure。
func (*AliyunOSS) DeleteFiles(ctx context.Context, keys []string) ([]DeleteFailure, error) {
	if len(keys) == 0 {
		return nil, nil
	}

	bucket, err := NewBucket()
	if err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("function AliyunOSS.NewBucket() Failed")
		return nil, errors.New("function AliyunOSS.NewBucket() Failed, err:" + err.Error())
	}

	result, err := bucket.DeleteObjects(keys)
	if err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("function bucket.DeleteObjects() failed")
		return nil, errors.New("function bucket.DeleteObjects() failed, err:" + err.Error())
	}

	// 默认非 quiet 模式，OSS 会返回已删除的 key 列表，与之对比找出未删项。
	deleted := make(map[string]struct{}, len(result.DeletedObjects))
	for _, k := range result.DeletedObjects {
		deleted[k] = struct{}{}
	}

	failed := make([]DeleteFailure, 0)
	for _, key := range keys {
		if _, ok := deleted[key]; !ok {
			failed = append(failed, DeleteFailure{Key: key, Err: errors.New("对象未被删除")})
		}
	}
	return failed, nil
}

// ListFiles 按前缀列举对象，cursor 映射为 OSS 的 marker。
func (*AliyunOSS) ListFiles(ctx context.Context, prefix, cursor string, limit int) ([]FileInfo, string, bool, error) {
	if limit <= 0 {
		limit = 100
	}

	bucket, err := NewBucket()
	if err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("function AliyunOSS.NewBucket() Failed")
		return nil, "", false, errors.New("function AliyunOSS.NewBucket() Failed, err:" + err.Error())
	}

	result, err := bucket.ListObjects(oss.Prefix(prefix), oss.Marker(cursor), oss.MaxKeys(limit))
	if err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("function bucket.ListObjects() failed")
		return nil, "", false, errors.New("function bucket.ListObjects() failed, err:" + err.Error())
	}

	files := make([]FileInfo, 0, len(result.Objects))
	for _, object := range result.Objects {
		files = append(files, FileInfo{
			Key:          object.Key,
			Size:         object.Size,
			LastModified: object.LastModified,
		})
	}

	nextCursor := ""
	if result.IsTruncated {
		nextCursor = result.NextMarker
	}
	return files, nextCursor, result.IsTruncated, nil
}
