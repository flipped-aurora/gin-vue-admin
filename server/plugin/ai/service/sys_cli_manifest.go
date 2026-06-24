package service

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
	autoModel "github.com/flipped-aurora/gin-vue-admin/server/plugin/ai/model"
	autoRes "github.com/flipped-aurora/gin-vue-admin/server/plugin/ai/model/response"
)

type sysCliManifestBinding struct {
	Binding autoModel.SysCliApi
	Api     sysModel.SysApi
}

type swaggerSpec struct {
	BasePath       string                            `json:"basePath"`
	Paths          map[string]swaggerPathItem        `json:"paths"`
	Definitions    map[string]swaggerDefinition      `json:"definitions"`
	RawDefinitions map[string]map[string]interface{} `json:"-"`
}

type swaggerPathItem map[string]swaggerOperation

type swaggerOperation struct {
	Summary     string                     `json:"summary"`
	Description string                     `json:"description"`
	Consumes    []string                   `json:"consumes"`
	Produces    []string                   `json:"produces"`
	Parameters  []swaggerParameter         `json:"parameters"`
	Responses   map[string]swaggerResponse `json:"responses"`
}

type swaggerParameter struct {
	Name        string                 `json:"name"`
	In          string                 `json:"in"`
	Description string                 `json:"description"`
	Required    bool                   `json:"required"`
	Type        string                 `json:"type"`
	Schema      map[string]interface{} `json:"schema"`
}

type swaggerResponse struct {
	Description string                 `json:"description"`
	Schema      map[string]interface{} `json:"schema"`
}

type swaggerDefinition struct {
	Properties map[string]swaggerProperty `json:"properties"`
}

type swaggerProperty struct {
	Type        string `json:"type"`
	Format      string `json:"format"`
	Description string `json:"description"`
	Ref         string `json:"$ref"`
}

var (
	sysCliNonAlnumRegexp    = regexp.MustCompile(`[^a-z0-9]+`)
	sysCliCamelBoundaryExpr = regexp.MustCompile(`([a-z0-9])([A-Z])`)
)

