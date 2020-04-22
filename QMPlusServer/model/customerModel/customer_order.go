package customerModel

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/coffeeModel"
	"gin-vue-admin/model/modelInterface"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type CustomerOrder struct {
	gorm.Model
	OrderDetail []OrderDetail `json:"orderDetail" gorm:"ForeignKey:OrderId;AssociationForeignKey:OrderId"`
	OrderId     uuid.UUID     `json:"orderId"`      // 订单号
	Consignee   string        `json:"consignee"`    // 收货人
	UserId      uuid.UUID     `json:"userId"`       // 用户id
	Value       float64       `json:"value"`        // 价格
	SpecAddress string        `json:"spec_address"` // 详细地址
	Phone       string        `json:"phone"`
	OrderType   int           `json:"orderType"` // 订单类型
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
func (co *CustomerOrder) GetInfoListByOrderType(info modelInterface.PageInfo, orderType int, userId uuid.UUID) (error, interface{}, int) {
	err, db, total := servers.PagingServer(co, info)
	var orderList []CustomerOrder
	if err != nil {
		return err, nil, 0
	} else {
		if orderType != 0 {
			err = db.Preload("OrderDetail").Where("user_id = ?", userId).Where("order_type = ?", orderType).Find(&orderList).Error
			for i := 0; i < len(orderList); i++ {
				for j := 0; j < len(orderList[i].OrderDetail); j++ {
					var coffee coffeeModel.Coffee
					coffee.GetCoffeeByUUID(orderList[i].OrderDetail[j].CoffeeId)
					orderList[i].OrderDetail[j].Coffee = coffee
				}
			}
		} else {
			err = db.Preload("OrderDetail").Where("user_id = ?", userId).Find(&orderList).Error
			for i := 0; i < len(orderList); i++ {
				for j := 0; j < len(orderList[i].OrderDetail); j++ {
					var coffee coffeeModel.Coffee
					coffee.GetCoffeeByUUID(orderList[i].OrderDetail[j].CoffeeId)
					orderList[i].OrderDetail[j].Coffee = coffee
				}
			}
		}

		return err, orderList, total
	}
}
func (co *CustomerOrder) AddOrder(cartList []Cart) (orderId uuid.UUID, err error) {
	co.OrderId = uuid.NewV4()
	co.Value = 0
	// 将购物车添加到订单
	for i := 0; i < len(cartList); i++ {
		if cartList[i].IsCheck == 1 {
			orderDetail := OrderDetail{
				Model:    gorm.Model{},
				OrderId:  co.OrderId,
				CoffeeId: cartList[i].CoffeeId,
				Count:    cartList[i].Count,
				Value:    cartList[i].Value,
				Spec:     cartList[i].Spec,
			}
			co.Value += orderDetail.Value
			co.OrderDetail = append(co.OrderDetail, orderDetail)
			err = qmsql.DEFAULTDB.Create(&orderDetail).Error
			err = qmsql.DEFAULTDB.Where("id = ?", cartList[i].ID).Delete(cartList[i]).Error
		}
	}
	err = qmsql.DEFAULTDB.Create(&co).Error
	return co.OrderId, err
}

func (co *CustomerOrder) DeleteOrder(orderId uuid.UUID) (err error) {
	var orderList []OrderDetail
	err = qmsql.DEFAULTDB.Where("order_id = ?", orderId).Delete(&orderList).Error
	if err != nil {
		return
	}
	err = qmsql.DEFAULTDB.Where("order_id = ?", orderId).Delete(&co).Error
	return
}

func (co *CustomerOrder) GetOrderDetail(orderId uuid.UUID) (err error) {
	err = qmsql.DEFAULTDB.Preload("OrderDetail").Where("order_id = ?", orderId).Find(&co).Error
	var coffee coffeeModel.Coffee
	for i := 0; i < len(co.OrderDetail); i++ {
		err = coffee.GetCoffeeByUUID(co.OrderDetail[i].CoffeeId)
		if err == nil {
			co.OrderDetail[i].Coffee = coffee
		}
	}
	return
}
func (co *CustomerOrder) FinishOrder(orderId uuid.UUID) {

}
