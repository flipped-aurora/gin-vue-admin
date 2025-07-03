{{- $db := "" }}
{{- if eq .BusinessDB "" }}
 {{- $db = "global.GVA_DB" }}
{{- else}}
 {{- $db =  printf "global.MustGetGlobalDBByDBName(\"%s\")" .BusinessDB   }}
{{- end}}

{{- if .IsAdd}}

// Get{{.StructName}}InfoList 新增搜索语句

    {{ GenerateSearchConditions .Fields }}

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
{{$db}}.Table("{{$value.Table}}"){{- if $value.HasDeletedAt}}.Where("deleted_at IS NULL"){{ end }}.Select("{{$value.Label}} as label,{{$value.Value}} as value").Scan(&{{$key}})
res["{{$key}}"] = {{$key}}
	{{- end }}
{{- end }}
{{- else}}
package service

import (
{{- if not .OnlyTemplate }}
	"context"
	"{{.Module}}/global"
	"{{.Module}}/plugin/{{.Package}}/model"
	{{- if not .IsTree }}
    "{{.Module}}/plugin/{{.Package}}/model/request"
    {{- else }}
    "errors"
    {{- end }}
    {{- if .AutoCreateResource }}
    "gorm.io/gorm"
    {{- end}}
{{- if .IsTree }}
    "{{.Module}}/utils"
{{- end }}
{{- end }}
)

var {{.StructName}} = new({{.Abbreviation}})

type {{.Abbreviation}} struct {}

{{- $db := "" }}
{{- if eq .BusinessDB "" }}
 {{- $db = "global.GVA_DB" }}
{{- else}}
 {{- $db =  printf "global.MustGetGlobalDBByDBName(\"%s\")" .BusinessDB   }}
{{- end}}
{{- if not .OnlyTemplate }}
// Create{{.StructName}} 创建{{.Description}}记录
// Author [yourname](https://github.com/yourname)
func (s *{{.Abbreviation}}) Create{{.StructName}}(ctx context.Context, {{.Abbreviation}} *model.{{.StructName}}) (err error) {
	err = {{$db}}.Create({{.Abbreviation}}).Error
	return err
}

// Delete{{.StructName}} 删除{{.Description}}记录
// Author [yourname](https://github.com/yourname)
func (s *{{.Abbreviation}}) Delete{{.StructName}}(ctx context.Context, {{.PrimaryField.FieldJson}} string{{- if .AutoCreateResource -}},userID uint{{- end -}}) (err error) {

	{{- if .IsTree }}
       var count int64
       err = {{$db}}.Find(&model.{{.StructName}}{},"parent_id = ?",{{.PrimaryField.FieldJson}}).Count(&count).Error
       if count > 0 {
          return errors.New("此节点存在子节点不允许删除")
       }
       if err != nil {
          return err
       }
    {{- end }}

	{{- if .AutoCreateResource }}
	err = {{$db}}.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&model.{{.StructName}}{}).Where("{{.PrimaryField.ColumnName}} = ?", {{.PrimaryField.FieldJson}}).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&model.{{.StructName}}{},"{{.PrimaryField.ColumnName}} = ?",{{.PrimaryField.FieldJson}}).Error; err != nil {
              return err
        }
        return nil
	})
    {{- else }}
	err = {{$db}}.Delete(&model.{{.StructName}}{},"{{.PrimaryField.ColumnName}} = ?",{{.PrimaryField.FieldJson}}).Error
	{{- end }}
	return err
}

