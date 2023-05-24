// Package config 小程序config配置
package config

import (
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/cache"
)

// Config .config for 小程序
type Config struct {
	AppID     string `json:"app_id"`     // appid
	AppSecret string `json:"app_secret"` // appSecret
	Cache     cache.Cache
}
