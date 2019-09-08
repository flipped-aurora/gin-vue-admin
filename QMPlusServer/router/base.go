package router

import (
	"github.com/gin-gonic/gin"
	"main/controller/api"
)

func InitBaseRouter(Router *gin.Engine) {
	UserRouter := Router.Group("base")
	{
		UserRouter.POST("regist", api.Regist)
		UserRouter.POST("login", api.Login)
	}
}
