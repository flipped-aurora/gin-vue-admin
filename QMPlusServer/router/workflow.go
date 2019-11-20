package router

import (
	"github.com/gin-gonic/gin"
	"main/controller/api"
)

func InitWorkflowRouter(Router *gin.Engine) {
	WorkflowRouter := Router.Group("workflow")
	//.Use(middleware.JWTAuth())
	{
		WorkflowRouter.POST("createWorkFlow", api.CreateWorkFlow) // 创建工作流
	}
}
