package upload

import (
	"errors"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
)

type CloudflareR2 struct{}

func (c *CloudflareR2) UploadFile(file *multipart.FileHeader) (fileUrl string, fileName string, err error) {
	session := c.newSession()
	client := s3manager.NewUploader(session)

	fileKey := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	fileName = fmt.Sprintf("%s/%s", global.GVA_CONFIG.CloudflareR2.Path, fileKey)
	f, openError := file.Open()
	if openError != nil {
		global.GVA_LOG.Error("function file.Open() failed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() failed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	input := &s3manager.UploadInput{
		Bucket: aws.String(global.GVA_CONFIG.CloudflareR2.Bucket),
		Key:    aws.String(fileName),
		Body:   f,
	}

	_, err = client.Upload(input)
	if err != nil {
		global.GVA_LOG.Error("function uploader.Upload() failed", zap.Any("err", err.Error()))
		return "", "", err
	}

	return fmt.Sprintf("%s/%s", global.GVA_CONFIG.CloudflareR2.BaseURL,
			fileName),
		fileKey,
		nil
}

func (c *CloudflareR2) DeleteFile(key string) error {
	session := newSession()
	svc := s3.New(session)
	filename := global.GVA_CONFIG.CloudflareR2.Path + "/" + key
	bucket := global.GVA_CONFIG.CloudflareR2.Bucket

	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		global.GVA_LOG.Error("function svc.DeleteObject() failed", zap.Any("err", err.Error()))
		return errors.New("function svc.DeleteObject() failed, err:" + err.Error())
	}

	_ = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	return nil
}

func (*CloudflareR2) newSession() *session.Session {
	endpoint := fmt.Sprintf("%s.r2.cloudflarestorage.com", global.GVA_CONFIG.CloudflareR2.AccountID)

	return session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("auto"),
		Endpoint: aws.String(endpoint),
		Credentials: credentials.NewStaticCredentials(
			global.GVA_CONFIG.CloudflareR2.AccessKeyID,
			global.GVA_CONFIG.CloudflareR2.SecretAccessKey,
			"",
		),
	}))
}
