package system

import (
	"archive/zip"
	"bytes"
	"fmt"
	"strings"
	"text/template"

	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
)

type skillParam struct {
	Flag        string
	Type        string
	Required    bool
	Description string
}

type skillCommand struct {
	CommandName string
	Summary     string
	Parameters  []skillParam
	Example     string
}

type skillRenderData struct {
	Command          string
	DisplayName      string
	Description      string
	SkillName        string
	SkillDescription string
	Commands         []skillCommand
}

const skillBodyTemplate = `# {{.Command}} — {{.DisplayName}} 命令行

{{if .Description}}{{.Description}}{{else}}由 Gin-Vue-Admin AI CLI 生成的命令行工具。{{end}}

## 安装与登录

1. 获取可执行文件：从 Gin-Vue-Admin「系统工具 → AI CLI管理 → 预览命令 → 编译下载」下载 {{.Command}}（Windows 为 {{.Command}}.exe），放到 PATH。
2. 登录（只需一次）：{{.Command}} login --token <JWT>。JWT 在 Gin-Vue-Admin 登录后获得，或使用系统工具里的 API Token。
3. 更改后台 API 地址：{{.Command}} set-base-url <API地址>（写入本地配置，之后命令都用新地址）。

## 命令一览

| 命令 | 说明 |
| --- | --- |
{{range .Commands}}| ` + "`{{.CommandName}}`" + ` | {{.Summary}} |
{{end}}
{{if .Commands}}
## 命令详情

{{range .Commands}}
### {{.CommandName}}

{{if .Summary}}{{.Summary}}{{else}}-{{end}}

参数：{{if .Parameters}}
{{range .Parameters}}- ` + "`--{{.Flag}}`" + ` ({{.Type}}{{if .Required}}, 必填{{end}}){{if .Description}} — {{.Description}}{{end}}
{{end}}{{else}}无{{end}}

示例：

` + "```bash" + `
{{.Example}}
` + "```" + `
{{end}}{{end}}
## 说明

- 用 {{.Command}} --help 查看所有命令；用 {{.Command}} <命令> --help 查看单个命令的完整参数。
- 登录态保存在 ~/.gva/config.json，更换账号时重新执行 login。
`

// buildSkillRenderData 从 CLI 实体与 manifest 派生渲染所需的数据结构。
func buildSkillRenderData(cli sysModel.SysCli, manifest systemRes.SysCliManifestResponse) skillRenderData {
	data := skillRenderData{
		Command:     strings.TrimSpace(cli.Command),
		DisplayName: strings.TrimSpace(cli.DisplayName),
		Description: strings.TrimSpace(cli.Description),
		SkillName:   strings.TrimSpace(cli.SkillName),
	}
	if data.SkillName == "" {
		slug := sanitizeSingleSegmentSlug(cli.Command)
		if slug == "" {
			slug = sanitizeSingleSegmentSlug(cli.Name)
		}
		if slug == "" {
			slug = "cli"
		}
		data.SkillName = slug + "-cli"
	}
	data.SkillDescription = strings.TrimSpace(cli.SkillDescription)
	if data.SkillDescription == "" {
		display := data.DisplayName
		if display == "" {
			display = data.Command
		}
		data.SkillDescription = fmt.Sprintf("通过 %s 命令行调用 %s 的后台 API。当需要在终端、脚本中或让 AI 助手代为执行这些接口时使用。", data.Command, display)
	}
	for _, cmd := range manifest.Commands {
		example := fmt.Sprintf("%s %s", data.Command, cmd.Name)
		for _, e := range cmd.Examples {
			if strings.TrimSpace(e) != "" {
				example = e
				break
			}
		}
		sc := skillCommand{
			CommandName: cmd.Name,
			Summary:     strings.TrimSpace(cmd.Summary),
			Example:     example,
		}
		for _, p := range cmd.Parameters {
			sc.Parameters = append(sc.Parameters, skillParam{
				Flag:        p.Flag,
				Type:        paramTypeLabel(p.Type),
				Required:    p.Required,
				Description: strings.TrimSpace(p.Description),
			})
		}
		data.Commands = append(data.Commands, sc)
	}
	return data
}

func paramTypeLabel(typeName string) string {
	typeName = strings.TrimSpace(typeName)
	if typeName == "" {
		return "string"
	}
	return typeName
}

// renderSkillBody 渲染命令行使用说明正文（SKILL.md 和 README.md 共用）。
func renderSkillBody(data skillRenderData) string {
	tpl := template.Must(template.New("skillBody").Parse(skillBodyTemplate))
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, data); err != nil {
		return ""
	}
	return buf.String()
}

// renderSkillMarkdown 渲染完整 SKILL.md，带 frontmatter。
func renderSkillMarkdown(data skillRenderData) string {
	frontmatter := fmt.Sprintf("---\nname: %s\ndescription: %s\n---\n\n", data.SkillName, escapeSkillDescription(data.SkillDescription))
	return frontmatter + renderSkillBody(data)
}

// escapeSkillDescription 让 description 单行化，避免 frontmatter 解析问题。
func escapeSkillDescription(desc string) string {
	desc = strings.ReplaceAll(desc, "\r", " ")
	desc = strings.ReplaceAll(desc, "\n", " ")
	return strings.TrimSpace(desc)
}

// BuildCliSkill 生成该 CLI 的 AI 使用说明并和编译好的二进制一起打包成 zip：
// SKILL.md + references/README.md + references/manifest.json + cli 二进制。
func (s *cliService) BuildCliSkill(req systemReq.BuildSysCliBinaryRequest) (string, []byte, error) {
	goos, goarch, err := normalizeBuildTarget(req.GOOS, req.GOARCH)
	if err != nil {
		return "", nil, err
	}
	cli, bindings, err := s.getCliAndBindings(req.CliID)
	if err != nil {
		return "", nil, err
	}
	manifest, err := buildSysCliManifest(cli, bindings)
	if err != nil {
		return "", nil, err
	}
	applyCliBuildBaseURL(&manifest, req.BaseURL)
	manifestBytes, err := marshalSysCliManifest(manifest)
	if err != nil {
		return "", nil, err
	}
	binaryName, binaryBytes, err := s.compileCliBinary(cli, manifestBytes, goos, goarch)
	if err != nil {
		return "", nil, err
	}
	data := buildSkillRenderData(cli, manifest)
	skillMarkdown := renderSkillMarkdown(data)
	readmeMarkdown := "# " + data.SkillName + "\n\n" + renderSkillBody(data)

	folderName := sanitizeSingleSegmentSlug(data.SkillName)
	if folderName == "" {
		folderName = "cli"
	}
	zipBytes, err := writeSkillPackageZip(folderName, skillMarkdown, readmeMarkdown, manifestBytes, binaryName, binaryBytes)
	if err != nil {
		return "", nil, err
	}
	return folderName + ".zip", zipBytes, nil
}

func writeSkillPackageZip(folderName, skillMarkdown, readmeMarkdown string, manifestBytes []byte, binaryName string, binaryBytes []byte) ([]byte, error) {
	buf := &bytes.Buffer{}
	w := zip.NewWriter(buf)
	files := []struct {
		name    string
		content []byte
	}{
		{folderName + "/SKILL.md", []byte(skillMarkdown)},
		{folderName + "/references/README.md", []byte(readmeMarkdown)},
		{folderName + "/references/manifest.json", manifestBytes},
		{folderName + "/" + binaryName, binaryBytes},
	}
	for _, f := range files {
		fw, err := w.Create(f.name)
		if err != nil {
			return nil, err
		}
		if _, err := fw.Write(f.content); err != nil {
			return nil, err
		}
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
