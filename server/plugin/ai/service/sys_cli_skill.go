package service

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"text/template"

	autoModel "github.com/flipped-aurora/gin-vue-admin/server/plugin/ai/model"
	autoReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/ai/model/request"
	autoRes "github.com/flipped-aurora/gin-vue-admin/server/plugin/ai/model/response"
)

type skillParam struct {
	Flag        string
	Type        string
	Required    bool
	Description string
}

type skillResponseField struct {
	Name        string
	Description string
}

type skillCommand struct {
	CommandName string
	Summary     string
	Parameters  []skillParam
	Response    []skillResponseField
	Example     string
}

type skillRenderData struct {
	Command           string
	DisplayName       string
	Description       string
	SkillName         string
	SkillDescription  string
	Commands          []skillCommand
	ScenariosMarkdown string
}

const skillBodyTemplate = `# {{.Command}} — {{.DisplayName}} 命令行

{{if .Description}}{{.Description}}{{else}}由 Gin-Vue-Admin AI CLI 生成的命令行工具。{{end}}

## 安装与登录

1. 可执行文件已包含在本 skill 包目录下（与 SKILL.md 同级）：{{.Command}}（Windows 为 {{.Command}}.exe）。可直接调用，或将其所在目录加入 PATH。
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
{{if .Response}}
返回：
{{range .Response}}- ` + "`{{.Name}}`" + `{{if .Description}} — {{.Description}}{{end}}
{{end}}{{end}}
示例：

` + "```bash" + `
{{.Example}}
` + "```" + `
{{end}}{{end}}
{{if .ScenariosMarkdown}}
{{.ScenariosMarkdown}}
{{end}}
## 说明

- 用 {{.Command}} --help 查看所有命令；用 {{.Command}} <命令> --help 查看单个命令的完整参数。
- 登录态保存在 ~/.gva/config.json，更换账号时重新执行 login。
`

// buildSkillRenderData 从 CLI 实体与 manifest 派生渲染所需的数据结构。
func buildSkillRenderData(cli autoModel.SysCli, manifest autoRes.SysCliManifestResponse) skillRenderData {
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
		for _, r := range cmd.Response {
			sc.Response = append(sc.Response, skillResponseField{
				Name:        r.Name,
				Description: strings.TrimSpace(r.Description),
			})
		}
		data.Commands = append(data.Commands, sc)
	}
	// 场景链路：解析 → 按命令名过滤 → 渲染为 markdown 注入模板
	scenarios, _ := parseCliScenarios(cli.ScenariosJSON)
	validCmds := make(map[string]bool, len(manifest.Commands))
	for _, c := range manifest.Commands {
		validCmds[c.Name] = true
	}
	data.ScenariosMarkdown = renderScenariosMarkdown(filterScenarios(scenarios, validCmds))
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

// parseCliScenarios 解析 SysCli.ScenariosJSON；空或非法一律返回空切片与 nil，绝不阻塞 skill 生成。
func parseCliScenarios(raw string) ([]autoModel.CliScenario, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil, nil
	}
	var scenarios []autoModel.CliScenario
	if err := json.Unmarshal([]byte(raw), &scenarios); err != nil {
		return nil, nil
	}
	return scenarios, nil
}

// filterScenarios 过滤图：剔除引用不存在命令的 command 节点（decision 保留），并删除悬空边；
// 节点全被过滤的场景整体剔除。
func filterScenarios(scenarios []autoModel.CliScenario, valid map[string]bool) []autoModel.CliScenario {
	out := make([]autoModel.CliScenario, 0, len(scenarios))
	for _, s := range scenarios {
		s.Nodes, s.Edges = filterScenarioGraph(s.Nodes, s.Edges, valid)
		if len(s.Nodes) == 0 {
			continue
		}
		out = append(out, s)
	}
	return out
}

func filterScenarioGraph(nodes []autoModel.CliScenarioNode, edges []autoModel.CliScenarioEdge, valid map[string]bool) ([]autoModel.CliScenarioNode, []autoModel.CliScenarioEdge) {
	keep := make(map[string]bool, len(nodes))
	filtered := make([]autoModel.CliScenarioNode, 0, len(nodes))
	for _, n := range nodes {
		switch n.Type {
		case "command":
			if !valid[n.CommandName] {
				continue
			}
			keep[n.ID] = true
			filtered = append(filtered, n)
		case "decision":
			keep[n.ID] = true
			filtered = append(filtered, n)
		default:
			// 未知类型忽略
		}
	}
	edgeOut := make([]autoModel.CliScenarioEdge, 0, len(edges))
	for _, e := range edges {
		if keep[e.From] && keep[e.To] {
			edgeOut = append(edgeOut, e)
		}
	}
	return filtered, edgeOut
}

