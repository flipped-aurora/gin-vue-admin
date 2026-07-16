package upload

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"

	"github.com/tencentyun/cos-go-sdk-v5"
)

type TencentCOS struct{}

// UploadFile upload file to COS
func (*TencentCOS) UploadFile(ctx context.Context, file *multipart.FileHeader) (string, string, error) {
	client := NewClient()
	f, openError := file.Open()
	if openError != nil {
		logger.WithCtx(ctx).Mod("upload").Err(openError).Error("function file.Open() failed")
		return "", "", errors.New("function file.Open() failed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)

	_, err := client.Object.Put(context.Background(), global.GVA_CONFIG.TencentCOS.PathPrefix+"/"+fileKey, f, nil)
	if err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("function client.Object.Put() failed")
		return "", "", errors.New("function client.Object.Put() failed, err:" + err.Error())
	}
	return global.GVA_CONFIG.TencentCOS.BaseURL + "/" + global.GVA_CONFIG.TencentCOS.PathPrefix + "/" + fileKey, fileKey, nil
}

// DeleteFile delete file form COS
func (*TencentCOS) DeleteFile(ctx context.Context, key string) error {
	client := NewClient()
	name := global.GVA_CONFIG.TencentCOS.PathPrefix + "/" + key
	_, err := client.Object.Delete(context.Background(), name)
	if err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("function bucketManager.Delete() failed")
		return errors.New("function bucketManager.Delete() failed, err:" + err.Error())
	}
	return nil
}

// Exists 通过 Object.Head 检查对象是否存在；404 统一降级为 (false, nil)。
func (*TencentCOS) Exists(ctx context.Context, key string) (bool, error) {
	client := NewClient()
	name := global.GVA_CONFIG.TencentCOS.PathPrefix + "/" + key
	_, err := client.Object.Head(ctx, name, nil)
	if err != nil {
		if cos.IsNotFoundError(err) {
			return false, nil
		}
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("function client.Object.Head() failed")
		return false, errors.New("function client.Object.Head() failed, err:" + err.Error())
	}
	return true, nil
}

// DeleteFiles 批量删除：Object.DeleteMulti 一次提交，返回 Errors 逐个失败项。
// PathPrefix 拼接方式与 DeleteFile 一致。
func (*TencentCOS) DeleteFiles(ctx context.Context, keys []string) (failed []DeleteFailure, err error) {
	client := NewClient()
	objects := make([]cos.Object, 0, len(keys))
	for _, k := range keys {
		objects = append(objects, cos.Object{Key: global.GVA_CONFIG.TencentCOS.PathPrefix + "/" + k})
	}

	res, _, err := client.Object.DeleteMulti(ctx, &cos.ObjectDeleteMultiOptions{
		Objects: objects,
		Quiet:   true,
	})
	if err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("function client.Object.DeleteMulti() failed")
		return nil, errors.New("function client.Object.DeleteMulti() failed, err:" + err.Error())
	}

	if res != nil {
		for _, e := range res.Errors {
			failed = append(failed, DeleteFailure{
				Key: e.Key,
				Err: fmt.Errorf("cos delete %s failed: code=%s, message=%s", e.Key, e.Code, e.Message),
			})
		}
	}
	return failed, nil
}

// ListFiles 按前缀列举对象，marker 分页：Marker=cursor，NextMarker→nextCursor，IsTruncated→hasMore。
func (*TencentCOS) ListFiles(ctx context.Context, prefix, cursor string, limit int) (files []FileInfo, nextCursor string, hasMore bool, err error) {
	if limit <= 0 {
		limit = 100
	}

	client := NewClient()
	res, _, err := client.Bucket.Get(ctx, &cos.BucketGetOptions{
		Prefix:  prefix,
		Marker:  cursor,
		MaxKeys: limit,
	})
	if err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("function client.Bucket.Get() failed")
		return nil, "", false, errors.New("function client.Bucket.Get() failed, err:" + err.Error())
	}

	if res != nil {
		for _, c := range res.Contents {
			files = append(files, FileInfo{
				Key:          c.Key,
				Size:         c.Size,
				LastModified: parseCOSLastModified(c.LastModified),
			})
		}
		hasMore = res.IsTruncated
		if hasMore && res.NextMarker != "" {
			nextCursor = res.NextMarker
		} else if hasMore && len(files) > 0 {
			// COS 未返回 NextMarker 时回退用最后一条 key 作为下次 marker
			nextCursor = files[len(files)-1].Key
		}
	}
	return files, nextCursor, hasMore, nil
}

// parseCOSLastModified 解析 COS 返回的时间字符串（RFC3339 / ISO8601），失败返回零值。
func parseCOSLastModified(s string) time.Time {
	if s == "" {
		return time.Time{}
	}
	if t, perr := time.Parse(time.RFC3339, s); perr == nil {
		return t
	}
	return time.Time{}
}

// NewClient init COS client
func NewClient() *cos.Client {
	urlStr, _ := url.Parse("https://" + global.GVA_CONFIG.TencentCOS.Bucket + ".cos." + global.GVA_CONFIG.TencentCOS.Region + ".myqcloud.com")
	baseURL := &cos.BaseURL{BucketURL: urlStr}
	client := cos.NewClient(baseURL, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  global.GVA_CONFIG.TencentCOS.SecretID,
			SecretKey: global.GVA_CONFIG.TencentCOS.SecretKey,
		},
	})
	return client
}
