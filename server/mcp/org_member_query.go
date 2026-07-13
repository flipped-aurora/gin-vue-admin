package mcpTool

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
)

func init() {
	RegisterTool(&OrgMemberQuery{})
}

type OrgMemberQuery struct{}

type orgMemberInfo struct {
	UserID        uint     `json:"userId"`
	UserName      string   `json:"userName"`
	NickName      string   `json:"nickName"`
	PrimaryDeptID uint     `json:"primaryDeptId"`
	Departments   []idName `json:"departments"`
	Positions     []idName `json:"positions"`
}

type orgMemberResponse struct {
	Success       bool            `json:"success"`
	Message       string          `json:"message"`
	Dimension     string          `json:"dimension"`
	TargetID      uint            `json:"targetId"`
	MemberCount   int             `json:"memberCount"`
	Members       []orgMemberInfo `json:"members,omitempty"`
	MemberIDs     []uint          `json:"memberIds"`
	UnresolvedIDs []uint          `json:"unresolvedIds,omitempty"`
}

func (o *OrgMemberQuery) New() mcp.Tool {
	return mcp.NewTool("query_org_members",
		mcp.WithDescription(`按维度查询成员用户：部门成员、岗位成员或角色关联用户。

**功能说明：**
- dimension=department：查询部门直属成员（不含子部门）
- dimension=position：查询岗位成员
- dimension=authority：查询拥有指定角色的用户
- 默认补全用户信息（用户名/昵称/主部门/多部门/多岗位）；上游用户列表不支持按ID过滤，
  补全通过翻页扫描实现，超出扫描上限的ID会列在 unresolvedIds 中原样返回
- 组织/用户为系统表(sys_*)，数据权限引擎不对其做行级过滤；返回的是主服务记录的完整成员归属，
  读取本身仍需调用者 token 具备对应接口的 casbin 授权

**适用场景：**
- 分配/调整用户组织归属前查看现状
- 审批人排查：确认某部门/岗位/角色下有哪些人`),
		mcp.WithString("dimension",
			mcp.Required(),
			mcp.Description("查询维度：department(部门) | position(岗位) | authority(角色)"),
		),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("对应维度的ID：部门ID / 岗位ID / 角色ID"),
		),
		mcp.WithBoolean("withUserInfo",
			mcp.Description("是否补全用户信息，默认true；false 时仅返回用户ID列表（免翻页扫描，速度快）"),
			mcp.DefaultBool(true),
		),
	)
}

func (o *OrgMemberQuery) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := request.GetArguments()

	dimension, _ := args["dimension"].(string)
	dimension = strings.ToLower(strings.TrimSpace(dimension))

	targetID, err := parseUintParam(args["id"], "id")
	if err != nil {
		return nil, err
	}

	withUserInfo := true
	if value, ok := args["withUserInfo"].(bool); ok {
		withUserInfo = value
	}

	var endpoint, idParam string
	switch dimension {
	case "department":
		endpoint, idParam = "/department/getDepartmentUsers", "deptId"
	case "position":
		endpoint, idParam = "/position/getPositionUsers", "positionId"
	case "authority":
		endpoint, idParam = "/authority/getUsersByAuthority", "authorityId"
	default:
		return nil, errors.New("dimension 仅支持 department|position|authority")
	}

	query := url.Values{}
	query.Set(idParam, strconv.FormatUint(uint64(targetID), 10))
	idsResp, err := getUpstream[[]uint](ctx, endpoint, query)
	if err != nil {
		return nil, fmt.Errorf("获取成员ID列表失败: %w", err)
	}
	memberIDs := idsResp.Data

	result := orgMemberResponse{
		Success:     true,
		Message:     fmt.Sprintf("共 %d 名成员", len(memberIDs)),
		Dimension:   dimension,
		TargetID:    targetID,
		MemberCount: len(memberIDs),
		MemberIDs:   memberIDs,
	}

	if withUserInfo && len(memberIDs) > 0 {
		found, unresolved, err := scanUsersByIDs(ctx, memberIDs)
		if err != nil {
			return nil, err
		}
		for _, id := range memberIDs {
			user, ok := found[id]
			if !ok {
				continue
			}
			info := orgMemberInfo{
				UserID:        user.ID,
				UserName:      user.Username,
				NickName:      user.NickName,
				PrimaryDeptID: user.DeptId,
			}
			for _, dept := range user.Departments {
				info.Departments = append(info.Departments, idName{ID: dept.ID, Name: dept.Name})
			}
			for _, pos := range user.Positions {
				info.Positions = append(info.Positions, idName{ID: pos.ID, Name: pos.Name})
			}
			result.Members = append(result.Members, info)
		}
		result.UnresolvedIDs = unresolved
		if len(unresolved) > 0 {
			result.Message += fmt.Sprintf("，其中 %d 名超出翻页扫描上限(前%d条)未补全信息，ID见 unresolvedIds", len(unresolved), orgUserScanMaxPages*orgUserScanPageSize)
		}
	}

	return textResultWithJSON("组织成员查询结果：", result)
}
