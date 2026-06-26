package media

import "github.com/gin-gonic/gin"

type MediaUploadRouter struct{}

func (r *MediaUploadRouter) InitMediaUploadRouter(Router *gin.RouterGroup) {
	g := Router.Group("mediaUpload")
	{
		g.POST("init", mediaUploadApi.Init)
		g.POST("chunk", mediaUploadApi.Chunk)
		g.POST("complete", mediaUploadApi.Complete)
		g.DELETE(":uploadId", mediaUploadApi.Cancel)
	}
}
