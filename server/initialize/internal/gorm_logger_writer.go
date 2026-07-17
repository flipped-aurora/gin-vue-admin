package internal

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/config"
	applog "github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
	"gorm.io/gorm/logger"
)

// GormLogger 实现 gorm logger.Interface，把 SQL 日志接到链式 logger（带 request_id）
type GormLogger struct {
	slowThreshold time.Duration
	level         logger.LogLevel
}

func NewGormLogger(cfg config.GeneralDB) *GormLogger {
	return &GormLogger{slowThreshold: 200 * time.Millisecond, level: cfg.LogLevel()}
}

func (g *GormLogger) LogMode(l logger.LogLevel) logger.Interface {
	ng := *g
	ng.level = l
	return &ng
}

func (g *GormLogger) Info(ctx context.Context, msg string, data ...any) {
	applog.WithCtx(ctx).Mod("sql").Info(fmt.Sprintf(msg, data...))
}

func (g *GormLogger) Warn(ctx context.Context, msg string, data ...any) {
	applog.WithCtx(ctx).Mod("sql").Warn(fmt.Sprintf(msg, data...))
}

func (g *GormLogger) Error(ctx context.Context, msg string, data ...any) {
	applog.WithCtx(ctx).Mod("sql").Error(fmt.Sprintf(msg, data...))
}

func (g *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if g.level <= logger.Silent {
		return
	}
	elapsed := time.Since(begin)
	sql, rows := fc()
	b := applog.WithCtx(ctx).Mod("sql").
		Field("sql", sql).
		Field("rows", rows).
		Field("elapsed_ms", elapsed.Milliseconds())
	switch {
	case err != nil && g.level >= logger.Error && !errors.Is(err, logger.ErrRecordNotFound):
		b.Err(err).Error("SQL 执行错误")
	case elapsed > g.slowThreshold && g.level >= logger.Warn:
		b.Warn("SQL 慢查询")
	case g.level >= logger.Info:
		b.Info("SQL")
	}
}
