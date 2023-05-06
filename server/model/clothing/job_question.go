// 自动生成模板JobQuestion
package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// JobQuestion 结构体
type JobQuestion struct {
	global.GVA_MODEL
	JobID     uint   `json:"jobID" form:"jobID" gorm:"column:job_id;comment:;"`
	UserID    uint   `json:"userID" form:"userID" gorm:"column:user_id;comment:;"`
	Content   string `json:"content" form:"content" gorm:"column:content;comment:;"`
	Status    *bool  `json:"status" form:"status" gorm:"column:status;comment:;"`
	CreatedBy uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName JobQuestion 表名
func (JobQuestion) TableName() string {
	return "job_question"
}
