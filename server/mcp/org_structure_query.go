package mcpTool

import (
	"context"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/mark3labs/mcp-go/mcp"
)

func init() {
	RegisterTool(&OrgStructureQuery{})
}

type OrgStructureQuery struct{}

type orgStructureResponse struct {
	Success        bool                   `json:"success"`
	Message        string                 `json:"message"`
	DepartmentTree []system.SysDepartment `json:"departmentTree,omitempty"`
	DepartmentFlat bool                   `json:"departmentFlat,omitempty"` // name 过滤时上游返回扁平匹配列表(无子级树),true 提示调用方勿把有子部门的节点当叶子
	Positions      []system.SysPosition   `json:"positions,omitempty"`
	PositionTotal  int64                  `json:"positionTotal,omitempty"`
	Page           int                    `json:"page,omitempty"`
	PageSize       int                    `json:"pageSize,omitempty"`
}

func (o *OrgStructureQuery) New() mcp.Tool {
	return mcp.NewTool("query_org_structure",
		mcp.WithDescription(`查询组织架构：部门树与岗位列表。

**功能说明：**
- scope=department 返回完整部门树（含祖级链 ancestors、负责人 leaderId、排序、启用状态）
- scope=position 分页返回岗位列表（岗位与角色正交，仅作组织身份/职务）
- scope=all 同时返回两者
- name 参数对部门/岗位均为名称模糊过滤；**部门带 name 过滤时上游返回扁平匹配列表(departmentFlat=true)，不组装子级树**，此时勿把有子部门的节点当叶子

**适用场景：**
- 分配用户部门/岗位前查询可用的部门ID与岗位ID
- 设置角色自定义数据权限(档位5)前确认部门ID真实存在
- 了解企业组织结构与汇报关系（部门 leaderId 为负责人用户ID）`),
		mcp.WithString("scope",
			mcp.Description("查询范围：department(部门树) | position(岗位) | all(两者)，默认 all"),
			mcp.DefaultString("all"),
		),
		mcp.WithString("name",
			mcp.Description("名称模糊过滤，可选"),
		),
		mcp.WithNumber("page",
			mcp.Description("岗位分页页码，默认1"),
		),
		mcp.WithNumber("pageSize",
			mcp.Description("岗位分页大小，默认100（上游上限100）"),
		),
	)
}

func (o *OrgStructureQuery) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := request.GetArguments()

	scope := "all"
	if value, ok := args["scope"].(string); ok && strings.TrimSpace(value) != "" {
		scope = strings.ToLower(strings.TrimSpace(value))
	}
	name := ""
	if value, ok := args["name"].(string); ok {
		name = strings.TrimSpace(value)
	}

	result := orgStructureResponse{Success: true, Message: "查询成功"}
	if scope != "department" && scope != "position" && scope != "all" {
		scope = "all"
		result.Message = "scope 仅支持 department|position|all，已按 all 处理"
	}

	if scope == "department" || scope == "all" {
		deptResp, err := postUpstream[[]system.SysDepartment](ctx, "/department/getDepartmentList", map[string]any{
			"name": name,
		})
		if err != nil {
			return nil, err
		}
		result.DepartmentTree = deptResp.Data
		if name != "" {
			// 上游 GetSysDepartmentTree 带 name 时平铺返回匹配项(Children 为空),如实标注避免误导
			result.DepartmentFlat = true
			result.Message = "查询成功(部门为 name 过滤后的扁平匹配列表,非完整树)"
		}
	}

	if scope == "position" || scope == "all" {
		page := parseOptionalPositiveInt(args["page"], 1)
		pageSize := parseOptionalPositiveInt(args["pageSize"], orgUserScanPageSize)
		posResp, err := postUpstream[pageResultData[[]system.SysPosition]](ctx, "/position/getPositionList", map[string]any{
			"page":     page,
			"pageSize": pageSize,
			"name":     name,
		})
		if err != nil {
			return nil, err
		}
		result.Positions = posResp.Data.List
		result.PositionTotal = posResp.Data.Total
		result.Page = posResp.Data.Page
		result.PageSize = posResp.Data.PageSize
	}

	return textResultWithJSON("组织架构查询结果：", result)
}
