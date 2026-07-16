package mcpTool

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/stretchr/testify/require"
)

// ---------- 纯函数 ----------

func TestParseUintList(t *testing.T) {
	cases := []struct {
		name    string
		input   any
		want    []uint
		wantErr bool
	}{
		{name: "JSON数组", input: []any{float64(1), float64(2)}, want: []uint{1, 2}},
		{name: "数组内字符串", input: []any{"3", float64(4)}, want: []uint{3, 4}},
		{name: "逗号分隔字符串", input: "5, 6,7", want: []uint{5, 6, 7}},
		{name: "去重保序", input: "8,9,8", want: []uint{8, 9}},
		{name: "nil返回空", input: nil, want: nil},
		{name: "空字符串返回空", input: "  ", want: nil},
		{name: "非法元素", input: "a,b", wantErr: true},
		{name: "零值拒绝", input: []any{float64(0)}, wantErr: true},
		{name: "非法类型", input: map[string]any{}, wantErr: true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := parseUintList(tc.input, "ids")
			if tc.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.want, got)
		})
	}
}

func TestMergeUintSets(t *testing.T) {
	merged, added := mergeUintSets([]uint{1, 2}, []uint{2, 3, 4})
	require.Equal(t, []uint{1, 2, 3, 4}, merged)
	require.Equal(t, []uint{3, 4}, added)

	merged, added = mergeUintSets([]uint{1, 2}, []uint{1, 2})
	require.Equal(t, []uint{1, 2}, merged)
	require.Empty(t, added)
}

func TestSameUintSet(t *testing.T) {
	require.True(t, sameUintSet([]uint{1, 2, 2}, []uint{2, 1}))
	require.False(t, sameUintSet([]uint{1}, []uint{1, 2}))
	require.False(t, sameUintSet([]uint{1, 3}, []uint{1, 2}))
	require.True(t, sameUintSet(nil, nil))
}

func TestFlattenDeptTree(t *testing.T) {
	tree := []system.SysDepartment{{
		Name: "总部",
		Children: []system.SysDepartment{
			{Name: "研发部"},
		},
	}}
	tree[0].ID = 1
	tree[0].Children[0].ID = 2

	index := make(map[uint]string)
	flattenDeptTree(tree, index)
	require.Equal(t, map[uint]string{1: "总部", 2: "研发部"}, index)
}

func TestParseAPIItems(t *testing.T) {
	items, err := parseAPIItems([]any{
		map[string]any{"path": "/user/getUserList", "method": "POST"},
		map[string]any{"path": "/menu/getMenuList"},
	})
	require.NoError(t, err)
	require.Len(t, items, 2)
	require.Equal(t, "/user/getUserList", items[0].Path)
	require.Equal(t, "", items[1].Method) // 缺省交给 normalizePolicy 补POST

	items, err = parseAPIItems("POST /a/b, GET /c/d, /e/f")
	require.NoError(t, err)
	require.Len(t, items, 3)
	require.Equal(t, "GET", items[1].Method)
	require.Equal(t, "/e/f", items[2].Path)

	// 批内去重(方法缺省归一化后与显式POST等价)
	items, err = parseAPIItems("POST /a/b, /a/b")
	require.NoError(t, err)
	require.Len(t, items, 1)

	_, err = parseAPIItems([]any{map[string]any{"method": "POST"}})
	require.Error(t, err)

	_, err = parseAPIItems("GET POST /a/b")
	require.Error(t, err)
}

// ---------- Handle 路径(httptest 伪造上游) ----------

type upstreamMock struct {
	mu       sync.Mutex
	server   *httptest.Server
	handlers map[string]http.HandlerFunc
	bodies   map[string][]byte // 按端点记录最近一次请求体
	methods  map[string]string // 按端点记录最近一次 HTTP 方法
}

