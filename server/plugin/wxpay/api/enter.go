package api

type ApiGroup struct {
	WxPay
}

var WxApiGroupApp = new(ApiGroup)

const TokenUrl = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
const Urlstr = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_base&state=123#wechat_redirect"

// AccessToken 微信授权返回数据结构 用来获取openid
type AccessToken struct {
	Token          string `json:"access_token"`
	ExpiresIn      int    `json:"expires_in"`
	RefreshToken   string `json:"refresh_token"`
	OpenID         string `json:"openId"`
	Scope          string `json:"scope"`
	IsSnapshotUser int    `json:"is_snapshotuser"`
	UnionID        string `json:"unionid"`
}
