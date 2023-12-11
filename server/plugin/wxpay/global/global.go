package global

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wxpay/config"
	"github.com/go-pay/gopay/wechat/v3"
)

var GlobalConfig = new(config.WxPay)
var (
	WxClient *wechat.ClientV3
)
