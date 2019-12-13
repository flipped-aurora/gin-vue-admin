package router

import (
	"github.com/gin-gonic/gin"
	"main/controller/api"
	"main/middleware"
)

func InitCasbinRouter(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("casbin").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		BaseRouter.POST("casbinPUpdata", api.CasbinPUpdata)
		BaseRouter.POST("getPolicyPathByAuthorityId", api.GetPolicyPathByAuthorityId)

	}
}
