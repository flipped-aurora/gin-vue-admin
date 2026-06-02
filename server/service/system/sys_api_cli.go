package system

import (
	"bytes"
	"encoding/json"
	"fmt"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/flipped-aurora/gin-vue-admin/server/docs"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/autocode"
)

type apiCliService struct{}

type swaggerSpec struct {
	BasePath string                       `json:"basePath"`
	Paths    map[string]swaggerPathItem   `json:"paths"`
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

type bashTemplateData struct {
	ScriptName      string
	Method          string
	Path            string
	Summary         string
	Description     string
	Example         string
	BaseUrlEnvName  string
	TokenEnvName    string
	PathParams      []bashParam
	QueryParams     []bashParam
	HeaderParams    []bashParam
	Warnings        []string
}

type bashParam struct {
	Name        string
	FlagName    string
	VarName     string
	Required    bool
	Description string
}

var ApiCliService = new(apiCliService)

func (s *apiCliService) Preview(info systemReq.GenerateApiCliRequest) (systemRes.ApiCliPreviewResponse, error) {
	resolved, err := s.resolve(info)
	if err != nil {
		return systemRes.ApiCliPreviewResponse{}, err
	}
	return systemRes.ApiCliPreviewResponse{
		ScriptName:     resolved.ScriptName,
		ScriptContent:  resolved.ScriptContent,
		Summary:        resolved.Summary,
		Description:    resolved.Description,
		Warnings:       resolved.Warnings,
		DetectedParams: resolved.DetectedParams,
	}, nil
}

func (s *apiCliService) Download(info systemReq.GenerateApiCliRequest) (string, []byte, error) {
	resolved, err := s.resolve(info)
	if err != nil {
		return "", nil, err
	}
	return resolved.ScriptName, []byte(resolved.ScriptContent), nil
}

type resolvedCli struct {
	ScriptName     string
	ScriptContent  string
	Summary        string
	Description    string
	Warnings       []string
	DetectedParams systemRes.ApiCliDetectedParams
}

func (s *apiCliService) resolve(info systemReq.GenerateApiCliRequest) (resolvedCli, error) {
	if strings.TrimSpace(info.Path) == "" || strings.TrimSpace(info.Method) == "" {
		return resolvedCli{}, fmt.Errorf("path和method不能为空")
	}
	spec, err := s.loadSwaggerSpec()
	if err != nil {
		return resolvedCli{}, err
	}
	operation, resolvedPath, err := s.resolveOperation(spec, info.Path, info.Method)
	if err != nil {
		return resolvedCli{}, err
	}
	data, detected := s.buildTemplateData(spec.BasePath, resolvedPath, info, operation)
	script, err := s.renderScript(data)
	if err != nil {
		return resolvedCli{}, err
	}
	return resolvedCli{
		ScriptName:     data.ScriptName,
		ScriptContent:  script,
		Summary:        data.Summary,
		Description:    data.Description,
		Warnings:       data.Warnings,
		DetectedParams: detected,
	}, nil
}

func (s *apiCliService) loadSwaggerSpec() (*swaggerSpec, error) {
	doc := docs.SwaggerInfo.ReadDoc()
	var spec swaggerSpec
	if err := json.Unmarshal([]byte(doc), &spec); err != nil {
		return nil, fmt.Errorf("解析Swagger文档失败: %w", err)
	}
	return &spec, nil
}

func (s *apiCliService) resolveOperation(spec *swaggerSpec, path string, method string) (swaggerOperation, string, error) {
	item, ok := spec.Paths[path]
	if !ok {
		return swaggerOperation{}, "", fmt.Errorf("未找到可生成的Swagger定义")
	}
	operation, ok := item[strings.ToLower(method)]
	if !ok {
		return swaggerOperation{}, "", fmt.Errorf("未找到可生成的Swagger定义")
	}
	return operation, path, nil
}

func (s *apiCliService) buildTemplateData(basePath string, path string, info systemReq.GenerateApiCliRequest, operation swaggerOperation) (bashTemplateData, systemRes.ApiCliDetectedParams) {
	fullPath := normalizeSwaggerPath(basePath, path)
	warnings := make([]string, 0)
	data := bashTemplateData{
		Method:         strings.ToUpper(info.Method),
		Path:           fullPath,
		Summary:        fallbackText(operation.Summary, "Generated API shell wrapper"),
		Description:    operation.Description,
		BaseUrlEnvName: fallbackText(info.BaseUrlEnvName, "GVA_BASE_URL"),
		TokenEnvName:   fallbackText(info.TokenEnvName, "GVA_X_TOKEN"),
		Warnings:       warnings,
	}
	detected := systemRes.ApiCliDetectedParams{}
	for _, param := range operation.Parameters {
		typeName := param.Type
		if typeName == "" && len(param.Schema) > 0 {
			typeName = "object"
		}
		detectedParam := systemRes.ApiCliDetectedParam{
			Name:        param.Name,
			In:          param.In,
			Type:        typeName,
			Required:    param.Required,
			Description: param.Description,
		}
		safeName := toFlagName(param.Name)
		bashParam := bashParam{
			Name:        param.Name,
			FlagName:    safeName,
			VarName:     toVarName(param.Name),
			Required:    param.Required,
			Description: fallbackText(param.Description, param.Name),
		}
		switch param.In {
		case "path":
			data.PathParams = append(data.PathParams, bashParam)
			detected.Path = append(detected.Path, detectedParam)
		case "query":
			data.QueryParams = append(data.QueryParams, bashParam)
			detected.Query = append(detected.Query, detectedParam)
		case "header":
			if strings.EqualFold(param.Name, "x-token") {
				continue
			}
			data.HeaderParams = append(data.HeaderParams, bashParam)
			detected.Header = append(detected.Header, detectedParam)
		case "body", "formData":
			detected.HasBody = true
			warnings = append(warnings, "请求体通过 --data 或 --data-file 透传，未自动拆分字段。")
		default:
			warnings = append(warnings, fmt.Sprintf("参数 %s 的位置 %s 暂未专门处理，可能需要手动调整脚本。", param.Name, param.In))
		}
	}
	data.Warnings = uniqueStrings(warnings)
	sort.Slice(data.QueryParams, func(i, j int) bool { return data.QueryParams[i].FlagName < data.QueryParams[j].FlagName })
	sort.Slice(data.HeaderParams, func(i, j int) bool { return data.HeaderParams[i].FlagName < data.HeaderParams[j].FlagName })
	data.ScriptName = s.buildScriptName(info.WrapperName, info.Method, path)
	data.Example = s.buildExample(data)
	return data, detected
}

func (s *apiCliService) renderScript(data bashTemplateData) (string, error) {
	templatePath := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "resource", "api_cli", "bash.tpl")
	files, err := template.New(filepath.Base(templatePath)).Funcs(autocode.GetTemplateFuncMap()).ParseFiles(templatePath)
	if err != nil {
		return "", err
	}
	var builder bytes.Buffer
	if err = files.Execute(&builder, data); err != nil {
		return "", err
	}
	return builder.String(), nil
}

