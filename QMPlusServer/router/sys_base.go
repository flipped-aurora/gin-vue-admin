package router

import (
	"gin-vue-admin/controller/api"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("regist", api.Regist)
		BaseRouter.POST("login", api.Login)
		BaseRouter.GET("captcha",api.Captcha)
		BaseRouter.GET("/captcha/:id",api.CaptchaImg)
	}
	return BaseRouter
}
