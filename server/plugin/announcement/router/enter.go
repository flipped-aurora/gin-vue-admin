package router

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement/api"

var (
	Router  = new(router)
	apiInfo = api.Api.Info
)

type router struct{ Info info }
