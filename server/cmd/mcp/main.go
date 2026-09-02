package main

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	mcpTool "github.com/flipped-aurora/gin-vue-admin/server/mcp"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
)

func main() {
	configPath, err := loadStandaloneConfig()
	if err != nil {
		panic(err)
	}

	if err := initializeStandaloneLogger(); err != nil {
		panic(err)
	}

	addr := fmt.Sprintf(":%d", global.GVA_CONFIG.MCP.Addr)
	server := mcpTool.NewStreamableHTTPServer()

	global.GVA_LOG.Info("mcp独立服务启动",
		zap.String("config", configPath),
		zap.String("addr", addr),
		zap.String("path", global.GVA_CONFIG.MCP.Path),
		zap.String("upstream", global.GVA_CONFIG.MCP.UpstreamBaseURL),
	)

	if err := server.Start(addr); err != nil {
		global.GVA_LOG.Fatal("mcp独立服务启动失败", zap.Error(err))
	}
}