func newUpstreamMock(t *testing.T) *upstreamMock {
	t.Helper()
	m := &upstreamMock{
		handlers: make(map[string]http.HandlerFunc),
		bodies:   make(map[string][]byte),
		methods:  make(map[string]string),
	}
	m.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		m.mu.Lock()
		m.bodies[r.URL.Path] = body
		m.methods[r.URL.Path] = r.Method
		handler, ok := m.handlers[r.URL.Path]
		m.mu.Unlock()
		if !ok {
			t.Errorf("未预期的上游调用: %s %s", r.Method, r.URL.Path)
			http.NotFound(w, r)
			return
		}
		handler(w, r)
	}))
	t.Cleanup(m.server.Close)

	original := global.GVA_CONFIG.MCP
	global.GVA_CONFIG.MCP.UpstreamBaseURL = m.server.URL
	t.Cleanup(func() { global.GVA_CONFIG.MCP = original })
	return m
}

func (m *upstreamMock) on(path string, data any) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.handlers[path] = func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		payload, _ := json.Marshal(map[string]any{"code": 0, "data": data, "msg": "成功"})
		_, _ = w.Write(payload)
	}
}

// onError 让端点返回业务错误码(code!=0),用于覆盖上游报错透传路径
func (m *upstreamMock) onError(path string, code int, msg string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.handlers[path] = func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		payload, _ := json.Marshal(map[string]any{"code": code, "data": nil, "msg": msg})
		_, _ = w.Write(payload)
	}
}

func (m *upstreamMock) methodOf(path string) string {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.methods[path]
}

func (m *upstreamMock) requestBody(t *testing.T, path string, into any) {
	t.Helper()
	m.mu.Lock()
	body, ok := m.bodies[path]
	m.mu.Unlock()
	require.True(t, ok, "端点 %s 未被调用", path)
	require.NoError(t, json.Unmarshal(body, into))
}

func (m *upstreamMock) called(path string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	_, ok := m.bodies[path]
	return ok
}

func authedCtx() context.Context {
	return context.WithValue(context.Background(), authTokenContextKey, "test-token")
}

func callTool(t *testing.T, tool McpTool, ctx context.Context, args map[string]any, into any) {
	t.Helper()
	request := mcp.CallToolRequest{}
	request.Params.Arguments = args
	result, err := tool.Handle(ctx, request)
	require.NoError(t, err)
	require.NotEmpty(t, result.Content)
	text := result.Content[0].(mcp.TextContent).Text
	parts := strings.SplitN(text, "\n\n", 2)
	payload := parts[len(parts)-1]
	require.NoError(t, json.Unmarshal([]byte(payload), into))
}

func TestToolsRejectMissingToken(t *testing.T) {
	newUpstreamMock(t)
	request := mcp.CallToolRequest{}
	request.Params.Arguments = map[string]any{
		"authorityId": float64(888),
		"apis":        "POST /a/b",
	}
	_, err := (&RoleAPIBatchAssigner{}).Handle(context.Background(), request)
	require.ErrorContains(t, err, "缺少MCP鉴权请求头")
}

func TestBatchAssignAPIsHandle(t *testing.T) {
	m := newUpstreamMock(t)
	m.on("/casbin/getPolicyPathByAuthorityId", map[string]any{
		"paths": []systemReq.CasbinInfo{{Path: "/a/b", Method: "POST"}},
	})
	m.on("/casbin/updateCasbin", map[string]any{})

	var resp roleAPIBatchAssignResponse
	callTool(t, &RoleAPIBatchAssigner{}, authedCtx(), map[string]any{
		"authorityId": float64(888),
		"apis":        "POST /a/b, GET /c/d",
	}, &resp)

	require.True(t, resp.Success)
	require.Equal(t, 1, resp.AddedCount)
	require.Equal(t, 1, resp.SkippedCount)
	require.Equal(t, 2, resp.TotalPolicies)

	var written struct {
		AuthorityId uint                   `json:"authorityId"`
		CasbinInfos []systemReq.CasbinInfo `json:"casbinInfos"`
	}
	m.requestBody(t, "/casbin/updateCasbin", &written)
	require.Equal(t, uint(888), written.AuthorityId)
	require.Len(t, written.CasbinInfos, 2, "写回必须包含既有策略,不得丢权")
}

func TestBatchAssignAPIsIdempotentSkipsWrite(t *testing.T) {
	m := newUpstreamMock(t)
	m.on("/casbin/getPolicyPathByAuthorityId", map[string]any{
		"paths": []systemReq.CasbinInfo{{Path: "/a/b", Method: "POST"}},
	})

	var resp roleAPIBatchAssignResponse
	callTool(t, &RoleAPIBatchAssigner{}, authedCtx(), map[string]any{
		"authorityId": float64(888),
		"apis":        "/a/b",
	}, &resp)

	require.Equal(t, 0, resp.AddedCount)
	require.False(t, m.called("/casbin/updateCasbin"), "全部已存在时不得写回")
}

