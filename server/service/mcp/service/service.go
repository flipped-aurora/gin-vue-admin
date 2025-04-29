package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/initialize/mcp"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/mcp"
)

type McpService struct {
}

var Init = new(mcp.McpServer)

func (m *McpService) Run() {

	config := &model.McpConfig{
		Url:     `127.0.0.1:9999`,
		Name:    "mcpServer",
		Version: "1.0.0",
	}

	// 创建mcp服务器-sse方式
	Init.NewSSE(config)

	// 注册工具
	Init.BinTool()

	// 启动mcp服务器
	Init.Run()

}
