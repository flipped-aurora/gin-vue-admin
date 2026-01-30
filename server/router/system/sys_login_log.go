package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LoginLogRouter struct{}

func (s *LoginLogRouter) InitLoginLogRouter(Router *gin.RouterGroup) {
	loginLogRouter := Router.Group("sysLoginLog").Use(middleware.OperationRecord())
	loginLogRouterWithoutRecord := Router.Group("sysLoginLog")
	sysLoginLogApi := v1.ApiGroupApp.SystemApiGroup.LoginLogApi
	{
		loginLogRouter.DELETE("deleteLoginLog", sysLoginLogApi.DeleteLoginLog)           // 删除登录日志
		loginLogRouter.DELETE("deleteLoginLogByIds", sysLoginLogApi.DeleteLoginLogByIds) // 批量删除登录日志
	}
	{
		loginLogRouterWithoutRecord.GET("findLoginLog", sysLoginLogApi.FindLoginLog)       // 根据ID获取登录日志(详情)
		loginLogRouterWithoutRecord.GET("getLoginLogList", sysLoginLogApi.GetLoginLogList) // 获取登录日志列表
	}
}
