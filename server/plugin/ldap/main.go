package ldap

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/ldap/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/ldap/router"
	"github.com/gin-gonic/gin"
)

type ldapPlugin struct{}

func CreateLdaplPlug(host, baseDN, bindDN, bindPwd, filter, fieldMap string, tls bool) *ldapPlugin {
	global.GlobalConfig.Host = host
	global.GlobalConfig.BaseDN = baseDN
	global.GlobalConfig.BindDN = bindDN
	global.GlobalConfig.BindPwd = bindPwd
	global.GlobalConfig.Filter = filter
	global.GlobalConfig.FieldMap = fieldMap
	global.GlobalConfig.TLS = tls
	// 初始化Ldap链接
	global.InitLdap()
	return &ldapPlugin{}
}

func (*ldapPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitLdapRouter(group)
}

func (*ldapPlugin) RouterPath() string {
	return "ldap"
}
