package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitAutoCodeRouter(Router *gin.RouterGroup) {
	AutoCodeRouter := Router.Group("autoCode").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		AutoCodeRouter.POST("createTemp", v1.CreateTemp) //创建自动化代码
	}
}
