package model

import (
	"gorm.io/gorm"
	"time"
)

type ExaFileUploadAndDownload struct {
	gorm.Model
	Name string `json:"name" gorm:"comment:'文件名'"`
	Url  string `json:"url" gorm:"comment:'文件地址'"`
	Tag  string `json:"tag" gorm:"comment:'文件标签'"`
	Key  string `json:"key" gorm:"comment:'编号'"`
}

func ExaFileUploadAndDownloadData() []ExaFileUploadAndDownload {
	return []ExaFileUploadAndDownload{
		{gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "10.png", "http://qmplusimg.henrongyi.top/158787308910.png", "png", "158787308910.png"},
		{gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "logo.png", "http://qmplusimg.henrongyi.top/1587973709logo.png", "png", "1587973709logo.png"},
	}
}
