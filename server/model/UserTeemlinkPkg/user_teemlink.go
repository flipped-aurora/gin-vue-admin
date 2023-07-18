// 自动生成模板UserTeemlink
package UserTeemlinkPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// UserTeemlink 结构体
type UserTeemlink struct {
      global.GVA_MODEL
      UserID  *int `json:"userID" form:"userID" gorm:"column:user_id;comment:;"`
      Username  string `json:"username" form:"username" gorm:"column:username;comment:;"`
      TlUserID  string `json:"tlUserID" form:"tlUserID" gorm:"column:tl_user_id;comment:;"`
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName UserTeemlink 表名
func (UserTeemlink) TableName() string {
  return "user_teemlink"
}

