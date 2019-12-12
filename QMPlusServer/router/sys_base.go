package router

import (
	"github.com/gin-gonic/gin"
	"main/controller/api"
)

func InitBaseRouter(Router *gin.Engine)(R gin.IRoutes) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("regist", api.Regist)
		BaseRouter.POST("login", api.Login)
	}
	return BaseRouter
}
