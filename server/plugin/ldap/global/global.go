package global

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/ldap/config"
	ldapUtils "github.com/flipped-aurora/gin-vue-admin/server/plugin/ldap/utils"
	"go.uber.org/zap"
)

var GlobalConfig = new(config.Ldap)

var (
	Ldap *ldapUtils.Ldap
)

func InitLdap() {
	Ldap = ldapUtils.NewLdap(&ldapUtils.Option{
		BindDN:  GlobalConfig.BindDN,
		BindPwd: GlobalConfig.BindPwd,
		Host:    GlobalConfig.Host,
		TLS:     GlobalConfig.TLS,
		BaseDN:  GlobalConfig.BaseDN,
		Filter:  GlobalConfig.Filter,
	})
	err := Ldap.Connect()
	if err != nil {
		global.GVA_LOG.Error("ldap connect  failed, err:", zap.Error(err))
	} else {
		global.GVA_LOG.Info("ldap connect success", zap.Any("config", GlobalConfig))
	}
}
