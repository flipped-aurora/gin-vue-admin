package internal

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
	"time"
)

type ZapLogger struct {
	log *zap.Logger
}

func NewZapLogger() logger.Interface {
	return &ZapLogger{
		log: global.GVA_LOG,
	}
}

func (l *ZapLogger) LogMode(level logger.LogLevel) logger.Interface {
	// 在这里，我们忽略了LogMode的设置，因为zap的日志级别在创建zap logger时已经设置了
	return l
}

func (l *ZapLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	l.log.Info(fmt.Sprintf(msg, data...))
}

func (l *ZapLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.log.Warn(fmt.Sprintf(msg, data...))
}

func (l *ZapLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	l.log.Error(fmt.Sprintf(msg, data...))
}

func (l *ZapLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	switch {
	case err != nil && l.log.Core().Enabled(zap.ErrorLevel):
		sql, _ := fc()
		l.log.Error("trace", zap.String("elapsed", elapsed.String()), zap.String("sql", sql), zap.Error(err))
	case elapsed > 200*time.Millisecond && l.log.Core().Enabled(zap.WarnLevel):
		sql, _ := fc()
		l.log.Warn("trace", zap.String("elapsed", elapsed.String()), zap.String("sql", sql))
	case l.log.Core().Enabled(zap.InfoLevel):
		sql, _ := fc()
		l.log.Info("trace", zap.String("elapsed", elapsed.String()), zap.String("sql", sql))
	}
}
