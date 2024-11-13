package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CliMainprofitRouter struct{}

func (s *CliMainprofitRouter) InitCliMainprofitRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	climainprofitRouter := Router.Group("climainprofit").Use(middleware.OperationRecord())
	climainprofitRouterWithoutRecord := Router.Group("climainprofit")
	climainprofitRouterWithoutAuth := PublicRouter.Group("climainprofit")
	{
		climainprofitRouter.POST("createCliMainprofit", climainprofitApi.CreateCliMainprofit)
		climainprofitRouter.DELETE("deleteCliMainprofit", climainprofitApi.DeleteCliMainprofit)
		climainprofitRouter.DELETE("deleteCliMainprofitByIds", climainprofitApi.DeleteCliMainprofitByIds)
		climainprofitRouter.PUT("updateCliMainprofit", climainprofitApi.UpdateCliMainprofit)
	}
	{
		climainprofitRouterWithoutRecord.GET("findCliMainprofit", climainprofitApi.FindCliMainprofit)
		climainprofitRouterWithoutRecord.GET("getCliMainprofitList", climainprofitApi.GetCliMainprofitList)
	}
	{
		climainprofitRouterWithoutAuth.GET("getCliMainprofitPublic", climainprofitApi.GetCliMainprofitPublic)
		climainprofitRouterWithoutAuth.GET("pullprofit", climainprofitApi.PullProfit)
		climainprofitRouterWithoutAuth.GET("teamprofit", climainprofitApi.TeamProfit)
	}
}
