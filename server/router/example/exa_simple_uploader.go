package example

import (
	"gin-vue-admin/api/v1"
	"github.com/gin-gonic/gin"
)

type SimpleUploaderRouter struct{}

func (e *SimpleUploaderRouter) InitSimpleUploaderRouter(Router *gin.RouterGroup) {
	simpleUploaderRouter := Router.Group("simpleUploader")
	var simpleUploaderApi = v1.ApiGroupApp.ExampleApiGroup.SimpleUploaderApi
	{
		simpleUploaderRouter.POST("upload", simpleUploaderApi.SimpleUploaderUpload) // 上传功能
		simpleUploaderRouter.GET("checkFileMd5", simpleUploaderApi.CheckFileMd5)    // 文件完整度验证
		simpleUploaderRouter.GET("mergeFileMd5", simpleUploaderApi.MergeFileMd5)    // 合并文件
	}
}
