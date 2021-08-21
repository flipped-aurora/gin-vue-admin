package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

type JwtRouter struct {
}

func (s *JwtRouter) InitJwtRouter(Router *gin.RouterGroup) {
	jwtRouter := Router.Group("jwt").Use(middleware.OperationRecord())
	var jwtApi = v1.ApiGroupApp.SystemApiGroup.JwtApi
	{
		jwtRouter.POST("jsonInBlacklist", jwtApi.JsonInBlacklist) // jwt加入黑名单
	}
}
