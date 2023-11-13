package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/core/cache"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

func InitCache() {
	c, err := cache.Create(global.GVA_CONFIG.Cache)
	if err != nil {
		panic(err)
	}
	global.GVA_CACHE = c
}
