package order

import (
	"encoding/xml"
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/pay/notify"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

var queryGateway = "https://api.mch.weixin.qq.com/pay/orderquery"

// QueryParams 传入的参数
type QueryParams struct {
	OutTradeNo    string // 商户订单号
	SignType      string // 签名类型
	TransactionID string // 微信订单号
}

// queryRequest 接口请求参数
type queryRequest struct {
	AppID         string `xml:"appid"`               // 公众账号ID
	MchID         string `xml:"mch_id"`              // 商户号
	NonceStr      string `xml:"nonce_str"`           // 随机字符串
	Sign          string `xml:"sign"`                // 签名
	SignType      string `xml:"sign_type,omitempty"` // 签名类型
	TransactionID string `xml:"transaction_id"`      // 微信订单号
	OutTradeNo    string `xml:"out_trade_no"`        // 商户订单号
}

// QueryOrder 查询订单
func (o *Order) QueryOrder(p *QueryParams) (paidResult notify.PaidResult, err error) {
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
	params["transaction_id"] = p.TransactionID

	sign, err := util.ParamSign(params, o.Key)
	if err != nil {
		return
	}
	request := queryRequest{
		AppID:         o.AppID,
		MchID:         o.MchID,
		NonceStr:      nonceStr,
		Sign:          sign,
		OutTradeNo:    p.OutTradeNo,
		TransactionID: p.TransactionID,
		SignType:      p.SignType,
	}

	rawRet, err := util.PostXML(queryGateway, request)
	if err != nil {
		return
	}

	err = xml.Unmarshal(rawRet, &paidResult)
	if err != nil {
		return
	}

	if *paidResult.ReturnCode == SUCCESS {
		// query success
		if *paidResult.ResultCode == SUCCESS {
			err = nil
			return
		}
		err = errors.New(*paidResult.ErrCode + *paidResult.ErrCodeDes)
		return
	}
	err = errors.New("[msg : xmlUnmarshalError] [rawReturn : " + string(rawRet) + "] [sign : " + sign + "]")
	return
}
