package mcp

import (
	"context"
	"fmt"
	"time"

	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"github.com/ThinkInAIXYZ/go-mcp/server"
	"github.com/ThinkInAIXYZ/go-mcp/transport"

	model "github.com/flipped-aurora/gin-vue-admin/server/model/mcp"
	McpTool "github.com/flipped-aurora/gin-vue-admin/server/service/mcp/service/tool"
)

type McpServer struct{}

var mcpServer *server.Server

type tool struct {
	ListTool func() (name string, description string, inputReqStruct interface{})
	CallTool func(context.Context, *protocol.CallToolRequest) (*protocol.CallToolResult, error)
}

var (
	ToolCall = new(McpTool.ToolCall)
	ToolList = new(McpTool.ToolList)
)

func (s *McpServer) NewSSE(Config *model.McpConfig) error {
	transportServer, err := transport.NewSSEServerTransport(Config.Url)
	if err != nil {
		return fmt.Errorf("创建mcp服务器失败: %v", err)
	}
	mcpServer, err = server.NewServer(transportServer,
		server.WithServerInfo(protocol.Implementation{
			Name:    Config.Name,
			Version: Config.Version,
		}),
	)
	if err != nil {
		return fmt.Errorf("创建mcp服务器失败: %v", err)
	}
	return nil
}

func (s *McpServer) BinTool() {
	if mcpServer == nil {
		fmt.Println("mcp服务器: MCP服务器未初始化")
		return
	}

	b := []tool{
		{ToolList.CityWeather, ToolCall.CallCityWeather},
		{ToolList.CurrentTime, ToolCall.CallCurrentTime},
	}
	for _, v := range b {
		s.RegisterTool(v)
	}
	fmt.Println("mcp服务器: 注册工具成功 ")

}

func (s *McpServer) RegisterTool(B tool) {
	if mcpServer == nil {
		fmt.Println("mcp服务器: MCP服务器未初始化")
		return
	}

	tt, err := protocol.NewTool(B.ListTool())
	if err != nil {
		panic(fmt.Errorf("创建mcp工具失败: %v", err))
	}

	mcpServer.RegisterTool(tt, B.CallTool)
}

func (s *McpServer) Run() (err error) {

	//启动服务器
	go func() {
		err = mcpServer.Run()
	}()

	time.Sleep(1000 * time.Millisecond)
	if mcpServer == nil || err != nil {
		return fmt.Errorf("mcp服务器: MCP服务器未初始化")
	}
	fmt.Println("mcp服务器: 启动成功 ")
	select {} // 阻塞主线程,保持运行
}
