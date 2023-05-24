package officialaccount

import (
	stdcontext "context"
	"net/http"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/officialaccount/draft"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/officialaccount/freepublish"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/officialaccount/ocr"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/officialaccount/datacube"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/credential"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/officialaccount/basic"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/officialaccount/broadcast"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/officialaccount/config"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/officialaccount/context"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/officialaccount/customerservice"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/officialaccount/device"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/officialaccount/js"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/officialaccount/material"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/officialaccount/menu"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/officialaccount/message"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/officialaccount/oauth"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/officialaccount/server"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/officialaccount/user"
)

// OfficialAccount 微信公众号相关API
type OfficialAccount struct {
	ctx          *context.Context
	basic        *basic.Basic
	menu         *menu.Menu
	oauth        *oauth.Oauth
	material     *material.Material
	draft        *draft.Draft
	freepublish  *freepublish.FreePublish
	js           *js.Js
	user         *user.User
	templateMsg  *message.Template
	managerMsg   *message.Manager
	device       *device.Device
	broadcast    *broadcast.Broadcast
	datacube     *datacube.DataCube
	ocr          *ocr.OCR
	subscribeMsg *message.Subscribe
}

// NewOfficialAccount 实例化公众号API
func NewOfficialAccount(cfg *config.Config) *OfficialAccount {
	defaultAkHandle := credential.NewDefaultAccessToken(cfg.AppID, cfg.AppSecret, credential.CacheKeyOfficialAccountPrefix, cfg.Cache)
	ctx := &context.Context{
		Config:            cfg,
		AccessTokenHandle: defaultAkHandle,
	}
	return &OfficialAccount{ctx: ctx}
}

// SetAccessTokenHandle 自定义access_token获取方式
func (officialAccount *OfficialAccount) SetAccessTokenHandle(accessTokenHandle credential.AccessTokenHandle) {
	officialAccount.ctx.AccessTokenHandle = accessTokenHandle
}

// GetContext get Context
func (officialAccount *OfficialAccount) GetContext() *context.Context {
	return officialAccount.ctx
}

// GetBasic qr/url 相关配置
func (officialAccount *OfficialAccount) GetBasic() *basic.Basic {
	if officialAccount.basic == nil {
		officialAccount.basic = basic.NewBasic(officialAccount.ctx)
	}
	return officialAccount.basic
}

// GetMenu 菜单管理接口
func (officialAccount *OfficialAccount) GetMenu() *menu.Menu {
	if officialAccount.menu == nil {
		officialAccount.menu = menu.NewMenu(officialAccount.ctx)
	}
	return officialAccount.menu
}

// GetServer 消息管理：接收事件，被动回复消息管理
func (officialAccount *OfficialAccount) GetServer(req *http.Request, writer http.ResponseWriter) *server.Server {
	srv := server.NewServer(officialAccount.ctx)
	srv.Request = req
	srv.Writer = writer
	return srv
}

// GetAccessToken 获取access_token
func (officialAccount *OfficialAccount) GetAccessToken() (string, error) {
	return officialAccount.ctx.GetAccessToken()
}

// GetAccessTokenContext 获取access_token
func (officialAccount *OfficialAccount) GetAccessTokenContext(ctx stdcontext.Context) (string, error) {
	if c, ok := officialAccount.ctx.AccessTokenHandle.(credential.AccessTokenContextHandle); ok {
		return c.GetAccessTokenContext(ctx)
	}
	return officialAccount.ctx.GetAccessToken()
}

// GetOauth oauth2网页授权
func (officialAccount *OfficialAccount) GetOauth() *oauth.Oauth {
	if officialAccount.oauth == nil {
		officialAccount.oauth = oauth.NewOauth(officialAccount.ctx)
	}
	return officialAccount.oauth
}

// GetMaterial 素材管理
func (officialAccount *OfficialAccount) GetMaterial() *material.Material {
	if officialAccount.material == nil {
		officialAccount.material = material.NewMaterial(officialAccount.ctx)
	}
	return officialAccount.material
}

// GetDraft 草稿箱
func (officialAccount *OfficialAccount) GetDraft() *draft.Draft {
	if officialAccount.draft == nil {
		officialAccount.draft = draft.NewDraft(officialAccount.ctx)
	}
	return officialAccount.draft
}

// GetFreePublish 发布能力
func (officialAccount *OfficialAccount) GetFreePublish() *freepublish.FreePublish {
	if officialAccount.freepublish == nil {
		officialAccount.freepublish = freepublish.NewFreePublish(officialAccount.ctx)
	}
	return officialAccount.freepublish
}

// GetJs js-sdk配置
func (officialAccount *OfficialAccount) GetJs() *js.Js {
	if officialAccount.js == nil {
		officialAccount.js = js.NewJs(officialAccount.ctx)
	}
	return officialAccount.js
}

// GetUser 用户管理接口
func (officialAccount *OfficialAccount) GetUser() *user.User {
	if officialAccount.user == nil {
		officialAccount.user = user.NewUser(officialAccount.ctx)
	}
	return officialAccount.user
}

// GetTemplate 模板消息接口
func (officialAccount *OfficialAccount) GetTemplate() *message.Template {
	if officialAccount.templateMsg == nil {
		officialAccount.templateMsg = message.NewTemplate(officialAccount.ctx)
	}
	return officialAccount.templateMsg
}

// GetCustomerMessageManager 客服消息接口
func (officialAccount *OfficialAccount) GetCustomerMessageManager() *message.Manager {
	if officialAccount.managerMsg == nil {
		officialAccount.managerMsg = message.NewMessageManager(officialAccount.ctx)
	}
	return officialAccount.managerMsg
}

// GetDevice 获取智能设备的实例
func (officialAccount *OfficialAccount) GetDevice() *device.Device {
	if officialAccount.device == nil {
		officialAccount.device = device.NewDevice(officialAccount.ctx)
	}
	return officialAccount.device
}

// GetBroadcast 群发消息
// TODO 待完善
func (officialAccount *OfficialAccount) GetBroadcast() *broadcast.Broadcast {
	if officialAccount.broadcast == nil {
		officialAccount.broadcast = broadcast.NewBroadcast(officialAccount.ctx)
	}
	return officialAccount.broadcast
}

// GetDataCube 数据统计
func (officialAccount *OfficialAccount) GetDataCube() *datacube.DataCube {
	if officialAccount.datacube == nil {
		officialAccount.datacube = datacube.NewCube(officialAccount.ctx)
	}
	return officialAccount.datacube
}

// GetOCR OCR接口
func (officialAccount *OfficialAccount) GetOCR() *ocr.OCR {
	if officialAccount.ocr == nil {
		officialAccount.ocr = ocr.NewOCR(officialAccount.ctx)
	}
	return officialAccount.ocr
}

// GetSubscribe 公众号订阅消息
func (officialAccount *OfficialAccount) GetSubscribe() *message.Subscribe {
	if officialAccount.subscribeMsg == nil {
		officialAccount.subscribeMsg = message.NewSubscribe(officialAccount.ctx)
	}
	return officialAccount.subscribeMsg
}

// GetCustomerServiceManager 客服管理
func (officialAccount *OfficialAccount) GetCustomerServiceManager() *customerservice.Manager {
	return customerservice.NewCustomerServiceManager(officialAccount.ctx)
}
