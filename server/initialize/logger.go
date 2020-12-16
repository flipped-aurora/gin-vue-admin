package initialize

import (
	"context"
	"fmt"
	"gin-vue-admin/global"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var (
	Discard = New(log.New(ioutil.Discard, "", log.LstdFlags), log.New(ioutil.Discard, "", log.LstdFlags), GormConfig{})
	Default = New(log.New(os.Stdout, "\r\n", log.LstdFlags), log.New(os.Stdout, "\r\n", log.LstdFlags), GormConfig{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})
	Recorder = traceRecorder{Interface: Default, BeginAt: time.Now()}
)

type traceRecorder struct {
	logger.Interface
	BeginAt      time.Time
	SQL          string
	RowsAffected int64
	Err          error
}

func New(writer Writer, gormWriter logger.Writer, config GormConfig) logger.Interface {
	var (
		infoStr      = "%s\n[info] "
		warnStr      = "%s\n[warn] "
		errStr       = "%s\n[error] "
		traceStr     = "%s\n[%.3fms] [rows:%v] %s"
		traceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s"
		traceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s"
	)

	if config.Colorful {
		infoStr = logger.Green + "%s\n" + logger.Reset + logger.Green + "[info] " + logger.Reset
		warnStr = logger.BlueBold + "%s\n" + logger.Reset + logger.Magenta + "[warn] " + logger.Reset
		errStr = logger.Magenta + "%s\n" + logger.Reset + logger.Red + "[error] " + logger.Reset
		traceStr = logger.Green + "%s\n" + logger.Reset + logger.Yellow + "[%.3fms] " + logger.BlueBold + "[rows:%v]" + logger.Reset + " %s"
		traceWarnStr = logger.Green + "%s " + logger.Yellow + "%s\n" + logger.Reset + logger.RedBold + "[%.3fms] " + logger.Yellow + "[rows:%v]" + logger.Magenta + " %s" + logger.Reset
		traceErrStr = logger.RedBold + "%s " + logger.MagentaBold + "%s\n" + logger.Reset + logger.Yellow + "[%.3fms] " + logger.BlueBold + "[rows:%v]" + logger.Reset + " %s"
	}

	return &GormLogger{
		Writer:       writer,
		gormWriter:   gormWriter,
		GormConfig:   config,
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
	}
}

// Writer log writer interface
type Writer interface {
	Printf(string, ...interface{})
}

type GormConfig struct {
	SlowThreshold time.Duration
	Colorful      bool
	LogLevel      logger.LogLevel
}

type GormLogger struct {
	Writer
	gormWriter logger.Writer
	GormConfig
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
}

func (g *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *g
	newLogger.LogLevel = level
	return &newLogger
}

func (g *GormLogger) Info(ctx context.Context, message string, data ...interface{}) {
	if g.LogLevel >= logger.Info {
		g.Printf(g.infoStr+message, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

func (g *GormLogger) Warn(ctx context.Context, message string, data ...interface{}) {
	if g.LogLevel >= logger.Warn {
		g.Printf(g.warnStr+message, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

func (g *GormLogger) Error(ctx context.Context, message string, data ...interface{}) {
	if g.LogLevel >= logger.Error {
		g.Printf(g.errStr+message, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

func (g *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if g.LogLevel > 0 {
		elapsed := time.Since(begin)
		switch {
		case err != nil && g.LogLevel >= logger.Error:
			sql, rows := fc()
			if rows == -1 {
				g.Printf(g.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				g.Printf(g.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		case elapsed > g.SlowThreshold && g.SlowThreshold != 0 && g.LogLevel >= logger.Warn:
			sql, rows := fc()
			slowLog := fmt.Sprintf("SLOW SQL >= %v", g.SlowThreshold)
			if rows == -1 {
				g.Printf(g.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				g.Printf(g.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		case g.LogLevel >= logger.Info:
			sql, rows := fc()
			if rows == -1 {
				g.Printf(g.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				g.Printf(g.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		}
	}
}

func (g *GormLogger) Printf(message string, data ...interface{}) {
	if global.GVA_CONFIG.Mysql.LogZap == "Info" && !global.GVA_CONFIG.Mysql.LogMode {
		switch len(data) {
		case 0:
			global.GVA_LOG.Info(message)
		case 1:
			global.GVA_LOG.Info("gorm", zap.Any("src", data[0]))
		case 2:
			global.GVA_LOG.Info("gorm", zap.Any("src", data[0]), zap.Any("duration", data[1]))
		case 3:
			global.GVA_LOG.Info("gorm", zap.Any("src", data[0]), zap.Any("duration", data[1]), zap.Any("rows", data[2]))
		case 4:
			global.GVA_LOG.Info("gorm", zap.Any("src", data[0]), zap.Any("duration", data[1]), zap.Any("rows", data[2]), zap.Any("sql", data[3]))
		}
		return
	}
	switch len(data) {
	case 0:
		g.gormWriter.Printf(message, "")
	case 1:
		g.gormWriter.Printf(message, data[0].(string))
	case 2:
		g.gormWriter.Printf(message, data[0].(string), data[1].(float64))
	case 3:
		g.gormWriter.Printf(message, data[0].(string), data[1].(float64), data[2].(string))
	case 4:
		g.gormWriter.Printf(message, data[0].(string), data[1].(float64), data[2].(string), data[3].(string))
	}
}
