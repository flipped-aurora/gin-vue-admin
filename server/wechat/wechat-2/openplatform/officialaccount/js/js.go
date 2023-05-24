package js

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/credential"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/officialaccount/context"
	officialJs "github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/officialaccount/js"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

// Js wx jssdk
type Js struct {
	*context.Context
	credential.JsTicketHandle
}

// NewJs init
func NewJs(context *context.Context, appID string) *Js {
	js := new(Js)
	js.Context = context
	jsTicketHandle := credential.NewDefaultJsTicket(appID, credential.CacheKeyOfficialAccountPrefix, context.Cache)
	js.SetJsTicketHandle(jsTicketHandle)
	return js
}

// SetJsTicketHandle 自定义js ticket取值方式
func (js *Js) SetJsTicketHandle(ticketHandle credential.JsTicketHandle) {
	js.JsTicketHandle = ticketHandle
}

// GetConfig 第三方平台 - 获取jssdk需要的配置参数
// uri 为当前网页地址
func (js *Js) GetConfig(uri, appid string) (config *officialJs.Config, err error) {
	config = new(officialJs.Config)
	var accessToken string
	accessToken, err = js.GetAccessToken()
	if err != nil {
		return
	}
	var ticketStr string
	ticketStr, err = js.GetTicket(accessToken)
	if err != nil {
		return
	}

	nonceStr := util.RandomStr(16)
	timestamp := util.GetCurrTS()
	str := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s", ticketStr, nonceStr, timestamp, uri)
	sigStr := util.Signature(str)

	config.AppID = appid
	config.NonceStr = nonceStr
	config.Timestamp = timestamp
	config.Signature = sigStr
	return
}
