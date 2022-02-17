package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var Authority = new(authority)

type authority struct{}

func (a *authority) TableName() string {
	return "sys_authorities"
}

func (a *authority) Initialize() error {
	entities := []system.SysAuthority{
		{AuthorityId: "888", AuthorityName: global.Translate("system.authority.normalUsers"), ParentId: "0", DefaultRouter: "dashboard"},
		{AuthorityId: "9528", AuthorityName: global.Translate("system.authority.testRole"), ParentId: "0", DefaultRouter: "dashboard"},
		{AuthorityId: "8881", AuthorityName: global.Translate("system.authority.normalUserSubRole"), ParentId: "888", DefaultRouter: "dashboard"},
	}
	if err := global.GVA_DB.Create(&entities).Error; err != nil {
		return errors.Wrapf(err, "%s "+global.Translate("general.tabelDataInitFail"), a.TableName())
	}
	return nil
}

func (a *authority) CheckDataExist() bool {
	if errors.Is(global.GVA_DB.Where("authority_id = ?", "8881").First(&system.SysAuthority{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
