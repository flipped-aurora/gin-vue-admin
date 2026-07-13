package mcpTool

import (
	"context"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/mark3labs/mcp-go/mcp"
)

func init() {
	RegisterTool(&RoleMenuAssigner{})
}

type RoleMenuAssigner struct{}

type roleMenuAssignResponse struct {
	Success         bool     `json:"success"`
	Message         string   `json:"message"`
	AuthorityID     uint     `json:"authorityId"`
	Added           []idName `json:"added"`
	AlreadyExists   []idName `json:"alreadyExists"`
	ParentAutoAdded []idName `json:"parentAutoAdded"`
	TotalMenus      int      `json:"totalMenus"`
}

// menuParentChainMaxDepth 父链回溯深度上限,防脏数据成环
const menuParentChainMaxDepth = 32

func (r *RoleMenuAssigner) New() mcp.Tool {
	return mcp.NewTool("assign_menu_to_role",
		mcp.WithDescription(`将菜单授权给角色（仅追加，绝不清除角色已有菜单授权）。

**功能说明：**
- 上游接口是全量覆盖语义；本工具先读取角色当前菜单授权，与传入 menuIds 合并去重后全量写回，防止丢权
- 父级菜单自动补齐：授权子菜单时自动带上其完整父链，避免出现"有子无父"导致前端渲染不出
- 重复调用幂等：已授权的菜单返回在 alreadyExists 中，不重复写入
- 菜单ID可先用 list_all_menus 工具查询

**适用场景：**
- create_menu 创建菜单后授权给指定角色
- 给角色补齐缺失的菜单可见性
- 注意：本工具不做"取消授权"，如需回收请到前端角色管理页操作`),
		mcp.WithNumber("authorityId",
			mcp.Required(),
			mcp.Description("角色ID，如：888"),
		),
		mcp.WithString("menuIds",
			mcp.Required(),
			mcp.Description("要授权的菜单ID列表，JSON数组或逗号分隔字符串，如 \"12,15\"，单次上限50个"),
		),
	)
}

func (r *RoleMenuAssigner) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := request.GetArguments()

	authorityID, err := parseAuthorityID(args["authorityId"])
	if err != nil {
		return nil, err
	}
	menuIDs, err := parseUintList(args["menuIds"], "menuIds")
	if err != nil {
		return nil, err
	}
	if err := requireNonEmptyList(menuIDs, "menuIds"); err != nil {
		return nil, err
	}
	if len(menuIDs) > orgBatchLimit {
		return nil, fmt.Errorf("单次授权菜单数量不能超过 %d 个", orgBatchLimit)
	}

	// 全量基础菜单:存在性校验、父链补齐、写回时的完整菜单对象都从这里取
	allResp, err := postUpstream[[]system.SysBaseMenu](ctx, "/menu/getMenuList", map[string]any{})
	if err != nil {
		return nil, fmt.Errorf("获取菜单列表失败: %w", err)
	}
	menuIndex := make(map[uint]system.SysBaseMenu)
	flattenBaseMenus(allResp.Data, menuIndex)

	for _, id := range menuIDs {
		if _, ok := menuIndex[id]; !ok {
			return nil, fmt.Errorf("菜单 %d 不存在,可先用 list_all_menus 查询全部菜单", id)
		}
	}

	// 角色当前授权(先读,防覆盖丢权)
	currentResp, err := postUpstream[map[string][]system.SysMenu](ctx, "/menu/getMenuAuthority", map[string]any{
		"authorityId": authorityID,
	})
	if err != nil {
		return nil, fmt.Errorf("获取角色当前菜单授权失败: %w", err)
	}
	currentIDs := make([]uint, 0)
	for _, menu := range currentResp.Data["menus"] {
		id := menu.MenuId
		if id == 0 {
			id = menu.ID
		}
		if id != 0 && !containsUint(currentIDs, id) {
			currentIDs = append(currentIDs, id)
		}
	}

	// 请求集 = 传入菜单 ∪ 各自的完整父链
	requested := make([]uint, 0, len(menuIDs))
	parentAdded := make([]uint, 0)
	for _, id := range menuIDs {
		requested = append(requested, id)
		parentID := menuIndex[id].ParentId
		for depth := 0; parentID != 0 && depth < menuParentChainMaxDepth; depth++ {
			parent, ok := menuIndex[parentID]
			if !ok {
				break
			}
			if !containsUint(requested, parent.ID) && !containsUint(currentIDs, parent.ID) {
				requested = append(requested, parent.ID)
				parentAdded = append(parentAdded, parent.ID)
			}
			parentID = parent.ParentId
		}
	}

	merged, added := mergeUintSets(currentIDs, requested)

	result := roleMenuAssignResponse{
		Success:     true,
		AuthorityID: authorityID,
		TotalMenus:  len(merged),
	}
	nameOf := func(id uint) idName {
		if menu, ok := menuIndex[id]; ok {
			return idName{ID: id, Name: menu.Title}
		}
		return idName{ID: id, Name: "(菜单已不在基础菜单表中,原样保留授权)"}
	}
	for _, id := range menuIDs {
		if containsUint(added, id) {
			result.Added = append(result.Added, nameOf(id))
		} else {
			result.AlreadyExists = append(result.AlreadyExists, nameOf(id))
		}
	}
	for _, id := range parentAdded {
		if containsUint(added, id) {
			result.ParentAutoAdded = append(result.ParentAutoAdded, nameOf(id))
		}
	}

	if len(added) == 0 {
		result.Message = "全部菜单均已授权,无需重复写入"
		return textResultWithJSON("角色菜单授权结果：", result)
	}

	// 全量写回:当前授权中已被删除的基础菜单构造最小对象原样保留,不因本次授权而丢失
	payload := make([]system.SysBaseMenu, 0, len(merged))
	for _, id := range merged {
		if menu, ok := menuIndex[id]; ok {
			payload = append(payload, menu)
			continue
		}
		var placeholder system.SysBaseMenu
		placeholder.ID = id
		payload = append(payload, placeholder)
	}
	if _, err := postUpstream[map[string]any](ctx, "/menu/addMenuAuthority", map[string]any{
		"authorityId": authorityID,
		"menus":       payload,
	}); err != nil {
		return nil, fmt.Errorf("菜单授权写回失败: %w", err)
	}

	result.Message = fmt.Sprintf("成功为角色 %d 新增 %d 个菜单授权(含自动补齐父级 %d 个),当前共 %d 个",
		authorityID, len(added), len(result.ParentAutoAdded), len(merged))
	return textResultWithJSON("角色菜单授权结果：", result)
}

// flattenBaseMenus 把基础菜单树摊平成 id→菜单 映射
func flattenBaseMenus(list []system.SysBaseMenu, into map[uint]system.SysBaseMenu) {
	for i := range list {
		into[list[i].ID] = list[i]
		if len(list[i].Children) > 0 {
			flattenBaseMenus(list[i].Children, into)
		}
	}
}
