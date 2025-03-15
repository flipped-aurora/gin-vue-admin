package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysExportTemplateRouter struct {
}

// InitSysExportTemplateRouter 初始化 导出模板 路由信息
func (s *SysExportTemplateRouter) InitSysExportTemplateRouter(Router *gin.RouterGroup, pubRouter *gin.RouterGroup) {
	sysExportTemplateRouter := Router.Group("sysExportTemplate").Use(middleware.OperationRecord())
	sysExportTemplateRouterWithoutRecord := Router.Group("sysExportTemplate")
	sysExportTemplateRouterWithoutAuth := pubRouter.Group("sysExportTemplate")

	{
		sysExportTemplateRouter.POST("createSysExportTemplate", exportTemplateApi.CreateSysExportTemplate)             // 新建导出模板
		sysExportTemplateRouter.DELETE("deleteSysExportTemplate", exportTemplateApi.DeleteSysExportTemplate)           // 删除导出模板
		sysExportTemplateRouter.DELETE("deleteSysExportTemplateByIds", exportTemplateApi.DeleteSysExportTemplateByIds) // 批量删除导出模板
		sysExportTemplateRouter.PUT("updateSysExportTemplate", exportTemplateApi.UpdateSysExportTemplate)              // 更新导出模板
		sysExportTemplateRouter.POST("importExcel", exportTemplateApi.ImportExcel)                                     // 导入excel模板数据
	}
	{
		sysExportTemplateRouterWithoutRecord.GET("findSysExportTemplate", exportTemplateApi.FindSysExportTemplate)       // 根据ID获取导出模板
		sysExportTemplateRouterWithoutRecord.GET("getSysExportTemplateList", exportTemplateApi.GetSysExportTemplateList) // 获取导出模板列表
		sysExportTemplateRouterWithoutRecord.GET("exportExcel", exportTemplateApi.ExportExcel)                           // 获取导出token
		sysExportTemplateRouterWithoutRecord.GET("exportTemplate", exportTemplateApi.ExportTemplate)                     // 导出表格模板
	}
	{
		sysExportTemplateRouterWithoutAuth.GET("exportExcelByToken", exportTemplateApi.ExportExcelByToken)       // 通过token导出表格
		sysExportTemplateRouterWithoutAuth.GET("exportTemplateByToken", exportTemplateApi.ExportTemplateByToken) // 通过token导出模板
	}
}
