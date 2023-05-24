package notify

import (
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/pay/config"
)

// Notify 回调
type Notify struct {
	*config.Config
}

// NewNotify new
func NewNotify(cfg *config.Config) *Notify {
	return &Notify{cfg}
}
