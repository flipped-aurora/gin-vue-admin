package sysModel

import (
	"github.com/jinzhu/gorm"
	"main/init/qmsql"
)

type SysApiAuthority struct {
	gorm.Model
	AuthorityId string
	Authority   SysAuthority `gorm:"ForeignKey:AuthorityId;AssociationForeignKey:AuthorityId"` //其实没有关联的必要
	ApiId       uint
	Api         SysApi
}

//创建角色api关联关系
func (a *SysApiAuthority) SetAuthAndApi(authId string, apisid []uint) (err error) {
	for _, v := range apisid {
		err = qmsql.DEFAULTDB.Create(&SysApiAuthority{AuthorityId: authId, ApiId: v}).Error
		if err != nil {
			return err
		}
	}
	return nil
}

// 获取角色api关联关系
func (a *SysApiAuthority) GetAuthAndApi(authId string) (err error,apiIds []uint) {
	var apis []SysApiAuthority
	err = qmsql.DEFAULTDB.Where("authority_id = ?", authId).Find(&apis).Error
	for _, v := range apis {
		apiIds = append(apiIds,v.ApiId)
	}
	return nil,apiIds
}