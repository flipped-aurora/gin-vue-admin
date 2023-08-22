package orderinfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OrderInfoRouter struct {
}

// InitOrderInfoRouter 初始化 OrderInfo 路由信息
func (s *OrderInfoRouter) InitOrderInfoRouter(Router *gin.RouterGroup) {
	orderInfoRouter := Router.Group("orderInfo").Use(middleware.OperationRecord())
	orderInfoRouterWithoutRecord := Router.Group("orderInfo")
	var orderInfoApi = v1.ApiGroupApp.OrderinfoApiGroup.OrderInfoApi
	{
		orderInfoRouter.POST("createOrderInfo", orderInfoApi.CreateOrderInfo)   // 新建OrderInfo
		orderInfoRouter.DELETE("deleteOrderInfo", orderInfoApi.DeleteOrderInfo) // 删除OrderInfo
		orderInfoRouter.DELETE("deleteOrderInfoByIds", orderInfoApi.DeleteOrderInfoByIds) // 批量删除OrderInfo
		orderInfoRouter.PUT("updateOrderInfo", orderInfoApi.UpdateOrderInfo)    // 更新OrderInfo
	}
	{
		orderInfoRouterWithoutRecord.GET("findOrderInfo", orderInfoApi.FindOrderInfo)        // 根据ID获取OrderInfo
		orderInfoRouterWithoutRecord.GET("getOrderInfoList", orderInfoApi.GetOrderInfoList)  // 获取OrderInfo列表
	}
}
