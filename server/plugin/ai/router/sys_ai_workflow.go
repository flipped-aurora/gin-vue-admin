package router

import "github.com/gin-gonic/gin"

type AIWorkflowRouter struct{}

// InitAIWorkflowRouter 注册 AI 工作流相关路由，路径前缀沿用 autoCode，处理逻辑仍在 core 的 AIWorkflowSessionApi。
func (s *AIWorkflowRouter) InitAIWorkflowRouter(Router *gin.RouterGroup) {
	aiWorkflowRouter := Router.Group("autoCode")
	{
		aiWorkflowRouter.POST("saveAIWorkflowSession", aiWorkflowSessionApi.Save)
		aiWorkflowRouter.POST("getAIWorkflowSessionList", aiWorkflowSessionApi.GetList)
		aiWorkflowRouter.POST("getAIWorkflowSessionDetail", aiWorkflowSessionApi.GetDetail)
		aiWorkflowRouter.POST("deleteAIWorkflowSession", aiWorkflowSessionApi.Delete)
		aiWorkflowRouter.POST("dumpAIWorkflowMarkdown", aiWorkflowSessionApi.DumpMarkdown)
	}
}
