package mcpTool

import (
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func textResultWithJSON(title string, payload any) (*mcp.CallToolResult, error) {
	resultJSON, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("序列化结果失败: %w", err)
	}

	text := string(resultJSON)
	if title != "" {
		text = fmt.Sprintf("%s\n\n%s", title, text)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: text,
			},
		},
	}, nil
}
