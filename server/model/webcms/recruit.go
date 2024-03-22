// 自动生成模板Recruit
package webcms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Recruit 结构体
type Recruit struct {
	global.GVA_MODEL
	Siteid    string    `json:"siteid" form:"siteid" gorm:"column:siteid;comment:站点id;size:32;"`
	CateId    string    `json:"cateId" form:"cateId" gorm:"column:cate_id;comment:栏目分类id;size:10;"`
	Job       string    `json:"job" form:"job" gorm:"column:job;comment:;size:32;"`
	Count     string    `json:"count" form:"count" gorm:"column:count;comment:;size:4;"`
	Address   string    `json:"address" form:"address" gorm:"column:address;comment:;size:191;"`
	Content   string    `json:"content" form:"content" gorm:"column:content;comment:;type:text;"`
	Enable    *bool     `json:"enable" form:"enable" gorm:"column:enable;comment:状态 1显示;2 隐藏;"`
	CateMenus CateMenus `json:"catemenus" gorm:"foreignKey:CateId"`
	Url       string    `json:"url"  gorm:"-"`
}

// TableName Recruit 表名
func (Recruit) TableName() string {
	return "recruit"
}
