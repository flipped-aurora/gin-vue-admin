package initialize

import (
	"context"

	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Api(ctx context.Context) {
	entities := []model.SysApi{
		// 通知管理相关API
		{
			Path:        "/notice/notification/createNotification",
			Description: "创建通知",
			ApiGroup:    "通知管理",
			Method:      "POST",
		},
		{
			Path:        "/notice/notification/updateNotification",
			Description: "更新通知",
			ApiGroup:    "通知管理",
			Method:      "PUT",
		},
		{
			Path:        "/notice/notification/deleteNotification",
			Description: "删除通知",
			ApiGroup:    "通知管理",
			Method:      "DELETE",
		},
		{
			Path:        "/notice/notification/publishNotification/:id",
			Description: "发布通知",
			ApiGroup:    "通知管理",
			Method:      "POST",
		},
		{
			Path:        "/notice/notification/sendNotification",
			Description: "发送通知",
			ApiGroup:    "通知管理",
			Method:      "POST",
		},
		{
			Path:        "/notice/notification/getNotificationList",
			Description: "获取通知列表",
			ApiGroup:    "通知管理",
			Method:      "POST",
		},
		{
			Path:        "/notice/notification/getNotificationById/:id",
			Description: "根据ID获取通知",
			ApiGroup:    "通知管理",
			Method:      "GET",
		},
		{
			Path:        "/notice/notification/getUserNotifications",
			Description: "获取用户通知列表",
			ApiGroup:    "通知管理",
			Method:      "POST",
		},
		{
			Path:        "/notice/notification/markAsRead",
			Description: "标记通知已读",
			ApiGroup:    "通知管理",
			Method:      "POST",
		},
		{
			Path:        "/notice/notification/getNotificationStats",
			Description: "获取通知统计",
			ApiGroup:    "通知管理",
			Method:      "GET",
		},
		// 在线用户管理相关API
		{
			Path:        "/notice/onlineUser/removeOnlineUser/:userId",
			Description: "移除在线用户",
			ApiGroup:    "在线用户管理",
			Method:      "DELETE",
		},
		{
			Path:        "/notice/onlineUser/cleanOfflineUsers",
			Description: "清理离线用户",
			ApiGroup:    "在线用户管理",
			Method:      "POST",
		},
		{
			Path:        "/notice/onlineUser/getOnlineUsers",
			Description: "获取在线用户列表",
			ApiGroup:    "在线用户管理",
			Method:      "GET",
		},
		{
			Path:        "/notice/onlineUser/getOnlineUsers",
			Description: "获取在线用户列表（分页）",
			ApiGroup:    "在线用户管理",
			Method:      "POST",
		},
		{
			Path:        "/notice/onlineUser/getOnlineUsersByRole/:roleId",
			Description: "根据角色获取在线用户",
			ApiGroup:    "在线用户管理",
			Method:      "GET",
		},
		{
			Path:        "/notice/onlineUser/checkUserOnline/:userId",
			Description: "检查用户是否在线",
			ApiGroup:    "在线用户管理",
			Method:      "GET",
		},
		{
			Path:        "/notice/onlineUser/getOnlineUserCount",
			Description: "获取在线用户数量",
			ApiGroup:    "在线用户管理",
			Method:      "GET",
		},
		{
			Path:        "/notice/onlineUser/getOnlineUserStats",
			Description: "获取在线用户统计数据",
			ApiGroup:    "在线用户管理",
			Method:      "GET",
		},
		// Socket.IO相关API
		{
			Path:        "/socket.io/*any",
			Description: "Socket.IO连接",
			ApiGroup:    "实时通信",
			Method:      "GET",
		},
		{
			Path:        "/socket.io/*any",
			Description: "Socket.IO连接",
			ApiGroup:    "实时通信",
			Method:      "POST",
		},
	}
	utils.RegisterApis(entities...)
}
