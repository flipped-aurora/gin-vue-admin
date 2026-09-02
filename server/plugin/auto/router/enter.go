package router

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	AutoCodeRouter
	SkillsRouter
}

var (
	autoCodeApi          = api.ApiGroupApp.SystemApiGroup.AutoCodeApi
	autoCodePluginApi    = api.ApiGroupApp.SystemApiGroup.AutoCodePluginApi
	autocodeHistoryApi   = api.ApiGroupApp.SystemApiGroup.AutoCodeHistoryApi
	autoCodePackageApi   = api.ApiGroupApp.SystemApiGroup.AutoCodePackageApi
	autoCodeTemplateApi  = api.ApiGroupApp.SystemApiGroup.AutoCodeTemplateApi
	skillsApi            = api.ApiGroupApp.SystemApiGroup.SkillsApi
	aiWorkflowSessionApi = api.ApiGroupApp.SystemApiGroup.AIWorkflowSessionApi
)

var RouterGroupApp = new(RouterGroup)
