// 自动生成模板Category
package model

import (
    "github.com/jinzhu/gorm"
)

type Category struct {
      gorm.Model 
      CName  string `json:"c_name" gorm:"column:c_name"`
      KeyWord  string `json:"key_word" gorm:"column:key_word"`
      Desc  string `json:"desc" gorm:"column:desc"`
      IcoImage  string `json:"ico_image" gorm:"column:ico_image"`
      BgImage  string `json:"bg_image" gorm:"column:bg_image"`
      Sort  int `json:"sort" gorm:"column:sort"`
      Href  string `json:"href" gorm:"column:href"`
      Status  int `json:"status" gorm:"column:status"`
}