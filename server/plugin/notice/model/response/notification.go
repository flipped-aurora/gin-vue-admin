package response

import (
	"time"
)

// NotificationResponse 通知响应
type NotificationResponse struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Content     string     `json:"content"`
	Type        string     `json:"type"`
	Priority    string     `json:"priority"`
	SenderId    uint       `json:"senderId"`
	SenderName  string     `json:"senderName"`
	TargetType  string     `json:"targetType"`
	TargetIds   string     `json:"targetIds"`
	Status      string     `json:"status"`
	PublishTime *time.Time `json:"publishTime"`
	ExpireTime  *time.Time `json:"expireTime"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

// UserNotificationResponse 用户通知响应
type UserNotificationResponse struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Content     string     `json:"content"`
	Type        string     `json:"type"`
	Priority    string     `json:"priority"`
	SenderId    uint       `json:"senderId"`
	SenderName  string     `json:"senderName"`
	IsRead      bool       `json:"isRead"`
	ReadTime    *time.Time `json:"readTime"`
	PublishTime *time.Time `json:"publishTime"`
	ExpireTime  *time.Time `json:"expireTime"`
	CreatedAt   time.Time  `json:"createdAt"`
}

// NotificationStatsResponse 通知统计响应
type NotificationStatsResponse struct {
	TotalCount  int64 `json:"totalCount"`
	UnreadCount int64 `json:"unreadCount"`
	ReadCount   int64 `json:"readCount"`
}

// OnlineUserResponse 在线用户响应
type OnlineUserResponse struct {
	UserId         uint      `json:"userId"`         // 用户ID
	Username       string    `json:"username"`       // 用户名
	NickName       string    `json:"nickName"`       // 昵称
	SocketId       string    `json:"socketId"`       // Socket连接ID
	LastActiveTime time.Time `json:"lastActiveTime"` // 最后活跃时间
	Status         string    `json:"status"`         // 状态
	OnlineDuration int64     `json:"onlineDuration"` // 在线时长（分钟）
}

// OnlineUserStatsResponse 在线用户统计响应结构体
type OnlineUserStatsResponse struct {
	TotalOnline   int64 `json:"totalOnline"`   // 当前在线总数
	TodayLogin    int64 `json:"todayLogin"`    // 今日登录人数
	PeakOnline    int64 `json:"peakOnline"`    // 峰值在线人数
	AverageOnline int64 `json:"averageOnline"` // 平均在线人数
}

// SendNotificationResponse 发送通知响应
type SendNotificationResponse struct {
	NotificationId uint   `json:"notificationId"`
	SuccessCount   int    `json:"successCount"`
	FailCount      int    `json:"failCount"`
	Message        string `json:"message"`
}
