package order

import (
	"encoding/xml"
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

// https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_3
var closeGateway = "https://api.mch.weixin.qq.com/pay/closeorder"

// CloseParams 传入的参数
type CloseParams struct {
	OutTradeNo string // 商户订单号
	SignType   string // 签名类型
}

// closeRequest 接口请求参数
type closeRequest struct {
	AppID      string `xml:"appid"`               // 公众账号ID
	MchID      string `xml:"mch_id"`              // 商户号
	NonceStr   string `xml:"nonce_str"`           // 随机字符串
	Sign       string `xml:"sign"`                // 签名
	SignType   string `xml:"sign_type,omitempty"` // 签名类型
	OutTradeNo string `xml:"out_trade_no"`        // 商户订单号
}

// CloseResult 关闭订单返回结果
type CloseResult struct {
	ReturnCode *string `xml:"return_code"`
	ReturnMsg  *string `xml:"return_msg"`

	AppID      *string `xml:"appid" json:"appid"`
	MchID      *string `xml:"mch_id"`
	NonceStr   *string `xml:"nonce_str"`
	Sign       *string `xml:"sign"`
	ResultCode *string `xml:"result_code"`
	ResultMsg  *string `xml:"result_msg"`
	ErrCode    *string `xml:"err_code"`
	ErrCodeDes *string `xml:"err_code_des"`
}

// CloseOrder 关闭订单
func (o *Order) CloseOrder(p *CloseParams) (closeResult CloseResult, err error) {
	nonceStr := util.RandomStr(32)
	// 签名类型
	if p.SignType == "" {
		p.SignType = "MD5"
	}

	params := make(map[string]string)
	params["appid"] = o.AppID
	params["mch_id"] = o.MchID
	params["nonce_str"] = nonceStr
	params["out_trade_no"] = p.OutTradeNo
	params["sign_type"] = p.SignType

	var (
		sign   string
		rawRet []byte
	)

	sign, err = util.ParamSign(params, o.Key)
	if err != nil {
		return
	}
	request := closeRequest{
		AppID:      o.AppID,
		MchID:      o.MchID,
		NonceStr:   nonceStr,
		Sign:       sign,
		OutTradeNo: p.OutTradeNo,
		SignType:   p.SignType,
	}

	rawRet, err = util.PostXML(closeGateway, request)
	if err != nil {
		return
	}

	err = xml.Unmarshal(rawRet, &closeResult)
	if err != nil {
		return
	}

	if *closeResult.ReturnCode == SUCCESS {
		// close success
		if *closeResult.ResultCode == SUCCESS {
			err = nil
			return
		}
		err = errors.New(*closeResult.ErrCode + *closeResult.ErrCodeDes)
		return
	}
	err = errors.New("[msg : xmlUnmarshalError] [rawReturn : " + string(rawRet) + "] [sign : " + sign + "]")
	return
}
