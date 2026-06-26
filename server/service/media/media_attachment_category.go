package media

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/media"
	"gorm.io/gorm"
)

type AttachmentCategoryService struct{}

// AddCategory 创建/更新的分类
func (a *AttachmentCategoryService) AddCategory(req *media.AttachmentCategory) (err error) {
	// 检查是否已存在相同名称的分类
	if (!errors.Is(global.GVA_DB.Take(&media.AttachmentCategory{}, "name = ? and pid = ?", req.Name, req.Pid).Error, gorm.ErrRecordNotFound)) {
		return errors.New("分类名称已存在")
	}
	if req.ID > 0 {
		if err = global.GVA_DB.Model(&media.AttachmentCategory{}).Where("id = ?", req.ID).Updates(&media.AttachmentCategory{
			Name: req.Name,
			Pid:  req.Pid,
		}).Error; err != nil {
			return err
		}
	} else {
		if err = global.GVA_DB.Create(&media.AttachmentCategory{
			Name: req.Name,
			Pid:  req.Pid,
		}).Error; err != nil {
			return err
		}
	}
	return nil
}

// DeleteCategory 删除分类
func (a *AttachmentCategoryService) DeleteCategory(id *int) error {
	var childCount int64
	global.GVA_DB.Model(&media.AttachmentCategory{}).Where("pid = ?", id).Count(&childCount)
	if childCount > 0 {
		return errors.New("请先删除子级")
	}
	return global.GVA_DB.Where("id = ?", id).Unscoped().Delete(&media.AttachmentCategory{}).Error
}

// GetCategoryList 分类列表
func (a *AttachmentCategoryService) GetCategoryList() (res []*media.AttachmentCategory, err error) {
	var fileLists []media.AttachmentCategory
	err = global.GVA_DB.Model(&media.AttachmentCategory{}).Find(&fileLists).Error
	if err != nil {
		return res, err
	}
	return a.getChildrenList(fileLists, 0), nil
}

// getChildrenList 子类
func (a *AttachmentCategoryService) getChildrenList(categories []media.AttachmentCategory, parentID uint) []*media.AttachmentCategory {
	var tree []*media.AttachmentCategory
	for _, category := range categories {
		if category.Pid == parentID {
			category.Children = a.getChildrenList(categories, category.ID)
			tree = append(tree, &category)
		}
	}
	return tree
}
