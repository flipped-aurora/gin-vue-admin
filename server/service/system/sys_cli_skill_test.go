package system

import (
	"strings"
	"testing"

	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
)

func TestBuildSkillRenderDataFillsDefaults(t *testing.T) {
	cli := sysModel.SysCli{
		Name:        "user-manager",
		Command:     "opsctl",
		DisplayName: "运维工具",
		Description: "用于运维查询",
	}
	manifest := systemRes.SysCliManifestResponse{
		Name:    "opsctl",
		Version: "v1",
		Commands: []systemRes.SysCliManifestCommand{
			{Name: "user-list", Summary: "分页获取用户列表", Method: "POST", Path: "/user/list", Examples: []string{"opsctl user-list --page 1"}, Parameters: []systemRes.SysCliManifestParameter{
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
