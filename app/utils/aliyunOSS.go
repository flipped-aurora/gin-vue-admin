package utils

import (
	"app/global"
	"app/response"
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"strings"
	"time"
)

type FileStruct struct {
	Name string `json:"name" gorm:"comment:文件名"` // 文件名
	Url  string `json:"url" gorm:"comment:文件地址"` // 文件地址
	Tag  string `json:"tag" gorm:"comment:文件标签"` // 文件标签
	Key  string `json:"key" gorm:"comment:编号"`   // 编号
}

func UploadFileUseAliyun(c *gin.Context) {
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		response.FailWithDetailed("接受文件失败", err.Error(), c)
		return
	}
	file, err := UploadFile(header)
	if err != nil {
		response.FailWithDetailed("上传文件失败", err.Error(), c)
		return
	}
	response.OkWithDetailed(file, "上传成功", c)
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UploadFile
//@description: 根据配置文件判断是文件上传到本地或者七牛云
//@param: header *multipart.FileHeader, noSave string
//@return: file model.ExaFileUploadAndDownload, err error

func UploadFile(header *multipart.FileHeader) (file FileStruct, err error) {

	filePath, key, uploadErr := UploadFileService(header)

	if uploadErr != nil {

		panic(uploadErr)
	}
	s := strings.Split(header.Filename, ".")
	f := FileStruct{
		Url:  filePath,
		Name: header.Filename,
		Tag:  s[len(s)-1],
		Key:  key,
	}

	return f, nil
}

func UploadFileService(file *multipart.FileHeader) (string, string, error) {
	bucket, err := NewBucket()
	if err != nil {
		//global.GVA_LOG.Error("function AliyunOSS.NewBucket() Failed", zap.Any("err", err.Error()))
		return "", "", errors.New("function AliyunOSS.NewBucket() Failed, err:" + err.Error())
	}

	// 读取本地文件。
	f, openError := file.Open()
	if openError != nil {
		//global.GVA_LOG.Error("function file.Open() Failed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Failed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭
	// 上传阿里云路径 文件名格式 自己可以改 建议保证唯一性
	// yunFileTmpPath := filepath.Join("uploads", time.Now().Format("2006-01-02")) + "/" + file.Filename
	yunFileTmpPath := global.AliyunOSS.BasePath + "/" + "uploads" + "/" + time.Now().Format("2006-01-02") + "/" + fmt.Sprintf("%d", time.Now().Unix()) + file.Filename

	// 上传文件流。
	err = bucket.PutObject(yunFileTmpPath, f)
	if err != nil {
		//global.GVA_LOG.Error("function formUploader.Put() Failed", zap.Any("err", err.Error()))
		return "", "", errors.New("function formUploader.Put() Failed, err:" + err.Error())
	}

	return global.AliyunOSS.BucketUrl + "/" + yunFileTmpPath, yunFileTmpPath, nil
}
func NewBucket() (*oss.Bucket, error) {
	fmt.Println(global.AliyunOSS)
	// 创建OSSClient实例。
	client, err := oss.New(global.AliyunOSS.Endpoint, global.AliyunOSS.AccessKeyId, global.AliyunOSS.AccessKeySecret)
	if err != nil {
		return nil, err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(global.AliyunOSS.BucketName)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}
