package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	noticeApi "github.com/flipped-aurora/gin-vue-admin/server/plugin/notice/api"
	"github.com/gin-gonic/gin"
)

type NotificationRouter struct{}

// InitNotificationRouter 初始化通知路由
func (r *NotificationRouter) InitNotificationRouter(Router *gin.RouterGroup) {
	notificationRouter := Router.Group("notification")
	notificationRouterWithoutRecord := Router.Group("notification")

	// 需要记录操作日志的路由
	notificationRouter.Use(middleware.OperationRecord())
	{
		// 管理员接口
		notificationRouter.POST("createNotification", noticeApi.ApiGroupApp.NotificationApi.CreateNotification)       // 创建通知
		notificationRouter.PUT("updateNotification", noticeApi.ApiGroupApp.NotificationApi.UpdateNotification)        // 更新通知
		notificationRouter.DELETE("deleteNotification", noticeApi.ApiGroupApp.NotificationApi.DeleteNotification)     // 删除通知
		notificationRouter.POST("publishNotification/:id", noticeApi.ApiGroupApp.NotificationApi.PublishNotification) // 发布通知
		notificationRouter.POST("sendNotification", noticeApi.ApiGroupApp.NotificationApi.SendNotification)           // 发送通知
	}

	// 不需要记录操作日志的路由
	notificationRouterWithoutRecord.Use(middleware.JWTAuth())
	{
		// 查询接口
		notificationRouterWithoutRecord.POST("getNotificationList", noticeApi.ApiGroupApp.NotificationApi.GetNotificationList)    // 获取通知列表
		notificationRouterWithoutRecord.GET("getNotificationById/:id", noticeApi.ApiGroupApp.NotificationApi.GetNotificationById) // 根据ID获取通知

		// 用户接口
		notificationRouterWithoutRecord.POST("getUserNotifications", noticeApi.ApiGroupApp.NotificationApi.GetUserNotifications)                       // 获取用户通知列表
		notificationRouterWithoutRecord.POST("markAsRead", noticeApi.ApiGroupApp.NotificationApi.MarkNotificationAsRead)                               // 标记已读
		notificationRouterWithoutRecord.GET("getNotificationStats", noticeApi.ApiGroupApp.NotificationApi.GetNotificationStats)                        // 获取通知统计
		notificationRouterWithoutRecord.DELETE("deleteUserNotification/:notificationId", noticeApi.ApiGroupApp.NotificationApi.DeleteUserNotification) // 删除用户通知
	}

}
