package system

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user").Use(middleware.OperationRecord())
	var baseApi = v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		userRouter.POST("register", baseApi.Register)                     // 用户注册账号
		userRouter.POST("changePassword", baseApi.ChangePassword)         // 用户修改密码
		userRouter.POST("getUserList", baseApi.GetUserList)               // 分页获取用户列表
		userRouter.POST("setUserAuthority", baseApi.SetUserAuthority)     // 设置用户权限
		userRouter.DELETE("deleteUser", baseApi.DeleteUser)               // 删除用户
		userRouter.PUT("setUserInfo", baseApi.SetUserInfo)                // 设置用户信息
		userRouter.POST("setUserAuthorities", baseApi.SetUserAuthorities) // 设置用户权限组
		userRouter.GET("getUserInfo", baseApi.GetUserInfo)                // 获取自身信息
	}
}
