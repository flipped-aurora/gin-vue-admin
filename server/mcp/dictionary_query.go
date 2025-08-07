package mcpTool

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/mark3labs/mcp-go/mcp"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 注册工具
func init() {
	RegisterTool(&DictionaryQuery{})
}

// DictionaryInfo 字典信息结构
type DictionaryInfo struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`   // 字典名（中）
	Type   string `json:"type"`   // 字典名（英）
	Status *bool  `json:"status"` // 状态
	Desc   string `json:"desc"`   // 描述
	Details []DictionaryDetailInfo `json:"details"` // 字典详情
}

// DictionaryDetailInfo 字典详情信息结构
type DictionaryDetailInfo struct {
	ID     uint   `json:"id"`
	Label  string `json:"label"`  // 展示值
	Value  string `json:"value"`  // 字典值
	Extend string `json:"extend"` // 扩展值
	Status *bool  `json:"status"` // 启用状态
	Sort   int    `json:"sort"`   // 排序标记
}

// DictionaryQueryResponse 字典查询响应结构
type DictionaryQueryResponse struct {
	Success     bool             `json:"success"`
	Message     string           `json:"message"`
	Total       int              `json:"total"`
	Dictionaries []DictionaryInfo `json:"dictionaries"`
}

// DictionaryQuery 字典查询工具
type DictionaryQuery struct{}

// New 创建字典查询工具
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

// Handle 处理字典查询请求
func (d *DictionaryQuery) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := request.GetArguments()
	
	// 获取参数
	dictType := ""
	if val, ok := args["dictType"].(string); ok {
		dictType = val
	}
	
	includeDisabled := false
	if val, ok := args["includeDisabled"].(bool); ok {
		includeDisabled = val
	}
	
	detailsOnly := false
	if val, ok := args["detailsOnly"].(bool); ok {
		detailsOnly = val
	}
	
	// 获取字典服务
	dictionaryService := service.ServiceGroupApp.SystemServiceGroup.DictionaryService
	
	var dictionaries []DictionaryInfo
	var err error
	
	if dictType != "" {
		// 查询指定类型的字典
		var status *bool
		if !includeDisabled {
			status = &[]bool{true}[0]
		}
		
		sysDictionary, err := dictionaryService.GetSysDictionary(dictType, 0, status)
		if err != nil {
			global.GVA_LOG.Error("查询字典失败", zap.Error(err))
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					mcp.NewTextContent(fmt.Sprintf(`{"success": false, "message": "查询字典失败: %v", "total": 0, "dictionaries": []}`, err.Error())),
				},
			}, nil
		}
		
		// 转换为响应格式
		dictInfo := DictionaryInfo{
			ID:     sysDictionary.ID,
			Name:   sysDictionary.Name,
			Type:   sysDictionary.Type,
			Status: sysDictionary.Status,
			Desc:   sysDictionary.Desc,
		}
		
		// 获取字典详情
		for _, detail := range sysDictionary.SysDictionaryDetails {
			if includeDisabled || (detail.Status != nil && *detail.Status) {
				dictInfo.Details = append(dictInfo.Details, DictionaryDetailInfo{
					ID:     detail.ID,
					Label:  detail.Label,
					Value:  detail.Value,
					Extend: detail.Extend,
					Status: detail.Status,
					Sort:   detail.Sort,
				})
			}
		}
		
		dictionaries = append(dictionaries, dictInfo)
	} else {
		// 查询所有字典
		var sysDictionaries []system.SysDictionary
		db := global.GVA_DB.Model(&system.SysDictionary{})
		
		if !includeDisabled {
			db = db.Where("status = ?", true)
		}
		
		err = db.Preload("SysDictionaryDetails", func(db *gorm.DB) *gorm.DB {
			if includeDisabled {
				return db.Order("sort")
			} else {
				return db.Where("status = ?", true).Order("sort")
			}
		}).Find(&sysDictionaries).Error
		
		if err != nil {
			global.GVA_LOG.Error("查询字典列表失败", zap.Error(err))
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					mcp.NewTextContent(fmt.Sprintf(`{"success": false, "message": "查询字典列表失败: %v", "total": 0, "dictionaries": []}`, err.Error())),
				},
			}, nil
		}
		
		// 转换为响应格式
		for _, dict := range sysDictionaries {
			dictInfo := DictionaryInfo{
				ID:     dict.ID,
				Name:   dict.Name,
				Type:   dict.Type,
				Status: dict.Status,
				Desc:   dict.Desc,
			}
			
			// 获取字典详情
			for _, detail := range dict.SysDictionaryDetails {
				if includeDisabled || (detail.Status != nil && *detail.Status) {
					dictInfo.Details = append(dictInfo.Details, DictionaryDetailInfo{
						ID:     detail.ID,
						Label:  detail.Label,
						Value:  detail.Value,
						Extend: detail.Extend,
						Status: detail.Status,
						Sort:   detail.Sort,
					})
				}
			}
			
			dictionaries = append(dictionaries, dictInfo)
		}
	}
	
	// 如果只需要详情信息，则提取所有详情
	if detailsOnly {
		var allDetails []DictionaryDetailInfo
		for _, dict := range dictionaries {
			allDetails = append(allDetails, dict.Details...)
		}
		
		response := map[string]interface{}{
			"success": true,
			"message": "查询字典详情成功",
			"total":   len(allDetails),
			"details": allDetails,
		}
		
		responseJSON, _ := json.Marshal(response)
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.NewTextContent(string(responseJSON)),
			},
		}, nil
	}
	
	// 构建响应
	response := DictionaryQueryResponse{
		Success:      true,
		Message:      "查询字典成功",
		Total:        len(dictionaries),
		Dictionaries: dictionaries,
	}
	
	responseJSON, err := json.Marshal(response)
	if err != nil {
		global.GVA_LOG.Error("序列化响应失败", zap.Error(err))
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.NewTextContent(fmt.Sprintf(`{"success": false, "message": "序列化响应失败: %v", "total": 0, "dictionaries": []}`, err.Error())),
			},
		}, nil
	}
	
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.NewTextContent(string(responseJSON)),
		},
	}, nil
}