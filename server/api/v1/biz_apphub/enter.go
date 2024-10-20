package biz_apphub

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	BizAppHubApi
	BizToolCmdSrvApiApi
	BizCloudFunctionApi
	BizFuncPackageApi
}

var (
	bizAppHubService        = service.ServiceGroupApp.Biz_apphubServiceGroup.BizAppHubService
	bizToolCmdSrvApiService = service.ServiceGroupApp.Biz_apphubServiceGroup.BizToolCmdSrvApiService
	bizCloudFunctionService = service.ServiceGroupApp.Biz_apphubServiceGroup.BizCloudFunctionService
	bizFuncPackageService   = service.ServiceGroupApp.Biz_apphubServiceGroup.BizFuncPackageService
)
