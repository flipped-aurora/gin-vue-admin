package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CompanyApplySearch struct {
	clothing.CompanyApply
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type OptCompanyApply struct {
	Status int `json:"status" form:"status"`
	ID     int `json:"ID" form:"ID"`
}
