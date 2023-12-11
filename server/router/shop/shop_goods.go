package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ShopGoodsRouter struct {
}

// InitShopGoodsRouter 初始化 shopGoods表 路由信息
func (s *ShopGoodsRouter) InitShopGoodsRouter(Router *gin.RouterGroup) {
	shopGoodsRouter := Router.Group("shopGoods").Use(middleware.OperationRecord())
	shopGoodsRouterWithoutRecord := Router.Group("shopGoods")
	var shopGoodsApi = v1.ApiGroupApp.ShopApiGroup.ShopGoodsApi
	{
		shopGoodsRouter.POST("createShopGoods", shopGoodsApi.CreateShopGoods)             // 新建shopGoods表
		shopGoodsRouter.DELETE("deleteShopGoods", shopGoodsApi.DeleteShopGoods)           // 删除shopGoods表
		shopGoodsRouter.DELETE("deleteShopGoodsByIds", shopGoodsApi.DeleteShopGoodsByIds) // 批量删除shopGoods表
		shopGoodsRouter.PUT("updateShopGoods", shopGoodsApi.UpdateShopGoods)              // 更新shopGoods表
	}
	{
		shopGoodsRouterWithoutRecord.GET("findShopGoods", shopGoodsApi.FindShopGoods)       // 根据ID获取shopGoods表
		shopGoodsRouterWithoutRecord.GET("getShopGoodsList", shopGoodsApi.GetShopGoodsList) // 获取shopGoods表列表
	}
}
