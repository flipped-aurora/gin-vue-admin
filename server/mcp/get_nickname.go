package mcpTool

import (
	"context"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/mark3labs/mcp-go/mcp"
)

func init() {
	RegisterTool(&GetNickname{})
}

type GetNickname struct{}

// 根据用户username获取nickname
func (t *GetNickname) New() mcp.Tool {
	return mcp.NewTool("getNickname",
		mcp.WithDescription("根据用户username获取nickname"),
		mcp.WithString("username",
			mcp.Required(),
			mcp.Description("用户的username"),
		))
}

// 根据用户username获取nickname
func (t *GetNickname) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, ok := request.Params.Arguments["username"].(string)
	if !ok {
		return nil, errors.New("name must be a string")
	}

	var user system.SysUser

	err := global.GVA_DB.First(&user, "username = ?", name).Error

	if err != nil {
		return nil, err
	}

	//构造回复信息
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: fmt.Sprintf("%s的昵称是%s", name, user.NickName),
			},
		},
	}, nil
}
