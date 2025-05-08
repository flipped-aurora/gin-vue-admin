package tool

import (
	"context"
	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"github.com/ThinkInAIXYZ/go-mcp/server"
	"github.com/flipped-aurora/gin-vue-admin/server/mcp/tool/current_time"
	"github.com/flipped-aurora/gin-vue-admin/server/mcp/tool/weather"
)

func init() {
	// 注册工具
	RegisterTool(&current_time.Tool{})
	RegisterTool(&weather.Tool{})
}

// McpTool 定义了MCP工具必须实现的接口
type McpTool interface {
	// Handle 返回工具调用信息
	Handle(ctx context.Context, request *protocol.CallToolRequest) (*protocol.CallToolResult, error)
	// New 返回工具注册信息
	New() (name string, description string, inputReqStruct interface{})
}

// 工具注册表
var toolRegister = make(map[string]McpTool)

// RegisterTool 供工具在init时调用，将自己注册到工具注册表中
func RegisterTool(tool McpTool) {
	name, _, _ := tool.New()
	toolRegister[name] = tool
}

// RegisterAllTools 将所有注册的工具注册到MCP服务中
func RegisterAllTools(mcpServer *server.Server) {
	for _, tool := range toolRegister {
		t, e := protocol.NewTool(tool.New())
		if e != nil {
			panic("tool register error: " + e.Error())
		}
		mcpServer.RegisterTool(t, tool.Handle)
	}
}
