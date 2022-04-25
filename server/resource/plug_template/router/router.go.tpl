package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/{{ .Snake}}/api"
	"github.com/gin-gonic/gin"
)

type {{ .PlugName}}Router struct {
}

func (s *{{ .PlugName}}Router) Init{{ .PlugName}}Router(Router *gin.RouterGroup) {
	plugRouter := Router
	plugApi := api.ApiGroupApp.{{ .PlugName}}Api
	{
		plugRouter.POST("routerName", plugApi.ApiName)
	}
}
