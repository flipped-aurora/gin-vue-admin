package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BannerRouter struct {
}

// InitBannerRouter 初始化 Banner 路由信息
func (s *BannerRouter) InitBannerRouter(Router *gin.RouterGroup) {
	bannerRouter := Router.Group("banner").Use(middleware.OperationRecord())
	bannerRouterWithoutRecord := Router.Group("banner")
	h5BannerRouterWithoutRecord := Router.Group(global.GetAppApi() + "banner")
	var bannerApi = v1.ApiGroupApp.ClothingApiGroup.BannerApi
	{
		bannerRouter.POST("createBanner", bannerApi.CreateBanner)             // 新建Banner
		bannerRouter.DELETE("deleteBanner", bannerApi.DeleteBanner)           // 删除Banner
		bannerRouter.DELETE("deleteBannerByIds", bannerApi.DeleteBannerByIds) // 批量删除Banner
		bannerRouter.PUT("updateBanner", bannerApi.UpdateBanner)              // 更新Banner
	}
	{
		bannerRouterWithoutRecord.GET("findBanner", bannerApi.FindBanner)       // 根据ID获取Banner
		bannerRouterWithoutRecord.GET("getBannerList", bannerApi.GetBannerList) // 获取Banner列表
	}
	{
		h5BannerRouterWithoutRecord.GET("findBanner", bannerApi.FindBanner)       // 根据ID获取Banner
		h5BannerRouterWithoutRecord.GET("getBannerList", bannerApi.GetBannerList) // 获取Banner列表
	}
}
