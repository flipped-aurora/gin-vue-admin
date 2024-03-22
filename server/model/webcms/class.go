// 自动生成模板Class
package webcms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Class 结构体
type Class struct {
	global.GVA_MODEL
	Siteid    string    `json:"siteid" form:"siteid" gorm:"column:siteid;comment:站点id;size:32;"`
	CateId    string    `json:"cateId" form:"cateId" gorm:"column:cate_id;comment:栏目分类id;size:10;"`
	Title     string    `json:"title" form:"title" gorm:"column:title;comment:标题;size:191;"`
	Thumb     string    `json:"thumb" form:"thumb" gorm:"column:thumb;comment:缩略图;size:191;"`
	Desc      string    `json:"desc" form:"desc" gorm:"column:desc;comment:介绍;size:1024;"`
	Content   string    `json:"content" form:"content" gorm:"column:content;comment:介绍;type:text;"`
	Enable    *bool     `json:"enable" form:"enable" gorm:"column:enable;comment:状态 1显示;2 隐藏;"`
	CateMenus CateMenus `json:"catemenus" gorm:"foreignKey:CateId"`
	Url       string    `json:"url"  gorm:"-"`
}

// TableName Class 表名
func (Class) TableName() string {
	return "class"
}
