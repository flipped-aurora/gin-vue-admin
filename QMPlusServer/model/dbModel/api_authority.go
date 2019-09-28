package dbModel

import (
	"github.com/jinzhu/gorm"
	"main/init/qmsql"
)

type ApiAuthority struct {
	gorm.Model
	AuthorityId string
	Authority   Authority `gorm:"ForeignKey:AuthorityId;AssociationForeignKey:AuthorityId"` //其实没有关联的必要
	ApiId       uint
	Api         Api
}

//创建角色api关联关系
func (a *ApiAuthority) SetAuthAndApi(authId string, apis []Api) (err error) {
	err = qmsql.DEFAULTDB.Where("authority_id = ?", authId).Delete(&ApiAuthority{}).Error
	for _, v := range apis {
		err = qmsql.DEFAULTDB.Create(&ApiAuthority{AuthorityId: authId, ApiId: v.ID}).Error
		if err != nil {
			return err
		}
	}
	return nil
}
