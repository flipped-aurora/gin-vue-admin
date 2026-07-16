package internal

import (
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// With 必须保留 *ZapCore 包装:zap 的 logger.With 会用返回值替换原 core,
// 一旦返回内层 core,Write 的错误入库/控制台路由/business 子目录路由全部旁路
// (回归自 core/zap.go 的 logger.With(node/app_id/env) 调用)。
func TestZapCoreWithKeepsWrapper(t *testing.T) {
	z := NewZapCore(zapcore.ErrorLevel)
	w := z.With([]zapcore.Field{zap.String("node", "n1")})
	zc, ok := w.(*ZapCore)
	if !ok {
		t.Fatalf("With should return *ZapCore, got %T", w)
	}
	if len(zc.fields) != 1 || zc.fields[0].Key != "node" {
		t.Fatalf("With fields not retained: %+v", zc.fields)
	}
	if zc.level != zapcore.ErrorLevel {
		t.Fatalf("With should keep level, got %v", zc.level)
	}

	// 链式 With 累积,且不污染原实例
	w2 := zc.With([]zapcore.Field{zap.String("env", "prod")})
	if zc2 := w2.(*ZapCore); len(zc2.fields) != 2 {
		t.Fatalf("chained With should accumulate fields: %+v", zc2.fields)
	}
	if len(zc.fields) != 1 {
		t.Fatalf("With must not mutate receiver: %+v", zc.fields)
	}
}
