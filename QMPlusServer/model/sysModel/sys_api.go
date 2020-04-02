package sysModel

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type SysApi struct {
	gorm.Model
	Path        string `json:"path"`
	Description string `json:"description"`
	Group       string `json:"group"`
	Method      string `json:"method" gorm:"default:'POST'"`
}

//新增基础api
func (a *SysApi) CreateApi() (err error) {
	findOne := qmsql.DEFAULTDB.Where("path = ?", a.Path).Find(&SysMenu{}).Error
	if findOne == nil {
		return errors.New("存在相同api")
	} else {
		err = qmsql.DEFAULTDB.Create(a).Error
	}
	return err
}

//删除基础api
func (a *SysApi) DeleteApi() (err error) {
	err = qmsql.DEFAULTDB.Delete(a).Error
	new(CasbinModel).clearCasbin(1, a.Path)
	return err
}

//更新api
func (a *SysApi) UpdataApi() (err error) {
	var oldA SysApi
	flag := qmsql.DEFAULTDB.Where("path = ?", a.Path).RecordNotFound()
	if !flag {
		return errors.New("存在相同api路径")
	}
	err = qmsql.DEFAULTDB.Where("id = ?", a.ID).First(&oldA).Error
	if err != nil {
		return err
	} else {
		err = new(CasbinModel).CasbinApiUpdata(oldA.Path, a.Path)
		if err != nil {
			return err
		} else {
			err = qmsql.DEFAULTDB.Save(a).Error
		}
	}
	return err
}

//获取选中角色所拥有的api
func (a *SysApi) GetApiById(id float64) (err error, api SysApi) {
	err = qmsql.DEFAULTDB.Where("id = ?", id).First(&api).Error
	return
}

// 获取所有api信息
func (a *SysApi) GetAllApis() (err error, apis []SysApi) {
	err = qmsql.DEFAULTDB.Find(&apis).Error
	return
}

// 分页获取数据  需要分页实现这个接口即可
func (a *SysApi) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(a, info)
	if err != nil {
		return
	} else {
		var apiList []SysApi
		model := qmsql.DEFAULTDB.Model(info)
		if a.Path != "" {
			model = model.Where("path LIKE ?", "%"+a.Path+"%")
			db = db.Where("path LIKE ?", "%"+a.Path+"%")
		}

		if a.Description != "" {
			model = model.Where("description LIKE ?", "%"+a.Description+"%")
			db = db.Where("description LIKE ?", "%"+a.Description+"%")
		}

		if a.Method != "" {
			model = model.Where("method = ?", a.Method)
			db = db.Where("method = ?", a.Method)
		}
		err = model.Find(&apiList).Count(&total).Error
		if err != nil {
			return err, apiList, total
		} else {
			err = db.Order("group", true).Find(&apiList).Error
		}
		return err, apiList, total
	}
}
