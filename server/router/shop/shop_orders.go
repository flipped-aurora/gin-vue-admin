package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ShopOrdersRouter struct {
}

// InitShopOrdersRouter 初始化 shopOrders表 路由信息
func (s *ShopOrdersRouter) InitShopOrdersRouter(Router *gin.RouterGroup) {
	shopOrdersRouter := Router.Group("shopOrders").Use(middleware.OperationRecord())
	shopOrdersRouterWithoutRecord := Router.Group("shopOrders")
	var shopOrdersApi = v1.ApiGroupApp.ShopApiGroup.ShopOrdersApi
	{
		shopOrdersRouter.POST("createShopOrders", shopOrdersApi.CreateShopOrders)             // 新建shopOrders表
		shopOrdersRouter.DELETE("deleteShopOrders", shopOrdersApi.DeleteShopOrders)           // 删除shopOrders表
		shopOrdersRouter.DELETE("deleteShopOrdersByIds", shopOrdersApi.DeleteShopOrdersByIds) // 批量删除shopOrders表
		shopOrdersRouter.PUT("updateShopOrders", shopOrdersApi.UpdateShopOrders)              // 更新shopOrders表
	}
	{
		shopOrdersRouterWithoutRecord.GET("findShopOrders", shopOrdersApi.FindShopOrders)       // 根据ID获取shopOrders表
		shopOrdersRouterWithoutRecord.GET("getShopOrdersList", shopOrdersApi.GetShopOrdersList) // 获取shopOrders表列表
	}
}
