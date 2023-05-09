package clothing

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type StyleRouter struct {
}

// InitStyleRouter 初始化 Style 路由信息
func (s *StyleRouter) InitStyleRouter(Router *gin.RouterGroup) {
	styleRouter := Router.Group("style").Use(middleware.OperationRecord())
	styleRouterWithoutRecord := Router.Group("style")
	h5StyleRouterWithoutRecord := Router.Group(global.GetAppApi() + "style")
	var styleApi = v1.ApiGroupApp.ClothingApiGroup.StyleApi
	{
		styleRouter.POST("createStyle", styleApi.CreateStyle)             // 新建Style
		styleRouter.DELETE("deleteStyle", styleApi.DeleteStyle)           // 删除Style
		styleRouter.DELETE("deleteStyleByIds", styleApi.DeleteStyleByIds) // 批量删除Style
		styleRouter.PUT("updateStyle", styleApi.UpdateStyle)              // 更新Style
	}
	{
		styleRouterWithoutRecord.GET("findStyle", styleApi.FindStyle)       // 根据ID获取Style
		styleRouterWithoutRecord.GET("getStyleList", styleApi.GetStyleList) // 获取Style列表
	}
	{
		h5StyleRouterWithoutRecord.POST("createStyle", styleApi.CreateStyle)   // 新建Style
		h5StyleRouterWithoutRecord.DELETE("deleteStyle", styleApi.DeleteStyle) // 删除Style
		h5StyleRouterWithoutRecord.PUT("updateStyle", styleApi.UpdateStyle)    // 更新Style
		h5StyleRouterWithoutRecord.GET("findStyle", styleApi.FindStyle)        // 根据ID获取Style
		h5StyleRouterWithoutRecord.GET("getStyleList", styleApi.GetStyleList)  // 获取Style列表
	}
}
