package core

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/core/internal"
	"os"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Zap 获取 zap.Logger
// Author [SliverHorn](https://github.com/SliverHorn)
func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.GVA_CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.GVA_CONFIG.Zap.Director)
		_ = os.Mkdir(global.GVA_CONFIG.Zap.Director, os.ModePerm)
	}

	cores := make([]zapcore.Core, 0, 7)
	debugLevel := getEncoderCore(zap.DebugLevel)
	infoLevel := getEncoderCore(zap.InfoLevel)
	warnLevel := getEncoderCore(zap.WarnLevel)
	errorLevel := getEncoderCore(zap.ErrorLevel)
	dPanicLevel := getEncoderCore(zap.DPanicLevel)
	panicLevel := getEncoderCore(zap.PanicLevel)
	fatalLevel := getEncoderCore(zap.FatalLevel)
	switch global.GVA_CONFIG.Zap.Level {
	case "debug", "DEBUG":
		cores = append(cores, debugLevel, infoLevel, warnLevel, errorLevel, dPanicLevel, panicLevel, fatalLevel)
	case "info", "INFO":
		cores = append(cores, infoLevel, warnLevel, errorLevel, dPanicLevel, panicLevel, fatalLevel)
	case "warn", "WARN":
		cores = append(cores, warnLevel, errorLevel, dPanicLevel, panicLevel, fatalLevel)
	case "error", "ERROR":
		cores = append(cores, errorLevel, dPanicLevel, panicLevel, fatalLevel)
	case "dpanic", "DPANIC":
		cores = append(cores, dPanicLevel, panicLevel, fatalLevel)
	case "panic", "PANIC":
		cores = append(cores, panicLevel, fatalLevel)
	case "fatal", "FATAL":
		cores = append(cores, panicLevel, fatalLevel)
	default:
		cores = append(cores, debugLevel, infoLevel, warnLevel, errorLevel, dPanicLevel, panicLevel, fatalLevel)
	}
	logger = zap.New(zapcore.NewTee(cores...), zap.AddCaller())

	if global.GVA_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

// getEncoderConfig 获取zapcore.EncoderConfig
// Author [SliverHorn](https://github.com/SliverHorn)
func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.GVA_CONFIG.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
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

// getEncoder 获取zapcore.Encoder
// Author [SliverHorn](https://github.com/SliverHorn)
func getEncoder() zapcore.Encoder {
	if global.GVA_CONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderCore 获取Encoder的zapcore.Core
// Author [SliverHorn](https://github.com/SliverHorn)
func getEncoderCore(level zapcore.Level) (core zapcore.Core) {
	writer, err := internal.FileRotatelogs.GetWriteSyncer(level.String()) // 使用file-rotatelogs进行日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return
	}
	return zapcore.NewCore(getEncoder(), writer, level)
}

// CustomTimeEncoder 自定义日志输出时间格式
// Author [SliverHorn](https://github.com/SliverHorn)
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.GVA_CONFIG.Zap.Prefix + "2006/01/02 - 15:04:05.000"))
}
