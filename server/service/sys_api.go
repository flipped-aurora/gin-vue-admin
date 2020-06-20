package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateApi
// @description   create base apis, 新增基础api
// @auth                     （2020/04/05  20:22）
// @param     api             model.SysApi
// @return                    error

func CreateApi(api model.SysApi) (err error) {
	findOne := global.GVA_DB.Where("path = ? AND method = ?", api.Path, api.Method).Find(&model.SysApi{}).Error
	if findOne == nil {
		return errors.New("存在相同api")
	} else {
		err = global.GVA_DB.Create(&api).Error
	}
	return err
}

// @title    DeleteApi
// @description   delete a base api, 删除基础api
// @param     api             model.SysApi
// @auth                     （2020/04/05  20:22）
// @return                    error

func DeleteApi(api model.SysApi) (err error) {
	err = global.GVA_DB.Delete(api).Error
	ClearCasbin(1, api.Path, api.Method)
	return err
}

// @title    GetInfoList
// @description   get apis by pagination, 分页获取数据
// @auth                     （2020/04/05  20:22）
// @param     api             model.SysApi
// @param     info            request.PageInfo
// @param     order           string
// @param     desc            bool
// @return    err             error
// @return    list            interface{}
// @return    total           int

func GetAPIInfoList(api model.SysApi, info request.PageInfo, order string, desc bool) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.SysApi{})
	var apiList []model.SysApi

	if api.Path != "" {
		db = db.Where("path LIKE ?", "%"+api.Path+"%")
	}

	if api.Description != "" {
		db = db.Where("description LIKE ?", "%"+api.Description+"%")
	}

	if api.Method != "" {
		db = db.Where("method = ?", api.Method)
	}

	if api.ApiGroup != "" {
		db = db.Where("api_group = ?", api.ApiGroup)
	}

	err = db.Count(&total).Error

	if err != nil {
		return err, apiList, total
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			if desc {
				OrderStr = order + " desc"
			} else {
				OrderStr = order
			}
			err = db.Order(OrderStr, true).Find(&apiList).Error
		} else {
			err = db.Order("api_group", true).Find(&apiList).Error
		}
	}
	return err, apiList, total
}

// @title    GetAllApis
// @description   get all apis, 获取所有的api
// @auth                     （2020/04/05  20:22）
// @return       err          error
// @return       apis         []SysApi

func GetAllApis() (err error, apis []model.SysApi) {
	err = global.GVA_DB.Find(&apis).Error
	return
}

// @title    GetApiById
// @description   根据id获取api
// @auth                     （2020/04/05  20:22）
// @param     api             model.SysApi
// @param     id              float64
// @return                    error

func GetApiById(id float64) (err error, api model.SysApi) {
	err = global.GVA_DB.Where("id = ?", id).First(&api).Error
	return
}

// @title    UpdateApi
// @description   update a base api, update api
// @auth                     （2020/04/05  20:22）
// @param     api             model.SysApi
// @return                    error

func UpdateApi(api model.SysApi) (err error) {
	var oldA model.SysApi

	err = global.GVA_DB.Where("id = ?", api.ID).First(&oldA).Error

	if oldA.Path != api.Path || oldA.Method != api.Method {
		flag := global.GVA_DB.Where("path = ? AND method = ?", api.Path, api.Method).Find(&model.SysApi{}).RecordNotFound()
		if !flag {
			return errors.New("存在相同api路径")
		}
	}
	if err != nil {
		return err
	} else {
		err = UpdateCasbinApi(oldA.Path, api.Path, oldA.Method, api.Method)
		if err != nil {
			return err
		} else {
			err = global.GVA_DB.Save(&api).Error
		}
	}
	return err
}
