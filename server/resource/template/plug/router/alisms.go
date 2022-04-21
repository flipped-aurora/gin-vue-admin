package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/api"
	"github.com/gin-gonic/gin"
)

type AliSmsRouter struct {
}

func (s *AliSmsRouter) InitAliSmsRouter(Router *gin.RouterGroup) {
	emailRouter := Router
	aliSmsApi := api.ApiGroupApp.AliSmsApi
	{
		emailRouter.POST("sendSms", aliSmsApi.SendAliSms) // 发送测试邮件
	}
}
