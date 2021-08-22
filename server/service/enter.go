package service

import (
	"github.com/flipped-aurora/gin-vue-admin/service/autocode"
	"github.com/flipped-aurora/gin-vue-admin/service/example"
	"github.com/flipped-aurora/gin-vue-admin/service/system"
)

type ServiceGroup struct {
	ExampleServiceGroup  example.ServiceGroup
	SystemServiceGroup   system.ServiceGroup
	AutoCodeServiceGroup autocode.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
