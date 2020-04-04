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

//新增基础api
func (a *SysApi) CreateApi() (err error) {
	findOne := global.GVA_DB.Where("path = ?", a.Path).Find(&SysApi{}).Error
	if findOne == nil {
		return errors.New("存在相同api")
	} else {
		err = global.GVA_DB.Create(a).Error
	}
	return err
}

//删除基础api
func (a *SysApi) DeleteApi() (err error) {
	err = global.GVA_DB.Delete(a).Error
	new(CasbinModel).clearCasbin(1, a.Path)
	return err
}

//更新api
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
		err = new(CasbinModel).CasbinApiUpdate(oldA.Path, a.Path)
		if err != nil {
			return err
		} else {
			err = global.GVA_DB.Save(a).Error
		}
	}
	return err
}

//获取选中角色所拥有的api
func (a *SysApi) GetApiById(id float64) (err error, api SysApi) {
	err = global.GVA_DB.Where("id = ?", id).First(&api).Error
	return
}

// 获取所有api信息
func (a *SysApi) GetAllApis() (err error, apis []SysApi) {
	err = global.GVA_DB.Find(&apis).Error
	return
}

// 分页获取数据
func (a *SysApi) GetInfoList(info PageInfo) (err error, list interface{}, total int) {
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
			err = db.Limit(limit).Offset(offset).Order("api_group", true).Find(&apiList).Error
		}
		return err, apiList, total
	}
}
