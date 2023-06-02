// 自动生成模板JobApply
package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// JobApply 结构体
type JobApply struct {
	global.GVA_MODEL
	CroppingID     uint           `json:"croppingID" form:"croppingID" gorm:"column:cropping_id;comment:;"`
	UserID         uint           `json:"userID" form:"userID" gorm:"column:user_id;comment:;"`
	TeamID         uint           `json:"teamID" form:"teamID" gorm:"column:team_id;comment:;"`
	ProcessID      uint           `json:"processID" form:"processID" gorm:"column:process_id;comment:;"`
	ProcessName    string         `json:"processName" form:"processName" gorm:"column:process_name;comment:;"`
	Price          float64        `json:"price" form:"price" gorm:"column:price;type:decimal(10,2);comment:;"`
	Size           string         `json:"size" form:"size" gorm:"column:size;comment:;"`
	Quantity       int            `json:"quantity" form:"quantity" gorm:"column:quantity;comment:;"`
	JobType        int            `json:"jobType" form:"jobType" gorm:"column:job_type;comment:;"`
	Status         *int           `json:"status" form:"status" gorm:"column:status;comment:;"`
	CreatedBy      uint           `gorm:"column:created_by;comment:创建者"`
	UpdatedBy      uint           `gorm:"column:updated_by;comment:更新者"`
	DeletedBy      uint           `gorm:"column:deleted_by;comment:删除者"`
	Team           Team           `json:"team" gorm:"foreignKey:TeamID;"`
	Process        Process        `json:"process" gorm:"foreignKey:ProcessID;"`
	User           AppUser        `json:"user" gorm:"foreignKey:UserID;"`
	CroppingRecord CroppingRecord `json:"croppingRecord" gorm:"foreignKey:CroppingID;"`
}

// TableName JobApply 表名
func (JobApply) TableName() string {
	return "job_apply"
}
