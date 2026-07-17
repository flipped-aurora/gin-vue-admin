package service

import (
	"archive/zip"
	"bytes"
	"io"
	"strings"
	"testing"

	autoModel "github.com/flipped-aurora/gin-vue-admin/server/plugin/ai/model"
	autoRes "github.com/flipped-aurora/gin-vue-admin/server/plugin/ai/model/response"
)

func TestBuildSkillRenderDataFillsDefaults(t *testing.T) {
	cli := autoModel.SysCli{
		Name:        "user-manager",
		Command:     "opsctl",
		DisplayName: "运维工具",
		Description: "用于运维查询",
	}
	manifest := autoRes.SysCliManifestResponse{
		Name:    "opsctl",
		Version: "v1",
		Commands: []autoRes.SysCliManifestCommand{
			{Name: "user-list", Summary: "分页获取用户列表", Method: "POST", Path: "/user/list", Examples: []string{"opsctl user-list --page 1"}, Parameters: []autoRes.SysCliManifestParameter{
				{Flag: "page", Type: "integer", Required: true, Description: "页码"},
				{Flag: "page-size", Type: "integer", Description: "每页大小"},
			}},
		},
	}
	data := buildSkillRenderData(cli, manifest)
	if data.SkillName != "opsctl-cli" {
		t.Fatalf("SkillName default = %q, want %q", data.SkillName, "opsctl-cli")
	}
	if !strings.Contains(data.SkillDescription, "opsctl") || !strings.Contains(data.SkillDescription, "运维工具") {
		t.Fatalf("SkillDescription = %q, want contains command and displayName", data.SkillDescription)
	}
	if len(data.Commands) != 1 || data.Commands[0].CommandName != "user-list" {
		t.Fatalf("commands not mapped: %+v", data.Commands)
	}
	if len(data.Commands[0].Parameters) != 2 {
		t.Fatalf("parameters not mapped: %+v", data.Commands[0].Parameters)
	}
}

func TestRenderSkillBodyContainsCommandsAndParams(t *testing.T) {
	data := skillRenderData{
		Command:          "opsctl",
		DisplayName:      "运维工具",
		Description:      "用于运维查询",
		SkillName:        "opsctl-cli",
		SkillDescription: "通过 opsctl 调用运维工具的 API。",
		Commands: []skillCommand{
			{
				CommandName: "user-list",
				Summary:     "分页获取用户列表",
				Parameters: []skillParam{
					{Flag: "page", Type: "integer", Required: true, Description: "页码"},
					{Flag: "page-size", Type: "integer", Description: "每页大小"},
				},
				Example: "opsctl user-list --page 1",
			},
		},
	}
	body := renderSkillBody(data)
	for _, want := range []string{"opsctl", "运维工具", "user-list", "--page", "page-size", "opsctl user-list --page 1", "必填"} {
		if !strings.Contains(body, want) {
			t.Fatalf("renderSkillBody missing %q\n%s", want, body)
		}
	}
}

func TestRenderSkillBodyContainsResponseSection(t *testing.T) {
	data := skillRenderData{
		Command:     "opsctl",
		DisplayName: "运维工具",
		SkillName:   "opsctl-cli",
		Commands: []skillCommand{
			{
				CommandName: "user-list",
				Summary:     "分页获取用户列表",
				Response: []skillResponseField{
					{Name: "total", Description: "总数"},
					{Name: "list", Description: "用户列表"},
				},
			},
		},
	}
	body := renderSkillBody(data)
	for _, want := range []string{"返回：", "total", "list", "总数"} {
		if !strings.Contains(body, want) {
			t.Fatalf("renderSkillBody missing %q\n%s", want, body)
		}
	}
}

func TestRenderSkillBodyOmitsResponseSectionWhenEmpty(t *testing.T) {
	data := skillRenderData{
		Command:     "opsctl",
		DisplayName: "运维工具",
		SkillName:   "opsctl-cli",
		Commands: []skillCommand{
			{CommandName: "user-list", Summary: "分页获取用户列表"},
		},
	}
	body := renderSkillBody(data)
	if strings.Contains(body, "返回：") {
		t.Fatalf("should not render response section when empty\n%s", body)
	}
}

