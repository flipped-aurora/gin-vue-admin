package system

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"gopkg.in/yaml.v3"
)

const (
	skillFileName            = "SKILL.md"
	globalConstraintFileName = "README.md"
)

var skillToolOrder = []string{"copilot", "claude", "cursor", "trae", "codex"}

var skillToolDirs = map[string]string{
	"copilot": ".aone_copilot",
	"claude":  ".claude",
	"trae":    ".trae",
	"codex":   ".codex",
	"cursor":  ".cursor",
}

var skillToolLabels = map[string]string{
	"copilot": "Copilot",
	"claude":  "Claude",
	"trae":    "Trae",
	"codex":   "Codex",
	"cursor":  "Cursor",
}

const defaultSkillMarkdown = "## 技能用途\n请在这里描述技能的目标、适用场景与限制条件。\n\n## 输入\n- 请补充输入格式与示例。\n\n## 输出\n- 请补充输出格式与示例。\n\n## 关键步骤\n1. 第一步\n2. 第二步\n\n## 示例\n在此补充一到两个典型示例。\n"

const defaultResourceMarkdown = "# 资源说明\n请在这里补充资源内容。\n"

const defaultReferenceMarkdown = "# 参考资料\n请在这里补充参考资料内容。\n"

const defaultTemplateMarkdown = "# 模板\n请在这里补充模板内容。\n"

const defaultGlobalConstraintMarkdown = "# 全局约束\n请在这里补充该工具的统一约束与使用规范。\n"

type SkillsService struct{}

func (s *SkillsService) Tools(_ context.Context) ([]system.SkillTool, error) {
	tools := make([]system.SkillTool, 0, len(skillToolOrder))
	for _, key := range skillToolOrder {
		if _, err := s.toolSkillsDir(key); err != nil {
			return nil, err
		}
		tools = append(tools, system.SkillTool{Key: key, Label: skillToolLabels[key]})
	}
	return tools, nil
}

func (s *SkillsService) List(_ context.Context, tool string) ([]string, error) {
	skillsDir, err := s.toolSkillsDir(tool)
	if err != nil {
		return nil, err
	}
	entries, err := os.ReadDir(skillsDir)
	if err != nil {
		return nil, err
	}
	var skills []string
	for _, entry := range entries {
		if entry.IsDir() {
			skills = append(skills, entry.Name())
		}
	}
	sort.Strings(skills)
	return skills, nil
}

func (s *SkillsService) Detail(_ context.Context, tool, skill string) (system.SkillDetail, error) {
	var detail system.SkillDetail
	if !isSafeName(skill) {
		return detail, errors.New("技能名称不合法")
	}
	detail.Tool = tool
	detail.Skill = skill

	skillDir, err := s.skillDir(tool, skill)
	if err != nil {
		return detail, err
	}

	skillFilePath := filepath.Join(skillDir, skillFileName)
	content, err := os.ReadFile(skillFilePath)
	if err != nil {
		if !os.IsNotExist(err) {
			return detail, err
		}
		detail.Meta = system.SkillMeta{Name: skill}
		detail.Markdown = defaultSkillMarkdown
	} else {
		meta, body, parseErr := parseSkillContent(string(content))
		if parseErr != nil {
			meta = system.SkillMeta{Name: skill}
			body = string(content)
		}
		if meta.Name == "" {
			meta.Name = skill
		}
		detail.Meta = meta
		detail.Markdown = body
	}

	detail.Scripts = listFiles(filepath.Join(skillDir, "scripts"))
	detail.Resources = listFiles(filepath.Join(skillDir, "resources"))
	detail.References = listFiles(filepath.Join(skillDir, "references"))
	detail.Templates = listFiles(filepath.Join(skillDir, "templates"))
	return detail, nil
}

