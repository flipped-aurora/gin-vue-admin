package customerModel

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Address struct {
	gorm.Model
	UserId      uuid.UUID `json:"userId"`
	Province    string    `json:"province"`
	City        string    `json:"city"`
	Town        string    `json:"town"`
	SpecAddress string    `json:"specAddress"`
}

func (a *Address) GetInfoList(info modelInterface.PageInfo) (error, interface{}, int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(a, info)
	if err != nil {
		return err, nil, 0
	} else {
		var addressList []Address
		err = db.Find(&addressList).Error
		return err, addressList, total
	}
}

func (a *Address) AddAddress() (err error) {
	err = qmsql.DEFAULTDB.Create(&a).Error
	return
}

func (a *Address) UpdateAddress() (err error) {
	updataMap := make(map[string]interface{})
	updataMap["province"] = a.Province
	updataMap["city"] = a.City
	updataMap["spec_address"] = a.SpecAddress

	err = qmsql.DEFAULTDB.Where("user_id = ?", a.UserId).Updates(&updataMap).Error

	return
}

func (a *Address) DeleteAddress(id int64) (err error) {
	err = qmsql.DEFAULTDB.Where("id = ?", id).Delete(&a).Error
	return
}
