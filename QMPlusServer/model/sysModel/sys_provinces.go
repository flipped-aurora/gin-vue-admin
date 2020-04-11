package sysModel

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"github.com/jinzhu/gorm"
)

type ChProvinces struct {
	gorm.Model
	Province string `json:"province"`
}

func (c *ChProvinces) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(c, info)
	if err != nil {
		return
	} else {
		var provList []ChProvinces
		err = db.Find(&provList).Error
		return err, provList, total
	}
}
