package service

import (
	"github.com/gzpz/golf-sales-system/server/service/example"
	"github.com/gzpz/golf-sales-system/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
