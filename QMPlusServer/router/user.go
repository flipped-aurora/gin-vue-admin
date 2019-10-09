package router

import (
	"github.com/gin-gonic/gin"
	"main/controller/api"
	"main/middleware"
)

func InitUserRouter(Router *gin.Engine) {
	UserRouter := Router.Group("user").Use(middleware.JWTAuth())
	{
		UserRouter.POST("changePassword", api.ChangePassword)   // 修改密码
		UserRouter.POST("uploadHeaderImg", api.UploadHeaderImg) //上传头像
		UserRouter.POST("getUserList", api.GetUserList)         // 分页获取用户列表
	}
}