func (s *apiCliService) buildScriptName(wrapperName string, method string, path string) string {
	name := strings.TrimSpace(wrapperName)
	if name == "" {
		name = strings.ToLower(method) + "_" + path
	}
	name = slugPattern.ReplaceAllString(strings.ToLower(name), "-")
	name = strings.Trim(name, "-")
	if name == "" {
		name = "api-cli"
	}
	return name + ".sh"
}

func (s *apiCliService) buildExample(data bashTemplateData) string {
	parts := []string{"./" + data.ScriptName}
	for _, param := range data.PathParams {
		parts = append(parts, "--"+param.FlagName, sampleValue(param.Name))
	}
	for _, param := range data.QueryParams {
		if param.Required {
			parts = append(parts, "--"+param.FlagName, sampleValue(param.Name))
		}
	}
	if len(data.PathParams) == 0 && !hasRequiredQuery(data.QueryParams) {
		parts = append(parts, "--help")
	}
	return strings.Join(parts, " ")
}

var slugPattern = regexp.MustCompile(`[^a-z0-9]+`)

func normalizeSwaggerPath(basePath string, path string) string {
	base := strings.TrimRight(basePath, "/")
	if base == "" || base == "/" {
		return path
	}
	return base + path
}

func fallbackText(value string, fallback string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return fallback
	}
	return value
}

func toFlagName(name string) string {
	name = strings.ReplaceAll(name, "_", "-")
	name = strings.ReplaceAll(name, ".", "-")
	name = slugPattern.ReplaceAllString(strings.ToLower(name), "-")
	return strings.Trim(name, "-")
}

func toVarName(name string) string {
	cleaned := strings.ToUpper(strings.ReplaceAll(toFlagName(name), "-", "_"))
	if cleaned == "" {
		return "PARAM"
	}
	return cleaned
}

func sampleValue(name string) string {
	switch strings.ToLower(name) {
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

func hasRequiredQuery(params []bashParam) bool {
	for _, param := range params {
		if param.Required {
			return true
		}
	}
	return false
}

func uniqueStrings(items []string) []string {
	seen := make(map[string]struct{}, len(items))
	result := make([]string, 0, len(items))
	for _, item := range items {
		if strings.TrimSpace(item) == "" {
			continue
		}
		if _, ok := seen[item]; ok {
			continue
		}
		seen[item] = struct{}{}
		result = append(result, item)
	}
	return result
}
