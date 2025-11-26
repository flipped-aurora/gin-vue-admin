package system

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gin-gonic/gin"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gorm.io/gorm"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateSysDictionary
//@description: 创建字典数据
//@param: sysDictionary model.SysDictionary
//@return: err error

type DictionaryService struct{}

var DictionaryServiceApp = new(DictionaryService)

func (dictionaryService *DictionaryService) CreateSysDictionary(sysDictionary system.SysDictionary) (err error) {
	if (!errors.Is(global.GVA_DB.First(&system.SysDictionary{}, "type = ?", sysDictionary.Type).Error, gorm.ErrRecordNotFound)) {
		return errors.New("存在相同的type，不允许创建")
	}
	err = global.GVA_DB.Create(&sysDictionary).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSysDictionary
//@description: 删除字典数据
//@param: sysDictionary model.SysDictionary
//@return: err error

func (dictionaryService *DictionaryService) DeleteSysDictionary(sysDictionary system.SysDictionary) (err error) {
	err = global.GVA_DB.Where("id = ?", sysDictionary.ID).Preload("SysDictionaryDetails").First(&sysDictionary).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("请不要搞事")
	}
	if err != nil {
		return err
	}
	err = global.GVA_DB.Delete(&sysDictionary).Error
	if err != nil {
		return err
	}

	if sysDictionary.SysDictionaryDetails != nil {
		return global.GVA_DB.Where("sys_dictionary_id=?", sysDictionary.ID).Delete(sysDictionary.SysDictionaryDetails).Error
	}
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateSysDictionary
//@description: 更新字典数据
//@param: sysDictionary *model.SysDictionary
//@return: err error

func (dictionaryService *DictionaryService) UpdateSysDictionary(sysDictionary *system.SysDictionary) (err error) {
	var dict system.SysDictionary
	sysDictionaryMap := map[string]interface{}{
		"Name":     sysDictionary.Name,
		"Type":     sysDictionary.Type,
		"Status":   sysDictionary.Status,
		"Desc":     sysDictionary.Desc,
		"ParentID": sysDictionary.ParentID,
	}
	err = global.GVA_DB.Where("id = ?", sysDictionary.ID).First(&dict).Error
	if err != nil {
		global.GVA_LOG.Debug(err.Error())
		return errors.New("查询字典数据失败")
	}
	if dict.Type != sysDictionary.Type {
		if !errors.Is(global.GVA_DB.First(&system.SysDictionary{}, "type = ?", sysDictionary.Type).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同的type，不允许创建")
		}
	}

	// 检查是否会形成循环引用
	if sysDictionary.ParentID != nil && *sysDictionary.ParentID != 0 {
		if err := dictionaryService.checkCircularReference(sysDictionary.ID, *sysDictionary.ParentID); err != nil {
			return err
		}
	}

	err = global.GVA_DB.Model(&dict).Updates(sysDictionaryMap).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSysDictionary
//@description: 根据id或者type获取字典单条数据
//@param: Type string, Id uint
//@return: err error, sysDictionary model.SysDictionary

func (dictionaryService *DictionaryService) GetSysDictionary(Type string, Id uint, status *bool) (sysDictionary system.SysDictionary, err error) {
	var flag = false
	if status == nil {
		flag = true
	} else {
		flag = *status
	}
	err = global.GVA_DB.Where("(type = ? OR id = ?) and status = ?", Type, Id, flag).Preload("SysDictionaryDetails", func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ? and deleted_at is null", true).Order("sort")
	}).First(&sysDictionary).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: GetSysDictionaryInfoList
//@description: 分页获取字典列表
//@param: info request.SysDictionarySearch
//@return: err error, list interface{}, total int64

func (dictionaryService *DictionaryService) GetSysDictionaryInfoList(c *gin.Context, req request.SysDictionarySearch) (list interface{}, err error) {
	var sysDictionarys []system.SysDictionary
	query := global.GVA_DB.WithContext(c)
	if req.Name != "" {
		query = query.Where("name LIKE ? OR type LIKE ?", "%"+req.Name+"%", "%"+req.Name+"%")
	}
	// 预加载子字典
	query = query.Preload("Children")
	err = query.Find(&sysDictionarys).Error
	return sysDictionarys, err
}

// checkCircularReference 检查是否会形成循环引用
func (dictionaryService *DictionaryService) checkCircularReference(currentID uint, parentID uint) error {
	if currentID == parentID {
		return errors.New("不能将字典设置为自己的父级")
	}

	// 递归检查父级链条
	var parent system.SysDictionary
	err := global.GVA_DB.Where("id = ?", parentID).First(&parent).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil // 父级不存在，允许设置
		}
		return err
	}

	// 如果父级还有父级，继续检查
	if parent.ParentID != nil && *parent.ParentID != 0 {
		return dictionaryService.checkCircularReference(currentID, *parent.ParentID)
	}

	return nil
}

