package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysVersionRouter struct{}

// InitSysVersionRouter 初始化 版本管理 路由信息
func (s *SysVersionRouter) InitSysVersionRouter(Router *gin.RouterGroup) {
	sysVersionRouter := Router.Group("sysVersion").Use(middleware.OperationRecord())
	sysVersionRouterWithoutRecord := Router.Group("sysVersion")
	{
		sysVersionRouter.DELETE("deleteSysVersion", sysVersionApi.DeleteSysVersion)           // 删除版本管理
		sysVersionRouter.DELETE("deleteSysVersionByIds", sysVersionApi.DeleteSysVersionByIds) // 批量删除版本管理
		sysVersionRouter.POST("exportVersion", sysVersionApi.ExportVersion)                   // 导出版本数据
		sysVersionRouter.POST("importVersion", sysVersionApi.ImportVersion)                   // 导入版本数据
	}
	{
		sysVersionRouterWithoutRecord.GET("findSysVersion", sysVersionApi.FindSysVersion)           // 根据ID获取版本管理
		sysVersionRouterWithoutRecord.GET("getSysVersionList", sysVersionApi.GetSysVersionList)     // 获取版本管理列表
		sysVersionRouterWithoutRecord.GET("downloadVersionJson", sysVersionApi.DownloadVersionJson) // 下载版本JSON数据
	}
}
