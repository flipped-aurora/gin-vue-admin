package main

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	mcpTool "github.com/flipped-aurora/gin-vue-admin/server/mcp"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
	_ "go.uber.org/automaxprocs"
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

	logger.Bg().Mod("cli").
		Field("config", configPath).
		Field("addr", addr).
		Field("path", global.GVA_CONFIG.MCP.Path).
		Field("upstream", global.GVA_CONFIG.MCP.UpstreamBaseURL).
		Info("mcp独立服务启动")

	if err := server.Start(addr); err != nil {
		logger.Bg().Mod("cli").Err(err).Error("mcp独立服务启动失败")
		panic(err)
	}
}
