package router

import (
	"gin-vue-admin/api/v1"
	"github.com/gin-gonic/gin"
)

func InitFileUploadAndDownloadRouter(Router *gin.RouterGroup) {
	FileUploadAndDownloadGroup := Router.Group("fileUploadAndDownload")
	{
		FileUploadAndDownloadGroup.POST("/upload", v1.UploadFile)                                 // 上传文件
		FileUploadAndDownloadGroup.POST("/getFileList", v1.GetFileList)                           // 获取上传文件列表
		FileUploadAndDownloadGroup.POST("/deleteFile", v1.DeleteFile)                             // 删除指定文件
		FileUploadAndDownloadGroup.POST("/breakpointContinue", v1.BreakpointContinue)             // 断点续传
		FileUploadAndDownloadGroup.GET("/findFile", v1.FindFile)                                  // 查询当前文件成功的切片
		FileUploadAndDownloadGroup.POST("/breakpointContinueFinish", v1.BreakpointContinueFinish) // 查询当前文件成功的切片
		FileUploadAndDownloadGroup.POST("/removeChunk", v1.RemoveChunk)                           // 查询当前文件成功的切片
		FileUploadAndDownloadGroup.POST("/importExcel", v1.ImportExcel)                          // 导入Excel
		FileUploadAndDownloadGroup.GET("/loadExcel", v1.LoadExcel)                                // 加载Excel数据
		FileUploadAndDownloadGroup.POST("/exportExcel", v1.ExportExcel)                           // 导出Excel
		FileUploadAndDownloadGroup.GET("/downloadTemplate", v1.DownloadTemplate)                  // 下载模板文件
	}
}
