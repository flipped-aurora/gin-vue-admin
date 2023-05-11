package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/reservation"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup      system.ApiGroup
	ExampleApiGroup     example.ApiGroup
	ReservationApiGroup reservation.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
