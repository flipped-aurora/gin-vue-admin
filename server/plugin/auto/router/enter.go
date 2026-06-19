package router

import (
	api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	autoApi "github.com/flipped-aurora/gin-vue-admin/server/plugin/auto/api"
)

type RouterGroup struct {
	AutoCodeRouter
	SkillsRouter
	CliRouter
}

var (
	autoCodeApi          = api.ApiGroupApp.SystemApiGroup.AutoCodeApi
	autoCodePluginApi    = api.ApiGroupApp.SystemApiGroup.AutoCodePluginApi
	autocodeHistoryApi   = api.ApiGroupApp.SystemApiGroup.AutoCodeHistoryApi
	autoCodePackageApi   = api.ApiGroupApp.SystemApiGroup.AutoCodePackageApi
	autoCodeTemplateApi  = api.ApiGroupApp.SystemApiGroup.AutoCodeTemplateApi
	skillsApi            = api.ApiGroupApp.SystemApiGroup.SkillsApi
	aiWorkflowSessionApi = api.ApiGroupApp.SystemApiGroup.AIWorkflowSessionApi
	cliApi               = autoApi.ApiGroupApp.CliApi
)

var RouterGroupApp = new(RouterGroup)
