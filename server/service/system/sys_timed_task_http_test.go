// server/service/system/sys_timed_task_http_test.go
package system

import (
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestIsDisallowedIP(t *testing.T) {
	cases := map[string]bool{
		"127.0.0.1":   true,  // 环回
		"10.0.0.5":    true,  // 私网
		"172.16.1.1":  true,  // 私网
		"192.168.1.1": true,  // 私网
		"169.254.1.1": true,  // 链路本地
		"0.0.0.0":     true,  // 未指定
		"::1":         true,  // v6 环回
		"8.8.8.8":     false, // 公网
	}
	for ipStr, want := range cases {
		if got := isDisallowedIP(net.ParseIP(ipStr)); got != want {
			t.Errorf("isDisallowedIP(%s) = %v, want %v", ipStr, got, want)
		}
	}
}

func TestSSRFBlocksLoopback(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(200)
	}))
	defer srv.Close()

	// 默认禁内网: 请求 127.0.0.1 必须被拒
	client := newTimedTaskHTTPClient(false, 5*time.Second)
	_, err := client.Get(srv.URL)
	if err == nil || !strings.Contains(err.Error(), "SSRF") {
		t.Fatalf("期望 SSRF 拒绝, got %v", err)
	}

	// 任务级豁免: allowPrivate=true 放行
	clientAllow := newTimedTaskHTTPClient(true, 5*time.Second)
	resp, err := clientAllow.Get(srv.URL)
	if err != nil {
		t.Fatalf("allowPrivate 应放行: %v", err)
	}
	resp.Body.Close()
}
