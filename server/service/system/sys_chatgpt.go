package system

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/sashabaranov/go-openai"
	"gorm.io/gorm"
)

type ChatGptService struct{}

func (chat *ChatGptService) CreateSK(option system.SysChatGptOption) error {
	_, err := chat.GetSK()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return global.GVA_DB.Create(option).Error
		}
		return err
	}
	return errors.New("已经存在sk")
}

func (chat *ChatGptService) GetSK() (option system.SysChatGptOption, err error) {
	err = global.GVA_DB.First(&option).Error
	return
}

func (chat *ChatGptService) DeleteSK() error {
	option, err := chat.GetSK()
	if err != nil {
		return err
	}
	return global.GVA_DB.Delete(option, "sk = ?", option.SK).Error
}

func (chat *ChatGptService) GetTable(req request.ChatGptRequest) (sql string, results []map[string]interface{}, err error) {
	if req.DBName == "" {
		return "", nil, errors.New("未选择db")
	}
	var tablesInfo []system.ChatField
	global.GVA_DB.Table("information_schema.columns").Where("TABLE_SCHEMA = ?", req.DBName).Scan(&tablesInfo)
	b, err := json.Marshal(tablesInfo)
	if err != nil {
		return
	}
	option, err := chat.GetSK()
	if err != nil {
		return "", nil, err
	}
	client := openai.NewClient(option.SK)
	ctx := context.Background()

	chatReq := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: fmt.Sprintf("数据库所有字段用json表示,表名为TABLE_NAME,列名为COLUMN_NAME,列描述为COLUMN_COMMENT,%s,根据语句帮我生成单纯的查询sql,,不要提示语\n+%s", string(b), req.Chat),
			},
		},
	}

	resp, err := client.CreateChatCompletion(ctx, chatReq)
	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
		return
	}

	err = global.GVA_DB.Raw(resp.Choices[0].Message.Content).Scan(&results).Error
	return resp.Choices[0].Message.Content, results, err
}
