package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/notify/api"
	"github.com/gin-gonic/gin"
)

type NotifyRouter struct {
}

func (s *NotifyRouter) InitRouter(Router *gin.RouterGroup) {
	emailRouter := Router.Use(middleware.OperationRecord())
	var Controller = api.ApiGroupApp.Api.NotifyController
	{
		emailRouter.POST("dingding", Controller)
	}
}
