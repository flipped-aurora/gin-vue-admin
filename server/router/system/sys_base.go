package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	userRouter := Router.Group("user").Use(middleware.OperationRecord())
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("login", baseApi.Login)
		userRouter.POST("register", baseApi.RegisterRegularUser) // 普通用户注册账号
		baseRouter.POST("captcha", baseApi.Captcha)
	}
	return baseRouter
}
