package router

import (
	"gin-vue-admin/controller/api"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSystemRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("system").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		UserRouter.POST("getSystemConfig", api.GetSystemConfig)     // 获取配置文件内容
		UserRouter.POST("setSystemConfig", api.SetSystemConfig)     // 设置配置文件内容
	}
}
