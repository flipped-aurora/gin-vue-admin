package notify

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/xml"
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

// reference: https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_16&index=10

// RefundedResult 退款回调
type RefundedResult struct {
	ReturnCode *string `xml:"return_code"`
	ReturnMsg  *string `xml:"return_msg"`

	AppID    *string `xml:"appid"`
	MchID    *string `xml:"mch_id"`
	NonceStr *string `xml:"nonce_str"`
	ReqInfo  *string `xml:"req_info"`
}

// RefundedReqInfo 退款结果（明文）
type RefundedReqInfo struct {
	TransactionID       *string `xml:"transaction_id"`
	OutTradeNO          *string `xml:"out_trade_no"`
	RefundID            *string `xml:"refund_id"`
	OutRefundNO         *string `xml:"out_refund_no"`
	TotalFee            *int    `xml:"total_fee"`
	SettlementTotalFee  *int    `xml:"settlement_total_fee"`
	RefundFee           *int    `xml:"refund_fee"`
	SettlementRefundFee *int    `xml:"settlement_refund_fee"`
	RefundStatus        *string `xml:"refund_status"`
	SuccessTime         *string `xml:"success_time"`
	RefundRecvAccount   *string `xml:"refund_recv_accout"`
	RefundAccount       *string `xml:"refund_account"`
	RefundRequestSource *string `xml:"refund_request_source"`
}

// RefundedResp 消息通知返回
type RefundedResp struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
}

// DecryptReqInfo 对退款结果进行解密
func (notify *Notify) DecryptReqInfo(result *RefundedResult) (*RefundedReqInfo, error) {
	if result == nil || result.ReqInfo == nil {
		return nil, errors.New("empty refunded_result or req_info")
	}

	base64Decode, err := base64.StdEncoding.DecodeString(*result.ReqInfo)
	if err != nil {
		return nil, err
	}

	hash := md5.New()
	if _, err = hash.Write([]byte(notify.Key)); err != nil {
		return nil, err
	}
	md5APIKey := hex.EncodeToString(hash.Sum(nil))

	data, err := util.AesECBDecrypt(base64Decode, []byte(md5APIKey))
	if err != nil {
		return nil, err
	}

	res := &RefundedReqInfo{}
	if err = xml.Unmarshal(data, res); err != nil {
		return nil, err
	}
	return res, nil
}
