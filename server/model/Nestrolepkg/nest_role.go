// 自动生成模板NestRole
package Nestrolepkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// NestRole 结构体
type NestRole struct {
	global.GVA_MODEL
	Roleid       uint                `json:"roleid" form:"roleid" gorm:"column:roleid;not null;comment:;"`
	Nestid       string              `json:"nestid" form:"nestid" gorm:"column:nestid;comment:;"`
	Authority    system.SysAuthority `json:"authority" gorm:"foreignKey:Roleid;references:AuthorityId;comment:用户角色"`
	CreatedBy    uint                `gorm:"column:created_by;comment:创建者"`
	UpdatedBy    uint                `gorm:"column:updated_by;comment:更新者"`
	DeletedBy    uint                `gorm:"column:deleted_by;comment:删除者"`
	RoleidSearch string              `json:"roleidsearch" form:"roleidsearch"`
}

// TableName NestRole 表名
func (NestRole) TableName() string {
	return "nest_role"
}
