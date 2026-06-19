package service

import (
	"encoding/json"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	autoModel "github.com/flipped-aurora/gin-vue-admin/server/plugin/auto/model"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

func TestParseParamsOverride(t *testing.T) {
	// 空串回退
	if _, ok := parseParamsOverride(""); ok {
		t.Fatal("empty should return false")
	}
	if _, ok := parseParamsOverride("   "); ok {
		t.Fatal("blank should return false")
	}
	// 非法 JSON 回退
	if _, ok := parseParamsOverride("not-json"); ok {
		t.Fatal("invalid json should return false")
	}
	// 合法覆盖
	raw := `[{"name":"keyword","flag":"keyword","type":"string","field":"keyword","location":"query","description":"关键字"}]`
	params, ok := parseParamsOverride(raw)
	if !ok {
		t.Fatal("valid override should return true")
	}
	if len(params) != 1 || params[0].Flag != "keyword" || params[0].Description != "关键字" {
		t.Fatalf("unexpected params: %+v", params)
	}
}

func TestBuildSysCliManifestCommandHonorsOverrides(t *testing.T) {
	originalMCP := global.GVA_CONFIG.MCP
	originalSystem := global.GVA_CONFIG.System
	global.GVA_CONFIG.MCP.BaseURL = ""
	global.GVA_CONFIG.MCP.UpstreamBaseURL = ""
	global.GVA_CONFIG.System.RouterPrefix = ""
	defer func() {
		global.GVA_CONFIG.MCP = originalMCP
		global.GVA_CONFIG.System = originalSystem
	}()

	cli := autoModel.SysCli{Name: "demo", Command: "demo"}
	overrideJSON, _ := json.Marshal([]sysCliParamJSON{
		{Flag: "keyword", Type: "string", Description: "搜索关键字", Field: "keyword", Location: "query"},
	})
	binding := sysCliManifestBinding{
		Binding: autoModel.SysCliApi{
			ApiID:          7,
			Enabled:        true,
			CommandDesc:    "按关键字搜索字典",
			ParamsOverride: string(overrideJSON),
		},
		Api: sysModel.SysApi{Path: "/sysDictionaryDetail/getSysDictionaryDetailList", Method: "GET", Description: "自动说明", ApiGroup: "字典"},
	}
	used := map[string]int{}
	cmd, err := buildSysCliManifestCommand(cli, &swaggerSpec{}, binding, used)
	if err != nil {
		t.Fatalf("buildSysCliManifestCommand: %v", err)
	}
	if cmd.Summary != "按关键字搜索字典" || cmd.Description != "按关键字搜索字典" {
		t.Fatalf("summary/description = %q/%q, want override", cmd.Summary, cmd.Description)
	}
	if len(cmd.Parameters) != 1 || cmd.Parameters[0].Flag != "keyword" || cmd.Parameters[0].Description != "搜索关键字" {
		t.Fatalf("parameters not from override: %+v", cmd.Parameters)
	}
}

type sysCliParamJSON struct {
	Name        string `json:"name,omitempty"`
	Flag        string `json:"flag"`
	Type        string `json:"type"`
	Required    bool   `json:"required,omitempty"`
	Description string `json:"description,omitempty"`
	Location    string `json:"location"`
	Field       string `json:"field"`
}

func TestParseResponseOverride(t *testing.T) {
	if _, ok := parseResponseOverride(""); ok {
		t.Fatal("empty should return false")
	}
	if _, ok := parseResponseOverride("   "); ok {
		t.Fatal("blank should return false")
	}
	if _, ok := parseResponseOverride("not-json"); ok {
		t.Fatal("invalid json should return false")
	}
	raw := `[{"name":"total","description":"总数"}]`
	fields, ok := parseResponseOverride(raw)
	if !ok {
		t.Fatal("valid override should return true")
	}
	if len(fields) != 1 || fields[0].Name != "total" || fields[0].Description != "总数" {
		t.Fatalf("unexpected fields: %+v", fields)
	}
}

type sysCliResponseFieldJSON struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func TestBuildSysCliManifestCommandApiBriefOverridesSummary(t *testing.T) {
	originalMCP := global.GVA_CONFIG.MCP
	originalSystem := global.GVA_CONFIG.System
	global.GVA_CONFIG.MCP.BaseURL = ""
	global.GVA_CONFIG.MCP.UpstreamBaseURL = ""
	global.GVA_CONFIG.System.RouterPrefix = ""
	defer func() {
		global.GVA_CONFIG.MCP = originalMCP
		global.GVA_CONFIG.System = originalSystem
	}()

	cli := autoModel.SysCli{Name: "demo", Command: "demo"}
	binding := sysCliManifestBinding{
		Binding: autoModel.SysCliApi{ApiID: 7, Enabled: true, ApiBrief: "简洁名", CommandDesc: "详细说明"},
		Api:     sysModel.SysApi{Path: "/sysDictionaryDetail/getSysDictionaryDetailList", Method: "GET", Description: "自动说明", ApiGroup: "字典"},
	}
	cmd, err := buildSysCliManifestCommand(cli, &swaggerSpec{}, binding, map[string]int{})
	if err != nil {
		t.Fatalf("buildSysCliManifestCommand: %v", err)
	}
	if cmd.Summary != "简洁名" {
		t.Fatalf("summary = %q, want %q", cmd.Summary, "简洁名")
	}
	if cmd.Description != "详细说明" {
		t.Fatalf("description = %q, want %q", cmd.Description, "详细说明")
	}
}

func TestBuildSysCliManifestCommandKeepsSummaryPriorityWithoutApiBrief(t *testing.T) {
	cli := autoModel.SysCli{Name: "demo", Command: "demo"}
	binding := sysCliManifestBinding{
		Binding: autoModel.SysCliApi{ApiID: 7, Enabled: true},
		Api:     sysModel.SysApi{Path: "/x/y", Method: "GET", Description: "API列表简介"},
	}
	cmd, err := buildSysCliManifestCommand(cli, &swaggerSpec{}, binding, map[string]int{})
	if err != nil {
		t.Fatalf("buildSysCliManifestCommand: %v", err)
	}
	if cmd.Summary != "API列表简介" {
		t.Fatalf("summary = %q, want %q (original priority, no override)", cmd.Summary, "API列表简介")
	}
}

func TestBuildSysCliManifestCommandResponseOverride(t *testing.T) {
	overrideJSON, _ := json.Marshal([]sysCliResponseFieldJSON{
		{Name: "total", Description: "总数"},
	})
	cli := autoModel.SysCli{Name: "demo", Command: "demo"}
	binding := sysCliManifestBinding{
		Binding: autoModel.SysCliApi{ApiID: 7, Enabled: true, ResponseOverride: string(overrideJSON)},
		Api:     sysModel.SysApi{Path: "/x/y", Method: "GET", Description: "d"},
	}
	cmd, err := buildSysCliManifestCommand(cli, &swaggerSpec{}, binding, map[string]int{})
	if err != nil {
		t.Fatalf("buildSysCliManifestCommand: %v", err)
	}
	if len(cmd.Response) != 1 || cmd.Response[0].Name != "total" {
		t.Fatalf("response not from override: %+v", cmd.Response)
	}
}
