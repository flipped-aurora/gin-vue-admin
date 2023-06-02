package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type InventoryRouter struct {
}

// InitInventoryRouter 初始化 Inventory 路由信息
func (s *InventoryRouter) InitInventoryRouter(Router *gin.RouterGroup) {
	inventoryRouter := Router.Group("inventory").Use(middleware.OperationRecord())
	inventoryRouterWithoutRecord := Router.Group("inventory")
	h5InventoryRouterWithoutRecord := Router.Group(global.GetAppApi() + "inventory")
	var inventoryApi = v1.ApiGroupApp.ClothingApiGroup.InventoryApi
	{
		inventoryRouter.POST("createInventory", inventoryApi.CreateInventory)             // 新建Inventory
		inventoryRouter.DELETE("deleteInventory", inventoryApi.DeleteInventory)           // 删除Inventory
		inventoryRouter.DELETE("deleteInventoryByIds", inventoryApi.DeleteInventoryByIds) // 批量删除Inventory
		inventoryRouter.PUT("updateInventory", inventoryApi.UpdateInventory)              // 更新Inventory
	}
	{
		inventoryRouterWithoutRecord.GET("findInventory", inventoryApi.FindInventory)       // 根据ID获取Inventory
		inventoryRouterWithoutRecord.GET("getInventoryList", inventoryApi.GetInventoryList) // 获取Inventory列表
	}
	{
		h5InventoryRouterWithoutRecord.GET("getInventoryList", inventoryApi.GetInventoryList) // 获取Inventory列表
	}
}
