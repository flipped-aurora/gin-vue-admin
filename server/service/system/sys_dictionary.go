package system

import (
	"encoding/json"
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

//@author: [pixelMax]
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

	// 清空字典详情中的ID、创建时间、更新时间等字段
	var cleanDetails []map[string]interface{}
	for _, detail := range dictionary.SysDictionaryDetails {
		cleanDetail := map[string]interface{}{
			"label":  detail.Label,
			"value":  detail.Value,
			"extend": detail.Extend,
			"status": detail.Status,
			"sort":   detail.Sort,
			"level":  detail.Level,
			"path":   detail.Path,
		}
		cleanDetails = append(cleanDetails, cleanDetail)
	}

	// 构造导出数据
	exportData = map[string]interface{}{
		"name":                 dictionary.Name,
		"type":                 dictionary.Type,
		"status":               dictionary.Status,
		"desc":                 dictionary.Desc,
		"sysDictionaryDetails": cleanDetails,
	}

	return exportData, nil
}

//@author: [pixelMax]
//@function: ImportSysDictionary
//@description: 导入字典JSON（包含字典详情）
//@param: jsonStr string
//@return: err error

func (dictionaryService *DictionaryService) ImportSysDictionary(jsonStr string) error {
	// 直接解析到 SysDictionary 结构体
	var importData system.SysDictionary
	if err := json.Unmarshal([]byte(jsonStr), &importData); err != nil {
		return errors.New("JSON 格式错误: " + err.Error())
	}

	// 验证必填字段
	if importData.Name == "" {
		return errors.New("字典名称不能为空")
	}
	if importData.Type == "" {
		return errors.New("字典类型不能为空")
	}

	// 检查字典类型是否已存在
	if !errors.Is(global.GVA_DB.First(&system.SysDictionary{}, "type = ?", importData.Type).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同的type，不允许导入")
	}

	// 创建字典（清空导入数据的ID和时间戳）
	dictionary := system.SysDictionary{
		Name:   importData.Name,
		Type:   importData.Type,
		Status: importData.Status,
		Desc:   importData.Desc,
	}

	// 开启事务
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 创建字典
		if err := tx.Create(&dictionary).Error; err != nil {
			return err
		}

		// 处理字典详情
		if len(importData.SysDictionaryDetails) > 0 {
			// 创建一个映射来跟踪旧ID到新ID的对应关系
			idMap := make(map[uint]uint)

			// 第一遍：创建所有详情记录
			for _, detail := range importData.SysDictionaryDetails {
				// 验证必填字段
				if detail.Label == "" || detail.Value == "" {
					continue
				}

				// 记录旧ID
				oldID := detail.ID

				// 创建新的详情记录（ID会被GORM自动设置）
				detailRecord := system.SysDictionaryDetail{
					Label:           detail.Label,
					Value:           detail.Value,
					Extend:          detail.Extend,
					Status:          detail.Status,
					Sort:            detail.Sort,
					Level:           detail.Level,
					Path:            detail.Path,
					SysDictionaryID: int(dictionary.ID),
				}

				// 创建详情记录
				if err := tx.Create(&detailRecord).Error; err != nil {
					return err
				}

				// 记录旧ID到新ID的映射
				if oldID > 0 {
					idMap[oldID] = detailRecord.ID
				}
			}

			// 第二遍：更新parent_id关系
			for _, detail := range importData.SysDictionaryDetails {
				if detail.ParentID != nil && *detail.ParentID > 0 && detail.ID > 0 {
					if newID, exists := idMap[detail.ID]; exists {
						if newParentID, parentExists := idMap[*detail.ParentID]; parentExists {
							if err := tx.Model(&system.SysDictionaryDetail{}).
								Where("id = ?", newID).
								Update("parent_id", newParentID).Error; err != nil {
								return err
							}
						}
					}
				}
			}
		}

		return nil
	})
}
