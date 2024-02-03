package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserInfoRouter struct {
}

// InitUserInfoRouter 初始化 用户信息 路由信息
func (s *UserInfoRouter) InitUserInfoRouter(Router *gin.RouterGroup) {
	userDataRouter := Router.Group("userData").Use(middleware.OperationRecord())
	userDataRouterWithoutRecord := Router.Group("userData")
	var userDataApi = v1.ApiGroupApp.UserApiGroup.UserInfoApi
	{
		userDataRouter.POST("createUserInfo", userDataApi.CreateUserInfo)   // 新建用户信息
		userDataRouter.DELETE("deleteUserInfo", userDataApi.DeleteUserInfo) // 删除用户信息
		userDataRouter.DELETE("deleteUserInfoByIds", userDataApi.DeleteUserInfoByIds) // 批量删除用户信息
		userDataRouter.PUT("updateUserInfo", userDataApi.UpdateUserInfo)    // 更新用户信息
	}
	{
		userDataRouterWithoutRecord.GET("findUserInfo", userDataApi.FindUserInfo)        // 根据ID获取用户信息
		userDataRouterWithoutRecord.GET("getUserInfoList", userDataApi.GetUserInfoList)  // 获取用户信息列表
	}
}
