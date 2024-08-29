package biz_apphub

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	BizAppHubApi
	BizCmdToolApiApi
}

var (
	bizAppHubService     = service.ServiceGroupApp.Biz_apphubServiceGroup.BizAppHubService
	bizCmdToolApiService = service.ServiceGroupApp.Biz_apphubServiceGroup.BizCmdToolApiService
)
