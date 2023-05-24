package context

import (
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/credential"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
