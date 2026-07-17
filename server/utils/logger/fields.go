package logger

import "context"

type ctxKey struct{}

// 日志字段名常量:builder 输出与 core/internal/zap_core 提取共用,避免字符串隐式耦合。
// 命名与 OTel 语义约定一一对应(snake_case 形式):
// trace_id/span_id ↔ W3C Trace Context,http_method ↔ http.request.method,
// http_path ↔ url.path,client_ip ↔ client.address
const (
	FieldMod          = "mod"
	FieldRequestID    = "request_id"
	FieldTraceID      = "trace_id"
	FieldSpanID       = "span_id"
	FieldParentSpanID = "parent_span_id"
	FieldDeviceID     = "device_id"
	FieldClientIP     = "client_ip"
	FieldHTTPMethod   = "http_method"
	FieldHTTPPath     = "http_path"
	FieldErrorMsg     = "error_msg"
	FieldErrorStack   = "error_stack"
	FieldDetail       = "detail"
)

// Fields 请求级身份字段，由 request-meta 中间件注入 ctx
type Fields struct {
	RequestID    string
	TraceID      string
	SpanID       string // 本请求 server span(W3C 16hex),由 request-meta 生成
	ParentSpanID string // 上游 traceparent 的 span-id 段,无上游则空
	DeviceID     string
	ClientIP     string
	HTTPMethod   string
	HTTPPath     string
}

func WithFields(ctx context.Context, f *Fields) context.Context {
	return context.WithValue(ctx, ctxKey{}, f)
}

func FromCtx(ctx context.Context) *Fields {
	if ctx == nil {
		return nil
	}
	f, _ := ctx.Value(ctxKey{}).(*Fields)
	return f
}

func (f *Fields) GetRequestID() string {
	if f == nil {
		return ""
	}
	return f.RequestID
}

func (f *Fields) GetTraceID() string {
	if f == nil {
		return ""
	}
	return f.TraceID
}

func (f *Fields) GetSpanID() string {
	if f == nil {
		return ""
	}
	return f.SpanID
}

func (f *Fields) GetDeviceID() string {
	if f == nil {
		return ""
	}
	return f.DeviceID
}
