package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/article/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/article/model/request"
)

var Category = new(category)

type category struct{}

func (s *category) CreateCategory(c *model.Category) error {
	return global.GVA_DB.Create(c).Error
}

func (s *category) DeleteCategory(id string) error {
	return global.GVA_DB.Delete(&model.Category{}, "id = ?", id).Error
}

func (s *category) UpdateCategory(c model.Category) error {
	return global.GVA_DB.Model(&model.Category{}).Where("id = ?", c.ID).Updates(&c).Error
}

func (s *category) GetCategory(id string) (model.Category, error) {
	var c model.Category
	err := global.GVA_DB.Where("id = ?", id).First(&c).Error
	return c, err
}

func (s *category) GetCategoryList(info request.CategorySearch) (list []model.Category, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.Category{})
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Lang != "" {
		db = db.Where("lang = ?", info.Lang)
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}
	if limit > 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Order("sort ASC, created_at DESC").Find(&list).Error
	return
}

func (s *category) GetAllCategories() (list []model.Category, err error) {
	err = global.GVA_DB.Order("sort ASC, created_at DESC").Find(&list).Error
	return
}
