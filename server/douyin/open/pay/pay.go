package pay

import "github.com/flipped-aurora/gin-vue-admin/server/douyin/open/context"

// Pay 支付服务
type Pay struct {
	*context.Context
}

// NewPay .
func NewPay(context *context.Context) *Pay {
	pay := new(Pay)
	pay.Context = context
	return pay
}
