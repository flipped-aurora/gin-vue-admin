package internal

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/natefinch/lumberjack"
	"github.com/songzhibin97/gkit/timeout"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"runtime"
	"strings"
	"time"
)

var Zap = new(_zap)

type _zap struct{}

// GetZapCore 跟所有日志级别获取所有的 zapcore.Core
// Author SliverHorn
func (z _zap) GetZapCore() []zapcore.Core {
	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	}) // 调试级别

	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	}) // 日志级别

	warnPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.WarnLevel
	}) // 警告级别

	errorPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.ErrorLevel
	}) // 错误级别

	dPanicPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DPanicLevel
	}) // 特别重要的错误级别

	panicPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.PanicLevel
	}) // panic级别

	fatalPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.FatalLevel
	}) // panic级别

	now := time.Now().Local().Format(timeout.DateFormat)
	cores := []zapcore.Core{
		z.GetEncoderCore(fmt.Sprintf("./%s/debug/%s.log", global.GVA_CONFIG.Zap.Director, now), debugPriority),
		z.GetEncoderCore(fmt.Sprintf("./%s/info/%s.log", global.GVA_CONFIG.Zap.Director, now), infoPriority),
		z.GetEncoderCore(fmt.Sprintf("./%s/warn/%s.log", global.GVA_CONFIG.Zap.Director, now), warnPriority),
		z.GetEncoderCore(fmt.Sprintf("./%s/error/%s.log", global.GVA_CONFIG.Zap.Director, now), errorPriority),
		z.GetEncoderCore(fmt.Sprintf("./%s/dpanic/%s.log", global.GVA_CONFIG.Zap.Director, now), dPanicPriority),
		z.GetEncoderCore(fmt.Sprintf("./%s/panic/%s.log", global.GVA_CONFIG.Zap.Director, now), panicPriority),
		z.GetEncoderCore(fmt.Sprintf("./%s/fatal/%s.log", global.GVA_CONFIG.Zap.Director, now), fatalPriority),
	}
	return cores
}

// GetWriteSyncer zap logger中加入lumberjack lumberjack 切割
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *_zap) GetWriteSyncer(file string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file, //日志文件的位置
		MaxSize:    10,   //在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 200,  //保留旧文件的最大个数
		MaxAge:     30,   //保留旧文件的最大天数
		Compress:   true, //是否压缩/归档旧文件
	}

	if global.GVA_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}
	return zapcore.AddSync(lumberJackLogger)
}

// GetEncoderConfig 获取zapcore.EncoderConfig
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *_zap) GetEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.GVA_CONFIG.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     z.CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	if runtime.GOOS == "windows" {
		if strings.Contains(global.GVA_CONFIG.Zap.EncodeLevel, "Color") {
			global.GVA_CONFIG.Zap.EncodeLevel = strings.Replace(global.GVA_CONFIG.Zap.EncodeLevel, "Color", "", -1)
		}
		if strings.Contains(global.GVA_CONFIG.Zap.EncodeLevel, "color") {
			global.GVA_CONFIG.Zap.EncodeLevel = strings.Replace(global.GVA_CONFIG.Zap.EncodeLevel, "color", "", -1)
		}
	}
	switch {
	case global.GVA_CONFIG.Zap.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case global.GVA_CONFIG.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case global.GVA_CONFIG.Zap.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case global.GVA_CONFIG.Zap.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// GetEncoder 获取zapcore.Encoder
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *_zap) GetEncoder() zapcore.Encoder {
	if global.GVA_CONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(z.GetEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(z.GetEncoderConfig())
}

// GetEncoderCore 获取Encoder的zapcore.Core
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *_zap) GetEncoderCore(fileName string, level zapcore.LevelEnabler) (core zapcore.Core) {
	writer := z.GetWriteSyncer(fileName) // 使用file-rotatelogs进行日志分割
	return zapcore.NewCore(z.GetEncoder(), writer, level)
}

// CustomTimeEncoder 自定义日志输出时间格式
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *_zap) CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.GVA_CONFIG.Zap.Prefix + timeout.DateTimeFormat))
}
