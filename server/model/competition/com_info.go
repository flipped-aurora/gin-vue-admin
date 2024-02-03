// 自动生成模板ComInfo
package competition

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
	
)

// 比赛信息 结构体  ComInfo
type ComInfo struct {
      global.GVA_MODEL
      ComTitle  string `json:"comTitle" form:"comTitle" gorm:"column:com_title;comment:比赛标题;size:100;"binding:"required"`  //比赛标题 
      ComIntroduction  string `json:"comIntroduction" form:"comIntroduction" gorm:"column:com_introduction;comment:比赛简介;size:300;"`  //比赛简介 
      ComPicture  string `json:"comPicture" form:"comPicture" gorm:"column:com_picture;comment:比赛图片;size:100;"`  //比赛图片 
      ComModel  *int `json:"comModel" form:"comModel" gorm:"column:com_model;comment:比赛等级;"`  //比赛等级 
      ComStart  *time.Time `json:"comStart" form:"comStart" gorm:"column:com_start;comment:比赛开始时间;"`  //比赛开始时间 
      ComEnd  *time.Time `json:"comEnd" form:"comEnd" gorm:"column:com_end;comment:比赛结束时间;"`  //比赛结束时间 
      ComNumber  *int `json:"comNumber" form:"comNumber" gorm:"column:com_number;comment:;"`  //比赛参加人数 
      ComType  *int `json:"comType" form:"comType" gorm:"column:com_type;comment:比赛类型;"`  //比赛类型 
}


// TableName 比赛信息 ComInfo自定义表名 comInfo
func (ComInfo) TableName() string {
  return "comInfo"
}

