package mcpTool

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/mark3labs/mcp-go/mcp"
)

func init() {
	RegisterTool(&OrgMemberAssigner{})
}

type OrgMemberAssigner struct{}

type orgAssignChange struct {
	Before []uint `json:"before"`
	After  []uint `json:"after"`
	Added  []uint `json:"added"`
}

type orgMemberAssignResponse struct {
	Success       bool             `json:"success"`
	Message       string           `json:"message"`
	UserID        uint             `json:"userId"`
	UserName      string           `json:"userName"`
	Departments   *orgAssignChange `json:"departments,omitempty"`
	Positions     *orgAssignChange `json:"positions,omitempty"`
	PrimaryDeptID uint             `json:"primaryDeptId,omitempty"`
	AlreadyExists bool             `json:"alreadyExists"`
}

func (o *OrgMemberAssigner) New() mcp.Tool {
	return mcp.NewTool("assign_user_org",
		mcp.WithDescription(`给指定用户追加分配部门/岗位归属（仅追加，绝不清除用户已有归属）。

**功能说明：**
- 上游接口是全量覆盖语义；本工具内部先读取用户当前部门/岗位集合，合并去重后全量写回，实现追加幂等
- 重复调用安全：已有归属不会重复写入，也不会丢失
- 主部门(数据归属/盖章用)默认保留用户现有主部门；仅当显式传 primaryDeptId 时才变更
- deptIds 会先经部门树校验存在性；positionIds 逐个经岗位查询校验（单次各上限50个）
- 返回 before/after/added 对比明细

**适用场景：**
- 新人入职：把用户加入部门与岗位
- 组织调整：给用户追加兼管部门或兼任岗位
- 注意：本工具不做"移除归属"，如需移除请到前端组织管理页操作`),
		mcp.WithNumber("userId",
			mcp.Required(),
			mcp.Description("目标用户ID"),
		),
		mcp.WithString("username",
			mcp.Description("目标用户登录名，可选。大用户量场景传入可精确定位，避免翻页扫描"),
		),
		mcp.WithString("deptIds",
			mcp.Description("要追加的部门ID列表，JSON数组或逗号分隔字符串，如 \"3,5\"；与 positionIds 至少传一个"),
		),
		mcp.WithString("positionIds",
			mcp.Description("要追加的岗位ID列表，JSON数组或逗号分隔字符串；与 deptIds 至少传一个"),
		),
		mcp.WithNumber("primaryDeptId",
			mcp.Description("显式指定主部门ID(数据归属/盖章)，可选；必须属于合并后的部门集合。不传则保留用户现有主部门"),
		),
	)
}

func (o *OrgMemberAssigner) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := request.GetArguments()

	userID, err := parseUintParam(args["userId"], "userId")
	if err != nil {
		return nil, err
	}
	usernameHint, _ := args["username"].(string)

	deptIDs, err := parseUintList(args["deptIds"], "deptIds")
	if err != nil {
		return nil, err
	}
	positionIDs, err := parseUintList(args["positionIds"], "positionIds")
	if err != nil {
		return nil, err
	}
	if len(deptIDs) == 0 && len(positionIDs) == 0 {
		return nil, errors.New("deptIds 与 positionIds 至少需要传一个")
	}
	if len(deptIDs) > orgBatchLimit || len(positionIDs) > orgBatchLimit {
		return nil, fmt.Errorf("单次分配的部门/岗位数量各不能超过 %d 个", orgBatchLimit)
	}
	// primaryDeptId 只在设置部门(deptIds)时生效;仅传 primaryDeptId 不传 deptIds 会被静默丢弃,
	// 显式挡回避免调用方误以为主部门已变更
	if value, ok := args["primaryDeptId"]; ok && value != nil && len(deptIDs) == 0 {
		return nil, errors.New("primaryDeptId 需与 deptIds 一起传入(主部门必须属于本次追加或已有的部门集合);仅调整主部门请到前端组织管理页操作")
	}

	// 校验目标部门/岗位真实存在,把无效ID挡在写入之前
	if len(deptIDs) > 0 {
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
	for _, id := range positionIDs {
		query := url.Values{}
		query.Set("id", strconv.FormatUint(uint64(id), 10))
		if _, err := getUpstream[system.SysPosition](ctx, "/position/findPosition", query); err != nil {
			return nil, fmt.Errorf("岗位 %d 校验失败: %w", id, err)
		}
	}

	user, err := findUserByID(ctx, userID, usernameHint)
	if err != nil {
		return nil, err
	}

	currentDeptIDs := make([]uint, 0, len(user.Departments))
	for _, dept := range user.Departments {
		currentDeptIDs = append(currentDeptIDs, dept.ID)
	}
	currentPositionIDs := make([]uint, 0, len(user.Positions))
	for _, pos := range user.Positions {
		currentPositionIDs = append(currentPositionIDs, pos.ID)
	}

	result := orgMemberAssignResponse{
		Success:  true,
		UserID:   user.ID,
		UserName: user.Username,
	}
	var messages []string

	if len(deptIDs) > 0 {
		merged, added := mergeUintSets(currentDeptIDs, deptIDs)

		// 主部门语义:显式传参优先,否则保留用户现有主部门;上游对 0 值取集合首个
		primary := user.DeptId
		if value, ok := args["primaryDeptId"]; ok && value != nil {
			explicit, err := parseUintParam(value, "primaryDeptId")
			if err != nil {
				return nil, err
			}
			primary = explicit
		}
		if primary != 0 && !containsUint(merged, primary) {
			return nil, fmt.Errorf("primaryDeptId %d 必须属于合并后的部门集合 %v", primary, merged)
		}

		if len(added) > 0 || (primary != user.DeptId && primary != 0) {
			if _, err := postUpstream[map[string]any](ctx, "/user/setUserDepartments", map[string]any{
				"ID":            user.ID,
				"deptIds":       merged,
				"primaryDeptId": primary,
			}); err != nil {
				return nil, fmt.Errorf("设置用户部门失败: %w", err)
			}
			messages = append(messages, fmt.Sprintf("部门新增 %v", added))
		} else {
			messages = append(messages, "部门均已存在,未写入")
		}
		result.Departments = &orgAssignChange{Before: currentDeptIDs, After: merged, Added: added}
		// 上游对 primaryDeptId=0 取部门集合首个;如实回报生效主部门,不因 0 值被 omitempty 隐去而误导
		result.PrimaryDeptID = primary
		if primary == 0 && len(merged) > 0 {
			result.PrimaryDeptID = merged[0]
		}
	}

	if len(positionIDs) > 0 {
		merged, added := mergeUintSets(currentPositionIDs, positionIDs)
		if len(added) > 0 {
			if _, err := postUpstream[map[string]any](ctx, "/user/setUserPositions", map[string]any{
				"ID":          user.ID,
				"positionIds": merged,
			}); err != nil {
				return nil, fmt.Errorf("设置用户岗位失败: %w", err)
			}
			messages = append(messages, fmt.Sprintf("岗位新增 %v", added))
		} else {
			messages = append(messages, "岗位均已存在,未写入")
		}
		result.Positions = &orgAssignChange{Before: currentPositionIDs, After: merged, Added: added}
	}

	result.AlreadyExists = (result.Departments == nil || len(result.Departments.Added) == 0) &&
		(result.Positions == nil || len(result.Positions.Added) == 0)
	result.Message = fmt.Sprintf("用户 %s(ID:%d) 组织归属分配完成：%s", user.Username, user.ID, strings.Join(messages, "；"))

	return textResultWithJSON("用户组织归属分配结果：", result)
}
