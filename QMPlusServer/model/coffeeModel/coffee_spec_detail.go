package coffeeModel

import (
	"gin-vue-admin/init/qmsql"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type CoffeeSpecDetail struct {
	gorm.Model
	SpecId     uuid.UUID `json:"spec_id"`
	CoffeeId   uuid.UUID `json:"coffee_id"`
	Value      string    `json:"value"`
	PriceIncre float64   `json:"price_incre"`
}

func (c *CoffeeSpecDetail) AddCoffeeSpecDetail() (err error) {
	err = qmsql.DEFAULTDB.Create(&c).Error
	return
}
