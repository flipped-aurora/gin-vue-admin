package urllink

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/context"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

// URLLink 小程序 URL Link
type URLLink struct {
	*context.Context
}

// NewURLLink 实例化
func NewURLLink(ctx *context.Context) *URLLink {
	return &URLLink{Context: ctx}
}

const generateURL = "https://api.weixin.qq.com/wxa/generate_urllink"

// TExpireType 失效类型 (指定时间戳/指定间隔)
type TExpireType int

const (
	// ExpireTypeTime 指定时间戳后失效
	ExpireTypeTime TExpireType = 0

	// ExpireTypeInterval 间隔指定天数后失效
	ExpireTypeInterval TExpireType = 1
)

// ULParams 请求参数
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-link/urllink.generate.html#请求参数
type ULParams struct {
	Path  string `json:"path"`
	Query string `json:"query"`
	// envVersion 要打开的小程序版本。正式版为 "release"，体验版为 "trial"，开发版为 "develop"
	EnvVersion     string      `json:"env_version,omitempty"`
	IsExpire       bool        `json:"is_expire"`
	ExpireType     TExpireType `json:"expire_type"`
	ExpireTime     int64       `json:"expire_time"`
	ExpireInterval int         `json:"expire_interval"`
}

// ULResult 返回的结果
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-link/urllink.generate.html#返回值
type ULResult struct {
	util.CommonError

	URLLink string `json:"url_link"`
}

// Generate 生成url link
func (u *URLLink) Generate(params *ULParams) (string, error) {
	accessToken, err := u.GetAccessToken()
	if err != nil {
		return "", err
	}

	uri := fmt.Sprintf("%s?access_token=%s", generateURL, accessToken)
	response, err := util.PostJSON(uri, params)
	if err != nil {
		return "", err
	}
	var resp ULResult
	err = util.DecodeWithError(response, &resp, "URLLink.Generate")
	if err != nil {
		return "", err
	}
	return resp.URLLink, nil
}
