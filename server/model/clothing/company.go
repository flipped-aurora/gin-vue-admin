// 自动生成模板Company
package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Company 结构体
type Company struct {
	global.GVA_MODEL
	UserID    uint   `json:"userID" form:"userID" gorm:"column:user_id;comment:;"`
	Name      string `json:"name" form:"name" gorm:"column:name;comment:;"`
	Status    *int   `json:"status" form:"status" gorm:"column:status;comment:;"`
	CreatedBy uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName Company 表名
func (Company) TableName() string {
	return "company"
}
