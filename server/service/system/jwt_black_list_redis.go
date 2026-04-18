package system

// 分布式 JWT 黑名单的 Redis 层。
//
// 历史：JwtService 只把黑名单写进 global.BlackCache (sync.Map)，
//       在多实例部署下，A 实例拉黑的 token 对 B 实例无效 —— 管理员的「踢下线」失效。
//
// 设计：
//   - DB (system.JwtBlacklist) 仍然是 source of truth，用于冷启动时的 LoadAll。
//   - Redis 是分布式 L2，JsonInBlacklist 写入时同步刷 Redis。
//   - BlackCache 是进程级 L1，命中即短路；未命中再查 Redis。
//   - Redis 未配置或短暂不可用时自动退化为 BlackCache-only（与未打此补丁时一致）。
//
// Key 设计：
//   gva:jwt:blacklist:<sha256(token)>   —— 固定长度 64 hex 字符，避免把 2KB+ 的 token
//   直接当 key；hash 对齐 Redis 的 key 大小最佳实践。
//
// TTL：
//   贴近 token 自身 exp；无法解析的 token 走保守默认值；已过期的 token 只写 1s
//   用来挡住并发 race；上限 30 天兜底。

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	jwtlib "github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

const (
	jwtBlacklistRedisPrefix  = "gva:jwt:blacklist:"
	jwtBlacklistDefaultTTL   = 24 * time.Hour
	jwtBlacklistMinTTL       = time.Second
	jwtBlacklistMaxTTL       = 30 * 24 * time.Hour
	jwtBlacklistReadTimeout  = 200 * time.Millisecond
	jwtBlacklistWriteTimeout = 500 * time.Millisecond
)

// ErrBlacklistRedisUnavailable 表示 Redis 不可用；调用方应退化到 BlackCache-only。
var ErrBlacklistRedisUnavailable = errors.New("jwt blacklist: redis unavailable")

func jwtBlacklistRedisKey(token string) string {
	sum := sha256.Sum256([]byte(token))
	return jwtBlacklistRedisPrefix + hex.EncodeToString(sum[:])
}

// deriveBlacklistTTL 不校验签名，只读 exp 算剩余寿命。
// ParseUnverified 是 jwt/v5 的合法 API —— 签名校验在上游鉴权链路已经完成，
// 此处只用来推断要给 Redis key 多长的 TTL。
func deriveBlacklistTTL(token string) time.Duration {
	parser := jwtlib.NewParser()
	claims := jwtlib.RegisteredClaims{}
	_, _, err := parser.ParseUnverified(token, &claims)
	if err != nil || claims.ExpiresAt == nil {
		return jwtBlacklistDefaultTTL
	}
	remaining := time.Until(claims.ExpiresAt.Time)
	if remaining <= 0 {
		return jwtBlacklistMinTTL
	}
	if remaining > jwtBlacklistMaxTTL {
		return jwtBlacklistMaxTTL
	}
	return remaining
}

// SetJWTBlacklistRedis 把 token 写入分布式黑名单。
//
//   - global.GVA_REDIS == nil（未配置 Redis）时返回 nil，让调用方无痛降级。
//   - 返回非 nil 表示 Redis 可用但写失败；调用方应记录 warn 但不应阻塞主流程
//     （DB 落盘成功即可；其他实例下次读路径会从 DB->Redis 走一遍）。
func SetJWTBlacklistRedis(ctx context.Context, token string) error {
	r := global.GVA_REDIS
	if r == nil {
		return nil
	}
	if ctx == nil {
		ctx = context.Background()
	}
	ctx, cancel := context.WithTimeout(ctx, jwtBlacklistWriteTimeout)
	defer cancel()
	ttl := deriveBlacklistTTL(token)
	return r.Set(ctx, jwtBlacklistRedisKey(token), "1", ttl).Err()
}

// IsJWTBlacklistedRedis 查询 Redis 黑名单。
//
//   - (true,  nil)                            token 在黑名单。
//   - (false, nil)                            token 不在黑名单。
//   - (false, ErrBlacklistRedisUnavailable)   Redis 未初始化 / 超时 / 连接错；
//     调用方应退化到 BlackCache-only（等价于打补丁之前的单实例行为）。
func IsJWTBlacklistedRedis(ctx context.Context, token string) (bool, error) {
	r := global.GVA_REDIS
	if r == nil {
		return false, ErrBlacklistRedisUnavailable
	}
	if ctx == nil {
		ctx = context.Background()
	}
	ctx, cancel := context.WithTimeout(ctx, jwtBlacklistReadTimeout)
	defer cancel()
	_, err := r.Get(ctx, jwtBlacklistRedisKey(token)).Result()
	if err == nil {
		return true, nil
	}
	if errors.Is(err, redis.Nil) {
		return false, nil
	}
	return false, ErrBlacklistRedisUnavailable
}

// HydrateRedisBlacklistFromDB 把 DB 里未过期的 JWT 黑名单记录批量回填到 Redis。
//
// 目的：
//   - 新节点启动时快速恢复分布式状态；
//   - Redis 经历重建 / 清空 / 故障恢复后也能自愈。
//
// 调用点：LoadAll() 之后异步触发。
// 幂等：同一 token 被 Set 多次仅覆盖 TTL，无副作用。
func HydrateRedisBlacklistFromDB(ctx context.Context, tokens []string) {
	r := global.GVA_REDIS
	if r == nil || len(tokens) == 0 {
		return
	}
	if ctx == nil {
		ctx = context.Background()
	}
	var okCount, failCount int
	for _, tok := range tokens {
		if tok == "" {
			continue
		}
		cctx, cancel := context.WithTimeout(ctx, jwtBlacklistWriteTimeout)
		err := r.Set(cctx, jwtBlacklistRedisKey(tok), "1", deriveBlacklistTTL(tok)).Err()
		cancel()
		if err != nil {
			failCount++
			continue
		}
		okCount++
	}
	if global.GVA_LOG != nil {
		global.GVA_LOG.Info("jwt blacklist hydrated to redis",
			zap.Int("ok", okCount),
			zap.Int("fail", failCount),
			zap.Int("total", len(tokens)),
		)
	}
}
