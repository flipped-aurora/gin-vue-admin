package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var UserAuthority = new(userAuthority)

type userAuthority struct{}

func (a *userAuthority) TableName() string {
	var entity system.SysUseAuthority
	return entity.TableName()
}

func (a *userAuthority) Initialize() error {
	entities := []system.SysUseAuthority{
		{SysUserId: 1, SysAuthorityAuthorityId: "888"},
		{SysUserId: 1, SysAuthorityAuthorityId: "8881"},
		{SysUserId: 1, SysAuthorityAuthorityId: "9528"},
		{SysUserId: 2, SysAuthorityAuthorityId: "888"},
	}
	if err := global.GVA_DB.Create(&entities).Error; err != nil {
		return errors.Wrap(err, a.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (a *userAuthority) CheckDataExist() bool {
	if errors.Is(global.GVA_DB.Where("sys_user_id = ? AND sys_authority_authority_id = ?", 2, "888").First(&system.SysUseAuthority{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
