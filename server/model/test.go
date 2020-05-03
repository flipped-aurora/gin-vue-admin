// 自动生成模板TestTest
package model

import (
	"github.com/jinzhu/gorm"
)

type TestTest struct {
	gorm.Model
	UserName string `json:"user_name" gorm:"column:user_name"`
	Pwd      string `json:"pwd" gorm:"column:pwd"`
	Sex      int    `json:"sex" gorm:"column:sex"`
	Content  string `json:"content" gorm:"column:content"`
}
