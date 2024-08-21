package oss

import (
	"context"
	"errors"
	"fmt"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"os"
)

func NewDefaultQiNiu() *QiNiu {
	return &QiNiu{
		Bucket:    "geeleo",
		AccessKey: "ehF_E4x_EyO_wSN_nwqExyhXPe5hGl5Xjo89_cZ6",
		SecretKey: "FjfIpqUevEcVx9bQxdgiuX9Di-CUOrKFkR88CZAj",
	}
}

type QiNiu struct {
	Bucket    string `json:"bucket"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	ImgPath   string `json:"img_path"`
}

func qiniuConfig() *storage.Config {
	cfg := storage.Config{UseHTTPS: false, UseCdnDomains: false}
	return &cfg
}

type File struct {
	Path     string `json:"path"`
	Filename string `json:"filename"`
}

// UploadLocalFile 上传本地文件到七牛云
func (q *QiNiu) UploadLocalFile(localFilePath string, ossPath string) (*File, error) {
	file, err := os.Open(localFilePath)
	if err != nil {
		return nil, err
	}
	putPolicy := storage.PutPolicy{Scope: q.Bucket}
	mac := qbox.NewMac(q.AccessKey, q.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := qiniuConfig()
	formUploader := storage.NewFormUploader(cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{Params: map[string]string{"x:name": "github logo"}} // 创建文件 defer 关闭
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	fileKey := fmt.Sprintf("%s", ossPath) // 文件名格式 自己可以改 建议保证唯一性
	putErr := formUploader.Put(context.Background(), &ret, upToken, fileKey, file, stat.Size(), &putExtra)
	if putErr != nil {
		return nil, errors.New("function formUploader.Put() failed, err:" + putErr.Error())
	}
	return &File{
		Path:     q.ImgPath + "/" + ret.Key,
		Filename: ret.Key}, nil
}

// GetUploadToken 获取上传token
func (q *QiNiu) GetUploadToken() string {
	putPolicy := storage.PutPolicy{
		Scope: q.Bucket,
		//InsertOnly:1,
	}
	mac := qbox.NewMac(q.AccessKey, q.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}
