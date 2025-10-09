package mcpTool

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/mark3labs/mcp-go/mcp"
)

// McpTool 定义了MCP工具必须实现的接口
type McpTool interface {
	// Handle 返回工具调用信息
	Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error)
	// New 返回工具注册信息
	New() mcp.Tool
}

// RegisterTool 供工具在init时调用，将自己注册到工具注册表中
func RegisterTool(tool McpTool) {
	global.GVA_MCP_SERVER.AddTool(tool.New(), tool.Handle)
}
