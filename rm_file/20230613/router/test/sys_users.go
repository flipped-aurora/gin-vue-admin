package test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysUsersRouter struct {
}

// InitSysUsersRouter 初始化 SysUsers 路由信息
func (s *SysUsersRouter) InitSysUsersRouter(Router *gin.RouterGroup) {
	sysUsersRouter := Router.Group("sysUsers").Use(middleware.OperationRecord())
	sysUsersRouterWithoutRecord := Router.Group("sysUsers")
	var sysUsersApi = v1.ApiGroupApp.TestApiGroup.SysUsersApi
	{
		sysUsersRouter.POST("createSysUsers", sysUsersApi.CreateSysUsers)   // 新建SysUsers
		sysUsersRouter.DELETE("deleteSysUsers", sysUsersApi.DeleteSysUsers) // 删除SysUsers
		sysUsersRouter.DELETE("deleteSysUsersByIds", sysUsersApi.DeleteSysUsersByIds) // 批量删除SysUsers
		sysUsersRouter.PUT("updateSysUsers", sysUsersApi.UpdateSysUsers)    // 更新SysUsers
	}
	{
		sysUsersRouterWithoutRecord.GET("findSysUsers", sysUsersApi.FindSysUsers)        // 根据ID获取SysUsers
		sysUsersRouterWithoutRecord.GET("getSysUsersList", sysUsersApi.GetSysUsersList)  // 获取SysUsers列表
	}
}