// Delete{{.StructName}}ByIds 批量删除{{.Description}}记录
// Author [yourname](https://github.com/yourname)
func (s *{{.Abbreviation}}) Delete{{.StructName}}ByIds(ctx context.Context, {{.PrimaryField.FieldJson}}s []string {{- if .AutoCreateResource }},deleted_by uint{{- end}}) (err error) {
	{{- if .AutoCreateResource }}
	err = {{$db}}.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&model.{{.StructName}}{}).Where("{{.PrimaryField.ColumnName}} in ?", {{.PrimaryField.FieldJson}}s).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("{{.PrimaryField.ColumnName}} in ?", {{.PrimaryField.FieldJson}}s).Delete(&model.{{.StructName}}{}).Error; err != nil {
            return err
        }
        return nil
    })
    {{- else}}
	err = {{$db}}.Delete(&[]model.{{.StructName}}{},"{{.PrimaryField.ColumnName}} in ?",{{.PrimaryField.FieldJson}}s).Error
    {{- end}}
	return err
}

// Update{{.StructName}} 更新{{.Description}}记录
// Author [yourname](https://github.com/yourname)
func (s *{{.Abbreviation}}) Update{{.StructName}}(ctx context.Context, {{.Abbreviation}} model.{{.StructName}}) (err error) {
	err = {{$db}}.Model(&model.{{.StructName}}{}).Where("{{.PrimaryField.ColumnName}} = ?",{{.Abbreviation}}.{{.PrimaryField.FieldName}}).Updates(&{{.Abbreviation}}).Error
	return err
}

// Get{{.StructName}} 根据{{.PrimaryField.FieldJson}}获取{{.Description}}记录
// Author [yourname](https://github.com/yourname)
func (s *{{.Abbreviation}}) Get{{.StructName}}(ctx context.Context, {{.PrimaryField.FieldJson}} string) ({{.Abbreviation}} model.{{.StructName}}, err error) {
	err = {{$db}}.Where("{{.PrimaryField.ColumnName}} = ?", {{.PrimaryField.FieldJson}}).First(&{{.Abbreviation}}).Error
	return
}


{{- if .IsTree }}
// Get{{.StructName}}InfoList 分页获取{{.Description}}记录,Tree模式下不添加分页和搜索
// Author [yourname](https://github.com/yourname)
func (s *{{.Abbreviation}}) Get{{.StructName}}InfoList(ctx context.Context) (list []*model.{{.StructName}},err error) {
    // 创建db
	db := {{$db}}.Model(&model.{{.StructName}}{})
    var {{.Abbreviation}}s []*model.{{.StructName}}

	err = db.Find(&{{.Abbreviation}}s).Error

	return utils.BuildTree({{.Abbreviation}}s), err
}
{{- else }}
// Get{{.StructName}}InfoList 分页获取{{.Description}}记录
// Author [yourname](https://github.com/yourname)
func (s *{{.Abbreviation}}) Get{{.StructName}}InfoList(ctx context.Context, info request.{{.StructName}}Search) (list []model.{{.StructName}}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := {{$db}}.Model(&model.{{.StructName}}{})
    var {{.Abbreviation}}s []model.{{.StructName}}
    // 如果有条件搜索 下方会自动创建搜索语句
{{- if .GvaModel }}
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
{{- end }}
  {{ GenerateSearchConditions .Fields }}
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
    {{- if .NeedSort}}
        var OrderStr string
        orderMap := make(map[string]bool)
      {{- if .GvaModel }}
        orderMap["id"] = true
        orderMap["created_at"] = true
      {{- end }}
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
func (s *{{.Abbreviation}})Get{{.StructName}}DataSource(ctx context.Context) (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)
	{{range $key, $value := .DataSourceMap}}
	   {{$key}} := make([]map[string]any, 0)
	   {{$db}}.Table("{{$value.Table}}"){{- if $value.HasDeletedAt}}.Where("deleted_at IS NULL"){{ end }}.Select("{{$value.Label}} as label,{{$value.Value}} as value").Scan(&{{$key}})
	   res["{{$key}}"] = {{$key}}
	{{- end }}
	return
}
{{- end }}
{{- end }}

func (s *{{.Abbreviation}})Get{{.StructName}}Public(ctx context.Context) {

}
{{- end }}
