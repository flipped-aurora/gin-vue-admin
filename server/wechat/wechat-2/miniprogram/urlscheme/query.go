package urlscheme

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	querySchemeURL = "https://api.weixin.qq.com/wxa/queryscheme?access_token=%s"
)

// QueryScheme 获取小程序访问scheme
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-scheme/urlscheme.query.html#参数
type QueryScheme struct {
	// 小程序 scheme 码
	Scheme string `json:"scheme"`
}

// SchemeInfo scheme 配置
type SchemeInfo struct {
	// 小程序 appid。
	AppID string `json:"appid"`
	// 小程序页面路径。
	Path string `json:"path"`
	// 小程序页面query。
	Query string `json:"query"`
	// 创建时间，为 Unix 时间戳。
	CreateTime int64 `json:"create_time"`
	// 到期失效时间，为 Unix 时间戳，0 表示永久生效
	ExpireTime int64 `json:"expire_time"`
	// 要打开的小程序版本。正式版为"release"，体验版为"trial"，开发版为"develop"。
	EnvVersion EnvVersion `json:"env_version"`
}

// resQueryScheme 返回结构体
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-scheme/urlscheme.query.html#参数
type resQueryScheme struct {
	// 通用错误
	*util.CommonError
	// scheme 配置
	SchemeInfo SchemeInfo `json:"scheme_info"`
	// 访问该链接的openid，没有用户访问过则为空字符串
	VisitOpenid string `json:"visit_openid"`
}

// QueryScheme 查询小程序 scheme 码
func (u *URLScheme) QueryScheme(querySchemeParams QueryScheme) (schemeInfo SchemeInfo, visitOpenid string, err error) {
	var accessToken string
	accessToken, err = u.GetAccessToken()
	if err != nil {
		return
	}

	urlStr := fmt.Sprintf(querySchemeURL, accessToken)
	var response []byte
	response, err = util.PostJSON(urlStr, querySchemeParams)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	var res resQueryScheme
	err = util.DecodeWithError(response, &res, "QueryScheme")
	if err != nil {
		return
	}

	return res.SchemeInfo, res.VisitOpenid, nil
}
