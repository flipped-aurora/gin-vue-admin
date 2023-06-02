// 自动生成模板Inventory
package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Inventory 结构体
type Inventory struct {
	global.GVA_MODEL
	CroppingRecordID uint   `json:"croppingRecordID" form:"croppingRecordID" gorm:"column:cropping_record_id;comment:;"`
	ProcessID        uint   `json:"processID" form:"processID" gorm:"column:process_id;comment:;"`
	Size             string `json:"size" form:"size" gorm:"column:size;comment:;"`
	Total            uint   `json:"total" form:"total" gorm:"column:total;comment:;"`
	Margin           uint   `json:"margin" form:"margin" gorm:"column:margin;comment:;"`
}

// TableName Inventory 表名
func (Inventory) TableName() string {
	return "inventory"
}
