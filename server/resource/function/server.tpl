{{- $db := "" }}
{{- if eq .BusinessDB "" }}
 {{- $db = "global.GVA_DB" }}
{{- else}}
 {{- $db =  printf "global.MustGetGlobalDBByDBName(\"%s\")" .BusinessDB   }}
{{- end}}

// Get{{.StructName}}InfoList 分页获取{{.Description}}记录
// Author [yourname](https://github.com/yourname)
func ({{.Abbreviation}}Service *{{.StructName}}Service){{.FuncName}}() (err error) {
	// 请在这里实现自己的业务逻辑
	db := {{$db}}.Model(&{{.Package}}.{{.StructName}}{})
    return db.Error
}