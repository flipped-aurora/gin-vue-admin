package xiao

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	CliLoadRouter
	CliUserRouter
	CliTreeRouter
	CliOrderRouter
	CliMainorderRouter
	CliMainprofitRouter
	CliProfitRouter
	CliWithdrawRouter
	CliMainwithRouter
	CliSetvipRouter
	CliSetRouter
}

var (
	cliLoadApi       = api.ApiGroupApp.XiaoApiGroup.CliLoadApi
	cliUserApi       = api.ApiGroupApp.XiaoApiGroup.CliUserApi
	cliTreeApi       = api.ApiGroupApp.XiaoApiGroup.CliTreeApi
	cliOrderApi      = api.ApiGroupApp.XiaoApiGroup.CliOrderApi
	climainorderApi  = api.ApiGroupApp.XiaoApiGroup.CliMainorderApi
	climainprofitApi = api.ApiGroupApp.XiaoApiGroup.CliMainprofitApi
	cliprofitApi     = api.ApiGroupApp.XiaoApiGroup.CliProfitApi
	cliwithdrawApi   = api.ApiGroupApp.XiaoApiGroup.CliWithdrawApi
	climainwithApi   = api.ApiGroupApp.XiaoApiGroup.CliMainwithApi
	clisetvipApi     = api.ApiGroupApp.XiaoApiGroup.CliSetvipApi
	clisetApi        = api.ApiGroupApp.XiaoApiGroup.CliSetApi
)
