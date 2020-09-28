package upload

import (
	"fmt"
	"mime/multipart"
	"time"
)

const (
	// LOCAL 本地存储
	LOCAL = iota
	// QINIU 七牛云
	QINIU
	// MINIO 对象存储
	MINIO
	// ALIYUN 对象存储
	ALIYUN
)

// Uploader 文件上传接口
type Uploader interface {
	Upload(file *multipart.FileHeader) (error, string, string) // 上传文件
	DeleteFile(key string) error                               // 删除文件
}

// UploadUtils 对象存储Client
var UploadUtils Uploader

// New 生成对象存储Client
func New(ossType int) {
	if ossType == LOCAL {
		// TODO
	} else if ossType == QINIU {
		UploadUtils = &Qiniu{}
	} else if ossType == MINIO {
		UploadUtils = &Minio{}
	} else if ossType == ALIYUN {
		UploadUtils = &Aliyun{}
	} else {
		// TODO ?local?
	}
}

// getObjectName 生成文件名
func getObjectName(filename string) string {
	folder := time.Now().Format("20060102")
	return fmt.Sprintf("%s/%d%s", folder, time.Now().Unix(), filename) // 文件名格式 自己可以改 建议保证唯一性
}
