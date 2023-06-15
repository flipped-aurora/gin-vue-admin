// 自动生成模板SysUsers
package test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// SysUsers 结构体
type SysUsers struct {
      global.GVA_MODEL
      Uuid  * `json:"uuid" form:"uuid" gorm:"column:uuid;comment:;"`
      Username  * `json:"username" form:"username" gorm:"column:username;comment:;"`
      Password  * `json:"password" form:"password" gorm:"column:password;comment:;"`
      NickName  * `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:;"`
      SideMode  * `json:"sideMode" form:"sideMode" gorm:"column:side_mode;comment:;"`
      HeaderImg  * `json:"headerImg" form:"headerImg" gorm:"column:header_img;comment:;"`
      BaseColor  * `json:"baseColor" form:"baseColor" gorm:"column:base_color;comment:;"`
      ActiveColor  * `json:"activeColor" form:"activeColor" gorm:"column:active_color;comment:;"`
      AuthorityId  * `json:"authorityId" form:"authorityId" gorm:"column:authority_id;comment:;"`
      Phone  * `json:"phone" form:"phone" gorm:"column:phone;comment:;"`
      Email  * `json:"email" form:"email" gorm:"column:email;comment:;"`
      Enable  * `json:"enable" form:"enable" gorm:"column:enable;comment:;"`
}


// TableName SysUsers 表名
func (SysUsers) TableName() string {
  return "sys_users"
}

