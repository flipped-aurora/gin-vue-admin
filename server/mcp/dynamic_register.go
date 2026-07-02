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

	autoRes "github.com/flipped-aurora/gin-vue-admin/server/plugin/ai/model/response"
)

// dynamicBindingsResponse 是 GVA ListMcpBindingsPublic 接口的响应结构（复用 GVA 标准信封）。
type dynamicBindingsResponse struct {
	Code int `json:"code"`
	Data struct {
		Commands []autoRes.SysCliManifestCommand `json:"commands"`
	} `json:"data"`
	Msg string `json:"msg"`
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
	endpoint := strings.TrimRight(upstreamBaseURL(), "/") + "/mcpApi/listBindingsPublic"
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(timeoutCtx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("构建请求失败: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求 GVA 绑定接口失败: %w", err)
	}
	defer resp.Body.Close()

	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("GVA 绑定接口返回状态码 %d: %s", resp.StatusCode, string(rawBody))
	}

	var result dynamicBindingsResponse
	if err := json.Unmarshal(rawBody, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}
	if result.Code != 0 {
		return nil, fmt.Errorf("GVA 绑定接口业务错误: %s", result.Msg)
	}
	return result.Data.Commands, nil
}

// 确保 mcp 包编译时引用了被使用的类型（避免 IDE 报未使用）
var _ mcp.CallToolRequest