func buildSysCliManifest(cli autoModel.SysCli, bindings []sysCliManifestBinding) (autoRes.SysCliManifestResponse, error) {
	enabled := enabledSysCliManifestBindings(bindings)
	spec, err := loadSysCliSwaggerSpec()
	if err != nil {
		return autoRes.SysCliManifestResponse{}, err
	}
	applySysCliDefaults(&cli)
	commands := make([]autoRes.SysCliManifestCommand, 0, len(enabled))
	usedNames := make(map[string]int)
	for _, binding := range enabled {
		command, err := buildSysCliManifestCommand(cli, spec, binding, usedNames)
		if err != nil {
			return autoRes.SysCliManifestResponse{}, err
		}
		commands = append(commands, command)
	}
	return autoRes.SysCliManifestResponse{
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

func buildSysCliManifestServer(basePath string) autoRes.SysCliManifestServer {
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
	return autoRes.SysCliManifestServer{
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

func buildSysCliManifestCommand(cli autoModel.SysCli, spec *swaggerSpec, binding sysCliManifestBinding, usedNames map[string]int) (autoRes.SysCliManifestCommand, error) {
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
	// 用户自定义的命令说明覆盖 Swagger 派生的 summary/description
	if desc := strings.TrimSpace(binding.Binding.CommandDesc); desc != "" {
		summary = desc
		description = desc
	}
	// ApiBrief 非空时仅覆盖 summary（保留原优先级，零回归）
	if brief := strings.TrimSpace(binding.Binding.ApiBrief); brief != "" {
		summary = brief
	}
	contentType := "application/json"
	if len(operation.Consumes) > 0 && strings.TrimSpace(operation.Consumes[0]) != "" {
		contentType = strings.TrimSpace(operation.Consumes[0])
	}
	// 用户自定义的参数定义覆盖 Swagger 派生；无覆盖时回退自动派生
	var params []autoRes.SysCliManifestParameter
	if override, ok := parseParamsOverride(binding.Binding.ParamsOverride); ok && len(override) > 0 {
		params = override
	} else {
		var paramWarnings []string
		params, paramWarnings = swaggerParamsToManifest(operation.Parameters, spec)
		warnings = append(warnings, paramWarnings...)
	}
	// 返回字段：用户覆盖优先，否则从 Swagger 响应派生
	var responseFields []autoRes.SysCliManifestResponseField
	if override, ok := parseResponseOverride(binding.Binding.ResponseOverride); ok && len(override) > 0 {
		responseFields = override
	} else {
		var responseWarnings []string
		responseFields, responseWarnings = swaggerResponseToFields(operation, spec)
		warnings = append(warnings, responseWarnings...)
	}
	fullPath := normalizeManifestPath(spec.BasePath, resolvedPath)
	examples := []string{buildSysCliExample(cli, name, params)}
	return autoRes.SysCliManifestCommand{
		Name:        name,
		Summary:     summary,
		Description: description,
		Method:      strings.ToUpper(strings.TrimSpace(binding.Api.Method)),
		Path:        fullPath,
		ContentType: contentType,
		Source: autoRes.SysCliManifestSource{
			ApiID:    binding.Binding.ApiID,
			ApiGroup: binding.Api.ApiGroup,
			From:     "mcp+swagger",
		},
		Parameters: params,
		Response:   responseFields,
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

func swaggerParamsToManifest(params []swaggerParameter, spec *swaggerSpec) ([]autoRes.SysCliManifestParameter, []string) {
	result := make([]autoRes.SysCliManifestParameter, 0, len(params))
	warnings := make([]string, 0)
	for _, param := range params {
		if strings.EqualFold(param.Name, "x-token") {
			continue
		}
		if expanded := expandBodySchemaParameters(param, spec); len(expanded) > 0 {
			result = append(result, expanded...)
			continue
		}
		typeName := detectSwaggerType(param)
		result = append(result, autoRes.SysCliManifestParameter{
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

// expandBodySchemaParameters 把 body/formData 中引用模型的参数按模型字段展开为多个 CLI 入参，
// 让 --help 能看到具体字段而非一个笼统的 JSON 入参。无法展开时返回 nil。
func expandBodySchemaParameters(param swaggerParameter, spec *swaggerSpec) []autoRes.SysCliManifestParameter {
	if spec == nil || len(param.Schema) == 0 {
		return nil
	}
	defName := resolveSchemaRef(param.Schema)
	if defName == "" {
		return nil
	}
	def, ok := spec.Definitions[defName]
	if !ok || len(def.Properties) == 0 {
		return nil
	}
	propNames := make([]string, 0, len(def.Properties))
	for name := range def.Properties {
		propNames = append(propNames, name)
	}
	sort.Strings(propNames)
	result := make([]autoRes.SysCliManifestParameter, 0, len(propNames))
	for _, name := range propNames {
		prop := def.Properties[name]
		result = append(result, autoRes.SysCliManifestParameter{
			Name:        name,
			Flag:        sanitizeSingleSegmentSlug(name),
			Type:        detectPropertyType(prop),
			Description: strings.TrimSpace(prop.Description),
			Location:    param.In,
			Field:       name,
		})
	}
	return result
}

// swaggerResponseToFields 从操作的 2xx 响应解析返回字段：
// GVA 响应为 allOf[response.Response 信封, {data: $ref|object|array}]，
// 只展开内层 data 一层。无法解析时返回空。
func swaggerResponseToFields(operation swaggerOperation, spec *swaggerSpec) ([]autoRes.SysCliManifestResponseField, []string) {
	if spec == nil {
		return nil, nil
	}
	resp := pickSysCli2xxResponse(operation.Responses)
	if resp == nil || len(resp.Schema) == 0 {
		return nil, nil
	}
	dataSchema := extractSysCliResponseData(resp.Schema)
	if dataSchema == nil {
		return nil, nil
	}
	return expandSysCliResponseFields(dataSchema, spec)
}

func pickSysCli2xxResponse(responses map[string]swaggerResponse) *swaggerResponse {
	codes := make([]string, 0, len(responses))
	for code := range responses {
		if len(code) == 3 && code[0] == '2' {
			codes = append(codes, code)
		}
	}
	sort.Strings(codes)
	if len(codes) == 0 {
		return nil
	}
	r := responses[codes[0]]
	return &r
}

// extractSysCliResponseData 从响应 schema 中定位 data 属性的 schema，
// 支持 allOf（GVA 信封）与顶层 properties.data。
func extractSysCliResponseData(schema map[string]interface{}) map[string]interface{} {
	candidates := []map[string]interface{}{schema}
	if allOf, ok := schema["allOf"].([]interface{}); ok {
		for _, raw := range allOf {
			if obj, ok := raw.(map[string]interface{}); ok {
				candidates = append(candidates, obj)
			}
		}
	}
	for _, c := range candidates {
		props, ok := c["properties"].(map[string]interface{})
		if !ok {
			continue
		}
		if data, ok := props["data"].(map[string]interface{}); ok {
			return data
		}
	}
	return nil
}

const sysCliResponseMaxDepth = 3

// expandSysCliResponseFields 把 data schema 递归展开为点路径返回字段（最大 3 段，数组下钻一层）。
func expandSysCliResponseFields(dataSchema map[string]interface{}, spec *swaggerSpec) ([]autoRes.SysCliManifestResponseField, []string) {
	return expandResponseNode(dataSchema, spec, "", 1), nil
}

// expandResponseNode 展开单个 schema 节点的子属性；无子属性（标量/空）返回 nil。
func expandResponseNode(schema map[string]interface{}, spec *swaggerSpec, prefix string, depth int) []autoRes.SysCliManifestResponseField {
	props := resolveResponseProps(schema, spec)
	if len(props) == 0 {
		return nil
	}
	return expandResponseProperties(props, spec, prefix, depth)
}

// expandResponseProperties 生成属性级字段：path 段数 == depth；depth < maxDepth 时对子对象/数组下钻。
func expandResponseProperties(props map[string]interface{}, spec *swaggerSpec, prefix string, depth int) []autoRes.SysCliManifestResponseField {
	if depth > sysCliResponseMaxDepth {
		return nil
	}
	names := make([]string, 0, len(props))
	for name := range props {
		names = append(names, name)
	}
	sort.Strings(names)
	fields := make([]autoRes.SysCliManifestResponseField, 0, len(names))
	for _, name := range names {
		propMap, _ := props[name].(map[string]interface{})
		path := joinResponsePath(prefix, name)
		if depth < sysCliResponseMaxDepth {
			if child := expandResponseNode(propMap, spec, path, depth+1); len(child) > 0 {
				fields = append(fields, child...)
				continue
			}
		}
		fields = append(fields, autoRes.SysCliManifestResponseField{
			Name:        path,
			Description: strings.TrimSpace(stringFromMap(propMap, "description")),
		})
	}
	return fields
}

// resolveResponseProps 解析一个 schema 节点的属性集合：
// $ref → RawDefinitions 的 properties；内联 object → 其 properties；array → items 的属性（递归一层）。
func resolveResponseProps(schema map[string]interface{}, spec *swaggerSpec) map[string]interface{} {
	if len(schema) == 0 {
		return nil
	}
	if name := resolveSchemaRef(schema); name != "" {
		def, ok := spec.RawDefinitions[name]
		if !ok {
			return nil
		}
		return propsOf(def)
	}
	if props := propsOf(schema); len(props) > 0 {
		return props
	}
	if t, _ := schema["type"].(string); t == "array" {
		items, _ := schema["items"].(map[string]interface{})
		if len(items) == 0 {
			return nil
		}
		// 数组只下钻一层：取 items 的属性（内联 properties 或 $ref 定义），不再处理 items 仍为数组的情况。
		if props := propsOf(items); len(props) > 0 {
			return props
		}
		if name := resolveSchemaRef(items); name != "" {
			if def, ok := spec.RawDefinitions[name]; ok {
				return propsOf(def)
			}
		}
		return nil
	}
	return nil
}

func propsOf(m map[string]interface{}) map[string]interface{} {
	if len(m) == 0 {
		return nil
	}
	props, _ := m["properties"].(map[string]interface{})
	return props
}

func joinResponsePath(prefix, name string) string {
	if prefix == "" {
		return name
	}
	return prefix + "." + name
}

func stringFromMap(m map[string]interface{}, key string) string {
	if v, ok := m[key].(string); ok {
		return v
	}
	return ""
}

func resolveSchemaRef(schema map[string]interface{}) string {
	if len(schema) == 0 {
		return ""
	}
	ref, ok := schema["$ref"].(string)
	if !ok || strings.TrimSpace(ref) == "" {
		return ""
	}
	ref = strings.TrimSpace(ref)
	if idx := strings.LastIndex(ref, "/"); idx >= 0 {
		return ref[idx+1:]
	}
	return ref
}

func detectPropertyType(prop swaggerProperty) string {
	if t := strings.TrimSpace(prop.Type); t != "" {
		return t
	}
	if strings.TrimSpace(prop.Ref) != "" {
		return "json"
	}
	return "string"
}

// parseParamsOverride 解析用户自定义的参数定义 JSON，空或非法时返回 false 以回退自动派生。
func parseParamsOverride(raw string) ([]autoRes.SysCliManifestParameter, bool) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil, false
	}
	var params []autoRes.SysCliManifestParameter
	if err := json.Unmarshal([]byte(raw), &params); err != nil {
		return nil, false
	}
	return params, true
}

// parseResponseOverride 解析用户自定义的返回字段 JSON，空或非法时返回 false 以回退自动派生。
func parseResponseOverride(raw string) ([]autoRes.SysCliManifestResponseField, bool) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil, false
	}
	var fields []autoRes.SysCliManifestResponseField
	if err := json.Unmarshal([]byte(raw), &fields); err != nil {
		return nil, false
	}
	return fields, true
}

func buildSysCliExample(cli autoModel.SysCli, commandName string, params []autoRes.SysCliManifestParameter) string {
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

func sysCliManifestFileName(cli autoModel.SysCli) string {
	applySysCliDefaults(&cli)
	name := sanitizeSingleSegmentSlug(cli.Command)
	if name == "" {
		name = fmt.Sprintf("cli-%d", cli.ID)
	}
	return name + ".manifest.json"
}

func marshalSysCliManifest(manifest autoRes.SysCliManifestResponse) ([]byte, error) {
	return json.MarshalIndent(manifest, "", "  ")
}

func loadSysCliSwaggerSpec() (*swaggerSpec, error) {
	doc := docs.SwaggerInfo.ReadDoc()
	var spec swaggerSpec
	if err := json.Unmarshal([]byte(doc), &spec); err != nil {
		return nil, err
	}
	var rawDefs struct {
		Definitions map[string]map[string]interface{} `json:"definitions"`
	}
	if err := json.Unmarshal([]byte(doc), &rawDefs); err != nil {
		return nil, err
	}
	spec.RawDefinitions = rawDefs.Definitions
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
func applyCliBuildBaseURL(manifest *autoRes.SysCliManifestResponse, baseURL string) {
	baseURL = strings.TrimSpace(baseURL)
	if baseURL == "" {
		return
	}
	manifest.Server.BaseURL = strings.TrimRight(baseURL, "/")
}

var _ = config.MCP{}
