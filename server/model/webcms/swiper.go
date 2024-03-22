// 自动生成模板Swiper
package webcms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Swiper 结构体
type Swiper struct {
	global.GVA_MODEL
	Siteid   string `json:"siteid" form:"siteid" gorm:"column:siteid;comment:站点id;size:32;"`
	Title    string `json:"title" form:"title" gorm:"column:title;comment:标题;size:32;"`
	Desc     string `json:"desc" form:"desc" gorm:"column:desc;comment:描述;size:255;"`
	Thumb    string `json:"thumb" form:"thumb" gorm:"column:thumb;comment:轮播图;size:255;"`
	WapThumb string `json:"wapThumb" form:"wapThumb" gorm:"column:wap_thumb;comment:wap轮播图;size:255;"`
	Url      string `json:"url" form:"url" gorm:"column:url;comment:跳转链接;size:255;"`
	Enable   *bool  `json:"enable" form:"enable" gorm:"column:enable;comment:是否显示 1是 2否;"`
}

// TableName Swiper 表名
func (Swiper) TableName() string {
	return "swiper"
}
