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
		WorkflowProcessRouter.GET("findWorkflowStep", v1.FindWorkflowStep)                        // 根据ID获取工作流步骤
		WorkflowProcessRouter.GET("getWorkflowProcessList", v1.GetWorkflowProcessList)            // 获取WorkflowProcess列表
		WorkflowProcessRouter.POST("startWorkflow", v1.StartWorkflow)                             // 开启工作流
		WorkflowProcessRouter.GET("getMyStated", v1.GetMyStated)                                  // 获取我发起的工作流
		WorkflowProcessRouter.GET("getMyNeed", v1.GetMyNeed)                                      // 获取我的待办
	}
}
