package context

import (
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/openplatform/config"
)

// Context struct
type Context struct {
	*config.Config
}
