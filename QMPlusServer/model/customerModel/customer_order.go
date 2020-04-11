package customerModel

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type CustomerOrder struct {
	gorm.Model
	UserId      uuid.UUID `json:"userId"`
	CoffeeName  string    `json:"coffee_name"`
	CoffeeSpec  string    `json:"coffee_spec"`
	Value       float32   `json:"value"`
	SpecAddress string    `json:"spec_address"`
}

func (co *CustomerOrder) GetInfoList(info modelInterface.PageInfo) (error, interface{}, int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(co, info)
	if err != nil {
		return err, nil, 0
	} else {
		var orderList []CustomerOrder
		err = db.Find(&orderList).Error
		return err, orderList, total
	}
}

func (co *CustomerOrder) AddOrder() (err error) {
	err = qmsql.DEFAULTDB.Create(&co).Error
	return
}

func (co *CustomerOrder) UpdateOrder() (err error) {
	updataMap := make(map[string]interface{})
	updataMap["coffee_name"] = co.CoffeeName
	updataMap["coffee_spec"] = co.CoffeeSpec
	updataMap["spec_address"] = co.SpecAddress

	err = qmsql.DEFAULTDB.Where("user_id = ?", co.UserId).Updates(&updataMap).Error

	return
}

func (co *CustomerOrder) DeleteOrder(id int64) (err error) {
	err = qmsql.DEFAULTDB.Where("id = ?", id).Delete(&co).Error
	return
}
