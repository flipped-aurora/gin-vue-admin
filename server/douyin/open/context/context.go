package context

import (
	"github.com/flipped-aurora/gin-vue-admin/server/douyin/open/config"
	"github.com/flipped-aurora/gin-vue-admin/server/douyin/open/credential"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
