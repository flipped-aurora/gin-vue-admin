package service

import (
	"gin-vue-admin/service/autocode"
	"gin-vue-admin/service/example"
	"gin-vue-admin/service/system"
)

type ServiceGroup struct {
	ExampleServiceGroup  example.ServiceGroup
	SystemServiceGroup   system.ServiceGroup
	AutoCodeServiceGroup autocode.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
