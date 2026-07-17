package logger

import (
	"strings"
	"testing"
)

func TestNewTraceIDFormat(t *testing.T) {
	seen := map[string]bool{}
	for i := 0; i < 100; i++ {
		id := NewTraceID()
		if !IsValidTraceID(id) {
			t.Fatalf("invalid trace id: %q", id)
		}
		if seen[id] {
			t.Fatalf("duplicated trace id: %q", id)
		}
		seen[id] = true
	}
}

func TestNewSpanIDFormat(t *testing.T) {
	id := NewSpanID()
	if !isLowerHex(id, spanIDHexLen) || isAllZero(id) {
		t.Fatalf("invalid span id: %q", id)
	}
}

func TestParseTraceparentValid(t *testing.T) {
	tid, psid, ok := ParseTraceparent("00-4bf92f3577b34da6a3ce929d0e0e4736-00f067aa0ba902b7-01")
	if !ok || tid != "4bf92f3577b34da6a3ce929d0e0e4736" || psid != "00f067aa0ba902b7" {
		t.Fatalf("valid traceparent rejected: %q %q %v", tid, psid, ok)
	}
}

func TestParseTraceparentFutureVersionExtraParts(t *testing.T) {
	// W3C 向前兼容:更高 version 允许尾部附加段,取前 4 段
	_, _, ok := ParseTraceparent("01-4bf92f3577b34da6a3ce929d0e0e4736-00f067aa0ba902b7-01-extra")
	if !ok {
		t.Fatal("future version with extra parts should be accepted")
	}
}

func TestParseTraceparentInvalid(t *testing.T) {
	cases := map[string]string{
		"empty":            "",
		"garbage":          "not-a-traceparent",
		"tooFewParts":      "00-4bf92f3577b34da6a3ce929d0e0e4736-01",
		"version00Extra":   "00-4bf92f3577b34da6a3ce929d0e0e4736-00f067aa0ba902b7-01-extra",
		"versionFF":        "ff-4bf92f3577b34da6a3ce929d0e0e4736-00f067aa0ba902b7-01",
		"zeroTraceID":      "00-00000000000000000000000000000000-00f067aa0ba902b7-01",
		"zeroSpanID":       "00-4bf92f3577b34da6a3ce929d0e0e4736-0000000000000000-01",
		"shortTraceID":     "00-4bf92f3577b34da6-00f067aa0ba902b7-01",
		"upperHexTraceID":  "00-4BF92F3577B34DA6A3CE929D0E0E4736-00f067aa0ba902b7-01",
		"nonHexFlags":      "00-4bf92f3577b34da6a3ce929d0e0e4736-00f067aa0ba902b7-zz",
		"nonHexVersion":    "zz-4bf92f3577b34da6a3ce929d0e0e4736-00f067aa0ba902b7-01",
		"whitespacePrefix": " 00-4bf92f3577b34da6a3ce929d0e0e4736-00f067aa0ba902b7-01",
	}
	for name, tp := range cases {
		if _, _, ok := ParseTraceparent(tp); ok {
			t.Errorf("%s: malformed traceparent accepted: %q", name, tp)
		}
	}
}

func TestBuildTraceparentRoundTrip(t *testing.T) {
	tid, sid := NewTraceID(), NewSpanID()
	tp := BuildTraceparent(tid, sid)
	if !strings.HasPrefix(tp, "00-") || !strings.HasSuffix(tp, "-01") {
		t.Fatalf("bad traceparent shape: %q", tp)
	}
	gotTid, gotSid, ok := ParseTraceparent(tp)
	if !ok || gotTid != tid || gotSid != sid {
		t.Fatalf("round-trip failed: %q -> %q %q %v", tp, gotTid, gotSid, ok)
	}
}

func TestIsValidTraceID(t *testing.T) {
	if IsValidTraceID("short") || IsValidTraceID(strings.Repeat("0", 32)) || IsValidTraceID(strings.Repeat("G", 32)) {
		t.Fatal("invalid trace id accepted")
	}
	if !IsValidTraceID(strings.Repeat("a", 32)) {
		t.Fatal("valid trace id rejected")
	}
}
