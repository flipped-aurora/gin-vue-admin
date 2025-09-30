package initialize

import (
	"context"

	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Menu(ctx context.Context) {
	entities := []model.SysBaseMenu{
		{
			MenuLevel: 0,
			ParentId:  0,
			Path:      "notice",
			Name:      "notice",
			Hidden:    false,
			Component: "view/routerHolder.vue",
			Sort:      5,
			Meta: model.Meta{
				Title: "通知中心",
				Icon:  "notification",
			},
		},
		{
			MenuLevel: 0,
			ParentId:  0, // 这里需要在运行时动态设置为父菜单ID
			Path:      "notificationManage",
			Name:      "notificationManage",
			Hidden:    false,
			Component: "view/notice/notification/index.vue",
			Sort:      1,
			Meta: model.Meta{
				Title: "通知管理",
				Icon:  "message",
			},
		},
		{
			MenuLevel: 0,
			ParentId:  0, // 这里需要在运行时动态设置为父菜单ID
			Path:      "onlineUserManage",
			Name:      "onlineUserManage",
			Hidden:    false,
			Component: "view/notice/onlineUser/index.vue",
			Sort:      2,
			Meta: model.Meta{
				Title: "在线用户",
				Icon:  "user",
			},
		},
		{
			MenuLevel: 0,
			ParentId:  0, // 这里需要在运行时动态设置为父菜单ID
			Path:      "userCenter",
			Name:      "userCenter",
			Hidden:    false,
			Component: "view/notice/userCenter/index.vue",
			Sort:      3,
			Meta: model.Meta{
				Title: "我的通知",
				Icon:  "bell",
			},
		},
	}
	utils.RegisterMenus(entities...)
}