//@author: [yourname]
//@function: ExportSysDictionary
//@description: 导出字典JSON（包含字典详情）
//@param: id uint
//@return: exportData map[string]interface{}, err error

func (dictionaryService *DictionaryService) ExportSysDictionary(id uint) (exportData map[string]interface{}, err error) {
	var dictionary system.SysDictionary
	// 查询字典及其所有详情
	err = global.GVA_DB.Where("id = ?", id).Preload("SysDictionaryDetails", func(db *gorm.DB) *gorm.DB {
		return db.Order("sort")
	}).First(&dictionary).Error
	if err != nil {
		return nil, err
	}

	// 构造导出数据
	exportData = map[string]interface{}{
		"name":    dictionary.Name,
		"type":    dictionary.Type,
		"status":  dictionary.Status,
		"desc":    dictionary.Desc,
		"details": dictionary.SysDictionaryDetails,
	}

	return exportData, nil
}

//@author: [yourname]
//@function: ImportSysDictionary
//@description: 导入字典JSON（包含字典详情）
//@param: importData map[string]interface{}
//@return: err error

func (dictionaryService *DictionaryService) ImportSysDictionary(importData map[string]interface{}) error {
	// 解析基本字典信息
	name, ok := importData["name"].(string)
	if !ok || name == "" {
		return errors.New("字典名称不能为空")
	}

	dictType, ok := importData["type"].(string)
	if !ok || dictType == "" {
		return errors.New("字典类型不能为空")
	}

	// 检查字典类型是否已存在
	if !errors.Is(global.GVA_DB.First(&system.SysDictionary{}, "type = ?", dictType).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同的type，不允许导入")
	}

	// 创建字典
	dictionary := system.SysDictionary{
		Name: name,
		Type: dictType,
	}

	// 处理status字段
	if status, ok := importData["status"].(bool); ok {
		dictionary.Status = &status
	}

	// 处理desc字段
	if desc, ok := importData["desc"].(string); ok {
		dictionary.Desc = desc
	}

	// 开启事务
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 创建字典
		if err := tx.Create(&dictionary).Error; err != nil {
			return err
		}

		// 处理字典详情
		if details, ok := importData["details"].([]interface{}); ok && len(details) > 0 {
			// 创建一个映射来跟踪旧ID到新ID的对应关系
			idMap := make(map[uint]uint)

			// 第一遍：创建所有详情记录（不设置parent_id）
			for _, detail := range details {
				detailMap, ok := detail.(map[string]interface{})
				if !ok {
					continue
				}

				label, _ := detailMap["label"].(string)
				value, _ := detailMap["value"].(string)

				if label == "" || value == "" {
					continue
				}

				detailRecord := system.SysDictionaryDetail{
					Label:           label,
					Value:           value,
					SysDictionaryID: int(dictionary.ID),
				}

				// 处理extend字段
				if extend, ok := detailMap["extend"].(string); ok {
					detailRecord.Extend = extend
				}

				// 处理status字段
				if status, ok := detailMap["status"].(bool); ok {
					detailRecord.Status = &status
				}

				// 处理sort字段
				if sort, ok := detailMap["sort"].(float64); ok {
					detailRecord.Sort = int(sort)
				}

				// 创建详情记录
				if err := tx.Create(&detailRecord).Error; err != nil {
					return err
				}

				// 记录ID映射（如果有原始ID）
				if oldID, ok := detailMap["ID"].(float64); ok {
					idMap[uint(oldID)] = detailRecord.ID
				}
			}

			// 第二遍：更新parent_id关系
			for i, detail := range details {
				detailMap, ok := detail.(map[string]interface{})
				if !ok {
					continue
				}

				// 如果有parentID，更新它
				if oldParentID, ok := detailMap["parentID"].(float64); ok && oldParentID > 0 {
					if newParentID, exists := idMap[uint(oldParentID)]; exists {
						// 获取新创建的记录ID（按顺序）
						if oldID, ok := detailMap["ID"].(float64); ok {
							if newID, exists := idMap[uint(oldID)]; exists {
								if err := tx.Model(&system.SysDictionaryDetail{}).Where("id = ?", newID).Update("parent_id", newParentID).Error; err != nil {
									return err
								}
							}
						} else {
							// 如果没有ID，使用索引来查找
							var allDetails []system.SysDictionaryDetail
							if err := tx.Where("sys_dictionary_id = ?", dictionary.ID).Order("id").Find(&allDetails).Error; err != nil {
								return err
							}
							if i < len(allDetails) {
								if err := tx.Model(&system.SysDictionaryDetail{}).Where("id = ?", allDetails[i].ID).Update("parent_id", newParentID).Error; err != nil {
									return err
								}
							}
						}
					}
				}
			}
		}

		return nil
	})
}
