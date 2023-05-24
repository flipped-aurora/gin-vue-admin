package account

import "github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/openplatform/context"

// Account 开放平台张哈管理
// TODO 实现方法
type Account struct {
	*context.Context
}

// NewAccount new
func NewAccount(ctx *context.Context) *Account {
	return &Account{ctx}
}

// Create 创建开放平台帐号并绑定公众号/小程序
func (account *Account) Create(appID string) (string, error) {
	return "", nil
}

// Bind 将公众号/小程序绑定到开放平台帐号下
func (account *Account) Bind(appID string) error {
	return nil
}

// Unbind 将公众号/小程序从开放平台帐号下解绑
func (account *Account) Unbind(appID string, openAppID string) error {
	return nil
}

// Get 获取公众号/小程序所绑定的开放平台帐号
func (account *Account) Get(appID string) (string, error) {
	return "", nil
}
