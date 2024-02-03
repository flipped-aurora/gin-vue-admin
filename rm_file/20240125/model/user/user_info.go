// 自动生成模板UserInfo
package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 用户信息 结构体  UserInfo
type UserInfo struct {
      global.GVA_MODEL
      UserWxopenid  string `json:"userWxopenid" form:"userWxopenid" gorm:"column:user_wxopenid;comment:用户opemid;size:50;"`  //用户openid 
      UserNickname  string `json:"userNickname" form:"userNickname" gorm:"column:user_nickname;comment:用户昵称;size:40;"`  //用户昵称 
      UserAvatarUrl  string `json:"userAvatarUrl" form:"userAvatarUrl" gorm:"column:user_avatar_url;comment:用户头像;size:50;"`  //用户头像 
      UserGender  *int `json:"userGender" form:"userGender" gorm:"column:user_gender;comment:用户性别;size:2;"`  //用户性别 
      UserGrade  *int `json:"userGrade" form:"userGrade" gorm:"column:user_grade;comment:用户年级;"`  //用户年级 
      UserProfession  string `json:"userProfession" form:"userProfession" gorm:"column:user_profession;comment:用户专业;size:50;"`  //用户专业 
      UserIntroduction  string `json:"userIntroduction" form:"userIntroduction" gorm:"column:user_introduction;comment:用户简介;size:200;"`  //用户简介 
      UserLabel  string `json:"userLabel" form:"userLabel" gorm:"column:user_label;comment:用户标签;size:100;"`  //用户标签 
      UserCity  string `json:"userCity" form:"userCity" gorm:"column:user_city;comment:用户城市;size:40;"`  //用户城市 
      LoveNumber  *int `json:"loveNumber" form:"loveNumber" gorm:"column:love_number;comment:用户获赞数量;"`  //用户获赞数量 
}


// TableName 用户信息 UserInfo自定义表名 userInfo
func (UserInfo) TableName() string {
  return "userInfo"
}
