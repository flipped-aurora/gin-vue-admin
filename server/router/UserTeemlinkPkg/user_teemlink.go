package UserTeemlinkPkg

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserTeemlinkRouter struct {
}

// InitUserTeemlinkRouter 初始化 UserTeemlink 路由信息
func (s *UserTeemlinkRouter) InitUserTeemlinkRouter(Router *gin.RouterGroup) {
	usertlRouter := Router.Group("usertl").Use(middleware.OperationRecord())
	usertlRouterWithoutRecord := Router.Group("usertl")
	var usertlApi = v1.ApiGroupApp.UserTeemlinkPkgApiGroup.UserTeemlinkApi
	{
		usertlRouter.POST("createUserTeemlink", usertlApi.CreateUserTeemlink)             // 新建UserTeemlink
		usertlRouter.DELETE("deleteUserTeemlink", usertlApi.DeleteUserTeemlink)           // 删除UserTeemlink
		usertlRouter.DELETE("deleteUserTeemlinkByIds", usertlApi.DeleteUserTeemlinkByIds) // 批量删除UserTeemlink
		usertlRouter.PUT("updateUserTeemlink", usertlApi.UpdateUserTeemlink)              // 更新UserTeemlink
	}
	{
		usertlRouterWithoutRecord.GET("findUserTeemlink", usertlApi.FindUserTeemlink)       // 根据ID获取UserTeemlink
		usertlRouterWithoutRecord.GET("getUserTeemlinkList", usertlApi.GetUserTeemlinkList) // 获取UserTeemlink列表
		// usertlRouterWithoutRecord.GET("getGvaToken", usertlApi.GetGvaToken)
	}
}
