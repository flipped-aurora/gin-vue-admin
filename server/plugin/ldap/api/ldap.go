package api

import (
	"errors"
	systemApi "github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	systemModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	ldapGlobal "github.com/flipped-aurora/gin-vue-admin/server/plugin/ldap/global"
	ldapSvc "github.com/flipped-aurora/gin-vue-admin/server/plugin/ldap/service"
	ldapUtils "github.com/flipped-aurora/gin-vue-admin/server/plugin/ldap/utils"

	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-ldap/ldap/v3"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"time"
)

var store = base64Captcha.DefaultMemStore

type LdapApi struct {
	systemApi.BaseApi
}

// Login
// @Tags     Base
// @Summary  用户登录
// @Produce   application/json
// @Param    data  body      systemReq.Login                                             true  "用户名, 密码, 验证码"
// @Success  200   {object}  response.Response{data=systemRes.LoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /ldap/login [post]
func (b *LdapApi) LdapLogin(c *gin.Context) {
	var l systemReq.Login
	err := c.ShouldBindJSON(&l)
	key := c.ClientIP()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(l, utils.LoginVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 判断验证码是否开启
	openCaptcha := global.GVA_CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global.GVA_CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}

	var oc bool = openCaptcha == 0 || openCaptcha < ldapUtils.InterfaceConvertInt(v)
	if !oc || store.Verify(l.CaptchaId, l.Captcha, true) {
		var user systemModel.SysUser
		// 获取ldap用户信息
		err = ldapGlobal.Ldap.SearchFunc(
			ldapGlobal.Ldap.BuildSearchRequest(l.Username),
			func(result *ldap.SearchResult) error {
				if result == nil || len(result.Entries) == 0 {
					return errors.New("ldap searchResult not found")
				}
				entry := result.Entries[0]
				if err := ldapGlobal.Ldap.Bind(entry.DN, l.Password).Error; err != nil {
					global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
					return err
				}
				// ldap用户信息入库
				regUser := systemModel.SysUser{
					Username: l.Username,
					Password: l.Password,
				}
				if err := ldapUtils.FillSysUser(&regUser, entry, ldapGlobal.GlobalConfig.FieldMap); err != nil {
					return err
				}
				user, err = ldapSvc.ServiceGroupApp.LdapRegister(regUser)
				return err
			}).Error
		if err != nil {
			global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
			global.BlackCache.Increment(key, 1)
			response.FailWithMessage("用户名不存在或者密码错误", c)
			return
		}
		if user.Enable != 1 {
			global.GVA_LOG.Error("登陆失败! 用户被禁止登录!")
			// 验证码次数+1
			global.BlackCache.Increment(key, 1)
			response.FailWithMessage("用户被禁止登录", c)
			return
		}
		user.Authority.DefaultRouter = "dashboard"
		b.TokenNext(c, user)
		return
	}
	// 验证码次数+1
	global.BlackCache.Increment(key, 1)
	response.FailWithMessage("验证码错误", c)
}
