package mcpTool

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	mcpServer "github.com/mark3labs/mcp-go/server"
)

// promptDef 对应 GVA ListMcpPromptsPublic 接口返回的单条 prompt 定义。
type promptDef struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Markdown    string `json:"markdown"`
}

// promptsData 是 GVA ListMcpPromptsPublic 接口 data 字段的负载（信封由 upstreamEnvelope 提供）。
type promptsData struct {
	Prompts []promptDef `json:"prompts"`
}

// registerDynamicPrompts 从 GVA 主服务读取编排 prompt 定义，注册成 MCP prompt。
// 在 NewMCPServer 里 registerDynamicTools 之后调用。
func registerDynamicPrompts(s *mcpServer.MCPServer) {
	defs, err := fetchDynamicPrompts()
	if err != nil {
		// prompt 注册失败不阻止 server 启动
		fmt.Printf("[mcp] 动态 prompt 注册失败(将仅使用 tool): %v\n", err)
		return
	}
	if len(defs) == 0 {
		return
	}
	count := 0
	for _, def := range defs {
		// 闭包捕获当前 def（避免循环变量问题）
		d := def
		prompt := mcp.NewPrompt(d.Name, mcp.WithPromptDescription(d.Description))
		handler := func(ctx context.Context, request mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
			return &mcp.GetPromptResult{
				Description: d.Description,
				Messages: []mcp.PromptMessage{
					{
						Role:    mcp.RoleUser,
						Content: mcp.NewTextContent(d.Markdown),
					},
				},
			}, nil
		}
		s.AddPrompt(prompt, handler)
		count++
	}
	fmt.Printf("[mcp] 动态注册了 %d 个编排 prompt\n", count)
}

// fetchDynamicPrompts 通过 HTTP 回打 GVA 的 ListMcpPromptsPublic 接口读取编排定义。
func fetchDynamicPrompts() ([]promptDef, error) {
	result, err := fetchPublicUpstream[promptsData]("/mcpApi/listPromptsPublic")
	if err != nil {
		return nil, err
	}
	return result.Data.Prompts, nil
}
