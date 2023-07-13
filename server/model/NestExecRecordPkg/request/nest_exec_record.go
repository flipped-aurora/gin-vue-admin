package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/NestExecRecordPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type NestExecRecordSearch struct {
	NestExecRecordPkg.NestExecRecord
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
