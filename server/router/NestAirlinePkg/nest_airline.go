package NestAirlinePkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type NestAirlineRouter struct {
}

// InitNestAirlineRouter 初始化 NestAirline 路由信息
func (s *NestAirlineRouter) InitNestAirlineRouter(Router *gin.RouterGroup) {
	NtAirlineRouter := Router.Group("NtAirline").Use(middleware.OperationRecord())
	NtAirlineRouterWithoutRecord := Router.Group("NtAirline")
	var NtAirlineApi = v1.ApiGroupApp.NestAirlinePkgApiGroup.NestAirlineApi
	{
		NtAirlineRouter.POST("createNestAirline", NtAirlineApi.CreateNestAirline)             // 新建NestAirline
		NtAirlineRouter.DELETE("deleteNestAirline", NtAirlineApi.DeleteNestAirline)           // 删除NestAirline
		NtAirlineRouter.DELETE("deleteNestAirlineByIds", NtAirlineApi.DeleteNestAirlineByIds) // 批量删除NestAirline
		NtAirlineRouter.PUT("updateNestAirline", NtAirlineApi.UpdateNestAirline)              // 更新NestAirline
	}
	{
		NtAirlineRouterWithoutRecord.GET("findNestAirline", NtAirlineApi.FindNestAirline)       // 根据ID获取NestAirline
		NtAirlineRouterWithoutRecord.GET("getNestAirlineList", NtAirlineApi.GetNestAirlineList) // 获取NestAirline列表
	}
}
