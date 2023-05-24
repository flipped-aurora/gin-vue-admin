package mini

import (
	"encoding/json"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/douyin/open/context"
	"github.com/flipped-aurora/gin-vue-admin/server/douyin/open/credential"
	"github.com/flipped-aurora/gin-vue-admin/server/douyin/util"
)

const (
	accessTokenURL    = "https://developer.toutiao.com/api/apps/v2/token"
	jscode2sessionURL = "https://developer.toutiao.com/api/apps/v2/jscode2session"
)

type Mini struct {
	*context.Context
}

// NewOauth 实例化授权信息
func NewMini(context *context.Context) *Mini {
	mini := new(Mini)
	mini.Context = context
	return mini
}

type commonReq struct {
	AppID  string `json:"appid"`
	Secret string `json:"secret"`
}

type commonRes struct {
	ErrNo   int64  `json:"err_no"`
	ErrTips string `json:"err_tips"`
}

type accessTokenReq struct {
	commonReq
	GrantType string `json:"grant_type"`
}

type accessTokenRes struct {
	commonRes
	Data struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int64  `json:"expires_in"`
		ExpiresAt   int64  `json:"expiresAt"`
	} `json:"data"`
}

func (mini *Mini) GetUserAccessToken() (accessToken credential.AccessToken, err error) {
	request := accessTokenReq{
		commonReq: commonReq{
			AppID:  mini.ClientKey,
			Secret: mini.ClientSecret,
		},
		GrantType: "client_credential",
	}
	var response []byte
	data, _ := json.Marshal(request)
	response, err = util.HTTPPost(accessTokenURL, string(data))
	if err != nil {
		return
	}
	var result accessTokenRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrNo != 0 {
		err = fmt.Errorf("GetUserAccessToken error : errcode=%v , errmsg=%v", result.ErrNo, result.ErrTips)
		return
	}
	accessToken.ErrCode = result.ErrNo
	accessToken.ErrMsg = result.ErrTips
	accessToken.ExpiresIn = result.Data.ExpiresIn
	accessToken.AccessToken = result.Data.AccessToken
	err = mini.SetAccessToken(&accessToken)
	if err != nil {
		return
	}
	return
}

type sessionReq struct {
	commonReq
	Code string `json:"code"`
}

type sessionRes struct {
	commonRes
	Data Session `json:"data"`
}

type Session struct {
	SessionKey      string `json:"session_key"`
	OpenID          string `json:"openid"`
	AnonymousOpenID string `json:"anonymous_openid"`
	Unionid         string `json:"unionid"`
}

func (mini *Mini) GetSession(code string) (session Session, err error) {
	request := sessionReq{
		commonReq: commonReq{
			AppID:  mini.ClientKey,
			Secret: mini.ClientSecret,
		},
		Code: code,
	}
	var response []byte
	data, _ := json.Marshal(request)
	response, err = util.HTTPPost(jscode2sessionURL, string(data))
	if err != nil {
		return
	}
	var result sessionRes
	fmt.Println(string(response))
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrNo != 0 {
		err = fmt.Errorf("GetUserAccessToken error : errcode=%v , errmsg=%v", result.ErrNo, result.ErrTips)
		return
	}
	return result.Data, nil
}
