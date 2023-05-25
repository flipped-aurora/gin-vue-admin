// 自动生成模板Process
package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Process 结构体
type Process struct {
	global.GVA_MODEL
	StyleID   uint    `json:"styleID" form:"styleID" gorm:"column:style_id;comment:;"`
	Name      string  `json:"name" form:"name" gorm:"column:name;comment:;"`
	Price     float64 `json:"price" form:"price" gorm:"column:price;type:decimal(10,2);comment:;"`
	Percent   float64 `json:"percent" form:"percent" gorm:"column:percent;type:decimal(10,2);comment:;"`
	Sort      uint    `json:"sort" form:"sort" gorm:"column:sort;comment:;"`
	CreatedBy uint    `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint    `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint    `gorm:"column:deleted_by;comment:删除者"`
	Style     Style   `json:"style" form:"style"  gorm:"foreignKey:StyleID;"`
}

// TableName Process 表名
func (Process) TableName() string {
	return "process"
}
