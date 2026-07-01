package logger

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func setupObserver(t *testing.T) *observer.ObservedLogs {
	t.Helper()
	core, logs := observer.New(zapcore.DebugLevel)
	global.GVA_LOG = zap.New(core, zap.AddCaller())
	return logs
}

func TestBuilderEnvelopeFromCtx(t *testing.T) {
	logs := setupObserver(t)
	ctx := WithFields(context.Background(), &Fields{RequestID: "rid-9", ClientIP: "9.9.9.9"})
	WithCtx(ctx).Mod("biz").Info("hello")

	if logs.Len() != 1 {
		t.Fatalf("want 1 log, got %d", logs.Len())
	}
	e := logs.All()[0]
	m := e.ContextMap()
	if e.Message != "hello" || m["mod"] != "biz" || m["request_id"] != "rid-9" || m["client_ip"] != "9.9.9.9" {
		t.Fatalf("bad fields: msg=%s map=%+v", e.Message, m)
	}
}

func TestBuilderErrAndDetail(t *testing.T) {
	logs := setupObserver(t)
	WithCtx(context.Background()).Mod("biz").Err(errors.New("boom")).ErrorDetail("failed", map[string]int{"n": 1})
	e := logs.All()[0]
	m := e.ContextMap()
	if m["error_msg"] != "boom" || m["detail"] == nil || e.Level != zapcore.ErrorLevel {
		t.Fatalf("bad err/detail: %+v lvl=%s", m, e.Level)
	}
}

func TestBuilderBgEmptyEnvelope(t *testing.T) {
	logs := setupObserver(t)
	Bg().Mod("system").Info("startup")
	m := logs.All()[0].ContextMap()
	if m["request_id"] != "" || m["mod"] != "system" {
		t.Fatalf("bg envelope should be empty request_id: %+v", m)
	}
}

func TestBuilderCallerPointsToCallSite(t *testing.T) {
	logs := setupObserver(t)
	WithCtx(context.Background()).Mod("biz").Info("x") // 本行即调用点
	caller := logs.All()[0].Caller.String()
	if !strings.Contains(caller, "builder_test.go") {
		t.Fatalf("caller should point to test file, got %q (调整 callerSkip)", caller)
	}
}
