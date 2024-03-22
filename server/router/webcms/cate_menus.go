package webcms

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CateMenusRouter struct {
}

// InitCateMenusRouter 初始化 CateMenus 路由信息
func (s *CateMenusRouter) InitCateMenusRouter(Router *gin.RouterGroup) {
	cateMenusRouter := Router.Group("cateMenus").Use(middleware.OperationRecord())
	cateMenusRouterWithoutRecord := Router.Group("cateMenus")
	var cateMenusApi = v1.ApiGroupApp.WebcmsApiGroup.CateMenusApi
	{
		cateMenusRouter.POST("createCateMenus", cateMenusApi.CreateCateMenus)             // 新建CateMenus
		cateMenusRouter.DELETE("deleteCateMenus", cateMenusApi.DeleteCateMenus)           // 删除CateMenus
		cateMenusRouter.DELETE("deleteCateMenusByIds", cateMenusApi.DeleteCateMenusByIds) // 批量删除CateMenus
		cateMenusRouter.PUT("updateCateMenus", cateMenusApi.UpdateCateMenus)              // 更新CateMenus
	}
	{
		cateMenusRouterWithoutRecord.GET("findCateMenus", cateMenusApi.FindCateMenus)       // 根据ID获取CateMenus
		cateMenusRouterWithoutRecord.GET("getCateMenusList", cateMenusApi.GetCateMenusList) // 获取CateMenus列表
		cateMenusRouterWithoutRecord.GET("getTemplateList", cateMenusApi.GetTemplateList)   // 获取模板名称列表
		cateMenusRouterWithoutRecord.GET("getModelsList", cateMenusApi.GetModelsList)       // 获取模型列表
	}
}
