// 自动生成模板AppUser
package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// AppUser 结构体
type AppUser struct {
      global.GVA_MODEL
      UserName  string `json:"userName" form:"userName" gorm:"column:user_name;comment:;"`
      Password  string `json:"password" form:"password" gorm:"column:password;comment:;"`
      RoleID  *int `json:"roleID" form:"roleID" gorm:"column:role_id;comment:;"`
      Nickname  string `json:"nickname" form:"nickname" gorm:"column:nickname;comment:;"`
      CompanyID  *int `json:"companyID" form:"companyID" gorm:"column:company_id;comment:;"`
      Wages  *float64 `json:"wages" form:"wages" gorm:"column:wages;comment:;"`
      WorkType  *int `json:"workType" form:"workType" gorm:"column:work_type;comment:;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:;"`
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName AppUser 表名
func (AppUser) TableName() string {
  return "app_user"
}

