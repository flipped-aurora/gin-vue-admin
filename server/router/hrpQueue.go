package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitHrpQueueRouter(Router *gin.RouterGroup) {
	HrpQueueRouter := Router.Group("hrpQueue").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		HrpQueueRouter.POST("createHrpQueue", v1.CreateHrpQueue)   // 新建HrpQueue
		HrpQueueRouter.DELETE("deleteHrpQueue", v1.DeleteHrpQueue) // 删除HrpQueue
		HrpQueueRouter.DELETE("deleteHrpQueueByIds", v1.DeleteHrpQueueByIds) // 批量删除HrpQueue
		HrpQueueRouter.PUT("updateHrpQueue", v1.UpdateHrpQueue)    // 更新HrpQueue
		HrpQueueRouter.GET("findHrpQueue", v1.FindHrpQueue)        // 根据ID获取HrpQueue
		HrpQueueRouter.GET("getHrpQueueList", v1.GetHrpQueueList)  // 获取HrpQueue列表
	}
}