func TestRoleMenuAssignerHandle(t *testing.T) {
	parent := system.SysBaseMenu{}
	parent.ID = 1
	parent.Title = "系统管理"
	child := system.SysBaseMenu{ParentId: 1}
	child.ID = 2
	child.Title = "菜单管理"
	parent.Children = []system.SysBaseMenu{child}

	existing := system.SysMenu{MenuId: 3}
	existing.Title = "已有菜单"

	m := newUpstreamMock(t)
	m.on("/menu/getMenuList", []system.SysBaseMenu{parent})
	m.on("/menu/getMenuAuthority", map[string]any{"menus": []system.SysMenu{existing}})
	m.on("/menu/addMenuAuthority", map[string]any{})

	var resp roleMenuAssignResponse
	callTool(t, &RoleMenuAssigner{}, authedCtx(), map[string]any{
		"authorityId": float64(888),
		"menuIds":     "2",
	}, &resp)

	require.True(t, resp.Success)
	require.Len(t, resp.Added, 1)
	require.Len(t, resp.ParentAutoAdded, 1, "授权子菜单必须自动补齐父级")
	require.Equal(t, 3, resp.TotalMenus)

	var written struct {
		AuthorityId uint                 `json:"authorityId"`
		Menus       []system.SysBaseMenu `json:"menus"`
	}
	m.requestBody(t, "/menu/addMenuAuthority", &written)
	ids := make([]uint, 0, len(written.Menus))
	for _, menu := range written.Menus {
		ids = append(ids, menu.ID)
	}
	require.ElementsMatch(t, []uint{1, 2, 3}, ids, "写回必须包含既有授权3、目标2与自动补齐的父级1")
}

func TestRoleMenuAssignerIdempotent(t *testing.T) {
	menu := system.SysBaseMenu{}
	menu.ID = 2
	existing := system.SysMenu{MenuId: 2}

	m := newUpstreamMock(t)
	m.on("/menu/getMenuList", []system.SysBaseMenu{menu})
	m.on("/menu/getMenuAuthority", map[string]any{"menus": []system.SysMenu{existing}})

	var resp roleMenuAssignResponse
	callTool(t, &RoleMenuAssigner{}, authedCtx(), map[string]any{
		"authorityId": float64(888),
		"menuIds":     "2",
	}, &resp)

	require.Len(t, resp.AlreadyExists, 1)
	require.False(t, m.called("/menu/addMenuAuthority"), "全部已授权时不得写回")
}

func TestSetRoleDataScopeHandle(t *testing.T) {
	dept := system.SysDepartment{Name: "研发部"}
	dept.ID = 3

	m := newUpstreamMock(t)
	m.on("/authority/getAuthorityList", []system.SysAuthority{{AuthorityId: 888, AuthorityName: "超管", DataScope: 1}})
	m.on("/authority/getDataScopeDepts", []uint{})
	m.on("/department/getDepartmentList", []system.SysDepartment{dept})
	m.on("/authority/setDataScope", map[string]any{})

	var resp roleDataScopeResponse
	callTool(t, &RoleDataScopeSetter{}, authedCtx(), map[string]any{
		"authorityId": float64(888),
		"dataScope":   float64(5),
		"deptIds":     "3",
	}, &resp)

	require.True(t, resp.Success)
	require.True(t, resp.Changed)
	require.Equal(t, 1, resp.Before.DataScope)
	require.Equal(t, 5, resp.After.DataScope)
	require.Equal(t, []uint{3}, resp.After.DeptIDs)

	var written systemReq.SetDataScope
	m.requestBody(t, "/authority/setDataScope", &written)
	require.Equal(t, uint(888), written.AuthorityId)
	require.Equal(t, 5, written.DataScope)
	require.Equal(t, []uint{3}, written.DeptIds)
}

