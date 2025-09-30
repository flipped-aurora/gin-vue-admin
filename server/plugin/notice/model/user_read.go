package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// UserRead 用户通知读取状态模型
type UserRead struct {
	global.GVA_MODEL
	UserId         uint       `json:"userId" gorm:"column:user_id;not null;index;comment:用户ID"`
	NotificationId uint       `json:"notificationId" gorm:"column:notification_id;not null;index;comment:通知ID"`
	IsRead         bool       `json:"isRead" gorm:"column:is_read;not null;default:false;comment:是否已读"`
	ReadTime       *time.Time `json:"readTime" gorm:"column:read_time;comment:通知读取时间"`
}

// TableName 设置表名
func (UserRead) TableName() string {
	return "notice_user_reads"
}
