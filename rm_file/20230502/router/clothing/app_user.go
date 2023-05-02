package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppUserRouter struct {
}

// InitAppUserRouter 初始化 AppUser 路由信息
func (s *AppUserRouter) InitAppUserRouter(Router *gin.RouterGroup) {
	appUserRouter := Router.Group("appUser").Use(middleware.OperationRecord())
	appUserRouterWithoutRecord := Router.Group("appUser")
	var appUserApi = v1.ApiGroupApp.ClothingApiGroup.AppUserApi
	{
		appUserRouter.POST("createAppUser", appUserApi.CreateAppUser)   // 新建AppUser
		appUserRouter.DELETE("deleteAppUser", appUserApi.DeleteAppUser) // 删除AppUser
		appUserRouter.DELETE("deleteAppUserByIds", appUserApi.DeleteAppUserByIds) // 批量删除AppUser
		appUserRouter.PUT("updateAppUser", appUserApi.UpdateAppUser)    // 更新AppUser
	}
	{
		appUserRouterWithoutRecord.GET("findAppUser", appUserApi.FindAppUser)        // 根据ID获取AppUser
		appUserRouterWithoutRecord.GET("getAppUserList", appUserApi.GetAppUserList)  // 获取AppUser列表
	}
}
