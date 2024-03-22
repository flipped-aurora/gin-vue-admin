// 自动生成模板CateMenus
package webcms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CateMenus 结构体
type CateMenus struct {
	global.GVA_MODEL
	Siteid       string `json:"siteid" form:"siteid" gorm:"column:siteid;comment:站点id;size:32;"`
	MenuLevel    int    `json:"menuLevel" form:"menuLevel" gorm:"column:menu_level;comment:;size:10;"`
	ParentId     uint   `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:父菜单ID;size:10;"`
	Name         string `json:"name" form:"name" gorm:"column:name;comment:栏目名称;size:191;"`
	Short        string `json:"short" form:"short" gorm:"column:short;comment:栏目简称;size:191;"`
	Hidden       bool   `json:"hidden" form:"hidden" gorm:"column:hidden;comment:是否隐藏;"`
	Sort         *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序标记;size:19;"`
	Islink       string `json:"islink" form:"islink" gorm:"column:islink;comment:排序标记;size:191;"`
	ModeType     int    `json:"modeType" form:"modeType" gorm:"column:mode_type;comment:模型管理 1 产品管理 2 文章管理;size:1;"`
	CateType     int    `json:"cateType" form:"cateType" gorm:"column:cate_type;comment:栏目类型 1 封面 2 列表 3 链接;size:19;"`
	ListTemplate string `json:"listTemplate" form:"listTemplate" gorm:"column:list_template;comment:列表页模板;size:191;"`
	ShowTemplate string `json:"showTemplate" form:"showTemplate" gorm:"column:show_template;comment:内容页模板;size:191;"`
	CateThumb    string `json:"cateThumb" form:"cateThumb" gorm:"column:cate_thumb;comment:封面缩略图;size:191;"`
	ListThumb    string `json:"listThumb" form:"listThumb" gorm:"column:list_thumb;comment:列表缩略图;size:191;"`
	Desc         string `json:"desc" form:"desc" gorm:"column:desc;comment:栏目简介;type:text;"`
	IsJump       bool   `json:"isJump" form:"isJump" gorm:"column:isjump;comment:是否跳转子栏目;"`
	OrderType    int    `json:"orderType" form:"orderType" gorm:"column:order_type;comment:排序方式 1 升序 2 降序;size:4;default:1;"`
	PgSize       int    `json:"pageSize" form:"pageSize" gorm:"column:page_size;comment:每页数量;size:5;default:10;"`
	Url          string `json:"url"  gorm:"-"`

	Children []CateMenus `json:"children" gorm:"-"`
}

// TableName CateMenus 表名
func (CateMenus) TableName() string {
	return "cate_menus"
}
