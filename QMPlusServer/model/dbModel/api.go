package dbModel

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"main/controller/servers"
	"main/init/qmsql"
	"main/model/modelInterface"
)

type Api struct {
	gorm.Model
	Path        string `json:"path"`
	Description string `json:"description"`
}

func (a *Api) CreateApi() (err error) {
	findOne := qmsql.DEFAULTDB.Where("path = ?", a.Path).Find(&Menu{}).Error
	if findOne == nil {
		return errors.New("存在相同api")
	} else {
		err = qmsql.DEFAULTDB.Create(a).Error
	}
	return err
}

func (a *Api) DeleteApi() (err error) {
	err = qmsql.DEFAULTDB.Where("path = ?", a.Path).Delete(a).Delete(&ApiAuthority{}).Error
	return err
}

func (a *Api) UpdataApi() (err error) {
	err = qmsql.DEFAULTDB.Save(a).Error
	return err
}

func (a *Api) GetApiById(id float64)(err error,api Api){
	err = qmsql.DEFAULTDB.Where("id = ?",id).First(&api).Error
	return
}

// 分页获取数据  需要分页实现这个接口即可
func (a *Api) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(a, info)
	if err != nil {
		return
	} else {
		var apiList []Api
		err = db.Find(&apiList).Error
		return err, apiList, total
	}
}
