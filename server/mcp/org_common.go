package mcpTool

import (
	"context"
	"errors"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// idName 组织实体的最小展示单元(部门/岗位)
type idName struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

const (
	// orgUserScanPageSize 上游 LimitOffset 的 MaxPageSize=100,取满
	orgUserScanPageSize = 100
	// orgUserScanMaxPages 按ID补全用户信息时的翻页上限:上游 getUserList 不支持按ID过滤,
	// 只能翻页扫描;超出上限的ID原样返回并标注未解析,避免大用户量下打爆主服务
	orgUserScanMaxPages = 20
	// orgBatchLimit 写类工具单次批量上限,控制误操作影响面
	orgBatchLimit = 50
)

// parseUintParam 解析单个正整数参数,兼容 MCP 客户端把数字发成字符串的情况
func parseUintParam(v any, name string) (uint, error) {
	switch value := v.(type) {
	case float64:
		if value <= 0 {
			return 0, fmt.Errorf("%s 必须大于0", name)
		}
		if value != math.Trunc(value) {
			return 0, fmt.Errorf("%s 必须为整数", name)
		}
		return uint(value), nil
	case string:
		id := strings.TrimSpace(value)
		if id == "" {
			return 0, fmt.Errorf("%s 参数是必需的", name)
		}
		parsed, err := strconv.ParseUint(id, 10, 64)
		if err != nil || parsed == 0 {
			return 0, fmt.Errorf("%s 格式不正确", name)
		}
		return uint(parsed), nil
	default:
		return 0, fmt.Errorf("%s 参数是必需的", name)
	}
}

// parseUintList 解析正整数ID列表,同时支持 JSON 数组与逗号分隔字符串两种格式
// (不同 MCP 客户端对数组参数支持不一),去重且保序;nil 返回空列表
func parseUintList(v any, name string) ([]uint, error) {
	if v == nil {
		return nil, nil
	}
	var raw []any
	switch value := v.(type) {
	case []any:
		raw = value
	case string:
		trimmed := strings.TrimSpace(value)
		if trimmed == "" {
			return nil, nil
		}
		for part := range strings.SplitSeq(trimmed, ",") {
			if strings.TrimSpace(part) == "" {
				continue // 容忍尾随/多余逗号(如 "1,2,"),不因空段整体报错
			}
			raw = append(raw, part)
		}
	default:
		return nil, fmt.Errorf("%s 需为ID数组或逗号分隔字符串", name)
	}

	seen := make(map[uint]struct{}, len(raw))
	ids := make([]uint, 0, len(raw))
	for _, item := range raw {
		id, err := parseUintParam(item, name)
		if err != nil {
			return nil, err
		}
		if _, ok := seen[id]; ok {
			continue
		}
		seen[id] = struct{}{}
		ids = append(ids, id)
	}
	return ids, nil
}

// mergeUintSets 把 add 合并进 current:返回合并后的全量集合(current 保序在前)与真正新增的部分
func mergeUintSets(current, add []uint) (merged []uint, added []uint) {
	seen := make(map[uint]struct{}, len(current)+len(add))
	merged = make([]uint, 0, len(current)+len(add))
	for _, id := range current {
		if _, ok := seen[id]; ok {
			continue
		}
		seen[id] = struct{}{}
		merged = append(merged, id)
	}
	for _, id := range add {
		if _, ok := seen[id]; ok {
			continue
		}
		seen[id] = struct{}{}
		merged = append(merged, id)
		added = append(added, id)
	}
	return merged, added
}

func containsUint(ids []uint, target uint) bool {
	return slices.Contains(ids, target)
}

// sameUintSet 忽略顺序与重复,判断两个ID集合是否等价
func sameUintSet(a, b []uint) bool {
	setA := make(map[uint]struct{}, len(a))
	for _, id := range a {
		setA[id] = struct{}{}
	}
	setB := make(map[uint]struct{}, len(b))
	for _, id := range b {
		setB[id] = struct{}{}
	}
	if len(setA) != len(setB) {
		return false
	}
	for id := range setA {
		if _, ok := setB[id]; !ok {
			return false
		}
	}
	return true
}

// flattenDeptTree 把部门树摊平成 id→名称 映射,供存在性校验与名称补全
func flattenDeptTree(list []system.SysDepartment, into map[uint]string) {
	for i := range list {
		into[list[i].ID] = list[i].Name
		if len(list[i].Children) > 0 {
			flattenDeptTree(list[i].Children, into)
		}
	}
}

// fetchDeptIndex 拉取部门树并摊平,一次上游调用完成全部部门ID校验
func fetchDeptIndex(ctx context.Context) (map[uint]string, error) {
	resp, err := postUpstream[[]system.SysDepartment](ctx, "/department/getDepartmentList", map[string]any{})
	if err != nil {
		return nil, fmt.Errorf("获取部门树失败: %w", err)
	}
	index := make(map[uint]string)
	flattenDeptTree(resp.Data, index)
	return index, nil
}

// findAuthorityInTree 在角色树中按 authorityId 查找角色
func findAuthorityInTree(list []system.SysAuthority, id uint) *system.SysAuthority {
	for i := range list {
		if list[i].AuthorityId == id {
			return &list[i]
		}
		if found := findAuthorityInTree(list[i].Children, id); found != nil {
			return found
		}
	}
	return nil
}

// fetchAuthority 拉取角色树并定位目标角色(getAuthorityList 返回调用者可见的角色树)
func fetchAuthority(ctx context.Context, authorityID uint) (*system.SysAuthority, error) {
	resp, err := postUpstream[[]system.SysAuthority](ctx, "/authority/getAuthorityList", map[string]any{})
	if err != nil {
		return nil, fmt.Errorf("获取角色列表失败: %w", err)
	}
	authority := findAuthorityInTree(resp.Data, authorityID)
	if authority == nil {
		return nil, fmt.Errorf("角色 %d 不存在或当前token无权查看", authorityID)
	}
	return authority, nil
}

// scanUsersByIDs 按ID集合补全用户信息:上游 getUserList 不支持按ID过滤,翻页扫描至集齐或达页数上限,
// 返回 命中映射 与 未解析ID 列表(行级可见范围由主服务侧 datascope 引擎决定)
func scanUsersByIDs(ctx context.Context, ids []uint) (map[uint]system.SysUser, []uint, error) {
	want := make(map[uint]struct{}, len(ids))
	for _, id := range ids {
		want[id] = struct{}{}
	}
	found := make(map[uint]system.SysUser, len(ids))

	for page := 1; page <= orgUserScanMaxPages && len(found) < len(want); page++ {
		resp, err := postUpstream[pageResultData[[]system.SysUser]](ctx, "/user/getUserList", map[string]any{
			"page":     page,
			"pageSize": orgUserScanPageSize,
		})
		if err != nil {
			return nil, nil, fmt.Errorf("获取用户列表失败: %w", err)
		}
		for i := range resp.Data.List {
			user := resp.Data.List[i]
			if _, ok := want[user.ID]; ok {
				found[user.ID] = user
			}
		}
		if len(resp.Data.List) < orgUserScanPageSize {
			break
		}
	}

	var unresolved []uint
	for _, id := range ids {
		if _, ok := found[id]; !ok {
			unresolved = append(unresolved, id)
		}
	}
	return found, unresolved, nil
}

// findUserByID 按ID定位单个用户;usernameHint 可选,提供时先按用户名过滤精确命中,避免全量翻页
func findUserByID(ctx context.Context, userID uint, usernameHint string) (*system.SysUser, error) {
	if hint := strings.TrimSpace(usernameHint); hint != "" {
		resp, err := postUpstream[pageResultData[[]system.SysUser]](ctx, "/user/getUserList", map[string]any{
			"page":     1,
			"pageSize": orgUserScanPageSize,
			"username": hint,
		})
		if err != nil {
			return nil, fmt.Errorf("获取用户列表失败: %w", err)
		}
		for i := range resp.Data.List {
			if resp.Data.List[i].ID == userID {
				return &resp.Data.List[i], nil
			}
		}
		// hint 仅是加速命中的优化;username 为 LIKE 模糊匹配且只查首页,未命中(名不匹配或
		// 命中过多致目标不在首页)时不误报,回退全量翻页扫描兜底
	}

	found, _, err := scanUsersByIDs(ctx, []uint{userID})
	if err != nil {
		return nil, err
	}
	if user, ok := found[userID]; ok {
		return &user, nil
	}
	return nil, fmt.Errorf("未找到 ID 为 %d 的用户(已扫描前 %d 条,大用户量场景请传 username 参数精确定位)", userID, orgUserScanMaxPages*orgUserScanPageSize)
}

// requireNonEmptyList 语义化的必填列表校验
func requireNonEmptyList(ids []uint, name string) error {
	if len(ids) == 0 {
		return errors.New(name + " 参数是必需的,且至少包含一个ID")
	}
	return nil
}

// parseOptionalPositiveInt 解析可选的正整数分页参数,兼容数字与字符串两种形式
// (与本工具集"数字参数兼容字符串"的既定设计一致);缺省或非法时返回 def
func parseOptionalPositiveInt(v any, def int) int {
	switch value := v.(type) {
	case float64:
		if value >= 1 {
			return int(value)
		}
	case string:
		if n, err := strconv.Atoi(strings.TrimSpace(value)); err == nil && n >= 1 {
			return n
		}
	}
	return def
}

// parseOptionalUint 解析可选的非负整数参数(允许0,如菜单 parentId=0 表示根),兼容 float64 与字符串;
// 缺省或非法返回 def。用于那些"0 是合法值"故不能走拒绝0的 parseUintParam 的场景
func parseOptionalUint(v any, def uint) uint {
	switch value := v.(type) {
	case float64:
		if value >= 0 && value == math.Trunc(value) {
			return uint(value)
		}
	case string:
		if n, err := strconv.ParseUint(strings.TrimSpace(value), 10, 64); err == nil {
			return uint(n)
		}
	}
	return def
}
