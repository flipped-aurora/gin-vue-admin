package client

import (
	"fmt"
	"github.com/ThinkInAIXYZ/go-mcp/client"
	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"github.com/ThinkInAIXYZ/go-mcp/transport"
)

type Client struct{}

var mcpClient *client.Client

// 链接mcp服务器
func (c *Client) New() error {
	transportClient, err := transport.NewSSEClientTransport("http://localhost:8888/sse")
	if err != nil {
		return fmt.Errorf("创建mcp客户端失败SSE: %v", err)
	}
	mcpClient, err = client.NewClient(transportClient, client.WithClientInfo(protocol.Implementation{
		Name:    "client",
		Version: "v0.0.1",
	}))
	if err != nil {
		return fmt.Errorf("创建mcp客户端失败: %v", err)
	}
	defer mcpClient.Close()
	fmt.Println("MCP客户端已启动")

	return nil
}
