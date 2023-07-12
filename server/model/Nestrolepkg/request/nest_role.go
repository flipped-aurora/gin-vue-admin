package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/Nestrolepkg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type NestRoleSearch struct {
	Nestrolepkg.NestRole
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
