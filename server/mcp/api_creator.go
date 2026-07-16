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
	RegisterTool(&ApiCreator{})
}

type ApiCreateRequest struct {
	Path        string `json:"path"`
	Description string `json:"description"`
	ApiGroup    string `json:"apiGroup"`
	Method      string `json:"method"`
}

type ApiCreateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	ApiID   uint   `json:"apiId"`
	Path    string `json:"path"`
	Method  string `json:"method"`
}

type ApiCreator struct{}

func (a *ApiCreator) New() mcp.Tool {
	return mcp.NewTool("create_api",
		mcp.WithDescription(`创建后端API记录，用于AI编辑器自动添加API接口时自动创建对应的API权限记录。

**两种调用模式（二选一）：**
- 单条：传 path + description + apiGroup（+可选 method，默认POST）
- 批量：传 apis（JSON数组字符串），一次创建多条；此模式下不需要 path/description/apiGroup

**重要限制：**
- 当使用gva_auto_generate工具且needCreatedModules=true时，模块创建会自动生成API权限，不应调用此工具
- 仅在以下情况使用：1) 单独创建API（不涉及模块创建）；2) AI编辑器自动添加API；3) router下的文件产生路径变化时`),
		mcp.WithString("path",
			mcp.Description("API路径，如：/user/create（单条模式必填；批量模式忽略）"),
		),
		mcp.WithString("description",
			mcp.Description("API中文描述，如：创建用户（单条模式必填；批量模式忽略）"),
		),
		mcp.WithString("apiGroup",
			mcp.Description("API组名称，用于分类管理，如：用户管理（单条模式必填；批量模式忽略）"),
		),
		mcp.WithString("method",
			mcp.Description("HTTP方法"),
			mcp.DefaultString("POST"),
		),
		mcp.WithString("apis",
			mcp.Description("批量创建API的JSON字符串，格式：[{\"path\":\"/user/create\",\"description\":\"创建用户\",\"apiGroup\":\"用户管理\",\"method\":\"POST\"}]"),
		),
	)
}

func (a *ApiCreator) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := request.GetArguments()

	var apis []ApiCreateRequest
	if apisStr, ok := args["apis"].(string); ok && apisStr != "" {
		if err := json.Unmarshal([]byte(apisStr), &apis); err != nil {
			return nil, fmt.Errorf("apis 参数格式错误: %w", err)
		}
	} else {
		path, ok := args["path"].(string)
		if !ok || path == "" {
			return nil, errors.New("path 参数是必需的")
		}
		description, ok := args["description"].(string)
		if !ok || description == "" {
			return nil, errors.New("description 参数是必需的")
		}
		apiGroup, ok := args["apiGroup"].(string)
		if !ok || apiGroup == "" {
			return nil, errors.New("apiGroup 参数是必需的")
		}

		method := "POST"
		if value, ok := args["method"].(string); ok && value != "" {
			method = value
		}

		apis = append(apis, ApiCreateRequest{
			Path:        path,
			Description: description,
			ApiGroup:    apiGroup,
			Method:      method,
		})
	}

	if len(apis) == 0 {
		return nil, errors.New("没有要创建的API")
	}

	responses := make([]ApiCreateResponse, 0, len(apis))
	successCount := 0

	for _, apiReq := range apis {
		// 批量模式给缺省 method 兜底 POST,与单条模式及 schema 默认值一致;
		// 否则空 method 会被上游 ApiVerify 判"Method 不能为空"导致该条静默失败
		if apiReq.Method == "" {
			apiReq.Method = "POST"
		}
		// createApi 现回传创建后的实体(含自增 ID),直接取用,免去此前逐条回查 getApiList 的 2N 次调用
		createResp, err := postUpstream[system.SysApi](ctx, "/api/createApi", system.SysApi{
			Path:        apiReq.Path,
			Description: apiReq.Description,
			ApiGroup:    apiReq.ApiGroup,
			Method:      apiReq.Method,
		})
		if err != nil {
			responses = append(responses, ApiCreateResponse{
				Success: false,
				Message: fmt.Sprintf("创建API失败: %v", err),
				Path:    apiReq.Path,
				Method:  apiReq.Method,
			})
			continue
		}

		responses = append(responses, ApiCreateResponse{
			Success: true,
			Message: fmt.Sprintf("成功创建API %s %s", apiReq.Method, apiReq.Path),
			ApiID:   createResp.Data.ID,
			Path:    apiReq.Path,
			Method:  apiReq.Method,
		})
		successCount++
	}

	result := map[string]any{
		"success":      successCount > 0,
		"totalCount":   len(apis),
		"successCount": successCount,
		"failedCount":  len(apis) - successCount,
		"details":      responses,
	}

	return textResultWithJSON("API创建结果：", result)
}
