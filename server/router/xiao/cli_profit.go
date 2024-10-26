package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CliProfitRouter struct{}

// InitCliProfitRouter 初始化 结算详情 路由信息
func (s *CliProfitRouter) InitCliProfitRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	cliprofitRouter := Router.Group("cliprofit").Use(middleware.OperationRecord())
	cliprofitRouterWithoutRecord := Router.Group("cliprofit")
	cliprofitRouterWithoutAuth := PublicRouter.Group("cliprofit")
	{
		cliprofitRouter.POST("createCliProfit", cliprofitApi.CreateCliProfit)             // 新建结算详情
		cliprofitRouter.DELETE("deleteCliProfit", cliprofitApi.DeleteCliProfit)           // 删除结算详情
		cliprofitRouter.DELETE("deleteCliProfitByIds", cliprofitApi.DeleteCliProfitByIds) // 批量删除结算详情
		cliprofitRouter.PUT("updateCliProfit", cliprofitApi.UpdateCliProfit)              // 更新结算详情
	}
	{
		cliprofitRouterWithoutRecord.GET("findCliProfit", cliprofitApi.FindCliProfit)       // 根据ID获取结算详情
		cliprofitRouterWithoutRecord.GET("getCliProfitList", cliprofitApi.GetCliProfitList) // 获取结算详情列表
	}
	{
		cliprofitRouterWithoutAuth.GET("getCliProfitPublic", cliprofitApi.GetCliProfitPublic) // 结算详情开放接口
	}
}
