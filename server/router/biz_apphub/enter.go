package biz_apphub

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	BizAppHubRouter
	BizCmdToolApiRouter
}

var (
	bizAppHubApi     = api.ApiGroupApp.Biz_apphubApiGroup.BizAppHubApi
	bizCmdToolApiApi = api.ApiGroupApp.Biz_apphubApiGroup.BizCmdToolApiApi
)
