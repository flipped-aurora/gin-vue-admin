package material

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	// UploadImgURL 上传图片
	UploadImgURL = "https://qyapi.weixin.qq.com/cgi-bin/media/uploadimg?access_token=%s"

	// UploadTempFile 上传临时素材
	UploadTempFile = "https://qyapi.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s"
)

// UploadImgResponse 上传图片响应
type UploadImgResponse struct {
	util.CommonError
	URL string `json:"url"`
}

// UploadTempFileResponse 上传临时素材响应
type UploadTempFileResponse struct {
	util.CommonError
	MediaID  string `json:"media_id"`
	CreateAt string `json:"created_at"`
	Type     string `json:"type"`
}

// UploadImg 上传图片
// @see https://developer.work.weixin.qq.com/document/path/90256
func (r *Client) UploadImg(filename string) (*UploadImgResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostFile("media", filename, fmt.Sprintf(UploadImgURL, accessToken)); err != nil {
		return nil, err
	}
	result := &UploadImgResponse{}
	if err = util.DecodeWithError(response, result, "UploadImg"); err != nil {
		return nil, err
	}
	return result, nil
}

// UploadTempFile 上传临时素材
// @see https://developer.work.weixin.qq.com/document/path/90253
// @mediaType 媒体文件类型，分别有图片（image）、语音（voice）、视频（video），普通文件（file）
func (r *Client) UploadTempFile(filename string, mediaType string) (*UploadTempFileResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostFile("media", filename, fmt.Sprintf(UploadTempFile, accessToken, mediaType)); err != nil {
		return nil, err
	}
	result := &UploadTempFileResponse{}
	if err = util.DecodeWithError(response, result, "UploadTempFile"); err != nil {
		return nil, err
	}
	return result, nil
}
