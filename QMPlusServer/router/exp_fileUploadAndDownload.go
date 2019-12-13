package router

import (
	"github.com/gin-gonic/gin"
	"main/controller/api"
	"main/middleware"
)

func InitFileUploadAndDownloadRouter(Router *gin.RouterGroup) {
	FileUploadAndDownloadGroup := Router.Group("fileUploadAndDownload").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	//.Use(middleware.JWTAuth())
	{
		FileUploadAndDownloadGroup.POST("/upload", api.UploadFile)       // 上传文件
		FileUploadAndDownloadGroup.POST("/getFileList", api.GetFileList) // 获取上传文件列表
		FileUploadAndDownloadGroup.POST("/deleteFile", api.DeleteFile)   // 删除指定文件
	}
}
