package oauth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/officialaccount/context"
	officialOauth "github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/officialaccount/oauth"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	platformRedirectOauthURL = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s&component_appid=%s#wechat_redirect"
	platformAccessTokenURL   = "https://api.weixin.qq.com/sns/oauth2/component/access_token?appid=%s&code=%s&grant_type=authorization_code&component_appid=%s&component_access_token=%s"
)

// Oauth 平台代发起oauth2网页授权
type Oauth struct {
	*context.Context
}

// NewOauth 实例化平台代发起oauth2网页授权
func NewOauth(context *context.Context) *Oauth {
	auth := new(Oauth)
	auth.Context = context
	return auth
}

// GetRedirectURL 第三方平台 - 获取跳转的url地址
func (oauth *Oauth) GetRedirectURL(redirectURI, scope, state, appID string) (string, error) {
	// url encode
	urlStr := url.QueryEscape(redirectURI)
	return fmt.Sprintf(platformRedirectOauthURL, appID, urlStr, scope, state, oauth.AppID), nil
}

// Redirect 第三方平台 - 跳转到网页授权
func (oauth *Oauth) Redirect(writer http.ResponseWriter, req *http.Request, redirectURI, scope, state, appID string) error {
	location, err := oauth.GetRedirectURL(redirectURI, scope, state, appID)
	if err != nil {
		return err
	}
	http.Redirect(writer, req, location, http.StatusFound)
	return nil
}

// GetUserAccessToken 第三方平台 - 通过网页授权的code 换取access_token(区别于context中的access_token)
func (oauth *Oauth) GetUserAccessToken(code, appID, componentAccessToken string) (result officialOauth.ResAccessToken, err error) {
	urlStr := fmt.Sprintf(platformAccessTokenURL, appID, code, oauth.AppID, componentAccessToken)
	var response []byte
	response, err = util.HTTPGet(urlStr)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetUserAccessToken error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}
