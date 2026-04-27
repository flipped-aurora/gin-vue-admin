package system

import (
	"context"
	"strconv"
	"sync/atomic"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	jwtlib "github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/songzhibin97/gkit/cache/local_cache"
)

var testJWTCounter atomic.Uint64

// JWT 黑名单的 Redis 层单元测试：
//   - 写路径（SetJWTBlacklistRedis）和读路径（IsJWTBlacklistedRedis）的 round-trip。
//   - Redis 故障时的 fail-open 降级。
//   - IsInBlacklist 分层查询的命中顺序。
//   - TTL 派生逻辑（正常 token / 无法解析的 token / 已过期 token）。
//   - HydrateRedisBlacklistFromDB 的批量回填。
//
// 约束：不依赖真实 MySQL（JsonInBlacklist 写 DB 不在本文件覆盖范围），
//       只验证 "Redis 层 + BlackCache 层" 的组合行为。

func setupRedisBlacklistTest(t *testing.T) (*miniredis.Miniredis, func()) {
	t.Helper()
	srv, err := miniredis.Run()
	if err != nil {
		t.Fatalf("miniredis: %v", err)
	}
	rdb := redis.NewClient(&redis.Options{Addr: srv.Addr()})

	origRedis := global.GVA_REDIS
	origCache := global.BlackCache
	global.GVA_REDIS = rdb
	global.BlackCache = local_cache.NewCache(local_cache.SetDefaultExpire(5 * time.Second))

	cleanup := func() {
		_ = rdb.Close()
		srv.Close()
		global.GVA_REDIS = origRedis
		global.BlackCache = origCache
	}
	return srv, cleanup
}

func signTestJWT(t *testing.T, ttl time.Duration) string {
	t.Helper()
	token := jwtlib.NewWithClaims(jwtlib.SigningMethodHS256, jwtlib.RegisteredClaims{
		ExpiresAt: jwtlib.NewNumericDate(time.Now().Add(ttl)),
		ID:        "test-" + strconv.FormatUint(testJWTCounter.Add(1), 10),
	})
	signed, err := token.SignedString([]byte("test-secret"))
	if err != nil {
		t.Fatalf("sign jwt: %v", err)
	}
	return signed
}

func TestDeriveBlacklistTTL_UsesJWTExp(t *testing.T) {
	tok := signTestJWT(t, 1*time.Hour)
	got := deriveBlacklistTTL(tok)
	if got < 59*time.Minute || got > 61*time.Minute {
		t.Errorf("ttl from 1h-exp jwt want ~1h, got %s", got)
	}
}

func TestDeriveBlacklistTTL_BadToken(t *testing.T) {
	got := deriveBlacklistTTL("not-a-real-jwt")
	if got != jwtBlacklistDefaultTTL {
		t.Errorf("unparseable token ttl want default %s, got %s", jwtBlacklistDefaultTTL, got)
	}
}

func TestDeriveBlacklistTTL_AlreadyExpired(t *testing.T) {
	tok := signTestJWT(t, -1*time.Hour)
	got := deriveBlacklistTTL(tok)
	if got != jwtBlacklistMinTTL {
		t.Errorf("expired token ttl want min %s, got %s", jwtBlacklistMinTTL, got)
	}
}

func TestSetAndIsBlacklistedRedis(t *testing.T) {
	_, cleanup := setupRedisBlacklistTest(t)
	defer cleanup()

	tok := signTestJWT(t, 1*time.Hour)
	ctx := context.Background()

	if hit, err := IsJWTBlacklistedRedis(ctx, tok); err != nil || hit {
		t.Fatalf("precondition: token not yet blacklisted, got hit=%v err=%v", hit, err)
	}
	if err := SetJWTBlacklistRedis(ctx, tok); err != nil {
		t.Fatalf("Set: %v", err)
	}
	hit, err := IsJWTBlacklistedRedis(ctx, tok)
	if err != nil {
		t.Fatalf("Is: %v", err)
	}
	if !hit {
		t.Error("expected token to be blacklisted after Set")
	}
}

func TestIsJWTBlacklistedRedis_NoRedis(t *testing.T) {
	orig := global.GVA_REDIS
	global.GVA_REDIS = nil
	defer func() { global.GVA_REDIS = orig }()

	hit, err := IsJWTBlacklistedRedis(context.Background(), "any-token")
	if hit {
		t.Error("expected hit=false when Redis is nil")
	}
	if err == nil || err != ErrBlacklistRedisUnavailable {
		t.Errorf("want ErrBlacklistRedisUnavailable, got %v", err)
	}
}

func TestJwtService_IsInBlacklist_Layered(t *testing.T) {
	_, cleanup := setupRedisBlacklistTest(t)
	defer cleanup()

	svc := &JwtService{}
	ctx := context.Background()

	tokL1 := signTestJWT(t, 1*time.Hour)
	global.BlackCache.SetDefault(tokL1, struct{}{})
	if !svc.IsInBlacklist(ctx, tokL1) {
		t.Error("L1 BlackCache hit should short-circuit to true")
	}

	tokL2 := signTestJWT(t, 1*time.Hour)
	if err := SetJWTBlacklistRedis(ctx, tokL2); err != nil {
		t.Fatal(err)
	}
	if !svc.IsInBlacklist(ctx, tokL2) {
		t.Error("Redis hit should return true via IsInBlacklist")
	}
	if _, ok := global.BlackCache.Get(tokL2); !ok {
		t.Error("Redis hit should backfill BlackCache L1")
	}

	tokClean := signTestJWT(t, 1*time.Hour)
	if svc.IsInBlacklist(ctx, tokClean) {
		t.Error("unknown token must not be reported as blacklisted")
	}

	if svc.IsInBlacklist(ctx, "") {
		t.Error("empty token must not be reported as blacklisted")
	}
}

func TestJwtService_IsInBlacklist_RedisDownFailsOpen(t *testing.T) {
	origCache := global.BlackCache
	global.BlackCache = local_cache.NewCache(local_cache.SetDefaultExpire(5 * time.Second))
	defer func() { global.BlackCache = origCache }()

	origRedis := global.GVA_REDIS
	global.GVA_REDIS = nil
	defer func() { global.GVA_REDIS = origRedis }()

	svc := &JwtService{}
	tok := signTestJWT(t, 1*time.Hour)

	global.BlackCache.SetDefault(tok, struct{}{})
	if !svc.IsInBlacklist(context.Background(), tok) {
		t.Error("with Redis down, L1 hit must still short-circuit")
	}

	other := signTestJWT(t, 1*time.Hour)
	if svc.IsInBlacklist(context.Background(), other) {
		t.Error("with Redis down and L1 miss, default is fail-open (false)")
	}
}

func TestHydrateRedisBlacklistFromDB(t *testing.T) {
	_, cleanup := setupRedisBlacklistTest(t)
	defer cleanup()

	toks := []string{
		signTestJWT(t, 1*time.Hour),
		signTestJWT(t, 2*time.Hour),
		"",
	}
	HydrateRedisBlacklistFromDB(context.Background(), toks)

	for i, tok := range toks {
		if tok == "" {
			continue
		}
		hit, err := IsJWTBlacklistedRedis(context.Background(), tok)
		if err != nil {
			t.Errorf("token[%d] Is err: %v", i, err)
		}
		if !hit {
			t.Errorf("token[%d] should be hydrated into Redis", i)
		}
	}
}
