package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppRoleRouter struct {
}

// InitAppRoleRouter 初始化 AppRole 路由信息
func (s *AppRoleRouter) InitAppRoleRouter(Router *gin.RouterGroup) {
	appRoleRouter := Router.Group("appRole").Use(middleware.OperationRecord())
	appRoleRouterWithoutRecord := Router.Group("appRole")
	h5AppRoleRouterWithoutRecord := Router.Group(global.GetAppApi() + "appRole")
	var appRoleApi = v1.ApiGroupApp.ClothingApiGroup.AppRoleApi
	{
		appRoleRouter.POST("createAppRole", appRoleApi.CreateAppRole)             // 新建AppRole
		appRoleRouter.DELETE("deleteAppRole", appRoleApi.DeleteAppRole)           // 删除AppRole
		appRoleRouter.DELETE("deleteAppRoleByIds", appRoleApi.DeleteAppRoleByIds) // 批量删除AppRole
		appRoleRouter.PUT("updateAppRole", appRoleApi.UpdateAppRole)              // 更新AppRole
	}
	{
		appRoleRouterWithoutRecord.GET("findAppRole", appRoleApi.FindAppRole)       // 根据ID获取AppRole
		appRoleRouterWithoutRecord.GET("getAppRoleList", appRoleApi.GetAppRoleList) // 获取AppRole列表
	}
	{
		h5AppRoleRouterWithoutRecord.GET("getAppRoleList", appRoleApi.GetAppRoleList) // 获取AppRole列表
	}
}
