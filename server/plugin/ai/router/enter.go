package router

import (
	api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	aiApi "github.com/flipped-aurora/gin-vue-admin/server/plugin/ai/api"
)

type RouterGroup struct {
	CliRouter
	SkillsRouter
	McpRouter
	AIWorkflowRouter
}

var (
	cliApi               = aiApi.ApiGroupApp.CliApi
	skillsApi            = api.ApiGroupApp.SystemApiGroup.SkillsApi
	autoCodeTemplateApi  = api.ApiGroupApp.SystemApiGroup.AutoCodeTemplateApi
	aiWorkflowSessionApi = api.ApiGroupApp.SystemApiGroup.AIWorkflowSessionApi
)

var RouterGroupApp = new(RouterGroup)
