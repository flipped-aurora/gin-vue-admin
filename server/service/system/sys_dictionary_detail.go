package system

import (
	"fmt"
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

/*
 * 字典详情管理服务
 *
 * 本文件实现了字典详情（SysDictionaryDetail）的完整管理功能
 * 字典详情支持树形结构，可以构建多级字典数据
 * 主要功能包括：
 * 1. 字典详情的创建、更新、删除、查询
 * 2. 树形结构的层级管理和路径计算
 * 3. 循环引用检测和子项管理
 * 4. 多种查询方式：按字典ID、按类型、按值、树形结构等
 * 5. 字典路径查询功能
 *
 * 使用示例：
 *   // 创建顶级字典详情
 *   detail := system.SysDictionaryDetail{
 *       SysDictionaryID: 1,
 *       Label:           "启用",
 *       Value:           "1",
 *       Status:          &[]bool{true}[0],
 *       Sort:            1,
 *   }
 *   err := DictionaryDetailServiceApp.CreateSysDictionaryDetail(detail)
 *
 *   // 创建子级字典详情
 *   parentID := uint(1)
 *   childDetail := system.SysDictionaryDetail{
 *       SysDictionaryID: 1,
 *       ParentID:        &parentID,
 *       Label:           "子项",
 *       Value:           "1-1",
 *       Status:          &[]bool{true}[0],
 *       Sort:            1,
 *   }
 *   err := DictionaryDetailServiceApp.CreateSysDictionaryDetail(childDetail)
 *
 *   // 获取字典树形结构
 *   treeList, err := DictionaryDetailServiceApp.GetDictionaryTreeList(1)
 *
 *   // 根据值获取字典详情
 *   detail, err := DictionaryDetailServiceApp.GetDictionaryInfoByValue(1, "1")
 */

// DictionaryDetailService 字典详情服务结构体
// 提供字典详情相关的所有业务逻辑方法
type DictionaryDetailService struct{}

// DictionaryDetailServiceApp 字典详情服务实例
// 全局单例，用于在项目各处调用字典详情服务方法
var DictionaryDetailServiceApp = new(DictionaryDetailService)

// CreateSysDictionaryDetail 创建字典详情数据
//
// 功能说明：
//
//	创建一个新的字典详情记录，支持树形结构
//	如果指定了父级ID，会自动计算层级（Level）和路径（Path）
//	层级从0开始，每增加一级层级+1
//	路径格式：父级ID,祖父级ID,...（用逗号分隔）
//
// 参数：
//
//	sysDictionaryDetail: 要创建的字典详情对象，必须包含：
//	  - SysDictionaryID: 所属字典ID
//	  - Label: 显示标签
//	  - Value: 字典值
//	  - ParentID: 父级ID（可选，nil表示顶级项）
//	  - Status: 状态（可选）
//	  - Sort: 排序值（可选）
//
// 返回值：
//
//	err: 错误信息，如果父级不存在则返回错误
//
// 执行流程：
//  1. 如果指定了父级ID，查询父级信息
//  2. 计算当前项的层级：父级层级 + 1
//  3. 计算当前项的路径：父级路径 + 父级ID
//  4. 如果没有父级，设置为顶级（Level=0, Path=""）
//  5. 创建字典详情记录
//
// 使用示例：
//
//	// 创建顶级字典详情
//	detail := system.SysDictionaryDetail{
//	    SysDictionaryID: 1,
//	    Label:           "启用",
//	    Value:           "1",
//	    Status:          &[]bool{true}[0],
//	    Sort:            1,
//	    // ParentID 为 nil，表示顶级项
//	}
//	err := DictionaryDetailServiceApp.CreateSysDictionaryDetail(detail)
//	if err != nil {
//	    log.Printf("创建字典详情失败: %v", err)
//	    return
//	}
//	log.Println("字典详情创建成功")
//
//	// 创建子级字典详情
//	parentID := uint(1) // 父级ID
//	childDetail := system.SysDictionaryDetail{
//	    SysDictionaryID: 1,
//	    ParentID:        &parentID, // 指定父级
//	    Label:           "子项",
//	    Value:           "1-1",
//	    Status:          &[]bool{true}[0],
//	    Sort:            1,
//	}
//	err = DictionaryDetailServiceApp.CreateSysDictionaryDetail(childDetail)
//	if err != nil {
//	    log.Printf("创建子级字典详情失败: %v", err)
//	    return
//	}
//	log.Println("子级字典详情创建成功，Level=1, Path=1")
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

// DeleteSysDictionaryDetail 删除字典详情数据
//
// 功能说明：
//
//	删除指定的字典详情记录
//	删除前会进行安全检查，确保该字典详情下没有子项
//	如果有子项，则不允许删除，防止数据不一致
//
// 参数：
//
//	sysDictionaryDetail: 要删除的字典详情对象，必须包含：
//	  - ID: 字典详情ID
//
// 返回值：
//
//	err: 错误信息，如果存在子项则返回错误
//
// 安全检查：
//  1. 检查是否存在子项（parent_id = 当前ID）
//  2. 如果存在子项，返回错误："该字典详情下还有子项，无法删除"
//  3. 如果没有子项，执行删除操作
//
// 使用示例：
//
//	detail := system.SysDictionaryDetail{ID: 1}
//	err := DictionaryDetailServiceApp.DeleteSysDictionaryDetail(detail)
//	if err != nil {
//	    log.Printf("删除字典详情失败: %v", err)
//	    return
//	}
//	log.Println("字典详情删除成功")
//
//	// 如果存在子项，删除会失败
//	detailWithChildren := system.SysDictionaryDetail{ID: 1} // 假设ID=1有子项
//	err = DictionaryDetailServiceApp.DeleteSysDictionaryDetail(detailWithChildren)
//	if err != nil {
//	    log.Printf("删除失败: %v", err) // 输出: 该字典详情下还有子项，无法删除
//	}
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

// UpdateSysDictionaryDetail 更新字典详情数据
//
// 功能说明：
//
//	更新字典详情的所有字段信息
//	如果更新了父级ID，会自动重新计算层级和路径
//	同时会递归更新所有子项的层级和路径，确保数据一致性
//	更新前会检查循环引用，防止将项设置为自己或其子项的父级
//
// 参数：
//
//	sysDictionaryDetail: 要更新的字典详情对象指针，必须包含：
//	  - ID: 字典详情ID（必需）
//	  - 其他要更新的字段（可选）
//	  - ParentID: 新的父级ID（可选，如果更新则重新计算层级）
//
// 返回值：
//
//	err: 错误信息，如果存在循环引用或父级不存在则返回错误
//
// 执行流程：
//  1. 如果更新了父级ID，查询新的父级信息
//  2. 检查循环引用（防止将项设置为自己或其子项的父级）
//  3. 重新计算层级和路径
//  4. 更新字典详情记录
//  5. 递归更新所有子项的层级和路径
//
// 使用示例：
//
//	// 更新字典详情的基本信息（不改变父级）
//	detail := &system.SysDictionaryDetail{
//	    ID:              1,
//	    Label:           "更新后的标签",
//	    Value:           "updated_value",
//	    Status:          &[]bool{true}[0],
//	    Sort:            2,
//	    // ParentID 不设置，保持原有父级
//	}
//	err := DictionaryDetailServiceApp.UpdateSysDictionaryDetail(detail)
//	if err != nil {
//	    log.Printf("更新字典详情失败: %v", err)
//	    return
//	}
//	log.Println("字典详情更新成功")
//
//	// 更新父级（会重新计算层级和路径）
//	newParentID := uint(2)
//	detailWithNewParent := &system.SysDictionaryDetail{
//	    ID:       1,
//	    ParentID: &newParentID, // 设置新的父级
//	    Label:    "更新后的标签",
//	}
//	err = DictionaryDetailServiceApp.UpdateSysDictionaryDetail(detailWithNewParent)
//	if err != nil {
//	    log.Printf("更新失败: %v", err)
//	    return
//	}
//	log.Println("父级更新成功，层级和路径已重新计算")
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

// checkCircularReference 检查循环引用（私有方法）
//
// 功能说明：
//
//	递归检查是否存在循环引用
//	防止将字典详情设置为自己或其子项的父级
//	这是更新父级时的安全检查方法
//
// 参数：
//
//	id: 当前字典详情ID
//	parentID: 要设置的父级ID
//
// 返回值：
//
//	bool: true表示存在循环引用，false表示不存在
//
// 检查逻辑：
//  1. 如果 id == parentID，存在循环引用（自己作为自己的父级）
//  2. 查询父级信息
//  3. 如果父级还有父级，递归检查父级的父级是否等于id
//  4. 如果找到id在父级链中，说明存在循环引用
//
// 使用示例：
//
//	// 通常在 UpdateSysDictionaryDetail 中调用
//	// 检查将ID=1的项设置为ID=2的父级是否会导致循环引用
//	hasCircular := dictionaryDetailService.checkCircularReference(1, 2)
//	if hasCircular {
//	    log.Println("存在循环引用，不能设置")
//	}
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

// updateChildrenLevelAndPath 更新子项的层级和路径（私有方法）
//
// 功能说明：
//
//	递归更新指定父级下所有子项的层级和路径
//	当父级的层级或路径发生变化时，需要同步更新所有子项
//	确保整个树形结构的数据一致性
//
// 参数：
//
//	parentID: 父级字典详情ID
//
// 返回值：
//
//	error: 错误信息
//
// 执行流程：
//  1. 查询父级的所有直接子项
//  2. 查询父级信息，获取父级的层级和路径
//  3. 更新每个子项的层级：父级层级 + 1
//  4. 更新每个子项的路径：父级路径 + 父级ID
//  5. 递归更新每个子项的子项
//
// 使用示例：
//
//	// 通常在 UpdateSysDictionaryDetail 中调用
//	// 当父级的层级或路径变化时，更新所有子项
//	err := dictionaryDetailService.updateChildrenLevelAndPath(1)
//	if err != nil {
//	    log.Printf("更新子项层级和路径失败: %v", err)
//	    return
//	}
//	log.Println("所有子项的层级和路径已更新")
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

// GetSysDictionaryDetail 根据ID获取字典详情单条数据
//
// 功能说明：
//
//	根据字典详情ID查询单条记录
//	用于获取字典详情的完整信息
//
// 参数：
//
//	id: 字典详情ID
//
// 返回值：
//
//	sysDictionaryDetail: 查询到的字典详情对象
//	err: 错误信息，如果记录不存在则返回错误
//
// 使用示例：
//
//	detail, err := DictionaryDetailServiceApp.GetSysDictionaryDetail(1)
//	if err != nil {
//	    log.Printf("获取字典详情失败: %v", err)
//	    return
//	}
//	log.Printf("字典详情: %+v", detail)
//	log.Printf("标签: %s, 值: %s", detail.Label, detail.Value)
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

// loadChildren 递归加载子项（私有方法）
//
// 功能说明：
//
//	递归加载指定字典详情的所有子项
//	将子项填充到 detail.Children 字段中
//	同时设置每个子项的 disabled 属性
//	这是一个辅助方法，用于构建树形结构
//
// 参数：
//
//	detail: 字典详情对象指针，查询结果会填充到其 Children 字段中
//
// 返回值：
//
//	error: 错误信息
//
// 执行流程：
//  1. 查询当前项的所有直接子项（parent_id = detail.ID）
//  2. 设置每个子项的 disabled 属性（status为false时disabled为true）
//  3. 对每个子项递归调用此方法，继续加载子项的子项
//  4. 将子项列表赋值给 detail.Children
//
// 使用示例：
//
//	// 通常在 GetDictionaryTreeList 中调用
//	detail := &system.SysDictionaryDetail{ID: 1}
//	err := dictionaryDetailService.loadChildren(detail)
//	if err != nil {
//	    log.Printf("加载子项失败: %v", err)
//	    return
//	}
//	// detail.Children 现在包含所有子项及其子项
//	log.Printf("子项数量: %d", len(detail.Children))
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
//
// 功能说明：
//
//	根据父级ID获取字典详情列表
//	支持两种模式：
//	1. 获取指定父级下的直接子项（IncludeChildren = false）
//	2. 获取指定父级下的所有子项，包括子项的子项（IncludeChildren = true）
//	自动设置 disabled 属性
//
// 参数：
//
//	req: 查询请求对象，包含：
//	  - SysDictionaryID: 字典ID（必需）
//	  - ParentID: 父级ID（可选，nil表示查询顶级项）
//	  - IncludeChildren: 是否包含子项的子项（可选，默认false）
//
// 返回值：
//
//	list: 字典详情列表
//	err: 错误信息
//
// 查询模式说明：
//   - ParentID 为 nil：查询顶级项（parent_id IS NULL）
//   - ParentID 不为 nil：查询指定父级下的直接子项
//   - IncludeChildren = true：递归加载所有层级的子项
//   - IncludeChildren = false：只返回直接子项
//
// 使用示例：
//
//	// 获取字典ID为1的顶级项
//	req1 := request.GetDictionaryDetailsByParentRequest{
//	    SysDictionaryID: 1,
//	    ParentID:        nil, // nil表示顶级项
//	    IncludeChildren: false,
//	}
//	list1, err := DictionaryDetailServiceApp.GetDictionaryDetailsByParent(req1)
//	if err != nil {
//	    log.Printf("查询失败: %v", err)
//	    return
//	}
//	log.Printf("顶级项数量: %d", len(list1))
//
//	// 获取父级ID为1的所有子项（包括子项的子项）
//	parentID := uint(1)
//	req2 := request.GetDictionaryDetailsByParentRequest{
//	    SysDictionaryID: 1,
//	    ParentID:        &parentID,
//	    IncludeChildren: true, // 包含所有层级的子项
//	}
//	list2, err := DictionaryDetailServiceApp.GetDictionaryDetailsByParent(req2)
//	if err != nil {
//	    log.Printf("查询失败: %v", err)
//	    return
//	}
//	log.Printf("子项数量: %d（包含所有层级）", len(list2))
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

// GetDictionaryListByType 按照字典类型获取字典全部内容
//
// 功能说明：
//
//	根据字典类型（type）获取该类型下所有字典的字典详情记录
//	通过 JOIN 关联 sys_dictionaries 表查询
//	返回扁平化的列表，不包含树形结构
//
// 参数：
//
//	t: 字典类型（type字段值）
//
// 返回值：
//
//	list: 字典详情列表（扁平化，所有层级）
//	err: 错误信息
//
// 使用示例：
//
//	// 获取类型为"status"的所有字典详情
//	list, err := DictionaryDetailServiceApp.GetDictionaryListByType("status")
//	if err != nil {
//	    log.Printf("获取字典列表失败: %v", err)
//	    return
//	}
//	log.Printf("类型为'status'的字典详情共 %d 条", len(list))
//	for _, detail := range list {
//	    log.Printf("字典ID: %d, 标签: %s, 值: %s", detail.SysDictionaryID, detail.Label, detail.Value)
//	}
func (dictionaryDetailService *DictionaryDetailService) GetDictionaryListByType(t string) (list []system.SysDictionaryDetail, err error) {
	var sysDictionaryDetails []system.SysDictionaryDetail
	db := global.GVA_DB.Model(&system.SysDictionaryDetail{}).Joins("JOIN sys_dictionaries ON sys_dictionaries.id = sys_dictionary_details.sys_dictionary_id")
	err = db.Find(&sysDictionaryDetails, "type = ?", t).Error
	return sysDictionaryDetails, err
}

// GetDictionaryTreeListByType 根据字典类型获取树形结构
//
// 功能说明：
//
//	根据字典类型（type）获取该类型下所有字典的树形结构列表
//	通过 JOIN 关联 sys_dictionaries 表查询
//	只返回顶级项（parent_id为NULL），每个顶级项包含其所有子项
//	自动设置 disabled 属性
//
// 参数：
//
//	t: 字典类型（type字段值）
//
// 返回值：
//
//	list: 字典详情树形结构列表（只包含顶级项及其子项）
//	err: 错误信息
//
// 树形结构说明：
//   - 只返回顶级项（parent_id IS NULL）
//   - 每个项通过 Children 字段包含其子项
//   - 子项也包含其子项，形成递归树形结构
//   - 所有项按 sort 字段排序
//
// 使用示例：
//
//	// 获取类型为"status"的树形结构
//	treeList, err := DictionaryDetailServiceApp.GetDictionaryTreeListByType("status")
//	if err != nil {
//	    log.Printf("获取字典树失败: %v", err)
//	    return
//	}
//
//	// 遍历树形结构
//	for _, root := range treeList {
//	    log.Printf("顶级项: %s (字典ID: %d)", root.Label, root.SysDictionaryID)
//	    // 遍历子项
//	    for _, child := range root.Children {
//	        log.Printf("  子项: %s", child.Label)
//	    }
//	}
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

// GetDictionaryInfoByValue 按照字典ID和值获取单条字典内容
//
// 功能说明：
//
//	根据字典ID和字典值（value）精确匹配查询单条字典详情记录
//	适用于通过值查找字典详情的场景
//
// 参数：
//
//	dictionaryID: 字典ID
//	value: 字典值（value字段值）
//
// 返回值：
//
//	detail: 查询到的字典详情对象
//	err: 错误信息，如果记录不存在则返回错误
//
// 使用示例：
//
//	// 获取字典ID为1，值为"1"的字典详情
//	detail, err := DictionaryDetailServiceApp.GetDictionaryInfoByValue(1, "1")
//	if err != nil {
//	    log.Printf("获取字典详情失败: %v", err)
//	    return
//	}
//	log.Printf("找到字典详情: ID=%d, 标签=%s, 值=%s", detail.ID, detail.Label, detail.Value)
//
//	// 实际应用场景：根据字典值获取显示标签
//	statusDetail, err := DictionaryDetailServiceApp.GetDictionaryInfoByValue(1, "1")
//	if err == nil {
//	    log.Printf("状态值'1'对应的标签是: %s", statusDetail.Label) // 例如：启用
//	}
func (dictionaryDetailService *DictionaryDetailService) GetDictionaryInfoByValue(dictionaryID uint, value string) (detail system.SysDictionaryDetail, err error) {
	var sysDictionaryDetail system.SysDictionaryDetail
	err = global.GVA_DB.First(&sysDictionaryDetail, "sys_dictionary_id = ? and value = ?", dictionaryID, value).Error
	return sysDictionaryDetail, err
}

// GetDictionaryInfoByTypeValue 按照字典类型和值获取单条字典内容
//
// 功能说明：
//
//	根据字典类型（type）和字典值（value）精确匹配查询单条字典详情记录
//	通过 JOIN 关联 sys_dictionaries 表查询
//	适用于只知道字典类型和值，不知道字典ID的场景
//
// 参数：
//
//	t: 字典类型（type字段值）
//	value: 字典值（value字段值）
//
// 返回值：
//
//	detail: 查询到的字典详情对象
//	err: 错误信息，如果记录不存在则返回错误
//
// 使用示例：
//
//	// 获取类型为"status"，值为"1"的字典详情
//	detail, err := DictionaryDetailServiceApp.GetDictionaryInfoByTypeValue("status", "1")
//	if err != nil {
//	    log.Printf("获取字典详情失败: %v", err)
//	    return
//	}
//	log.Printf("找到字典详情: ID=%d, 标签=%s, 值=%s", detail.ID, detail.Label, detail.Value)
//
//	// 实际应用场景：根据类型和值获取显示标签
//	statusDetail, err := DictionaryDetailServiceApp.GetDictionaryInfoByTypeValue("status", "1")
//	if err == nil {
//	    log.Printf("状态值'1'对应的标签是: %s", statusDetail.Label) // 例如：启用
//	}
func (dictionaryDetailService *DictionaryDetailService) GetDictionaryInfoByTypeValue(t string, value string) (detail system.SysDictionaryDetail, err error) {
	var sysDictionaryDetails system.SysDictionaryDetail
	db := global.GVA_DB.Model(&system.SysDictionaryDetail{}).Joins("JOIN sys_dictionaries ON sys_dictionaries.id = sys_dictionary_details.sys_dictionary_id")
	err = db.First(&sysDictionaryDetails, "sys_dictionaries.type = ? and sys_dictionary_details.value = ?", t, value).Error
	return sysDictionaryDetails, err
}

// GetDictionaryPath 获取字典详情的完整路径
//
// 功能说明：
//
//	根据字典详情ID获取从根节点到当前节点的完整路径
//	返回路径数组，第一个元素是根节点，最后一个元素是当前节点
//	通过递归查询父级节点构建完整路径
//
// 参数：
//
//	id: 字典详情ID
//
// 返回值：
//
//	path: 路径数组，从根节点到当前节点的所有字典详情对象
//	err: 错误信息，如果记录不存在则返回错误
//
// 路径说明：
//   - 路径数组的第一个元素是根节点（顶级项）
//   - 路径数组的最后一个元素是当前节点
//   - 中间元素是按层级顺序的父级节点
//   - 例如：[根节点, 父级1, 父级2, 当前节点]
//
// 使用示例：
//
//	// 获取ID为5的字典详情的完整路径
//	path, err := DictionaryDetailServiceApp.GetDictionaryPath(5)
//	if err != nil {
//	    log.Printf("获取路径失败: %v", err)
//	    return
//	}
//
//	// 打印完整路径
//	log.Println("完整路径:")
//	for i, detail := range path {
//	    if i > 0 {
//	        log.Printf(" -> ")
//	    }
//	    log.Printf("%s", detail.Label)
//	}
//	// 输出示例: 根节点 -> 父级1 -> 父级2 -> 当前节点
//
//	// 获取路径的层级信息
//	log.Printf("路径深度: %d", len(path))
//	log.Printf("根节点: %s", path[0].Label)
//	log.Printf("当前节点: %s", path[len(path)-1].Label)
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
//
// 功能说明：
//
//	根据字典ID和字典值获取从根节点到当前节点的完整路径
//	先通过字典ID和值查找字典详情，然后获取其完整路径
//	这是 GetDictionaryInfoByValue 和 GetDictionaryPath 的组合方法
//
// 参数：
//
//	dictionaryID: 字典ID
//	value: 字典值（value字段值）
//
// 返回值：
//
//	path: 路径数组，从根节点到当前节点的所有字典详情对象
//	err: 错误信息，如果记录不存在则返回错误
//
// 执行流程：
//  1. 根据字典ID和值查找字典详情
//  2. 如果找到，调用 GetDictionaryPath 获取完整路径
//  3. 返回路径数组
//
// 使用示例：
//
//	// 获取字典ID为1，值为"1-1-1"的字典详情的完整路径
//	path, err := DictionaryDetailServiceApp.GetDictionaryPathByValue(1, "1-1-1")
//	if err != nil {
//	    log.Printf("获取路径失败: %v", err)
//	    return
//	}
//
//	// 打印完整路径
//	log.Println("完整路径:")
//	for i, detail := range path {
//	    if i > 0 {
//	        log.Printf(" -> ")
//	    }
//	    log.Printf("%s (值: %s)", detail.Label, detail.Value)
//	}
//	// 输出示例: 根节点 (值: 1) -> 父级1 (值: 1-1) -> 当前节点 (值: 1-1-1)
//
//	// 实际应用场景：显示面包屑导航
//	breadcrumb := make([]string, len(path))
//	for i, detail := range path {
//	    breadcrumb[i] = detail.Label
//	}
//	log.Printf("面包屑: %v", breadcrumb) // 例如: [根节点, 父级1, 当前节点]
func (dictionaryDetailService *DictionaryDetailService) GetDictionaryPathByValue(dictionaryID uint, value string) (path []system.SysDictionaryDetail, err error) {
	detail, err := dictionaryDetailService.GetDictionaryInfoByValue(dictionaryID, value)
	if err != nil {
		return nil, err
	}

	return dictionaryDetailService.GetDictionaryPath(detail.ID)
}
