// 自动生成模板ComInfo
package competition

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 比赛信息 结构体  ComInfo
type ComInfo struct {
      global.GVA_MODEL
      ComTitle  string `json:"comTitle" form:"comTitle" gorm:"column:com_title;comment:比赛标题;size:100;"binding:"required"`  //比赛标题 
      ComIntruduction  string `json:"comIntruduction" form:"comIntruduction" gorm:"column:com_intruduction;comment:比赛简介;size:300;"`  //比赛简介 
      ComPicture  string `json:"comPicture" form:"comPicture" gorm:"column:com_picture;comment:比赛图片;size:100;"`  //比赛图片 
}


// TableName 比赛信息 ComInfo自定义表名 comInfo
func (ComInfo) TableName() string {
  return "comInfo"
}

