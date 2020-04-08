package model

type SysMenu struct {
	SysBaseMenu
	MenuId      string    `json:"menuId"`
	AuthorityId string    `json:"-"`
	Children    []SysMenu `json:"children"`
}