package service

import (
	"encoding/json"
	"github.com/go-pay/gopay"
	"time"
)

// 设置微信Native预下单body
func WxNativeBody(appid, mchid, tradeno, goodsname, notifyurl, attach string, goodsprice int64) gopay.BodyMap {
	expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)
	bm := make(gopay.BodyMap)
	bm.Set("appid", appid).
		Set("mchid", mchid).
		Set("out_trade_no", tradeno).
		Set("description", goodsname).
		Set("time_expire", expire).
		Set("notify_url", notifyurl).
		Set("attach", attach).
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", goodsprice).
				Set("currency", "CNY")
		})
	return bm
}

// 设置微信JsApi预下单body
func WxJsapiBody(appid, mchid, tradeno, goodsname, notifyurl, attach, openid string, goodsprice int64) gopay.BodyMap {
	expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)
	bm := make(gopay.BodyMap)
	bm.Set("appid", appid).
		Set("mchid", mchid).
		Set("description", goodsname).
		Set("out_trade_no", tradeno).
		Set("time_expire", expire).
		Set("notify_url", notifyurl).
		Set("attach", attach).
		Set("attach", attach).
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", goodsprice).
				Set("currency", "CNY")
		}).
		SetBodyMap("payer", func(bm gopay.BodyMap) {
			bm.Set("openid", openid)
		})
	return bm
}

// 字符串转json
func StrToJson(s, m string) (r string) {
	var jsonMap map[string]interface{}
	err := json.Unmarshal([]byte(s), &jsonMap)
	if err != nil {
		return ""
	}
	r = jsonMap[m].(string)
	return r

}
