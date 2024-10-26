package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CliOrderRouter struct{}

// InitCliOrderRouter 初始化 订单详情 路由信息
func (s *CliOrderRouter) InitCliOrderRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	cliOrderRouter := Router.Group("cliOrder").Use(middleware.OperationRecord())
	cliOrderRouterWithoutRecord := Router.Group("cliOrder")
	cliOrderRouterWithoutAuth := PublicRouter.Group("cliOrder")
	{
		cliOrderRouter.POST("createCliOrder", cliOrderApi.CreateCliOrder)             // 新建订单详情
		cliOrderRouter.DELETE("deleteCliOrder", cliOrderApi.DeleteCliOrder)           // 删除订单详情
		cliOrderRouter.DELETE("deleteCliOrderByIds", cliOrderApi.DeleteCliOrderByIds) // 批量删除订单详情
		cliOrderRouter.PUT("updateCliOrder", cliOrderApi.UpdateCliOrder)              // 更新订单详情
	}
	{
		cliOrderRouterWithoutRecord.GET("findCliOrder", cliOrderApi.FindCliOrder)       // 根据ID获取订单详情
		cliOrderRouterWithoutRecord.GET("getCliOrderList", cliOrderApi.GetCliOrderList) // 获取订单详情列表
	}
	{
		cliOrderRouterWithoutAuth.GET("getCliOrderPublic", cliOrderApi.GetCliOrderPublic) // 订单详情开放接口
	}
}
