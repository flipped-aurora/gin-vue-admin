package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CliMainorderRouter struct{}

func (s *CliMainorderRouter) InitCliMainorderRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	climainorderRouter := Router.Group("climainorder").Use(middleware.OperationRecord())
	climainorderRouterWithoutRecord := Router.Group("climainorder")
	myjwtRouter := PublicRouter.Group("climainorder").Use(middleware.JWTAuth())
	climainorderRouterWithoutAuth := PublicRouter.Group("climainorder")
	{
		climainorderRouter.POST("createCliMainorder", climainorderApi.CreateCliMainorder)
		climainorderRouter.DELETE("deleteCliMainorder", climainorderApi.DeleteCliMainorder)
		climainorderRouter.DELETE("deleteCliMainorderByIds", climainorderApi.DeleteCliMainorderByIds)
		climainorderRouter.PUT("updateCliMainorder", climainorderApi.UpdateCliMainorder)

	}
	{
		climainorderRouterWithoutRecord.GET("findCliMainorder", climainorderApi.FindCliMainorder)
		climainorderRouterWithoutRecord.GET("getCliMainorderList", climainorderApi.GetCliMainorderList)
	}
	{
		myjwtRouter.POST("buy", climainorderApi.Buy)
	}
	{
		climainorderRouterWithoutAuth.GET("getCliMainorderPublic", climainorderApi.GetCliMainorderPublic)
	}
}
