// 自动生成模板Computation
package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Computation 结构体
type Computation struct {
	global.GVA_MODEL
	ClothWidth     float64 `json:"clothWidth" form:"clothWidth" gorm:"column:cloth_width;type:decimal(10,2);comment:;"`
	BarLength      float64 `json:"barLength" form:"barLength" gorm:"column:bar_length;type:decimal(10,2);comment:;"`
	BarWidth       float64 `json:"barWidth" form:"barWidth" gorm:"column:bar_width;type:decimal(10,2);comment:;"`
	CountNum       int     `json:"countNum" form:"countNum" gorm:"column:count_num;comment:;"`
	BarLength45    float64 `json:"barLength45" form:"barLength45" gorm:"column:bar_length45;type:decimal(10,2);comment:;"`
	BarLength90    float64 `json:"barLength90" form:"barLength90" gorm:"column:bar_length90;type:decimal(10,2);comment:;"`
	BarLength180   float64 `json:"barLength180" form:"barLength180" gorm:"column:bar_length180;type:decimal(10,2);comment:;"`
	ClothLength45  float64 `json:"clothLength45" form:"clothLength45" gorm:"column:cloth_length45;type:decimal(10,2);comment:;"`
	ClothLength90  float64 `json:"clothLength90" form:"clothLength90" gorm:"column:cloth_length90;type:decimal(10,2);comment:;"`
	ClothLength180 float64 `json:"clothLength180" form:"clothLength180" gorm:"column:cloth_length180;type:decimal(10,2);comment:;"`
	UserID         uint    `json:"userID" form:"userID" gorm:"column:user_id;comment:;"`
	CreatedBy      uint    `gorm:"column:created_by;comment:创建者"`
	UpdatedBy      uint    `gorm:"column:updated_by;comment:更新者"`
	DeletedBy      uint    `gorm:"column:deleted_by;comment:删除者"`
}

// TableName Computation 表名
func (Computation) TableName() string {
	return "computation"
}
