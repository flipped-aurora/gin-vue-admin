package sysModel

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"github.com/jinzhu/gorm"
)

type ChCity struct {
	gorm.Model
	Province string
	City     string
	Code     string
}

func (c *ChCity) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(c, info)
	if err != nil {
		return
	} else {
		var cityList []ChCity
		err = db.Find(&cityList).Error
		return err, cityList, total
	}
}

func (c *ChCity) GetCityByProvince(province string) (cityList []ChCity, err error) {
	err = qmsql.DEFAULTDB.Where("province = ?", province).Find(&cityList).Error
	return
}
