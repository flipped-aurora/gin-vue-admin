package upload

import (
	"fmt"
	"gin-vue-admin/global"
	"mime/multipart"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go.uber.org/zap"
)

// Aliyun 对象存储
type Aliyun struct {
}

// Upload 上传、接收参数 一个文件流
func (aliyun *Aliyun) Upload(file *multipart.FileHeader) (err error, path string, key string) {
	config := global.GVA_CONFIG.OSS
	client, err := oss.New(config.Endpoint, config.AccessKeyID, config.SecretAccessKey,
		oss.Timeout(10, 120))
	if err != nil {
		fmt.Println("Error:", err)
		return err, "", ""
	}

	// 获取存储空间。
	bucket, err := client.Bucket(config.Bucket)
	if err != nil {
		fmt.Println("Error:", err)
		return err, "", ""
	}

	// 指定存储类型为标准存储，缺省也为标准存储。
	storageType := oss.ObjectStorageClass(oss.StorageStandard)

	// 指定存储类型为归档存储。
	// storageType := oss.ObjectStorageClass(oss.StorageArchive)

	// 读取文件
	f, e := file.Open()
	if e != nil {
		fmt.Println(e)
		return e, "", ""
	}

	// 获取文件类型
	contentType := file.Header.Get("content-type")
	objectType := oss.ContentType(contentType)

	// 指定访问权限为公共读，缺省为继承bucket的权限。
	objectAcl := oss.ObjectACL(oss.ACLPublicRead)

	// 文件对象名
	objectName := getObjectName(file.Filename)

	// 上传
	err = bucket.PutObject(objectName, f, storageType, objectType, objectAcl)
	if err != nil {
		global.GVA_LOG.Error("upload file fail:", zap.Any("err", err))
		return e, "", ""
	}

	return err, config.Path + "/" + objectName, objectName
}

// DeleteFile 删除文件
func (aliyun *Aliyun) DeleteFile(key string) error {
	config := global.GVA_CONFIG.OSS
	client, err := oss.New(config.Endpoint, config.AccessKeyID, config.SecretAccessKey,
		oss.Timeout(10, 120))
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(config.Bucket)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	// 删除单个文件。objectName表示删除OSS文件时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	// 如需删除文件夹，请将objectName设置为对应的文件夹名称。如果文件夹非空，则需要将文件夹下的所有object删除后才能删除该文件夹。
	err = bucket.DeleteObject(key)
	if err != nil {
		global.GVA_LOG.Error("delete file fail:", zap.Any("err", err))
		return err
	}

	return nil
}
