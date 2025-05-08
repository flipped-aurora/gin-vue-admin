package tool

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

func init() {
	RegisterTool(&GetNickname{})
}

type GetNickname struct {
	Request struct {
		Username string `json:"username" description:"用户的username"`
	}
}

// 获取天气
func (t *GetNickname) New() (name string, description string, inputReqStruct interface{}) {
	name = "GetNickname"
	description = "根据用户username获取nickname"
	inputReqStruct = t.Request //自动解析参数
	return
}

// 获取天气
func (t *GetNickname) Handle(ctx context.Context, request *protocol.CallToolRequest) (*protocol.CallToolResult, error) {
	// 解析请求参数
	if err := json.Unmarshal(request.RawArguments, &t.Request); err != nil {
		return nil, err
	}

	var user system.SysUser

	err := global.GVA_DB.First(&user, "username = ?", t.Request.Username).Error

	if err != nil {
		return nil, err
	}

	//构造回复信息
	return &protocol.CallToolResult{
		Content: []protocol.Content{
			protocol.TextContent{
				Type: "text",
				Text: fmt.Sprintf("%s的昵称是%s", t.Request.Username, user.NickName),
			},
		},
	}, nil
}
