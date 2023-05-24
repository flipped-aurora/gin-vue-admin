package config

import "github.com/flipped-aurora/gin-vue-admin/server/douyin/open/cache"

// Config config for 抖音开放平台
type Config struct {
	ClientKey    string
	ClientSecret string
	RedirectURL  string
	Scopes       string
	Cache        cache.Cache
}
