package upload

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"github.com/pkg/errors"
	"mime/multipart"
)

var HuaWeiObs = new(_obs)

type _obs struct{}

func NewHuaWeiObsClient() (client *obs.ObsClient, err error) {
	return obs.New(global.GVA_CONFIG.HuaWeiObs.AccessKey, global.GVA_CONFIG.HuaWeiObs.SecretKey, global.GVA_CONFIG.HuaWeiObs.Endpoint)
}

func (o *_obs) UploadFile(file *multipart.FileHeader) (filename string, filepath string, err error) {
	var open multipart.File
	open, err = file.Open()
	if err != nil {
		return filename, filepath, err
	}
	filename = file.Filename
	input := &obs.PutObjectInput{
		PutObjectBasicInput: obs.PutObjectBasicInput{
			ObjectOperationInput: obs.ObjectOperationInput{
				Bucket: global.GVA_CONFIG.HuaWeiObs.Bucket,
				Key:    filename,
			},
			ContentType: file.Header.Get("content-type"),
		},
		Body: open,
	}

	var client *obs.ObsClient
	client, err = NewHuaWeiObsClient()
	if err != nil {
		return filepath, filename, errors.Wrap(err, "获取华为对象存储对象失败!")
	}

	_, err = client.PutObject(input)
	if err != nil {
		return filepath, filename, errors.Wrap(err, "文件上传失败!")
	}
	filepath = global.GVA_CONFIG.HuaWeiObs.Path + "/" + filename
	return filepath, filename, err
}

func (o *_obs) DeleteFile(key string) error {
	client, err := NewHuaWeiObsClient()
	if err != nil {
		return errors.Wrap(err, "获取华为对象存储对象失败!")
	}
	input := &obs.DeleteObjectInput{
		Bucket: global.GVA_CONFIG.HuaWeiObs.Bucket,
		Key:    key,
	}
	var output *obs.DeleteObjectOutput
	output, err = client.DeleteObject(input)
	if err != nil {
		return errors.Wrapf(err, "删除对象(%s)失败!, output: %v", key, output)
	}
	return nil
}
