package image

import (
	"encoding/json"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/douyin/open/context"
	"github.com/flipped-aurora/gin-vue-admin/server/douyin/util"
)

const (
	// 上传图片
	imageUploadURL string = "https://open.douyin.com/image/upload?access_token=%s&open_id=%s"
	// 发布图片
	imageCreateURL string = "https://open.douyin.com/image/create?access_token=%s&open_id=%s"
)

// Image 图片
type Image struct {
	*context.Context
}

// NewImage .
func NewImage(context *context.Context) *Image {
	image := new(Image)
	image.Context = context
	return image
}

// Info 视频信息.
type Info struct {
	util.CommonError

	Image struct {
		ImageID string `json:"image_id"`
		Height  int64  `json:"height"`
		Width   int64  `json:"width"`
	} `json:"image"`
}

type uploadImageRes struct {
	Message string `json:"message"`
	Data    Info   `json:"data"`
}

// Upload 图片上传
func (image *Image) Upload(openid string, filename string) (imageInfo Info, err error) {
	accessToken, err := image.GetAccessToken(openid)
	if err != nil {
		return
	}

	uri := fmt.Sprintf(imageUploadURL, accessToken, openid)
	var response []byte
	response, err = util.PostFile("image", filename, uri)
	if err != nil {
		return
	}

	var result uploadImageRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.Data.ErrCode != 0 {
		err = fmt.Errorf("Upload error : errcode=%v , errmsg=%v", result.Data.ErrCode, result.Data.ErrMsg)
		return
	}
	imageInfo = result.Data
	return
}

// CreateImageReq .
type CreateImageReq struct {
	ImageID       string   `json:"image_id"`
	PoiID         string   `json:"poi_id"`
	PoiName       string   `json:"poi_name"`
	Text          string   `json:"text"`
	AtUsers       []string `json:"at_users"`
	MicroAppURL   string   `json:"micro_app_url"`
	MicroAppID    string   `json:"micro_app_id"`
	MicroAppTitle string   `json:"micro_app_title"`
}

// CreateInfo .
type CreateInfo struct {
	util.CommonError

	ItemID string `json:"item_id"`
}

type createRes struct {
	Message string     `json:"message"`
	Data    CreateInfo `json:"data"`
}

// Create 图片发布.
func (image *Image) Create(openid string, imageInfo *CreateImageReq) (info CreateInfo, err error) {
	accessToken, err := image.GetAccessToken(openid)
	if err != nil {
		return
	}

	uri := fmt.Sprintf(imageCreateURL, accessToken, openid)
	var response []byte
	response, err = util.PostJSON(uri, imageInfo)
	if err != nil {
		return
	}

	var result createRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.Data.ErrCode != 0 {
		err = fmt.Errorf("Create error : errcode=%v , errmsg=%v", result.Data.ErrCode, result.Data.ErrMsg)
		return
	}
	info = result.Data
	return
}
