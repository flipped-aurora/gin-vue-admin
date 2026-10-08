package example

import (
	"github.com/gin-gonic/gin"
)

type FileUploadAndDownloadRouter struct{}

func (e *FileUploadAndDownloadRouter) InitFileUploadAndDownloadRouter(PrivateRouter *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	fileUploadAndDownloadRouter := PrivateRouter.Group("fileUploadAndDownload")
	scanUploadRouter := PublicRouter.Group("fileUploadAndDownload")
	{
		fileUploadAndDownloadRouter.POST("upload", exaFileUploadAndDownloadApi.UploadFile)                                 // 上传文件
		fileUploadAndDownloadRouter.POST("createScanUploadToken", exaFileUploadAndDownloadApi.CreateScanUploadToken)       // 创建扫码上传凭证
		fileUploadAndDownloadRouter.POST("getFileList", exaFileUploadAndDownloadApi.GetFileList)                           // 获取上传文件列表
		fileUploadAndDownloadRouter.POST("deleteFile", exaFileUploadAndDownloadApi.DeleteFile)                             // 删除指定文件
		fileUploadAndDownloadRouter.POST("editFileName", exaFileUploadAndDownloadApi.EditFileName)                         // 编辑文件名或者备注
		fileUploadAndDownloadRouter.POST("breakpointContinue", exaFileUploadAndDownloadApi.BreakpointContinue)             // 断点续传
		fileUploadAndDownloadRouter.GET("findFile", exaFileUploadAndDownloadApi.FindFile)                                  // 查询当前文件成功的切片
		fileUploadAndDownloadRouter.POST("breakpointContinueFinish", exaFileUploadAndDownloadApi.BreakpointContinueFinish) // 切片传输完成
		fileUploadAndDownloadRouter.POST("removeChunk", exaFileUploadAndDownloadApi.RemoveChunk)                           // 删除切片
		fileUploadAndDownloadRouter.POST("importURL", exaFileUploadAndDownloadApi.ImportURL)                               // 导入URL
	}
	{
		scanUploadRouter.POST("uploadByScanToken", exaFileUploadAndDownloadApi.UploadFileByScanToken) // 使用一次性凭证扫码上传
	}
}
