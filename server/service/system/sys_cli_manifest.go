package system

import (
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/docs"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
)

type sysCliManifestBinding struct {
	Binding sysModel.SysCliApi
	Api     sysModel.SysApi
}

type swaggerSpec struct {
	BasePath string                     `json:"basePath"`
	Paths    map[string]swaggerPathItem `json:"paths"`
}

type swaggerPathItem map[string]swaggerOperation

type swaggerOperation struct {
	Summary     string             `json:"summary"`
	Description string             `json:"description"`
	Consumes    []string           `json:"consumes"`
	Produces    []string           `json:"produces"`
	Parameters  []swaggerParameter `json:"parameters"`
}

type swaggerParameter struct {
	Name        string                 `json:"name"`
	In          string                 `json:"in"`
	Description string                 `json:"description"`
	Required    bool                   `json:"required"`
	Type        string                 `json:"type"`
	Schema      map[string]interface{} `json:"schema"`
}

var (
	sysCliNonAlnumRegexp    = regexp.MustCompile(`[^a-z0-9]+`)
	sysCliCamelBoundaryExpr = regexp.MustCompile(`([a-z0-9])([A-Z])`)
)

func buildSysCliManifest(cli sysModel.SysCli, bindings []sysCliManifestBinding) (systemRes.SysCliManifestResponse, error) {
	enabled := enabledSysCliManifestBindings(bindings)
	spec, err := loadSysCliSwaggerSpec()
	if err != nil {
		return systemRes.SysCliManifestResponse{}, err
	}
	applySysCliDefaults(&cli)
	commands := make([]systemRes.SysCliManifestCommand, 0, len(enabled))
	usedNames := make(map[string]int)
	for _, binding := range enabled {
		command, err := buildSysCliManifestCommand(cli, spec, binding, usedNames)
		if err != nil {
			return systemRes.SysCliManifestResponse{}, err
		}
		commands = append(commands, command)
	}
	return systemRes.SysCliManifestResponse{
		Name:     strings.TrimSpace(cli.Command),
		Version:  strings.TrimSpace(cli.Version),
		Server:   buildSysCliManifestServer(spec.BasePath),
		Commands: commands,
	}, nil
}

func enabledSysCliManifestBindings(bindings []sysCliManifestBinding) []sysCliManifestBinding {
	enabled := make([]sysCliManifestBinding, 0, len(bindings))
	for _, binding := range bindings {
		if !binding.Binding.Enabled {
			continue
		}
		enabled = append(enabled, binding)
	}
	sort.SliceStable(enabled, func(i, j int) bool {
		if enabled[i].Binding.Sort != enabled[j].Binding.Sort {
			return enabled[i].Binding.Sort < enabled[j].Binding.Sort
		}
		if enabled[i].Binding.ID != enabled[j].Binding.ID {
			return enabled[i].Binding.ID < enabled[j].Binding.ID
		}
		return enabled[i].Binding.ApiID < enabled[j].Binding.ApiID
	})
	return enabled
}

func buildSysCliManifestServer(basePath string) systemRes.SysCliManifestServer {
	authHeader := strings.TrimSpace(global.GVA_CONFIG.MCP.AuthHeader)
	if authHeader == "" {
		authHeader = "x-token"
	}
	baseURL := strings.TrimSpace(global.GVA_CONFIG.MCP.UpstreamBaseURL)
	if baseURL == "" {
		baseURL = strings.TrimSpace(global.GVA_CONFIG.MCP.BaseURL)
	}
	if baseURL == "" {
		baseURL = buildDefaultServerBaseURL(basePath)
	}
	return systemRes.SysCliManifestServer{
		BaseURL:    strings.TrimRight(baseURL, "/"),
		AuthHeader: authHeader,
	}
}

func buildDefaultServerBaseURL(basePath string) string {
	prefix := strings.TrimSpace(global.GVA_CONFIG.System.RouterPrefix)
	prefix = strings.Trim(prefix, "/")
	path := ""
	if prefix != "" {
		path += "/" + prefix
	}
	if strings.TrimSpace(basePath) != "" {
		path += normalizeManifestPath("", basePath)
	}
	if path == "" {
		return fmt.Sprintf("http://127.0.0.1:%d", global.GVA_CONFIG.System.Addr)
	}
	return fmt.Sprintf("http://127.0.0.1:%d%s", global.GVA_CONFIG.System.Addr, path)
}

