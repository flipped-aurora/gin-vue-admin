package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysErrorRouter struct{}

// InitSysErrorRouter 初始化 错误日志 路由信息
func (s *SysErrorRouter) InitSysErrorRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
    sysErrorRouter := Router.Group("sysError").Use(middleware.OperationRecord())
    sysErrorRouterWithoutRecord := Router.Group("sysError")
    sysErrorRouterWithoutAuth := PublicRouter.Group("sysError")
    {
        sysErrorRouter.DELETE("deleteSysError", sysErrorApi.DeleteSysError)           // 删除错误日志
        sysErrorRouter.DELETE("deleteSysErrorByIds", sysErrorApi.DeleteSysErrorByIds) // 批量删除错误日志
        sysErrorRouter.PUT("updateSysError", sysErrorApi.UpdateSysError)              // 更新错误日志
        sysErrorRouter.GET("getSysErrorSolution", sysErrorApi.GetSysErrorSolution)    // 触发错误日志处理
    }
    {
        sysErrorRouterWithoutRecord.GET("findSysError", sysErrorApi.FindSysError)       // 根据ID获取错误日志
        sysErrorRouterWithoutRecord.GET("getSysErrorList", sysErrorApi.GetSysErrorList) // 获取错误日志列表
    }
    {
        sysErrorRouterWithoutAuth.POST("createSysError", sysErrorApi.CreateSysError) // 新建错误日志
    }
}
