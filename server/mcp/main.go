package mcp

import (
	"fmt"
	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"github.com/ThinkInAIXYZ/go-mcp/server"
	"github.com/ThinkInAIXYZ/go-mcp/transport"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/mcp/tool"
)

func Run() (transport.ServerTransport, *transport.SSEHandler, error) {
	// 初始化SSE服务器
	config := global.GVA_CONFIG.MCP
	transportServer, mcpHandle, err := transport.NewSSEServerTransportAndHandler(config.MessagePath,
		transport.WithSSEServerTransportAndHandlerOptionLogger(global.GVA_LOG.Sugar()),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("创建SSE服务器失败: %v", err)
	}
	mcpServer, err := server.NewServer(transportServer,
		server.WithServerInfo(protocol.Implementation{
			Name:    config.Name,
			Version: config.Version,
		}),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("创建MCP服务器失败: %v", err)
	}

	tool.RegisterAllTools(mcpServer)

	//启动服务器
	go func() {
		err = mcpServer.Run()
	}()

	return transportServer, mcpHandle, err
}
