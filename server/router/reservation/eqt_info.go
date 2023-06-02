package reservation

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EqtInfoRouter struct {
}

// InitEqtInfoRouter 初始化 EqtInfo 路由信息
func (s *EqtInfoRouter) InitEqtInfoRouter(Router *gin.RouterGroup) {
	eqtInfoRouter := Router.Group("eqtInfo").Use(middleware.OperationRecord())
	eqtInfoRouterWithoutRecord := Router.Group("eqtInfo")
	var eqtInfoApi = v1.ApiGroupApp.ReservationApiGroup.EqtInfoApi
	{
		eqtInfoRouter.POST("createEqtInfo", eqtInfoApi.CreateEqtInfo)              // 新建EqtInfo
		eqtInfoRouter.DELETE("deleteEqtInfo", eqtInfoApi.DeleteEqtInfo)            // 删除EqtInfo
		eqtInfoRouter.DELETE("deleteEqtInfoByIds", eqtInfoApi.DeleteEqtInfoByIds)  // 批量删除EqtInfo
		eqtInfoRouter.PUT("updateEqtInfo", eqtInfoApi.UpdateEqtInfo)               // 更新EqtInfo
		eqtInfoRouter.PUT("UpdateEqtStatusOpen", eqtInfoApi.UpdateEqtStatusOpen)   // 开放预约
		eqtInfoRouter.PUT("UpdateEqtStatusClose", eqtInfoApi.UpdateEqtStatusClose) // 关闭预约
	}
	{
		eqtInfoRouterWithoutRecord.GET("findEqtInfo", eqtInfoApi.FindEqtInfo)       // 根据ID获取EqtInfo
		eqtInfoRouterWithoutRecord.GET("getEqtInfoList", eqtInfoApi.GetEqtInfoList) // 获取EqtInfo列表
	}
}
