package logger

import (
	"net/http"
	"strings"
)

const (
	maskValue     = "***"
	truncatedMark = "[超出记录长度]"
	MultipartMark = "[文件]"
	MaxBodyBytes  = 1024
)

// 需脱敏的请求头（小写比较）
var sensitiveHeaders = map[string]struct{}{
	"authorization": {}, "cookie": {}, "set-cookie": {}, "x-token": {},
}

func SanitizeHeaders(h http.Header) map[string]string {
	out := make(map[string]string, len(h))
	for k, v := range h {
		if _, ok := sensitiveHeaders[strings.ToLower(k)]; ok {
			out[k] = maskValue
			continue
		}
		out[k] = strings.Join(v, ",")
	}
	return out
}

func Truncate(s string, max int) string {
	if len(s) > max {
		return truncatedMark
	}
	return s
}
