package router

import "github.com/gin-gonic/gin"

type McpRouter struct{}

// InitMcpRouter 注册 MCP Tools 相关路由，路径前缀沿用 autoCode，处理逻辑仍在 core 的 AutoCodeTemplateApi。
func (s *McpRouter) InitMcpRouter(Router *gin.RouterGroup) {
	mcpRouter := Router.Group("autoCode")
	{
		mcpRouter.POST("mcp", autoCodeTemplateApi.MCP)
		mcpRouter.POST("mcpStatus", autoCodeTemplateApi.MCPStatus)
		mcpRouter.POST("mcpStart", autoCodeTemplateApi.MCPStart)
		mcpRouter.POST("mcpStop", autoCodeTemplateApi.MCPStop)
		mcpRouter.POST("mcpList", autoCodeTemplateApi.MCPList)
		mcpRouter.POST("mcpRoutes", autoCodeTemplateApi.MCPRoutes)
		mcpRouter.POST("mcpTest", autoCodeTemplateApi.MCPTest)
	}
}
