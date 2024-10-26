package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CliLoadSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	Address        string     `json:"address" form:"address" `
	Usdt           *float64   `json:"usdt" form:"usdt" `
	Status         string     `json:"status" form:"status" `
	Desc           string     `json:"desc" form:"desc" `
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
