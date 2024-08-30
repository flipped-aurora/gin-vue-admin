package biz_apphub

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	BizAppHubApi
	BizToolCmdSrvApiApi
}

var (
	bizAppHubService        = service.ServiceGroupApp.Biz_apphubServiceGroup.BizAppHubService
	bizToolCmdSrvApiService = service.ServiceGroupApp.Biz_apphubServiceGroup.BizToolCmdSrvApiService
)
