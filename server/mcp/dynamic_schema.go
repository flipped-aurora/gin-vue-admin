package mcpTool

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	mcpServer "github.com/mark3labs/mcp-go/server"

	autoRes "github.com/flipped-aurora/gin-vue-admin/server/plugin/ai/model/response"
)

// manifestCommandToTool 把 CLI 管线产出的 command 定义转成 MCP Tool（用 raw JSON schema）。
// 参数定义（SysCliManifestParameter）→ JSON Schema properties，几乎一一对应。
func manifestCommandToTool(cmd autoRes.SysCliManifestCommand) (mcp.Tool, error) {
	properties := make(map[string]any, len(cmd.Parameters))
	required := make([]string, 0, len(cmd.Parameters))
	for _, p := range cmd.Parameters {
		prop := map[string]any{
			"type":        manifestTypeToJSONType(p.Type),
			"description": p.Description,
		}
		properties[p.Name] = prop
		if p.Required {
			required = append(required, p.Name)
		}
	}

	schema := map[string]any{
		"type":       "object",
		"properties": properties,
	}
	if len(required) > 0 {
		schema["required"] = required
	}

	rawSchema, err := json.Marshal(schema)
	if err != nil {
		return mcp.Tool{}, fmt.Errorf("序列化 tool schema 失败: %w", err)
	}

	description := cmd.Description
	if description == "" {
		description = cmd.Summary
	}
	if description == "" {
		description = fmt.Sprintf("%s %s", strings.ToUpper(cmd.Method), cmd.Path)
	}

	tool := mcp.NewToolWithRawSchema(cmd.Name, description, rawSchema)
	return tool, nil
}

// manifestTypeToJSONType 把 CLI manifest 的参数类型映射到 JSON Schema 类型。
func manifestTypeToJSONType(t string) string {
	switch strings.ToLower(strings.TrimSpace(t)) {
	case "int", "int8", "int16", "int32", "int64", "uint", "integer":
		return "integer"
	case "bool", "boolean":
		return "boolean"
	case "float", "float32", "float64", "number":
		return "number"
	case "[]string", "array":
		return "array"
	default:
		return "string"
	}
}

// makeDynamicHandler 为一个 command 定义创建 MCP tool handler。
// 执行时：从 MCP 调用参数取值 → 按 parameter.Location 组 path/query/body → 调 GVA。
func makeDynamicHandler(cmd autoRes.SysCliManifestCommand) mcpServer.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, _ := request.Params.Arguments.(map[string]any)
		if args == nil {
			args = map[string]any{}
		}

		path := cmd.Path
		query := url.Values{}
		body := map[string]any{}

		for _, p := range cmd.Parameters {
			val, ok := args[p.Name]
			if !ok {
				continue
			}
			switch strings.ToLower(p.Location) {
			case "path":
				// 替换路径参数 :field 和 {field}
				path = strings.ReplaceAll(path, ":"+p.Field, fmt.Sprintf("%v", val))
				path = strings.ReplaceAll(path, "{"+p.Field+"}", fmt.Sprintf("%v", val))
			case "query":
				query.Set(p.Field, fmt.Sprintf("%v", val))
			default:
				// body / formData / header / 空 都放进 body
				body[p.Field] = val
			}
		}

		method := strings.ToUpper(cmd.Method)
		var bodyArg any
		if len(body) > 0 && (method == "POST" || method == "PUT" || method == "PATCH") {
			bodyArg = body
		}

		statusCode, rawBody, err := doUpstreamRaw(ctx, method, path, query, bodyArg)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("调用上游失败: %v", err)), nil
		}
		if statusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("上游返回错误状态码 %d: %s", statusCode, string(rawBody))), nil
		}

		return mcp.NewToolResultText(string(rawBody)), nil
	}
}
