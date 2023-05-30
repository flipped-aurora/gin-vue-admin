// 自动生成模板Style
package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Style 结构体
type Style struct {
	global.GVA_MODEL
	CompanyID   uint    `json:"companyID" form:"companyID" gorm:"column:company_id;comment:;"`
	Name        string  `json:"name" form:"name" gorm:"column:name;comment:;"`
	Price       float64 `json:"price" form:"price" gorm:"column:price;type:decimal(10,2);comment:;"`
	CroppingNum int     `json:"croppingNum" form:"croppingNum" gorm:"column:cropping_num;default:1;comment:;"`
	CreatedBy   uint    `gorm:"column:created_by;comment:创建者"`
	UpdatedBy   uint    `gorm:"column:updated_by;comment:更新者"`
	DeletedBy   uint    `gorm:"column:deleted_by;comment:删除者"`
}

// TableName Style 表名
func (Style) TableName() string {
	return "style"
}
