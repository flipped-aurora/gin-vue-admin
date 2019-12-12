package router

import (
	"github.com/gin-gonic/gin"
	"main/controller/api"
	"main/middleware"
)

func InitWorkflowRouter(Router *gin.Engine)(R gin.IRoutes)  {
	WorkflowRouter := Router.Group("workflow").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		WorkflowRouter.POST("createWorkFlow", api.CreateWorkFlow) // 创建工作流
	}
	return WorkflowRouter
}