func TestSetRoleDataScopeValidation(t *testing.T) {
	newUpstreamMock(t)

	request := mcp.CallToolRequest{}
	request.Params.Arguments = map[string]any{"authorityId": float64(888), "dataScope": float64(7)}
	_, err := (&RoleDataScopeSetter{}).Handle(authedCtx(), request)
	require.ErrorContains(t, err, "1-5")

	request.Params.Arguments = map[string]any{"authorityId": float64(888), "dataScope": float64(5)}
	_, err = (&RoleDataScopeSetter{}).Handle(authedCtx(), request)
	require.ErrorContains(t, err, "档位5")
}

func TestAssignUserOrgHandle(t *testing.T) {
	deptA := system.SysDepartment{Name: "总部"}
	deptA.ID = 2
	deptB := system.SysDepartment{Name: "研发部"}
	deptB.ID = 3

	currentDept := system.SysDepartment{}
	currentDept.ID = 2
	currentPos := system.SysPosition{}
	currentPos.ID = 7

	user := system.SysUser{
		Username:    "demo",
		DeptId:      2,
		Departments: []system.SysDepartment{currentDept},
		Positions:   []system.SysPosition{currentPos},
	}
	user.ID = 10

	m := newUpstreamMock(t)
	m.on("/department/getDepartmentList", []system.SysDepartment{deptA, deptB})
	m.on("/user/getUserList", pageResultData[[]system.SysUser]{List: []system.SysUser{user}, Total: 1, Page: 1, PageSize: 100})
	m.on("/position/findPosition", currentPos)
	m.on("/user/setUserDepartments", map[string]any{})

	var resp orgMemberAssignResponse
	callTool(t, &OrgMemberAssigner{}, authedCtx(), map[string]any{
		"userId":      float64(10),
		"deptIds":     "3",
		"positionIds": "7",
	}, &resp)

	require.True(t, resp.Success)
	require.Equal(t, []uint{3}, resp.Departments.Added)
	require.Empty(t, resp.Positions.Added, "岗位7已存在,必须幂等跳过")
	require.False(t, m.called("/user/setUserPositions"), "岗位无新增时不得写回")

	var written systemReq.SetUserDepartments
	m.requestBody(t, "/user/setUserDepartments", &written)
	require.Equal(t, uint(10), written.ID)
	require.ElementsMatch(t, []uint{2, 3}, written.DeptIds, "写回必须保留既有部门2")
	require.Equal(t, uint(2), written.PrimaryDeptId, "未显式指定时必须保留现有主部门")
}

func TestOrgMemberQueryHandle(t *testing.T) {
	member := system.SysUser{Username: "demo", NickName: "演示", DeptId: 2}
	member.ID = 10

	m := newUpstreamMock(t)
	m.on("/department/getDepartmentUsers", []uint{10, 99})
	m.on("/user/getUserList", pageResultData[[]system.SysUser]{List: []system.SysUser{member}, Total: 1, Page: 1, PageSize: 100})

	var resp orgMemberResponse
	callTool(t, &OrgMemberQuery{}, authedCtx(), map[string]any{
		"dimension": "department",
		"id":        float64(2),
	}, &resp)

	require.Equal(t, 2, resp.MemberCount)
	require.Len(t, resp.Members, 1)
	require.Equal(t, "demo", resp.Members[0].UserName)
	require.Equal(t, []uint{99}, resp.UnresolvedIDs, "扫描不到的ID必须显式标注")
}

func TestOrgStructureQueryHandle(t *testing.T) {
	dept := system.SysDepartment{Name: "总部"}
	dept.ID = 1
	pos := system.SysPosition{Name: "工程师"}
	pos.ID = 7

	m := newUpstreamMock(t)
	m.on("/department/getDepartmentList", []system.SysDepartment{dept})
	m.on("/position/getPositionList", pageResultData[[]system.SysPosition]{List: []system.SysPosition{pos}, Total: 1, Page: 1, PageSize: 100})

	var resp orgStructureResponse
	callTool(t, &OrgStructureQuery{}, authedCtx(), map[string]any{"scope": "all"}, &resp)

	require.True(t, resp.Success)
	require.Len(t, resp.DepartmentTree, 1)
	require.False(t, resp.DepartmentFlat, "无 name 过滤时是完整树,不应标记 flat")
	require.Len(t, resp.Positions, 1)
	require.Equal(t, int64(1), resp.PositionTotal)
}

