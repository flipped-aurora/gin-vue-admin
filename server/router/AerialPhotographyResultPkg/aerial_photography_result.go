package AerialPhotographyResultPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AerialPhotographyResultRouter struct {
}

// InitAerialPhotographyResultRouter 初始化 AerialPhotographyResult 路由信息
func (s *AerialPhotographyResultRouter) InitAerialPhotographyResultRouter(Router *gin.RouterGroup) {
	ALPhotographyResultRouter := Router.Group("ALPhotographyResult").Use(middleware.OperationRecord())
	ALPhotographyResultRouterWithoutRecord := Router.Group("ALPhotographyResult")
	var ALPhotographyResultApi = v1.ApiGroupApp.AerialPhotographyResultPkgApiGroup.AerialPhotographyResultApi
	{
		ALPhotographyResultRouter.POST("createAerialPhotographyResult", ALPhotographyResultApi.CreateAerialPhotographyResult)             // 新建AerialPhotographyResult
		ALPhotographyResultRouter.DELETE("deleteAerialPhotographyResult", ALPhotographyResultApi.DeleteAerialPhotographyResult)           // 删除AerialPhotographyResult
		ALPhotographyResultRouter.DELETE("deleteAerialPhotographyResultByIds", ALPhotographyResultApi.DeleteAerialPhotographyResultByIds) // 批量删除AerialPhotographyResult
		ALPhotographyResultRouter.PUT("updateAerialPhotographyResult", ALPhotographyResultApi.UpdateAerialPhotographyResult)              // 更新AerialPhotographyResult
	}
	{
		ALPhotographyResultRouterWithoutRecord.GET("findAerialPhotographyResult", ALPhotographyResultApi.FindAerialPhotographyResult)       // 根据ID获取AerialPhotographyResult
		ALPhotographyResultRouterWithoutRecord.GET("getAerialPhotographyResultList", ALPhotographyResultApi.GetAerialPhotographyResultList) // 获取AerialPhotographyResult列表
		ALPhotographyResultRouterWithoutRecord.GET("queryAerialPhotographyResult", ALPhotographyResultApi.QueryAerialPhotographyResult)     // 查询并处理航摄成果数据
	}
}
