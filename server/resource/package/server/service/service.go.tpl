{{- $db := "" }}
{{- if eq .BusinessDB "" }}
 {{- $db = "global.GVA_DB" }}
{{- else}}
 {{- $db =  printf "global.MustGetGlobalDBByDBName(\"%s\")" .BusinessDB   }}
{{- end}}

{{- if .IsAdd}}

// Get{{.StructName}}InfoList 新增搜索语句
        {{- range .Fields}}
            {{- if .FieldSearchType}}
                {{- if or (eq .FieldType "enum") (eq .FieldType "pictures") (eq .FieldType "picture") (eq .FieldType "video") (eq .FieldType "json") }}
if info.{{.FieldName}} != "" {
        {{- if or (eq .FieldType "enum") }}
    db = db.Where("{{.ColumnName}} {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+ {{ end }}*info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
        {{- else}}
// 数据类型为复杂类型，请根据业务需求自行实现复杂类型的查询业务
        {{- end}}
}
    {{- else if eq .FieldSearchType "BETWEEN" "NOT BETWEEN"}}
if info.Start{{.FieldName}} != nil && info.End{{.FieldName}} != nil {
    db = db.Where("{{.ColumnName}} {{.FieldSearchType}} ? AND ? ",info.Start{{.FieldName}},info.End{{.FieldName}})
}
    {{- else}}
if info.{{.FieldName}} != nil{{- if eq .FieldType "string" }} && *info.{{.FieldName}} != ""{{- end }} {
    db = db.Where("{{.ColumnName}} {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}*info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
}
            {{- end }}
        {{- end }}
    {{- end }}


// Get{{.StructName}}InfoList 新增排序语句 请自行在搜索语句中添加orderMap内容
       {{- range .Fields}}
            {{- if .Sort}}
orderMap["{{.ColumnName}}"] = true
         	{{- end}}
       {{- end}}


{{- if .HasDataSource }}
//  Get{{.StructName}}DataSource()方法新增关联语句
	{{range $key, $value := .DataSourceMap}}
{{$key}} := make([]map[string]any, 0)
{{ $dataDB := "" }}
{{- if eq $value.DBName "" }}
{{ $dataDB = $db }}
{{- else}}
{{ $dataDB = printf "global.MustGetGlobalDBByDBName(\"%s\")" $value.DBName }}
{{- end}}
{{$dataDB}}.Table("{{$value.Table}}"){{- if $value.HasDeletedAt}}.Where("deleted_at IS NULL"){{ end }}.Select("{{$value.Label}} as label,{{$value.Value}} as value").Scan(&{{$key}})
res["{{$key}}"] = {{$key}}
	{{- end }}
{{- end }}
{{- else}}
package {{.Package}}

import (
{{- if not .OnlyTemplate }}
	"{{.Module}}/global"
	"{{.Module}}/model/{{.Package}}"
	{{- if not .IsTree}}
    {{.Package}}Req "{{.Module}}/model/{{.Package}}/request"
    {{- else }}
    "{{.Module}}/utils"
    "errors"
    {{- end }}
    {{- if .AutoCreateResource }}
    "gorm.io/gorm"
    {{- end}}
{{- end }}
)

type {{.StructName}}Service struct {}

{{- if not .OnlyTemplate }}
// Create{{.StructName}} 创建{{.Description}}记录
// Author [yourname](https://github.com/yourname)
func ({{.Abbreviation}}Service *{{.StructName}}Service) Create{{.StructName}}({{.Abbreviation}} *{{.Package}}.{{.StructName}}) (err error) {
	err = {{$db}}.Create({{.Abbreviation}}).Error
	return err
}

// Delete{{.StructName}} 删除{{.Description}}记录
// Author [yourname](https://github.com/yourname)
func ({{.Abbreviation}}Service *{{.StructName}}Service)Delete{{.StructName}}({{.PrimaryField.FieldJson}} string{{- if .AutoCreateResource -}},userID uint{{- end -}}) (err error) {
	{{- if .IsTree }}
       var count int64
	   err = {{$db}}.Find(&{{.Package}}.{{.StructName}}{},"parent_id = ?",{{.PrimaryField.FieldJson}}).Count(&count).Error
	   if count > 0 {
           return errors.New("此节点存在子节点不允许删除")
       }
       if err != nil {
           return err
       }
	{{- end }}

	{{- if .AutoCreateResource }}
	err = {{$db}}.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&{{.Package}}.{{.StructName}}{}).Where("{{.PrimaryField.ColumnName}} = ?", {{.PrimaryField.FieldJson}}).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&{{.Package}}.{{.StructName}}{},"{{.PrimaryField.ColumnName}} = ?",{{.PrimaryField.FieldJson}}).Error; err != nil {
              return err
        }
        return nil
	})
    {{- else }}
	err = {{$db}}.Delete(&{{.Package}}.{{.StructName}}{},"{{.PrimaryField.ColumnName}} = ?",{{.PrimaryField.FieldJson}}).Error
	{{- end }}
	return err
}

// Delete{{.StructName}}ByIds 批量删除{{.Description}}记录
// Author [yourname](https://github.com/yourname)
func ({{.Abbreviation}}Service *{{.StructName}}Service)Delete{{.StructName}}ByIds({{.PrimaryField.FieldJson}}s []string {{- if .AutoCreateResource }},deleted_by uint{{- end}}) (err error) {
	{{- if .AutoCreateResource }}
	err = {{$db}}.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&{{.Package}}.{{.StructName}}{}).Where("{{.PrimaryField.ColumnName}} in ?", {{.PrimaryField.FieldJson}}s).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("{{.PrimaryField.ColumnName}} in ?", {{.PrimaryField.FieldJson}}s).Delete(&{{.Package}}.{{.StructName}}{}).Error; err != nil {
            return err
        }
        return nil
    })
    {{- else}}
	err = {{$db}}.Delete(&[]{{.Package}}.{{.StructName}}{},"{{.PrimaryField.ColumnName}} in ?",{{.PrimaryField.FieldJson}}s).Error
    {{- end}}
	return err
}

