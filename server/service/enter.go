package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/NestInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/service/Nestrolepkg"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup      system.ServiceGroup
	ExampleServiceGroup     example.ServiceGroup
	NestInfoServiceGroup    NestInfo.ServiceGroup
	NestrolepkgServiceGroup Nestrolepkg.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
