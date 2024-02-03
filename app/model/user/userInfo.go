// 自动生成模板UserInfo
package user

import (
	"app/model/request"
	"time"
)
import "gorm.io/gorm"

// 用户信息 结构体  UserInfo
type UserInfo struct {
	ID               uint           `gorm:"primarykey" json:"ID"` // 主键ID
	CreatedAt        time.Time      // 创建时间
	UpdatedAt        time.Time      // 更新时间
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`                                                                                  // 删除时间
	UserWxopenid     string         `json:"userWxopenid" form:"userWxopenid" gorm:"column:user_wxopenid;comment:用户opemid;size:50;"`          //用户openid
	UserNickname     string         `json:"userNickname" form:"userNickname" gorm:"column:user_nickname;comment:用户昵称;size:40;"`              //用户昵称
	UserAvatarUrl    string         `json:"userAvatarUrl" form:"userAvatarUrl" gorm:"column:user_avatar_url;comment:用户头像;size:300;"`         //用户头像
	UserGender       int            `json:"userGender" form:"userGender" gorm:"column:user_gender;comment:用户性别;size:2;"`                     //用户性别
	UserGrade        int            `json:"userGrade" form:"userGrade" gorm:"column:user_grade;comment:用户年级;"`                               //用户年级
	UserProfession   string         `json:"userProfession" form:"userProfession" gorm:"column:user_profession;comment:用户专业;size:50;"`        //用户专业
	UserIntroduction string         `json:"userIntroduction" form:"userIntroduction" gorm:"column:user_introduction;comment:用户简介;size:200;"` //用户简介
	UserLabel        string         `json:"userLabel" form:"userLabel" gorm:"column:user_label;comment:用户标签;size:100;"`                      //用户标签
	UserCity         string         `json:"userCity" form:"userCity" gorm:"column:user_city;comment:用户城市;size:40;"`                          //用户城市
	LoveNumber       int            `json:"loveNumber" form:"loveNumber" gorm:"column:love_number;comment:用户获赞数量;"`                          //用户获赞数量
	SessionKey       string         `json:"sessionKey" gorm:"_"`
	UserModel        int            `json:"userModel" form:"userModel" gorm:"column:user_model;comment:用户身份;"` //用户身份
}

// TableName 用户信息 UserInfo自定义表名 userInfo
func (UserInfo) TableName() string {
	return "userInfo"
}

type UserInfoSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	UserNickname   string     `json:"userNickname" form:"userNickname" `
	UserGender     *int       `json:"userGender" form:"userGender" `
	UserGrade      *int       `json:"userGrade" form:"userGrade" `
	UserProfession string     `json:"userProfession" form:"userProfession" `
	UserCity       string     `json:"userCity" form:"userCity" `
	UserModel      *int       `json:"userModel" form:"userModel" `
	request.PageInfo
}
