package request

import "gin-vue-admin/model"

// Add menu authority info structure
type AddMenuAuthorityInfo struct {
	Menus       []model.SysBaseMenu
	AuthorityId string
}
