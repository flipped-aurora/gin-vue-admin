package mcpTool

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	mcpServer "github.com/mark3labs/mcp-go/server"
)

// promptDef 对应 GVA ListMcpPromptsPublic 接口返回的单条 prompt 定义。
type promptDef struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Markdown    string `json:"markdown"`
}

// promptsResponse 是 GVA ListMcpPromptsPublic 接口的响应结构（复用 GVA 标准信封）。
type promptsResponse struct {
	Code int `json:"code"`
	Data struct {
		Prompts []promptDef `json:"prompts"`
	} `json:"data"`
	Msg string `json:"msg"`
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
	endpoint := strings.TrimRight(upstreamBaseURL(), "/") + "/mcpApi/listPromptsPublic"
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(timeoutCtx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("构建请求失败: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求 GVA prompt 接口失败: %w", err)
	}
	defer resp.Body.Close()

	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("GVA prompt 接口返回状态码 %d: %s", resp.StatusCode, string(rawBody))
	}

	var result promptsResponse
	if err := json.Unmarshal(rawBody, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}
	if result.Code != 0 {
		return nil, fmt.Errorf("GVA prompt 接口业务错误: %s", result.Msg)
	}
	return result.Data.Prompts, nil
}
