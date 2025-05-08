package tool

import (
	"context"
	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"time"
)

func init() {
	RegisterTool(&CurrentTime{})
}

type CurrentTime struct {
	Request struct {
		Timezone string `json:"timezone" description:"current time timezone"`
	}
}

// 获取当前系统时间
func (t *CurrentTime) Handle(ctx context.Context, request *protocol.CallToolRequest) (*protocol.CallToolResult, error) {
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

func (t *CurrentTime) New() (name string, description string, inputReqStruct interface{}) {
	name = "CurrentTime"
	description = "获取当前系统时间"
	inputReqStruct = t.Request
	return
}
