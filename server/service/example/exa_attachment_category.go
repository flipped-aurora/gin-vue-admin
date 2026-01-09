package example

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"gorm.io/gorm"
)

type AttachmentCategoryService struct{}

// AddCategory 创建/更新的分类
func (a *AttachmentCategoryService) AddCategory(req *example.ExaAttachmentCategory) (err error) {
	// 检查是否已存在相同名称的分类
	if (!errors.Is(global.GVA_DB.Take(&example.ExaAttachmentCategory{}, "name = ? and pid = ?", req.Name, req.Pid).Error, gorm.ErrRecordNotFound)) {
		return errors.New("分类名称已存在")
	}
	if req.ID > 0 {
		if err = global.GVA_DB.Model(&example.ExaAttachmentCategory{}).Where("id = ?", req.ID).Updates(&example.ExaAttachmentCategory{
			Name: req.Name,
			Pid:  req.Pid,
		}).Error; err != nil {
			return err
		}
	} else {
		if err = global.GVA_DB.Create(&example.ExaAttachmentCategory{
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
	global.GVA_DB.Model(&example.ExaAttachmentCategory{}).Where("pid = ?", id).Count(&childCount)
	if childCount > 0 {
		return errors.New("请先删除子级")
	}
	return global.GVA_DB.Where("id = ?", id).Unscoped().Delete(&example.ExaAttachmentCategory{}).Error
}

// GetCategoryList 分类列表
func (a *AttachmentCategoryService) GetCategoryList() (res []*example.ExaAttachmentCategory, err error) {
	var fileLists []example.ExaAttachmentCategory
	err = global.GVA_DB.Model(&example.ExaAttachmentCategory{}).Find(&fileLists).Error
	if err != nil {
		return res, err
	}
	return a.getChildrenList(fileLists, 0), nil
}

// getChildrenList 子类
func (a *AttachmentCategoryService) getChildrenList(categories []example.ExaAttachmentCategory, parentID uint) []*example.ExaAttachmentCategory {
	var tree []*example.ExaAttachmentCategory
	for _, category := range categories {
		if category.Pid == parentID {
			// 重要：创建新变量避免循环变量地址问题
			// 为什么需要 category := category：
			// 1. 在 Go 1.22 之前，循环变量 category 在每次迭代中使用同一个内存地址
			//    如果直接使用 &category，所有指针都会指向最后一次迭代的值
			// 2. 虽然 Go 1.22+ 已修复此问题（每次迭代创建新变量），但显式创建新变量
			//    可以保证代码在所有 Go 版本中都能正确工作，提高可移植性
			// 3. 代码意图更明确：明确表示我们要为每个节点创建独立的内存空间
			category := category // 创建新变量，复制当前迭代的值
			category.Children = a.getChildrenList(categories, category.ID)
			tree = append(tree, &category)
		}
	}
	return tree
}
