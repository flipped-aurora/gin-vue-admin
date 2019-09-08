package dbModel

import (
	"github.com/jinzhu/gorm"
	"main/init/qmsql"
)

type Authority struct {
	gorm.Model    `json:"-"`
	AuthorityId   uint   `json:"authorityId"`
	AuthorityName string `json:"authorityName"`
	Menus         []Menu `json:"_"`
}

func (a *Authority) CreateAuthority() (err error, authority *Authority) {
	err = qmsql.DEFAULTDB.Create(a).Error
	return err, a
}
