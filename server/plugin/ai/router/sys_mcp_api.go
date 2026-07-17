package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type McpApiRouter struct{}

func (s *McpApiRouter) InitMcpApiRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	mcpApiRouter := Router.Group("mcpApi").Use(middleware.OperationRecord())
	mcpApiRouterWithoutRecord := Router.Group("mcpApi")
	{
		mcpApiRouter.POST("createMcp", mcpApi.CreateMcp)
		mcpApiRouter.POST("updateMcp", mcpApi.UpdateMcp)
		mcpApiRouter.POST("deleteMcp", mcpApi.DeleteMcp)
		mcpApiRouter.POST("addMcpApis", mcpApi.AddMcpApis)
		mcpApiRouter.POST("removeMcpApis", mcpApi.RemoveMcpApis)
	}
	{
		mcpApiRouterWithoutRecord.POST("getMcpList", mcpApi.GetMcpList)
		mcpApiRouterWithoutRecord.POST("getMcpDetail", mcpApi.GetMcpDetail)
		mcpApiRouterWithoutRecord.POST("previewManifest", mcpApi.PreviewMcpManifest)
		mcpApiRouterWithoutRecord.POST("previewPrompt", mcpApi.PreviewMcpPrompt)
		mcpApiRouterWithoutRecord.POST("previewApiCommand", mcpApi.PreviewApiCommand)
	}
	// 公开接口（不走鉴权）：供 MCP 独立进程启动时读取动态 tool 的能力定义
	{
		RouterPub.GET("mcpApi/listBindingsPublic", mcpApi.ListMcpBindingsPublic)
		RouterPub.GET("mcpApi/listPromptsPublic", mcpApi.ListMcpPromptsPublic)
	}
}
