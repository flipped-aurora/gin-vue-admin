package pprof

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/pprof/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/pprof/router"
	"github.com/gin-gonic/gin"
)

const (
	// default url prefix of pprof
	DefaultPrefix = "debug/pprof"
)

type pprofPlugin struct{}

func CreatePprofPlug(Prefix string) *pprofPlugin {
	global.PprofConfig.Prefix = Prefix
	return &pprofPlugin{}
}

func (*pprofPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitPprofRouter(group)
}

func (*pprofPlugin) RouterPath() string {
	prefix := DefaultPrefix
	if len(global.PprofConfig.Prefix) > 0 {
		prefix = global.PprofConfig.Prefix
	}
	return prefix
}
