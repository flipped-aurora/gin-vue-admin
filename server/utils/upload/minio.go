package upload

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/component"
	"mime/multipart"
)

/*
* @Author <bypanghu> (bypanghu@163.com)
* @Date 2024/1/29 09:38
**/

type MinioOSS struct{}

func (o *MinioOSS) UploadFile(file *multipart.FileHeader) (string, string, error) {
	return component.MinioPkg.UploadByHeader(context.Background(), file, "www")
}
// DeleteFile 请自行实现
func (o *MinioOSS) DeleteFile(key string) error {
	return nil
}
