package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SecurityConfigRouter struct{}

// InitSecurityConfigRouter 初始化 安全配置 路由信息
func (s *SecurityConfigRouter) InitSecurityConfigRouter(Router *gin.RouterGroup) {
	securityConfigRouter := Router.Group("securityConfig").Use(middleware.OperationRecord())
	securityConfigRouterWithoutRecord := Router.Group("securityConfig")
	{
		securityConfigRouter.POST("setSecurityConfig", securityConfigApi.SetSecurityConfig) // 设置安全配置
	}
	{
		securityConfigRouterWithoutRecord.GET("getSecurityConfig", securityConfigApi.GetSecurityConfig) // 获取安全配置
	}
}