func buildSysCliManifestCommand(cli sysModel.SysCli, spec *swaggerSpec, binding sysCliManifestBinding, usedNames map[string]int) (systemRes.SysCliManifestCommand, error) {
	operation, resolvedPath, warnings := resolveSwaggerOperation(spec, binding.Api.Path, binding.Api.Method)
	name := uniqueSysCliCommandName(sysCliCommandName(binding), usedNames)
	summary := strings.TrimSpace(binding.Api.Description)
	if strings.TrimSpace(operation.Summary) != "" {
		summary = strings.TrimSpace(operation.Summary)
	}
	description := strings.TrimSpace(operation.Description)
	if description == "" {
		description = summary
	}
	contentType := "application/json"
	if len(operation.Consumes) > 0 && strings.TrimSpace(operation.Consumes[0]) != "" {
		contentType = strings.TrimSpace(operation.Consumes[0])
	}
	params, paramWarnings := swaggerParamsToManifest(operation.Parameters)
	warnings = append(warnings, paramWarnings...)
	fullPath := normalizeManifestPath(spec.BasePath, resolvedPath)
	examples := []string{buildSysCliExample(cli, name, params)}
	return systemRes.SysCliManifestCommand{
		Name:        name,
		Summary:     summary,
		Description: description,
		Method:      strings.ToUpper(strings.TrimSpace(binding.Api.Method)),
		Path:        fullPath,
		ContentType: contentType,
		Source: systemRes.SysCliManifestSource{
			ApiID:    binding.Binding.ApiID,
			ApiGroup: binding.Api.ApiGroup,
			From:     "mcp+swagger",
		},
		Parameters: params,
		Examples:   examples,
		Warnings:   uniqueStrings(warnings),
	}, nil
}

func resolveSwaggerOperation(spec *swaggerSpec, path string, method string) (swaggerOperation, string, []string) {
	if spec == nil {
		return swaggerOperation{}, path, []string{"Swagger 文档不可用，参数定义为空。"}
	}
	if item, ok := spec.Paths[path]; ok {
		if op, exists := item[strings.ToLower(method)]; exists {
			return op, path, nil
		}
	}
	return swaggerOperation{}, path, []string{"未找到对应 Swagger 定义，参数定义为空。"}
}

func uniqueSysCliCommandName(name string, used map[string]int) string {
	base := sanitizeSingleSegmentSlug(name)
	if base == "" {
		base = "command"
	}
	count := used[base]
	used[base] = count + 1
	if count == 0 {
		return base
	}
	return fmt.Sprintf("%s-%d", base, count+1)
}

func sysCliCommandName(binding sysCliManifestBinding) string {
	if override := sanitizeSingleSegmentSlug(binding.Binding.CommandName); override != "" {
		return override
	}
	return deriveSingleSegmentSlugFromAPI(binding.Api)
}

func deriveSingleSegmentSlugFromAPI(api sysModel.SysApi) string {
	segments := strings.Split(strings.Trim(strings.TrimSpace(api.Path), "/"), "/")
	if len(segments) >= 2 {
		resource := sanitizeSingleSegmentSlug(segments[0])
		actionRaw := trimAPIVerbPrefix(segments[len(segments)-1])
		action := sanitizeSingleSegmentSlug(actionRaw)
		if resource != "" && action != "" {
			if action == resource {
				return resource
			}
			if strings.HasPrefix(action, resource+"-") {
				action = strings.TrimPrefix(action, resource+"-")
			}
			if action != "" {
				return resource + "-" + action
			}
		}
	}
	for i := len(segments) - 1; i >= 0; i-- {
		segment := trimAPIVerbPrefix(segments[i])
		if slug := sanitizeSingleSegmentSlug(segment); slug != "" {
			return slug
		}
	}
	fallback := sanitizeSingleSegmentSlug(strings.ToLower(strings.TrimSpace(api.Method)))
	if fallback == "" {
		fallback = "command"
	}
	if api.ID > 0 {
		return fmt.Sprintf("%s-%d", fallback, api.ID)
	}
	return fallback
}

func trimAPIVerbPrefix(value string) string {
	prefixes := []string{"get", "create", "update", "delete", "set", "find", "list"}
	trimmed := strings.TrimSpace(value)
	lower := strings.ToLower(trimmed)
	for _, prefix := range prefixes {
		if strings.HasPrefix(lower, prefix) && len(trimmed) > len(prefix) {
			return trimmed[len(prefix):]
		}
	}
	return trimmed
}

func sanitizeSingleSegmentSlug(input string) string {
	value := strings.TrimSpace(input)
	if value == "" {
		return ""
	}
	value = strings.TrimPrefix(value, ":")
	value = strings.Trim(value, "{}")
	value = sysCliCamelBoundaryExpr.ReplaceAllString(value, `${1}-${2}`)
	value = strings.ToLower(value)
	value = strings.ReplaceAll(value, "_", "-")
	value = sysCliNonAlnumRegexp.ReplaceAllString(value, "-")
	value = strings.Trim(value, "-")
	return value
}