func (s *SkillsService) Save(_ context.Context, req request.SkillSaveRequest) error {
	if !isSafeName(req.Skill) {
		return errors.New("技能名称不合法")
	}
	skillDir, err := s.ensureSkillDir(req.Tool, req.Skill)
	if err != nil {
		return err
	}
	if req.Meta.Name == "" {
		req.Meta.Name = req.Skill
	}
	content, err := buildSkillContent(req.Meta, req.Markdown)
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(skillDir, skillFileName), []byte(content), 0644); err != nil {
		return err
	}

	if len(req.SyncTools) > 0 {
		for _, tool := range req.SyncTools {
			if tool == req.Tool {
				continue
			}
			targetDir, err := s.ensureSkillDir(tool, req.Skill)
			if err != nil {
				return err
			}
			if err := copySkillDir(skillDir, targetDir); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *SkillsService) CreateScript(_ context.Context, req request.SkillScriptCreateRequest) (string, string, error) {
	if !isSafeName(req.Skill) {
		return "", "", errors.New("技能名称不合法")
	}
	fileName, lang, err := buildScriptFileName(req.FileName, req.ScriptType)
	if err != nil {
		return "", "", err
	}
	if lang == "" {
		return "", "", errors.New("脚本类型不支持")
	}
	skillDir, err := s.ensureSkillDir(req.Tool, req.Skill)
	if err != nil {
		return "", "", err
	}
	filePath := filepath.Join(skillDir, "scripts", fileName)
	if _, err := os.Stat(filePath); err == nil {
		return "", "", errors.New("脚本已存在")
	}
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return "", "", err
	}
	content := scriptTemplate(lang)
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		return "", "", err
	}
	return fileName, content, nil
}

func (s *SkillsService) GetScript(_ context.Context, req request.SkillFileRequest) (string, error) {
	return s.readSkillFile(req.Tool, req.Skill, "scripts", req.FileName)
}

func (s *SkillsService) SaveScript(_ context.Context, req request.SkillFileSaveRequest) error {
	return s.writeSkillFile(req.Tool, req.Skill, "scripts", req.FileName, req.Content)
}

func (s *SkillsService) CreateResource(_ context.Context, req request.SkillResourceCreateRequest) (string, string, error) {
	return s.createMarkdownFile(req.Tool, req.Skill, "resources", req.FileName, defaultResourceMarkdown, "资源")
}

func (s *SkillsService) GetResource(_ context.Context, req request.SkillFileRequest) (string, error) {
	return s.readSkillFile(req.Tool, req.Skill, "resources", req.FileName)
}

func (s *SkillsService) SaveResource(_ context.Context, req request.SkillFileSaveRequest) error {
	return s.writeSkillFile(req.Tool, req.Skill, "resources", req.FileName, req.Content)
}

func (s *SkillsService) CreateReference(_ context.Context, req request.SkillReferenceCreateRequest) (string, string, error) {
	return s.createMarkdownFile(req.Tool, req.Skill, "references", req.FileName, defaultReferenceMarkdown, "参考")
}

func (s *SkillsService) GetReference(_ context.Context, req request.SkillFileRequest) (string, error) {
	return s.readSkillFile(req.Tool, req.Skill, "references", req.FileName)
}

func (s *SkillsService) SaveReference(_ context.Context, req request.SkillFileSaveRequest) error {
	return s.writeSkillFile(req.Tool, req.Skill, "references", req.FileName, req.Content)
}

func (s *SkillsService) CreateTemplate(_ context.Context, req request.SkillTemplateCreateRequest) (string, string, error) {
	return s.createMarkdownFile(req.Tool, req.Skill, "templates", req.FileName, defaultTemplateMarkdown, "模板")
}

func (s *SkillsService) GetTemplate(_ context.Context, req request.SkillFileRequest) (string, error) {
	return s.readSkillFile(req.Tool, req.Skill, "templates", req.FileName)
}

func (s *SkillsService) SaveTemplate(_ context.Context, req request.SkillFileSaveRequest) error {
	return s.writeSkillFile(req.Tool, req.Skill, "templates", req.FileName, req.Content)
}

func (s *SkillsService) GetGlobalConstraint(_ context.Context, tool string) (string, bool, error) {
	skillsDir, err := s.toolSkillsDir(tool)
	if err != nil {
		return "", false, err
	}
	filePath := filepath.Join(skillsDir, globalConstraintFileName)
	content, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return defaultGlobalConstraintMarkdown, false, nil
		}
		return "", false, err
	}
	return string(content), true, nil
}

