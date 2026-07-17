package logger

import (
	"net/http"
	"strings"
	"testing"
)

func TestSanitizeHeaders(t *testing.T) {
	h := http.Header{}
	h.Set("Authorization", "Bearer secret")
	h.Set("X-Token", "tok")
	h.Set("Accept", "application/json")
	got := SanitizeHeaders(h)
	if got["Authorization"] != maskValue || got["X-Token"] != maskValue {
		t.Fatalf("sensitive headers not masked: %+v", got)
	}
	if got["Accept"] != "application/json" {
		t.Fatalf("normal header changed: %+v", got)
	}
}

func TestTruncate(t *testing.T) {
	if Truncate("short", 100) != "short" {
		t.Fatal("short should be unchanged")
	}
	long := strings.Repeat("a", 2000)
	if Truncate(long, 1024) != truncatedMark {
		t.Fatal("long should be replaced with mark")
	}
}