func normalizeManifestPath(basePath string, path string) string {
	base := strings.TrimSpace(basePath)
	body := strings.TrimSpace(path)
	if base == "" {
		if body == "" {
			return "/"
		}
		if strings.HasPrefix(body, "/") {
			return body
		}
		return "/" + body
	}
	base = "/" + strings.Trim(base, "/")
	if body == "" || body == "/" {
		return base
	}
	body = "/" + strings.Trim(body, "/")
	return base + body
}

func swaggerParamsToManifest(params []swaggerParameter) ([]systemRes.SysCliManifestParameter, []string) {
	result := make([]systemRes.SysCliManifestParameter, 0, len(params))
	warnings := make([]string, 0)
	for _, param := range params {
		if strings.EqualFold(param.Name, "x-token") {
			continue
		}
		typeName := detectSwaggerType(param)
		result = append(result, systemRes.SysCliManifestParameter{
			Name:        param.Name,
			Flag:        sanitizeSingleSegmentSlug(param.Name),
			Type:        typeName,
			Required:    param.Required,
			Description: strings.TrimSpace(param.Description),
			Location:    param.In,
			Field:       param.Name,
		})
		if param.In == "body" || param.In == "formData" {
			warnings = append(warnings, fmt.Sprintf("参数 %s 使用 %s 方式传入，复杂结构需按 JSON 透传。", param.Name, param.In))
		}
	}
	return result, warnings
}

func detectSwaggerType(param swaggerParameter) string {
	if strings.TrimSpace(param.Type) != "" {
		return strings.TrimSpace(param.Type)
	}
	if _, ok := param.Schema["$ref"]; ok {
		return "json"
	}
	if len(param.Schema) > 0 {
		if t, ok := param.Schema["type"].(string); ok && strings.TrimSpace(t) != "" {
			return strings.TrimSpace(t)
		}
		return "json"
	}
	return "string"
}

func buildSysCliExample(cli sysModel.SysCli, commandName string, params []systemRes.SysCliManifestParameter) string {
	applySysCliDefaults(&cli)
	parts := []string{strings.TrimSpace(cli.Command), commandName}
	for _, param := range params {
		if !param.Required {
			continue
		}
		parts = append(parts, "--"+param.Flag, sampleManifestValue(param.Flag))
	}
	if len(parts) == 2 {
		parts = append(parts, "--help")
	}
	return strings.Join(parts, " ")
}

func sampleManifestValue(flag string) string {
	switch strings.ToLower(flag) {
	case "id":
		return "123"
	case "page":
		return "1"
	case "pagesize", "page-size":
		return "10"
	default:
		return "value"
	}
}

func sysCliManifestFileName(cli sysModel.SysCli) string {
	applySysCliDefaults(&cli)
	name := sanitizeSingleSegmentSlug(cli.Command)
	if name == "" {
		name = fmt.Sprintf("cli-%d", cli.ID)
	}
	return name + ".manifest.json"
}

func marshalSysCliManifest(manifest systemRes.SysCliManifestResponse) ([]byte, error) {
	return json.MarshalIndent(manifest, "", "  ")
}

func loadSysCliSwaggerSpec() (*swaggerSpec, error) {
	doc := docs.SwaggerInfo.ReadDoc()
	var spec swaggerSpec
	if err := json.Unmarshal([]byte(doc), &spec); err != nil {
		return nil, err
	}
	if spec.BasePath == "" {
		spec.BasePath = strings.TrimSpace(global.GVA_CONFIG.System.RouterPrefix)
		if spec.BasePath != "" && !strings.HasPrefix(spec.BasePath, "/") {
			spec.BasePath = "/" + spec.BasePath
		}
	}
	return &spec, nil
}

func uniqueStrings(items []string) []string {
	seen := make(map[string]struct{}, len(items))
	result := make([]string, 0, len(items))
	for _, item := range items {
		trimmed := strings.TrimSpace(item)
		if trimmed == "" {
			continue
		}
		if _, ok := seen[trimmed]; ok {
			continue
		}
		seen[trimmed] = struct{}{}
		result = append(result, trimmed)
	}
	return result
}

// applyCliBuildBaseURL 用编译时指定的 baseURL 覆盖 manifest 的服务地址，空值保留原值。
func applyCliBuildBaseURL(manifest *systemRes.SysCliManifestResponse, baseURL string) {
	baseURL = strings.TrimSpace(baseURL)
	if baseURL == "" {
		return
	}
	manifest.Server.BaseURL = strings.TrimRight(baseURL, "/")
}

var _ = config.MCP{}
