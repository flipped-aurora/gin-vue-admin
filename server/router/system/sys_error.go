package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysErrorRouter struct{}

// InitSysErrorRouter 初始化 错误日志 路由信息
func (s *SysErrorRouter) InitSysErrorRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	sysErrprRouter := Router.Group("sysErrpr").Use(middleware.OperationRecord())
	sysErrprRouterWithoutRecord := Router.Group("sysErrpr")
	sysErrprRouterWithoutAuth := PublicRouter.Group("sysErrpr")
	{
		sysErrprRouter.DELETE("deleteSysError", sysErrprApi.DeleteSysError)           // 删除错误日志
		sysErrprRouter.DELETE("deleteSysErrorByIds", sysErrprApi.DeleteSysErrorByIds) // 批量删除错误日志
		sysErrprRouter.PUT("updateSysError", sysErrprApi.UpdateSysError)              // 更新错误日志
	}
	{
		sysErrprRouterWithoutRecord.GET("findSysError", sysErrprApi.FindSysError)       // 根据ID获取错误日志
		sysErrprRouterWithoutRecord.GET("getSysErrorList", sysErrprApi.GetSysErrorList) // 获取错误日志列表
	}
	{
		sysErrprRouterWithoutAuth.POST("createSysError", sysErrprApi.CreateSysError) // 新建错误日志
	}
}