// name 过滤时上游平铺返回,工具须标记 departmentFlat;且分页参数兼容字符串形式
func TestOrgStructureQueryNameFilterFlatAndStringPaging(t *testing.T) {
	dept := system.SysDepartment{Name: "研发部"}
	dept.ID = 5

	m := newUpstreamMock(t)
	m.on("/department/getDepartmentList", []system.SysDepartment{dept})
	m.on("/position/getPositionList", pageResultData[[]system.SysPosition]{List: nil, Total: 0, Page: 2, PageSize: 20})

	var resp orgStructureResponse
	callTool(t, &OrgStructureQuery{}, authedCtx(), map[string]any{
		"scope": "all", "name": "研发",
		"page": "2", "pageSize": "20", // 字符串形式,验证 parseOptionalPositiveInt 容错
	}, &resp)

	require.True(t, resp.DepartmentFlat, "带 name 过滤必须标记为扁平列表")

	var posReq struct {
		Page     int `json:"page"`
		PageSize int `json:"pageSize"`
	}
	m.requestBody(t, "/position/getPositionList", &posReq)
	require.Equal(t, 2, posReq.Page, "字符串 page 应被解析为 2")
	require.Equal(t, 20, posReq.PageSize, "字符串 pageSize 应被解析为 20")
}

// ---------- 补充: 纯函数边界 ----------

func TestParseUintListTolerantAndIntOnly(t *testing.T) {
	// 尾随/多余逗号应被容忍,而非整体报"参数是必需的"
	got, err := parseUintList("1,2,", "ids")
	require.NoError(t, err)
	require.Equal(t, []uint{1, 2}, got)

	got, err = parseUintList("1, ,3", "ids")
	require.NoError(t, err)
	require.Equal(t, []uint{1, 3}, got)

	// 非整数 float 必须拒绝,避免权限写操作落到被截断的错误ID上
	_, err = parseUintParam(float64(3.5), "id")
	require.ErrorContains(t, err, "整数")
}

func TestParseAPIItemsJSONArrayAsString(t *testing.T) {
	// 只支持字符串参数的客户端会把 JSON 数组序列化成字符串传入,不能被逗号切成垃圾策略
	items, err := parseAPIItems(`[{"path":"/user/getUserList","method":"POST"},{"path":"/menu/getMenuList"}]`)
	require.NoError(t, err)
	require.Len(t, items, 2)
	require.Equal(t, "/user/getUserList", items[0].Path)
	require.Equal(t, "POST", items[0].Method)
	require.Equal(t, "/menu/getMenuList", items[1].Path)
}

// ---------- 补充: 上游报错透传 ----------

func TestUpstreamErrorPassthrough(t *testing.T) {
	m := newUpstreamMock(t)
	m.onError("/department/getDepartmentUsers", 7, "无权访问该部门")

	request := mcp.CallToolRequest{}
	request.Params.Arguments = map[string]any{"dimension": "department", "id": float64(2), "withUserInfo": false}
	_, err := (&OrgMemberQuery{}).Handle(authedCtx(), request)
	require.Error(t, err)
	require.ErrorContains(t, err, "无权访问该部门", "上游业务错误必须原样透传给调用方")
}

// ---------- 补充: query_org_members 三维度与非法维度 ----------

func TestOrgMemberQueryDimensions(t *testing.T) {
	t.Run("position", func(t *testing.T) {
		m := newUpstreamMock(t)
		m.on("/position/getPositionUsers", []uint{11, 12})
		var resp orgMemberResponse
		callTool(t, &OrgMemberQuery{}, authedCtx(), map[string]any{
			"dimension": "position", "id": float64(7), "withUserInfo": false,
		}, &resp)
		require.Equal(t, "position", resp.Dimension)
		require.Equal(t, []uint{11, 12}, resp.MemberIDs)
		require.Equal(t, http.MethodGet, m.methodOf("/position/getPositionUsers"), "成员ID端点应为 GET")
	})

	t.Run("authority", func(t *testing.T) {
		m := newUpstreamMock(t)
		m.on("/authority/getUsersByAuthority", []uint{21})
		var resp orgMemberResponse
		callTool(t, &OrgMemberQuery{}, authedCtx(), map[string]any{
			"dimension": "authority", "id": float64(888), "withUserInfo": false,
		}, &resp)
		require.Equal(t, []uint{21}, resp.MemberIDs)
		require.Equal(t, http.MethodGet, m.methodOf("/authority/getUsersByAuthority"))
	})

	t.Run("非法维度", func(t *testing.T) {
		newUpstreamMock(t)
		request := mcp.CallToolRequest{}
		request.Params.Arguments = map[string]any{"dimension": "foo", "id": float64(1)}
		_, err := (&OrgMemberQuery{}).Handle(authedCtx(), request)
		require.ErrorContains(t, err, "department|position|authority")
	})
}

