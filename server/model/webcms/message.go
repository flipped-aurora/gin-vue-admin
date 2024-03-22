// 自动生成模板Message
package webcms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Message 结构体
type Message struct {
	global.GVA_MODEL
	Name      string `json:"name" form:"name" gorm:"column:name;comment:姓名;size:32;"`
	Email     string `json:"email" form:"email" gorm:"column:email;comment:邮箱;size:32;"`
	Tel       string `json:"tel" form:"tel" gorm:"column:tel;comment:电话;size:32;"`
	Country   string `json:"country" form:"country" gorm:"column:country;comment:国家;size:32;"`
	Company   string `json:"company" form:"company" gorm:"column:company;comment:公司;size:32;"`
	Content   string `json:"content" form:"content" gorm:"column:content;comment:内容;"`
	Captcha   string `json:"captcha" gorm:"-"`   //验证码
	Captchaid string `json:"captchaid" gorm:"-"` //验证码id
}

// TableName Message 表名
func (Message) TableName() string {
	return "message"
}
