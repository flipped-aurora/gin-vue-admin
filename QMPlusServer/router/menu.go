package router

import (
	"github.com/gin-gonic/gin"
	"main/controller/api"
	"main/middleware"
)

func InitMenuRouter(Router *gin.Engine) {
	MenuRouter := Router.Group("menu").Use(middleware.JWTAuth())
	{
		MenuRouter.POST("getMenu", api.GetMenu)
	}
}
