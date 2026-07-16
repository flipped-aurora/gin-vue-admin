package logger

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
)

// W3C Trace Context(https://www.w3.org/TR/trace-context/)纯逻辑实现,零外部依赖。
// trace-id 为 16 字节 32 位小写 hex,span-id 为 8 字节 16 位小写 hex,全零均非法。
// 生成的 id 与 OTel/Datadog/SkyWalking 等所有 tracing 后端原生格式一致,
// 后续接入任何观测平台无需改造埋点。

const (
	traceIDHexLen = 32
	spanIDHexLen  = 16
)

// NewTraceID 生成 W3C 格式 trace-id(32 位小写 hex)
func NewTraceID() string { return randomHex(16) }

// NewSpanID 生成 W3C 格式 span-id(16 位小写 hex)
func NewSpanID() string { return randomHex(8) }

func randomHex(n int) string {
	b := make([]byte, n)
	_, _ = rand.Read(b) // crypto/rand 实际不会失败(Go 1.24 起有保证)
	return hex.EncodeToString(b)
}

// ParseTraceparent 解析并严格校验 traceparent 头(格式: version-traceid-spanid-flags)。
// 任一段不合法(长度、非小写 hex、trace/span 全零、version=ff)即整体拒绝,
// 避免畸形上游头污染链路 id。version 00 要求恰好 4 段;更高版本按 W3C 向前兼容
// 规则允许尾部有附加段,仅取前 4 段。
func ParseTraceparent(s string) (traceID, parentSpanID string, ok bool) {
	parts := strings.Split(s, "-")
	if len(parts) < 4 {
		return "", "", false
	}
	version, tid, sid, flags := parts[0], parts[1], parts[2], parts[3]
	if !isLowerHex(version, 2) || version == "ff" {
		return "", "", false
	}
	if version == "00" && len(parts) != 4 {
		return "", "", false
	}
	if !isLowerHex(tid, traceIDHexLen) || isAllZero(tid) {
		return "", "", false
	}
	if !isLowerHex(sid, spanIDHexLen) || isAllZero(sid) {
		return "", "", false
	}
	if !isLowerHex(flags, 2) {
		return "", "", false
	}
	return tid, sid, true
}

// BuildTraceparent 组装 version 00、sampled 置位的 traceparent 头
func BuildTraceparent(traceID, spanID string) string {
	return "00-" + traceID + "-" + spanID + "-01"
}

// IsValidTraceID 判断是否为合法 W3C trace-id(32 位小写 hex 且非全零)。
// 用于决定透传的 X-Trace-Id 能否回写成 traceparent。
func IsValidTraceID(s string) bool { return isLowerHex(s, traceIDHexLen) && !isAllZero(s) }

// HeaderIDMaxLen 与 sys_errors/sys_operation_records 的 request_id/trace_id
// 列宽 varchar(64) 对齐:超过列宽的值在 MySQL 严格模式下会让整行 INSERT 失败,
// 而错误日志入库路径吞掉该错误——恰好在出错的请求上静默丢失错误记录。
const HeaderIDMaxLen = 64

// SaneHeaderID 校验宽松透传的上游标识头(X-Request-Id/X-Trace-Id):非空、
// 不超过落库列宽、仅可见 ASCII。不满足按"无上游"处理,由调用方本地重新生成,
// 避免超长/脏值污染日志与落库。主服务 request-meta 中间件与 MCP 进程共用。
func SaneHeaderID(s string) bool {
	if s == "" || len(s) > HeaderIDMaxLen {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] <= 0x20 || s[i] > 0x7e {
			return false
		}
	}
	return true
}

func isLowerHex(s string, n int) bool {
	if len(s) != n {
		return false
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if (c < '0' || c > '9') && (c < 'a' || c > 'f') {
			return false
		}
	}
	return true
}

func isAllZero(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != '0' {
			return false
		}
	}
	return true
}
