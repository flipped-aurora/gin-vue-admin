package logger

import (
	"context"
	"testing"
)

func TestWithFieldsRoundTrip(t *testing.T) {
	ctx := WithFields(context.Background(), &Fields{RequestID: "rid-1", ClientIP: "1.2.3.4"})
	got := FromCtx(ctx)
	if got == nil || got.RequestID != "rid-1" || got.ClientIP != "1.2.3.4" {
		t.Fatalf("round-trip failed: %+v", got)
	}
}

func TestFromCtxNil(t *testing.T) {
	var nilCtx context.Context
	if FromCtx(nilCtx) != nil {
		t.Fatal("FromCtx(nil) should be nil")
	}
	if FromCtx(context.Background()) != nil {
		t.Fatal("FromCtx without fields should be nil")
	}
}

func TestGetRequestIDNilSafe(t *testing.T) {
	var f *Fields
	if f.GetRequestID() != "" {
		t.Fatal("nil Fields.GetRequestID should be empty")
	}
}
