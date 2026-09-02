package mcpTool

import (
	"net/http"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	mcpServer "github.com/mark3labs/mcp-go/server"
)

func NewMCPServer() *mcpServer.MCPServer {
	config := global.GVA_CONFIG.MCP

	s := mcpServer.NewMCPServer(
		config.Name,
		config.Version,
	)

	global.GVA_MCP_SERVER = s
	RegisterAllTools(s)

	return s
}

func NewStreamableHTTPServer() *mcpServer.StreamableHTTPServer {
	config := global.GVA_CONFIG.MCP
	path := config.Path
	if path == "" {
		path = "/mcp"
	}
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	mux := http.NewServeMux()
	httpSrv := &http.Server{
		Handler: mux,
	}

	handler := mcpServer.NewStreamableHTTPServer(
		NewMCPServer(),
		mcpServer.WithHTTPContextFunc(WithHTTPRequestContext),
		mcpServer.WithStreamableHTTPServer(httpSrv),
	)
	mux.Handle(path, handler)
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	return handler
}
