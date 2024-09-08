package biz_apphub

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	BizAppHubRouter
	BizToolCmdSrvApiRouter
	BizCloudFunctionRouter
}

var (
	bizAppHubApi        = api.ApiGroupApp.Biz_apphubApiGroup.BizAppHubApi
	bizToolCmdSrvApiApi = api.ApiGroupApp.Biz_apphubApiGroup.BizToolCmdSrvApiApi
	bizCloudFunctionApi = api.ApiGroupApp.Biz_apphubApiGroup.BizCloudFunctionApi
)
