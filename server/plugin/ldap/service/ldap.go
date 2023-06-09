package service

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	systemModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"strings"
)

type LdapService struct{}

func (ls *LdapService) LdapRegister(u systemModel.SysUser) (userInter systemModel.SysUser, err error) {
	var user systemModel.SysUser
	err = global.GVA_DB.Where("username = ?", u.Username).First(&user).Error
	if err == nil {
		// ldap可能会变化的信息做更新
		isChange := false
		if strings.EqualFold(user.Phone, u.Phone) {
			user.Phone = u.Phone
			isChange = true
		}
		if strings.EqualFold(user.Email, u.Email) {
			user.Email = u.Email
			isChange = true
		}
		if strings.EqualFold(user.Username, u.Username) {
			user.Username = u.Username
			isChange = true
		}
		if !utils.BcryptCheck(u.Password, user.Password) {
			user.Password = utils.BcryptHash(u.Password)
			isChange = true
		}
		if isChange {
			err = global.GVA_DB.Where("id = ?", user.ID).UpdateColumns(&user).Error
		}
		return user, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 否则 附加uuid 密码hash加密 注册
		u.Password = utils.BcryptHash(u.Password)
		u.Enable = 1
		u.UUID = uuid.Must(uuid.NewV4())
		err = global.GVA_DB.Create(&u).Error
	}
	return u, err
}