func (s *SkillsService) SaveGlobalConstraint(_ context.Context, req request.SkillGlobalConstraintSaveRequest) error {
	if strings.TrimSpace(req.Tool) == "" {
		return errors.New("工具类型不能为空")
	}
	writeConstraint := func(tool, content string) error {
		skillsDir, err := s.toolSkillsDir(tool)
		if err != nil {
			return err
		}
		filePath := filepath.Join(skillsDir, globalConstraintFileName)
		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return err
		}
		return os.WriteFile(filePath, []byte(content), 0644)
	}
	if err := writeConstraint(req.Tool, req.Content); err != nil {
		return err
	}
	if len(req.SyncTools) == 0 {
		return nil
	}
	for _, tool := range req.SyncTools {
		if tool == "" || tool == req.Tool {
			continue
		}
		if err := writeConstraint(tool, req.Content); err != nil {
			return err
		}
	}
	return nil
}

func (s *SkillsService) toolSkillsDir(tool string) (string, error) {
	toolDir, ok := skillToolDirs[tool]
	if !ok {
		return "", errors.New("工具类型不支持")
	}
	root := strings.TrimSpace(global.GVA_CONFIG.AutoCode.Root)
	if root == "" {
		root = "."
	}
	skillsDir := filepath.Join(root, toolDir, "skills")
	if err := os.MkdirAll(skillsDir, os.ModePerm); err != nil {
		return "", err
	}
	return skillsDir, nil
}

func (s *SkillsService) skillDir(tool, skill string) (string, error) {
	skillsDir, err := s.toolSkillsDir(tool)
	if err != nil {
		return "", err
	}
	return filepath.Join(skillsDir, skill), nil
}

func (s *SkillsService) ensureSkillDir(tool, skill string) (string, error) {
	if !isSafeName(skill) {
		return "", errors.New("技能名称不合法")
	}
	skillDir, err := s.skillDir(tool, skill)
	if err != nil {
		return "", err
	}
	if err := os.MkdirAll(skillDir, os.ModePerm); err != nil {
		return "", err
	}
	return skillDir, nil
}

func (s *SkillsService) createMarkdownFile(tool, skill, subDir, fileName, defaultContent, label string) (string, string, error) {
	if !isSafeName(skill) {
		return "", "", errors.New("技能名称不合法")
	}
	cleanName, err := buildResourceFileName(fileName)
	if err != nil {
		return "", "", err
	}
	skillDir, err := s.ensureSkillDir(tool, skill)
	if err != nil {
		return "", "", err
	}
	filePath := filepath.Join(skillDir, subDir, cleanName)
	if _, err := os.Stat(filePath); err == nil {
		if label == "" {
			label = "文件"
		}
		return "", "", fmt.Errorf("%s已存在", label)
	}
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return "", "", err
	}
	content := defaultContent
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		return "", "", err
	}
	return cleanName, content, nil
}

func (s *SkillsService) readSkillFile(tool, skill, subDir, fileName string) (string, error) {
	if !isSafeName(skill) {
		return "", errors.New("技能名称不合法")
	}
	if !isSafeFileName(fileName) {
		return "", errors.New("文件名不合法")
	}
	skillDir, err := s.skillDir(tool, skill)
	if err != nil {
		return "", err
	}
	filePath := filepath.Join(skillDir, subDir, fileName)
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (s *SkillsService) writeSkillFile(tool, skill, subDir, fileName, content string) error {
	if !isSafeName(skill) {
		return errors.New("技能名称不合法")
	}
	if !isSafeFileName(fileName) {
		return errors.New("文件名不合法")
	}
	skillDir, err := s.ensureSkillDir(tool, skill)
	if err != nil {
		return err
	}
	filePath := filepath.Join(skillDir, subDir, fileName)
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}
	return os.WriteFile(filePath, []byte(content), 0644)
}

