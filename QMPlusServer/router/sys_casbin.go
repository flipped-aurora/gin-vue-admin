package router

import (
	"github.com/gin-gonic/gin"
	"main/controller/api"
	"main/middleware"
)

func InitCasbinRouter(Router *gin.Engine)(R gin.IRoutes) {
	BaseRouter := Router.Group("casbin").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		BaseRouter.POST("casbinPUpdata", api.CasbinPUpdata)
		BaseRouter.POST("getPolicyPathByAuthorityId", api.GetPolicyPathByAuthorityId)


	}
	return BaseRouter
}
