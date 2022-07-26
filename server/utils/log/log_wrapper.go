package log

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common"
	"go.uber.org/zap"
)

type LogWrapper struct {
	ZapLogger *zap.Logger
}

var GvaLog *LogWrapper

func (gvaLog *LogWrapper) Debug(tag string, fields ...zap.Field) {
	GvaLog.ZapLogger.Debug(tag, fields...)
}

func (gvaLog *LogWrapper) DebugF(ctx context.Context, tag string, fields ...zap.Field) {
	trace := ctx.Value(common.TraceKey).(*common.Trace)
	GvaLog.ZapLogger.Debug(tag,
		append(fields, zap.String("TraceId", trace.TraceId), zap.Int("UserId", trace.UserId))...,
	)
}

func (gvaLog *LogWrapper) Info(tag string, fields ...zap.Field) {
	GvaLog.ZapLogger.Info(tag, fields...)
}

func (gvaLog *LogWrapper) InfoF(ctx context.Context, tag string, fields ...zap.Field) {
	trace := ctx.Value(common.TraceKey).(*common.Trace)
	GvaLog.ZapLogger.Info(tag,
		append(fields, zap.String("TraceId", trace.TraceId), zap.Int("UserId", trace.UserId))...,
	)
}

func (gvaLog *LogWrapper) Warn(tag string, fields ...zap.Field) {
	GvaLog.ZapLogger.Warn(tag, fields...)
}

func (gvaLog *LogWrapper) WarnF(ctx context.Context, tag string, fields ...zap.Field) {
	trace := ctx.Value(common.TraceKey).(*common.Trace)
	GvaLog.ZapLogger.Warn(tag,
		append(fields, zap.String("TraceId", trace.TraceId), zap.Int("UserId", trace.UserId))...,
	)
}

func (gvaLog *LogWrapper) Error(tag string, fields ...zap.Field) {
	GvaLog.ZapLogger.Error(tag, fields...)
}

func (gvaLog *LogWrapper) ErrorF(ctx context.Context, tag string, fields ...zap.Field) {
	trace := ctx.Value(common.TraceKey).(*common.Trace)
	GvaLog.ZapLogger.Error(tag,
		append(fields, zap.String("TraceId", trace.TraceId), zap.Int("UserId", trace.UserId))...,
	)
}

func (gvaLog *LogWrapper) Fatal(tag string, fields ...zap.Field) {
	GvaLog.ZapLogger.Fatal(tag, fields...)
}

func (gvaLog *LogWrapper) FatalF(ctx context.Context, tag string, fields ...zap.Field) {
	trace := ctx.Value(common.TraceKey).(*common.Trace)
	GvaLog.ZapLogger.Fatal(tag,
		append(fields, zap.String("TraceId", trace.TraceId), zap.Int("UserId", trace.UserId))...,
	)
}
