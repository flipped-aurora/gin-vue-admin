package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/system"

type SysMenusResponse struct {
	Menus []system.SysMenu `json:"menus" comment:"系统菜单详情列表"`
}

type SysBaseMenusResponse struct {
	Menus []system.SysBaseMenu `json:"menus" comment:"系统菜单列表"`
}

type SysBaseMenuResponse struct {
	Menu system.SysBaseMenu `json:"menu" comment:"系统菜单列表"`
}
