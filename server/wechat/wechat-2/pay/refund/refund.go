package refund

import (
	"encoding/xml"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/pay/config"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

var refundGateway = "https://api.mch.weixin.qq.com/secapi/pay/refund"

// Refund struct extends context
type Refund struct {
	*config.Config
}

// NewRefund return an instance of refund package
func NewRefund(cfg *config.Config) *Refund {
	refund := Refund{cfg}
	return &refund
}

// Params 调用参数
type Params struct {
	TransactionID string
	OutRefundNo   string
	OutTradeNo    string
	TotalFee      string
	RefundFee     string
	RefundDesc    string
	RootCa        string // ca证书
	NotifyURL     string
	SignType      string
}

// request 接口请求参数
type request struct {
	AppID         string `xml:"appid"`
	MchID         string `xml:"mch_id"`
	NonceStr      string `xml:"nonce_str"`
	Sign          string `xml:"sign"`
	SignType      string `xml:"sign_type,omitempty"`
	TransactionID string `xml:"transaction_id,omitempty"`
	OutTradeNo    string `xml:"out_trade_no,omitempty"`
	OutRefundNo   string `xml:"out_refund_no"`
	TotalFee      string `xml:"total_fee"`
	RefundFee     string `xml:"refund_fee"`
	RefundDesc    string `xml:"refund_desc,omitempty"`
	NotifyURL     string `xml:"notify_url,omitempty"`
}

// Response 接口返回
type Response struct {
	ReturnCode          string `xml:"return_code"`
	ReturnMsg           string `xml:"return_msg"`
	AppID               string `xml:"appid,omitempty"`
	MchID               string `xml:"mch_id,omitempty"`
	NonceStr            string `xml:"nonce_str,omitempty"`
	Sign                string `xml:"sign,omitempty"`
	ResultCode          string `xml:"result_code,omitempty"`
	ErrCode             string `xml:"err_code,omitempty"`
	ErrCodeDes          string `xml:"err_code_des,omitempty"`
	TransactionID       string `xml:"transaction_id,omitempty"`
	OutTradeNo          string `xml:"out_trade_no,omitempty"`
	OutRefundNo         string `xml:"out_refund_no,omitempty"`
	RefundID            string `xml:"refund_id,omitempty"`
	RefundFee           string `xml:"refund_fee,omitempty"`
	SettlementRefundFee string `xml:"settlement_refund_fee,omitempty"`
	TotalFee            string `xml:"total_fee,omitempty"`
	SettlementTotalFee  string `xml:"settlement_total_fee,omitempty"`
	FeeType             string `xml:"fee_type,omitempty"`
	CashFee             string `xml:"cash_fee,omitempty"`
	CashFeeType         string `xml:"cash_fee_type,omitempty"`
}

// Refund 退款申请
func (refund *Refund) Refund(p *Params) (rsp Response, err error) {
	param := refund.GetSignParam(p)

	sign, err := util.ParamSign(param, refund.Key)
	if err != nil {
		return
	}

	req := request{
		AppID:       param["appid"],
		MchID:       param["mch_id"],
		NonceStr:    param["nonce_str"],
		Sign:        sign,
		SignType:    param["sign_type"],
		OutRefundNo: param["out_refund_no"],
		TotalFee:    param["total_fee"],
		RefundFee:   param["refund_fee"],
		RefundDesc:  param["refund_desc"],
		NotifyURL:   param["notify_url"],
	}
	if p.OutTradeNo != "" {
		req.OutTradeNo = p.OutTradeNo
	}
	if p.TransactionID != "" {
		req.TransactionID = p.TransactionID
	}

	rawRet, err := util.PostXMLWithTLS(refundGateway, req, p.RootCa, refund.MchID)
	if err != nil {
		return
	}
	err = xml.Unmarshal(rawRet, &rsp)
	if err != nil {
		return
	}
	if rsp.ReturnCode == "SUCCESS" {
		if rsp.ResultCode == "SUCCESS" {
			err = nil
			return
		}
		err = fmt.Errorf("refund error, errcode=%s,errmsg=%s", rsp.ErrCode, rsp.ErrCodeDes)
		return
	}
	err = fmt.Errorf("[msg : xmlUnmarshalError] [rawReturn : %s] [sign : %s]", string(rawRet), sign)
	return
}

// GetSignParam 获取签名的参数
func (refund *Refund) GetSignParam(p *Params) (param map[string]string) {
	nonceStr := util.RandomStr(32)
	param = make(map[string]string)
	param["appid"] = refund.AppID
	param["mch_id"] = refund.MchID
	param["nonce_str"] = nonceStr
	param["out_refund_no"] = p.OutRefundNo
	param["refund_desc"] = p.RefundDesc
	param["refund_fee"] = p.RefundFee
	param["total_fee"] = p.TotalFee

	if p.SignType == "" {
		param["sign_type"] = util.SignTypeMD5
	}
	if p.OutTradeNo != "" {
		param["out_trade_no"] = p.OutTradeNo
	}
	if p.TransactionID != "" {
		param["transaction_id"] = p.TransactionID
	}
	if p.NotifyURL != "" {
		param["notify_url"] = p.NotifyURL
	}

	return param
}
