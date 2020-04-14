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
	OrderDetail []OrderDetail `json:"orderDetail"`
	OrderId     uuid.UUID     `json:"orderId" gorm:"ForeignKey:OrderId;AssociationForeignKey:OrderId"` // 订单号
	UserId      uuid.UUID     `json:"userId"`                                                          // 用户id
	Value       float64       `json:"value"`                                                           // 价格
	SpecAddress string        `json:"spec_address"`                                                    // 详细地址
	OrderType   int           `json:"orderType"`                                                       // 订单类型
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
	if err != nil {
		return err, nil, 0
	} else {
		var orderList []CustomerOrder
		err = db.Preload("OrderDetail").Where("user_id = ?", userId).Where("order_type = ?", orderType).Find(&orderList).Error
		return err, orderList, total
	}
}
func (co *CustomerOrder) AddOrder(cartList []Cart) (err error) {
	co.OrderId = uuid.NewV4()
	err = qmsql.DEFAULTDB.Create(&co).Error
	if err != nil {
		return
	}
	// 将购物车添加到订单
	for i := 0; i < len(cartList); i++ {
		if cartList[i].IsCheck == 1 {
			co.OrderDetail = append(co.OrderDetail, OrderDetail{
				Model:   gorm.Model{},
				OrderId: co.OrderId,
				Coffee:  cartList[i].Coffee,
				Count:   cartList[i].Count,
				Value:   cartList[i].Value,
			})
			err = qmsql.DEFAULTDB.Create(&co.OrderDetail).Error
		}
	}
	return
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
