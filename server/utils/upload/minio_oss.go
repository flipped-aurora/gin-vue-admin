package upload

import (
	"bytes"
	"context"
	"errors"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// Minio 与其他 OSS 实现一致，采用空结构体：每次上传/删除时从 global.GVA_CONFIG 现读配置现建 client，
// 不缓存全局单例，支持配置热重载，避免并发首次初始化竞争。
type Minio struct{}

// newMinioClient 每次调用时根据当前 config 新建 minio client。
// endpoint 兼容带 http:// / https:// 前缀的写法（自动剥离，协议由 use-ssl 决定），对齐 aws-s3 的 endpoint 容错行为。
func newMinioClient() (*minio.Client, error) {
	cfg := global.GVA_CONFIG.Minio
	endpoint := strings.TrimPrefix(strings.TrimPrefix(cfg.Endpoint, "https://"), "http://")

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKeyId, cfg.AccessKeySecret, ""),
		Secure: cfg.UseSSL, // Set to true if using https
	})
	if err != nil {
		return nil, err
	}
	return minioClient, nil
}

func (m *Minio) UploadFile(ctx context.Context, file *multipart.FileHeader) (filePathres, key string, uploadErr error) {
	client, err := newMinioClient()
	if err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("minio client 初始化失败")
		return "", "", errors.New("minio client 初始化失败, err:" + err.Error())
	}

	f, openError := file.Open()
	// mutipart.File to os.File
	if openError != nil {
		logger.WithCtx(ctx).Mod("upload").Err(openError).Error("function file.Open() Failed")
		return "", "", errors.New("function file.Open() Failed, err:" + openError.Error())
	}

	filecontent := bytes.Buffer{}
	_, err = io.Copy(&filecontent, f)
	if err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("读取文件失败")
		return "", "", errors.New("读取文件失败, err:" + err.Error())
	}
	f.Close() // 创建文件 defer 关闭

	// 对文件名进行加密存储
	ext := filepath.Ext(file.Filename)
	filename := utils.MD5V([]byte(strings.TrimSuffix(file.Filename, ext))) + ext
	if global.GVA_CONFIG.Minio.BasePath == "" {
		filePathres = "uploads" + "/" + time.Now().Format("2006-01-02") + "/" + filename
	} else {
		filePathres = global.GVA_CONFIG.Minio.BasePath + "/" + time.Now().Format("2006-01-02") + "/" + filename
	}

	// 根据文件扩展名检测 MIME 类型
	contentType := mime.TypeByExtension(ext)
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	// 设置超时10分钟
	putCtx, cancel := context.WithTimeout(ctx, time.Minute*10)
	defer cancel()

	// Upload the file with PutObject   大文件自动切换为分片上传
	info, err := client.PutObject(putCtx, global.GVA_CONFIG.Minio.BucketName, filePathres, &filecontent, file.Size, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("上传文件到minio失败")
		return "", "", errors.New("上传文件到minio失败, err:" + err.Error())
	}
	return global.GVA_CONFIG.Minio.BucketUrl + "/" + info.Key, filePathres, nil
}

func (m *Minio) DeleteFile(ctx context.Context, key string) error {
	client, err := newMinioClient()
	if err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("minio client 初始化失败")
		return errors.New("minio client 初始化失败, err:" + err.Error())
	}

	delCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	// Delete the object from MinIO
	err = client.RemoveObject(delCtx, global.GVA_CONFIG.Minio.BucketName, key, minio.RemoveObjectOptions{})
	return err
}

// Exists 通过 StatObject 检查对象是否存在；404 统一降级为 (false, nil)。
func (m *Minio) Exists(ctx context.Context, key string) (bool, error) {
	client, err := newMinioClient()
	if err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("minio client 初始化失败")
		return false, errors.New("minio client 初始化失败, err:" + err.Error())
	}

	statCtx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	_, err = client.StatObject(statCtx, global.GVA_CONFIG.Minio.BucketName, key, minio.StatObjectOptions{})
	if err != nil {
		// minio channel/响应错误统一用 ToErrorResponse 还原，不存在视为正常状态
		errResp := minio.ToErrorResponse(err)
		if errResp.StatusCode == http.StatusNotFound {
			return false, nil
		}
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("minio StatObject 失败")
		return false, err
	}
	return true, nil
}

// DeleteFiles 批量删除：minio RemoveObjects 走 channel 模型，range errCh 收集逐个失败项。
func (m *Minio) DeleteFiles(ctx context.Context, keys []string) (failed []DeleteFailure, err error) {
	client, err := newMinioClient()
	if err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("minio client 初始化失败")
		return nil, errors.New("minio client 初始化失败, err:" + err.Error())
	}

	bucket := global.GVA_CONFIG.Minio.BucketName
	objCh := make(chan minio.ObjectInfo)
	go func() {
		defer close(objCh)
		for _, key := range keys {
			objCh <- minio.ObjectInfo{Key: key}
		}
	}()

	errCh := client.RemoveObjects(ctx, bucket, objCh, minio.RemoveObjectsOptions{})
	for e := range errCh {
		failed = append(failed, DeleteFailure{Key: e.ObjectName, Err: e.Err})
	}
	return failed, nil
}

// ListFiles 按前缀列举对象。minio 的 ListObjects 是 channel 迭代器型，不支持外部游标，
// 首版策略：cursor 非空时映射到 StartAfter 做起点；limit<=0 默认 100；读够 limit 且 channel 未关闭则 hasMore=true。
func (m *Minio) ListFiles(ctx context.Context, prefix, cursor string, limit int) (files []FileInfo, nextCursor string, hasMore bool, err error) {
	if limit <= 0 {
		limit = 100
	}

	client, err := newMinioClient()
	if err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("minio client 初始化失败")
		return nil, "", false, errors.New("minio client 初始化失败, err:" + err.Error())
	}

	bucket := global.GVA_CONFIG.Minio.BucketName
	objCh := client.ListObjects(ctx, bucket, minio.ListObjectsOptions{
		Prefix:     prefix,
		Recursive:  true,
		MaxKeys:    limit,
		StartAfter: cursor, // cursor 为空即从头列举
	})

	count := 0
	for info := range objCh {
		if info.Err != nil {
			err = info.Err
			logger.WithCtx(ctx).Mod("upload").Err(err).Error("minio ListObjects 迭代失败")
			return files, "", false, err
		}
		files = append(files, FileInfo{
			Key:          info.Key,
			Size:         info.Size,
			LastModified: info.LastModified,
			ContentType:  info.ContentType,
		})
		count++
		if count >= limit {
			break
		}
	}

	// 读够 limit：以最后一条 key 作为下次的 StartAfter 游标；
	// 再阻塞读一次判定是否还有剩余，读到的额外项丢弃，并排空 channel 避免泄漏 ListObjects goroutine。
	if count >= limit && len(files) > 0 {
		nextCursor = files[len(files)-1].Key
		_, hasMore = <-objCh
		for range objCh { // 排空剩余
		}
	}
	return files, nextCursor, hasMore, nil
}
