package model

import (
	"gin-vue-admin/global"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type SysApi struct {
	gorm.Model
	Path        string `json:"path"`
	Description string `json:"description"`
	ApiGroup    string `json:"apiGroup"`
	Method      string `json:"method" gorm:"default:'POST'"`
}

type CreateApiParams struct {
	Path        string `json:"path"`
	Description string `json:"description"`
}

type DeleteApiParams struct {
	ID uint `json:"id"`
}

// @title    CreateApi
// @description   create base apis, 新增基础api
// @auth                     （2020/04/05  20:22 ）
// @param     FileMd5         string
// @param     FileName        string
// @param     FilePath        string
// @return                    error
func (a *SysApi) CreateApi() (err error) {
	findOne := global.GVA_DB.Where("path = ?", a.Path).Find(&SysApi{}).Error
	if findOne == nil {
		return errors.New("存在相同api")
	} else {
		err = global.GVA_DB.Create(a).Error
	}
	return err
}

// @title    DeleteApi
// @description   delete base apis, 删除基础api
// @auth                     （2020/04/05  20:22 ）
// @return                    error
func (a *SysApi) DeleteApi() (err error) {
	err = global.GVA_DB.Delete(a).Error
	new(CasbinModel).ClearCasbin(1, a.Path)
	return err
}

// @title    UpdateApi
// @description   update a base api, update api
// @auth                     （2020/04/05  20:22 ）
// @return                    error
func (a *SysApi) UpdateApi() (err error) {
	var oldA SysApi
	flag := global.GVA_DB.Where("path = ?", a.Path).Find(&SysApi{}).RecordNotFound()
	if !flag {
		return errors.New("存在相同api路径")
	}
	err = global.GVA_DB.Where("id = ?", a.ID).First(&oldA).Error
	if err != nil {
		return err
	} else {
		err = new(CasbinModel).UpdateCasbinApi(oldA.Path, a.Path)
		if err != nil {
			return err
		} else {
			err = global.GVA_DB.Save(a).Error
		}
	}
	return err
}

// @title    GetApiById
// @description   get the apis of the selected user, 获取选中角色所拥有的api
// @auth                     （2020/04/05  20:22 ）
// @param     id              float64
// @return                    error
func (a *SysApi) GetApiById(id float64) (err error, api SysApi) {
	err = global.GVA_DB.Where("id = ?", id).First(&api).Error
	return
}

// @title    GetAllApis
// @description   get all apis, 获取所有的api
// @auth                     （2020/04/05  20:22 ）
// @return       err          error
// @return       apis         []SysApi
func (a *SysApi) GetAllApis() (err error, apis []SysApi) {
	err = global.GVA_DB.Find(&apis).Error
	return
}

// @title    GetInfoList
// @description   get apis by pagination, 分页获取数据
// @auth                     （2020/04/05  20:22 ）
// @param     info            PageInfo
// @return    err             error
// @return    list            interface{}
// @return    total           int
func (a *SysApi) GetInfoList(info PageInfo, Order string, Desc bool) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB

	if err != nil {
		return
	} else {
		var apiList []SysApi

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
				err = db.Order(Order+" desc", true).Find(&apiList).Error
			} else {
				err = db.Order("api_group", true).Find(&apiList).Error
			}
		}
		return err, apiList, total
	}
}
