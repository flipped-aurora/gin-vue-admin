package mcpTool

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	mcpServer "github.com/mark3labs/mcp-go/server"

	autoRes "github.com/flipped-aurora/gin-vue-admin/server/plugin/ai/model/response"
)

// scalarToString 把 MCP 入参(JSON 解码后数字均为 float64)转成用于 path/query 的字符串。
// 不能用 fmt.Sprintf("%v"):float64 大整数会被打成科学计数法(1000000→"1e+06",时间戳必中),
// 切片会被打成 "[a b]"。整数用 FormatFloat(-1 精度)输出真实数字,布尔/字符串原样。
func scalarToString(v any) string {
	switch value := v.(type) {
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(value)
	case string:
		return value
	default:
		return fmt.Sprintf("%v", value)
	}
}

// manifestCommandToTool 把 CLI 管线产出的 command 定义转成 MCP Tool（用 raw JSON schema）。
// 参数定义（SysCliManifestParameter）→ JSON Schema properties，几乎一一对应。
func manifestCommandToTool(cmd autoRes.SysCliManifestCommand) (mcp.Tool, error) {
	properties := make(map[string]any, len(cmd.Parameters))
	required := make([]string, 0, len(cmd.Parameters))
	for _, p := range cmd.Parameters {
		jsonType := manifestTypeToJSONType(p.Type)
		prop := map[string]any{
			"type":        jsonType,
			"description": p.Description,
		}
		// JSON Schema 的 array 需带 items 才完整;元素类型从 manifest 的 []xxx 推断
		// (如 []uint→integer),避免统一硬编码成 string 误导 AI 对数值ID数组发字符串
		if jsonType == "array" {
			prop["items"] = map[string]any{"type": arrayItemJSONType(p.Type)}
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
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "integer":
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

// arrayItemJSONType 从形如 []xxx 的 manifest 类型推断数组元素的 JSON Schema 类型;
// 无法识别(如泛化的 "array"、未知元素类型)时回退 string
func arrayItemJSONType(t string) string {
	trimmed := strings.TrimSpace(t)
	if strings.HasPrefix(trimmed, "[]") {
		if item := manifestTypeToJSONType(strings.TrimPrefix(trimmed, "[]")); item != "array" {
			return item
		}
	}
	return "string"
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
				// 替换路径参数 :field 和 {field}(标量化避免大整数科学计数法)
				seg := scalarToString(val)
				path = strings.ReplaceAll(path, ":"+p.Field, seg)
				path = strings.ReplaceAll(path, "{"+p.Field+"}", seg)
			case "query":
				// 数组值逐个追加成多值 query(field=a&field=b),而非被格式化成 "[a b]"
				if arr, ok := val.([]any); ok {
					for _, item := range arr {
						query.Add(p.Field, scalarToString(item))
					}
				} else {
					query.Set(p.Field, scalarToString(val))
				}
			default:
				// body / formData / header / 空 都放进 body
				body[p.Field] = val
			}
		}

		method := strings.ToUpper(cmd.Method)
		var bodyArg any
		// DELETE 也纳入可带 body 的方法:GVA 大量删除接口(deleteApisByIds 等)从 JSON body 读参数,
		// 与项目 deleteUpstream(DELETE+body)语义一致;排除它会让绑到删除接口的动态工具发空体
		if len(body) > 0 && (method == "POST" || method == "PUT" || method == "PATCH" || method == "DELETE") {
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
