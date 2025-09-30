package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

// NotificationSearch 通知搜索请求
type NotificationSearch struct {
	request.PageInfo
	Title      string     `json:"title" form:"title"`
	Type       string     `json:"type" form:"type"`
	Priority   string     `json:"priority" form:"priority"`
	Status     string     `json:"status" form:"status"`
	SenderId   uint       `json:"senderId" form:"senderId"`
	TargetType string     `json:"targetType" form:"targetType"`
	StartTime  *time.Time `json:"startTime" form:"startTime"`
	EndTime    *time.Time `json:"endTime" form:"endTime"`
}

// CreateNotificationRequest 创建通知请求
type CreateNotificationRequest struct {
	Title       string     `json:"title" binding:"required" form:"title"`
	Content     string     `json:"content" binding:"required" form:"content"`
	Type        string     `json:"type" binding:"required" form:"type"`
	Priority    string     `json:"priority" form:"priority"`
	TargetType  string     `json:"targetType" binding:"required" form:"targetType"`
	TargetIds   []uint     `json:"targetIds" form:"targetIds"`
	PublishTime *time.Time `json:"publishTime" form:"publishTime"`
	ExpireTime  *time.Time `json:"expireTime" form:"expireTime"`
}

// UpdateNotificationRequest 更新通知请求
type UpdateNotificationRequest struct {
	ID          uint       `json:"id" binding:"required"`
	Title       string     `json:"title" binding:"required"`
	Content     string     `json:"content" binding:"required"`
	Type        string     `json:"type" binding:"required"`
	Priority    string     `json:"priority" binding:"required"`
	TargetType  string     `json:"targetType" binding:"required"`
	TargetIds   []uint     `json:"targetIds"`
	Status      string     `json:"status" binding:"required"`
	PublishTime *time.Time `json:"publishTime"`
	ExpireTime  *time.Time `json:"expireTime"`
}

// SendNotificationRequest 发送通知请求
type SendNotificationRequest struct {
	Title      string `json:"title" binding:"required"`
	Content    string `json:"content" binding:"required"`
	Type       string `json:"type" binding:"required"`
	Priority   string `json:"priority" binding:"required"`
	TargetType string `json:"targetType" binding:"required"`
	TargetIds  []uint `json:"targetIds"`
	RoleIds    []uint `json:"roleIds"`
}

// MarkReadRequest 标记已读请求
type MarkReadRequest struct {
	NotificationIds []uint `json:"notificationIds" binding:"required"`
}

// UserNotificationSearch 用户通知搜索请求
type UserNotificationSearch struct {
	request.PageInfo
	IsRead    *bool      `json:"isRead" form:"isRead"`
	Type      string     `json:"type" form:"type"`
	Priority  string     `json:"priority" form:"priority"`
	StartTime *time.Time `json:"startTime" form:"startTime"`
	EndTime   *time.Time `json:"endTime" form:"endTime"`
}
