// 自动生成模板Cloth
package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Cloth 结构体
type Cloth struct {
	global.GVA_MODEL
	CompanyID uint    `json:"companyID" form:"companyID" gorm:"column:company_id;comment:;"`
	StyleID   uint    `json:"styleID" form:"styleID" gorm:"column:style_id;comment:;"`
	Color     string  `json:"color" form:"color" gorm:"column:color;comment:;"`
	Length    float64 `json:"length" form:"length" gorm:"column:length;type:decimal(10,2);comment:;"`
	CreatedBy uint    `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint    `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint    `gorm:"column:deleted_by;comment:删除者"`
}

// TableName Cloth 表名
func (Cloth) TableName() string {
	return "cloth"
}
