package media

import (
	"github.com/gin-gonic/gin"
)

type FileUploadAndDownloadRouter struct{}

func (e *FileUploadAndDownloadRouter) InitFileUploadAndDownloadRouter(Router *gin.RouterGroup) {
	fileUploadAndDownloadRouter := Router.Group("fileUploadAndDownload")
	{
		fileUploadAndDownloadRouter.POST("upload", fileUploadAndDownloadApi.UploadFile)                                 // 上传文件
		fileUploadAndDownloadRouter.POST("getFileList", fileUploadAndDownloadApi.GetFileList)                           // 获取上传文件列表
		fileUploadAndDownloadRouter.POST("deleteFile", fileUploadAndDownloadApi.DeleteFile)                             // 删除指定文件
		fileUploadAndDownloadRouter.POST("deleteFiles", fileUploadAndDownloadApi.DeleteFiles)                           // 批量删除文件
		fileUploadAndDownloadRouter.GET("findFile", fileUploadAndDownloadApi.FindFile)                                  // 查找单个文件详情
		fileUploadAndDownloadRouter.POST("listOssFiles", fileUploadAndDownloadApi.ListOssFiles)                         // 列举存储桶文件
		fileUploadAndDownloadRouter.POST("editFileName", fileUploadAndDownloadApi.EditFileName)                         // 编辑文件名或者备注
		fileUploadAndDownloadRouter.POST("importURL", fileUploadAndDownloadApi.ImportURL)                               // 导入URL
	}
}
