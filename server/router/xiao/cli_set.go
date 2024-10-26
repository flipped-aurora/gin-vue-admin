package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CliSetRouter struct{}

// InitCliSetRouter 初始化 结算设置 路由信息
func (s *CliSetRouter) InitCliSetRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	clisetRouter := Router.Group("cliset").Use(middleware.OperationRecord())
	clisetRouterWithoutRecord := Router.Group("cliset")
	clisetRouterWithoutAuth := PublicRouter.Group("cliset")
	{
		clisetRouter.POST("createCliSet", clisetApi.CreateCliSet)             // 新建结算设置
		clisetRouter.DELETE("deleteCliSet", clisetApi.DeleteCliSet)           // 删除结算设置
		clisetRouter.DELETE("deleteCliSetByIds", clisetApi.DeleteCliSetByIds) // 批量删除结算设置
		clisetRouter.PUT("updateCliSet", clisetApi.UpdateCliSet)              // 更新结算设置
	}
	{
		clisetRouterWithoutRecord.GET("findCliSet", clisetApi.FindCliSet)       // 根据ID获取结算设置
		clisetRouterWithoutRecord.GET("getCliSetList", clisetApi.GetCliSetList) // 获取结算设置列表
	}
	{
		clisetRouterWithoutAuth.GET("getCliSetPublic", clisetApi.GetCliSetPublic) // 结算设置开放接口
	}
}
