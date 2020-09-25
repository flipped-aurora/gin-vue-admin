package upload

import (
	"gin-vue-admin/global"
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

func init() {
	if global.GVA_CONFIG.System.OssType == LOCAL {
		// TODO
	} else if global.GVA_CONFIG.System.OssType == QINIU {
		// TODO
	} else if global.GVA_CONFIG.System.OssType == MINIO {
		Uploader = &Minio{}
	} else {
		// TODO ?local?
	}
}
