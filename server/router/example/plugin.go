package example

import (
	"github.com/flipped-aurora/gin-vue-admin/utils/plugin"

	"github.com/gin-gonic/gin"
)

var Plugin []plugin.Plugin = []plugin.Plugin{&PluginExample}

var PluginExample = pluginExample{}

type pluginExample struct {
}

func (*pluginExample) Register(group *gin.RouterGroup) {
	group.GET("hello", func(context *gin.Context) {
		context.JSON(200, "hello world")
	})
}

func (*pluginExample) RouterPath() string {
	return "group"
}

func PluginInit(group *gin.RouterGroup) {
	for i := range Plugin {
		PluginGroup := group.Group(Plugin[i].RouterPath())
		Plugin[i].Register(PluginGroup)
	}
}
