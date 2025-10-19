package system

import (
	"fmt"
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateSysDictionaryDetail
//@description: 创建字典详情数据
//@param: sysDictionaryDetail model.SysDictionaryDetail
//@return: err error

type DictionaryDetailService struct{}

var DictionaryDetailServiceApp = new(DictionaryDetailService)

func (dictionaryDetailService *DictionaryDetailService) CreateSysDictionaryDetail(sysDictionaryDetail system.SysDictionaryDetail) (err error) {
	// 计算层级和路径
	if sysDictionaryDetail.ParentID != nil {
		var parent system.SysDictionaryDetail
		err = global.GVA_DB.First(&parent, *sysDictionaryDetail.ParentID).Error
		if err != nil {
			return err
		}
		sysDictionaryDetail.Level = parent.Level + 1
		if parent.Path == "" {
			sysDictionaryDetail.Path = strconv.Itoa(int(parent.ID))
		} else {
			sysDictionaryDetail.Path = parent.Path + "," + strconv.Itoa(int(parent.ID))
		}
	} else {
		sysDictionaryDetail.Level = 0
		sysDictionaryDetail.Path = ""
	}

	err = global.GVA_DB.Create(&sysDictionaryDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSysDictionaryDetail
//@description: 删除字典详情数据
//@param: sysDictionaryDetail model.SysDictionaryDetail
//@return: err error

func (dictionaryDetailService *DictionaryDetailService) DeleteSysDictionaryDetail(sysDictionaryDetail system.SysDictionaryDetail) (err error) {
	// 检查是否有子项
	var count int64
	err = global.GVA_DB.Model(&system.SysDictionaryDetail{}).Where("parent_id = ?", sysDictionaryDetail.ID).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("该字典详情下还有子项，无法删除")
	}

	err = global.GVA_DB.Delete(&sysDictionaryDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateSysDictionaryDetail
//@description: 更新字典详情数据
//@param: sysDictionaryDetail *model.SysDictionaryDetail
//@return: err error

func (dictionaryDetailService *DictionaryDetailService) UpdateSysDictionaryDetail(sysDictionaryDetail *system.SysDictionaryDetail) (err error) {
	// 如果更新了父级ID，需要重新计算层级和路径
	if sysDictionaryDetail.ParentID != nil {
		var parent system.SysDictionaryDetail
		err = global.GVA_DB.First(&parent, *sysDictionaryDetail.ParentID).Error
		if err != nil {
			return err
		}

		// 检查循环引用
		if dictionaryDetailService.checkCircularReference(sysDictionaryDetail.ID, *sysDictionaryDetail.ParentID) {
			return fmt.Errorf("不能将字典详情设置为自己或其子项的父级")
		}

		sysDictionaryDetail.Level = parent.Level + 1
		if parent.Path == "" {
			sysDictionaryDetail.Path = strconv.Itoa(int(parent.ID))
		} else {
			sysDictionaryDetail.Path = parent.Path + "," + strconv.Itoa(int(parent.ID))
		}
	} else {
		sysDictionaryDetail.Level = 0
		sysDictionaryDetail.Path = ""
	}

	err = global.GVA_DB.Save(sysDictionaryDetail).Error
	if err != nil {
		return err
	}

	// 更新所有子项的层级和路径
	return dictionaryDetailService.updateChildrenLevelAndPath(sysDictionaryDetail.ID)
}

// checkCircularReference 检查循环引用
func (dictionaryDetailService *DictionaryDetailService) checkCircularReference(id, parentID uint) bool {
	if id == parentID {
		return true
	}

	var parent system.SysDictionaryDetail
	err := global.GVA_DB.First(&parent, parentID).Error
	if err != nil {
		return false
	}

	if parent.ParentID == nil {
		return false
	}

	return dictionaryDetailService.checkCircularReference(id, *parent.ParentID)
}

// updateChildrenLevelAndPath 更新子项的层级和路径
func (dictionaryDetailService *DictionaryDetailService) updateChildrenLevelAndPath(parentID uint) error {
	var children []system.SysDictionaryDetail
	err := global.GVA_DB.Where("parent_id = ?", parentID).Find(&children).Error
	if err != nil {
		return err
	}

	var parent system.SysDictionaryDetail
	err = global.GVA_DB.First(&parent, parentID).Error
	if err != nil {
		return err
	}

	for _, child := range children {
		child.Level = parent.Level + 1
		if parent.Path == "" {
			child.Path = strconv.Itoa(int(parent.ID))
		} else {
			child.Path = parent.Path + "," + strconv.Itoa(int(parent.ID))
		}

		err = global.GVA_DB.Save(&child).Error
		if err != nil {
			return err
		}

		// 递归更新子项的子项
		err = dictionaryDetailService.updateChildrenLevelAndPath(child.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSysDictionaryDetail
//@description: 根据id获取字典详情单条数据
//@param: id uint
//@return: sysDictionaryDetail system.SysDictionaryDetail, err error

func (dictionaryDetailService *DictionaryDetailService) GetSysDictionaryDetail(id uint) (sysDictionaryDetail system.SysDictionaryDetail, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysDictionaryDetail).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSysDictionaryDetailInfoList
//@description: 分页获取字典详情列表
//@param: info request.SysDictionaryDetailSearch
//@return: list interface{}, total int64, err error

func (dictionaryDetailService *DictionaryDetailService) GetSysDictionaryDetailInfoList(info request.SysDictionaryDetailSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&system.SysDictionaryDetail{})
	var sysDictionaryDetails []system.SysDictionaryDetail
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Label != "" {
		db = db.Where("label LIKE ?", "%"+info.Label+"%")
	}
	if info.Value != "" {
		db = db.Where("value = ?", info.Value)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.SysDictionaryID != 0 {
		db = db.Where("sys_dictionary_id = ?", info.SysDictionaryID)
	}
	if info.ParentID != nil {
		db = db.Where("parent_id = ?", *info.ParentID)
	}
	if info.Level != nil {
		db = db.Where("level = ?", *info.Level)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("sort").Order("id").Find(&sysDictionaryDetails).Error
	return sysDictionaryDetails, total, err
}

// 按照字典id获取字典全部内容的方法
func (dictionaryDetailService *DictionaryDetailService) GetDictionaryList(dictionaryID uint) (list []system.SysDictionaryDetail, err error) {
	var sysDictionaryDetails []system.SysDictionaryDetail
	err = global.GVA_DB.Find(&sysDictionaryDetails, "sys_dictionary_id = ?", dictionaryID).Error
	return sysDictionaryDetails, err
}

// GetDictionaryTreeList 获取字典树形结构列表
func (dictionaryDetailService *DictionaryDetailService) GetDictionaryTreeList(dictionaryID uint) (list []system.SysDictionaryDetail, err error) {
	var sysDictionaryDetails []system.SysDictionaryDetail
	// 只获取顶级项目（parent_id为空）
	err = global.GVA_DB.Where("sys_dictionary_id = ? AND parent_id IS NULL", dictionaryID).Order("sort").Find(&sysDictionaryDetails).Error
	if err != nil {
		return nil, err
	}

	// 递归加载子项并设置disabled属性
	for i := range sysDictionaryDetails {
		// 设置disabled属性：当status为false时，disabled为true
		if sysDictionaryDetails[i].Status != nil {
			sysDictionaryDetails[i].Disabled = !*sysDictionaryDetails[i].Status
		} else {
			sysDictionaryDetails[i].Disabled = false // 默认不禁用
		}
		
		err = dictionaryDetailService.loadChildren(&sysDictionaryDetails[i])
		if err != nil {
			return nil, err
		}
	}

	return sysDictionaryDetails, nil
}

// loadChildren 递归加载子项
func (dictionaryDetailService *DictionaryDetailService) loadChildren(detail *system.SysDictionaryDetail) error {
	var children []system.SysDictionaryDetail
	err := global.GVA_DB.Where("parent_id = ?", detail.ID).Order("sort").Find(&children).Error
	if err != nil {
		return err
	}

	for i := range children {
		// 设置disabled属性：当status为false时，disabled为true
		if children[i].Status != nil {
			children[i].Disabled = !*children[i].Status
		} else {
			children[i].Disabled = false // 默认不禁用
		}
		
		err = dictionaryDetailService.loadChildren(&children[i])
		if err != nil {
			return err
		}
	}

	detail.Children = children
	return nil
}

// GetDictionaryDetailsByParent 根据父级ID获取字典详情
func (dictionaryDetailService *DictionaryDetailService) GetDictionaryDetailsByParent(req request.GetDictionaryDetailsByParentRequest) (list []system.SysDictionaryDetail, err error) {
	db := global.GVA_DB.Model(&system.SysDictionaryDetail{}).Where("sys_dictionary_id = ?", req.SysDictionaryID)

	if req.ParentID != nil {
		db = db.Where("parent_id = ?", *req.ParentID)
	} else {
		db = db.Where("parent_id IS NULL")
	}

	err = db.Order("sort").Find(&list).Error
	if err != nil {
		return list, err
	}

	// 设置disabled属性
	for i := range list {
		if list[i].Status != nil {
			list[i].Disabled = !*list[i].Status
		} else {
			list[i].Disabled = false // 默认不禁用
		}
	}

	// 如果需要包含子级数据，使用递归方式加载所有层级的子项
	if req.IncludeChildren {
		for i := range list {
			err = dictionaryDetailService.loadChildren(&list[i])
			if err != nil {
				return list, err
			}
		}
	}

	return list, err
}

// 按照字典type获取字典全部内容的方法
func (dictionaryDetailService *DictionaryDetailService) GetDictionaryListByType(t string) (list []system.SysDictionaryDetail, err error) {
	var sysDictionaryDetails []system.SysDictionaryDetail
	db := global.GVA_DB.Model(&system.SysDictionaryDetail{}).Joins("JOIN sys_dictionaries ON sys_dictionaries.id = sys_dictionary_details.sys_dictionary_id")
	err = db.Find(&sysDictionaryDetails, "type = ?", t).Error
	return sysDictionaryDetails, err
}

// GetDictionaryTreeListByType 根据字典类型获取树形结构
func (dictionaryDetailService *DictionaryDetailService) GetDictionaryTreeListByType(t string) (list []system.SysDictionaryDetail, err error) {
	var sysDictionaryDetails []system.SysDictionaryDetail
	db := global.GVA_DB.Model(&system.SysDictionaryDetail{}).
		Joins("JOIN sys_dictionaries ON sys_dictionaries.id = sys_dictionary_details.sys_dictionary_id").
		Where("sys_dictionaries.type = ? AND sys_dictionary_details.parent_id IS NULL", t).
		Order("sys_dictionary_details.sort")

	err = db.Find(&sysDictionaryDetails).Error
	if err != nil {
		return nil, err
	}

	// 递归加载子项并设置disabled属性
	for i := range sysDictionaryDetails {
		// 设置disabled属性：当status为false时，disabled为true
		if sysDictionaryDetails[i].Status != nil {
			sysDictionaryDetails[i].Disabled = !*sysDictionaryDetails[i].Status
		} else {
			sysDictionaryDetails[i].Disabled = false // 默认不禁用
		}
		
		err = dictionaryDetailService.loadChildren(&sysDictionaryDetails[i])
		if err != nil {
			return nil, err
		}
	}

	return sysDictionaryDetails, nil
}

// 按照字典id+字典内容value获取单条字典内容
func (dictionaryDetailService *DictionaryDetailService) GetDictionaryInfoByValue(dictionaryID uint, value string) (detail system.SysDictionaryDetail, err error) {
	var sysDictionaryDetail system.SysDictionaryDetail
	err = global.GVA_DB.First(&sysDictionaryDetail, "sys_dictionary_id = ? and value = ?", dictionaryID, value).Error
	return sysDictionaryDetail, err
}

// 按照字典type+字典内容value获取单条字典内容
func (dictionaryDetailService *DictionaryDetailService) GetDictionaryInfoByTypeValue(t string, value string) (detail system.SysDictionaryDetail, err error) {
	var sysDictionaryDetails system.SysDictionaryDetail
	db := global.GVA_DB.Model(&system.SysDictionaryDetail{}).Joins("JOIN sys_dictionaries ON sys_dictionaries.id = sys_dictionary_details.sys_dictionary_id")
	err = db.First(&sysDictionaryDetails, "sys_dictionaries.type = ? and sys_dictionary_details.value = ?", t, value).Error
	return sysDictionaryDetails, err
}

// GetDictionaryPath 获取字典详情的完整路径
func (dictionaryDetailService *DictionaryDetailService) GetDictionaryPath(id uint) (path []system.SysDictionaryDetail, err error) {
	var detail system.SysDictionaryDetail
	err = global.GVA_DB.First(&detail, id).Error
	if err != nil {
		return nil, err
	}

	path = append(path, detail)

	if detail.ParentID != nil {
		parentPath, err := dictionaryDetailService.GetDictionaryPath(*detail.ParentID)
		if err != nil {
			return nil, err
		}
		path = append(parentPath, path...)
	}

	return path, nil
}

// GetDictionaryPathByValue 根据值获取字典详情的完整路径
func (dictionaryDetailService *DictionaryDetailService) GetDictionaryPathByValue(dictionaryID uint, value string) (path []system.SysDictionaryDetail, err error) {
	detail, err := dictionaryDetailService.GetDictionaryInfoByValue(dictionaryID, value)
	if err != nil {
		return nil, err
	}

	return dictionaryDetailService.GetDictionaryPath(detail.ID)
}
