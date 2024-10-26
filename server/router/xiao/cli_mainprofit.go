package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CliMainprofitRouter struct{}

// InitCliMainprofitRouter 初始化 结算总表 路由信息
func (s *CliMainprofitRouter) InitCliMainprofitRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	climainprofitRouter := Router.Group("climainprofit").Use(middleware.OperationRecord())
	climainprofitRouterWithoutRecord := Router.Group("climainprofit")
	climainprofitRouterWithoutAuth := PublicRouter.Group("climainprofit")
	{
		climainprofitRouter.POST("createCliMainprofit", climainprofitApi.CreateCliMainprofit)             // 新建结算总表
		climainprofitRouter.DELETE("deleteCliMainprofit", climainprofitApi.DeleteCliMainprofit)           // 删除结算总表
		climainprofitRouter.DELETE("deleteCliMainprofitByIds", climainprofitApi.DeleteCliMainprofitByIds) // 批量删除结算总表
		climainprofitRouter.PUT("updateCliMainprofit", climainprofitApi.UpdateCliMainprofit)              // 更新结算总表
	}
	{
		climainprofitRouterWithoutRecord.GET("findCliMainprofit", climainprofitApi.FindCliMainprofit)       // 根据ID获取结算总表
		climainprofitRouterWithoutRecord.GET("getCliMainprofitList", climainprofitApi.GetCliMainprofitList) // 获取结算总表列表
	}
	{
		climainprofitRouterWithoutAuth.GET("getCliMainprofitPublic", climainprofitApi.GetCliMainprofitPublic) // 结算总表开放接口
	}
}
