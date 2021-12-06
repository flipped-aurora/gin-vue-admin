package example

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type ExcelRouter struct{}

func (e *ExcelRouter) InitExcelRouter(Router *gin.RouterGroup) {
	excelRouter := Router.Group("excel")
	exaExcelApi := v1.ApiGroupApp.ExampleApiGroup.ExcelApi
	{
		excelRouter.POST("importExcel", exaExcelApi.ImportExcel)          // 导入Excel
		excelRouter.GET("loadExcel", exaExcelApi.LoadExcel)               // 加载Excel数据
		excelRouter.POST("exportExcel", exaExcelApi.ExportExcel)          // 导出Excel
		excelRouter.GET("downloadTemplate", exaExcelApi.DownloadTemplate) // 下载模板文件
	}
}
