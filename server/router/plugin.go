package router

import (
	"github.com/eyotang/game-proxy/server/api/v1"
	"github.com/eyotang/game-proxy/server/middleware"
	"github.com/gin-gonic/gin"
)

func InitProductPluginRouter(Router *gin.RouterGroup) {
	ProductPluginRouter := Router.Group("productPlugin").Use(middleware.OperationRecord())
	{
		ProductPluginRouter.POST("createProductPlugin", v1.CreateProductPlugin)   // 新建ProductPlugin
		ProductPluginRouter.DELETE("deleteProductPlugin", v1.DeleteProductPlugin) // 删除ProductPlugin
		ProductPluginRouter.DELETE("deleteProductPluginByIds", v1.DeleteProductPluginByIds) // 批量删除ProductPlugin
		ProductPluginRouter.PUT("updateProductPlugin", v1.UpdateProductPlugin)    // 更新ProductPlugin
		ProductPluginRouter.GET("findProductPlugin", v1.FindProductPlugin)        // 根据ID获取ProductPlugin
		ProductPluginRouter.GET("getProductPluginList", v1.GetProductPluginList)  // 获取ProductPlugin列表
	}
}