func TestWriteSkillPackageZipContainsAllFiles(t *testing.T) {
	zipBytes, err := writeSkillPackageZip("opsctl-cli", "# skill", "readme body", []byte(`{"name":"cpt"}`), "cpt.exe", []byte("BINARY"))
	if err != nil {
		t.Fatalf("writeSkillPackageZip: %v", err)
	}
	r, err := zip.NewReader(bytes.NewReader(zipBytes), int64(len(zipBytes)))
	if err != nil {
		t.Fatalf("zip.NewReader: %v", err)
	}
	want := map[string]bool{
		"opsctl-cli/SKILL.md":                 false,
		"opsctl-cli/references/README.md":     false,
		"opsctl-cli/references/manifest.json": false,
		"opsctl-cli/cpt.exe":                  false,
	}
	for _, f := range r.File {
		if _, ok := want[f.Name]; ok {
			want[f.Name] = true
		}
	}
	for name, found := range want {
		if !found {
			t.Errorf("missing %q in zip", name)
		}
	}
	for _, f := range r.File {
		if f.Name != "opsctl-cli/cpt.exe" {
			continue
		}
		rc, err := f.Open()
		if err != nil {
			t.Fatalf("open cpt.exe: %v", err)
		}
		data, _ := io.ReadAll(rc)
		rc.Close()
		if string(data) != "BINARY" {
			t.Errorf("cpt.exe content = %q, want %q", string(data), "BINARY")
		}
	}
}

func TestParseCliScenarios(t *testing.T) {
	if got, err := parseCliScenarios(""); err != nil || len(got) != 0 {
		t.Fatalf("empty: got=%v err=%v", got, err)
	}
	if got, err := parseCliScenarios("{not json"); err != nil || len(got) != 0 {
		t.Fatalf("invalid json should be tolerated: got=%v err=%v", got, err)
	}
	raw := `[{"name":"开通","description":"d","sort":1,"nodes":[{"id":"n1","type":"command","commandName":"user-create","note":"记录id"}],"edges":[]}]`
	got, err := parseCliScenarios(raw)
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if len(got) != 1 || got[0].Name != "开通" || len(got[0].Nodes) != 1 || got[0].Nodes[0].CommandName != "user-create" {
		t.Fatalf("parsed wrong: %+v", got)
	}
}

func TestFilterScenariosDropsUnknownNodesAndDanglingEdges(t *testing.T) {
	scenarios := []autoModel.CliScenario{{
		Name: "s",
		Nodes: []autoModel.CliScenarioNode{
			{ID: "n1", Type: "command", CommandName: "user-create"},
			{ID: "n2", Type: "command", CommandName: "ghost"},
			{ID: "n3", Type: "decision"},
		},
		Edges: []autoModel.CliScenarioEdge{
			{From: "n1", To: "n2"},
			{From: "n1", To: "n3"},
		},
	}}
	valid := map[string]bool{"user-create": true}
	got := filterScenarios(scenarios, valid)
	if len(got) != 1 || len(got[0].Nodes) != 2 {
		t.Fatalf("nodes filter wrong: %+v", got)
	}
	if len(got[0].Edges) != 1 || got[0].Edges[0].To != "n3" {
		t.Fatalf("dangling edge not removed: %+v", got[0].Edges)
	}
}

func TestRenderScenariosMarkdownEmpty(t *testing.T) {
	if got := renderScenariosMarkdown(nil, true); got != "" {
		t.Fatalf("empty should render empty, got %q", got)
	}
}

func TestRenderScenariosMarkdownLinearDecisionMerge(t *testing.T) {
	scenarios := []autoModel.CliScenario{{
		Name:        "开通用户",
		Description: "按状态分支",
		Sort:        1,
		Nodes: []autoModel.CliScenarioNode{
			{ID: "n1", Type: "command", CommandName: "user-create", Note: "创建用户"},
			{ID: "n2", Type: "decision", Note: "基于 第1步.status"},
			{ID: "n3", Type: "command", CommandName: "notify-send", Note: "发通知", InputNote: "第1步 id"},
			{ID: "n4", Type: "command", CommandName: "log-record", Note: "记日志"},
		},
		Edges: []autoModel.CliScenarioEdge{
			{From: "n1", To: "n2"},
			{From: "n2", To: "n3", Condition: "status=active"},
			{From: "n2", To: "n4", Condition: "status=pending"},
		},
	}}
	got := renderScenariosMarkdown(scenarios, true)
	for _, want := range []string{
		"## 典型场景", "### 开通用户", "按状态分支",
		"1. `user-create`", "创建用户",
		"2. 判断：基于 第1步.status",
		"若 status=active → `notify-send`",
		"若 status=pending → `log-record`",
		"3. `notify-send`", "入参：第1步 id",
		"4. `log-record`",
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("missing %q in\n%s", want, got)
		}
	}
}

