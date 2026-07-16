package upload

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
)

type CloudflareR2 struct{}

func (c *CloudflareR2) UploadFile(ctx context.Context, file *multipart.FileHeader) (fileUrl string, fileName string, err error) {
	client, err := c.newR2Client()
	if err != nil {
		return "", "", err
	}
	uploader := manager.NewUploader(client)

	fileKey := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	fileName = fmt.Sprintf("%s/%s", global.GVA_CONFIG.CloudflareR2.Path, fileKey)
	f, openError := file.Open()
	if openError != nil {
		logger.WithCtx(ctx).Mod("upload").Err(openError).Error("function file.Open() failed")
		return "", "", errors.New("function file.Open() failed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(global.GVA_CONFIG.CloudflareR2.Bucket),
		Key:    aws.String(fileName),
		Body:   f,
	})
	if err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("function uploader.Upload() failed")
		return "", "", err
	}

	return fmt.Sprintf("%s/%s", global.GVA_CONFIG.CloudflareR2.BaseURL, fileName), fileKey, nil
}

func (c *CloudflareR2) DeleteFile(ctx context.Context, key string) error {
	client, err := c.newR2Client()
	if err != nil {
		return err
	}
	filename := global.GVA_CONFIG.CloudflareR2.Path + "/" + key
	bucket := global.GVA_CONFIG.CloudflareR2.Bucket

	_, err = client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("function client.DeleteObject() failed")
		return errors.New("function client.DeleteObject() failed, err:" + err.Error())
	}

	waiter := s3.NewObjectNotExistsWaiter(client)
	_ = waiter.Wait(context.TODO(), &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	}, 30*time.Second)

	return nil
}

// Exists 检查对象是否存在。404 统一降级为 (false, nil)。
func (c *CloudflareR2) Exists(ctx context.Context, key string) (bool, error) {
	client, err := c.newR2Client()
	if err != nil {
		return false, err
	}
	bucket := global.GVA_CONFIG.CloudflareR2.Bucket
	fullKey := global.GVA_CONFIG.CloudflareR2.Path + "/" + key

	_, err = client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fullKey),
	})
	if err != nil {
		// 优先用 smithy 的 ResponseError 判定 HTTP 状态码；兜底用字符串匹配。
		var re *smithyhttp.ResponseError
		if errors.As(err, &re) && re.HTTPStatusCode() == 404 {
			return false, nil
		}
		if strings.Contains(err.Error(), "404") {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// DeleteFiles 批量删除，返回逐个失败项；整体 err 仅用于请求本身致命错误。
func (c *CloudflareR2) DeleteFiles(ctx context.Context, keys []string) (failed []DeleteFailure, err error) {
	client, err := c.newR2Client()
	if err != nil {
		return nil, err
	}
	bucket := global.GVA_CONFIG.CloudflareR2.Bucket
	prefix := global.GVA_CONFIG.CloudflareR2.Path

	objects := make([]types.ObjectIdentifier, 0, len(keys))
	for _, key := range keys {
		fullKey := key
		// 入参可能是短 key（不含前缀），也可能是 List 出来的全路径 key；统一拼接前缀时避免重复。
		if prefix != "" && !strings.HasPrefix(key, prefix+"/") {
			fullKey = prefix + "/" + key
		}
		objects = append(objects, types.ObjectIdentifier{Key: aws.String(fullKey)})
	}

	out, err := client.DeleteObjects(ctx, &s3.DeleteObjectsInput{
		Bucket: aws.String(bucket),
		Delete: &types.Delete{
			Objects: objects,
			Quiet:   aws.Bool(true),
		},
	})
	if err != nil {
		return nil, errors.New("function client.DeleteObjects() failed, err:" + err.Error())
	}

	for _, e := range out.Errors {
		key := ""
		if e.Key != nil {
			key = *e.Key
		}
		msg := "unknown error"
		if e.Message != nil {
			msg = *e.Message
		}
		failed = append(failed, DeleteFailure{Key: key, Err: errors.New(msg)})
	}
	return failed, nil
}

// ListFiles 按前缀列举存储对象，cursor 映射到 ContinuationToken。
func (c *CloudflareR2) ListFiles(ctx context.Context, prefix, cursor string, limit int) (files []FileInfo, nextCursor string, hasMore bool, err error) {
	client, err := c.newR2Client()
	if err != nil {
		return nil, "", false, err
	}
	bucket := global.GVA_CONFIG.CloudflareR2.Bucket

	if limit <= 0 {
		limit = 100
	}

	input := &s3.ListObjectsV2Input{
		Bucket:  aws.String(bucket),
		Prefix:  aws.String(prefix),
		MaxKeys: aws.Int32(int32(limit)),
	}
	if cursor != "" {
		input.ContinuationToken = aws.String(cursor)
	}

	out, err := client.ListObjectsV2(ctx, input)
	if err != nil {
		return nil, "", false, errors.New("function client.ListObjectsV2() failed, err:" + err.Error())
	}

	for _, obj := range out.Contents {
		fi := FileInfo{}
		if obj.Key != nil {
			fi.Key = *obj.Key
		}
		if obj.Size != nil {
			fi.Size = *obj.Size
		}
		if obj.LastModified != nil {
			fi.LastModified = *obj.LastModified
		}
		files = append(files, fi)
	}

	if out.NextContinuationToken != nil {
		nextCursor = *out.NextContinuationToken
	}
	if out.IsTruncated != nil {
		hasMore = *out.IsTruncated
	}
	return files, nextCursor, hasMore, nil
}

func (*CloudflareR2) newR2Client() (*s3.Client, error) {
	endpoint := fmt.Sprintf("https://%s.r2.cloudflarestorage.com", global.GVA_CONFIG.CloudflareR2.AccountID)

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("auto"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			global.GVA_CONFIG.CloudflareR2.AccessKeyID,
			global.GVA_CONFIG.CloudflareR2.SecretAccessKey,
			"",
		)),
	)
	if err != nil {
		logger.Bg().Mod("upload").Err(err).Error("function config.LoadDefaultConfig() failed")
		return nil, errors.New("function config.LoadDefaultConfig() failed, err:" + err.Error())
	}

	return s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(endpoint)
	}), nil
}
