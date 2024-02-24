package cinema

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CinemaOrderRouter struct {
}

// InitCinemaOrderRouter 初始化 cinemaOrder表 路由信息
func (s *CinemaOrderRouter) InitCinemaOrderRouter(Router *gin.RouterGroup) {
	cinemaOrderRouterWithoutRecord := Router.Group("cinemaOrder")
	var cinemaOrderApi = v1.ApiGroupApp.CinemaApiGroup.CinemaOrderApi
	{
		cinemaOrderRouterWithoutRecord.GET("findCinemaOrder", cinemaOrderApi.FindCinemaOrder)       // 根据ID获取cinemaOrder表
		cinemaOrderRouterWithoutRecord.GET("getStatistics", cinemaOrderApi.GetStatistics)           // 根据ID获取cinemaOrder表
		cinemaOrderRouterWithoutRecord.GET("getCinemaOrderList", cinemaOrderApi.GetCinemaOrderList) // 获取cinemaOrder表列表
	}
}