func parseSkillContent(content string) (system.SkillMeta, string, error) {
	clean := strings.TrimPrefix(content, "\ufeff")
	lines := strings.Split(clean, "\n")
	if len(lines) == 0 || strings.TrimSpace(lines[0]) != "---" {
		return system.SkillMeta{}, clean, nil
	}
	end := -1
	for i := 1; i < len(lines); i++ {
		if strings.TrimSpace(lines[i]) == "---" {
			end = i
			break
		}
	}
	if end == -1 {
		return system.SkillMeta{}, clean, nil
	}
	yamlText := strings.Join(lines[1:end], "\n")
	body := strings.Join(lines[end+1:], "\n")
	var meta system.SkillMeta
	if err := yaml.Unmarshal([]byte(yamlText), &meta); err != nil {
		return system.SkillMeta{}, body, err
	}
	return meta, body, nil
}

func buildSkillContent(meta system.SkillMeta, markdown string) (string, error) {
	if meta.Name == "" {
		return "", errors.New("name不能为空")
	}
	data, err := yaml.Marshal(meta)
	if err != nil {
		return "", err
	}
	yamlText := strings.TrimRight(string(data), "\n")
	body := strings.TrimLeft(markdown, "\n")
	if body != "" {
		body = body + "\n"
	}
	return fmt.Sprintf("---\n%s\n---\n%s", yamlText, body), nil
}

func listFiles(dir string) []string {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return []string{}
	}
	files := make([]string, 0, len(entries))
	for _, entry := range entries {
		if entry.Type().IsRegular() {
			files = append(files, entry.Name())
		}
	}
	sort.Strings(files)
	return files
}

func isSafeName(name string) bool {
	if strings.TrimSpace(name) == "" {
		return false
	}
	if strings.Contains(name, "..") {
		return false
	}
	if strings.ContainsAny(name, "/\\") {
		return false
	}
	return name == filepath.Base(name)
}

func isSafeFileName(name string) bool {
	if strings.TrimSpace(name) == "" {
		return false
	}
	if strings.Contains(name, "..") {
		return false
	}
	if strings.ContainsAny(name, "/\\") {
		return false
	}
	return name == filepath.Base(name)
}

func buildScriptFileName(fileName, scriptType string) (string, string, error) {
	clean := strings.TrimSpace(fileName)
	if clean == "" {
		return "", "", errors.New("文件名不能为空")
	}
	if !isSafeFileName(clean) {
		return "", "", errors.New("文件名不合法")
	}
	base := strings.TrimSuffix(clean, filepath.Ext(clean))
	if base == "" {
		return "", "", errors.New("文件名不合法")
	}

	switch strings.ToLower(scriptType) {
	case "py", "python":
		return base + ".py", "python", nil
	case "js", "javascript", "script":
		return base + ".js", "javascript", nil
	case "sh", "shell", "bash":
		return base + ".sh", "sh", nil
	default:
		return "", "", errors.New("脚本类型不支持")
	}
}

func buildResourceFileName(fileName string) (string, error) {
	clean := strings.TrimSpace(fileName)
	if clean == "" {
		return "", errors.New("文件名不能为空")
	}
	if !isSafeFileName(clean) {
		return "", errors.New("文件名不合法")
	}
	base := strings.TrimSuffix(clean, filepath.Ext(clean))
	if base == "" {
		return "", errors.New("文件名不合法")
	}
	return base + ".md", nil
}

func scriptTemplate(lang string) string {
	switch lang {
	case "python":
		return "# -*- coding: utf-8 -*-\n# TODO: 在这里实现脚本逻辑\n"
	case "javascript":
		return "// TODO: 在这里实现脚本逻辑\n"
	case "sh":
		return "#!/usr/bin/env bash\nset -euo pipefail\n\n# TODO: 在这里实现脚本逻辑\n"
	default:
		return ""
	}
}

func copySkillDir(src, dst string) error {
	return filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		if rel == "." {
			return nil
		}
		target := filepath.Join(dst, rel)
		if d.IsDir() {
			return os.MkdirAll(target, os.ModePerm)
		}
		if !d.Type().IsRegular() {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if err := os.MkdirAll(filepath.Dir(target), os.ModePerm); err != nil {
			return err
		}
		return os.WriteFile(target, data, 0644)
	})
}
