package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppUserSearch struct {
	clothing.AppUser
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	StartWages     *float64   `json:"startWages" form:"startWages"`
	EndWages       *float64   `json:"endWages" form:"endWages"`
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}

type Login struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}
