package customerModel

import (
	"errors"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/coffeeModel"
	"gin-vue-admin/model/modelInterface"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Cart struct {
	gorm.Model
	Coffee   coffeeModel.Coffee `json:"coffee" gorm:"ForeignKey:CoffeeId;AssociationForeignKey:CoffeeId"`
	CoffeeId uuid.UUID          `json:"coffee_id"`
	Spec     string             `json:"spec"`
	Value    float64            `json:"value"`
	Count    int                `json:"count"`
	Customer Customers          `json:"customer" gorm:"ForeignKey:UserId;AssociationForeignKey:UserId"`
	UserId   uuid.UUID          `json:"user_id"`
	IsCheck  int                `json:"isCheck"`
}

func (c *Cart) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(c, info)
	if err != nil {
		return
	} else {
		var cartList []Cart
		err = db.Find(&cartList).Error
		return err, cartList, total
	}
}

func (c *Cart) GetInfoListByUserId(info modelInterface.PageInfo, userId uuid.UUID) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(c, info)
	if err != nil {
		return
	} else {
		var cartList []Cart
		err = db.Where("user_id = ?", userId).Find(&cartList).Error
		for i, _ := range cartList {
			var coffee coffeeModel.Coffee
			coffee.GetCoffeeByUUID(cartList[i].CoffeeId)
			cartList[i].Coffee = coffee
		}
		return err, cartList, total
	}
}

func (c *Cart) AddCart(userId uuid.UUID, coffeeId uuid.UUID, spec string, value float64) (err error) {
	findOne := qmsql.DEFAULTDB.Where("user_id = ?", userId).Where("coffee_id = ?", coffeeId).Where("spec = ?", spec).Find(&c).Error
	var coffee coffeeModel.Coffee
	err = coffee.GetCoffeeByUUID(coffeeId)
	var customer Customers
	err = customer.GetCustomerById(userId)
	c.Coffee = coffee
	c.Customer = customer
	if findOne != nil {
		c.Spec = spec
		c.Value = value
		c.Count = 1
		c.IsCheck = 1
		c.CoffeeId = coffeeId
		c.UserId = customer.UUID
		err = qmsql.DEFAULTDB.Create(&c).Error
	} else {
		upDataMap := make(map[string]interface{})
		upDataMap["value"] = c.Value + value
		upDataMap["count"] = c.Count + 1
		upDataMap["isCheck"] = c.IsCheck
		upDataMap["spec"] = c.Spec
		err = qmsql.DEFAULTDB.Where("coffee_id = ?", coffeeId).Where("user_id = ?", userId).Find(&c).Updates(upDataMap).Error
	}
	return
}

func (c *Cart) ReduceCart(userId uuid.UUID, coffeeId uuid.UUID, spec string) (err error) {
	findOne := qmsql.DEFAULTDB.Where("user_id = ?", userId).Where("coffee_id = ?", coffeeId).Where("spec = ?", spec).Find(&c).Error
	var coffee coffeeModel.Coffee
	err = coffee.GetCoffeeByUUID(coffeeId)
	var customer Customers
	err = customer.GetCustomerById(userId)
	c.Coffee = coffee
	c.Customer = customer
	if findOne != nil {
		return errors.New("没有该咖啡")
	} else {
		upDataMap := make(map[string]interface{})
		upDataMap["value"] = c.Value - coffee.Value
		upDataMap["count"] = c.Count - 1
		upDataMap["isCheck"] = c.IsCheck
		upDataMap["spec"] = c.Spec
		if c.Count-1 == 0 || c.Value-coffee.Value <= 0 {
			err = qmsql.DEFAULTDB.Where("coffee_id = ?", coffeeId).Where("user_id = ?", userId).Where("spec = ?", spec).Delete(&c).Error
		} else {
			err = qmsql.DEFAULTDB.Where("coffee_id = ?", coffeeId).Where("user_id = ?", userId).Where("spec = ?", spec).Find(&c).Updates(upDataMap).Error
		}
	}
	return
}

func (c *Cart) DeleteCart(userId uuid.UUID) (err error) {
	err = qmsql.DEFAULTDB.Where("user_id = ?", userId).Delete(&c).Error
	return
}

func (c *Cart) CheckStatus(id int64, isCheck bool) (err error) {

	upDataMap := make(map[string]interface{})
	upDataMap["isCheck"] = isCheck
	err = qmsql.DEFAULTDB.Where("id = ?", id).Find(&c).Updates(upDataMap).Error
	return
}
