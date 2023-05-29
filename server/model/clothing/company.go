// 自动生成模板Company
package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Company 结构体
type Company struct {
	global.GVA_MODEL
	UserID          uint       `json:"userID" form:"userID" gorm:"column:user_id;comment:;"`
	AgentID         uint       `json:"agentID" form:"agentID" gorm:"column:agent_id;comment:;"`
	Name            string     `json:"name" form:"name" gorm:"column:name;comment:;unique;"`
	Status          *bool      `json:"status" form:"status" gorm:"column:status;comment:;"`
	ClerkCount      int        `json:"clerkCount" form:"clerkCount" gorm:"column:clerk_count;comment:职员数量;"`
	ClerkCountLimit int        `json:"clerkCountLimit" form:"clerkCountLimit" gorm:"column:clerk_count_limit;comment:职员数量上限"`
	ExpirationAt    *time.Time `json:"expirationAt" form:"expirationAt" gorm:"column:expiration_at;comment:到期时间"`
	BeanRate        float64    `json:"beanRate" form:"beanRate" gorm:"column:bean_bean;comment:汇率"`
	CreatedBy       uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy       uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy       uint       `gorm:"column:deleted_by;comment:删除者"`
	User            AppUser    `json:"user" form:"user" gorm:"foreignKey:UserID"`
	Agent           Agent      `json:"agent" form:"agent" gorm:"foreignKey:AgentID"`
}

// TableName Company 表名
func (Company) TableName() string {
	return "company"
}
