package FlyResultPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FlyResultRouter struct {
}

// InitFlyResultRouter 初始化 FlyResult 路由信息
func (s *FlyResultRouter) InitFlyResultRouter(Router *gin.RouterGroup) {
	FlyRtRouter := Router.Group("FlyRt").Use(middleware.OperationRecord())
	FlyRtRouterWithoutRecord := Router.Group("FlyRt")
	var FlyRtApi = v1.ApiGroupApp.FlyResultPkgApiGroup.FlyResultApi
	{
		FlyRtRouter.POST("createFlyResult", FlyRtApi.CreateFlyResult)             // 新建FlyResult
		FlyRtRouter.DELETE("deleteFlyResult", FlyRtApi.DeleteFlyResult)           // 删除FlyResult
		FlyRtRouter.DELETE("deleteFlyResultByIds", FlyRtApi.DeleteFlyResultByIds) // 批量删除FlyResult
		FlyRtRouter.PUT("updateFlyResult", FlyRtApi.UpdateFlyResult)              // 更新FlyResult
	}
	{
		FlyRtRouterWithoutRecord.GET("findFlyResult", FlyRtApi.FindFlyResult)                             // 根据ID获取FlyResult
		FlyRtRouterWithoutRecord.GET("getFlyResultList", FlyRtApi.GetFlyResultList)                       // 获取FlyResult列表
		FlyRtRouterWithoutRecord.GET("queryAirlineRecordFlyResult", FlyRtApi.QueryAirlineRecordFlyResult) // 查询航线下所有作业记录及成果
		FlyRtRouterWithoutRecord.POST("dataResultDownload", FlyRtApi.DataResultDownload)                  // 数据成果打包下载
	}
}
