package router

import (
	"github.com/gin-gonic/gin"
	"main/controller/api"
	"main/middleware"
)

func InitApiRouter(Router *gin.Engine) {
	UserRouter := Router.Group("api").Use(middleware.JWTAuth())
	{
		UserRouter.POST("createApi", api.CreateApi)
		UserRouter.POST("deleteApi", api.DeleteApi)
	}
}
