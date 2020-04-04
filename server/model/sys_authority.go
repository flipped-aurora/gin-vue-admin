package model

import (
	"gin-vue-admin/global"
	"github.com/pkg/errors"
	"time"
)

type SysAuthority struct {
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time     `sql:"index"`
	AuthorityId     string         `json:"authorityId" gorm:"not null;unique;primary_key"`
	AuthorityName   string         `json:"authorityName"`
	ParentId        string         `json:"parentId"`
	DataAuthorityId []SysAuthority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id;association_jointable_foreignkey:data_authority_id"`
	Children        []SysAuthority `json:"children"`
	SysBaseMenus    []SysBaseMenu  `json:"menus" gorm:"many2many:sys_authority_menus;"`
}

// 创建角色
func (a *SysAuthority) CreateAuthority() (err error, authority *SysAuthority) {
	err = global.GVA_DB.Create(a).Error
	return err, a
}

// 删除角色
func (a *SysAuthority) DeleteAuthority() (err error) {
	err = global.GVA_DB.Where("authority_id = ?", a.AuthorityId).Find(&SysUser{}).Error
	if err != nil {
		err = global.GVA_DB.Where("parent_id = ?", a.AuthorityId).Find(&SysAuthority{}).Error
		if err != nil {
			err = global.GVA_DB.Where("authority_id = ?", a.AuthorityId).First(a).Unscoped().Delete(a).Error
			new(CasbinModel).clearCasbin(0, a.AuthorityId)
		} else {
			err = errors.New("此角色存在子角色不允许删除")
		}
	} else {
		err = errors.New("此角色有用户正在使用禁止删除")
	}
	return err
}

// 分页获取数据
func (a *SysAuthority) GetInfoList(info PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB
	if err != nil {
		return
	} else {
		var authority []SysAuthority
		err = db.Limit(limit).Offset(offset).Preload("DataAuthorityId").Where("parent_id = 0").Find(&authority).Error
		if len(authority) > 0 {
			for k, _ := range authority {
				err = findChildrenAuthority(&authority[k])
			}
		}
		return err, authority, total
	}
}

func findChildrenAuthority(authority *SysAuthority) (err error) {
	err = global.GVA_DB.Preload("DataAuthorityId").Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k, _ := range authority.Children {
			err = findChildrenAuthority(&authority.Children[k])
		}
	}
	return err
}

func (a *SysAuthority) SetDataAuthority() error {
	var s SysAuthority
	global.GVA_DB.Preload("DataAuthorityId").First(&s, "authority_id = ?", a.AuthorityId)
	err := global.GVA_DB.Model(&s).Association("DataAuthorityId").Replace(&a.DataAuthorityId).Error
	return err
}

func (a *SysAuthority) SetMuneAuthority() error {
	var s SysAuthority
	global.GVA_DB.Preload("SysBaseMenus").First(&s, "authority_id = ?", a.AuthorityId)
	err := global.GVA_DB.Model(&s).Association("SysBaseMenus").Replace(&a.SysBaseMenus).Error
	return err
}

func (a *SysAuthority) GetAuthorityInfo() (err error, sa SysAuthority) {
	err = global.GVA_DB.Preload("DataAuthorityId").Where("authority_id = ?", a.AuthorityId).First(&sa).Error
	return err, sa
}
