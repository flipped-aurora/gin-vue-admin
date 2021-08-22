package router

import (
	"github.com/flipped-aurora/gin-vue-admin/middleware"
	"github.com/flipped-aurora/gin-vue-admin/plugin/email/api"
	"github.com/gin-gonic/gin"
)

type EmailRouter struct {
}

func (s *EmailRouter) InitEmailRouter(Router *gin.RouterGroup) {
	emailRouter := Router.Use(middleware.OperationRecord())
	var EmailApi = api.ApiGroupApp.EmailApi.EmailTest
	{
		emailRouter.POST("emailTest", EmailApi) // 发送测试邮件
	}
}
