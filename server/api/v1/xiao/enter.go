package xiao

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	CliLoadApi
	CliUserApi
	CliTreeApi
	CliOrderApi
	CliMainorderApi
	CliMainprofitApi
	CliProfitApi
	CliWithdrawApi
	CliMainwithApi
	CliSetvipApi
	CliSetApi
}

var (
	cliLoadService       = service.ServiceGroupApp.XiaoServiceGroup.CliLoadService
	cliUserService       = service.ServiceGroupApp.XiaoServiceGroup.CliUserService
	cliTreeService       = service.ServiceGroupApp.XiaoServiceGroup.CliTreeService
	cliOrderService      = service.ServiceGroupApp.XiaoServiceGroup.CliOrderService
	climainorderService  = service.ServiceGroupApp.XiaoServiceGroup.CliMainorderService
	climainprofitService = service.ServiceGroupApp.XiaoServiceGroup.CliMainprofitService
	cliprofitService     = service.ServiceGroupApp.XiaoServiceGroup.CliProfitService
	cliwithdrawService   = service.ServiceGroupApp.XiaoServiceGroup.CliWithdrawService
	climainwithService   = service.ServiceGroupApp.XiaoServiceGroup.CliMainwithService
	clisetvipService     = service.ServiceGroupApp.XiaoServiceGroup.CliSetvipService
	clisetService        = service.ServiceGroupApp.XiaoServiceGroup.CliSetService
)
