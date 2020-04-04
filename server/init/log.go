package init

// Custom Logger
type Logger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warning(v ...interface{})
	Error(v ...interface{})
	Critical(v ...interface{})
	Fatal(v ...interface{})
}

var L Logger

func SetLogger(logger Logger) {
	L = logger
}
