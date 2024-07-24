package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
)

// Info 公告 结构体
type Info struct {
	global.GVA_MODEL
	Title       string         `json:"title" form:"title" gorm:"column:title;comment:公告标题;"`                                             //标题
	Content     string         `json:"content" form:"content" gorm:"column:content;comment:公告内容;type:text;"`                             //内容
	UserID      *int           `json:"userID" form:"userID" gorm:"column:user_id;comment:发布者;"`                                          //作者
	Attachments datatypes.JSON `json:"attachments" form:"attachments" gorm:"column:attachments;comment:相关附件;"swaggertype:"array,object"` //附件
}

// TableName 公告 Info自定义表名 gva_announcements_info
func (Info) TableName() string {
	return "gva_announcements_info"
}
