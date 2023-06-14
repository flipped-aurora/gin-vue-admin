package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ProcessSearch struct {
	clothing.Process
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type CreateProcess struct {
	Type    int  `json:"type"`
	StyleID uint `json:"styleID"`
	Detail  []struct {
		Name    string  `json:"name"`
		Price   float64 `json:"price"`
		Percent float64 `json:"percent"`
	} `json:"detail"`
}
