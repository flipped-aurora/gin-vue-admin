package mcpTool

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/mark3labs/mcp-go/mcp"
)

func init() {
	RegisterTool(&RoleDataScopeSetter{})
}

type RoleDataScopeSetter struct{}

type dataScopeState struct {
	DataScope int    `json:"dataScope"`
	Label     string `json:"label"`
	DeptIDs   []uint `json:"deptIds,omitempty"`
}

type roleDataScopeResponse struct {
	Success       bool           `json:"success"`
	Message       string         `json:"message"`
	AuthorityID   uint           `json:"authorityId"`
	AuthorityName string         `json:"authorityName"`
	Before        dataScopeState `json:"before"`
	After         dataScopeState `json:"after"`
	Changed       bool           `json:"changed"`
}

var dataScopeLabels = map[int]string{
	1: "全部数据",
	2: "本部门及以下",
	3: "本部门",
	4: "仅本人",
	5: "自定义部门",
}

func (r *RoleDataScopeSetter) New() mcp.Tool {
	return mcp.NewTool("set_role_data_scope",
		mcp.WithDescription(`设置角色的数据权限档位（行级数据可见范围，由数据权限引擎在查询时自动过滤）。

**五档位语义（dataScope）：**
- 1 全部数据：不做行级过滤，可见所有数据
- 2 本部门及以下：可见本部门与全部子级部门的数据
- 3 本部门：仅可见本部门数据（不含子级）
- 4 仅本人：仅可见自己创建的数据
- 5 自定义部门：可见 deptIds 指定部门集合的数据（此档位必须传非空 deptIds，工具会先校验部门ID真实存在）

**功能说明：**
- 写入前读取当前配置，返回变更前后对比
- 档位2/3的"本部门"判定基于用户的**全部归属部门集合**(多部门，含主部门)，不是仅主部门；
  主部门(primaryDeptId)仅决定新建数据时盖的 dept_id 归属。接口权限走 casbin，与本档位正交

**适用场景：**
- 新建角色后配置其数据可见范围
- 排查"用户看不到某条数据"时调整角色档位`),
		mcp.WithNumber("authorityId",
			mcp.Required(),
			mcp.Description("角色ID，如：888"),
		),
		mcp.WithNumber("dataScope",
			mcp.Required(),
			mcp.Description("数据权限档位：1全部 2本部门及以下 3本部门 4仅本人 5自定义部门"),
		),
		mcp.WithString("deptIds",
			mcp.Description("自定义部门ID集合，JSON数组或逗号分隔字符串；仅档位5必传，其它档位忽略"),
		),
	)
}

func (r *RoleDataScopeSetter) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := request.GetArguments()

	authorityID, err := parseAuthorityID(args["authorityId"])
	if err != nil {
		return nil, err
	}

	dataScope, err := parseUintParam(args["dataScope"], "dataScope")
	if err != nil {
		return nil, err
	}
	if dataScope < 1 || dataScope > 5 {
		return nil, fmt.Errorf("dataScope 必须在 1-5 之间：1全部 2本部门及以下 3本部门 4仅本人 5自定义部门，收到 %d", dataScope)
	}
	scope := int(dataScope)

	var deptIDs []uint
	if scope == 5 {
		deptIDs, err = parseUintList(args["deptIds"], "deptIds")
		if err != nil {
			return nil, err
		}
		if err := requireNonEmptyList(deptIDs, "deptIds"); err != nil {
			return nil, fmt.Errorf("档位5(自定义部门)必须指定部门集合: %w", err)
		}
		deptIndex, err := fetchDeptIndex(ctx)
		if err != nil {
			return nil, err
		}
		for _, id := range deptIDs {
			if _, ok := deptIndex[id]; !ok {
				return nil, fmt.Errorf("部门 %d 不存在,可先用 query_org_structure 查询部门树", id)
			}
		}
	}

	// 读取变更前状态:角色档位来自角色树,自定义部门集来自专用接口
	authority, err := fetchAuthority(ctx, authorityID)
	if err != nil {
		return nil, err
	}
	query := url.Values{}
	query.Set("authorityId", strconv.FormatUint(uint64(authorityID), 10))
	beforeDeptsResp, err := getUpstream[[]uint](ctx, "/authority/getDataScopeDepts", query)
	if err != nil {
		return nil, fmt.Errorf("读取角色当前自定义部门集失败: %w", err)
	}

	before := dataScopeState{
		DataScope: authority.DataScope,
		Label:     dataScopeLabels[authority.DataScope],
		DeptIDs:   beforeDeptsResp.Data,
	}

	if _, err := postUpstream[map[string]any](ctx, "/authority/setDataScope", map[string]any{
		"authorityId": authorityID,
		"dataScope":   scope,
		"deptIds":     deptIDs,
	}); err != nil {
		return nil, fmt.Errorf("设置数据权限失败: %w", err)
	}

	after := dataScopeState{
		DataScope: scope,
		Label:     dataScopeLabels[scope],
		DeptIDs:   deptIDs,
	}

	return textResultWithJSON("角色数据权限设置结果：", roleDataScopeResponse{
		Success:       true,
		Message:       fmt.Sprintf("角色 %s(ID:%d) 数据权限已设置为「%s」", authority.AuthorityName, authorityID, after.Label),
		AuthorityID:   authorityID,
		AuthorityName: authority.AuthorityName,
		Before:        before,
		After:         after,
		Changed:       before.DataScope != after.DataScope || !sameUintSet(before.DeptIDs, after.DeptIDs),
	})
}
