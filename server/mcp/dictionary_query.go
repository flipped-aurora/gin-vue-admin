package mcpTool

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/mark3labs/mcp-go/mcp"
)

func init() {
	RegisterTool(&DictionaryQuery{})
}

type DictionaryPre struct {
	Type string `json:"type"`
	Desc string `json:"desc"`
}

type DictionaryInfo struct {
	ID      uint                   `json:"id"`
	Name    string                 `json:"name"`
	Type    string                 `json:"type"`
	Status  *bool                  `json:"status"`
	Desc    string                 `json:"desc"`
	Details []DictionaryDetailInfo `json:"details"`
}

type DictionaryDetailInfo struct {
	ID     uint   `json:"id"`
	Label  string `json:"label"`
	Value  string `json:"value"`
	Extend string `json:"extend"`
	Status *bool  `json:"status"`
	Sort   int    `json:"sort"`
}

type DictionaryQueryResponse struct {
	Success      bool             `json:"success"`
	Message      string           `json:"message"`
	Total        int              `json:"total"`
	Dictionaries []DictionaryInfo `json:"dictionaries"`
}

type DictionaryQuery struct{}

func (d *DictionaryQuery) New() mcp.Tool {
	return mcp.NewTool("query_dictionaries",
		mcp.WithDescription("查询系统中所有的字典和字典属性，用于AI生成逻辑时了解可用的字典选项"),
		mcp.WithString("dictType",
			mcp.Description("可选：指定字典类型进行精确查询，如果不提供则返回所有字典"),
		),
		mcp.WithBoolean("includeDisabled",
			mcp.Description("是否包含已禁用的字典和字典项，默认为false（只返回启用的）"),
		),
		mcp.WithBoolean("detailsOnly",
			mcp.Description("是否只返回字典详情信息（不包含字典基本信息），默认为false"),
		),
	)
}

func (d *DictionaryQuery) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := request.GetArguments()

	dictType := stringValue(args["dictType"])
	includeDisabled, _ := args["includeDisabled"].(bool)
	detailsOnly, _ := args["detailsOnly"].(bool)

	dictionaries, err := fetchDictionaryList(ctx, dictType)
	if err != nil {
		return nil, err
	}

	result := make([]DictionaryInfo, 0)
	for _, dictionary := range dictionaries {
		if dictType != "" && dictionary.Type != dictType {
			continue
		}
		if !includeDisabled && dictionary.Status != nil && !*dictionary.Status {
			continue
		}

		dictInfo, err := buildDictionaryInfo(ctx, dictionary, includeDisabled)
		if err != nil {
			return nil, err
		}
		result = append(result, dictInfo)
	}

	if detailsOnly {
		details := make([]DictionaryDetailInfo, 0)
		for _, dictionary := range result {
			details = append(details, dictionary.Details...)
		}
		return textResultWithJSON("", map[string]any{
			"success": true,
			"message": "查询字典详情成功",
			"total":   len(details),
			"details": details,
		})
	}

	return textResultWithJSON("", DictionaryQueryResponse{
		Success:      true,
		Message:      "查询字典成功",
		Total:        len(result),
		Dictionaries: result,
	})
}

func buildDictionaryInfo(ctx context.Context, dictionary system.SysDictionary, includeDisabled bool) (DictionaryInfo, error) {
	exported, err := exportDictionary(ctx, dictionary.ID)
	if err != nil {
		return DictionaryInfo{}, err
	}

	info := DictionaryInfo{
		ID:     dictionary.ID,
		Name:   exported.Name,
		Type:   exported.Type,
		Status: exported.Status,
		Desc:   exported.Desc,
	}

	for _, detail := range exported.SysDictionaryDetails {
		if !includeDisabled && detail.Status != nil && !*detail.Status {
			continue
		}
		info.Details = append(info.Details, DictionaryDetailInfo{
			ID:     detail.ID,
			Label:  detail.Label,
			Value:  detail.Value,
			Extend: detail.Extend,
			Status: detail.Status,
			Sort:   detail.Sort,
		})
	}

	return info, nil
}
