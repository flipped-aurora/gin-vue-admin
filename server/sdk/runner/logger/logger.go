package logger

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/jsonx"
	"runtime"
	"strconv"
	"time"
)

type Logger struct {
	Level   string
	DataMap map[string]interface{} `json:"data_map"`
}

func (l *Logger) logPrint(level string, msg string) {
	_, file, line, _ := runtime.Caller(2)
	l.DataMap["ts"] = time.Now().Unix()
	l.DataMap["msg"] = msg
	l.DataMap["stack"] = file + ":" + strconv.Itoa(line)
	l.DataMap["level"] = level

	fmt.Println("<Logger>" + jsonx.JSONString(l.DataMap) + "</Logger>")
}

func (l *Logger) logPrintf(level string, formatMsg string, args any) {
	_, file, line, _ := runtime.Caller(1)
	l.DataMap["ts"] = time.Now().Unix()
	l.DataMap["msg"] = fmt.Sprintf(formatMsg, args)
	l.DataMap["stack"] = file + ":" + strconv.Itoa(line)
	if l.Level != "" {
		l.DataMap["level"] = l.Level
	} else {
		l.DataMap["level"] = level
	}
	fmt.Println("<Logger>" + jsonx.JSONString(l.DataMap) + "</Logger>")
}

// SetLevel 可以自定义level
func (l *Logger) SetLevel(level string) *Logger {
	l.Level = level
	return l
}
func (l *Logger) Info(msg string) {
	l.logPrint("INFO", msg)
}
func (l *Logger) Infof(formatMsg string, args ...any) {
	l.logPrintf("INFO", formatMsg, args)
}

func (l *Logger) Warn(msg string) {
	l.logPrint("WARN", msg)
}
func (l *Logger) Warnf(formatMsg string, args ...any) {
	l.logPrintf("WARN", formatMsg, args)
}

func (l *Logger) Error(msg string) {
	l.logPrint("ERROR", msg)
}
func (l *Logger) Errorf(formatMsg string, args ...any) {
	l.logPrintf("ERROR", formatMsg, args)
}
func (l *Logger) Panic(msg string) {
	l.logPrint("PANIC", msg)
}
func (l *Logger) Panicf(formatMsg string, args ...any) {
	l.logPrintf("PANIC", formatMsg, args)
}
func (l *Logger) Debug(msg string) {
	l.logPrint("DEBUG", msg)
}
func (l *Logger) Debugf(formatMsg string, args ...any) {
	l.logPrintf("DEBUG", formatMsg, args)
}

func (l *Logger) Fatal(msg string) {
	l.logPrint("FATAL", msg)
}

func (l *Logger) Fatalf(formatMsg string, args ...any) {
	l.logPrintf("FATAL", formatMsg, args)
}

func (l *Logger) WithField(key string, value interface{}) *Logger {
	if _, ok := l.DataMap[key]; !ok {
		l.DataMap[key] = value
	} else {
		l.DataMap["runner."+key] = value
	}
	return l
}
func (l *Logger) WithFields(fields map[string]interface{}) *Logger {
	for k, v := range fields {
		if _, ok := l.DataMap[k]; !ok {
			l.DataMap[k] = v
		} else {
			l.DataMap["runner."+k] = v
		}
	}
	return l
}
