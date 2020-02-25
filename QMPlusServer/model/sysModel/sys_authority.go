package sysModel

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type SysAuthority struct {
	gorm.Model
	AuthorityId     string         `json:"authorityId" gorm:"not null;unique"`
	AuthorityName   string         `json:"authorityName"`
	ParentId        string         `json:"parentId"`
	DataAuthorityId []SysAuthority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id;association_jointable_foreignkey:data_id"`
	Children        []SysAuthority `json:"children"`
}

// 创建角色
func (a *SysAuthority) CreateAuthority() (err error, authority *SysAuthority) {
	err = qmsql.DEFAULTDB.Create(a).Error
	return err, a
}

// 删除角色
func (a *SysAuthority) DeleteAuthority() (err error) {
	err = qmsql.DEFAULTDB.Where("authority_id = ?", a.AuthorityId).Find(&SysUser{}).Error
	if err != nil {
		err = qmsql.DEFAULTDB.Where("parentId = ?", a.AuthorityId).Find(&SysAuthority{}).Error
		if err != nil {
			err = qmsql.DEFAULTDB.Where("authority_id = ?", a.AuthorityId).First(a).Unscoped().Delete(a).Error
			new(CasbinModel).clearCasbin(0, a.AuthorityId)
		} else {
			err = errors.New("此角色存在子角色不允许删除")
		}
	} else {
		err = errors.New("此角色有用户正在使用禁止删除")
	}
	return err
}

// 分页获取数据  需要分页实现这个接口即可
func (a *SysAuthority) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(a, info)
	if err != nil {
		return
	} else {
		var authority []SysAuthority
		err = db.Preload("DataAuthorityId").Where("parent_id = 0").Find(&authority).Error
		if len(authority) > 0 {
			for k, _ := range authority {
				err = findChildrenAuthority(&authority[k])
			}
		}
		return err, authority, total
	}
}

func findChildrenAuthority(authority *SysAuthority) (err error) {
	err = qmsql.DEFAULTDB.Preload("DataAuthorityId").Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k, _ := range authority.Children {
			err = findChildrenAuthority(&authority.Children[k])
		}
	}
	return err
}

func (a *SysAuthority) SetDataAuthority() error {
	var s SysAuthority
	qmsql.DEFAULTDB.Preload("DataAuthorityId").First(&s, "id = ?", a.ID)
	err := qmsql.DEFAULTDB.Model(&s).Association("DataAuthorityId").Replace(&a.DataAuthorityId).Error
	return err
}

func (a *SysAuthority) GetAuthorityInfo() (err error,sa SysAuthority) {
	err = qmsql.DEFAULTDB.Preload("DataAuthorityId").Where("authority_id = ?",a.AuthorityId).First(&sa).Error
	return err,sa
}