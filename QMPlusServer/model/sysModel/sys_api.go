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
}

func (a *SysApi) CreateApi() (err error) {
	findOne := qmsql.DEFAULTDB.Where("path = ?", a.Path).Find(&SysMenu{}).Error
	if findOne == nil {
		return errors.New("存在相同api")
	} else {
		err = qmsql.DEFAULTDB.Create(a).Error
	}
	return err
}

func (a *SysApi) DeleteApi() (err error) {
	err = qmsql.DEFAULTDB.Delete(a).Error
	new(CasbinModel).clearCasbin(1, a.Path)
	return err
}

func (a *SysApi) UpdataApi() (err error) {
	var oldA SysApi
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
		err = db.Order("group", true).Where("path LIKE ?", "%"+a.Path+"%").Find(&apiList).Error
		return err, apiList, total
	}
}
