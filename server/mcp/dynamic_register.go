package mcpTool

import (
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	mcpServer "github.com/mark3labs/mcp-go/server"

	autoRes "github.com/flipped-aurora/gin-vue-admin/server/plugin/ai/model/response"
)

// dynamicBindingsData 是 GVA ListMcpBindingsPublic 接口 data 字段的负载（信封由 upstreamEnvelope 提供）。
type dynamicBindingsData struct {
	Commands []autoRes.SysCliManifestCommand `json:"commands"`
}

// registerDynamicTools 从 GVA 主服务读取动态 tool 绑定，注册成 MCP tool。
// 在 NewMCPServer 里 RegisterAllTools 之后调用。
func registerDynamicTools(s *mcpServer.MCPServer) {
	commands, err := fetchDynamicCommands()
	if err != nil {
		// 动态注册失败不阻止 server 启动（静态 tool 仍可用），只记录
		fmt.Printf("[mcp] 动态 tool 注册失败(将仅使用静态 tool): %v\n", err)
		return
	}
	if len(commands) == 0 {
		return
	}
	count := 0
	for _, cmd := range commands {
		tool, err := manifestCommandToTool(cmd)
		if err != nil {
			fmt.Printf("[mcp] 跳过动态 tool %q: %v\n", cmd.Name, err)
			continue
		}
		s.AddTool(tool, makeDynamicHandler(cmd))
		count++
	}
	fmt.Printf("[mcp] 动态注册了 %d 个 tool\n", count)
}

// fetchDynamicCommands 通过 HTTP 回打 GVA 的 ListMcpBindingsPublic 接口读取能力定义。
func fetchDynamicCommands() ([]autoRes.SysCliManifestCommand, error) {
	result, err := fetchPublicUpstream[dynamicBindingsData]("/mcpApi/listBindingsPublic")
	if err != nil {
		return nil, err
	}
	return result.Data.Commands, nil
}

// 确保 mcp 包编译时引用了被使用的类型（避免 IDE 报未使用）
var _ mcp.CallToolRequest
