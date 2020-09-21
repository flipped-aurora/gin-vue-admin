package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitHrpTriagedQueueRouter(Router *gin.RouterGroup) {
	HrpTriagedQueueRouter := Router.Group("hrpTriagedQueue").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		HrpTriagedQueueRouter.POST("createHrpTriagedQueue", v1.CreateHrpTriagedQueue)   // 新建HrpTriagedQueue
		HrpTriagedQueueRouter.DELETE("deleteHrpTriagedQueue", v1.DeleteHrpTriagedQueue) // 删除HrpTriagedQueue
		HrpTriagedQueueRouter.DELETE("deleteHrpTriagedQueueByIds", v1.DeleteHrpTriagedQueueByIds) // 批量删除HrpTriagedQueue
		HrpTriagedQueueRouter.PUT("updateHrpTriagedQueue", v1.UpdateHrpTriagedQueue)    // 更新HrpTriagedQueue
		HrpTriagedQueueRouter.GET("findHrpTriagedQueue", v1.FindHrpTriagedQueue)        // 根据ID获取HrpTriagedQueue
		HrpTriagedQueueRouter.GET("getHrpTriagedQueueList", v1.GetHrpTriagedQueueList)  // 获取HrpTriagedQueue列表
	}
}
