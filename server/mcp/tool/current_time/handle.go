package current_time

import (
	"context"
	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"time"
)

// 获取当前系统时间
func (t *Tool) Handle(ctx context.Context, request *protocol.CallToolRequest) (*protocol.CallToolResult, error) {
	// 获取当前系统时间
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	//返回
	return &protocol.CallToolResult{
		Content: []protocol.Content{
			protocol.TextContent{
				Type: "text",
				Text: currentTime,
			},
		},
	}, nil
}
