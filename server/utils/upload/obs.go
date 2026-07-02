package upload

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	pkgerrors "github.com/pkg/errors"
)

var HuaWeiObs = new(Obs)

type Obs struct{}

func NewHuaWeiObsClient() (client *obs.ObsClient, err error) {
	return obs.New(global.GVA_CONFIG.HuaWeiObs.AccessKey, global.GVA_CONFIG.HuaWeiObs.SecretKey, global.GVA_CONFIG.HuaWeiObs.Endpoint)
}

func (o *Obs) UploadFile(ctx context.Context, file *multipart.FileHeader) (string, string, error) {
	// var open multipart.File
	open, err := file.Open()
	if err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("function file.Open() failed")
		return "", "", err
	}
	defer open.Close()
	// 加唯一时间前缀，避免同名文件覆盖（对齐 aws_s3 / cloudflare_r2 的命名风格）
	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	input := &obs.PutObjectInput{
		PutObjectBasicInput: obs.PutObjectBasicInput{
			ObjectOperationInput: obs.ObjectOperationInput{
				Bucket: global.GVA_CONFIG.HuaWeiObs.Bucket,
				Key:    filename,
			},
			HttpHeader: obs.HttpHeader{
				ContentType: file.Header.Get("content-type"),
			},
		},
		Body: open,
	}

	var client *obs.ObsClient
	client, err = NewHuaWeiObsClient()
	if err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("获取华为对象存储对象失败")
		return "", "", pkgerrors.Wrap(err, "获取华为对象存储对象失败!")
	}

	_, err = client.PutObject(input)
	if err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("文件上传失败")
		return "", "", pkgerrors.Wrap(err, "文件上传失败!")
	}
	filepath := global.GVA_CONFIG.HuaWeiObs.Path + "/" + filename
	return filepath, filename, err
}

func (o *Obs) DeleteFile(ctx context.Context, key string) error {
	client, err := NewHuaWeiObsClient()
	if err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("获取华为对象存储对象失败")
		return pkgerrors.Wrap(err, "获取华为对象存储对象失败!")
	}
	input := &obs.DeleteObjectInput{
		Bucket: global.GVA_CONFIG.HuaWeiObs.Bucket,
		Key:    key,
	}
	var output *obs.DeleteObjectOutput
	output, err = client.DeleteObject(input)
	if err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("删除对象失败")
		return pkgerrors.Wrapf(err, "删除对象(%s)失败!, output: %v", key, output)
	}
	return nil
}

// Exists 通过 GetObjectMetadata 检查对象是否存在；404 统一降级为 (false, nil)。
func (o *Obs) Exists(ctx context.Context, key string) (bool, error) {
	client, err := NewHuaWeiObsClient()
	if err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("获取华为对象存储对象失败")
		return false, pkgerrors.Wrap(err, "获取华为对象存储对象失败!")
	}

	_, err = client.GetObjectMetadata(&obs.GetObjectMetadataInput{
		Bucket: global.GVA_CONFIG.HuaWeiObs.Bucket,
		Key:    key,
	})
	if err != nil {
		// OBS 错误统一是 ObsError（内嵌 BaseModel.StatusCode），用 errors.As 还原
		var obsErr obs.ObsError
		if errors.As(err, &obsErr) && obsErr.StatusCode == http.StatusNotFound {
			return false, nil
		}
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("obs GetObjectMetadata 失败")
		return false, err
	}
	return true, nil
}

// DeleteFiles 批量删除：OBS DeleteObjects 一次提交，返回 Errors []obs.Error 逐个失败项。
func (o *Obs) DeleteFiles(ctx context.Context, keys []string) (failed []DeleteFailure, err error) {
	client, err := NewHuaWeiObsClient()
	if err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("获取华为对象存储对象失败")
		return nil, pkgerrors.Wrap(err, "获取华为对象存储对象失败!")
	}

	objects := make([]obs.ObjectToDelete, 0, len(keys))
	for _, k := range keys {
		objects = append(objects, obs.ObjectToDelete{Key: k})
	}

	output, err := client.DeleteObjects(&obs.DeleteObjectsInput{
		Bucket:  global.GVA_CONFIG.HuaWeiObs.Bucket,
		Objects: objects,
		Quiet:   true,
	})
	if err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("obs DeleteObjects 失败")
		return nil, err
	}

	if output != nil {
		for _, e := range output.Errors {
			failed = append(failed, DeleteFailure{
				Key: e.Key,
				Err: fmt.Errorf("obs delete %s failed: code=%s, message=%s", e.Key, e.Code, e.Message),
			})
		}
	}
	return failed, nil
}

// ListFiles 按前缀列举对象，marker 分页：Marker=cursor，NextMarker→nextCursor，IsTruncated→hasMore。
func (o *Obs) ListFiles(ctx context.Context, prefix, cursor string, limit int) (files []FileInfo, nextCursor string, hasMore bool, err error) {
	if limit <= 0 {
		limit = 100
	}

	client, err := NewHuaWeiObsClient()
	if err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("获取华为对象存储对象失败")
		return nil, "", false, pkgerrors.Wrap(err, "获取华为对象存储对象失败!")
	}

	output, err := client.ListObjects(&obs.ListObjectsInput{
		ListObjsInput: obs.ListObjsInput{
			Prefix:  prefix,
			MaxKeys: limit,
		},
		Bucket: global.GVA_CONFIG.HuaWeiObs.Bucket,
		Marker: cursor,
	})
	if err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("obs ListObjects 失败")
		return nil, "", false, err
	}

	if output != nil {
		for _, c := range output.Contents {
			files = append(files, FileInfo{
				Key:          c.Key,
				Size:         c.Size,
				LastModified: c.LastModified,
			})
		}
		hasMore = output.IsTruncated
		if hasMore && output.NextMarker != "" {
			nextCursor = output.NextMarker
		} else if hasMore && len(files) > 0 {
			// 个别情况下 OBS 不返回 NextMarker，回退用最后一条 key 作为下次 marker
			nextCursor = files[len(files)-1].Key
		}
	}
	return files, nextCursor, hasMore, nil
}
