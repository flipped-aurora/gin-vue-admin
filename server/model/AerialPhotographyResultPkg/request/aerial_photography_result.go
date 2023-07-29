package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/AerialPhotographyResultPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type AerialPhotographyResultSearch struct {
	AerialPhotographyResultPkg.AerialPhotographyResult
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
