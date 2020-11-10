package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitWorkflowProcessRouter(Router *gin.RouterGroup) {
	WorkflowProcessRouter := Router.Group("workflowProcess").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		WorkflowProcessRouter.POST("createWorkflowProcess", v1.CreateWorkflowProcess)             // 新建WorkflowProcess
		WorkflowProcessRouter.DELETE("deleteWorkflowProcess", v1.DeleteWorkflowProcess)           // 删除WorkflowProcess
		WorkflowProcessRouter.DELETE("deleteWorkflowProcessByIds", v1.DeleteWorkflowProcessByIds) // 批量删除WorkflowProcess
		WorkflowProcessRouter.PUT("updateWorkflowProcess", v1.UpdateWorkflowProcess)              // 更新WorkflowProcess
		WorkflowProcessRouter.GET("findWorkflowProcess", v1.FindWorkflowProcess)                  // 根据ID获取WorkflowProcess
		WorkflowProcessRouter.GET("findWorkflowCreateStep", v1.FindWorkflowCreateStep)            // 根据ID获取工作流开启步骤
		WorkflowProcessRouter.GET("getWorkflowProcessList", v1.GetWorkflowProcessList)            // 获取WorkflowProcess列表
	}
}
