package mcpTool

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/mark3labs/mcp-go/mcp"
)

func init() {
	RegisterTool(&DictionaryOptionsGenerator{})
}

type DictionaryOptionsGenerator struct{}

type DictionaryOption struct {
	Label string `json:"label"`
	Value string `json:"value"`
	Sort  int    `json:"sort"`
}

type DictionaryGenerateRequest struct {
	DictType    string             `json:"dictType"`
	FieldDesc   string             `json:"fieldDesc"`
	Options     []DictionaryOption `json:"options"`
	DictName    string             `json:"dictName"`
	Description string             `json:"description"`
}

type DictionaryGenerateResponse struct {
	Success      bool   `json:"success"`
	Message      string `json:"message"`
	DictType     string `json:"dictType"`
	OptionsCount int    `json:"optionsCount"`
}

func (d *DictionaryOptionsGenerator) New() mcp.Tool {
	return mcp.NewTool("generate_dictionary_options",
		mcp.WithDescription("智能生成字典选项并自动创建字典和字典详情"),
		mcp.WithString("dictType",
			mcp.Required(),
			mcp.Description("字典类型，用于标识字典的唯一性"),
		),
		mcp.WithString("fieldDesc",
			mcp.Required(),
			mcp.Description("字段描述，用于AI理解字段含义"),
		),
		mcp.WithString("options",
			mcp.Required(),
			mcp.Description("字典选项JSON字符串，格式：[{\"label\":\"显示名\",\"value\":\"值\",\"sort\":1}]"),
		),
		mcp.WithString("dictName",
			mcp.Description("字典名称，如果不提供将自动生成"),
		),
		mcp.WithString("description",
			mcp.Description("字典描述"),
		),
	)
}

func (d *DictionaryOptionsGenerator) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := request.GetArguments()

	dictType, ok := args["dictType"].(string)
	if !ok || dictType == "" {
		return nil, errors.New("dictType 参数是必需的")
	}
	fieldDesc, ok := args["fieldDesc"].(string)
	if !ok || fieldDesc == "" {
		return nil, errors.New("fieldDesc 参数是必需的")
	}
	optionsStr, ok := args["options"].(string)
	if !ok || optionsStr == "" {
		return nil, errors.New("options 参数是必需的")
	}

	var options []DictionaryOption
	if err := json.Unmarshal([]byte(optionsStr), &options); err != nil {
		return nil, fmt.Errorf("options 参数格式错误: %v", err)
	}
	if len(options) == 0 {
		return nil, errors.New("options 不能为空")
	}

	req := &DictionaryGenerateRequest{
		DictType:    dictType,
		FieldDesc:   fieldDesc,
		Options:     options,
		DictName:    stringValue(args["dictName"]),
		Description: stringValue(args["description"]),
	}

	result, err := d.createDictionaryWithOptions(ctx, req)
	if err != nil {
		return nil, err
	}

	return textResultWithJSON("字典选项生成结果：", result)
}

func (d *DictionaryOptionsGenerator) createDictionaryWithOptions(ctx context.Context, req *DictionaryGenerateRequest) (*DictionaryGenerateResponse, error) {
	existingDict, err := findDictionaryByType(ctx, req.DictType)
	if err != nil {
		return nil, fmt.Errorf("检查字典是否存在失败: %v", err)
	}
	if existingDict != nil {
		return &DictionaryGenerateResponse{
			Success:      false,
			Message:      fmt.Sprintf("字典 %s 已存在，跳过创建", req.DictType),
			DictType:     req.DictType,
			OptionsCount: 0,
		}, nil
	}

	dictName := req.DictName
	if dictName == "" {
		dictName = d.generateDictionaryName(req.DictType, req.FieldDesc)
	}

	if err := createDictionary(ctx, system.SysDictionary{
		Name:   dictName,
		Type:   req.DictType,
		Status: enabledBoolPointer(),
		Desc:   req.Description,
	}); err != nil {
		return nil, fmt.Errorf("创建字典失败: %v", err)
	}

	createdDict, err := findDictionaryByType(ctx, req.DictType)
	if err != nil {
		return nil, fmt.Errorf("获取创建的字典失败: %v", err)
	}
	if createdDict == nil {
		return nil, fmt.Errorf("获取创建的字典失败")
	}

	successCount := 0
	for _, option := range req.Options {
		err := createDictionaryDetail(ctx, system.SysDictionaryDetail{
			Label:           option.Label,
			Value:           option.Value,
			Status:          enabledBoolPointer(),
			Sort:            option.Sort,
			SysDictionaryID: int(createdDict.ID),
		})
		if err == nil {
			successCount++
		}
	}

	return &DictionaryGenerateResponse{
		Success:      true,
		Message:      fmt.Sprintf("成功创建字典 %s，包含 %d 个选项", req.DictType, successCount),
		DictType:     req.DictType,
		OptionsCount: successCount,
	}, nil
}

func (d *DictionaryOptionsGenerator) generateDictionaryName(dictType, fieldDesc string) string {
	if fieldDesc != "" {
		return fmt.Sprintf("%s字典", fieldDesc)
	}
	return fmt.Sprintf("%s字典", dictType)
}

func stringValue(value any) string {
	if str, ok := value.(string); ok {
		return str
	}
	return ""
}
