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
		BaseRouter.POST("captcha", api.Captcha)
		BaseRouter.GET("captcha/:captchaId", api.CaptchaImg)
		BaseRouter.POST("coffee", api.GetCoffeeList)
		BaseRouter.POST("coffeetype", api.GetCoffeeTypeList)
		BaseRouter.POST("cuslogin", api.CusLogin)
		BaseRouter.POST("coffeebycode", api.GetCoffeeListByCode)
		BaseRouter.POST("/getcoffeetype", api.GetCoffeeByCode)
		BaseRouter.POST("getcoffeebyid", api.GetCoffeeByUUID)
		BaseRouter.POST("/phone", api.SendPhoneCaptcha)
		BaseRouter.POST("/cusRegister", api.CusRegist)
	}
	return BaseRouter
}
