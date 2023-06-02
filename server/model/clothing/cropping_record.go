// 自动生成模板CroppingRecord
package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

// CroppingRecord 结构体
type CroppingRecord struct {
	global.GVA_MODEL
	CompanyID   uint        `json:"companyID" form:"companyID" gorm:"column:company_id;comment:;"`
	CroppingNum int         `json:"croppingNum" form:"croppingNum" gorm:"column:cropping_num;default:1;comment:;"`
	UserID      uint        `json:"userID" form:"userID" gorm:"column:user_id;comment:;"`
	StyleID     uint        `json:"styleID" form:"styleID" gorm:"column:style_id;comment:;"`
	Color       string      `json:"color" form:"color" gorm:"column:color;comment:;"`
	Length      float64     `json:"length" form:"length" gorm:"column:length;type:decimal(10,2);comment:;"`
	Quantity    uint        `json:"quantity" form:"quantity" gorm:"column:quantity;comment:;"`
	Remainder   float64     `json:"remainder" form:"remainder" gorm:"column:remainder;type:decimal(10,2);comment:;"`
	Usage       float64     `json:"usage" form:"usage" gorm:"column:usage;type:decimal(10,2);comment:;"`
	Step        uint        `json:"step" form:"step" gorm:"column:step;comment:;"`
	CreatedBy   uint        `gorm:"column:created_by;comment:创建者"`
	UpdatedBy   uint        `gorm:"column:updated_by;comment:更新者"`
	DeletedBy   uint        `gorm:"column:deleted_by;comment:删除者"`
	Style       Style       `json:"style" form:"style"  gorm:"foreignKey:StyleID;"`
	SizeList    []SizeList  `gorm:"foreignKey:CroppingRecordID"`
	Inventory   []Inventory `gorm:"foreignKey:CroppingRecordID"`
}

// TableName CroppingRecord 表名
func (CroppingRecord) TableName() string {
	return "cropping_record"
}

func (receiver *CroppingRecord) BeforeCreate(tx *gorm.DB) (err error) {
	var style Style
	tx.First(&style, receiver.StyleID)
	receiver.CroppingNum = style.CroppingNum
	tx.Model(&style).Update("cropping_num", gorm.Expr("cropping_num + ?", 1))
	return
}
