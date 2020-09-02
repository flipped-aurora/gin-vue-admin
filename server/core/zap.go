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

const (
	zapLogDir      = "log"
	zapLogSoftLink = "latest_log"
	zapModule      = "gin-vue-admin"
)

func init() {
	if global.GVA_CONFIG.Zap.File == "" {
		global.GVA_CONFIG.Zap.File = "DEBUG"
	}
	if ok, _ := utils.PathExists(zapLogDir); !ok {
		// directory not exist
		fmt.Println("create log directory")
		_ = os.Mkdir(zapLogDir, os.ModePerm)
	}
	var l = new(zapcore.Level)
	writeSyncer, err := getWriteSyncer()
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return
	}
	encoder := getEncoderConfig()
	if err := l.UnmarshalText([]byte(global.GVA_CONFIG.Zap.Level)); err != nil {
		fmt.Printf("Unmarshal Level Failed err:%v", err.Error())
		return
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)
	global.GVA_ZAP = zap.New(core, zap.AddCaller())
}

// getWriteSyncer zap logger中加入file-rotatelogs
func getWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := zaprotatelogs.New(
		zapLogDir + string(os.PathSeparator) + "%Y-%m-%d-%H-%M.log",
		zaprotatelogs.WithLinkName(zapLogSoftLink),
		zaprotatelogs.WithMaxAge(7*24*time.Hour),
		zaprotatelogs.WithRotationTime(24*time.Hour),
	)
	return zapcore.AddSync(fileWriter), err
}

// getEncoderConfig 获取zapcore.Encoder
func getEncoderConfig() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}



