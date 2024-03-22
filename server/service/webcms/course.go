package webcms

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/webcms"
	webcmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/webcms/request"
	"github.com/gofrs/uuid"
)

type CourseService struct {
}

// CreateCourse 创建Course记录
// Author [piexlmax](https://github.com/piexlmax)
func (courseService *CourseService) CreateCourse(course webcms.Course) (err error) {
	err = global.GVA_DB.Create(&course).Error
	return err
}

// DeleteCourse 删除Course记录
// Author [piexlmax](https://github.com/piexlmax)
func (courseService *CourseService) DeleteCourse(course webcms.Course) (err error) {
	err = global.GVA_DB.Delete(&course).Error
	return err
}

// DeleteCourseByIds 批量删除Course记录
// Author [piexlmax](https://github.com/piexlmax)
func (courseService *CourseService) DeleteCourseByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]webcms.Course{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCourse 更新Course记录
// Author [piexlmax](https://github.com/piexlmax)
func (courseService *CourseService) UpdateCourse(course webcms.Course) (err error) {
	global.GVA_DB.Model(&course).Association("CreatedBy").Clear()
	err = global.GVA_DB.Save(&course).Error
	return err
}

// GetCourse 根据id获取Course记录
// Author [piexlmax](https://github.com/piexlmax)
func (courseService *CourseService) GetCourse(id uint) (course webcms.Course, err error) {
	err = global.GVA_DB.Preload("SysUser").Where("id = ?", id).First(&course).Error
	return
}

// GetCourseByCatid 根据catid获取Course记录
// Author [piexlmax](https://github.com/piexlmax)
func (courseService *CourseService) GetCourseByCatid(catid uint) (course []webcms.Course, err error) {
	err = global.GVA_DB.Where("cate_id = ?", catid).Where("enable", 1).Find(&course).Error
	return
}

func (courseService *CourseService) GetAllInfoByCatid(catid uint) (course []webcms.Course, err error) {
	var catids []uint
	global.GVA_DB.Model(&webcms.CateMenus{}).Where("parent_id = ?", catid).Select("id").Find(&catids)
	err = global.GVA_DB.Where("cate_id in ?", catids).Where("enable", 1).Order("created_at desc").Limit(8).Find(&course).Error
	return
}

// GetCourseInfoList 分页获取Course记录
// Author [piexlmax](https://github.com/piexlmax)
func (courseService *CourseService) GetCourseInfoList(info webcmsReq.CourseSearch, AuthorityId uint, uuid uuid.UUID) (list []webcms.Course, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&webcms.Course{})
	db.Preload("SysUser")
	// 权限slice
	authorityIdSlice := []uint{9528, 888}
	if !InSlice(authorityIdSlice, AuthorityId) {
		db.Where("created_by", uuid)
	}
	var courses []webcms.Course
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Title != "" {
		db = db.Where("title like ?", "%"+info.Title+"%")
	}
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Order("sort DESC,created_at DESC").Limit(limit).Offset(offset).Find(&courses).Error
	return courses, total, err
}

// 获取内容
func (courseService *CourseService) GetCourseListByCatId(info request.PageInfo, catid string) (list []webcms.Course, total int64, err error) {

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&webcms.Course{}).Where("enable", 1)

	// 判断是否有子栏目
	var catids []uint
	global.GVA_DB.Model(&webcms.CateMenus{}).Where("parent_id = ?", catid).Select("id").Find(&catids)

	if len(catids) > 0 {
		db.Where("cate_id in ?", catids)
	} else {
		db.Where("cate_id = ?", catid)
	}
	var courses []webcms.Course
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Order(info.OrderType).Find(&courses).Error
	for k, v := range courses {
		courses[k].Url = fmt.Sprint("/show/", v.CateId, "/", v.ID)
	}
	return courses, total, err
}

// 获取上一页  下一页
func (courseService *CourseService) GetCourseNextPreviousById(catID, showID, ordertype int) (next, previous webcms.Class) {
	db1 := global.GVA_DB.Model(&webcms.Course{}).Select("id,title").Where("enable", 1).Where("cate_id = ?", catID)
	db2 := global.GVA_DB.Model(&webcms.Course{}).Select("id,title").Where("enable", 1).Where("cate_id = ?", catID)
	// 升序
	if ordertype == 1 {
		// 查询下一个ID
		db1.Where("id > ?", showID).Order("id ASC").First(&next)
		// 查询上一个ID
		db2.Where("id < ?", showID).Order("id ASC").First(&previous)
	} else {
		// 查询上一个ID
		db1.Where("id < ?", showID).Order("id DESC").First(&previous)
		// 查询下一个ID
		db2.Where("id > ?", showID).Order("id DESC").First(&next)
	}
	return
}
