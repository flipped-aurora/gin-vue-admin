package model

type SysMenu struct {
	SysBaseMenu
	MenuId      string                 `json:"menuId" gorm:"comment:'菜单ID'"`
	AuthorityId string                 `json:"-" gorm:"comment:'角色ID'"`
	Children    []SysMenu              `json:"children"`
	Parameters  []SysBaseMenuParameter `json:"parameters" gorm:"ForeignKey:MenuId"`
}

func (s SysMenu) TableName() string {
	return "authority_menu"
}
