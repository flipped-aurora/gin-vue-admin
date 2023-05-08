// 自动生成模板CompanyApply
package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CompanyApply 结构体
type CompanyApply struct {
	global.GVA_MODEL
	CompanyID uint    `json:"companyID" form:"companyID" gorm:"column:company_id;comment:;"`
	UserID    uint    `json:"userID" form:"userID" gorm:"column:user_id;comment:;"`
	RoleID    uint    `json:"roleID" form:"roleID" gorm:"column:role_id;comment:;"`
	Remark    string  `json:"remark" form:"remark" gorm:"column:remark;comment:;"`
	Status    *int    `json:"status" form:"status" gorm:"column:status;comment:;"`
	CreatedBy uint    `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint    `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint    `gorm:"column:deleted_by;comment:删除者"`
	Company   Company `json:"company" form:"company" gorm:"foreignKey:CompanyID"`
}

// TableName CompanyApply 表名
func (CompanyApply) TableName() string {
	return "company_apply"
}
