package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SizeListRouter struct {
}

// InitSizeListRouter 初始化 SizeList 路由信息
func (s *SizeListRouter) InitSizeListRouter(Router *gin.RouterGroup) {
	sizeListRouter := Router.Group("sizeList").Use(middleware.OperationRecord())
	sizeListRouterWithoutRecord := Router.Group("sizeList")
	var sizeListApi = v1.ApiGroupApp.ClothingApiGroup.SizeListApi
	{
		sizeListRouter.POST("createSizeList", sizeListApi.CreateSizeList)   // 新建SizeList
		sizeListRouter.DELETE("deleteSizeList", sizeListApi.DeleteSizeList) // 删除SizeList
		sizeListRouter.DELETE("deleteSizeListByIds", sizeListApi.DeleteSizeListByIds) // 批量删除SizeList
		sizeListRouter.PUT("updateSizeList", sizeListApi.UpdateSizeList)    // 更新SizeList
	}
	{
		sizeListRouterWithoutRecord.GET("findSizeList", sizeListApi.FindSizeList)        // 根据ID获取SizeList
		sizeListRouterWithoutRecord.GET("getSizeListList", sizeListApi.GetSizeListList)  // 获取SizeList列表
	}
}
