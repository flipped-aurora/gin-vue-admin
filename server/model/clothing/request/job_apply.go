package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type JobApplySearch struct {
	clothing.JobApply
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type OptJobApply struct {
	Status int `json:"status" form:"status"`
	ID     int `json:"ID" form:"ID"`
}
