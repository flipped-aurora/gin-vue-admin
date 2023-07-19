package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/FlyResultPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type FlyResultSearch struct {
	FlyResultPkg.FlyResult
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
