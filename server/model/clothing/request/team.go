package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type TeamSearch struct {
	clothing.Team
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type JoinTeam struct {
	UserID uint `json:"userID" form:"userID"`
	TeamID uint `json:"teamID" form:"teamID"`
}
