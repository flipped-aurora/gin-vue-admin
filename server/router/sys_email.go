package router

import (
	"gin-vue-admin/api/v1"
	"github.com/gin-gonic/gin"
)

func InitEmailRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("email")
	{
		UserRouter.POST("emailTest", v1.EmailTest) // 发送测试邮件
	}
}
