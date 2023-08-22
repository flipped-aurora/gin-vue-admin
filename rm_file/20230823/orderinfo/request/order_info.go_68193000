package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/orderinfo"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type OrderInfoSearch struct{
    orderinfo.OrderInfo
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
