package system

import (
	"encoding/json"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
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

	cli := sysModel.SysCli{Name: "user-manager", Command: "opsctl", Version: "v2", Description: "用户管理工具"}
	bindings := []sysCliManifestBinding{
		{Binding: sysModel.SysCliApi{ApiID: 3, CommandName: "custom-users", Enabled: true, Sort: 2}, Api: sysModel.SysApi{Path: "/user/getUserList", Method: "POST", Description: "获取用户列表", ApiGroup: "系统用户"}},
		{Binding: sysModel.SysCliApi{ApiID: 1, Enabled: true, Sort: 1}, Api: sysModel.SysApi{Path: "/user/getUserInfo", Method: "GET", Description: "获取用户信息", ApiGroup: "系统用户"}},
		{Binding: sysModel.SysCliApi{ApiID: 2, CommandName: "ignored", Enabled: false, Sort: 0}, Api: sysModel.SysApi{Path: "/user/deleteUser", Method: "DELETE", Description: "删除用户", ApiGroup: "系统用户"}},
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
	cli := sysModel.SysCli{Name: "ops-tool"}
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
	filename := sysCliManifestFileName(sysModel.SysCli{Name: "user-manager", Command: "opsctl"})
	if filename != "opsctl.manifest.json" {
		t.Fatalf("filename = %q, want %q", filename, "opsctl.manifest.json")
	}
}

func TestSysCliManifestDownloadPayloadIsJSON(t *testing.T) {
	manifest := systemRes.SysCliManifestResponse{
		Name:    "ops-tool",
		Version: "v1",
		Server: systemRes.SysCliManifestServer{BaseURL: "http://127.0.0.1:8889/mcp", AuthHeader: "x-token"},
		Commands: []systemRes.SysCliManifestCommand{{Name: "health", Method: "GET", Path: "/health"}},
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
