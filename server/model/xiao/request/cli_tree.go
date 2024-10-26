package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CliTreeSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	Address        string     `json:"address" form:"address" `
	Father         string     `json:"father" form:"father" `
	Invite         string     `json:"invite" form:"invite" `
	Description    string     `json:"description" form:"description" `
	request.PageInfo
}
