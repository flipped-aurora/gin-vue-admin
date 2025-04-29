package tool

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/mcp"
)

type ToolCall struct{}

// 获取当前系统时间
func (t *ToolCall) CallCurrentTime(ctx context.Context, request *protocol.CallToolRequest) (*protocol.CallToolResult, error) {
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

// 获取天气信息
// 参数：城市名称
// 日期: 2023-07-01

func (t *ToolCall) CallCityWeather(ctx context.Context, request *protocol.CallToolRequest) (*protocol.CallToolResult, error) {
	// 解析请求参数
	req := new(model.WeatherReq)
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
