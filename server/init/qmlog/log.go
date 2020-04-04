package qmlog

// Register logger
import (
	"errors"
	"fmt"
	"gin-vue-admin/init"
	"gin-vue-admin/utils"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	oplogging "github.com/op/go-logging"
	"io"
	"os"
	"strings"
	"time"
)

const (
	logDir      = "log"
	logSoftLink = "api.log"
	module      = "gin-vue-admin"
)

var (
	configNotFound = errors.New("logger prefix not found")

	defaultFormatter = `%{time:2006/01/02 - 15:04:05.000} %{longfile} %{color:bold}▶ [%{level:.6s}] %{message}%{color:reset}`
)

type Logger struct {
	logger *oplogging.Logger
}

func NewLogger() error {
	c := init.GinVueAdminconfig.Log
	if c.Prefix == "" {
		return configNotFound
	}
	logger := oplogging.MustGetLogger(module)
	var backends []oplogging.Backend
	backends = registerStdout(c, backends)
	backends = registerFile(c, backends)

	oplogging.SetBackend(backends...)
	init.SetLogger(logger)
	return nil
}

func registerStdout(c init.Log, backends []oplogging.Backend) []oplogging.Backend {
	for _, v := range c.Stdout {
		level, err := oplogging.LogLevel(v)
		if err != nil {
			fmt.Println(err)
			continue
		}
		backends = append(backends, createBackend(os.Stdout, c, level))
	}

	return backends
}

func registerFile(c init.Log, backends []oplogging.Backend) []oplogging.Backend {
	if len(c.File) > 0 {
		if ok, _ := utils.PathExists(logDir); !ok {
			// directory not exist
			fmt.Println("create log directory")
			_ = os.Mkdir(logDir, os.ModePerm)
		}
		apiLogPath := logDir + string(os.PathSeparator) + logSoftLink
		fileWriter, err := rotatelogs.New(
			apiLogPath+".%Y-%m-%d-%H-%M.log",
			// generate soft link, point to latest log file
			rotatelogs.WithLinkName(apiLogPath),
			// maximum time to save log files
			rotatelogs.WithMaxAge(7*24*time.Hour),
			// time period of log file switching
			rotatelogs.WithRotationTime(24*time.Hour),
		)
		if err != nil {
			fmt.Println(err)
			return backends
		}
		for _, v := range c.File {
			level, err := oplogging.LogLevel(v)
			if err != nil {
				fmt.Println(err)
				continue
			}
			backends = append(backends, createBackend(fileWriter, c, level))
		}
	}

	return backends
}

func createBackend(w io.Writer, c init.Log, level oplogging.Level) oplogging.Backend {
	backend := oplogging.NewLogBackend(w, c.Prefix, 0)
	stdoutWriter := false
	if w == os.Stdout {
		stdoutWriter = true
	}
	format := getLogFormatter(c, stdoutWriter)
	backendLeveled := oplogging.AddModuleLevel(oplogging.NewBackendFormatter(backend, format))
	backendLeveled.SetLevel(level, module)
	return backendLeveled
}

func getLogFormatter(c init.Log, stdoutWriter bool) oplogging.Formatter {
	pattern := defaultFormatter
	if !stdoutWriter {
		// Color is only required for console output
		// Other writers don't need %{color} tag
		pattern = strings.Replace(pattern, "%{color:bold}", "", -1)
		pattern = strings.Replace(pattern, "%{color:reset}", "", -1)
	}
	if !c.LogFile {
		// Remove %{logfile} tag
		pattern = strings.Replace(pattern, "%{longfile}", "", -1)
	}
	return oplogging.MustStringFormatter(pattern)
}

func (l Logger) Debug(v ...interface{}) {
	l.logger.Debug(v)
}

func (l Logger) Info(v ...interface{}) {
	l.logger.Info(v)
}

func (l Logger) Warning(v ...interface{}) {
	l.logger.Warning(v)
}

func (l Logger) Error(v ...interface{}) {
	l.logger.Error(v)
}

func (l Logger) Critical(v ...interface{}) {
	l.logger.Critical(v)
}

func (l Logger) Fatal(v ...interface{}) {
	l.logger.Fatal(v)
}
