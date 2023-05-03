// 自动生成模板Job
package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Job 结构体
type Job struct {
	global.GVA_MODEL
	CroppingID   uint    `json:"croppingID" form:"croppingID" gorm:"column:cropping_id;comment:;"`
	UserID       uint    `json:"userID" form:"userID" gorm:"column:user_id;comment:;"`
	ProcessName  string  `json:"processName" form:"processName" gorm:"column:process_name;comment:;"`
	Price        float64 `json:"price" form:"price" gorm:"column:price;type:decimal(10,2);comment:;"`
	Quantity     uint    `json:"quantity" form:"quantity" gorm:"column:quantity;comment:;"`
	Income       float64 `json:"income" form:"income" gorm:"column:income;type:decimal(10,2);comment:;"`
	RealQuantity uint    `json:"realQuantity" form:"realQuantity" gorm:"column:real_quantity;comment:;"`
	RealIncome   float64 `json:"realIncome" form:"realIncome" gorm:"column:real_income;type:decimal(10,2);comment:;"`
	Step         uint    `json:"step" form:"step" gorm:"column:step;comment:;"`
	JobType      uint    `json:"jobType" form:"jobType" gorm:"column:job_type;comment:;"`
	CreatedBy    uint    `gorm:"column:created_by;comment:创建者"`
	UpdatedBy    uint    `gorm:"column:updated_by;comment:更新者"`
	DeletedBy    uint    `gorm:"column:deleted_by;comment:删除者"`
}

// TableName Job 表名
func (Job) TableName() string {
	return "job"
}
