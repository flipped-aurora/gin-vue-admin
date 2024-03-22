// 自动生成模板Links
package webcms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Links 结构体
type Links struct {
	global.GVA_MODEL
	Title  string `json:"title" form:"title" gorm:"column:title;comment:名称;size:32;"`
	Url    string `json:"url" form:"url" gorm:"column:url;comment:跳转链接;size:32;"`
	Thumb  string `json:"thumb" form:"thumb" gorm:"column:thumb;comment:网站logo;size:191;"`
	Type   string `json:"type" form:"type" gorm:"column:type;type:enum('文字','图片');comment:类型 1 文字 2 图片;"`
	Enable *bool  `json:"enable" form:"enable" gorm:"column:enable;comment:用户是否显示 1正常 2冻结;"`
}

// TableName Links 表名
func (Links) TableName() string {
	return "links"
}
