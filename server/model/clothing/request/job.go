package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type JobSearch struct {
	clothing.Job
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type JobFilter struct {
	CroppingID uint `json:"croppingID" form:"croppingID" gorm:"column:cropping_id;comment:;"`
}

type PostJob struct {
	UserID    uint `json:"userId"`
	ProcessID uint `json:"processId"`
}

type JobList struct {
	CroppingID uint      `json:"croppingId"` // 裁剪单id
	TeamID     uint      `json:"teamId"`
	Jobs       []PostJob `json:"jobs"`
}

type JobAuditApply struct {
	ID           uint `json:"id"`
	Type         int  `json:"type"`
	TeamID       uint `json:"teamID"`
	StyleID      uint `json:"styleID"`
	ProcessID    uint `json:"processID"`
	RealQuantity uint `json:"realQuantity"`
}

type JobAuditOpt struct {
	ID     uint `json:"id"`
	Status bool `json:"status"`
}
