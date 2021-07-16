package system

import (
	"gin-vue-admin/api/v1"
	"github.com/gin-gonic/gin"
)

func (s *Router) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("login", v1.Login)
		BaseRouter.POST("captcha", v1.Captcha)
	}
	return BaseRouter
}
