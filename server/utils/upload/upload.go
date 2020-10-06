package upload

import (
	"gin-vue-admin/global"
	"mime/multipart"
)

var Oss OSS

type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}

func InitOss() {
	switch global.GVA_CONFIG.System.OssType {
	case "local":
		Oss = &Local{}
	case "qiniu":
		Oss = &Qiniu{}
	default:
		Oss = &Local{}
	}
}