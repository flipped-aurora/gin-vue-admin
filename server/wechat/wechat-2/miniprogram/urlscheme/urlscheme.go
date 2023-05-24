package urlscheme

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/context"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

// URLScheme 小程序 URL Scheme
type URLScheme struct {
	*context.Context
}

// NewURLScheme 实例化
func NewURLScheme(ctx *context.Context) *URLScheme {
	return &URLScheme{Context: ctx}
}

const generateURL = "https://api.weixin.qq.com/wxa/generatescheme"

// TExpireType 失效类型 (指定时间戳/指定间隔)
type TExpireType int

// EnvVersion 要打开的小程序版本
type EnvVersion string

const (
	// ExpireTypeTime 指定时间戳后失效
	ExpireTypeTime TExpireType = 0
	// ExpireTypeInterval 间隔指定天数后失效
	ExpireTypeInterval TExpireType = 1

	// EnvVersionRelease 正式版为"release"
	EnvVersionRelease EnvVersion = "release"
	// EnvVersionTrial 体验版为"trial"
	EnvVersionTrial EnvVersion = "trial"
	// EnvVersionDevelop 开发版为"develop"
	EnvVersionDevelop EnvVersion = "develop"
)

// JumpWxa 跳转到的目标小程序信息
type JumpWxa struct {
	Path  string `json:"path"`
	Query string `json:"query"`
	// envVersion 要打开的小程序版本。正式版为 "release"，体验版为 "trial"，开发版为 "develop"
	EnvVersion EnvVersion `json:"env_version,omitempty"`
}

// USParams 请求参数
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-scheme/urlscheme.generate.html#请求参数
type USParams struct {
	JumpWxa        *JumpWxa    `json:"jump_wxa"`
	ExpireType     TExpireType `json:"expire_type"`
	ExpireTime     int64       `json:"expire_time"`
	ExpireInterval int         `json:"expire_interval"`
}

// USResult 返回的结果
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-scheme/urlscheme.generate.html#返回值
type USResult struct {
	util.CommonError

	OpenLink string `json:"openlink"`
}

// Generate 生成url link
func (u *URLScheme) Generate(params *USParams) (string, error) {
	accessToken, err := u.GetAccessToken()
	if err != nil {
		return "", err
	}

	uri := fmt.Sprintf("%s?access_token=%s", generateURL, accessToken)
	response, err := util.PostJSON(uri, params)
	if err != nil {
		return "", err
	}
	var resp USResult
	err = util.DecodeWithError(response, &resp, "URLScheme.Generate")
	if err != nil {
		return "", err
	}
	return resp.OpenLink, nil
}
