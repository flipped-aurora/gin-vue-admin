package api

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement/service"

var (
	Api         = new(api)
	serviceInfo = service.Service.Info
)

type api struct{ Info info }
