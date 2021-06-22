package request

import (
	"gin-vue-admin/model"
)

// Add menu authority info structure
type AddMenuAuthorityInfo struct {
	Menus       []model.SysBaseMenu
	AuthorityId string // 角色ID
}

func DefaultMenu() []model.SysBaseMenu {
	return []model.SysBaseMenu{{
		ParentId:  "0",
		Path:      "dashboard",
		Name:      "dashboard",
		Component: "view/dashboard/index.vue",
		Sort:      1,
		Meta: model.Meta{
			Title: "仪表盘",
			Icon:  "setting",
		},
	}}
}
