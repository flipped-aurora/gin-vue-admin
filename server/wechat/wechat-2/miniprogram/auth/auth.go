package auth

import (
	context2 "context"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/context"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	code2SessionURL = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"

	checkEncryptedDataURL = "https://api.weixin.qq.com/wxa/business/checkencryptedmsg?access_token=%s"

	getPhoneNumber = "https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=%s"
)

const (
	code2SessionReserveURL = "https://api2.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"

	checkEncryptedDataReserveURL = "https://api2.weixin.qq.com/wxa/business/checkencryptedmsg?access_token=%s"

	getPhoneNumberReserve = "https://api2.weixin.qq.com/wxa/business/getuserphonenumber?access_token=%s"
)

// Auth 登录/用户信息
type Auth struct {
	*context.Context
}

// NewAuth new auth
func NewAuth(ctx *context.Context) *Auth {
	return &Auth{ctx}
}

// ResCode2Session 登录凭证校验的返回结果
type ResCode2Session struct {
	util.CommonError

	OpenID     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
	UnionID    string `json:"unionid"`     // 用户在开放平台的唯一标识符，在满足UnionID下发条件的情况下会返回
}

// RspCheckEncryptedData .
type RspCheckEncryptedData struct {
	util.CommonError

	Vaild      bool `json:"vaild"`       // 是否是合法的数据
	CreateTime uint `json:"create_time"` // 加密数据生成的时间戳
}

// Code2Session 登录凭证校验。
func (auth *Auth) Code2Session(jsCode string) (result ResCode2Session, err error) {
	return auth.Code2SessionContext(context2.Background(), jsCode)
}

// Code2SessionContext 登录凭证校验。
func (auth *Auth) Code2SessionContext(ctx context2.Context, jsCode string) (result ResCode2Session, err error) {
	var response []byte
	if response, err = util.HTTPGetContext(ctx, fmt.Sprintf(code2SessionURL, auth.AppID, auth.AppSecret, jsCode)); err != nil {
		if response, err = util.HTTPGetContext(ctx, fmt.Sprintf(code2SessionReserveURL, auth.AppID, auth.AppSecret, jsCode)); err != nil {
			return
		}
	}
	if err = json.Unmarshal(response, &result); err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("Code2Session error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}

// GetPaidUnionID 用户支付完成后，获取该用户的 UnionId，无需用户授权
func (auth *Auth) GetPaidUnionID() {
	// TODO
}

// CheckEncryptedData .检查加密信息是否由微信生成（当前只支持手机号加密数据），只能检测最近3天生成的加密数据
func (auth *Auth) CheckEncryptedData(encryptedMsgHash string) (result RspCheckEncryptedData, err error) {
	return auth.CheckEncryptedDataContext(context2.Background(), encryptedMsgHash)
}

// CheckEncryptedDataContext .检查加密信息是否由微信生成（当前只支持手机号加密数据），只能检测最近3天生成的加密数据
func (auth *Auth) CheckEncryptedDataContext(ctx context2.Context, encryptedMsgHash string) (result RspCheckEncryptedData, err error) {
	var response []byte
	var (
		at string
	)
	if at, err = auth.GetAccessToken(); err != nil {
		return
	}

	// 由于GetPhoneNumberContext需要传入JSON，所以HTTPPostContext入参改为[]byte
	if response, err = util.HTTPPostContext(ctx, fmt.Sprintf(checkEncryptedDataURL, at), []byte("encrypted_msg_hash="+encryptedMsgHash), nil); err != nil {
		if response, err = util.HTTPPostContext(ctx, fmt.Sprintf(checkEncryptedDataReserveURL, at), []byte("encrypted_msg_hash="+encryptedMsgHash), nil); err != nil {
			return
		}
	}
	if err = util.DecodeWithError(response, &result, "CheckEncryptedDataAuth"); err != nil {
		return
	}
	return
}

// GetPhoneNumberResponse 新版获取用户手机号响应结构体
type GetPhoneNumberResponse struct {
	util.CommonError

	PhoneInfo PhoneInfo `json:"phone_info"`
}

// PhoneInfo 获取用户手机号内容
type PhoneInfo struct {
	PhoneNumber     string `json:"phoneNumber"`     // 用户绑定的手机号
	PurePhoneNumber string `json:"purePhoneNumber"` // 没有区号的手机号
	CountryCode     string `json:"countryCode"`     // 区号
	WaterMark       struct {
		Timestamp int64  `json:"timestamp"`
		AppID     string `json:"appid"`
	} `json:"watermark"` // 数据水印
}

// GetPhoneNumberContext 小程序通过code获取用户手机号
func (auth *Auth) GetPhoneNumberContext(ctx context2.Context, code string) (*GetPhoneNumberResponse, error) {
	var response []byte
	var (
		at  string
		err error
	)
	if at, err = auth.GetAccessToken(); err != nil {
		return nil, err
	}
	body := map[string]interface{}{
		"code": code,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	header := map[string]string{"Content-Type": "application/json;charset=utf-8"}
	if response, err = util.HTTPPostContext(ctx, fmt.Sprintf(getPhoneNumber, at), bodyBytes, header); err != nil {
		if response, err = util.HTTPPostContext(ctx, fmt.Sprintf(getPhoneNumberReserve, at), bodyBytes, header); err != nil {
			return nil, err
		}
	}

	var result GetPhoneNumberResponse
	if err = util.DecodeWithError(response, &result, "phonenumber.getPhoneNumber"); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetPhoneNumber 小程序通过code获取用户手机号
func (auth *Auth) GetPhoneNumber(code string) (*GetPhoneNumberResponse, error) {
	return auth.GetPhoneNumberContext(context2.Background(), code)
}
