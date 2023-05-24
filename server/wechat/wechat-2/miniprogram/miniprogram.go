package miniprogram

import (
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/credential"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/analysis"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/auth"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/business"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/config"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/content"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/context"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/encryptor"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/message"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/privacy"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/qrcode"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/riskcontrol"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/security"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/shortlink"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/subscribe"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/tcb"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/urllink"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/urlscheme"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/werun"
)

// MiniProgram 微信小程序相关API
type MiniProgram struct {
	ctx *context.Context
}

// NewMiniProgram 实例化小程序API
func NewMiniProgram(cfg *config.Config) *MiniProgram {
	defaultAkHandle := credential.NewDefaultAccessToken(cfg.AppID, cfg.AppSecret, credential.CacheKeyMiniProgramPrefix, cfg.Cache)
	ctx := &context.Context{
		Config:            cfg,
		AccessTokenHandle: defaultAkHandle,
	}
	return &MiniProgram{ctx}
}

// SetAccessTokenHandle 自定义access_token获取方式
func (miniProgram *MiniProgram) SetAccessTokenHandle(accessTokenHandle credential.AccessTokenHandle) {
	miniProgram.ctx.AccessTokenHandle = accessTokenHandle
}

// GetContext get Context
func (miniProgram *MiniProgram) GetContext() *context.Context {
	return miniProgram.ctx
}

// GetEncryptor  小程序加解密
func (miniProgram *MiniProgram) GetEncryptor() *encryptor.Encryptor {
	return encryptor.NewEncryptor(miniProgram.ctx)
}

// GetAuth 登录/用户信息相关接口
func (miniProgram *MiniProgram) GetAuth() *auth.Auth {
	return auth.NewAuth(miniProgram.ctx)
}

// GetAnalysis 数据分析
func (miniProgram *MiniProgram) GetAnalysis() *analysis.Analysis {
	return analysis.NewAnalysis(miniProgram.ctx)
}

// GetBusiness 业务接口
func (miniProgram *MiniProgram) GetBusiness() *business.Business {
	return business.NewBusiness(miniProgram.ctx)
}

// GetPrivacy 小程序隐私协议相关API
func (miniProgram *MiniProgram) GetPrivacy() *privacy.Privacy {
	return privacy.NewPrivacy(miniProgram.ctx)
}

// GetQRCode 小程序码相关API
func (miniProgram *MiniProgram) GetQRCode() *qrcode.QRCode {
	return qrcode.NewQRCode(miniProgram.ctx)
}

// GetTcb 小程序云开发API
func (miniProgram *MiniProgram) GetTcb() *tcb.Tcb {
	return tcb.NewTcb(miniProgram.ctx)
}

// GetSubscribe 小程序订阅消息
func (miniProgram *MiniProgram) GetSubscribe() *subscribe.Subscribe {
	return subscribe.NewSubscribe(miniProgram.ctx)
}

// GetCustomerMessage 客服消息接口
func (miniProgram *MiniProgram) GetCustomerMessage() *message.Manager {
	return message.NewCustomerMessageManager(miniProgram.ctx)
}

// GetWeRun 微信运动接口
func (miniProgram *MiniProgram) GetWeRun() *werun.WeRun {
	return werun.NewWeRun(miniProgram.ctx)
}

// GetContentSecurity 内容安全接口
func (miniProgram *MiniProgram) GetContentSecurity() *content.Content {
	return content.NewContent(miniProgram.ctx)
}

// GetURLLink 小程序URL Link接口
func (miniProgram *MiniProgram) GetURLLink() *urllink.URLLink {
	return urllink.NewURLLink(miniProgram.ctx)
}

// GetRiskControl 安全风控接口
func (miniProgram *MiniProgram) GetRiskControl() *riskcontrol.RiskControl {
	return riskcontrol.NewRiskControl(miniProgram.ctx)
}

// GetSecurity 内容安全接口
func (miniProgram *MiniProgram) GetSecurity() *security.Security {
	return security.NewSecurity(miniProgram.ctx)
}

// GetShortLink 小程序短链接口
func (miniProgram *MiniProgram) GetShortLink() *shortlink.ShortLink {
	return shortlink.NewShortLink(miniProgram.ctx)
}

// GetSURLScheme 小程序URL Scheme接口
func (miniProgram *MiniProgram) GetSURLScheme() *urlscheme.URLScheme {
	return urlscheme.NewURLScheme(miniProgram.ctx)
}
