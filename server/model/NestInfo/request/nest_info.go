package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/NestInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type NestInfoSearch struct{
    NestInfo.NestInfo
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