func TestRenderScenariosMarkdownWithAlias(t *testing.T) {
	scenarios := []autoModel.CliScenario{{
		Name: "开通用户",
		Nodes: []autoModel.CliScenarioNode{
			{ID: "n1", Type: "command", CommandName: "user-create", Alias: "create", Note: "创建用户"},
			{ID: "n2", Type: "decision", Alias: "check", Note: "基于 create.status 合并判断"},
			{ID: "n3", Type: "command", CommandName: "notify-send", Alias: "notify", InputNote: "create.id"},
		},
		Edges: []autoModel.CliScenarioEdge{
			{From: "n1", To: "n2"},
			{From: "n2", To: "n3", Condition: "create.status=active"},
		},
	}}
	got := renderScenariosMarkdown(scenarios, true)
	for _, want := range []string{
		"1. `user-create`（别名：create）",
		"2. 判断：基于 create.status 合并判断（别名：check）",
		"3. `notify-send`（别名：notify）",
		"入参：create.id",
		"若 create.status=active → `notify-send`",
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("missing %q in\n%s", want, got)
		}
	}
}

// command 节点出边上的流转条件必须渲染出来（回归 bug：此前只有 decision 节点渲染出边条件，
// 挂在 command 节点连线上的条件被整段丢弃，场景无法体现该 step 何时被调用）。
func TestRenderScenariosMarkdownCommandBranchConditions(t *testing.T) {
	scenarios := []autoModel.CliScenario{{
		Name: "支付分支",
		Nodes: []autoModel.CliScenarioNode{
			{ID: "n1", Type: "command", CommandName: "order-create", Alias: "create", Note: "创建订单"},
			{ID: "n2", Type: "command", CommandName: "pay-notify", Note: "支付成功通知"},
			{ID: "n3", Type: "command", CommandName: "order-cancel", Note: "超时取消"},
		},
		Edges: []autoModel.CliScenarioEdge{
			{From: "n1", To: "n2", Condition: "create.status=paid"},
			{From: "n1", To: "n3", Condition: "create.status=timeout"},
		},
	}}
	got := renderScenariosMarkdown(scenarios, true)
	for _, want := range []string{
		"1. `order-create`（别名：create）",
		"若 create.status=paid → `pay-notify`",
		"若 create.status=timeout → `order-cancel`",
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("missing %q in\n%s", want, got)
		}
	}
}

// 只有单条带条件的 command 出边（“执行完这一步，满足条件才走到下一步”）也要渲染出条件。
func TestRenderScenariosMarkdownSingleCommandEdgeConditionRendered(t *testing.T) {
	scenarios := []autoModel.CliScenario{{
		Name: "单条件流转",
		Nodes: []autoModel.CliScenarioNode{
			{ID: "n1", Type: "command", CommandName: "user-create", Alias: "create"},
			{ID: "n2", Type: "command", CommandName: "welcome-send"},
		},
		Edges: []autoModel.CliScenarioEdge{
			{From: "n1", To: "n2", Condition: "create.isNew=true"},
		},
	}}
	got := renderScenariosMarkdown(scenarios, true)
	if !strings.Contains(got, "若 create.isNew=true → `welcome-send`") {
		t.Fatalf("single command-edge condition dropped:\n%s", got)
	}
}

// 单条无条件 command 出边保持隐式（靠编号顺序表达），不产生冗余箭头，避免线性链路噪声。
func TestRenderScenariosMarkdownSinglePlainCommandEdgeStaysImplicit(t *testing.T) {
	scenarios := []autoModel.CliScenario{{
		Name: "线性",
		Nodes: []autoModel.CliScenarioNode{
			{ID: "n1", Type: "command", CommandName: "step-a"},
			{ID: "n2", Type: "command", CommandName: "step-b"},
		},
		Edges: []autoModel.CliScenarioEdge{
			{From: "n1", To: "n2"},
		},
	}}
	got := renderScenariosMarkdown(scenarios, true)
	if strings.Contains(got, "→") {
		t.Fatalf("single unconditional command edge should stay implicit, got:\n%s", got)
	}
}

