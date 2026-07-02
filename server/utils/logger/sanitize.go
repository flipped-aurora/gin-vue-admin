package logger

import (
	"net/http"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

const (
	maskValue     = "***"
	truncatedMark = "[超出记录长度]"
	MultipartMark = "[文件]"
	MaxBodyBytes  = 1024 // 配置未设置时的兜底阈值
)

// AccessLogMaxBytes 返回访问日志/操作记录对请求体、响应体的最大记录字节数。
// 优先读 zap.access-log-max-bytes 配置；未配置或 <=0 时回退到 MaxBodyBytes(1024)。
func AccessLogMaxBytes() int {
	if global.GVA_CONFIG.Zap.AccessLogMaxBytes > 0 {
		return global.GVA_CONFIG.Zap.AccessLogMaxBytes
	}
	return MaxBodyBytes
}

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
