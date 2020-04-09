package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateApi
// @description   create base apis, 新增基础api
// @auth                     （2020/04/05  20:22 ）
// @param     FileMd5         string
// @param     FileName        string
// @param     FilePath        string
// @return                    error
func CreateApi(a model.SysApi) (err error) {
	findOne := global.GVA_DB.Where("path = ? AND method = ?", a.Path, a.Method).Find(&model.SysApi{}).Error
	if findOne == nil {
		return errors.New("存在相同api")
	} else {
		err = global.GVA_DB.Create(&a).Error
	}
	return err
}

// @title    DeleteApi
// @description   delete base apis, 删除基础api
// @auth                     （2020/04/05  20:22 ）
// @return                    error
func DeleteApi(a model.SysApi) (err error) {
	err = global.GVA_DB.Delete(a).Error
	ClearCasbin(1, a.Path,a.Method)
	return err
}

// @title    GetInfoList
// @description   get apis by pagination, 分页获取数据
// @auth                     （2020/04/05  20:22 ）
// @param     info            PageInfo
// @return    err             error
// @return    list            interface{}
// @return    total           int
func GetAPIInfoList(a model.SysApi, info request.PageInfo, Order string, Desc bool) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB

	if err != nil {
		return
	} else {
		var apiList []model.SysApi

		if a.Path != "" {
			db = db.Where("path LIKE ?", "%"+a.Path+"%")
		}

		if a.Description != "" {
			db = db.Where("description LIKE ?", "%"+a.Description+"%")
		}

		if a.Method != "" {
			db = db.Where("method = ?", a.Method)
		}

		err = db.Find(&apiList).Count(&total).Error

		if err != nil {
			return err, apiList, total
		} else {
			db = db.Limit(limit).Offset(offset)
			if Order != "" {
				var OrderStr string
				if Desc {
					OrderStr = Order + " desc"
				} else {
					OrderStr = Order
				}
				err = db.Order(OrderStr, true).Find(&apiList).Error
			} else {
				err = db.Order("api_group", true).Find(&apiList).Error
			}
		}
		return err, apiList, total
	}
}

// @title    GetAllApis
// @description   get all apis, 获取所有的api
// @auth                     （2020/04/05  20:22 ）
// @return       err          error
// @return       apis         []SysApi
func GetAllApis() (err error, apis []model.SysApi) {
	err = global.GVA_DB.Find(&apis).Error
	return
}

// @title    GetApiById
// @description   根据id获取api
// @auth                     （2020/04/05  20:22 ）
// @param     id              float64
// @return                    error
func GetApiById(id float64) (err error, api model.SysApi) {
	err = global.GVA_DB.Where("id = ?", id).First(&api).Error
	return
}

// @title    UpdateApi
// @description   update a base api, update api
// @auth                     （2020/04/05  20:22 ）
// @return                    error
func UpdateApi(a model.SysApi) (err error) {
	var oldA model.SysApi

	err = global.GVA_DB.Where("id = ?", a.ID).First(&oldA).Error

	if oldA.Path != a.Path || oldA.Method != a.Method{
		flag := global.GVA_DB.Where("path = ? AND method = ?", a.Path, a.Method).Find(&model.SysApi{}).RecordNotFound()
		if !flag {
			return errors.New("存在相同api路径")
		}
	}
	if err != nil {
		return err
	} else {
		err = UpdateCasbinApi(oldA.Path, a.Path,oldA.Method,a.Method)
		if err != nil {
			return err
		} else {
			err = global.GVA_DB.Save(a).Error
		}
	}
	return err
}
