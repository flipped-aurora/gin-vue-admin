package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/core/cache"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

func Cache() {
	c, err := cache.Create()
	if err != nil {
		panic(err)
	}
	global.GVA_CACHE = c
}
