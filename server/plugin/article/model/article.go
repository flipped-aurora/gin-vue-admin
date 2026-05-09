package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Article 文章
type Article struct {
	global.GVA_MODEL
	Title      string `json:"title" form:"title" gorm:"column:title;comment:文章标题;"`
	Summary    string `json:"summary" form:"summary" gorm:"column:summary;comment:文章摘要;"`
	Content    string `json:"content" form:"content" gorm:"column:content;comment:文章内容;type:text;"`
	Cover      string `json:"cover" form:"cover" gorm:"column:cover;comment:封面图URL;"`
	Lang       string `json:"lang" form:"lang" gorm:"column:lang;size:8;default:zh;comment:语言 zh/en;"`
	Status     int    `json:"status" form:"status" gorm:"column:status;default:1;comment:状态 1发布 0草稿;"`
	ReadCount  int    `json:"readCount" form:"readCount" gorm:"column:read_count;default:0;comment:阅读次数;"`
	CategoryID uint   `json:"categoryID" form:"categoryID" gorm:"column:category_id;comment:分类ID;"`
	UserID     uint   `json:"userID" form:"userID" gorm:"column:user_id;comment:作者ID;"`
	NickName   string `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:作者昵称;"`
}

func (Article) TableName() string {
	return "gva_article"
}

// Category 文章分类
type Category struct {
	global.GVA_MODEL
	Name string `json:"name" form:"name" gorm:"column:name;comment:分类名称;"`
	Lang string `json:"lang" form:"lang" gorm:"column:lang;size:8;default:zh;comment:语言;"`
	Sort int    `json:"sort" form:"sort" gorm:"column:sort;default:0;comment:排序;"`
}

func (Category) TableName() string {
	return "gva_article_category"
}
