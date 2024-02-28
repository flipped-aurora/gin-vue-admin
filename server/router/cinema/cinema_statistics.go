package cinema

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CinemaStatisticsRouter struct {
}

// InitCinemaStatisticsRouter 初始化 cinemaStatistics表 路由信息
func (s *CinemaStatisticsRouter) InitCinemaStatisticsRouter(Router *gin.RouterGroup) {
	cinemaStatisticsRouterWithoutRecord := Router.Group("cinemaStatistics")
	var cinemaStatisticsApi = v1.ApiGroupApp.CinemaApiGroup.CinemaStatisticsApi
	{
		cinemaStatisticsRouterWithoutRecord.GET("getCinemaStatisticsList", cinemaStatisticsApi.GetCinemaStatisticsList) // 获取cinemaStatistics表列表
	}
}
