package service

import (
	"encoding/json"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	autoModel "github.com/flipped-aurora/gin-vue-admin/server/plugin/auto/model"
	autoRes "github.com/flipped-aurora/gin-vue-admin/server/plugin/auto/model/response"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

func TestBuildSysCliManifestUsesEnabledBindingsAndCommandPrecedence(t *testing.T) {
	originalMCP := global.GVA_CONFIG.MCP
	originalSystem := global.GVA_CONFIG.System
	global.GVA_CONFIG.MCP.BaseURL = "http://manifest.example.com/mcp/"
	global.GVA_CONFIG.MCP.UpstreamBaseURL = ""
	global.GVA_CONFIG.MCP.AuthHeader = ""
	global.GVA_CONFIG.System.RouterPrefix = "api"
	defer func() {
		global.GVA_CONFIG.MCP = originalMCP
		global.GVA_CONFIG.System = originalSystem
	}()

	cli := autoModel.SysCli{Name: "user-manager", Command: "opsctl", Version: "v2", Description: "用户管理工具"}
	bindings := []sysCliManifestBinding{
		{Binding: autoModel.SysCliApi{ApiID: 3, CommandName: "custom-users", Enabled: true, Sort: 2}, Api: sysModel.SysApi{Path: "/user/getUserList", Method: "POST", Description: "获取用户列表", ApiGroup: "系统用户"}},
		{Binding: autoModel.SysCliApi{ApiID: 1, Enabled: true, Sort: 1}, Api: sysModel.SysApi{Path: "/user/getUserInfo", Method: "GET", Description: "获取用户信息", ApiGroup: "系统用户"}},
		{Binding: autoModel.SysCliApi{ApiID: 2, CommandName: "ignored", Enabled: false, Sort: 0}, Api: sysModel.SysApi{Path: "/user/deleteUser", Method: "DELETE", Description: "删除用户", ApiGroup: "系统用户"}},
	}

	manifest, err := buildSysCliManifest(cli, bindings)
	if err != nil {
		t.Fatalf("buildSysCliManifest() error = %v", err)
	}
	if manifest.Name != "opsctl" {
		t.Fatalf("manifest.Name = %q, want %q", manifest.Name, "opsctl")
	}
	if manifest.Version != "v2" {
		t.Fatalf("manifest.Version = %q, want %q", manifest.Version, "v2")
	}
	if manifest.Server.BaseURL != "http://manifest.example.com/mcp" {
		t.Fatalf("manifest.Server.BaseURL = %q, want %q", manifest.Server.BaseURL, "http://manifest.example.com/mcp")
	}
	if manifest.Server.AuthHeader != "x-token" {
		t.Fatalf("manifest.Server.AuthHeader = %q, want %q", manifest.Server.AuthHeader, "x-token")
	}
	if len(manifest.Commands) != 2 {
		t.Fatalf("len(manifest.Commands) = %d, want 2", len(manifest.Commands))
	}
	if manifest.Commands[0].Name != "user-info" {
		t.Fatalf("manifest.Commands[0].Name = %q, want %q", manifest.Commands[0].Name, "user-info")
	}
	if manifest.Commands[1].Name != "custom-users" {
		t.Fatalf("manifest.Commands[1].Name = %q, want %q", manifest.Commands[1].Name, "custom-users")
	}
	if len(manifest.Commands[0].Examples) == 0 || manifest.Commands[0].Examples[0] != "opsctl user-info --help" {
		t.Fatalf("manifest.Commands[0].Examples[0] = %q, want %q", manifest.Commands[0].Examples[0], "opsctl user-info --help")
	}
}

func TestApplySysCliDefaultsAndSlugFallback(t *testing.T) {
	cli := autoModel.SysCli{Name: "ops-tool"}
	applySysCliDefaults(&cli)
	if cli.Command != "ops-tool" {
		t.Fatalf("Command = %q, want %q", cli.Command, "ops-tool")
	}
	if cli.DisplayName != "ops-tool" {
		t.Fatalf("DisplayName = %q, want %q", cli.DisplayName, "ops-tool")
	}
	if cli.Version != "v1" {
		t.Fatalf("Version = %q, want %q", cli.Version, "v1")
	}
	if cli.Status != "enabled" {
		t.Fatalf("Status = %q, want %q", cli.Status, "enabled")
	}
	if cli.AuthMode != "jwt" {
		t.Fatalf("AuthMode = %q, want %q", cli.AuthMode, "jwt")
	}

	slug := deriveSingleSegmentSlugFromAPI(sysModel.SysApi{Path: "/user/getUserList", Method: "POST"})
	if slug != "user-list" {
		t.Fatalf("slug = %q, want %q", slug, "user-list")
	}
}

func TestSysCliManifestFileNameIsStable(t *testing.T) {
	filename := sysCliManifestFileName(autoModel.SysCli{Name: "user-manager", Command: "opsctl"})
	if filename != "opsctl.manifest.json" {
		t.Fatalf("filename = %q, want %q", filename, "opsctl.manifest.json")
	}
}

func TestSysCliManifestDownloadPayloadIsJSON(t *testing.T) {
	manifest := autoRes.SysCliManifestResponse{
		Name:     "ops-tool",
		Version:  "v1",
		Server:   autoRes.SysCliManifestServer{BaseURL: "http://127.0.0.1:8889/mcp", AuthHeader: "x-token"},
		Commands: []autoRes.SysCliManifestCommand{{Name: "health", Method: "GET", Path: "/health"}},
	}
	payload, err := marshalSysCliManifest(manifest)
	if err != nil {
		t.Fatalf("marshalSysCliManifest() error = %v", err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	if decoded["name"] != "ops-tool" {
		t.Fatalf("decoded name = %v, want %v", decoded["name"], "ops-tool")
	}
}

func TestSwaggerResponseToFieldsExpandsDataRef(t *testing.T) {
	spec := &swaggerSpec{
		RawDefinitions: map[string]map[string]interface{}{
			"response.SysAPIResponse": {"properties": map[string]interface{}{
				"api": map[string]interface{}{"type": "object", "description": "API对象"},
				"msg": map[string]interface{}{"type": "string", "description": "提示信息"},
			}},
		},
	}
	operation := swaggerOperation{
		Responses: map[string]swaggerResponse{
			"200": {Description: "成功", Schema: map[string]interface{}{
				"allOf": []interface{}{
					map[string]interface{}{"$ref": "#/definitions/response.Response"},
					map[string]interface{}{"type": "object", "properties": map[string]interface{}{
						"data": map[string]interface{}{"$ref": "#/definitions/response.SysAPIResponse"},
					}},
				},
			}},
		},
	}
	fields, warnings := swaggerResponseToFields(operation, spec)
	if len(warnings) != 0 {
		t.Fatalf("unexpected warnings: %v", warnings)
	}
	if len(fields) != 2 {
		t.Fatalf("fields len = %d, want 2: %+v", len(fields), fields)
	}
	if fields[0].Name != "api" || fields[0].Description != "API对象" {
		t.Fatalf("fields[0] = %+v", fields[0])
	}
	if fields[1].Name != "msg" {
		t.Fatalf("fields[1] = %+v", fields[1])
	}
}

func TestSwaggerResponseToFieldsEmptyWhenNoData(t *testing.T) {
	spec := &swaggerSpec{}
	operation := swaggerOperation{
		Responses: map[string]swaggerResponse{
			"200": {Schema: map[string]interface{}{"$ref": "#/definitions/response.Response"}},
		},
	}
	fields, _ := swaggerResponseToFields(operation, spec)
	if len(fields) != 0 {
		t.Fatalf("expected no fields, got %+v", fields)
	}
}

func TestSwaggerResponseToFieldsRecursiveDottedPaths(t *testing.T) {
	spec := &swaggerSpec{
		RawDefinitions: map[string]map[string]interface{}{
			"response.UserResponse": {"properties": map[string]interface{}{
				"user": map[string]interface{}{"$ref": "#/definitions/system.User", "description": "用户对象"},
				"msg":  map[string]interface{}{"type": "string", "description": "提示"},
			}},
			"system.User": {"properties": map[string]interface{}{
				"name": map[string]interface{}{"type": "string", "description": "姓名"},
				"age":  map[string]interface{}{"type": "integer", "description": "年龄"},
			}},
		},
	}
	operation := swaggerOperation{Responses: map[string]swaggerResponse{
		"200": {Schema: map[string]interface{}{
			"allOf": []interface{}{
				map[string]interface{}{"$ref": "#/definitions/response.Response"},
				map[string]interface{}{"type": "object", "properties": map[string]interface{}{
					"data": map[string]interface{}{"$ref": "#/definitions/response.UserResponse"},
				}},
			},
		}},
	}}
	fields, _ := swaggerResponseToFields(operation, spec)
	got := map[string]string{}
	for _, f := range fields {
		got[f.Name] = f.Description
	}
	for _, want := range []string{"msg", "user.name", "user.age"} {
		if _, ok := got[want]; !ok {
			t.Fatalf("missing path %q in %+v", want, fields)
		}
	}
	if got["user.name"] != "姓名" {
		t.Fatalf("user.name desc = %q", got["user.name"])
	}
}

func TestSwaggerResponseToFieldsArrayItemDottedPaths(t *testing.T) {
	spec := &swaggerSpec{
		RawDefinitions: map[string]map[string]interface{}{
			"response.ListResp": {"properties": map[string]interface{}{
				"list": map[string]interface{}{"type": "array", "items": map[string]interface{}{"$ref": "#/definitions/system.Item"}},
			}},
			"system.Item": {"properties": map[string]interface{}{
				"id":   map[string]interface{}{"type": "integer"},
				"name": map[string]interface{}{"type": "string"},
			}},
		},
	}
	operation := swaggerOperation{Responses: map[string]swaggerResponse{
		"200": {Schema: map[string]interface{}{
			"allOf": []interface{}{
				map[string]interface{}{"$ref": "#/definitions/response.Response"},
				map[string]interface{}{"type": "object", "properties": map[string]interface{}{
					"data": map[string]interface{}{"$ref": "#/definitions/response.ListResp"},
				}},
			},
		}},
	}}
	fields, _ := swaggerResponseToFields(operation, spec)
	got := map[string]bool{}
	for _, f := range fields {
		got[f.Name] = true
	}
	for _, want := range []string{"list.id", "list.name"} {
		if !got[want] {
			t.Fatalf("missing path %q in %+v", want, fields)
		}
	}
}

func TestSwaggerResponseToFieldsMaxDepthTruncates(t *testing.T) {
	spec := &swaggerSpec{
		RawDefinitions: map[string]map[string]interface{}{
			"response.A": {"properties": map[string]interface{}{"b": map[string]interface{}{"$ref": "#/definitions/response.B"}}},
			"response.B": {"properties": map[string]interface{}{"c": map[string]interface{}{"$ref": "#/definitions/response.C"}}},
			"response.C": {"properties": map[string]interface{}{"d": map[string]interface{}{"$ref": "#/definitions/response.D"}}},
			"response.D": {"properties": map[string]interface{}{"e": map[string]interface{}{"type": "string"}}},
		},
	}
	operation := swaggerOperation{Responses: map[string]swaggerResponse{
		"200": {Schema: map[string]interface{}{
			"allOf": []interface{}{
				map[string]interface{}{"$ref": "#/definitions/response.Response"},
				map[string]interface{}{"type": "object", "properties": map[string]interface{}{
					"data": map[string]interface{}{"$ref": "#/definitions/response.A"},
				}},
			},
		}},
	}}
	fields, _ := swaggerResponseToFields(operation, spec)
	if len(fields) != 1 || fields[0].Name != "b.c.d" {
		t.Fatalf("expected single truncated path b.c.d, got %+v", fields)
	}
}

func TestSwaggerResponseToFieldsMissingRefBecomesLeaf(t *testing.T) {
	spec := &swaggerSpec{
		RawDefinitions: map[string]map[string]interface{}{
			"response.R": {"properties": map[string]interface{}{
				"x": map[string]interface{}{"$ref": "#/definitions/response.Missing"},
			}},
		},
	}
	operation := swaggerOperation{Responses: map[string]swaggerResponse{
		"200": {Schema: map[string]interface{}{
			"allOf": []interface{}{
				map[string]interface{}{"$ref": "#/definitions/response.Response"},
				map[string]interface{}{"type": "object", "properties": map[string]interface{}{
					"data": map[string]interface{}{"$ref": "#/definitions/response.R"},
				}},
			},
		}},
	}}
	fields, _ := swaggerResponseToFields(operation, spec)
	if len(fields) != 1 || fields[0].Name != "x" {
		t.Fatalf("expected leaf x for missing ref, got %+v", fields)
	}
}

func TestSwaggerResponseToFieldsArrayDrillsOnlyOneLevel(t *testing.T) {
	spec := &swaggerSpec{
		RawDefinitions: map[string]map[string]interface{}{
			"system.Item": {"properties": map[string]interface{}{
				"id": map[string]interface{}{"type": "integer"},
			}},
			"response.Nested": {"properties": map[string]interface{}{
				"matrix": map[string]interface{}{"type": "array", "items": map[string]interface{}{
					"type":  "array",
					"items": map[string]interface{}{"$ref": "#/definitions/system.Item"},
				}},
			}},
		},
	}
	operation := swaggerOperation{Responses: map[string]swaggerResponse{
		"200": {Schema: map[string]interface{}{
			"allOf": []interface{}{
				map[string]interface{}{"$ref": "#/definitions/response.Response"},
				map[string]interface{}{"type": "object", "properties": map[string]interface{}{
					"data": map[string]interface{}{"$ref": "#/definitions/response.Nested"},
				}},
			},
		}},
	}}
	// 数组只下钻一层：matrix 的 items 仍是数组，不再下钻，matrix 作为叶子
	fields, _ := swaggerResponseToFields(operation, spec)
	if len(fields) != 1 || fields[0].Name != "matrix" {
		t.Fatalf("expected single leaf matrix (array drills only one level), got %+v", fields)
	}
}
