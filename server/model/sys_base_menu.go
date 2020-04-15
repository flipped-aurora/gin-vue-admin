package model

import (
	"github.com/jinzhu/gorm"
)

type SysBaseMenu struct {
	gorm.Model
	MenuLevel     uint   `json:"-"`
	ParentId      string `json:"parentId"`
	Path          string `json:"path"`
	Name          string `json:"name"`
	Hidden        bool   `json:"hidden"`
	Component     string `json:"component"`
	Sort          int    `json:"sort"`
	Meta          `json:"meta"`
	SysAuthoritys []SysAuthority `json:"authoritys" gorm:"many2many:sys_authority_menus;"`
	Children      []SysBaseMenu  `json:"children"`
}

type Meta struct {
	KeepAlive   bool   `json:"keepAlive"`
	DefaultMenu bool   `json:"defaultMenu"`
	Title       string `json:"title"`
	Icon        string `json:"icon"`
}
