package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// OnlineUser 在线用户模型
type OnlineUser struct {
	global.GVA_MODEL
	UserId         uint      `json:"userId" gorm:"column:user_id;not null;uniqueIndex;comment:用户ID"`
	SocketId       string    `json:"socketId" gorm:"column:socket_id;type:varchar(255);not null;comment:Socket连接ID"`
	LastActiveTime time.Time `json:"lastActiveTime" gorm:"column:last_active_time;not null;comment:最后活跃时间"`
	Status         string    `json:"status" gorm:"column:status;type:varchar(20);not null;default:online;comment:在线状态：online-在线,offline-离线"`
}

// TableName 设置表名
func (OnlineUser) TableName() string {
	return "notice_online_users"
}
