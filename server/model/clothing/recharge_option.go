// 自动生成模板RechargeOption
package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// RechargeOption 结构体
type RechargeOption struct {
	global.GVA_MODEL
	Day        int     `json:"day" form:"day" gorm:"column:day;comment:;"`
	ClerkCount int     `json:"clerkCount" form:"clerkCount" gorm:"column:clerk_count;comment:;"`
	Price      float64 `json:"price" form:"price" gorm:"column:price;comment:;type:decimal(10,2);"`
	Amount     float64 `json:"amount" form:"amount" gorm:"column:amount;comment:;type:decimal(10,2);"`
	CreatedBy  uint    `gorm:"column:created_by;comment:创建者"`
	UpdatedBy  uint    `gorm:"column:updated_by;comment:更新者"`
	DeletedBy  uint    `gorm:"column:deleted_by;comment:删除者"`
}

// TableName RechargeOption 表名
func (RechargeOption) TableName() string {
	return "recharge_option"
}
