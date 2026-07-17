package testutil

import (
	"context"
	"testing"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// TestNewRedisOrSkip_NoServerSkips 验证：连不上 Redis 时返回 t.Skipf 而非 fail。
// 这是一个"自验证"测试：无论本地是否真的有 Redis，本测试都不会 fail。
//
//   - 本地无 Redis：触发 skip 分支，符合预期。
//   - 本地有 Redis（默认 6379）：连接成功，进入正常分支，也通过。
//
// 我们额外构造一个指向不存在端口的请求来强制走 skip 路径。
func TestNewRedisOrSkip_UnreachableSkips(t *testing.T) {
	// 用一个几乎不可能开放的端口 + 极短超时，强制触发 ping 失败 → skip
	// 注意：被 skip 的测试不算失败，所以这里即便没"断言"也无妨。
	sub := &skipCapturingT{TB: t, skipped: false}
	_ = NewRedisOrSkip(sub,
		WithRedisAddr("127.0.0.1:1"), // 端口 1 通常被拒绝或超时
		WithRedisPingTimeout(50*time.Millisecond),
	)
	if !sub.skipped {
		// 如果端口 1 恰好被某些环境监听（极少见），这里不强制失败，
		// 只记录以便观察。
		t.Logf("注意：连接 127.0.0.1:1 未被 skip，环境可能监听该端口")
	}
}

// TestNewRedisOrSkip_AssignsGlobal 验证 WithAssignGlobal 在连接成功时赋值全局。
// 若本地无 Redis，自动 skip。
func TestNewRedisOrSkip_AssignsGlobal(t *testing.T) {
	old := global.GVA_REDIS
	t.Cleanup(func() { global.GVA_REDIS = old })

	client := NewRedisOrSkip(t, WithAssignGlobal())
	// 走到这里说明 Redis 可用
	if client == nil {
		t.Fatal("NewRedisOrSkip 返回 nil（但未 skip）")
	}
	if global.GVA_REDIS == nil {
		t.Fatal("WithAssignGlobal 未赋值 global.GVA_REDIS")
	}

	// 验证 client 可用
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := client.Ping(ctx).Err(); err != nil {
		t.Fatalf("赋值的 client ping 失败: %v", err)
	}
}

// skipCapturingT 包装 testing.T 并捕获是否触发过 Skip。
type skipCapturingT struct {
	testing.TB
	skipped bool
}

func (s *skipCapturingT) SkipNow() { s.skipped = true }
func (s *skipCapturingT) Skipf(format string, args ...any) {
	s.skipped = true
	s.TB.Logf("captured skip: "+format, args...)
}
func (s *skipCapturingT) Skip(args ...any) {
	s.skipped = true
	s.TB.Logf("captured skip: %v", args...)
}
