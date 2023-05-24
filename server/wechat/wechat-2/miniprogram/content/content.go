package content

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/context"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	checkTextURL  = "https://api.weixin.qq.com/wxa/msg_sec_check?access_token=%s"
	checkImageURL = "https://api.weixin.qq.com/wxa/img_sec_check?access_token=%s"
)

const (
	checkTextReserveURL  = "https://api2.weixin.qq.com/wxa/msg_sec_check?access_token=%s"
	checkImageReserveURL = "https://api2.weixin.qq.com/wxa/img_sec_check?access_token=%s"
)

// Content 内容安全
type Content struct {
	*context.Context
}

// NewContent 内容安全接口
func NewContent(ctx *context.Context) *Content {
	return &Content{ctx}
}

// CheckText 检测文字
// @text 需要检测的文字
// Deprecated
// 采用 security.MsgCheckV1 替代，返回值更加丰富
func (content *Content) CheckText(text string) error {
	accessToken, err := content.GetAccessToken()
	if err != nil {
		return err
	}
	var response []byte

	if response, err = util.PostJSON(
		fmt.Sprintf(checkTextURL, accessToken),
		map[string]string{
			"content": text,
		},
	); err != nil {
		if response, err = util.PostJSON(
			fmt.Sprintf(checkTextReserveURL, accessToken),
			map[string]string{
				"content": text,
			},
		); err != nil {
			return err
		}
	}
	return util.DecodeWithCommonError(response, "ContentCheckText")
}

// CheckImage 检测图片
// 所传参数为要检测的图片文件的绝对路径，图片格式支持PNG、JPEG、JPG、GIF, 像素不超过 750 x 1334，同时文件大小以不超过 300K 为宜，否则可能报错
// @media 图片文件的绝对路径
// Deprecated
// 采用 security.ImageCheckV1 替代，返回值更加丰富
func (content *Content) CheckImage(media string) error {
	accessToken, err := content.GetAccessToken()
	if err != nil {
		return err
	}
	var response []byte

	if response, err = util.PostFile(
		"media",
		media,
		fmt.Sprintf(checkImageURL, accessToken),
	); err != nil {
		if response, err = util.PostFile(
			"media",
			media,
			fmt.Sprintf(checkImageReserveURL, accessToken),
		); err != nil {
			return err
		}
	}
	return util.DecodeWithCommonError(response, "ContentCheckImage")
}
