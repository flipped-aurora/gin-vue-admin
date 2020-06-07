package model

import (
	"time"
)

type SysAuthority struct {
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time     `sql:"index"`
	AuthorityId     string         `json:"authorityId" gorm:"not null;unique;primary_key" gorm:"comment:'角色ID'"`
	AuthorityName   string         `json:"authorityName" gorm:"comment:'角色名'"`
	ParentId        string         `json:"parentId" gorm:"comment:'父角色ID'"`
	DataAuthorityId []SysAuthority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id;association_jointable_foreignkey:data_authority_id"`
	Children        []SysAuthority `json:"children"`
	SysBaseMenus    []SysBaseMenu  `json:"menus" gorm:"many2many:sys_authority_menus;"`
}
