package router

import (
	"github.com/gin-gonic/gin"
	"main/controller/api"
)

func InitUserRouter(Router *gin.Engine) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("regist", api.Regist)
		UserRouter.POST("login", api.Login)
		UserRouter.POST("changePassWord", api.ChangePassWord)
	}
}
