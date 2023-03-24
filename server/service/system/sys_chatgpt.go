package system

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/sashabaranov/go-openai"
)

type ChatGptService struct{}

func (chat *ChatGptService) CreateToken() {

}

func (chat *ChatGptService) DeleteToken() {

}

func (chat *ChatGptService) GetTable(req request.ChatGptRequest) (results []map[string]interface{}) {
	if req.ChatID != "" || req.DBName == "" {
		return
	}
	var tablesInfo []system.ChatField
	global.GVA_DB.Table("information_schema.columns").Where("TABLE_SCHEMA = ?", req.DBName).Scan(&tablesInfo)
	fmt.Println(tablesInfo)
	b, err := json.Marshal(tablesInfo)
	if err != nil {
		return
	}

	client := openai.NewClient("")
	ctx := context.Background()

	chatReq := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: fmt.Sprintf("现在有数据库所有字段用json表示如下,表名为TABLE_NAME字段,列名为COLUMN_NAME字段,列描述为COLUMN_COMMENT字段,%s,根据下方语句帮我生成单纯的查询sql语句，不要提示语\n+%s", string(b), req.Chat),
			},
		},
	}

	resp, err := client.CreateChatCompletion(ctx, chatReq)
	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
		return
	}
	fmt.Printf("用户语句:%s\n\n", req.Chat)

	fmt.Println(resp.Choices[0].Message.Content)

	global.GVA_DB.Raw(resp.Choices[0].Message.Content).Scan(&results)
	return results
}
