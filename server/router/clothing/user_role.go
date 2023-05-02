package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserRoleRouter struct {
}

// InitUserRoleRouter 初始化 UserRole 路由信息
func (s *UserRoleRouter) InitUserRoleRouter(Router *gin.RouterGroup) {
	userRoleRouter := Router.Group("userRole").Use(middleware.OperationRecord())
	userRoleRouterWithoutRecord := Router.Group("userRole")
	var userRoleApi = v1.ApiGroupApp.ClothingApiGroup.UserRoleApi
	{
		userRoleRouter.POST("createUserRole", userRoleApi.CreateUserRole)   // 新建UserRole
		userRoleRouter.DELETE("deleteUserRole", userRoleApi.DeleteUserRole) // 删除UserRole
		userRoleRouter.DELETE("deleteUserRoleByIds", userRoleApi.DeleteUserRoleByIds) // 批量删除UserRole
		userRoleRouter.PUT("updateUserRole", userRoleApi.UpdateUserRole)    // 更新UserRole
	}
	{
		userRoleRouterWithoutRecord.GET("findUserRole", userRoleApi.FindUserRole)        // 根据ID获取UserRole
		userRoleRouterWithoutRecord.GET("getUserRoleList", userRoleApi.GetUserRoleList)  // 获取UserRole列表
	}
}
