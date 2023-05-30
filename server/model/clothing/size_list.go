// 自动生成模板SizeList
package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SizeList 结构体
type SizeList struct {
	global.GVA_MODEL
	CroppingRecordID uint   `json:"croppingRecordID" form:"croppingRecordID" gorm:"column:cropping_record_id;comment:;"`
	Size             string `json:"size" form:"size" gorm:"column:size;comment:;"`
	Quantity         uint   `json:"quantity" form:"quantity" gorm:"column:quantity;comment:;"`
	CreatedBy        uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy        uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy        uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName SizeList 表名
func (SizeList) TableName() string {
	return "size_list"
}
