package testutil

import (
	"context"
	"testing"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/redis/go-redis/v9"
)

// DefaultRedisAddr 是 NewRedisOrSkip 默认连接的 Redis 地址。
const DefaultRedisAddr = "127.0.0.1:6379"

// DefaultRedisPingTimeout 是 ping 检测的默认超时，避免本地无 Redis 时阻塞 CI。
const DefaultRedisPingTimeout = 200 * time.Millisecond

// RedisOption 配置 NewRedisOrSkip 的可选行为。
type RedisOption func(*redisConfig)

type redisConfig struct {
	addr        string
	username    string
	password    string
	db          int
	pingTimeout time.Duration
	assignGlobal bool // 连上后赋值到 global.GVA_REDIS（并在 cleanup 还原）
}

// WithRedisAddr 自定义 Redis 地址（默认 127.0.0.1:6379）。
func WithRedisAddr(addr string) RedisOption {
	return func(c *redisConfig) {
		if addr != "" {
			c.addr = addr
		}
	}
}

// WithRedisCredentials 自定义用户名/密码。
func WithRedisCredentials(username, password string) RedisOption {
	return func(c *redisConfig) { c.username, c.password = username, password }
}

// WithRedisDB 自定义 db index。
func WithRedisDB(db int) RedisOption {
	return func(c *redisConfig) { c.db = db }
}

// WithRedisPingTimeout 自定义 ping 超时。
func WithRedisPingTimeout(d time.Duration) RedisOption {
	return func(c *redisConfig) {
		if d > 0 {
			c.pingTimeout = d
		}
	}
}

// WithAssignGlobal 让 NewRedisOrSkip 在连接成功后把 client 赋值到 global.GVA_REDIS，
// 并在 t.Cleanup 中还原旧值。默认不赋值，仅返回 client，避免污染其它测试。
func WithAssignGlobal() RedisOption {
	return func(c *redisConfig) { c.assignGlobal = true }
}

// NewRedisOrSkip 尝试连接 Redis；连不上则 t.Skipf（不阻塞 CI / 本地无 Redis 环境）。
// 返回新建的 redis.UniversalClient 供调用方直接使用。
//
// 复用 utils/gva_cache/redis_cache_test.go 中的 ping-or-skip 模式并泛化为通用 helper。
//
// 用法：
//
//	func TestFoo(t *testing.T) {
//	    client := testutil.NewRedisOrSkip(t)
//	    _ = client
//	}
//	// 自定义地址 + 同步全局：
//	client := testutil.NewRedisOrSkip(t,
//	    testutil.WithRedisAddr("redis:6379"),
//	    testutil.WithAssignGlobal(),
//	)
func NewRedisOrSkip(t testing.TB, opts ...RedisOption) redis.UniversalClient {
	t.Helper()
	cfg := redisConfig{
		addr:        DefaultRedisAddr,
		pingTimeout: DefaultRedisPingTimeout,
	}
	for _, opt := range opts {
		opt(&cfg)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.addr,
		Username: cfg.username,
		Password: cfg.password,
		DB:       cfg.db,
	})

	ctx, cancel := context.WithTimeout(context.Background(), cfg.pingTimeout)
	defer cancel()
	if err := client.Ping(ctx).Err(); err != nil {
		_ = client.Close()
		t.Skipf("testutil: 无可用 Redis (%s), 跳过测试: %v", cfg.addr, err)
		return nil
	}

	if cfg.assignGlobal {
		old := global.GVA_REDIS
		global.GVA_REDIS = client
		t.Cleanup(func() {
			global.GVA_REDIS = old
			_ = client.Close()
		})
	} else {
		t.Cleanup(func() { _ = client.Close() })
	}
	return client
}
