package mcpTool

import (
	"context"
	"net"
	"net/http"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
	"github.com/google/uuid"
)

type mcpContextKey string

const authTokenContextKey mcpContextKey = "mcp-auth-token"

func WithHTTPRequestContext(ctx context.Context, r *http.Request) context.Context {
	token := extractIncomingAuthToken(r.Header)
	ctx = context.WithValue(ctx, authTokenContextKey, token)
	// 提取/生成链路标识(与主服务 request-meta 中间件同语义):
	// 外部 AI 客户端 → MCP 进程 → GVA 主服务可串成同一条 trace
	return logger.WithFields(ctx, buildTraceFields(r))
}

// buildTraceFields 从入站 MCP HTTP 请求提取 W3C 链路标识,无上游则本地生成。
// 提取优先级与卫生校验与主服务 request-meta 中间件保持一致:
// 合法 traceparent → 通过 SaneHeaderID 的 X-Trace-Id → 本地生成。
func buildTraceFields(r *http.Request) *logger.Fields {
	clientIP := r.RemoteAddr
	if host, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
		clientIP = host // 与主服务 client_ip 字段同语义:纯 IP,不带端口
	}
	f := &logger.Fields{
		RequestID:  r.Header.Get("X-Request-Id"),
		ClientIP:   clientIP,
		HTTPMethod: r.Method,
		HTTPPath:   r.URL.Path,
	}
	// 脏值/超长值按无上游处理:否则会放大进 MCP 全量日志,且主服务侧会
	// 重新生成导致跨进程 request_id 关联静默断裂
	if !logger.SaneHeaderID(f.RequestID) {
		f.RequestID = uuid.NewString()
	}
	if tid, psid, ok := logger.ParseTraceparent(r.Header.Get("traceparent")); ok {
		f.TraceID, f.ParentSpanID = tid, psid
	} else if tid := r.Header.Get("X-Trace-Id"); logger.SaneHeaderID(tid) {
		f.TraceID = tid
	} else {
		f.TraceID = logger.NewTraceID()
	}
	f.SpanID = logger.NewSpanID()
	return f
}

func configuredAuthHeader() string {
	if header := strings.TrimSpace(global.GVA_CONFIG.MCP.AuthHeader); header != "" {
		return header
	}
	return "x-token"
}

func ConfiguredAuthHeader() string {
	return configuredAuthHeader()
}

func authTokenFromContext(ctx context.Context) string {
	token, _ := ctx.Value(authTokenContextKey).(string)
	return strings.TrimSpace(token)
}

func extractIncomingAuthToken(headers http.Header) string {
	candidates := []string{
		configuredAuthHeader(),
		"x-token",
		"token",
		"authorization",
	}

	seen := make(map[string]struct{}, len(candidates))
	for _, name := range candidates {
		key := strings.ToLower(strings.TrimSpace(name))
		if key == "" {
			continue
		}
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}

		value := strings.TrimSpace(headers.Get(name))
		if value == "" {
			continue
		}
		if key == "authorization" {
			return strings.TrimSpace(strings.TrimPrefix(value, "Bearer "))
		}
		return value
	}

	return ""
}
