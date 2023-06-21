package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type InventoryService struct {
}

// CreateInventory 创建Inventory记录
// Author [piexlmax](https://github.com/piexlmax)
func (inventoryService *InventoryService) CreateInventory(inventory *clothing.Inventory) (err error) {
	err = global.GVA_DB.Create(inventory).Error
	return err
}

// DeleteInventory 删除Inventory记录
// Author [piexlmax](https://github.com/piexlmax)
func (inventoryService *InventoryService) DeleteInventory(inventory clothing.Inventory) (err error) {
	err = global.GVA_DB.Delete(&inventory).Error
	return err
}

// DeleteInventoryByIds 批量删除Inventory记录
// Author [piexlmax](https://github.com/piexlmax)
func (inventoryService *InventoryService) DeleteInventoryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]clothing.Inventory{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateInventory 更新Inventory记录
// Author [piexlmax](https://github.com/piexlmax)
func (inventoryService *InventoryService) UpdateInventory(inventory clothing.Inventory) (err error) {
	err = global.GVA_DB.Save(&inventory).Error
	return err
}

// GetInventory 根据id获取Inventory记录
// Author [piexlmax](https://github.com/piexlmax)
func (inventoryService *InventoryService) GetInventory(id uint) (inventory clothing.Inventory, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&inventory).Error
	return
}

// GetInventoryInfoList 分页获取Inventory记录
// Author [piexlmax](https://github.com/piexlmax)
func (inventoryService *InventoryService) GetInventoryInfoList(info clothingReq.InventorySearch) (list []clothing.Inventory, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&clothing.Inventory{})
	var inventorys []clothing.Inventory
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if len(info.Size) > 0 {
		db = db.Where("size = ?", info.Size)
	}
	if info.CroppingRecordID != 0 {
		db = db.Where("cropping_record_id = ?", info.CroppingRecordID)
	}
	if info.ProcessID != 0 {
		db = db.Where("process_id = ?", info.ProcessID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&inventorys).Error
	return inventorys, total, err
}
