package qmlog

// 日志初始化包  调用qmlog.QMLog.Info 记录日志 24小时切割 日志保存7天 可自行设置
import (
	"fmt"
	"gin-vue-admin/tools"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var QMLog = logrus.New()

//禁止logrus的输出
func InitLog() *logrus.Logger {
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	QMLog.Out = src
	QMLog.SetLevel(logrus.DebugLevel)
	if ok, _ := tools.PathExists("./log"); !ok {
		// Directory not exist
		fmt.Println("Create log.")
		os.Mkdir("log", os.ModePerm)
	}
	apiLogPath := "./log/api.log"
	logWriter, err := rotatelogs.New(
		apiLogPath+".%Y-%m-%d-%H-%M.log",
		rotatelogs.WithLinkName(apiLogPath),       // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{})
	QMLog.AddHook(lfHook)
	return QMLog
}
