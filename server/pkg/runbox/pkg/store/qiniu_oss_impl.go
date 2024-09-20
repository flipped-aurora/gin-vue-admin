package store

import (
	"context"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/stringsx"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/httpx"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"os"
	"path/filepath"
	"strings"
)

func NewDefaultQiNiu() *QiNiu {

	return &QiNiu{
		Domain:           "http://cdn.geeleo.com",
		DownloadRootPath: os.TempDir(),
		Bucket:           "geeleo",
		AccessKey:        "ehF_E4x_EyO_wSN_nwqExyhXPe5hGl5Xjo89_cZ6",
		SecretKey:        "FjfIpqUevEcVx9bQxdgiuX9Di-CUOrKFkR88CZAj",
	}
}

type QiNiu struct {
	DownloadRootPath string `json:"download_root_path"` //下载文件的默认本地根路径
	Domain           string `json:"domain"`             //云存储的域名
	Bucket           string `json:"bucket"`             //空间
	AccessKey        string `json:"access_key"`
	SecretKey        string `json:"secret_key"`
	ImgPath          string `json:"img_path"` //存储路径
}

func (q *QiNiu) GetFullPath(savePath string) string {
	return q.Domain + savePath
}

func qiniuConfig() *storage.Config {
	cfg := storage.Config{UseHTTPS: false, UseCdnDomains: false}
	return &cfg
}

//type File struct {
//	Path     string `json:"path"`
//	Filename string `json:"filename"`
//}

// FileSave 上传本地文件到七牛云
func (q *QiNiu) FileSave(localFilePath string, ossPath string) (*FileSaveInfo, error) {
	file, err := os.Open(localFilePath)
	if err != nil {
		return nil, err
	}
	name := file.Name()
	fileType := stringsx.GetSuffix(name, ".")
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

	return &FileSaveInfo{
		SavePath:     q.ImgPath + "/" + ret.Key,
		SaveFullPath: q.Domain + q.ImgPath + "/" + ret.Key,
		FileName:     name,
		FileType:     fileType,
	}, nil
}

// GetFile 下载文件到本地
func (q *QiNiu) GetFile(savePath string) (*GetFileInfo, error) {
	if savePath == "" {
		return nil, fmt.Errorf("请正确填写文件地址")
	}
	if savePath[0] != '/' {
		savePath = "/" + savePath
	}
	tempPath := filepath.Join(q.DownloadRootPath, savePath)
	url := q.GetFullPath(savePath)
	err := httpx.DownloadFile(url, tempPath)
	if err != nil {
		return nil, err
	}
	p := strings.ReplaceAll(savePath, "\\", "/")
	fileName := stringsx.GetSuffix(p, "/")
	fileSuffix := stringsx.GetSuffix(p, ".")

	return &GetFileInfo{
		FileSaveInfo: FileSaveInfo{
			SavePath:     savePath,
			SaveFullPath: q.Domain + savePath,
			FileName:     fileName,
		},
		FileLocalPath: tempPath,
		FileType:      fileSuffix,
	}, nil
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

// DeleteFile 删除文件
func (q *QiNiu) DeleteFile(savePath string) error {
	savePath = strings.TrimPrefix(savePath, "/")
	mac := qbox.NewMac(q.AccessKey, q.SecretKey)
	cfg := qiniuConfig()
	bucketManager := storage.NewBucketManager(mac, cfg)
	if err := bucketManager.Delete(q.Bucket, savePath); err != nil {
		return errors.New("function bucketManager.Delete() failed, err:" + err.Error())
	}
	return nil
}
