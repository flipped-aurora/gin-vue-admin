package client

import (
	"fmt"
	"github.com/ThinkInAIXYZ/go-mcp/client"
	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"github.com/ThinkInAIXYZ/go-mcp/transport"
)

// 创建MCP客户端
func NewClient(url string, clientInfo protocol.Implementation) (*client.Client, error) {
	transportClient, err := transport.NewSSEClientTransport(url)
	if err != nil {
		return nil, fmt.Errorf("创建MCP客户端失败: %v", err)
	}
	mcpClient, err := client.NewClient(transportClient, client.WithClientInfo(clientInfo))
	if err != nil {
		return nil, fmt.Errorf("创建MCP客户端失败: %v", err)
	}
	return mcpClient, nil
}
