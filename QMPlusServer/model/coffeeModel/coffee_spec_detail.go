package coffeeModel

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type CoffeeSpecDetail struct {
	gorm.Model
	SpecId   uuid.UUID `json:"spec_id"`
	CoffeeId uuid.UUID `json:"coffee_id"`
	Value    string    `json:"value"`
	Price    float64   `json:"price"`
}
