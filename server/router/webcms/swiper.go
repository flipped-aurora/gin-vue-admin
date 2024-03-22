package webcms

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SwiperRouter struct {
}

// InitSwiperRouter 初始化 Swiper 路由信息
func (s *SwiperRouter) InitSwiperRouter(Router *gin.RouterGroup) {
	swiperRouter := Router.Group("swiper").Use(middleware.OperationRecord())
	swiperRouterWithoutRecord := Router.Group("swiper")
	var swiperApi = v1.ApiGroupApp.WebcmsApiGroup.SwiperApi
	{
		swiperRouter.POST("createSwiper", swiperApi.CreateSwiper)             // 新建Swiper
		swiperRouter.DELETE("deleteSwiper", swiperApi.DeleteSwiper)           // 删除Swiper
		swiperRouter.DELETE("deleteSwiperByIds", swiperApi.DeleteSwiperByIds) // 批量删除Swiper
		swiperRouter.PUT("updateSwiper", swiperApi.UpdateSwiper)              // 更新Swiper
	}
	{
		swiperRouterWithoutRecord.GET("findSwiper", swiperApi.FindSwiper)       // 根据ID获取Swiper
		swiperRouterWithoutRecord.GET("getSwiperList", swiperApi.GetSwiperList) // 获取Swiper列表
	}
}
