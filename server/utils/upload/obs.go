package upload

import (
	"fmt"
	"mime/multipart"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"github.com/pkg/errors"
)

var HuaWeiObs = new(Obs)

type Obs struct{}

func NewHuaWeiObsClient() (client *obs.ObsClient, err error) {
	return obs.New(global.GVA_CONFIG.HuaWeiObs.AccessKey, global.GVA_CONFIG.HuaWeiObs.SecretKey, global.GVA_CONFIG.HuaWeiObs.Endpoint)
}

func (o *Obs) UploadFile(file *multipart.FileHeader) (string, string, error) {
	// var open multipart.File
	open, err := file.Open()
	if err != nil {
		return "", "", err
	}
	defer open.Close()
	filename := file.Filename
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
		return "", "", errors.Wrap(err, global.Translate("utils.getHuaweiObjectFailed"))
	}

	_, err = client.PutObject(input)
	if err != nil {
		return "", "", errors.Wrap(err, global.Translate("utils.fileUploadFailed"))
	}
	filepath := global.GVA_CONFIG.HuaWeiObs.Path + "/" + filename
	return filepath, filename, err
}

func (o *Obs) DeleteFile(key string) error {
	client, err := NewHuaWeiObsClient()
	if err != nil {
		return errors.Wrap(err, global.Translate("utils.getHuaweiObjectFailedDuplicate"))
	}
	input := &obs.DeleteObjectInput{
		Bucket: global.GVA_CONFIG.HuaWeiObs.Bucket,
		Key:    key,
	}
	var output *obs.DeleteObjectOutput
	output, err = client.DeleteObject(input)
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf(global.Translate("deleteObjectFailed"), key, output))
	}
	return nil
}
