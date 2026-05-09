package initialize

import (
	"context"

	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

// Menu 注册五笔打字插件菜单
// menus[0] 为父菜单 placed under "插件系统" (id=9)
// menus[1:] 为子菜单 自动挂到父菜单下
func Menu(ctx context.Context) {
	entities := []model.SysBaseMenu{
		{
			ParentId:  9,
			Path:      "wubi",
			Name:      "wubi",
			Hidden:    false,
			Component: "view/routerHolder.vue",
			Sort:      10,
			Meta:      model.Meta{Title: "五笔打字", Icon: "edit"},
		},
		{
			Path:      "wubiPractice",
			Name:      "wubiPractice",
			Hidden:    false,
			Component: "plugin/wubi/view/practice.vue",
			Sort:      1,
			Meta:      model.Meta{Title: "五笔练习", Icon: "edit-pen"},
		},
		{
			Path:      "wubiLeaderboard",
			Name:      "wubiLeaderboard",
			Hidden:    false,
			Component: "plugin/wubi/view/leaderboard.vue",
			Sort:      2,
			Meta:      model.Meta{Title: "五笔排行", Icon: "trophy"},
		},
	}
	utils.RegisterMenus(entities...)
}
