package router

import (
	"github.com/gin-gonic/gin"
	"main/controller/api"
	"main/middleware"
)

func InitApiRouter(Router *gin.Engine) {
	ApiRouter := Router.Group("api").Use(middleware.JWTAuth())
	{
		ApiRouter.POST("createApi", api.CreateApi)
		ApiRouter.POST("deleteApi", api.DeleteApi)
		ApiRouter.POST("setAuthAndPath",api.SetAuthAndPath)
		ApiRouter.POST("getApiList",api.GetApiList)
	}
}
