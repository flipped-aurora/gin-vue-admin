// 自动生成模板TeamUser
package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// TeamUser 结构体
type TeamUser struct {
	global.GVA_MODEL
	TeamID    uint `json:"teamID" form:"teamID" gorm:"column:team_id;comment:;"`
	UserID    uint `json:"userID" form:"userID" gorm:"column:user_id;comment:;"`
	CreatedBy uint `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint `gorm:"column:deleted_by;comment:删除者"`
}

// TableName TeamUser 表名
func (TeamUser) TableName() string {
	return "team_user"
}
