package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CliTreeRouter struct{}

// InitCliTreeRouter 初始化 用户关系表 路由信息
func (s *CliTreeRouter) InitCliTreeRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	cliTreeRouter := Router.Group("cliTree").Use(middleware.OperationRecord())
	cliTreeRouterWithoutRecord := Router.Group("cliTree")
	cliTreeRouterWithoutAuth := PublicRouter.Group("cliTree")
	{
		cliTreeRouter.POST("createCliTree", cliTreeApi.CreateCliTree)             // 新建用户关系表
		cliTreeRouter.DELETE("deleteCliTree", cliTreeApi.DeleteCliTree)           // 删除用户关系表
		cliTreeRouter.DELETE("deleteCliTreeByIds", cliTreeApi.DeleteCliTreeByIds) // 批量删除用户关系表
		cliTreeRouter.PUT("updateCliTree", cliTreeApi.UpdateCliTree)              // 更新用户关系表
	}
	{
		cliTreeRouterWithoutRecord.GET("findCliTree", cliTreeApi.FindCliTree)       // 根据ID获取用户关系表
		cliTreeRouterWithoutRecord.GET("getCliTreeList", cliTreeApi.GetCliTreeList) // 获取用户关系表列表
	}
	{
		cliTreeRouterWithoutAuth.GET("getCliTreePublic", cliTreeApi.GetCliTreePublic) // 用户关系表开放接口
	}
}
