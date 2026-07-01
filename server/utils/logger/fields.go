package logger

import "context"

type ctxKey struct{}

// Fields 请求级身份字段，由 request-meta 中间件注入 ctx
type Fields struct {
	RequestID  string
	TraceID    string
	DeviceID   string
	ClientIP   string
	HTTPMethod string
	HTTPPath   string
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
