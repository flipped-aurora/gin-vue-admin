// 自动生成模板Course
package webcms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// Course 结构体
type Course struct {
	global.GVA_MODEL
	CateId         string         `json:"cateId" form:"cateId" gorm:"column:cate_id;comment:栏目分类id;size:10;"`
	Title          string         `json:"title" form:"title" gorm:"column:title;comment:产品名称;size:191;"`
	Thumb          string         `json:"thumb" form:"thumb" gorm:"column:thumb;comment:缩略图;size:191;"`
	ProductThumbs  string         `json:"productThumbs" form:"productThumbs" gorm:"column:product_thumbs;comment:产品图片集;size:9999;"`
	ProductDetails string         `json:"productDetails" form:"productDetails" gorm:"column:product_details;comment:产品详细信息;type:text;"`
	ProductClause  string         `json:"productClause" form:"productClause" gorm:"column:product_clause;comment:付款和装运条款;type:text;"`
	Desc           string         `json:"desc" form:"desc" gorm:"column:desc;comment:产品简介;type:text;"`
	CreatedBy      string         `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:上传用户id;size:191;"`
	Enable         *bool          `json:"enable" form:"enable" gorm:"default:1;column:enable;comment:状态 1 显示;2 隐藏;"`
	SysUser        system.SysUser `json:"sysUser" form:"sysUser" gorm:"foreignKey:CreatedBy;references:UUID"`
	Sort           *int           `json:"sort" form:"sort" gorm:"default:0;column:sort;comment:排序标记;size:19;"`
	Url            string         `json:"url"  gorm:"-"`
}

// TableName Course 表名
func (Course) TableName() string {
	return "course"
}
