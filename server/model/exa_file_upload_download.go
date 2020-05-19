package model

import (
	"github.com/jinzhu/gorm"
)

type ExaFileUploadAndDownload struct {
	gorm.Model
	Name string `json:"name"`
	Url  string `json:"url"`
	Tag  string `json:"tag"`
	Key  string `json:"key"`
}
