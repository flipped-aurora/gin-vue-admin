package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	mcpTool "github.com/flipped-aurora/gin-vue-admin/server/mcp"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
)

// unsafeListenHosts is the set of MCP bind hosts that the standalone MCP
// binary refuses to start on. The MCP endpoint exposes code-generation,
// database execution and route mutation tools, so binding to any public
// interface effectively grants RCE to anyone who can reach the port.
//
// Operators that deliberately want MCP reachable from outside localhost
// should pick a specific private IP (e.g. the Docker bridge or a VPN
// interface) rather than blanket 0.0.0.0 / ::.
var unsafeListenHosts = map[string]struct{}{
	"0.0.0.0": {},
	"::":      {},
	"*":       {},
}

func main() {
	configPath, err := loadStandaloneConfig()
	if err != nil {
		panic(err)
	}

	if err := initializeStandaloneLogger(); err != nil {
		panic(err)
	}

	host := strings.TrimSpace(global.GVA_CONFIG.MCP.ListenHost)
	if _, bad := unsafeListenHosts[host]; bad {
		panic(fmt.Errorf(
			"mcp.listen_host=%q is refused: binding MCP to any public interface exposes "+
				"code-generation and DB-execution tools. Set mcp.listen_host to 127.0.0.1 "+
				"(loopback) or a specific private interface, and front it with a reverse proxy "+
				"if remote access is required.",
			host,
		))
	}

	addr := net.JoinHostPort(host, strconv.Itoa(global.GVA_CONFIG.MCP.Addr))
	server := mcpTool.NewStreamableHTTPServer()

	global.GVA_LOG.Info("mcp独立服务启动",
		zap.String("config", configPath),
		zap.String("addr", addr),
		zap.String("listen_host", host),
		zap.String("path", global.GVA_CONFIG.MCP.Path),
		zap.String("upstream", global.GVA_CONFIG.MCP.UpstreamBaseURL),
	)

	if err := server.Start(addr); err != nil {
		global.GVA_LOG.Fatal("mcp独立服务启动失败", zap.Error(err))
	}
}
