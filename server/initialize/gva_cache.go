package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/gva_cache"
	"go.uber.org/zap"
)

// InitGvaCache 初始化通用缓存句柄 global.GVA_CACHE。
// 必须在 Redis 初始化之后调用：有 Redis 用 Redis 后端，否则用内存后端。
func InitGvaCache() {
	if global.GVA_REDIS != nil {
		global.GVA_CACHE = gva_cache.NewRedisCache(global.GVA_REDIS)
		global.GVA_LOG.Info("GVA_CACHE 使用 Redis 后端")
		return
	}
	dr, err := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
	if err != nil {
		// JWT 过期配置非法应在启动期暴露
		panic(err)
	}
	global.GVA_CACHE = gva_cache.NewMemoryCache(dr)
	global.GVA_LOG.Info("GVA_CACHE 使用内存后端", zap.Duration("defaultExpire", dr))
}