// renderScenariosMarkdown 把每个场景的图按拓扑序渲染为「## 典型场景」markdown；空返回空串。
func renderScenariosMarkdown(scenarios []autoModel.CliScenario) string {
	if len(scenarios) == 0 {
		return ""
	}
	var b strings.Builder
	b.WriteString("## 典型场景\n\n")
	for _, s := range scenarios {
		name := strings.TrimSpace(s.Name)
		if name == "" {
			name = "未命名场景"
		}
		b.WriteString("### " + name + "\n\n")
		if desc := strings.TrimSpace(s.Description); desc != "" {
			b.WriteString(desc + "\n\n")
		}
		if len(s.Nodes) == 0 {
			continue
		}
		writeScenarioGraph(&b, s.Nodes, s.Edges)
		b.WriteString("\n")
	}
	return strings.TrimRight(b.String(), "\n") + "\n"
}

func writeScenarioGraph(b *strings.Builder, nodes []autoModel.CliScenarioNode, edges []autoModel.CliScenarioEdge) {
	order := topologicalOrder(nodes, edges)
	nodeByID := make(map[string]autoModel.CliScenarioNode, len(nodes))
	for _, n := range nodes {
		nodeByID[n.ID] = n
	}
	outEdges := make(map[string][]autoModel.CliScenarioEdge, len(nodes))
	for _, e := range edges {
		outEdges[e.From] = append(outEdges[e.From], e)
	}
	for i, id := range order {
		n := nodeByID[id]
		switch n.Type {
		case "command":
			line := fmt.Sprintf("%d. `%s`", i+1, strings.TrimSpace(n.CommandName))
			if alias := strings.TrimSpace(n.Alias); alias != "" {
				line += "（别名：" + alias + "）"
			}
			if note := strings.TrimSpace(n.Note); note != "" {
				line += " — " + note
			}
			if in := strings.TrimSpace(n.InputNote); in != "" {
				line += "（入参：" + in + "）"
			}
			b.WriteString(line + "\n")
		case "decision":
			desc := strings.TrimSpace(n.Note)
			if desc == "" {
				desc = "判断"
			}
			head := fmt.Sprintf("%d. 判断：%s", i+1, desc)
			if alias := strings.TrimSpace(n.Alias); alias != "" {
				head += "（别名：" + alias + "）"
			}
			b.WriteString(head + "\n")
			for _, e := range outEdges[n.ID] {
				target := nodeLabel(nodeByID[e.To])
				if cond := strings.TrimSpace(e.Condition); cond != "" {
					b.WriteString(fmt.Sprintf("   - 若 %s → %s\n", cond, target))
				} else {
					b.WriteString(fmt.Sprintf("   - 否则 → %s\n", target))
				}
			}
		}
	}
}

func nodeLabel(n autoModel.CliScenarioNode) string {
	if n.Type == "decision" {
		return "判断"
	}
	return "`" + strings.TrimSpace(n.CommandName) + "`"
}

// topologicalOrder 返回节点 ID 的拓扑序列；入度为 0 的节点优先（按 nodes 顺序入队保证稳定）。
// 有环时，环内节点按 nodes 顺序降级追加到末尾，保证不阻塞渲染。
func topologicalOrder(nodes []autoModel.CliScenarioNode, edges []autoModel.CliScenarioEdge) []string {
	indeg := make(map[string]int, len(nodes))
	adj := make(map[string][]string, len(nodes))
	for _, n := range nodes {
		indeg[n.ID] = 0
	}
	for _, e := range edges {
		if _, ok := indeg[e.From]; !ok {
			continue
		}
		if _, ok := indeg[e.To]; !ok {
			continue
		}
		adj[e.From] = append(adj[e.From], e.To)
		indeg[e.To]++
	}
	queue := make([]string, 0)
	for _, n := range nodes {
		if indeg[n.ID] == 0 {
			queue = append(queue, n.ID)
		}
	}
	result := make([]string, 0, len(nodes))
	for len(queue) > 0 {
		id := queue[0]
		queue = queue[1:]
		result = append(result, id)
		for _, to := range adj[id] {
			indeg[to]--
			if indeg[to] == 0 {
				queue = append(queue, to)
			}
		}
	}
	if len(result) < len(nodes) {
		appended := make(map[string]bool, len(result))
		for _, id := range result {
			appended[id] = true
		}
		for _, n := range nodes {
			if !appended[n.ID] {
				result = append(result, n.ID)
				appended[n.ID] = true
			}
		}
	}
	return result
}

// BuildCliSkill 生成该 CLI 的 AI 使用说明并和编译好的二进制一起打包成 zip：
// SKILL.md + references/README.md + references/manifest.json + cli 二进制。
func (s *cliService) BuildCliSkill(req autoReq.BuildSysCliBinaryRequest) (string, []byte, error) {
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
