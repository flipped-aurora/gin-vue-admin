package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ClothRouter struct {
}

// InitClothRouter 初始化 Cloth 路由信息
func (s *ClothRouter) InitClothRouter(Router *gin.RouterGroup) {
	clothRouter := Router.Group("cloth").Use(middleware.OperationRecord())
	clothRouterWithoutRecord := Router.Group("cloth")
	var clothApi = v1.ApiGroupApp.ClothingApiGroup.ClothApi
	{
		clothRouter.POST("createCloth", clothApi.CreateCloth)   // 新建Cloth
		clothRouter.DELETE("deleteCloth", clothApi.DeleteCloth) // 删除Cloth
		clothRouter.DELETE("deleteClothByIds", clothApi.DeleteClothByIds) // 批量删除Cloth
		clothRouter.PUT("updateCloth", clothApi.UpdateCloth)    // 更新Cloth
	}
	{
		clothRouterWithoutRecord.GET("findCloth", clothApi.FindCloth)        // 根据ID获取Cloth
		clothRouterWithoutRecord.GET("getClothList", clothApi.GetClothList)  // 获取Cloth列表
	}
}
