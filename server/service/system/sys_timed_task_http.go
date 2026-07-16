// server/service/system/sys_timed_task_http.go
package system

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"syscall"
	"time"
)

// errPrivateAddr SSRF 防护拒绝(错误信息含"SSRF"关键字, 供日志/测试识别)
var errPrivateAddr = errors.New("目标解析为内网/环回/链路本地地址, 已被 SSRF 防护拒绝(可在任务上开启\"允许内网\"豁免)")

// isDisallowedIP 内网/环回/链路本地/未指定地址判定
func isDisallowedIP(ip net.IP) bool {
	return ip.IsLoopback() || ip.IsPrivate() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() || ip.IsUnspecified()
}

// ssrfControl 在拨号阶段(DNS 解析后、连接建立前)校验目标 IP:
// 每次连接都过检, 天然覆盖重定向与 DNS rebinding(TOCTOU 安全)。
func ssrfControl(allowPrivate bool) func(network, address string, c syscall.RawConn) error {
	return func(_ string, address string, _ syscall.RawConn) error {
		if allowPrivate {
			return nil
		}
		host, _, err := net.SplitHostPort(address)
		if err != nil {
			return fmt.Errorf("解析拨号地址失败: %w", err)
		}
		ip := net.ParseIP(host)
		if ip == nil {
			return fmt.Errorf("非法拨号 IP: %s", host)
		}
		if isDisallowedIP(ip) {
			return fmt.Errorf("%w: %s", errPrivateAddr, ip)
		}
		return nil
	}
}

// newTimedTaskHTTPClient 定时任务专用 HTTP 客户端:
// 整体超时 + 禁用环境代理(防经代理绕过 IP 校验) + 拨号层 SSRF 防护
func newTimedTaskHTTPClient(allowPrivate bool, timeout time.Duration) *http.Client {
	dialer := &net.Dialer{
		Timeout: 10 * time.Second,
		Control: ssrfControl(allowPrivate),
	}
	transport := &http.Transport{
		Proxy:       nil, // 显式禁用代理
		DialContext: dialer.DialContext,
	}
	return &http.Client{Timeout: timeout, Transport: transport}
}
