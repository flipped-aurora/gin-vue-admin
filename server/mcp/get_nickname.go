package mcpTool

import (
	"context"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/mark3labs/mcp-go/mcp"
	"gorm.io/gorm"
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

// Handle 处理获取昵称的请求
func (t *GetNickname) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 1. 参数验证
	username, ok := request.Params.Arguments["username"].(string)
	if !ok {
		return nil, errors.New("参数错误：username 必须是字符串类型")
	}

	if username == "" {
		return nil, errors.New("参数错误：username 不能为空")
	}

	// 2. 记录操作日志
	global.GVA_LOG.Info("getNickname 工具被调用")

	// 3. 优化查询，只选择需要的字段
	var user struct {
		NickName string
	}

	err := global.GVA_DB.Model(&system.SysUser{}).
		Select("nick_name").
		Where("username = ?", username).
		First(&user).Error

	// 4. 优化错误处理
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					mcp.TextContent{
						Type: "text",
						Text: fmt.Sprintf("用户 %s 不存在", username),
					},
				},
			}, nil
		}
		global.GVA_LOG.Error("数据库查询错误")
		return nil, errors.New("系统错误，请稍后再试")
	}

	// 构造回复信息
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: fmt.Sprintf("用户 %s 的昵称是 %s", username, user.NickName),
			},
		},
	}, nil
}
