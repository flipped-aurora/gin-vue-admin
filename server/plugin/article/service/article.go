package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/article/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/article/model/request"
	"gorm.io/gorm"
)

var Article = new(article)

type article struct{}

func (s *article) CreateArticle(a *model.Article) error {
	return global.GVA_DB.Create(a).Error
}

func (s *article) DeleteArticle(id string) error {
	return global.GVA_DB.Delete(&model.Article{}, "id = ?", id).Error
}

func (s *article) DeleteArticleByIds(ids []string) error {
	return global.GVA_DB.Delete(&[]model.Article{}, "id in ?", ids).Error
}

func (s *article) UpdateArticle(a model.Article) error {
	return global.GVA_DB.Model(&model.Article{}).Where("id = ?", a.ID).Updates(&a).Error
}

func (s *article) GetArticle(id string) (model.Article, error) {
	var a model.Article
	err := global.GVA_DB.Where("id = ?", id).First(&a).Error
	return a, err
}

func (s *article) IncrementReadCount(id string) error {
	return global.GVA_DB.Model(&model.Article{}).Where("id = ?", id).UpdateColumn("read_count", gorm.Expr("read_count + 1")).Error
}

func (s *article) GetArticleList(info request.ArticleSearch) (list []model.Article, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.Article{})
	if info.Title != "" {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
	}
	if info.Lang != "" {
		db = db.Where("lang = ?", info.Lang)
	}
	if info.Status != nil {
		db = db.Where("status = ?", *info.Status)
	}
	if info.CategoryID > 0 {
		db = db.Where("category_id = ?", info.CategoryID)
	}
	if info.StartTime != nil && info.EndTime != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartTime, info.EndTime)
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}
	if limit > 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Order("created_at DESC").Find(&list).Error
	return
}
