package logger

import (
	"encoding/json"
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

// 需脱敏的请求/响应体字段名（统一小写、去掉 _ - 后比较）
var sensitiveBodyKeys = map[string]struct{}{
	"password": {}, "newpassword": {}, "oldpassword": {}, "confirmpassword": {},
	"passwd": {}, "pwd": {}, "token": {}, "accesstoken": {}, "refreshtoken": {},
	"secret": {}, "clientsecret": {}, "apikey": {}, "privatekey": {}, "idcard": {},
}

func normalizeBodyKey(k string) string {
	k = strings.ToLower(k)
	k = strings.ReplaceAll(k, "_", "")
	k = strings.ReplaceAll(k, "-", "")
	return k
}

// SanitizeBody 对 JSON 请求/响应体做字段级脱敏：按字段名匹配 password / token / secret 等
// 敏感键，将其值替换为掩码。仅处理 JSON；非 JSON 或解析失败时原样返回，
// 保证脱敏逻辑本身不会影响到日志记录（例如截断后的半截 JSON 直接原样返回）。
func SanitizeBody(contentType, body string) string {
	if body == "" || !strings.Contains(strings.ToLower(contentType), "json") {
		return body
	}
	var v any
	if err := json.Unmarshal([]byte(body), &v); err != nil {
		return body
	}
	maskSensitiveBody(v)
	out, err := json.Marshal(v)
	if err != nil {
		return body
	}
	return string(out)
}

func maskSensitiveBody(v any) {
	switch t := v.(type) {
	case map[string]any:
		for k, val := range t {
			if _, ok := sensitiveBodyKeys[normalizeBodyKey(k)]; ok {
				t[k] = maskValue
				continue
			}
			maskSensitiveBody(val)
		}
	case []any:
		for _, item := range t {
			maskSensitiveBody(item)
		}
	}
}
