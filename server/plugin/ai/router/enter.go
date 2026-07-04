package router

import (
	api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	aiApi "github.com/flipped-aurora/gin-vue-admin/server/plugin/ai/api"
)

type RouterGroup struct {
	CliRouter
	SkillsRouter
	McpRouter
	McpApiRouter
}

var (
	cliApi              = aiApi.ApiGroupApp.CliApi
	mcpApi              = aiApi.ApiGroupApp.McpApi
	skillsApi           = api.ApiGroupApp.SystemApiGroup.SkillsApi
	autoCodeTemplateApi = api.ApiGroupApp.SystemApiGroup.AutoCodeTemplateApi
)

var RouterGroupApp = new(RouterGroup)
