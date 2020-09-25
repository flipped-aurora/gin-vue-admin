package upload

import (
	"mime/multipart"
)

const (
	LOCAL = iota
	QINIU
	MINIO
)

// 文件上传接口
type Uploader interface {
	Upload(file *multipart.FileHeader) (error, string, string)
	DeleteFile(key string) error
}

var UploadUtils Uploader

func New(ossType int) {
	if ossType == LOCAL {
		// TODO
	} else if ossType == QINIU {
		UploadUtils = &Qiniu{}
	} else if ossType == MINIO {
		UploadUtils = &Minio{}
	} else {
		// TODO ?local?
	}
}
