package example

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	CustomerApi
}

var (
	customerService = service.ServiceGroupApp.ExampleServiceGroup.CustomerService
)
