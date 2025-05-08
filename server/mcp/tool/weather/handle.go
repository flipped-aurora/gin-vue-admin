package weather

import (
	"context"
	"encoding/json"
	"github.com/ThinkInAIXYZ/go-mcp/protocol"
)

// 获取天气
func (t *Tool) Handle(ctx context.Context, request *protocol.CallToolRequest) (*protocol.CallToolResult, error) {
	var req Request
	// 解析请求参数
	if err := json.Unmarshal(request.RawArguments, &req); err != nil {
		return nil, err
	}

	//构造回复信息
	str := "天气信息: " + req.Date + " " + req.City + " 天气晴朗，温度为 25°C，相对湿度为 60%，风速为 10m/s。"

	return &protocol.CallToolResult{
		Content: []protocol.Content{
			protocol.TextContent{
				Type: "text",
				Text: str,
			},
		},
	}, nil
}
