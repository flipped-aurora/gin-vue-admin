package mcpTool

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// McpTool 定义了MCP工具必须实现的接口
type McpTool interface {
	// Handle 返回工具调用信息
	Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error)
	// New 返回工具注册信息
	New() mcp.Tool
}

// 工具注册表
var toolRegister = make(map[string]McpTool)

// RegisterTool 供工具在init时调用，将自己注册到工具注册表中
func RegisterTool(tool McpTool) {
	mcpTool := tool.New()
	toolRegister[mcpTool.Name] = tool
}

// RegisterAllTools 将所有注册的工具注册到MCP服务中
func RegisterAllTools(mcpServer *server.MCPServer) {
	for _, tool := range toolRegister {
		mcpServer.AddTool(tool.New(), tool.Handle)
	}
}
