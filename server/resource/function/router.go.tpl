{{if .IsPlugin}}
group.{{.Method}}("{{.FuncName}}", api{{.StructName}}.{{.FuncName}})  // 等待开发的方法
{{else}}
{{.Abbreviation}}RouterWithoutAuth.{{.Method}}("{{.FuncName}}", {{.Abbreviation}}Api.{{.FuncName}})  // 等待开发的方法
{{end}}





