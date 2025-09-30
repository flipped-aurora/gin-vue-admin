package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Notification 通知信息模型
type Notification struct {
	global.GVA_MODEL
	Title       string     `json:"title" gorm:"column:title;type:varchar(255);not null;comment:通知标题"`
	Content     string     `json:"content" gorm:"column:content;type:text;not null;comment:通知内容"`
	Type        string     `json:"type" gorm:"column:type;type:varchar(50);not null;default:system;comment:通知类型：system-系统通知,personal-个人通知,role-角色通知"`
	Priority    string     `json:"priority" gorm:"column:priority;type:varchar(20);not null;default:normal;comment:通知优先级：low-低,normal-普通,high-高,urgent-紧急"`
	SenderId    uint       `json:"senderId" gorm:"column:sender_id;not null;comment:发送者用户ID"`
	TargetType  string     `json:"targetType" gorm:"column:target_type;type:varchar(20);not null;default:all;comment:目标类型：all-全部用户,user-指定用户,role-指定角色"`
	TargetIds   string     `json:"targetIds" gorm:"column:target_ids;type:text;comment:目标用户或角色ID列表，JSON格式"`
	Status      string     `json:"status" gorm:"column:status;type:varchar(20);not null;default:draft;comment:通知状态：draft-草稿,published-已发布,expired-已过期"`
	PublishTime *time.Time `json:"publishTime" gorm:"column:publish_time;comment:通知发布时间"`
	ExpireTime  *time.Time `json:"expireTime" gorm:"column:expire_time;comment:通知过期时间"`
}

// TableName 设置表名
func (Notification) TableName() string {
	return "notice_notifications"
}
