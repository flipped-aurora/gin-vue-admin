package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysParamsRouter struct{}

// InitSysParamsRouter 初始化 参数 路由信息
func (s *SysParamsRouter) InitSysParamsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	sysParamsRouter := Router.Group("sysParams").Use(middleware.OperationRecord())
	sysParamsRouterWithoutRecord := Router.Group("sysParams")
	{
		sysParamsRouter.POST("createSysParams", sysParamsApi.CreateSysParams)             // 新建参数
		sysParamsRouter.DELETE("deleteSysParams", sysParamsApi.DeleteSysParams)           // 删除参数
		sysParamsRouter.DELETE("deleteSysParamsByIds", sysParamsApi.DeleteSysParamsByIds) // 批量删除参数
		sysParamsRouter.PUT("updateSysParams", sysParamsApi.UpdateSysParams)              // 更新参数
	}
	{
		sysParamsRouterWithoutRecord.GET("findSysParams", sysParamsApi.FindSysParams)       // 根据ID获取参数
		sysParamsRouterWithoutRecord.GET("getSysParamsList", sysParamsApi.GetSysParamsList) // 获取参数列表
		sysParamsRouterWithoutRecord.GET("getSysParam", sysParamsApi.GetSysParam)           // 根据Key获取参数
	}
}
