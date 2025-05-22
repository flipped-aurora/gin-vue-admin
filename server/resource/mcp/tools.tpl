package mcpTool

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

func init() {
	RegisterTool(&{{.Name | title}}{})
}

type {{.Name | title}} struct {
}

// {{.Description}}
func (t *{{.Name | title}}) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// TODO: 实现工具逻辑
	// 参数示例:
	// {{- range .Params}}
	// {{.Name}} := request.GetArguments()["{{.Name}}"]
	// {{- end}}
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			{{- range .Response}}
			mcp.{{.Type | title}}Content{
				Type: "{{.Type}}",
				// TODO: 填充{{.Type}}内容
			},
			{{- end}}
		},
	}, nil
}

func (t *{{.Name | title}}) New() mcp.Tool {
	return mcp.NewTool("{{.Name}}",
		mcp.WithDescription("{{.Description}}"),
		{{- range .Params}}
		mcp.With{{.Type | title}}("{{.Name}}",
			{{- if .Required}}mcp.Required(),{{end}}
			mcp.Description("{{.Description}}"),
			{{- if .Default}}
              {{- if eq .Type "string"}}
              mcp.DefaultString("{{.Default}}"),
              {{- else if eq .Type "number"}}
              mcp.DefaultNumber({{.Default}}),
              {{- else if eq .Type "boolean"}}
              mcp.DefaultBoolean({{if or (eq .Default "true") (eq .Default "True")}}true{{else}}false{{end}}),
              {{- else if eq .Type "array"}}
              // 注意：数组默认值需要在后端代码中预处理为正确的格式
              // mcp.DefaultArray({{.Default}}),
              {{- end}}
            {{- end}}
		),
		{{- end}}
	)
}
