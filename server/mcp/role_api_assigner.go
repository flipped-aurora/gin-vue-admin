package mcpTool

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/mark3labs/mcp-go/mcp"
)

func init() {
	RegisterTool(&RoleAPIAssigner{})
}

type RoleAPIAssigner struct{}

type roleAPIAssignResponse struct {
	Success       bool   `json:"success"`
	Message       string `json:"message"`
	AuthorityID   uint   `json:"authorityId"`
	Path          string `json:"path"`
	Method        string `json:"method"`
	Added         bool   `json:"added"`
	AlreadyExists bool   `json:"alreadyExists"`
	TotalPolicies int    `json:"totalPolicies"`
}

func (a *RoleAPIAssigner) New() mcp.Tool {
	return mcp.NewTool("assign_api_to_role",
		mcp.WithDescription(`将指定API权限追加分配给角色（仅追加，不覆盖原有权限）。

**功能说明：**
- 自动读取角色当前API权限
- 若目标权限不存在则追加并保存
- 若已存在则幂等返回成功，不重复写入

**适用场景：**
- create_api后将新API授权给指定角色
- 手动补齐某个角色缺失的API权限`),
		mcp.WithNumber("authorityId",
			mcp.Required(),
			mcp.Description("角色ID，如：888"),
		),
		mcp.WithString("path",
			mcp.Required(),
			mcp.Description("API路径，如：/user/getCustomerUserList"),
		),
		mcp.WithString("method",
			mcp.Description("HTTP方法，默认POST"),
			mcp.DefaultString("POST"),
		),
	)
}

func (a *RoleAPIAssigner) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := request.GetArguments()

	authorityID, err := parseAuthorityID(args["authorityId"])
	if err != nil {
		return nil, err
	}

	path, ok := args["path"].(string)
	if !ok || strings.TrimSpace(path) == "" {
		return nil, errors.New("path 参数是必需的")
	}
	method := "POST"
	if value, ok := args["method"].(string); ok && strings.TrimSpace(value) != "" {
		method = value
	}

	path, method = normalizePolicy(path, method)
	currentResp, err := postUpstream[map[string][]systemReq.CasbinInfo](ctx, "/casbin/getPolicyPathByAuthorityId", map[string]any{
		"authorityId": authorityID,
	})
	if err != nil {
		return nil, fmt.Errorf("获取角色当前API权限失败: %w", err)
	}

	current := currentResp.Data["paths"]
	updated, added := appendPolicyIfMissing(current, path, method)
	if added {
		if _, err = postUpstream[map[string]any](ctx, "/casbin/updateCasbin", map[string]any{
			"authorityId": authorityID,
			"casbinInfos": updated,
		}); err != nil {
			return nil, fmt.Errorf("分配API权限失败: %w", err)
		}
	}

	msg := "权限已存在，无需重复分配"
	if added {
		msg = fmt.Sprintf("成功为角色 %d 分配权限 %s %s", authorityID, method, path)
	}

	return textResultWithJSON("角色API权限分配结果：", roleAPIAssignResponse{
		Success:       true,
		Message:       msg,
		AuthorityID:   authorityID,
		Path:          path,
		Method:        method,
		Added:         added,
		AlreadyExists: !added,
		TotalPolicies: len(updated),
	})
}

func parseAuthorityID(v any) (uint, error) {
	switch value := v.(type) {
	case float64:
		if value <= 0 {
			return 0, errors.New("authorityId 必须大于0")
		}
		return uint(value), nil
	case string:
		id := strings.TrimSpace(value)
		if id == "" {
			return 0, errors.New("authorityId 参数是必需的")
		}
		parsed, err := strconv.ParseUint(id, 10, 64)
		if err != nil || parsed == 0 {
			return 0, errors.New("authorityId 格式不正确")
		}
		return uint(parsed), nil
	default:
		return 0, errors.New("authorityId 参数是必需的")
	}
}

func appendPolicyIfMissing(current []systemReq.CasbinInfo, path, method string) ([]systemReq.CasbinInfo, bool) {
	path, method = normalizePolicy(path, method)
	for _, policy := range current {
		p, m := normalizePolicy(policy.Path, policy.Method)
		if p == path && m == method {
			return current, false
		}
	}
	updated := make([]systemReq.CasbinInfo, 0, len(current)+1)
	updated = append(updated, current...)
	updated = append(updated, systemReq.CasbinInfo{
		Path:   path,
		Method: method,
	})
	return updated, true
}

func normalizePolicy(path, method string) (string, string) {
	path = strings.TrimSpace(path)
	if path == "" {
		path = "/"
	}
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	method = strings.ToUpper(strings.TrimSpace(method))
	if method == "" {
		method = "POST"
	}
	return path, method
}
