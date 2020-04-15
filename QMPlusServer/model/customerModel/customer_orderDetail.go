package customerModel

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/coffeeModel"
	"gin-vue-admin/model/modelInterface"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type OrderDetail struct {
	gorm.Model
	OrderId  uuid.UUID          `json:"orderId"`
	Coffee   coffeeModel.Coffee `json:"coffee"`
	CoffeeId uuid.UUID          `json:"coffee_id"`
	Count    int                `json:"count"`
	Value    float64            `json:"value"`
}

func (co *OrderDetail) GetInfoList(info modelInterface.PageInfo) (error, interface{}, int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(co, info)
	if err != nil {
		return err, nil, 0
	} else {
		var orderList []OrderDetail
		err = db.Find(&orderList).Error
		return err, orderList, total
	}
}

//func (co *OrderDetail) AddOrder() (err error) {
//	err = qmsql.DEFAULTDB.Create(&co).Error
//	return
//}

//func (co *OrderDetail) UpdateOrder() (err error) {
//	updataMap := make(map[string]interface{})
//	updataMap["spec_address"] = co.SpecAddress
//	updataMap["orderType"] = co.OrderType
//
//	err = qmsql.DEFAULTDB.Where("user_id = ?", co.UserId).Updates(&updataMap).Error
//
//	return
//}
//
//func (co *OrderDetail) DeleteOrder(id int64) (err error) {
//	err = qmsql.DEFAULTDB.Where("id = ?", id).Delete(&co).Error
//	return
//}
