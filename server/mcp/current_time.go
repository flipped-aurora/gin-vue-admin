package mcpTool

import (
	"context"
	"errors"
	"fmt"
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

	timezone, ok := request.Params.Arguments["timezone"].(string)

	if !ok {
		return nil, errors.New("参数错误：timezone 必须是字符串类型")
	}
	// 根据timezone参数加载对应时区
	loc, err := loadTimeZone(timezone)
	if err != nil {
		return nil, err
	}

	// 获取当前时间并转换为指定时区
	currentTime := time.Now().In(loc).Format("2006-01-02 15:04:05")
	//返回
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: fmt.Sprintf("%s 时区的当前时间是：%s", timezone, currentTime),
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

// 将简写时区转换为IANA标准时区
func loadTimeZone(timezone string) (*time.Location, error) {
	// 时区映射表
	timezoneMap := map[string]string{
		"UTC": "UTC",
		"CST": "Asia/Shanghai", // 中国标准时间
		"PST": "America/Los_Angeles",
		"EST": "America/New_York",
		"GMT": "GMT",
		"CET": "Europe/Paris",
		"JST": "Asia/Tokyo",
		"MST": "America/Denver",
		"IST": "Asia/Kolkata",
		"AST": "Asia/Riyadh", // 阿拉伯标准时间
		"HST": "Pacific/Honolulu",
	}

	// 获取标准时区名称
	tzName, exists := timezoneMap[timezone]
	if !exists {
		return nil, errors.New("不支持的时区: " + timezone)
	}

	// 加载时区
	return time.LoadLocation(tzName)
}
