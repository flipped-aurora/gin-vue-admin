package client

import (
	"context"
	"fmt"
	"slices"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
)

// 测试 MCP 客户端连接
func TestMcpClientConnection(t *testing.T) {
	c, err := NewClient("http://localhost:8888/sse", "test-client", "1.0.0", "gin-vue-admin MCP服务")
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
}

func TestTools(t *testing.T) {
	t.Run("currentTime", func(t *testing.T) {
		c, err := NewClient("http://localhost:8888/sse", "test-client", "1.0.0", "gin-vue-admin MCP服务")
		if err != nil {
			t.Fatalf("Failed to create client: %v", err)
		}
		defer c.Close()
		ctx := context.Background()

		request := mcp.CallToolRequest{}
		request.Params.Name = "currentTime"
		request.Params.Arguments = map[string]any{
			"timezone": "UTC+8",
		}

		result, err := c.CallTool(ctx, request)
		if err != nil {
			t.Fatalf("方法调用错误: %v", err)
		}

		if len(result.Content) != 1 {
			t.Errorf("应该有且仅返回1条信息，但是现在有 %d", len(result.Content))
		}
		if content, ok := result.Content[0].(mcp.TextContent); ok {
			t.Logf("成功返回信息%s", content.Text)
		} else {
			t.Logf("返回为止类型信息%+v", content)
		}
	})

	t.Run("getNickname", func(t *testing.T) {

		c, err := NewClient("http://localhost:8888/sse", "test-client", "1.0.0", "gin-vue-admin MCP服务")
		if err != nil {
			t.Fatalf("Failed to create client: %v", err)
		}
		defer c.Close()
		ctx := context.Background()

		// Initialize
		initRequest := mcp.InitializeRequest{}
		initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
		initRequest.Params.ClientInfo = mcp.Implementation{
			Name:    "test-client",
			Version: "1.0.0",
		}

		_, err = c.Initialize(ctx, initRequest)
		if err != nil {
			t.Fatalf("初始化失败: %v", err)
		}

		request := mcp.CallToolRequest{}
		request.Params.Name = "getNickname"
		request.Params.Arguments = map[string]any{
			"username": "admin",
		}

		result, err := c.CallTool(ctx, request)
		if err != nil {
			t.Fatalf("方法调用错误: %v", err)
		}

		if len(result.Content) != 1 {
			t.Errorf("应该有且仅返回1条信息，但是现在有 %d", len(result.Content))
		}
		if content, ok := result.Content[0].(mcp.TextContent); ok {
			t.Logf("成功返回信息%s", content.Text)
		} else {
			t.Logf("返回为止类型信息%+v", content)
		}
	})
}

func TestGetTools(t *testing.T) {
	c, err := NewClient("http://localhost:8888/sse", "test-client", "1.0.0", "gin-vue-admin MCP服务")
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	defer c.Close()
	ctx := context.Background()

	toolsRequest := mcp.ListToolsRequest{}

	toolListResult, err := c.ListTools(ctx, toolsRequest)
	if err != nil {
		t.Fatalf("获取工具列表失败: %v", err)
	}
	for i := range toolListResult.Tools {
		tool := toolListResult.Tools[i]
		fmt.Printf("工具名称: %s\n", tool.Name)
		fmt.Printf("工具描述: %s\n", tool.Description)

		// 打印参数信息
		if tool.InputSchema.Properties != nil {
			fmt.Println("参数列表:")
			for paramName, prop := range tool.InputSchema.Properties {
				required := "否"
				// 检查参数是否在必填列表中
				if slices.Contains(tool.InputSchema.Required, paramName) {
					required = "是"
				}
				fmt.Printf("  - %s (类型: %s, 描述: %s, 必填: %s)\n",
					paramName, prop.(map[string]any)["type"], prop.(map[string]any)["description"], required)
			}
		} else {
			fmt.Println("该工具没有参数")
		}
		fmt.Println("-------------------")
	}
}
