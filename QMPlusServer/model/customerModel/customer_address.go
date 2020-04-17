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
	Consignee   string    `json:"consignee"`
	Province    string    `json:"province"`
	City        string    `json:"city"`
	Town        string    `json:"town"`
	SpecAddress string    `json:"specAddress"`
	Phone       string    `json:"phone"`
	IsDefault   int       `json:"isDefault" gorm:"force"`
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

func (a *Address) GetInfoListByUserId(info modelInterface.PageInfo, userId uuid.UUID) (error, interface{}, int) {
	err, db, total := servers.PagingServer(a, info)
	if err == nil {
		var addressList []Address
		err = db.Where("user_id = ?", userId).Find(&addressList).Error
		return err, addressList, total
	} else {
		return err, nil, 0
	}
}

func (a *Address) AddAddress() (err error) {
	var temp Address
	if a.IsDefault == 1 {
		findOne := qmsql.DEFAULTDB.Where("is_default = ?", 1).Where("user_id = ?", a.UserId).Find(&temp).Error
		if findOne == nil {
			err = qmsql.DEFAULTDB.Model(&temp).Select("is_default").Update("is_default", gorm.Expr("is_default - 1")).Error
		}
	}
	err = qmsql.DEFAULTDB.Create(&a).Error
	return
}
func (a *Address) SetDefaultAddress() (err error) {
	var temp Address
	findOne := qmsql.DEFAULTDB.Where("is_default = ?", 1).Where("user_id = ?", a.UserId).Find(&temp).Error
	if findOne == nil {
		err = qmsql.DEFAULTDB.Model(&temp).Select("is_default").Update("is_default", gorm.Expr("is_default - 1")).Error
	}
	err = qmsql.DEFAULTDB.Model(&a).Select("is_default").Update("is_default", 1).Error
	return
}

func (a *Address) UpdateAddress() (err error) {
	updataMap := make(map[string]interface{})
	updataMap["consignee"] = a.Consignee
	updataMap["province"] = a.Province
	updataMap["city"] = a.City
	updataMap["spec_address"] = a.SpecAddress
	updataMap["Phone"] = a.Phone
	updataMap["isDefault"] = a.IsDefault
	err = qmsql.DEFAULTDB.Where("user_id = ?", a.UserId).Updates(&updataMap).Error

	return
}

func (a *Address) DeleteAddress(id int64) (err error) {
	err = qmsql.DEFAULTDB.Where("id = ?", id).Delete(&a).Error
	return
}
