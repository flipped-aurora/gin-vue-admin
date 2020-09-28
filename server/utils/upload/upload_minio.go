package upload

import (
	"context"
	"fmt"
	"gin-vue-admin/global"
	"log"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
)

// minio
type Minio struct {
}

// Upload 上传、接收参数 一个文件流
func (oss *Minio) Upload(file *multipart.FileHeader) (err error, path string, key string) {
	config := global.GVA_CONFIG.OSS
	// Initialize minio client object.
	minioClient, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyID, config.SecretAccessKey, ""),
		Secure: config.UseSSL,
	})
	if err != nil {
		fmt.Println(err)
		return err, "", ""
	}

	// Make a new bucket called mymusic.
	bucketName := config.Bucket

	ctx := context.Background()
	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: ""})

	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			fmt.Println(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	// 读取文件
	f, e := file.Open()
	if e != nil {
		fmt.Println(e)
		return e, "", ""
	}

	// 获取文件类型
	contentType := file.Header.Get("content-type")
	objectName := getObjectName(file.Filename)
	n, err := minioClient.PutObject(ctx, bucketName, objectName, f, file.Size, minio.PutObjectOptions{ContentType: contentType})

	if err != nil {
		global.GVA_LOG.Error("upload file fail:", zap.Any("err", err))
		return e, "", ""
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
	return err, config.Path + "/" + bucketName + "/" + objectName, objectName
}

// DeleteFile 删除文件
func (oss *Minio) DeleteFile(key string) error {

	config := global.GVA_CONFIG.OSS

	minioClient, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyID, config.SecretAccessKey, ""),
		Secure: config.UseSSL,
	})

	if err != nil {
		fmt.Println(err)
		return err
	}

	opts := minio.RemoveObjectOptions{
		GovernanceBypass: true,
	}

	err = minioClient.RemoveObject(context.Background(), config.Bucket, key, opts)
	if err != nil {
		global.GVA_LOG.Error("delete file fail:", zap.Any("err", err))
		return err
	}

	return nil
}
