package nestrolepkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type NestRoleRouter struct {
}

// InitNestRoleRouter 初始化 NestRole 路由信息
func (s *NestRoleRouter) InitNestRoleRouter(Router *gin.RouterGroup) {
	nestroleRouter := Router.Group("nestrole").Use(middleware.OperationRecord())
	nestroleRouterWithoutRecord := Router.Group("nestrole")
	var nestroleApi = v1.ApiGroupApp.NestrolepkgApiGroup.NestRoleApi
	{
		nestroleRouter.POST("createNestRole", nestroleApi.CreateNestRole)   // 新建NestRole
		nestroleRouter.DELETE("deleteNestRole", nestroleApi.DeleteNestRole) // 删除NestRole
		nestroleRouter.DELETE("deleteNestRoleByIds", nestroleApi.DeleteNestRoleByIds) // 批量删除NestRole
		nestroleRouter.PUT("updateNestRole", nestroleApi.UpdateNestRole)    // 更新NestRole
	}
	{
		nestroleRouterWithoutRecord.GET("findNestRole", nestroleApi.FindNestRole)        // 根据ID获取NestRole
		nestroleRouterWithoutRecord.GET("getNestRoleList", nestroleApi.GetNestRoleList)  // 获取NestRole列表
	}
}