// command 节点混合“带条件 + 无条件”出边：条件边渲染“若…”，无条件边渲染“默认流转”，不出现“否则”。
func TestRenderScenariosMarkdownCommandMixedEdges(t *testing.T) {
	scenarios := []autoModel.CliScenario{{
		Name: "混合分支",
		Nodes: []autoModel.CliScenarioNode{
			{ID: "n1", Type: "command", CommandName: "order-create", Alias: "create"},
			{ID: "n2", Type: "command", CommandName: "pay-notify"},
			{ID: "n3", Type: "command", CommandName: "log-record"},
		},
		Edges: []autoModel.CliScenarioEdge{
			{From: "n1", To: "n2", Condition: "create.status=paid"},
			{From: "n1", To: "n3"},
		},
	}}
	got := renderScenariosMarkdown(scenarios, true)
	for _, want := range []string{
		"若 create.status=paid → `pay-notify`",
		"默认流转 → `log-record`",
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("missing %q in\n%s", want, got)
		}
	}
	if strings.Contains(got, "否则") {
		t.Fatalf("unconditional edge should not render 否则:\n%s", got)
	}
}

// 悬空出边（目标节点不存在，MCP 预览路径不过滤）应渲染可读占位而非空 backtick，且不 panic。
func TestRenderScenariosMarkdownDanglingEdgeTarget(t *testing.T) {
	scenarios := []autoModel.CliScenario{{
		Name: "悬空边",
		Nodes: []autoModel.CliScenarioNode{
			{ID: "n1", Type: "command", CommandName: "order-create"},
			{ID: "n2", Type: "command", CommandName: "pay-notify"},
		},
		Edges: []autoModel.CliScenarioEdge{
			{From: "n1", To: "n2", Condition: "已支付"},
			{From: "n1", To: "ghost", Condition: "未知目标"},
		},
	}}
	got := renderScenariosMarkdown(scenarios, false)
	if !strings.Contains(got, "若 未知目标 → （未知节点）") {
		t.Fatalf("dangling target should render placeholder:\n%s", got)
	}
	if strings.Contains(got, "→ ``") {
		t.Fatalf("dangling target must not render empty backticks:\n%s", got)
	}
}

func TestRenderScenariosMarkdownCycleDoesNotPanic(t *testing.T) {
	scenarios := []autoModel.CliScenario{{
		Name: "环",
		Nodes: []autoModel.CliScenarioNode{
			{ID: "n1", Type: "command", CommandName: "cmd-a"},
			{ID: "n2", Type: "command", CommandName: "cmd-b"},
		},
		Edges: []autoModel.CliScenarioEdge{
			{From: "n1", To: "n2"},
			{From: "n2", To: "n1"},
		},
	}}
	got := renderScenariosMarkdown(scenarios, true)
	if !strings.Contains(got, "## 典型场景") || !strings.Contains(got, "`cmd-a`") {
		t.Fatalf("cycle should degrade gracefully, got %q", got)
	}
}

func TestBuildSkillRenderDataInjectsScenarios(t *testing.T) {
	cli := autoModel.SysCli{
		Command: "opsctl", DisplayName: "运维",
		ScenariosJSON: `[{"name":"开通","description":"d","sort":1,"nodes":[{"id":"n1","type":"command","commandName":"user-list","note":"n"}],"edges":[]}]`,
	}
	manifest := autoRes.SysCliManifestResponse{Commands: []autoRes.SysCliManifestCommand{{Name: "user-list"}}}
	data := buildSkillRenderData(cli, manifest)
	if !strings.Contains(data.ScenariosMarkdown, "### 开通") || !strings.Contains(data.ScenariosMarkdown, "`user-list`") {
		t.Fatalf("scenarios not injected: %q", data.ScenariosMarkdown)
	}
	cli.ScenariosJSON = `[{"name":"s","nodes":[{"id":"n1","type":"command","commandName":"ghost"}],"edges":[]}]`
	data2 := buildSkillRenderData(cli, manifest)
	if data2.ScenariosMarkdown != "" {
		t.Fatalf("unknown command should filter out scenario, got %q", data2.ScenariosMarkdown)
	}
}

func TestRenderSkillBodyContainsScenarioSection(t *testing.T) {
	data := skillRenderData{
		Command: "opsctl", SkillName: "opsctl-cli",
		ScenariosMarkdown: "## 典型场景\n\n### 开通\n\n1. `user-list`\n",
	}
	body := renderSkillBody(data)
	for _, want := range []string{"## 典型场景", "### 开通", "`user-list`"} {
		if !strings.Contains(body, want) {
			t.Fatalf("body missing %q\n%s", want, body)
		}
	}
}
