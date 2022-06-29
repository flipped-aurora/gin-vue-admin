package internal

import (
	"bytes"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
	"runtime"
	"strings"
	"time"
)

var Zap = new(_zap)

type _zap struct{}

// GetEncoder 获取 zapcore.Encoder
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *_zap) GetEncoder() zapcore.Encoder {
	if global.GVA_CONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(z.GetEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(z.GetEncoderConfig())
}

// GetEncoderConfig 获取zapcore.EncoderConfig
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *_zap) GetEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.GVA_CONFIG.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    global.GVA_CONFIG.Zap.ZapEncodeLevel(),
		EncodeTime:     z.CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   CallerEncoder,
	}
}

// GetEncoderCore 获取Encoder的 zapcore.Core
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *_zap) GetEncoderCore(l zapcore.Level, level zap.LevelEnablerFunc) zapcore.Core {
	writer, err := FileRotatelogs.GetWriteSyncer(l.String()) // 使用file-rotatelogs进行日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return nil
	}

	return zapcore.NewCore(&EscapeSeqJSONEncoder{z.GetEncoder()}, writer, level)
}

// CustomTimeEncoder 自定义日志输出时间格式
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *_zap) CustomTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(t.Format(global.GVA_CONFIG.Zap.Prefix + "2006/01/02 - 15:04:05.000"))
}

// GetZapCores 根据配置文件的Level获取 []zapcore.Core
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *_zap) GetZapCores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)
	for level := global.GVA_CONFIG.Zap.TransportLevel(); level <= zapcore.FatalLevel; level++ {
		cores = append(cores, z.GetEncoderCore(level, z.GetLevelPriority(level)))
	}
	return cores
}

// GetLevelPriority 根据 zapcore.Level 获取 zap.LevelEnablerFunc
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *_zap) GetLevelPriority(level zapcore.Level) zap.LevelEnablerFunc {
	switch level {
	case zapcore.DebugLevel:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	case zapcore.InfoLevel:
		return func(level zapcore.Level) bool { // 日志级别
			return level == zap.InfoLevel
		}
	case zapcore.WarnLevel:
		return func(level zapcore.Level) bool { // 警告级别
			return level == zap.WarnLevel
		}
	case zapcore.ErrorLevel:
		return func(level zapcore.Level) bool { // 错误级别
			return level == zap.ErrorLevel
		}
	case zapcore.DPanicLevel:
		return func(level zapcore.Level) bool { // dpanic级别
			return level == zap.DPanicLevel
		}
	case zapcore.PanicLevel:
		return func(level zapcore.Level) bool { // panic级别
			return level == zap.PanicLevel
		}
	case zapcore.FatalLevel:
		return func(level zapcore.Level) bool { // 终止级别
			return level == zap.FatalLevel
		}
	default:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	}
}

// FuncName 返回调用本函数的函数名称
// pc runtime.Caller 返回的第一个值
func FuncName(pc uintptr) string {
	funcName := runtime.FuncForPC(pc).Name()
	sFuncName := strings.Split(funcName, ".")
	return sFuncName[len(sFuncName)-1]
}

// CallerEncoder serializes a caller in package/file:funcname:line format
func CallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	shortCaller := caller.TrimmedPath()
	shortCallerSplited := strings.Split(shortCaller, ":")
	funcName := FuncName(caller.PC)
	result := shortCallerSplited[0] + ":" + funcName + ":" + shortCallerSplited[1]
	enc.AppendString(result)
}

type EscapeSeqJSONEncoder struct {
	zapcore.Encoder
}

// EncodeEntry 将方法zap.error中的errorVerbose的堆栈换行符修改
func (enc *EscapeSeqJSONEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	b, err := enc.Encoder.EncodeEntry(entry, fields)
	if err != nil {
		return nil, err
	}
	newb := buffer.NewPool().Get()

	b1 := bytes.Replace(b.Bytes(), []byte("\\n"), []byte("\n"), -1)
	b2 := bytes.Replace(b1, []byte("\\t"), []byte("\t"), -1)
	_, err = newb.Write(b2)
	if err != nil {
		return nil, err
	}
	return newb, nil
}
