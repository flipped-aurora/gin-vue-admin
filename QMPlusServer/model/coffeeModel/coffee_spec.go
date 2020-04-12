package coffeeModel

import (
	"errors"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type CoffeeSpec struct {
	gorm.Model
	CoffeeId uuid.UUID `json:"uuid" gorm:"not null;unique"`
	Name     string    `json:"name"`
	Sort     string    `json:"sort"`
}

func (c *CoffeeSpec) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(c, info)
	if err != nil {
		return
	} else {
		var coffeespecList []CoffeeSpec
		err = db.Preload("CoffeeId").Find(&coffeespecList).Error
		return err, coffeespecList, total
	}
}

func (c *CoffeeSpec) AddCoffeeSpec() (err error) {
	findOne := qmsql.DEFAULTDB.Where("name = ?", c.Name).Find(&CoffeeSpec{}).Error
	if findOne != nil {
		err = qmsql.DEFAULTDB.Create(c).Error
	} else {
		err = errors.New("存在重复名称，请修改名称")
	}
	return
}

func (c *CoffeeSpec) UpdateCoffeeSpec() (err error) {
	upDataMap := make(map[string]interface{})
	findOne := qmsql.DEFAULTDB.Where("name = ?", c.Name).Find(&CoffeeSpec{}).Error
	if findOne != nil {
		upDataMap["name"] = c.Name
		upDataMap["sort"] = c.Sort
		err = qmsql.DEFAULTDB.Where("id = ?", c.ID).Find(&CoffeeSpec{}).Updates(upDataMap).Error
	} else {
		err = errors.New("存在重复名称，请修改名称")
	}

	return
}

func (c *CoffeeSpec) DeleteCoffeeSpec(id int64) (err error) {
	err = qmsql.DEFAULTDB.Where("id = ?", id).Delete(&c).Error
	return
}

func (c *CoffeeSpec) GetCoffeeById(id int64) (err error) {
	err = qmsql.DEFAULTDB.Where("id = ?", id).Find(&c).Error
	return
}

func (c *CoffeeSpec) GetCoffeeSpecByCoffeeId(coffeeId uuid.UUID) (list []CoffeeSpec, err error) {
	err = qmsql.DEFAULTDB.Where("coffee_id = ?", coffeeId).Find(&list).Error

	return
}
