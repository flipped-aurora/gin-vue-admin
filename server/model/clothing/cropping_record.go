// 自动生成模板CroppingRecord
package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CroppingRecord 结构体
type CroppingRecord struct {
	global.GVA_MODEL
	CompanyID uint    `json:"companyID" form:"companyID" gorm:"column:company_id;comment:;"`
	UserID    uint    `json:"userID" form:"userID" gorm:"column:user_id;comment:;"`
	StyleID   uint    `json:"styleID" form:"styleID" gorm:"column:style_id;comment:;"`
	Color     string  `json:"color" form:"color" gorm:"column:color;comment:;"`
	Length    float64 `json:"length" form:"length" gorm:"column:length;type:decimal(10,2);comment:;"`
	Size      string  `json:"size" form:"size" gorm:"column:size;comment:;"`
	Quantity  uint    `json:"quantity" form:"quantity" gorm:"column:quantity;comment:;"`
	Step      uint    `json:"step" form:"step" gorm:"column:step;comment:;"`
	CreatedBy uint    `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint    `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint    `gorm:"column:deleted_by;comment:删除者"`
}

// TableName CroppingRecord 表名
func (CroppingRecord) TableName() string {
	return "cropping_record"
}
