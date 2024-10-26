package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CliLoadRouter struct{}

func (s *CliLoadRouter) InitCliLoadRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	cliLoadRouter := Router.Group("cliLoad").Use(middleware.OperationRecord())
	cliLoadRouterWithoutRecord := Router.Group("cliLoad")
	cliLoadRouterWithoutAuth := PublicRouter.Group("cliLoad")
	{
		cliLoadRouter.POST("createCliLoad", cliLoadApi.CreateCliLoad)
		cliLoadRouter.DELETE("deleteCliLoad", cliLoadApi.DeleteCliLoad)
		cliLoadRouter.DELETE("deleteCliLoadByIds", cliLoadApi.DeleteCliLoadByIds)
		cliLoadRouter.PUT("updateCliLoad", cliLoadApi.UpdateCliLoad)
	}
	{
		cliLoadRouterWithoutRecord.GET("findCliLoad", cliLoadApi.FindCliLoad)
		cliLoadRouterWithoutRecord.GET("getCliLoadList", cliLoadApi.GetCliLoadList)
	}
	{
		cliLoadRouterWithoutAuth.GET("getCliLoadPublic", cliLoadApi.GetCliLoadPublic)
		cliLoadRouterWithoutAuth.POST("login", cliLoadApi.Login)
	}
}
