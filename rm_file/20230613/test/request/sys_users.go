package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/test"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type SysUsersSearch struct{
    test.SysUsers
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
