package coffeeModel

import (
	"errors"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"github.com/jinzhu/gorm"
)

type CoffeeType struct {
	gorm.Model
	Code  string `json:"code" gorm:"not null;unique"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

func (c *CoffeeType) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(c, info)
	if err != nil {
		return
	} else {
		var coffeetypeList []CoffeeType
		err = db.Find(&coffeetypeList).Error
		return err, coffeetypeList, total
	}
}

func (c *CoffeeType) AddCoffeeType() (err error) {
	findOne := qmsql.DEFAULTDB.Where("name = ?", c.Name).Find(&CoffeeType{}).Error
	if findOne != nil {
		err = qmsql.DEFAULTDB.Create(c).Error
	} else {
		err = errors.New("存在重复名称，请修改名称")
	}
	return
}

func (c *CoffeeType) UpdateCoffeeType() (err error) {
	upDataMap := make(map[string]interface{})
	findOne := qmsql.DEFAULTDB.Where("name = ?", c.Name).Find(&CoffeeType{}).Error
	if findOne != nil {
		err = qmsql.DEFAULTDB.Where("code = ?", c.Code).Find(&CoffeeType{}).Error
		upDataMap["name"] = c.Name
		err = qmsql.DEFAULTDB.Where("id = ?", c.ID).Find(&CoffeeType{}).Updates(upDataMap).Error
	} else {
		err = errors.New("存在重复名称，请修改名称")
	}

	return
}

func (c *CoffeeType) DeleteCoffeeType(code string) (err error) {
	coffees := []Coffee{}
	find := qmsql.DEFAULTDB.Where("code = ?", code).Find(&coffees).Error
	if find != nil {
		err = qmsql.DEFAULTDB.Where("code = ?", code).Delete(&c).Error
	} else {
		for _, co := range coffees {
			err = qmsql.DEFAULTDB.Delete(&co).Error
		}
		err = qmsql.DEFAULTDB.Where("code = ?", code).Delete(&c).Error
	}

	return
}

func (c *CoffeeType) GetCoffeeTypeByCode(code string) (err error) {
	err = qmsql.DEFAULTDB.Where("code = ?", code).Find(&c).Error
	return
}
