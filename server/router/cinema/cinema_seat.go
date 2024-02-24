package cinema

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CinemaSeatRouter struct {
}

// InitCinemaSeatRouter 初始化 cinemaSeat表 路由信息
func (s *CinemaSeatRouter) InitCinemaSeatRouter(Router *gin.RouterGroup) {
	cinemaSeatRouter := Router.Group("cinemaSeat").Use(middleware.OperationRecord())
	cinemaSeatRouterWithoutRecord := Router.Group("cinemaSeat")
	var cinemaSeatApi = v1.ApiGroupApp.CinemaApiGroup.CinemaSeatApi
	{
		cinemaSeatRouter.POST("createCinemaSeat", cinemaSeatApi.CreateCinemaSeat)   // 新建cinemaSeat表
		cinemaSeatRouter.DELETE("deleteCinemaSeat", cinemaSeatApi.DeleteCinemaSeat) // 删除cinemaSeat表
	}
	{
		cinemaSeatRouterWithoutRecord.GET("findCinemaSeat", cinemaSeatApi.FindCinemaSeat)       // 根据ID获取cinemaSeat表
		cinemaSeatRouterWithoutRecord.GET("getCinemaSeatList", cinemaSeatApi.GetCinemaSeatList) // 获取cinemaSeat表列表
	}
}
