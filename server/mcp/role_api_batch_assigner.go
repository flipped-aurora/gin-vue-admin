package mcpTool

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/mark3labs/mcp-go/mcp"
)

func init() {
	RegisterTool(&RoleAPIBatchAssigner{})
}

type RoleAPIBatchAssigner struct{}

type batchAPIItemResult struct {
	Path          string `json:"path"`
	Method        string `json:"method"`
	Added         bool   `json:"added"`
	AlreadyExists bool   `json:"alreadyExists"`
}

type roleAPIBatchAssignResponse struct {
	Success       bool                 `json:"success"`
	Message       string               `json:"message"`
	AuthorityID   uint                 `json:"authorityId"`
	Items         []batchAPIItemResult `json:"items"`
	AddedCount    int                  `json:"addedCount"`
	SkippedCount  int                  `json:"skippedCount"`
	TotalPolicies int                  `json:"totalPolicies"`
}

func (r *RoleAPIBatchAssigner) New() mcp.Tool {
	return mcp.NewTool("batch_assign_apis_to_role",
		mcp.WithDescription(`将一批API权限追加分配给角色（仅追加，不覆盖原有权限）。单条分配请用 assign_api_to_role。

**功能说明：**
- 一次读取角色现有策略，逐条合并去重，仅一次写回——读-合并-写在单次调用内完成，缩小并发覆盖窗口
- 已存在的策略幂等跳过，返回逐条 added/alreadyExists 明细
- 单次上限50条

**apis 参数格式（两种任选）：**
1. JSON数组：[{"path":"/user/getUserList","method":"POST"},{"path":"/menu/getMenuList"}]（method 缺省为POST）
2. 逗号分隔字符串："POST /user/getUserList, GET /department/getDepartmentUsers"（省略方法时默认POST）

**适用场景：**
- 新建角色后批量开通一组API
- 代码生成器产出新模块后，把整组增删改查API授权给角色`),
		mcp.WithNumber("authorityId",
			mcp.Required(),
			mcp.Description("角色ID，如：888"),
		),
		mcp.WithString("apis",
			mcp.Required(),
			mcp.Description("API列表，JSON数组[{path,method}]或逗号分隔字符串\"METHOD /path, ...\""),
		),
	)
}

func (r *RoleAPIBatchAssigner) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := request.GetArguments()

	authorityID, err := parseAuthorityID(args["authorityId"])
	if err != nil {
		return nil, err
	}
	items, err := parseAPIItems(args["apis"])
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, errors.New("apis 参数是必需的,且至少包含一条")
	}
	if len(items) > orgBatchLimit {
		return nil, fmt.Errorf("单次批量授权不能超过 %d 条,收到 %d 条", orgBatchLimit, len(items))
	}

	currentResp, err := postUpstream[map[string][]systemReq.CasbinInfo](ctx, "/casbin/getPolicyPathByAuthorityId", map[string]any{
		"authorityId": authorityID,
	})
	if err != nil {
		return nil, fmt.Errorf("获取角色当前API权限失败: %w", err)
	}

	updated := currentResp.Data["paths"]
	result := roleAPIBatchAssignResponse{Success: true, AuthorityID: authorityID}
	anyAdded := false
	for _, item := range items {
		var added bool
		updated, added = appendPolicyIfMissing(updated, item.Path, item.Method)
		path, method := normalizePolicy(item.Path, item.Method)
		result.Items = append(result.Items, batchAPIItemResult{
			Path:          path,
			Method:        method,
			Added:         added,
			AlreadyExists: !added,
		})
		if added {
			result.AddedCount++
			anyAdded = true
		} else {
			result.SkippedCount++
		}
	}

	if anyAdded {
		if _, err := postUpstream[map[string]any](ctx, "/casbin/updateCasbin", map[string]any{
			"authorityId": authorityID,
			"casbinInfos": updated,
		}); err != nil {
			return nil, fmt.Errorf("批量分配API权限失败: %w", err)
		}
	}

	result.TotalPolicies = len(updated)
	result.Message = fmt.Sprintf("角色 %d 批量授权完成：新增 %d 条,已存在跳过 %d 条,当前共 %d 条策略",
		authorityID, result.AddedCount, result.SkippedCount, result.TotalPolicies)
	return textResultWithJSON("角色API权限批量分配结果：", result)
}

// parseAPIItems 解析 apis 参数,支持 JSON 数组[{path,method}]与"METHOD /path"逗号分隔字符串双格式,
// 批内按 归一化path+method 去重
func parseAPIItems(v any) ([]systemReq.CasbinInfo, error) {
	if v == nil {
		return nil, nil
	}
	var items []systemReq.CasbinInfo

	appendItem := func(path, method string) {
		items = append(items, systemReq.CasbinInfo{Path: path, Method: method})
	}

	switch value := v.(type) {
	case []any:
		for _, entry := range value {
			switch typed := entry.(type) {
			case map[string]any:
				path, _ := typed["path"].(string)
				if strings.TrimSpace(path) == "" {
					return nil, errors.New("apis 数组元素缺少 path 字段")
				}
				method, _ := typed["method"].(string)
				appendItem(path, method)
			case string:
				path, method, err := parseAPIItemString(typed)
				if err != nil {
					return nil, err
				}
				appendItem(path, method)
			default:
				return nil, errors.New("apis 数组元素需为 {path,method} 对象或 \"METHOD /path\" 字符串")
			}
		}
	case string:
		trimmed := strings.TrimSpace(value)
		if trimmed == "" {
			return nil, nil
		}
		// 兼容只支持字符串参数的 MCP 客户端把 JSON 数组序列化成字符串传入的情况:
		// 形如 [{"path":..,"method":..}] 若按逗号切分会碎成垃圾策略,先尝试按 JSON 数组解析
		if strings.HasPrefix(trimmed, "[") {
			var arr []any
			if err := json.Unmarshal([]byte(trimmed), &arr); err == nil {
				return parseAPIItems(arr)
			}
		}
		for part := range strings.SplitSeq(trimmed, ",") {
			if strings.TrimSpace(part) == "" {
				continue
			}
			path, method, err := parseAPIItemString(part)
			if err != nil {
				return nil, err
			}
			appendItem(path, method)
		}
	default:
		return nil, errors.New("apis 需为JSON数组或逗号分隔字符串")
	}

	// 批内去重(按归一化后的 path+method)
	seen := make(map[string]struct{}, len(items))
	deduped := items[:0]
	for _, item := range items {
		path, method := normalizePolicy(item.Path, item.Method)
		key := method + " " + path
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		deduped = append(deduped, item)
	}
	return deduped, nil
}

// parseAPIItemString 解析"METHOD /path"或"/path"形式的单条API,方法缺省为POST
func parseAPIItemString(s string) (path, method string, err error) {
	fields := strings.Fields(strings.TrimSpace(s))
	switch len(fields) {
	case 1:
		return fields[0], "", nil
	case 2:
		return fields[1], fields[0], nil
	default:
		return "", "", fmt.Errorf("无法解析API项 %q,期望 \"/path\" 或 \"METHOD /path\"", s)
	}
}
