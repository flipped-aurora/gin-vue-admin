package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	t.Run("CurrentTime", func(t *testing.T) {
		clientInfo := protocol.Implementation{Name: "TestClient"}
		mcpClient, err := NewClient("http://127.0.0.1:8888/sse", clientInfo)

		if err != nil {
			t.Error(err)
			return
		}

		tools, err := mcpClient.ListTools(context.Background())
		if err != nil {
			t.Error(err)
			return
		}
		for i := range tools.Tools {
			tool := tools.Tools[i]
			fmt.Printf("工具名称: %s\n", tool.Name)
			fmt.Printf("工具描述: %s\n", tool.Description)

			// 打印参数信息
			if tool.InputSchema.Properties != nil {
				fmt.Println("参数列表:")
				for paramName, prop := range tool.InputSchema.Properties {
					required := "否"
					// 检查参数是否在必填列表中
					for _, reqField := range tool.InputSchema.Required {
						if reqField == paramName {
							required = "是"
							break
						}
					}
					fmt.Printf("  - %s (类型: %s, 描述: %s, 必填: %s)\n",
						paramName, prop.Type, prop.Description, required)
				}
			} else {
				fmt.Println("该工具没有参数")
			}
			fmt.Println("-------------------")
		}

		// 测试调用工具
		res, err := mcpClient.CallTool(context.Background(), &protocol.CallToolRequest{
			Name: "CurrentTime",
			Arguments: map[string]interface{}{
				"timezone": "UTC",
			},
		})
		assert.NoError(t, err)
		assert.NotNil(t, res)
		for _, content := range res.Content {
			if textContent, ok := content.(protocol.TextContent); ok {
				fmt.Printf("文本内容: %s\n", textContent.Text)
			} else {
				fmt.Printf("不是文本: %v\n", content)
			}
		}
	})

	t.Run("GetNickname", func(t *testing.T) {
		clientInfo := protocol.Implementation{Name: "TestClient"}
		mcpClient, err := NewClient("http://127.0.0.1:8888/sse", clientInfo)

		if err != nil {
			t.Error(err)
			return
		}

		// 测试调用工具
		res, err := mcpClient.CallTool(context.Background(), &protocol.CallToolRequest{
			Name: "GetNickname",
			Arguments: map[string]interface{}{
				"username": "admin",
			},
		})

		if err != nil {
			t.Error(err)
			return
		}

		for _, content := range res.Content {
			if textContent, ok := content.(protocol.TextContent); ok {
				fmt.Printf("文本内容: %s\n", textContent.Text)
			} else {
				fmt.Printf("不是文本: %v\n", content)
			}
		}
	})
}
