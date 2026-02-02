package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type CertCategory struct {
	global.GVA_MODEL
	Name string `json:"name" gorm:"column:name;comment:分类名称;not null;uniqueIndex:idx_category_name"`
	Desc string `json:"desc" gorm:"column:desc;comment:分类描述"`
}

func (CertCategory) TableName() string {
	return "gva_cert_categories_list"
}
