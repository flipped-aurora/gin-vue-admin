// 自动生成模板Agent
package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Agent 结构体
type Agent struct {
	global.GVA_MODEL
	Name        string `json:"name" form:"name" gorm:"column:name;comment:;"`
	MemberCount uint   `json:"memberCount" form:"memberCount" gorm:"column:member_count;comment:;"`
	CreatedBy   uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy   uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy   uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName Agent 表名
func (Agent) TableName() string {
	return "agent"
}
