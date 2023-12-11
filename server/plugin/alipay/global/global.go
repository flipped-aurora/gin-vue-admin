package global

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/alipay/config"
	"github.com/go-pay/gopay/alipay"
)

var GlobalConfig = new(config.Alipay)

var (
	Client *alipay.Client
)
