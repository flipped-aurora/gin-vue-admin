package logger

import (
	"context"
	"net/http"
	"testing"
)

func newOutboundReq(t *testing.T) *http.Request {
	t.Helper()
	req, err := http.NewRequest(http.MethodGet, "http://downstream/api", nil)
	if err != nil {
		t.Fatal(err)
	}
	return req
}

func TestInjectTraceHeaders(t *testing.T) {
	tid, sid := NewTraceID(), NewSpanID()
	ctx := WithFields(context.Background(), &Fields{RequestID: "rid-1", TraceID: tid, SpanID: sid})
	req := newOutboundReq(t)
	InjectTraceHeaders(ctx, req)

	if req.Header.Get("X-Request-Id") != "rid-1" || req.Header.Get("X-Trace-Id") != tid {
		t.Fatalf("id headers wrong: %+v", req.Header)
	}
	gotTid, gotSid, ok := ParseTraceparent(req.Header.Get("traceparent"))
	if !ok || gotTid != tid {
		t.Fatalf("traceparent should carry same trace id: %q", req.Header.Get("traceparent"))
	}
	// parent 段必须是当前 server span(本进程日志已落盘该 span_id),
	// 下游的 parent_span_id 才能精确 join 回上游日志,不能现场生成幽灵 span
	if gotSid != sid {
		t.Fatalf("traceparent parent segment should be the current server span %q, got %q", sid, gotSid)
	}
}

func TestInjectTraceHeadersMissingSpanFallsBack(t *testing.T) {
	tid := NewTraceID()
	// ctx 里没有 span(如非 request-meta 构造的 Fields):仍需组装合法 traceparent
	ctx := WithFields(context.Background(), &Fields{RequestID: "rid-3", TraceID: tid})
	req := newOutboundReq(t)
	InjectTraceHeaders(ctx, req)
	gotTid, gotSid, ok := ParseTraceparent(req.Header.Get("traceparent"))
	if !ok || gotTid != tid || gotSid == "" {
		t.Fatalf("missing span should fall back to generated span id: %q", req.Header.Get("traceparent"))
	}
}

func TestInjectTraceHeadersNoFields(t *testing.T) {
	req := newOutboundReq(t)
	InjectTraceHeaders(context.Background(), req)
	if len(req.Header) != 0 {
		t.Fatalf("no fields in ctx should inject nothing: %+v", req.Header)
	}
}

func TestInjectTraceHeadersNonW3CTraceID(t *testing.T) {
	ctx := WithFields(context.Background(), &Fields{RequestID: "rid-2", TraceID: "legacy-gateway-id"})
	req := newOutboundReq(t)
	InjectTraceHeaders(ctx, req)
	if req.Header.Get("X-Trace-Id") != "legacy-gateway-id" {
		t.Fatal("X-Trace-Id should pass through")
	}
	if req.Header.Get("traceparent") != "" {
		t.Fatalf("non-W3C trace id must not produce traceparent: %q", req.Header.Get("traceparent"))
	}
}
