// 自动生成模板Team
package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Team 结构体
type Team struct {
      global.GVA_MODEL
      CompanyID  *int `json:"companyID" form:"companyID" gorm:"column:company_id;comment:;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;"`
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName Team 表名
func (Team) TableName() string {
  return "team"
}

