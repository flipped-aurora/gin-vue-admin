package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	//"gorm.io/datatypes"
)

type DisInfoSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	DisTitle       string     `json:"disTitle" form:"disTitle" `
	DisComId       *int       `json:"disComId" form:"disComId" `
	DisUserId      *int       `json:"disUserId" form:"disUserId" `
	request.PageInfo
}
