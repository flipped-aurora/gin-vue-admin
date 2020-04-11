package coffeeModel

import (
	"errors"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Coffee struct {
	gorm.Model
	UUID       uuid.UUID    `json:"uuid"`
	Name       string       `json:"name"`
	Value      float64      `json:"value"`
	Des        string       `json:"des"`
	Img        string       `json:"img"`
	Coffeetype CoffeeType   `json:"type" gorm:"ForeignKey:Code;AssociationForeignKey:Code"`
	Code       string       `json:"code"`
	CoffeeSpec []CoffeeSpec `json:"spec" gorm:"ForeignKey:CoffeeId;AssociationForeignKey:CoffeeId"`
}

func (c *Coffee) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(c, info)
	if err != nil {
		return
	} else {
		var coffeeList []Coffee
		err = db.Preload("Coffeetype").Find(&coffeeList).Error
		return err, coffeeList, total
	}
}
func (c *Coffee) GetInfoListByCode(info modelInterface.PageInfo, code string) (err error, list interface{}, total int) {
	err, db, total := servers.PagingServer(c, info)
	if err != nil {
		return
	} else {
		var coffeeList []Coffee
		err = db.Preload("Coffeetype").Where("code = ?", code).Find(&coffeeList).Error
		return err, coffeeList, total
	}
}
func (c *Coffee) AddCoffee() (err error) {
	findOne := qmsql.DEFAULTDB.Where("name = ?", c.Name).Find(&Coffee{}).Error
	if findOne != nil {
		c.UUID = uuid.NewV4()
		err = qmsql.DEFAULTDB.Create(c).Error
	} else {
		err = errors.New("存在重复名称，请修改名称")
	}
	return
}

func (c *Coffee) UpdateCoffee() (err error) {
	upDataMap := make(map[string]interface{})
	upDataMap["name"] = c.Name
	upDataMap["value"] = c.Value
	upDataMap["des"] = c.Des
	upDataMap["img"] = c.Img
	err = qmsql.DEFAULTDB.Where("id = ?", c.ID).Find(&Coffee{}).Updates(upDataMap).Error
	return
}

func (c *Coffee) UpdateCoffeeType(uuid uuid.UUID, code string) (err error) {
	err = qmsql.DEFAULTDB.Where("uuid = ?", uuid).First(&Coffee{}).Update("code", code).Error
	return
}

func (c *Coffee) DeleteCoffee(id uuid.UUID) (err error) {
	err = qmsql.DEFAULTDB.Where("uuid = ?", id).Delete(&c).Error
	return
}

func (c *Coffee) GetCoffeeByUUID(uuid uuid.UUID) (err error) {
	err = qmsql.DEFAULTDB.Where("uuid = ?", uuid).Find(&c).Error
	return
}
