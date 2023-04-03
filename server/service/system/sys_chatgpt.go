package system

import (
	"context"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/sashabaranov/go-openai"
	"gorm.io/gorm"
	"strings"
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
	var tableName string
	global.GVA_DB.Table("information_schema.columns").Where("TABLE_SCHEMA = ?", req.DBName).Scan(&tablesInfo)

	var tablesMap = make(map[string]bool)
	for i := range tablesInfo {
		tablesMap[tablesInfo[i].TABLE_NAME] = true
	}
	for i := range tablesMap {
		tableName += i + ","
	}
	option, err := chat.GetSK()
	if err != nil {
		return "", nil, err
	}
	client := openai.NewClient(option.SK)
	ctx := context.Background()

	tables, err := getTables(ctx, client, tableName, req.Chat)
	if err != nil {
		return "", nil, err
	}
	tableArr := strings.Split(tables, ",")
	if len(tableArr) != 0 {
		firstKeyArr := strings.Split(tableArr[0], ":")
		tableArr[0] = strings.Trim(firstKeyArr[len(firstKeyArr)-1], "\n")
	}
	sql, err = getSql(ctx, client, tableArr, tablesInfo, req.Chat)
	if err != nil {
		return "", nil, err
	}
	err = global.GVA_DB.Raw(sql).Scan(&results).Error
	return sql, results, err
}

func getTables(ctx context.Context, client *openai.Client, tables string, chat string) (string, error) {
	var tablePrompt = `You are a database administrator

Filter out the table names you might need from the tables I provided formatted as:

Table1,Table2,Table3

I will provide you with the following table configuration information:

Table1,Table2,Table3

Do not return information other than the table

Configured as:
%s

The problem is:
%s
`
	content := fmt.Sprintf(tablePrompt, tables, chat)
	chatReq := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: content,
			},
		},
	}

	resp, err := client.CreateChatCompletion(ctx, chatReq)
	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}

func getSql(ctx context.Context, client *openai.Client, tables []string, ChatField []system.ChatField, chat string) (string, error) {
	var sqlPrompt = `You are a database administrator

Give me an SQL statement based on my question

I will provide you with my current database table configuration information in the form below

Table Name | Column Name | Column Description

Do not return information other than SQL

Configured as:

%s

The problem is:

%s`
	var configured string
	for ii := range ChatField {
		for i := range tables {
			if strings.Index(tables[i], ChatField[ii].TABLE_NAME) > -1 {
				configured += fmt.Sprintf("%s | %s | %s \n", ChatField[ii].TABLE_NAME, ChatField[ii].COLUMN_NAME, ChatField[ii].COLUMN_COMMENT)
			}
		}
	}

	if configured == "" {
		return "", errors.New("未找到表")
	}
	chatReq := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: fmt.Sprintf(sqlPrompt, configured, chat),
			},
		},
	}

	resp, err := client.CreateChatCompletion(ctx, chatReq)
	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
		return "", err
	}
	sql := resp.Choices[0].Message.Content
	sqlArr := strings.Split(sql, ":")
	sql = sqlArr[len(sqlArr)-1]
	return sql, nil
}
