package client

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	model "github.com/flipped-aurora/gin-vue-admin/server/model/mcp"

	"github.com/ThinkInAIXYZ/go-mcp/client"
	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"github.com/ThinkInAIXYZ/go-mcp/transport"
)

type ClientStruct struct{}

var mcpClient *client.Client

// 链接mcp服务器
func (c *ClientStruct) NewSSE(Config *model.McpConfig) error {
	transportClient, err := transport.NewSSEClientTransport(Config.Url)
	if err != nil {
		return fmt.Errorf("创建mcp客户端失败SSE: %v", err)
	}
	mcpClient, err = client.NewClient(transportClient, client.WithClientInfo(protocol.Implementation{
		Name:    Config.Name,
		Version: Config.Version,
	}))
	if err != nil {
		return fmt.Errorf("创建mcp客户端失败: %v", err)
	}
	defer mcpClient.Close()
	fmt.Println("MCP客户端已启动")

	return nil
}

// 列出所有可用工具
func (c *ClientStruct) ListTools() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	toolsResult, err := mcpClient.ListTools(ctx)
	if err != nil {
		return "", fmt.Errorf("获取工具列表失败: %v", err)

	}
	b, _ := json.Marshal(toolsResult.Tools)
	return string(b), nil
}

// 调用工具
func (c *ClientStruct) CallTool(toolName string, inputReq map[string]interface{}) (string, error) {
	input := make(map[string]interface{})
	for k, v := range inputReq {
		input[k] = v
	}

	callToolResult, err := mcpClient.CallTool(context.Background(),
		protocol.NewCallToolRequest(toolName, input))
	if err != nil {
		return "", fmt.Errorf("使用工具失败: %v", err)
	}
	b, _ := json.Marshal(callToolResult)
	return string(b), nil
}

// 启动mcp client
func (c *ClientStruct) Run() {

	//连接McpServe服务器
	err := c.NewSSE(&model.McpConfig{
		Url:     "http://127.0.0.1:9999/sse",
		Name:    "mcpClient",
		Version: "1.0.0",
	})
	if err != nil || mcpClient == nil {
		fmt.Println("连接McpServe服务器失败")
		return
	}

	fmt.Println("McpServe服务器已启动")

	//(示例)
	{
		// 列出所有工具(示例)
		tools, err := c.ListTools()
		if err != nil {
			fmt.Println("获取工具列表失败:", err)
			return
		}
		fmt.Println("工具列表:", tools)

		// 调用工具(示例)
		result, err := c.CallTool("CurrentTime", nil)
		if err != nil {
			fmt.Println("使用工具失败:", err)
			return
		}
		fmt.Println("工具结果:", result)
	}
}
