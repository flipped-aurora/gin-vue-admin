package api

import aiService "github.com/flipped-aurora/gin-vue-admin/server/plugin/ai/service"

type ApiGroup struct {
	CliApi
	McpApi
}

var ApiGroupApp = new(ApiGroup)

var cliService = aiService.ServiceGroupApp.CliService
var mcpService = aiService.ServiceGroupApp.McpService
