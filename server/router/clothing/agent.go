package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AgentRouter struct {
}

// InitAgentRouter 初始化 Agent 路由信息
func (s *AgentRouter) InitAgentRouter(Router *gin.RouterGroup) {
	agentRouter := Router.Group("agent").Use(middleware.OperationRecord())
	agentRouterWithoutRecord := Router.Group("agent")
	var agentApi = v1.ApiGroupApp.ClothingApiGroup.AgentApi
	{
		agentRouter.POST("createAgent", agentApi.CreateAgent)   // 新建Agent
		agentRouter.DELETE("deleteAgent", agentApi.DeleteAgent) // 删除Agent
		agentRouter.DELETE("deleteAgentByIds", agentApi.DeleteAgentByIds) // 批量删除Agent
		agentRouter.PUT("updateAgent", agentApi.UpdateAgent)    // 更新Agent
	}
	{
		agentRouterWithoutRecord.GET("findAgent", agentApi.FindAgent)        // 根据ID获取Agent
		agentRouterWithoutRecord.GET("getAgentList", agentApi.GetAgentList)  // 获取Agent列表
	}
}
