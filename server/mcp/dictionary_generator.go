package mcpTool

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/mark3labs/mcp-go/mcp"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func init() {
	RegisterTool(&DictionaryOptionsGenerator{})
}

// DictionaryOptionsGenerator 字典选项生成器
type DictionaryOptionsGenerator struct{}

// DictionaryGenerateRequest 字典生成请求
type DictionaryGenerateRequest struct {
	DictType    string             `json:"dictType"`    // 字典类型
	FieldDesc   string             `json:"fieldDesc"`   // 字段描述
	Options     []DictionaryOption `json:"options"`     // AI生成的字典选项
	DictName    string             `json:"dictName"`    // 字典名称（可选）
	Description string             `json:"description"` // 字典描述（可选）
}

// DictionaryGenerateResponse 字典生成响应
type DictionaryGenerateResponse struct {
	Success      bool   `json:"success"`
	Message      string `json:"message"`
	DictType     string `json:"dictType"`
	OptionsCount int    `json:"optionsCount"`
}

// New 返回工具注册信息
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

// Name 返回工具名称
func (d *DictionaryOptionsGenerator) Name() string {
	return "generate_dictionary_options"
}

// Description 返回工具描述
func (d *DictionaryOptionsGenerator) Description() string {
	return `字典选项生成工具 - 让AI生成并创建字典选项

此工具允许AI根据字典类型和字段描述生成合适的字典选项，并自动创建字典和字典详情。

参数说明：
- dictType: 字典类型（必填）
- fieldDesc: 字段描述（必填）
- options: AI生成的字典选项数组（必填）
  - label: 选项标签
  - value: 选项值
  - sort: 排序号
- dictName: 字典名称（可选，默认根据fieldDesc生成）
- description: 字典描述（可选）

使用场景：
1. 在创建模块时，如果字段需要字典类型，AI可以根据字段描述智能生成合适的选项
2. 支持各种业务场景的字典选项生成，如状态、类型、等级等
3. 自动创建字典和字典详情，无需手动配置

示例调用：
{
  "dictType": "user_status",
  "fieldDesc": "用户状态",
  "options": [
    {"label": "正常", "value": "1", "sort": 1},
    {"label": "禁用", "value": "0", "sort": 2}
  ],
  "dictName": "用户状态字典",
  "description": "用于管理用户账户状态的字典"
}`
}

// InputSchema 返回输入参数的JSON Schema
func (d *DictionaryOptionsGenerator) InputSchema() map[string]interface{} {
	return map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"dictType": map[string]interface{}{
				"type":        "string",
				"description": "字典类型，用于标识字典的唯一性",
			},
			"fieldDesc": map[string]interface{}{
				"type":        "string",
				"description": "字段描述，用于生成字典名称和理解字典用途",
			},
			"options": map[string]interface{}{
				"type":        "array",
				"description": "AI生成的字典选项数组",
				"items": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"label": map[string]interface{}{
							"type":        "string",
							"description": "选项标签，显示给用户的文本",
						},
						"value": map[string]interface{}{
							"type":        "string",
							"description": "选项值，存储在数据库中的值",
						},
						"sort": map[string]interface{}{
							"type":        "integer",
							"description": "排序号，用于控制选项显示顺序",
						},
					},
					"required": []string{"label", "value", "sort"},
				},
			},
			"dictName": map[string]interface{}{
				"type":        "string",
				"description": "字典名称，可选，默认根据fieldDesc生成",
			},
			"description": map[string]interface{}{
				"type":        "string",
				"description": "字典描述，可选",
			},
		},
		"required": []string{"dictType", "fieldDesc", "options"},
	}
}

// Handle 处理工具调用
func (d *DictionaryOptionsGenerator) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 解析请求参数
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

	// 解析options JSON字符串
	var options []DictionaryOption
	if err := json.Unmarshal([]byte(optionsStr), &options); err != nil {
		return nil, fmt.Errorf("options 参数格式错误: %v", err)
	}

	if len(options) == 0 {
		return nil, errors.New("options 不能为空")
	}

	// 可选参数
	dictName, _ := args["dictName"].(string)
	description, _ := args["description"].(string)

	// 构建请求对象
	req := &DictionaryGenerateRequest{
		DictType:    dictType,
		FieldDesc:   fieldDesc,
		Options:     options,
		DictName:    dictName,
		Description: description,
	}

	// 创建字典
	response, err := d.createDictionaryWithOptions(ctx, req)
	if err != nil {
		return nil, err
	}

	// 构建响应
	resultJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("序列化结果失败: %v", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: fmt.Sprintf("字典选项生成结果：\n\n%s", string(resultJSON)),
			},
		},
	}, nil
}

// createDictionaryWithOptions 创建字典和字典选项
func (d *DictionaryOptionsGenerator) createDictionaryWithOptions(ctx context.Context, req *DictionaryGenerateRequest) (*DictionaryGenerateResponse, error) {
	// 检查字典是否已存在
	exists, err := d.checkDictionaryExists(req.DictType)
	if err != nil {
		return nil, fmt.Errorf("检查字典是否存在失败: %v", err)
	}

	if exists {
		return &DictionaryGenerateResponse{
			Success:      false,
			Message:      fmt.Sprintf("字典 %s 已存在，跳过创建", req.DictType),
			DictType:     req.DictType,
			OptionsCount: 0,
		}, nil
	}

	// 生成字典名称
	dictName := req.DictName
	if dictName == "" {
		dictName = d.generateDictionaryName(req.DictType, req.FieldDesc)
	}

	// 创建字典
	dictionaryService := service.ServiceGroupApp.SystemServiceGroup.DictionaryService
	dictionary := system.SysDictionary{
		Name:   dictName,
		Type:   req.DictType,
		Status: &[]bool{true}[0], // 默认启用
		Desc:   req.Description,
	}

	err = dictionaryService.CreateSysDictionary(dictionary)
	if err != nil {
		return nil, fmt.Errorf("创建字典失败: %v", err)
	}

	// 获取刚创建的字典ID
	var createdDict system.SysDictionary
	err = global.GVA_DB.Where("type = ?", req.DictType).First(&createdDict).Error
	if err != nil {
		return nil, fmt.Errorf("获取创建的字典失败: %v", err)
	}

	// 创建字典详情项
	dictionaryDetailService := service.ServiceGroupApp.SystemServiceGroup.DictionaryDetailService
	successCount := 0

	for _, option := range req.Options {
		dictionaryDetail := system.SysDictionaryDetail{
			Label:           option.Label,
			Value:           option.Value,
			Status:          &[]bool{true}[0], // 默认启用
			Sort:            option.Sort,
			SysDictionaryID: int(createdDict.ID),
		}

		err = dictionaryDetailService.CreateSysDictionaryDetail(dictionaryDetail)
		if err != nil {
			global.GVA_LOG.Warn("创建字典详情项失败", zap.Error(err))
		} else {
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

// checkDictionaryExists 检查字典是否存在
func (d *DictionaryOptionsGenerator) checkDictionaryExists(dictType string) (bool, error) {
	var dictionary system.SysDictionary
	err := global.GVA_DB.Where("type = ?", dictType).First(&dictionary).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil // 字典不存在
		}
		return false, err // 其他错误
	}
	return true, nil // 字典存在
}

// generateDictionaryName 生成字典名称
func (d *DictionaryOptionsGenerator) generateDictionaryName(dictType, fieldDesc string) string {
	if fieldDesc != "" {
		return fmt.Sprintf("%s字典", fieldDesc)
	}
	return fmt.Sprintf("%s字典", dictType)
}
