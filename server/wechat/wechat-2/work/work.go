package work

import (
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/credential"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/work/addresslist"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/work/config"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/work/context"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/work/externalcontact"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/work/kf"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/work/material"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/work/message"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/work/msgaudit"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/work/oauth"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/work/robot"
)

// Work 企业微信
type Work struct {
	ctx *context.Context
}

// NewWork init work
func NewWork(cfg *config.Config) *Work {
	defaultAkHandle := credential.NewWorkAccessToken(cfg.CorpID, cfg.CorpSecret, credential.CacheKeyWorkPrefix, cfg.Cache)
	ctx := &context.Context{
		Config:            cfg,
		AccessTokenHandle: defaultAkHandle,
	}
	return &Work{ctx: ctx}
}

// GetContext get Context
func (wk *Work) GetContext() *context.Context {
	return wk.ctx
}

// GetOauth get oauth
func (wk *Work) GetOauth() *oauth.Oauth {
	return oauth.NewOauth(wk.ctx)
}

// GetMsgAudit get msgAudit
func (wk *Work) GetMsgAudit() (*msgaudit.Client, error) {
	return msgaudit.NewClient(wk.ctx.Config)
}

// GetKF get kf
func (wk *Work) GetKF() (*kf.Client, error) {
	return kf.NewClient(wk.ctx.Config)
}

// GetExternalContact get external_contact
func (wk *Work) GetExternalContact() *externalcontact.Client {
	return externalcontact.NewClient(wk.ctx)
}

// GetAddressList get address_list
func (wk *Work) GetAddressList() *addresslist.Client {
	return addresslist.NewClient(wk.ctx)
}

// GetMaterial get material
func (wk *Work) GetMaterial() *material.Client {
	return material.NewClient(wk.ctx)
}

// GetRobot get robot
func (wk *Work) GetRobot() *robot.Client {
	return robot.NewClient(wk.ctx)
}

// GetMessage get robot
func (wk *Work) GetMessage() *message.Client {
	return message.NewClient(wk.ctx)
}
