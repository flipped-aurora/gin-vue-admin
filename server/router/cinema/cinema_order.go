package cinema

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CinemaOrderRouter struct {
}

// InitCinemaOrderRouter 初始化 cinemaOrder表 路由信息
func (s *CinemaOrderRouter) InitCinemaOrderRouter(Router *gin.RouterGroup) {
	cinemaOrderRouter := Router.Group("cinemaOrder").Use(middleware.OperationRecord())
	cinemaOrderRouterWithoutRecord := Router.Group("cinemaOrder")
	var cinemaOrderApi = v1.ApiGroupApp.CinemaApiGroup.CinemaOrderApi
	{
		cinemaOrderRouter.POST("createCinemaOrder", cinemaOrderApi.CreateCinemaOrder)   // 新建cinemaOrder表
		cinemaOrderRouter.DELETE("deleteCinemaOrder", cinemaOrderApi.DeleteCinemaOrder) // 删除cinemaOrder表
		cinemaOrderRouter.DELETE("deleteCinemaOrderByIds", cinemaOrderApi.DeleteCinemaOrderByIds) // 批量删除cinemaOrder表
		cinemaOrderRouter.PUT("updateCinemaOrder", cinemaOrderApi.UpdateCinemaOrder)    // 更新cinemaOrder表
	}
	{
		cinemaOrderRouterWithoutRecord.GET("findCinemaOrder", cinemaOrderApi.FindCinemaOrder)        // 根据ID获取cinemaOrder表
		cinemaOrderRouterWithoutRecord.GET("getCinemaOrderList", cinemaOrderApi.GetCinemaOrderList)  // 获取cinemaOrder表列表
	}
}
