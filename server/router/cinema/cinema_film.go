package cinema

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CinemaFilmRouter struct {
}

// InitCinemaFilmRouter 初始化 cinemaFilm表 路由信息
func (s *CinemaFilmRouter) InitCinemaFilmRouter(Router *gin.RouterGroup) {
	cinemaFilmRouter := Router.Group("cinemaFilm").Use(middleware.OperationRecord())
	cinemaFilmRouterWithoutRecord := Router.Group("cinemaFilm")
	var cinemaFilmApi = v1.ApiGroupApp.CinemaApiGroup.CinemaFilmApi
	{
		cinemaFilmRouter.POST("createCinemaFilm", cinemaFilmApi.CreateCinemaFilm)   // 新建cinemaFilm表
		cinemaFilmRouter.DELETE("deleteCinemaFilm", cinemaFilmApi.DeleteCinemaFilm) // 删除cinemaFilm表
		cinemaFilmRouter.DELETE("deleteCinemaFilmByIds", cinemaFilmApi.DeleteCinemaFilmByIds) // 批量删除cinemaFilm表
		cinemaFilmRouter.PUT("updateCinemaFilm", cinemaFilmApi.UpdateCinemaFilm)    // 更新cinemaFilm表
	}
	{
		cinemaFilmRouterWithoutRecord.GET("findCinemaFilm", cinemaFilmApi.FindCinemaFilm)        // 根据ID获取cinemaFilm表
		cinemaFilmRouterWithoutRecord.GET("getCinemaFilmList", cinemaFilmApi.GetCinemaFilmList)  // 获取cinemaFilm表列表
	}
}
