package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CliUserRouter struct{}

func (s *CliUserRouter) InitCliUserRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	cliUserRouter := Router.Group("cliUser").Use(middleware.OperationRecord())
	cliUserRouterWithoutRecord := Router.Group("cliUser")
	cliUserRouterWithoutAuth := PublicRouter.Group("cliUser")
	{
		cliUserRouter.POST("createCliUser", cliUserApi.CreateCliUser)
		cliUserRouter.DELETE("deleteCliUser", cliUserApi.DeleteCliUser)
		cliUserRouter.DELETE("deleteCliUserByIds", cliUserApi.DeleteCliUserByIds)
		cliUserRouter.PUT("updateCliUser", cliUserApi.UpdateCliUser)
	}
	{
		cliUserRouterWithoutRecord.GET("findCliUser", cliUserApi.FindCliUser)
		cliUserRouterWithoutRecord.GET("getCliUserList", cliUserApi.GetCliUserList)
	}
	{
		cliUserRouterWithoutAuth.GET("getCliUserPublic", cliUserApi.GetCliUserPublic)
		cliUserRouterWithoutAuth.POST("rigester", cliUserApi.Rigester)
	}
}
