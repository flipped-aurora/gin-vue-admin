package mcpTool

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
	"time"
)

func init() {
	RegisterTool(&CurrentTime{})
}

type CurrentTime struct {
}

// 获取当前系统时间
func (t *CurrentTime) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 获取当前系统时间
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	//返回
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: currentTime,
			},
		},
	}, nil
}

func (t *CurrentTime) New() mcp.Tool {
	return mcp.NewTool("currentTime",
		mcp.WithDescription("获取当前系统时间"),
		mcp.WithString("timezone",
			mcp.Required(),
			mcp.Description("时区"),
			mcp.Enum("UTC", "CST", "PST", "EST", "GMT", "CET", "JST", "MST", "IST", "AST", "HST"),
		))
}
