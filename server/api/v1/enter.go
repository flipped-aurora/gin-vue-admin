package v1

import (
	"gin-vue-admin/api/v1/autocode"
	"gin-vue-admin/api/v1/example"
	"gin-vue-admin/api/v1/system"
)

type ApiGroup struct {
	ExampleApiGroup  example.ApiGroup
	SystemApiGroup   system.ApiGroup
	AutoCodeApiGroup autocode.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
