package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user").Use(middleware.OperationRecord())
	userRouterWithoutRecord := Router.Group("user")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		userRouter.POST("admin_register", baseApi.Register)                   // 管理员注册账号
		userRouter.POST("changePassword", baseApi.ChangePassword)             // 用户修改密码
		userRouter.POST("setUserAuthority", baseApi.SetUserAuthority)         // 设置用户权限
		userRouter.DELETE("deleteUser", baseApi.DeleteUser)                   // 删除用户All
		userRouter.DELETE("deleteRegularUser", baseApi.DeleteRegularUser)     // 删除普通用户
		userRouter.PUT("setUserInfo", baseApi.SetUserInfo)                    // 设置用户信息
		userRouter.PUT("setSelfInfo", baseApi.SetSelfInfo)                    // 设置自身信息
		userRouter.POST("setUserAuthorities", baseApi.SetUserAuthorities)     // 设置用户权限组
		userRouter.POST("resetPassword", baseApi.ResetPassword)               // 重置密码All
		userRouter.POST("resetRegularPassword", baseApi.ResetRegularPassword) // 重置普通用户密码
	}
	{
		userRouterWithoutRecord.POST("getUserList", baseApi.GetUserList)               // 分页获取用户列表All
		userRouterWithoutRecord.POST("getAdminUserList", baseApi.GetAdminUserList)     // 分页获取管理员用户列表
		userRouterWithoutRecord.POST("getRegularUserList", baseApi.GetRegularUserList) // 分页获取普通用户列表
		userRouterWithoutRecord.GET("getUserInfo", baseApi.GetUserInfo)                // 获取自身信息
	}
}
