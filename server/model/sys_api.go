package model

import (
	"github.com/jinzhu/gorm"
)

type SysApi struct {
	gorm.Model
	Path        string `json:"path"`
	Description string `json:"description"`
	ApiGroup    string `json:"apiGroup"`
	Method      string `json:"method" gorm:"default:'POST'"`
}
