package internal

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

type config struct {
	SlowThreshold time.Duration
	Colorful      bool
	LogLevel      logger.LogLevel
}

var (
	Discard = New(log.New(ioutil.Discard, "", log.LstdFlags), config{})
	Default = New(log.New(os.Stdout, "\r\n", log.LstdFlags), config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})
	Recorder = traceRecorder{Interface: Default, BeginAt: time.Now()}
)

func New(writer logger.Writer, config config) logger.Interface {
	var (
		infoStr      = "%s\n[info] "
		warnStr      = "%s\n[warn] "
		errStr       = "%s\n[error] "
		traceStr     = "%s\n[%.3fms] [rows:%v] %s\n"
		traceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s\n"
		traceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s\n"
	)

	if config.Colorful {
		infoStr = logger.Green + "%s\n" + logger.Reset + logger.Green + "[info] " + logger.Reset
		warnStr = logger.BlueBold + "%s\n" + logger.Reset + logger.Magenta + "[warn] " + logger.Reset
		errStr = logger.Magenta + "%s\n" + logger.Reset + logger.Red + "[error] " + logger.Reset
		traceStr = logger.Green + "%s\n" + logger.Reset + logger.Yellow + "[%.3fms] " + logger.BlueBold + "[rows:%v]" + logger.Reset + " %s\n"
		traceWarnStr = logger.Green + "%s " + logger.Yellow + "%s\n" + logger.Reset + logger.RedBold + "[%.3fms] " + logger.Yellow + "[rows:%v]" + logger.Magenta + " %s\n" + logger.Reset
		traceErrStr = logger.RedBold + "%s " + logger.MagentaBold + "%s\n" + logger.Reset + logger.Yellow + "[%.3fms] " + logger.BlueBold + "[rows:%v]" + logger.Reset + " %s\n"
	}

	return &_logger{
		Writer:       writer,
		config:       config,
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
	}
}

type _logger struct {
	config
	logger.Writer
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
}

// LogMode log mode
func (c *_logger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *c
	newLogger.LogLevel = level
	return &newLogger
}

// Info print info
func (c *_logger) Info(ctx context.Context, message string, data ...interface{}) {
	if c.LogLevel >= logger.Info {
		c.Printf(c.infoStr+message, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Warn print warn messages
func (c *_logger) Warn(ctx context.Context, message string, data ...interface{}) {
	if c.LogLevel >= logger.Warn {
		c.Printf(c.warnStr+message, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Error print error messages
func (c *_logger) Error(ctx context.Context, message string, data ...interface{}) {
	if c.LogLevel >= logger.Error {
		c.Printf(c.errStr+message, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Trace print sql message
func (c *_logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if c.LogLevel > 0 {
		elapsed := time.Since(begin)
		switch {
		case err != nil && c.LogLevel >= logger.Error:
			sql, rows := fc()
			if rows == -1 {
				c.Printf(c.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				c.Printf(c.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		case elapsed > c.SlowThreshold && c.SlowThreshold != 0 && c.LogLevel >= logger.Warn:
			sql, rows := fc()
			slowLog := fmt.Sprintf("SLOW SQL >= %v", c.SlowThreshold)
			if rows == -1 {
				c.Printf(c.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				c.Printf(c.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		case c.LogLevel >= logger.Info:
			sql, rows := fc()
			if rows == -1 {
				c.Printf(c.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				c.Printf(c.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		}
	}
}

func (c *_logger) Printf(message string, data ...interface{}) {
	if global.GVA_CONFIG.Mysql.LogZap {
		global.GVA_LOG.Info(fmt.Sprintf(message, data...))
	} else {
		c.Writer.Printf(message, data...)
	}
}

type traceRecorder struct {
	logger.Interface
	BeginAt      time.Time
	SQL          string
	RowsAffected int64
	Err          error
}

func (t traceRecorder) New() *traceRecorder {
	return &traceRecorder{Interface: t.Interface, BeginAt: time.Now()}
}

func (t *traceRecorder) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	t.BeginAt = begin
	t.SQL, t.RowsAffected = fc()
	t.Err = err
}
