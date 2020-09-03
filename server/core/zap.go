package core

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/utils"
	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var (
	err    error
	level zapcore.Level
	writer zapcore.WriteSyncer
)

func init() {
	if ok, _ := utils.PathExists(global.GVA_CONFIG.Zap.Director); !ok { // 判断是否有logs文件夹
		fmt.Println("create logs directory") // directory not exist
		_ = os.Mkdir(global.GVA_CONFIG.Zap.Director, os.ModePerm)
	}

	switch global.GVA_CONFIG.Zap.Level {// 初始化配置文件的Level
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	writer, err = getWriteSyncer() // 使用file-rotatelogs进行日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return
	}

	if level == zap.DebugLevel || level == zap.ErrorLevel {
		global.GVA_ZAP = zap.New(getEncoderCore(), zap.AddStacktrace(level))
		if global.GVA_CONFIG.Zap.ShowLine {
			global.GVA_ZAP.WithOptions(zap.AddCaller())
		}
		return
	}
	global.GVA_ZAP = zap.New(getEncoderCore())
	if global.GVA_CONFIG.Zap.ShowLine {
		global.GVA_ZAP.WithOptions(zap.AddCaller())
	}

}

// getWriteSyncer zap logger中加入file-rotatelogs
func getWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := zaprotatelogs.New(
		global.GVA_CONFIG.Zap.Director+string(os.PathSeparator)+"%Y-%m-%d-%H-%M.log",
		zaprotatelogs.WithLinkName(global.GVA_CONFIG.Zap.LinkName),
		zaprotatelogs.WithMaxAge(7*24*time.Hour),
		zaprotatelogs.WithRotationTime(24*time.Hour),
	)
	if global.GVA_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	return config
}

// getEncoder 获取zapcore.Encoder
func getEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore() (core zapcore.Core) {
	return zapcore.NewCore(getEncoder(), writer, level)
}

// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.GVA_CONFIG.Zap.Prefix+"2006/01/02 - 15:04:05.000"))
}
