package logger

import (
	"context"
	"net/http"
)

// InjectTraceHeaders 把当前请求链路标识注入出站 HTTP 请求头,使下游服务可以
// 延续同一条 trace(W3C Trace Context 传播,下游是任何平台都能续联)。
// traceparent 的 parent 段直接携带当前 server span(本进程日志已落盘该
// span_id,下游日志的 parent_span_id 可精确 join 回来);仅当 ctx 里没有合法
// span 时才现场生成,避免制造全链路无任何日志记录的"幽灵 span"。
// ctx 无链路字段时不注入;trace_id 非 W3C 格式(如透传的旧网关自定义 id)时
// 只注入 X-Request-Id/X-Trace-Id,不组装 traceparent(避免产生非法头)。
func InjectTraceHeaders(ctx context.Context, req *http.Request) {
	if req == nil {
		return
	}
	f := FromCtx(ctx)
	if f == nil {
		return
	}
	if f.RequestID != "" {
		req.Header.Set("X-Request-Id", f.RequestID)
	}
	if f.TraceID == "" {
		return
	}
	req.Header.Set("X-Trace-Id", f.TraceID)
	if IsValidTraceID(f.TraceID) {
		spanID := f.SpanID
		if !isLowerHex(spanID, spanIDHexLen) || isAllZero(spanID) {
			spanID = NewSpanID()
		}
		req.Header.Set("traceparent", BuildTraceparent(f.TraceID, spanID))
	}
}
