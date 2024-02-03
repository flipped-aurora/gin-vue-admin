package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type ComInfoSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
                      ComTitle  string `json:"comTitle" form:"comTitle" `
                      ComModel  *int `json:"comModel" form:"comModel" `
                StartComStart  *time.Time  `json:"startComStart" form:"startComStart"`
                EndComStart  *time.Time  `json:"endComStart" form:"endComStart"`
                StartComEnd  *time.Time  `json:"startComEnd" form:"startComEnd"`
                EndComEnd  *time.Time  `json:"endComEnd" form:"endComEnd"`
                      ComType  *int `json:"comType" form:"comType" `
    request.PageInfo
}