// Update{{.StructName}} 更新{{.Description}}记录
// Author [yourname](https://github.com/yourname)
func ({{.Abbreviation}}Service *{{.StructName}}Service)Update{{.StructName}}({{.Abbreviation}} {{.Package}}.{{.StructName}}) (err error) {
	err = {{$db}}.Model(&{{.Package}}.{{.StructName}}{}).Where("{{.PrimaryField.ColumnName}} = ?",{{.Abbreviation}}.{{.PrimaryField.FieldName}}).Updates(&{{.Abbreviation}}).Error
	return err
}

// Get{{.StructName}} 根据{{.PrimaryField.FieldJson}}获取{{.Description}}记录
// Author [yourname](https://github.com/yourname)
func ({{.Abbreviation}}Service *{{.StructName}}Service)Get{{.StructName}}({{.PrimaryField.FieldJson}} string) ({{.Abbreviation}} {{.Package}}.{{.StructName}}, err error) {
	err = {{$db}}.Where("{{.PrimaryField.ColumnName}} = ?", {{.PrimaryField.FieldJson}}).First(&{{.Abbreviation}}).Error
	return
}


{{- if .IsTree }}
// Get{{.StructName}}InfoList 分页获取{{.Description}}记录,Tree模式下不添加分页和搜索
// Author [yourname](https://github.com/yourname)
func ({{.Abbreviation}}Service *{{.StructName}}Service)Get{{.StructName}}InfoList() (list []*{{.Package}}.{{.StructName}},err error) {
    // 创建db
	db := {{$db}}.Model(&{{.Package}}.{{.StructName}}{})
    var {{.Abbreviation}}s []*{{.Package}}.{{.StructName}}

	err = db.Find(&{{.Abbreviation}}s).Error

	return utils.BuildTree({{.Abbreviation}}s), err
}
{{- else }}
// Get{{.StructName}}InfoList 分页获取{{.Description}}记录
// Author [yourname](https://github.com/yourname)
func ({{.Abbreviation}}Service *{{.StructName}}Service)Get{{.StructName}}InfoList(info {{.Package}}Req.{{.StructName}}Search) (list []{{.Package}}.{{.StructName}}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := {{$db}}.Model(&{{.Package}}.{{.StructName}}{})
    var {{.Abbreviation}}s []{{.Package}}.{{.StructName}}
    // 如果有条件搜索 下方会自动创建搜索语句
{{- if .GvaModel }}
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
{{- end }}
        {{- range .Fields}}
            {{- if .FieldSearchType}}
                {{- if or (eq .FieldType "enum") (eq .FieldType "pictures") (eq .FieldType "picture") (eq .FieldType "video") (eq .FieldType "json") }}
    if info.{{.FieldName}} != "" {
        {{- if or (eq .FieldType "enum")}}
        db = db.Where("{{.ColumnName}} {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+ {{ end }}*info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
        {{- else}}
        // 数据类型为复杂类型，请根据业务需求自行实现复杂类型的查询业务
        {{- end}}
    }
    {{- else if eq .FieldSearchType "BETWEEN" "NOT BETWEEN"}}
        if info.Start{{.FieldName}} != nil && info.End{{.FieldName}} != nil {
            db = db.Where("{{.ColumnName}} {{.FieldSearchType}} ? AND ? ",info.Start{{.FieldName}},info.End{{.FieldName}})
        }
    {{- else}}
    if info.{{.FieldName}} != nil{{- if eq .FieldType "string" }} && *info.{{.FieldName}} != ""{{- end }} {
        db = db.Where("{{.ColumnName}} {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}*info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
    }
            {{- end }}
        {{- end }}
    {{- end }}
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
    {{- if .NeedSort}}
        var OrderStr string
        orderMap := make(map[string]bool)
       {{- range .Fields}}
            {{- if .Sort}}
         	orderMap["{{.ColumnName}}"] = true
         	{{- end}}
       {{- end}}
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }
    {{- end}}

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&{{.Abbreviation}}s).Error
	return  {{.Abbreviation}}s, total, err
}

{{- end }}

{{- if .HasDataSource }}
func ({{.Abbreviation}}Service *{{.StructName}}Service)Get{{.StructName}}DataSource() (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)
	{{range $key, $value := .DataSourceMap}}
	   {{$key}} := make([]map[string]any, 0)
	   {{ $dataDB := "" }}
	   {{- if eq $value.DBName "" }}
       {{ $dataDB = $db }}
       {{- else}}
       {{ $dataDB = printf "global.MustGetGlobalDBByDBName(\"%s\")" $value.DBName }}
       {{- end}}
       {{$dataDB}}.Table("{{$value.Table}}"){{- if $value.HasDeletedAt}}.Where("deleted_at IS NULL"){{ end }}.Select("{{$value.Label}} as label,{{$value.Value}} as value").Scan(&{{$key}})
	   res["{{$key}}"] = {{$key}}
	{{- end }}
	return
}
{{- end }}
{{- end }}
func ({{.Abbreviation}}Service *{{.StructName}}Service)Get{{.StructName}}Public() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
{{- end }}
