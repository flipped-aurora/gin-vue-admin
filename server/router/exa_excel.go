package router

import (
	"gin-vue-admin/api/v1"
	"github.com/gin-gonic/gin"
)

func InitExcelRouter(Router *gin.RouterGroup) {
	ExcelRouter := Router.Group("excel")
	{
		ExcelRouter.POST("/importExcel", v1.ImportExcel)          // 导入Excel
		ExcelRouter.GET("/loadExcel", v1.LoadExcel)               // 加载Excel数据
		ExcelRouter.POST("/exportExcel", v1.ExportExcel)          // 导出Excel
		ExcelRouter.GET("/downloadTemplate", v1.DownloadTemplate) // 下载模板文件
	}
}
