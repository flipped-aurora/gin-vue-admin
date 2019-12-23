package router

import (
	"gin-vue-admin/controller/api"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitCasbinRouter(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("casbin").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		BaseRouter.POST("casbinPUpdata", api.CasbinPUpdata)
		BaseRouter.POST("getPolicyPathByAuthorityId", api.GetPolicyPathByAuthorityId)

	}
}
