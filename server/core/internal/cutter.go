package internal

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
)

type Cutter struct {
	level    string        // 日志级别(debug, info, warn, error, dpanic, panic, fatal)
	format   string        // 时间格式(2006-01-02)
	Director string        // 日志文件夹
	file     *os.File      // 文件句柄
	mutex    *sync.RWMutex // 读写锁
}

type CutterOption func(*Cutter)

// WithCutterFormat 设置时间格式
func WithCutterFormat(format string) CutterOption {
	return func(c *Cutter) {
		c.format = format
	}
}

func NewCutter(director string, level string, options ...CutterOption) *Cutter {
	rotate := &Cutter{
		level:    level,
		Director: director,
		mutex:    new(sync.RWMutex),
	}
	for i := 0; i < len(options); i++ {
		options[i](rotate)
	}
	return rotate
}

// Write satisfies the io.Writer interface. It writes to the
// appropriate file handle that is currently being used.
// If we have reached rotation time, the target file gets
// automatically rotated, and also purged if necessary.
func (c *Cutter) Write(bytes []byte) (n int, err error) {
	c.mutex.Lock()
	defer func() {
		if c.file != nil {
			_ = c.file.Close()
			c.file = nil
		}
		c.mutex.Unlock()
	}()
	var business string
	if strings.Contains(string(bytes), "business") {
		var compile *regexp.Regexp
		compile, err = regexp.Compile(`{"business": "([^,]+)"}`)
		if err != nil {
			return 0, err
		}
		if compile.Match(bytes) {
			finds := compile.FindSubmatch(bytes)
			business = string(finds[len(finds)-1])
			bytes = compile.ReplaceAll(bytes, []byte(""))
		}
		compile, err = regexp.Compile(`"business": "([^,]+)"`)
		if err != nil {
			return 0, err
		}
		if compile.Match(bytes) {
			finds := compile.FindSubmatch(bytes)
			business = string(finds[len(finds)-1])
			bytes = compile.ReplaceAll(bytes, []byte(""))
		}
	}
	format := time.Now().Format(c.format)
	formats := make([]string, 0, 4)
	formats = append(formats, c.Director)
	if format != "" {
		formats = append(formats, format)
	}
	if business != "" {
		formats = append(formats, business)
	}
	formats = append(formats, c.level+".log")
	filename := filepath.Join(formats...)
	dirname := filepath.Dir(filename)
	err = os.MkdirAll(dirname, 0755)
	if err != nil {
		return 0, err
	}
	c.file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	return c.file.Write(bytes)
}
