package msgaudit

import (
	"fmt"
)

// 返回码	错误说明
// 10000	参数错误，请求参数错误
// 10001	网络错误，网络请求错误
// 10002	数据解析失败
// 10003	系统失败
// 10004	密钥错误导致加密失败
// 10005	fileid错误
// 10006	解密失败
// 10007 找不到消息加密版本的私钥，需要重新传入私钥对
// 10008 解析encrypt_key出错
// 10009 ip非法
// 10010 数据过期
// 10011	证书错误
const (
	SDKErrMsg               = "sdk failed"
	SDKParamsErrMsg         = "参数错误，请求参数错误"
	SDKNetworkErrMsg        = "网络错误，网络请求错误"
	SDKParseErrMsg          = "数据解析失败"
	SDKSystemErrMsg         = "系统失败"
	SDKSecretErrMsg         = "密钥错误导致加密失败"
	SDKFileIDErrMsg         = "fileid错误"
	SDKDecryptErrMsg        = "解密失败"
	SDKSecretMissErrMsg     = "找不到消息加密版本的私钥，需要重新传入私钥对"
	SDKEncryptKeyErrMsg     = "解析encrypt_key出错"
	SDKIPNotWhiteListErrMsg = "ip非法"
	SDKDataExpiredErrMsg    = "数据过期"
	SDKTokenExpiredErrMsg   = "证书过期"
)

// Error 错误
type Error struct {
	ErrCode int    `json:"errcode,omitempty"`
	ErrMsg  string `json:"errmsg,omitempty"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%d:%s", e.ErrCode, e.ErrMsg)
}

// NewSDKErr 初始化新的SDK错误
func NewSDKErr(code int) Error {
	msg := ""
	switch code {
	case 10000:
		msg = SDKParamsErrMsg
	case 10001:
		msg = SDKNetworkErrMsg
	case 10002:
		msg = SDKParseErrMsg
	case 10003:
		msg = SDKSystemErrMsg
	case 10004:
		msg = SDKSecretErrMsg
	case 10005:
		msg = SDKFileIDErrMsg
	case 10006:
		msg = SDKDecryptErrMsg
	case 10007:
		msg = SDKSecretMissErrMsg
	case 10008:
		msg = SDKEncryptKeyErrMsg
	case 10009:
		msg = SDKIPNotWhiteListErrMsg
	case 10010:
		msg = SDKDataExpiredErrMsg
	case 10011:
		msg = SDKTokenExpiredErrMsg
	}
	return Error{ErrCode: code, ErrMsg: msg}
}
