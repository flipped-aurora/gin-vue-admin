package dbModel

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"main/init/qmsql"
)

type Authority struct {
	gorm.Model    `json:"-"`
	AuthorityId   uint   `json:"authorityId" gorm:"not null;unique"`
	AuthorityName string `json:"authorityName"`
}

// 创建角色
func (a *Authority) CreateAuthority() (err error, authority *Authority) {
	err = qmsql.DEFAULTDB.Create(a).Error
	return err, a
}

// 删除角色
func (a *Authority) DeleteAuthority() (err error) {
	err = qmsql.DEFAULTDB.Where("authority_id = ?", a.AuthorityId).Find(&User{}).Error
	if err != nil {
		err = qmsql.DEFAULTDB.Where("authority_id = ?", a.AuthorityId).Delete(a).Error
	} else {
		err = errors.New("此角色有用户正在使用禁止删除")
	}
	return err
}
