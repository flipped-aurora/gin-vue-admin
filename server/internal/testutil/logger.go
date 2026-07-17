package testutil

import (
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
)

// InitNopLogger 当 global.GVA_LOG 为 nil 时，设为 zap.NewNop()，
// 并通过 zap.ReplaceGlobals 同步到 zap 的全局 logger。
// 已设置（非 nil）则不覆盖，保证幂等。返回最终使用的 logger。
//
// 许多代码路径在告警/无身份场景下会经由 GVA_LOG 输出日志，nil 会触发 panic。
// 在测试中通常用 nop logger 兜底，避免依赖文件系统路径（core.Zap() 需要合法
// 的 Zap 配置和日志目录）。
//
// 用法：
//
//	func TestFoo(t *testing.T) {
//	    testutil.InitNopLogger()
//	    // ...
//	}
//
// 注意：本函数会修改全局 GVA_LOG 与 zap 全局 logger，且不还原。
// 这是有意的：nop logger 对其它测试无副作用，全局只应初始化一次。
// 如需还原，请在调用方自行保存旧值并 t.Cleanup。
func InitNopLogger(t ...testing.TB) *zap.Logger {
	_ = t // 接受可选的 testing.TB 仅为 t.Helper()/future 用途，不强制要求
	if len(t) > 0 {
		t[0].Helper()
	}
	if global.GVA_LOG == nil {
		logger := zap.NewNop()
		global.GVA_LOG = logger
		zap.ReplaceGlobals(logger)
	}
	return global.GVA_LOG
}
