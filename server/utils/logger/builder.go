package logger

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// callerSkip 跳过 Builder 的 wrapper 帧（log + 级别方法 = 2），使 caller 指向真实调用点。
// 若 TestBuilderCallerPointsToCallSite 失败，按报错里的 caller 调整此值。
const callerSkip = 2

type Builder struct {
	ctx       context.Context
	mod       string
	fields    []zap.Field
	detail    any
	hasDetail bool
	err       error
}

// WithCtx 请求链路入口
func WithCtx(ctx context.Context) *Builder { return &Builder{ctx: ctx} }

// Bg 无 ctx 兜底（启动/定时器/CLI/utils）
func Bg() *Builder { return &Builder{} }

func (b *Builder) Mod(mod string) *Builder { b.mod = mod; return b }

func (b *Builder) Field(key string, val any) *Builder {
	b.fields = append(b.fields, zap.Any(key, val))
	return b
}

func (b *Builder) Err(err error) *Builder { b.err = err; return b }

func (b *Builder) assemble() []zap.Field {
	f := make([]zap.Field, 0, len(b.fields)+10)
	f = append(f, b.fields...)
	f = append(f, zap.String(FieldMod, b.mod))

	fields := FromCtx(b.ctx)
	if fields == nil {
		fields = &Fields{}
	}
	f = append(f,
		zap.String(FieldRequestID, fields.RequestID),
		zap.String(FieldTraceID, fields.TraceID),
		zap.String(FieldDeviceID, fields.DeviceID),
		zap.String(FieldClientIP, fields.ClientIP),
		zap.String(FieldHTTPMethod, fields.HTTPMethod),
		zap.String(FieldHTTPPath, fields.HTTPPath),
	)
	// span 字段仅在有值时输出,避免非请求链路日志出现空字段
	if fields.SpanID != "" {
		f = append(f, zap.String(FieldSpanID, fields.SpanID))
	}
	if fields.ParentSpanID != "" {
		f = append(f, zap.String(FieldParentSpanID, fields.ParentSpanID))
	}
	if b.err != nil {
		f = append(f, zap.String(FieldErrorMsg, b.err.Error()), zap.Stack(FieldErrorStack))
	}
	if b.hasDetail {
		f = append(f, zap.Any(FieldDetail, b.detail))
	}
	return f
}

func (b *Builder) log(lvl zapcore.Level, msg string) {
	l := global.GVA_LOG.WithOptions(zap.AddCallerSkip(callerSkip))
	if ce := l.Check(lvl, msg); ce != nil {
		ce.Write(b.assemble()...)
	}
}

func (b *Builder) Info(msg string)  { b.log(zapcore.InfoLevel, msg) }
func (b *Builder) Warn(msg string)  { b.log(zapcore.WarnLevel, msg) }
func (b *Builder) Error(msg string) { b.log(zapcore.ErrorLevel, msg) }
func (b *Builder) Debug(msg string) { b.log(zapcore.DebugLevel, msg) }

// *Detail 直接调 log（不经 Info），保证 caller 深度与非 Detail 版本一致
func (b *Builder) InfoDetail(msg string, detail any) {
	b.detail, b.hasDetail = detail, true
	b.log(zapcore.InfoLevel, msg)
}
func (b *Builder) WarnDetail(msg string, detail any) {
	b.detail, b.hasDetail = detail, true
	b.log(zapcore.WarnLevel, msg)
}
func (b *Builder) ErrorDetail(msg string, detail any) {
	b.detail, b.hasDetail = detail, true
	b.log(zapcore.ErrorLevel, msg)
}
func (b *Builder) DebugDetail(msg string, detail any) {
	b.detail, b.hasDetail = detail, true
	b.log(zapcore.DebugLevel, msg)
}
