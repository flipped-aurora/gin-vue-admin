package system

import (
	"context"

	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type JwtService struct{}

var JwtServiceApp = new(JwtService)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: JsonInBlacklist
//@description: 拉黑jwt
//@param: jwtList model.JwtBlacklist
//@return: err error

// 写路径：DB + Redis + BlackCache 三写。
//   - DB 是 source of truth（持久化）。
//   - Redis 是分布式 L2（让多实例立刻一致）。
//   - BlackCache 是本实例 L1（省一次 Redis 往返）。
//
// Redis 未配置 / 写失败都不阻塞流程：DB 已落盘；其他实例读路径命中 DB
// 重新同步 Redis 即可。
func (jwtService *JwtService) JsonInBlacklist(jwtList system.JwtBlacklist) (err error) {
	err = global.GVA_DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	if rerr := SetJWTBlacklistRedis(context.Background(), jwtList.Jwt); rerr != nil && global.GVA_LOG != nil {
		global.GVA_LOG.Warn("jwt blacklist: redis write failed, falling back to BlackCache only",
			zap.Error(rerr))
	}
	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetRedisJWT
//@description: 从redis取jwt
//@param: userName string
//@return: redisJWT string, err error

func (jwtService *JwtService) GetRedisJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = global.GVA_REDIS.Get(context.Background(), userName).Result()
	return redisJWT, err
}

// IsInBlacklist 分层查询 JWT 黑名单。
//
//   - L1 BlackCache 命中 → true（进程内即判定，零 IO）。
//   - L1 未命中且 Redis 可用 → 查 Redis；命中后顺手回填 L1。
//   - Redis 未配置或不可用 → 退化为 BlackCache-only（等价于历史 sync.Map 行为）。
//
// Fail-open 原则：Redis 抖动不得阻塞鉴权。最坏情况是被拉黑的 token 在
// 某节点多活一小段时间，直到该节点下次 LoadAll / HydrateRedisBlacklistFromDB。
func (jwtService *JwtService) IsInBlacklist(ctx context.Context, token string) bool {
	if token == "" {
		return false
	}
	if _, ok := global.BlackCache.Get(token); ok {
		return true
	}
	hit, err := IsJWTBlacklistedRedis(ctx, token)
	if err != nil {
		return false
	}
	if hit {
		global.BlackCache.SetDefault(token, struct{}{})
		return true
	}
	return false
}

// LoadAll 启动时从 DB 加载 JWT 黑名单。
//   - 全量灌进 BlackCache（本实例 L1，保持旧行为）。
//   - 异步回填 Redis（分布式 L2，新节点启动时迅速对齐其他节点）。
func LoadAll() {
	var data []string
	err := global.GVA_DB.Model(&system.JwtBlacklist{}).Select("jwt").Find(&data).Error
	if err != nil {
		global.GVA_LOG.Error("加载数据库jwt黑名单失败!", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		global.BlackCache.SetDefault(data[i], struct{}{})
	} // jwt黑名单 加入 BlackCache 中
	go HydrateRedisBlacklistFromDB(context.Background(), data)
}
