package remoteServer

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysRemotesRouter struct {
}

// InitSysRemotesRouter 初始化 远程服务器配置表 路由信息
func (s *SysRemotesRouter) InitSysRemotesRouter(Router *gin.RouterGroup) {
	sysRemotesRouter := Router.Group("sysRemotes").Use(middleware.OperationRecord())
	sysRemotesRouterWithoutRecord := Router.Group("sysRemotes")
	var sysRemotesApi = v1.ApiGroupApp.RemoteServerApiGroup.SysRemotesApi
	{
		sysRemotesRouter.POST("createSysRemotes", sysRemotesApi.CreateSysRemotes)   // 新建远程服务器配置表
		sysRemotesRouter.DELETE("deleteSysRemotes", sysRemotesApi.DeleteSysRemotes) // 删除远程服务器配置表
		sysRemotesRouter.DELETE("deleteSysRemotesByIds", sysRemotesApi.DeleteSysRemotesByIds) // 批量删除远程服务器配置表
		sysRemotesRouter.PUT("updateSysRemotes", sysRemotesApi.UpdateSysRemotes)    // 更新远程服务器配置表
	}
	{
		sysRemotesRouterWithoutRecord.GET("findSysRemotes", sysRemotesApi.FindSysRemotes)        // 根据ID获取远程服务器配置表
		sysRemotesRouterWithoutRecord.GET("getSysRemotesList", sysRemotesApi.GetSysRemotesList)  // 获取远程服务器配置表列表
	}
}
