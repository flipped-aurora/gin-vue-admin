package testutil

import (
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
)

func TestInitNopLogger_SetsWhenNil(t *testing.T) {
	// 保存并清理全局状态
	old := global.GVA_LOG
	global.GVA_LOG = nil
	t.Cleanup(func() { global.GVA_LOG = old })

	logger := InitNopLogger()
	if logger == nil {
		t.Fatal("InitNopLogger 返回 nil")
	}
	if global.GVA_LOG == nil {
		t.Fatal("InitNopLogger 未设置 global.GVA_LOG")
	}
	if global.GVA_LOG != logger {
		t.Fatal("global.GVA_LOG 与返回值不一致")
	}
}

func TestInitNopLogger_Idempotent(t *testing.T) {
	old := global.GVA_LOG
	t.Cleanup(func() { global.GVA_LOG = old })

	// 先放一个非 nil 的 logger（不一定是 nop）
	first := zap.NewNop()
	global.GVA_LOG = first

	got := InitNopLogger()
	if got != first {
		t.Fatal("InitNopLogger 在 GVA_LOG 非 nil 时不应覆盖")
	}
}
