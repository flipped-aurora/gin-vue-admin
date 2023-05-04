// 自动生成模板UserRole
package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// UserRole 结构体
type UserRole struct {
	global.GVA_MODEL
	UserID    uint    `json:"userID" form:"userID" gorm:"column:user_id;comment:;"`
	RoleID    uint    `json:"roleID" form:"roleID" gorm:"column:role_id;comment:;"`
	CompanyID uint    `json:"companyID" form:"companyID" gorm:"column:company_id;comment:;"`
	CreatedBy uint    `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint    `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint    `gorm:"column:deleted_by;comment:删除者"`
	Role      AppRole `json:"role" form:"role" gorm:"foreignKey:RoleID"`
	Company   Company `json:"company" form:"company" gorm:"foreignKey:CompanyID"`
}

// TableName UserRole 表名
func (UserRole) TableName() string {
	return "user_role"
}
