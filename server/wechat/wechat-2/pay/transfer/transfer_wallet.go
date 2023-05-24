package transfer

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/pay/config"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

// walletTransferGateway 付款到零钱
// https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_2
var walletTransferGateway = "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers"

// Transfer struct extends context
type Transfer struct {
	*config.Config
}

// NewTransfer return an instance of Transfer package
func NewTransfer(cfg *config.Config) *Transfer {
	return &Transfer{cfg}
}

// Params 调用参数
type Params struct {
	DeviceInfo     string
	PartnerTradeNo string
	OpenID         string
	CheckName      bool
	ReUserName     string
	Amount         int
	Desc           string
	SpbillCreateIP string
	RootCa         string // ca证书
}

// request 接口请求参数
type request struct {
	AppID          string `xml:"mch_appid"`
	MchID          string `xml:"mchid"`
	NonceStr       string `xml:"nonce_str"`
	Sign           string `xml:"sign"`
	DeviceInfo     string `xml:"device_info,omitempty"`
	PartnerTradeNo string `xml:"partner_trade_no"`
	OpenID         string `xml:"openid"`
	CheckName      string `xml:"check_name"`
	ReUserName     string `xml:"re_user_name,omitempty"`
	Amount         int    `xml:"amount"`
	Desc           string `xml:"desc"`
	SpbillCreateIP string `xml:"spbill_create_ip,omitempty"`
}

// Response 接口返回
type Response struct {
	ReturnCode     string `xml:"return_code"`
	ReturnMsg      string `xml:"return_msg"`
	AppID          string `xml:"appid,omitempty"`
	MchID          string `xml:"mch_id,omitempty"`
	DeviceInfo     string `xml:"device_info,omitempty"`
	NonceStr       string `xml:"nonce_str,omitempty"`
	ResultCode     string `xml:"result_code,omitempty"`
	ErrCode        string `xml:"err_code,omitempty"`
	ErrCodeDes     string `xml:"err_code_des,omitempty"`
	PartnerTradeNo string `xml:"partner_trade_no"`
	PaymentNo      string `xml:"payment_no"`
	PaymentTime    string `xml:"payment_time"`
}

// WalletTransfer 付款到零钱
func (transfer *Transfer) WalletTransfer(p *Params) (rsp *Response, err error) {
	nonceStr := util.RandomStr(32)
	param := make(map[string]string)
	param["mch_appid"] = transfer.AppID
	param["mchid"] = transfer.MchID
	param["nonce_str"] = nonceStr
	param["partner_trade_no"] = p.PartnerTradeNo
	param["openid"] = p.OpenID
	param["amount"] = strconv.Itoa(p.Amount)
	param["desc"] = p.Desc
	if p.DeviceInfo != "" {
		param["device_info"] = p.DeviceInfo
	}
	param["check_name"] = "NO_CHECK"
	if p.CheckName {
		param["check_name"] = "FORCE_CHECK"
		param["re_user_name"] = p.ReUserName
	}
	if p.SpbillCreateIP != "" {
		param["spbill_create_ip"] = p.SpbillCreateIP
	}

	sign, err := util.ParamSign(param, transfer.Key)
	if err != nil {
		return
	}

	req := request{
		AppID:          transfer.AppID,
		MchID:          transfer.MchID,
		NonceStr:       nonceStr,
		Sign:           sign,
		DeviceInfo:     p.DeviceInfo,
		PartnerTradeNo: p.PartnerTradeNo,
		OpenID:         p.OpenID,
		Amount:         p.Amount,
		Desc:           p.Desc,
		SpbillCreateIP: p.SpbillCreateIP,
	}
	req.CheckName = "NO_CHECK"
	if p.CheckName {
		req.CheckName = "FORCE_CHECK"
		req.ReUserName = p.ReUserName
	}
	rawRet, err := util.PostXMLWithTLS(walletTransferGateway, req, p.RootCa, transfer.MchID)
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
		err = fmt.Errorf("transfer error, errcode=%s,errmsg=%s", rsp.ErrCode, rsp.ErrCodeDes)
		return
	}
	err = fmt.Errorf("[msg : xmlUnmarshalError] [rawReturn : %s] [sign : %s]", string(rawRet), sign)
	return
}
