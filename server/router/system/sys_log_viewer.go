package system

import "github.com/gin-gonic/gin"

type LogViewerRouter struct{}

func (l *LogViewerRouter) InitLogViewerRouter(router *gin.RouterGroup) {
	logViewerRouter := router.Group("logViewer")
	{
		logViewerRouter.GET("dates", logViewerApi.GetLogDates)
		logViewerRouter.GET("files", logViewerApi.GetLogFiles)
		logViewerRouter.GET("content", logViewerApi.GetLogContent)
	}
}
