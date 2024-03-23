package webcms

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/webcms"
	webcmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/webcms/request"
)

type ClassService struct {
}

// CreateClass 创建Class记录
// Author [piexlmax](https://github.com/piexlmax)
func (classService *ClassService) CreateClass(class webcms.Class) (err error) {
	err = global.GVA_DB.Create(&class).Error
	return err
}

// DeleteClass 删除Class记录
// Author [piexlmax](https://github.com/piexlmax)
func (classService *ClassService) DeleteClass(class webcms.Class) (err error) {
	err = global.GVA_DB.Delete(&class).Error
	return err
}

// DeleteClassByIds 批量删除Class记录
// Author [piexlmax](https://github.com/piexlmax)
func (classService *ClassService) DeleteClassByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]webcms.Class{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateClass 更新Class记录
// Author [piexlmax](https://github.com/piexlmax)
func (classService *ClassService) UpdateClass(class webcms.Class) (err error) {
	err = global.GVA_DB.Save(&class).Error
	return err
}

// GetClass 根据id获取Class记录
// Author [piexlmax](https://github.com/piexlmax)
func (classService *ClassService) GetClass(id uint) (class webcms.Class, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&class).Error
	return
}

// GetClass 根据catid获取Class记录
func (classService *ClassService) GetClassByCatid(catid uint) (class []webcms.Class, err error) {
	err = global.GVA_DB.Where("cate_id = ?", catid).Find(&class).Error
	return
}

// GetClassInfoList 分页获取Class记录
// Author [piexlmax](https://github.com/piexlmax)
func (classService *ClassService) GetClassInfoList(info webcmsReq.ClassSearch, siteid any) (list []webcms.Class, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&webcms.Class{}).Where("siteid", siteid)
	var classs []webcms.Class
	db.Preload("CateMenus")
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	// if info.CateId != "0" {
	// 	db.Where("cate_id = ?", info.CateId)
	// }
	// 判断是否有子栏目
	var catids []uint
	global.GVA_DB.Model(&webcms.CateMenus{}).Where("parent_id = ?", info.CateId).Select("id").Find(&catids)

	if len(catids) > 0 {
		db.Where("cate_id in ?", catids)
	} else {
		db.Where("cate_id = ?", info.CateId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&classs).Error
	return classs, total, err
}

func (classService *ClassService) GetClassInfoListByCatId(info request.PageInfo, catid string) (list []webcms.Class, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&webcms.Class{})
	db.Preload("CateMenus").Where("enable", 1)
	// 判断是否有子栏目
	var catids []uint
	global.GVA_DB.Model(&webcms.CateMenus{}).Where("parent_id = ?", catid).Select("id").Find(&catids)

	if len(catids) > 0 {
		db.Where("cate_id in ?", catids)
	} else {
		db.Where("cate_id = ?", catid)
	}
	var classs []webcms.Class
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Order(info.OrderType).Find(&classs).Error
	for k, v := range classs {
		classs[k].Url = fmt.Sprint("/show/", v.CateId, "/", v.ID)
	}
	return classs, total, err
}

// 通过课程id获取Class列表
func (classService *ClassService) GetClassListById(course_id string) (list []webcms.Class, err error) {
	// 创建db
	db := global.GVA_DB.Model(&webcms.Class{})
	db.Select("id,title")
	var classs []webcms.Class
	err = db.Where("course_id", course_id).Find(&classs).Error
	return classs, err
}
