package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitWorkflowRouter(Router *gin.RouterGroup) {
	WorkflowRouter := Router.Group("workflow").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		WorkflowRouter.POST("createWorkFlow", v1.CreateWorkFlow) // 创建工作流
	}
}
