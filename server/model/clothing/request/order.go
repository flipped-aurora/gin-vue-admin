package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type OrderSearch struct {
	clothing.Order
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type CreateOrderReq struct {
	OptID     uint `json:"optId"`
	CompanyID uint `json:"companyId"`
}

type PayOrderReq struct {
	OrderNo string `json:"orderNo"`
	PayType int    `json:"payType"`
	OpenID  string `json:"openId"`
}