// ---------- 补充: 写类工具的校验拒绝分支 ----------

func TestAssignUserOrgRejections(t *testing.T) {
	t.Run("部门不存在", func(t *testing.T) {
		dept := system.SysDepartment{Name: "总部"}
		dept.ID = 2
		m := newUpstreamMock(t)
		m.on("/department/getDepartmentList", []system.SysDepartment{dept})
		request := mcp.CallToolRequest{}
		request.Params.Arguments = map[string]any{"userId": float64(10), "deptIds": "99"}
		_, err := (&OrgMemberAssigner{}).Handle(authedCtx(), request)
		require.ErrorContains(t, err, "部门 99 不存在")
	})

	t.Run("仅传primaryDeptId不传deptIds", func(t *testing.T) {
		newUpstreamMock(t)
		request := mcp.CallToolRequest{}
		request.Params.Arguments = map[string]any{"userId": float64(10), "positionIds": "7", "primaryDeptId": float64(3)}
		_, err := (&OrgMemberAssigner{}).Handle(authedCtx(), request)
		require.ErrorContains(t, err, "primaryDeptId 需与 deptIds", "只传主部门不传部门必须显式挡回,不能静默丢弃")
	})

	t.Run("primaryDeptId不在合并集", func(t *testing.T) {
		deptA := system.SysDepartment{Name: "总部"}
		deptA.ID = 2
		deptB := system.SysDepartment{Name: "研发部"}
		deptB.ID = 3
		currentDept := system.SysDepartment{}
		currentDept.ID = 2
		user := system.SysUser{Username: "demo", DeptId: 2, Departments: []system.SysDepartment{currentDept}}
		user.ID = 10

		m := newUpstreamMock(t)
		m.on("/department/getDepartmentList", []system.SysDepartment{deptA, deptB})
		m.on("/user/getUserList", pageResultData[[]system.SysUser]{List: []system.SysUser{user}, Total: 1, Page: 1, PageSize: 100})
		// 合并集为 {2(既有),3(追加)};主部门指定不在集合内的 999 应被拒绝
		request := mcp.CallToolRequest{}
		request.Params.Arguments = map[string]any{"userId": float64(10), "deptIds": "3", "primaryDeptId": float64(999)}
		_, err := (&OrgMemberAssigner{}).Handle(authedCtx(), request)
		require.ErrorContains(t, err, "必须属于合并后的部门集合")
	})
}

func TestRoleMenuAssignerRejectsUnknownMenu(t *testing.T) {
	menu := system.SysBaseMenu{}
	menu.ID = 1
	m := newUpstreamMock(t)
	m.on("/menu/getMenuList", []system.SysBaseMenu{menu})
	request := mcp.CallToolRequest{}
	request.Params.Arguments = map[string]any{"authorityId": float64(888), "menuIds": "999"}
	_, err := (&RoleMenuAssigner{}).Handle(authedCtx(), request)
	require.ErrorContains(t, err, "菜单 999 不存在")
}

func TestBatchAssignAPIsOverLimit(t *testing.T) {
	newUpstreamMock(t)
	parts := make([]string, 0, orgBatchLimit+1)
	for i := 0; i <= orgBatchLimit; i++ {
		parts = append(parts, fmt.Sprintf("/p/%d", i))
	}
	request := mcp.CallToolRequest{}
	request.Params.Arguments = map[string]any{"authorityId": float64(888), "apis": strings.Join(parts, ",")}
	_, err := (&RoleAPIBatchAssigner{}).Handle(authedCtx(), request)
	require.ErrorContains(t, err, "不能超过")
}
