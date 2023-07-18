package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/UserTeemlinkPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type UserTeemlinkSearch struct{
    UserTeemlinkPkg.UserTeemlink
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
