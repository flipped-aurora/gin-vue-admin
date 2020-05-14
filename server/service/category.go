package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateCategory
// @description   create a Category
// @param     cate               model.Category
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateCategory(cate model.Category) (err error) {
	err = global.GVA_DB.Create(&cate).Error
	return err
}

// @title    DeleteCategory
// @description   delete a Category
// @auth                     （2020/04/05  20:22）
// @param     cate               model.Category
// @return                    error

func DeleteCategory(cate model.Category) (err error) {
	err = global.GVA_DB.Delete(cate).Error
	return err
}

// @title    UpdateCategory
// @description   update a Category
// @param     cate          *model.Category
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateCategory(cate *model.Category) (err error) {
	err = global.GVA_DB.Save(cate).Error
	return err
}

// @title    GetCategory
// @description   get the info of a Category
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Category        Category

func GetCategory(id uint) (err error, cate model.Category) {
	err = global.GVA_DB.Where("id = ?", id).First(&cate).Error
	return
}

// @title    GetCategoryInfoList
// @description   get Category list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetCategoryInfoList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
    	offset := info.PageSize * (info.Page - 1)
    	db := global.GVA_DB
    	var cates []model.Category
    	err = db.Find(&cates).Count(&total).Error
    	err = db.Limit(limit).Offset(offset).Find(&cates).Error
    	return err,cates, total
}
