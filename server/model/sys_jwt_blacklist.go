package model

import (
	"github.com/jinzhu/gorm"
)

type JwtBlacklist struct {
	gorm.Model
	Jwt string `gorm:"type:text;comment:'jwt'"`
}
