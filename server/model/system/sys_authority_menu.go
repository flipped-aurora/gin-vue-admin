package system

type SysMenu struct {
	SysBaseMenu
	MenuId      string                 `json:"menuId" gorm:"comment:菜单ID" example:"menu_id"`
	AuthorityId string                 `json:"-" gorm:"comment:角色ID" example:"role_id"`
	Children    []SysMenu              `json:"children" gorm:"-"`
	Parameters  []SysBaseMenuParameter `json:"parameters" gorm:"foreignKey:SysBaseMenuID;references:MenuId"`
}

func (s SysMenu) TableName() string {
	return "authority_menu"
}
