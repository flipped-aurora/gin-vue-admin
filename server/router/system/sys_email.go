package system

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func (s *Router) InitEmailRouter(Router *gin.RouterGroup) {
	EmailRouter := Router.Group("email").Use(middleware.OperationRecord())
	{
		EmailRouter.POST("emailTest", v1.EmailTest) // 发送测试邮件
	}
}
