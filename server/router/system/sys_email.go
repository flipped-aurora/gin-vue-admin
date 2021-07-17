package system

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

type EmailRouter struct {
}

func (s *EmailRouter) InitEmailRouter(Router *gin.RouterGroup) {
	emailRouter := Router.Group("email").Use(middleware.OperationRecord())
	var systemApi = v1.ApiGroupApp.SystemApiGroup.SystemApi
	{
		emailRouter.POST("emailTest", systemApi.EmailTest) // 发送测试邮件
	}
}
