package dbModel

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"main/controller/servers"
	"main/init/qmsql"
	"main/model/modelInterface"
)

type Authority struct {
	gorm.Model    `json:"-"`
	AuthorityId   string   `json:"authorityId" gorm:"not null;unique"`
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



// 分页获取数据  需要分页实现这个接口即可
func (a *Authority) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(a, info)
	if err != nil {
		return
	} else {
		var authority []Authority
		err = db.Find(&authority).Error
		return err, authority, total
	}
}