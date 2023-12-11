package initialize

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/alipay"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wxpay"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/email"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/plugin"
	"github.com/gin-gonic/gin"
)

func PluginInit(group *gin.RouterGroup, Plugin ...plugin.Plugin) {
	for i := range Plugin {
		PluginGroup := group.Group(Plugin[i].RouterPath())
		Plugin[i].Register(PluginGroup)
	}
}

func InstallPlugin(Router *gin.Engine) {
	PublicGroup := Router.Group("")
	fmt.Println("无鉴权插件安装==》", PublicGroup)
	PrivateGroup := Router.Group("")
	fmt.Println("鉴权插件安装==》", PrivateGroup)
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	//  添加跟角色挂钩权限的插件 示例 本地示例模式于在线仓库模式注意上方的import 可以自行切换 效果相同
	PluginInit(PrivateGroup, email.CreateEmailPlug(
		global.GVA_CONFIG.Email.To,
		global.GVA_CONFIG.Email.From,
		global.GVA_CONFIG.Email.Host,
		global.GVA_CONFIG.Email.Secret,
		global.GVA_CONFIG.Email.Nickname,
		global.GVA_CONFIG.Email.Port,
		global.GVA_CONFIG.Email.IsSSL,
	))

	//微信支付插件 不需要鉴权
	PluginInit(PublicGroup, wxpay.CreatePayPlug(
		global.GVA_CONFIG.WxPay.MchID,
		global.GVA_CONFIG.WxPay.AppID,
		global.GVA_CONFIG.WxPay.Secret,
		global.GVA_CONFIG.WxPay.NotifyUrl,
		global.GVA_CONFIG.WxPay.MchCertificateSerialNumber,
		global.GVA_CONFIG.WxPay.MchAPIv3Key,
		global.GVA_CONFIG.WxPay.PemPath,
		global.GVA_CONFIG.WxPay.Load,
	))

	//支付宝支付插件 不需要鉴权
	PluginInit(PublicGroup, alipay.CreateAlipayPlug(
		global.GVA_CONFIG.AliPay.Appid,
		global.GVA_CONFIG.AliPay.EncryptKey,
		global.GVA_CONFIG.AliPay.PrivateKey,
		global.GVA_CONFIG.AliPay.AlipayPublicContentRSA2,
		global.GVA_CONFIG.AliPay.AlipayRootContent,
		global.GVA_CONFIG.AliPay.AppPublicContent,
		global.GVA_CONFIG.AliPay.NotifyUrl,
		global.GVA_CONFIG.AliPay.ReturnURL,
		global.GVA_CONFIG.AliPay.Load,
	))
}
