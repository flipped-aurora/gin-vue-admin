package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

type CasbinRouter struct {
}

func (s *CasbinRouter) InitCasbinRouter(Router *gin.RouterGroup) {
	casbinRouter := Router.Group("casbin").Use(middleware.OperationRecord())
	var casbinApi = v1.ApiGroupApp.SystemApiGroup.CasbinApi
	{
		casbinRouter.POST("updateCasbin", casbinApi.UpdateCasbin)
		casbinRouter.POST("getPolicyPathByAuthorityId", casbinApi.GetPolicyPathByAuthorityId)
	}
}
