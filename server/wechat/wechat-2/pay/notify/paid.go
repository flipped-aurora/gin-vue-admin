package notify

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/fatih/structs"
	"github.com/spf13/cast"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

// doc: https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_7&index=8

// PaidResult 下单回调
type PaidResult struct {
	ReturnCode *string `xml:"return_code"`
	ReturnMsg  *string `xml:"return_msg"`

	AppID              *string `xml:"appid" json:"appid"`
	MchID              *string `xml:"mch_id"`
	DeviceInfo         *string `xml:"device_info"`
	NonceStr           *string `xml:"nonce_str"`
	Sign               *string `xml:"sign"`
	SignType           *string `xml:"sign_type"`
	ResultCode         *string `xml:"result_code"`
	ErrCode            *string `xml:"err_code"`
	ErrCodeDes         *string `xml:"err_code_des"`
	OpenID             *string `xml:"openid"`
	IsSubscribe        *string `xml:"is_subscribe"`
	TradeType          *string `xml:"trade_type"`
	TradeState         *string `xml:"trade_state"`
	BankType           *string `xml:"bank_type"`
	TotalFee           *int    `xml:"total_fee"`
	SettlementTotalFee *int    `xml:"settlement_total_fee"`
	FeeType            *string `xml:"fee_type"`
	CashFee            *string `xml:"cash_fee"`
	CashFeeType        *string `xml:"cash_fee_type"`
	CouponFee          *int    `xml:"coupon_fee"`
	CouponCount        *int    `xml:"coupon_count"`

	// coupon_type_$n 这里只声明 3 个，如果有更多的可以自己组合
	CouponType0 *string `xml:"coupon_type_0"`
	CouponType1 *string `xml:"coupon_type_1"`
	CouponType2 *string `xml:"coupon_type_2"`
	CouponID0   *string `xml:"coupon_id_0"`
	CouponID1   *string `xml:"coupon_id_1"`
	CouponID2   *string `xml:"coupon_id_2"`
	CouponFee0  *string `xml:"coupon_fee_0"`
	CouponFee1  *string `xml:"coupon_fee_1"`
	CouponFee2  *string `xml:"coupon_fee_2"`

	TransactionID *string `xml:"transaction_id"`
	OutTradeNo    *string `xml:"out_trade_no"`
	Attach        *string `xml:"attach"`
	TimeEnd       *string `xml:"time_end"`
}

// PaidResp 消息通知返回
type PaidResp struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
}

// PaidVerifySign 支付成功结果验签
func (notify *Notify) PaidVerifySign(notifyRes PaidResult) bool {
	// STEP1, 转换 struct 为 map，并对 map keys 做排序
	resMap := structs.Map(notifyRes)

	sortedKeys := make([]string, 0, len(resMap))
	for k := range resMap {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)

	// STEP2, 对key=value的键值对用&连接起来，略过空值 & sign
	var signStrings string
	for _, k := range sortedKeys {
		value := fmt.Sprintf("%v", cast.ToString(resMap[k]))
		if value != "" && strings.ToLower(k) != "sign" {
			signStrings = signStrings + getTagKeyName(k, &notifyRes) + "=" + value + "&"
		}
	}

	// STEP3, 在键值对的最后加上key=API_KEY
	signStrings = signStrings + "key=" + notify.Key

	// STEP4, 根据SignType计算出签名
	var signType string
	if notifyRes.SignType != nil {
		signType = *notifyRes.SignType
	}
	sign, err := util.CalculateSign(signStrings, signType, notify.Key)
	if err != nil {
		return false
	}
	if sign != *notifyRes.Sign {
		return false
	}
	return true
}

func getTagKeyName(key string, notifyRes *PaidResult) string {
	s := reflect.TypeOf(notifyRes).Elem()
	f, _ := s.FieldByName(key)
	name := f.Tag.Get("xml")
	return name
}
