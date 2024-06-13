package base

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type BaseService struct {
	Model interface{}
}

// Create 创建记录
func (b *BaseService) Create(record interface{}) error {
	return global.GVA_DB.Create(record).Error
}

// Delete 删除记录
func (b *BaseService) Delete(ID string, userID uint) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(b.Model).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err := tx.Delete(b.Model, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
}

// DeleteByIds 批量删除记录
func (b *BaseService) DeleteByIds(IDs []string, deleted_by uint) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(b.Model).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(b.Model).Error; err != nil {
			return err
		}
		return nil
	})
}

// Update 更新记录
func (b *BaseService) Update(record interface{}) error {
	return global.GVA_DB.Model(b.Model).Where("id = ?", record.(interface{ GetID() uint }).GetID()).Updates(record).Error
}

// Find 根据ID查询记录
func (b *BaseService) Find(ID string) (interface{}, error) {
	record := b.Model
	err := global.GVA_DB.Where("id = ?", ID).First(record).Error
	return record, err
}

// List 分页查询记录
func (b *BaseService) List(pageInfo interface {
	GetPage() int
	GetPageSize() int
	GetCondition() map[string]interface{}
}) ([]interface{}, int64, error) {
	limit := pageInfo.GetPageSize()
	offset := pageInfo.GetPageSize() * (pageInfo.GetPage() - 1)

	db := global.GVA_DB.Model(b.Model)
	var total int64
	var records []interface{}

	condition := pageInfo.GetCondition()
	for k, v := range condition {
		db = db.Where(k, v)
	}

	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&records).Error
	return records, total, err
}
