// 自动生成模板TeamApply
package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// TeamApply 结构体
type TeamApply struct {
	global.GVA_MODEL
	TeamID    uint    `json:"teamID" form:"teamID" gorm:"column:team_id;comment:;"`
	UserID    uint    `json:"userID" form:"userID" gorm:"column:user_id;comment:;"`
	RoleID    uint    `json:"roleID" form:"roleID" gorm:"column:role_id;comment:;"`
	Status    *int    `json:"status" form:"status" gorm:"column:status;comment:;"`
	CreatedBy uint    `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint    `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint    `gorm:"column:deleted_by;comment:删除者"`
	Team      Team    `json:"team" gorm:"foreignKey:TeamID;"`
	User      AppUser `json:"user" gorm:"foreignKey:UserID;"`
}

// TableName TeamApply 表名
func (TeamApply) TableName() string {
	return "team_apply"
}
