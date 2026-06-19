package api

import autoService "github.com/flipped-aurora/gin-vue-admin/server/plugin/auto/service"

type ApiGroup struct {
	CliApi
}

var ApiGroupApp = new(ApiGroup)

var cliService = autoService.ServiceGroupApp.CliService
