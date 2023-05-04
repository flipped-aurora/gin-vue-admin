// 自动生成模板Banner
package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Banner 结构体
type Banner struct {
      global.GVA_MODEL
      Url  string `json:"url" form:"url" gorm:"column:url;comment:;"`
      Title  string `json:"title" form:"title" gorm:"column:title;comment:;"`
      Content  string `json:"content" form:"content" gorm:"column:content;comment:;"`
      Img  string `json:"img" form:"img" gorm:"column:img;comment:;"`
      Sort  *int `json:"sort" form:"sort" gorm:"column:sort;comment:;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:;"`
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName Banner 表名
func (Banner) TableName() string {
  return "banner"
}

