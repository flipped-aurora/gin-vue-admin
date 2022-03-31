package {{.Package}}

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/{{.Package}}"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    {{.Package}}Req "github.com/flipped-aurora/gin-vue-admin/server/model/{{.Package}}/request"
)

type {{.StructName}}Service struct {
}

// Create{{.StructName}} 创建{{.StructName}}记录
// Author [piexlmax](https://github.com/piexlmax)
func ({{.Abbreviation}}Service *{{.StructName}}Service) Create{{.StructName}}({{.Abbreviation}} {{.Package}}.{{.StructName}}) (err error) {
	err = global.GVA_DB.Create(&{{.Abbreviation}}).Error
	return err
}

// Delete{{.StructName}} 删除{{.StructName}}记录
// Author [piexlmax](https://github.com/piexlmax)
func ({{.Abbreviation}}Service *{{.StructName}}Service)Delete{{.StructName}}({{.Abbreviation}} {{.Package}}.{{.StructName}}) (err error) {
	err = global.GVA_DB.Delete(&{{.Abbreviation}}).Error
	return err
}

// Delete{{.StructName}}ByIds 批量删除{{.StructName}}记录
// Author [piexlmax](https://github.com/piexlmax)
func ({{.Abbreviation}}Service *{{.StructName}}Service)Delete{{.StructName}}ByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]{{.Package}}.{{.StructName}}{},"id in ?",ids.Ids).Error
	return err
}

// Update{{.StructName}} 更新{{.StructName}}记录
// Author [piexlmax](https://github.com/piexlmax)
func ({{.Abbreviation}}Service *{{.StructName}}Service)Update{{.StructName}}({{.Abbreviation}} {{.Package}}.{{.StructName}}) (err error) {
	err = global.GVA_DB.Save(&{{.Abbreviation}}).Error
	return err
}

// Get{{.StructName}} 根据id获取{{.StructName}}记录
// Author [piexlmax](https://github.com/piexlmax)
func ({{.Abbreviation}}Service *{{.StructName}}Service)Get{{.StructName}}(id uint) (err error, {{.Abbreviation}} {{.Package}}.{{.StructName}}) {
	err = global.GVA_DB.Where("id = ?", id).First(&{{.Abbreviation}}).Error
	return
}

// Get{{.StructName}}InfoList 分页获取{{.StructName}}记录
// Author [piexlmax](https://github.com/piexlmax)
func ({{.Abbreviation}}Service *{{.StructName}}Service)Get{{.StructName}}InfoList(info {{.Package}}Req.{{.StructName}}Search) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&{{.Package}}.{{.StructName}}{})
    var {{.Abbreviation}}s []{{.Package}}.{{.StructName}}
    // 如果有条件搜索 下方会自动创建搜索语句
        {{- range .Fields}}
            {{- if .FieldSearchType}}
                {{- if eq .FieldType "string" }}
    if info.{{.FieldName}} != "" {
        db = db.Where("{{.ColumnName}} {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+ {{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
    }
                {{- else if eq .FieldType "bool" }}
    if info.{{.FieldName}} != nil {
        db = db.Where("{{.ColumnName}} {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
    }
                {{- else if eq .FieldType "int" }}
    if info.{{.FieldName}} != nil {
        db = db.Where("{{.ColumnName}} {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
    }
                {{- else if eq .FieldType "float64" }}
    if info.{{.FieldName}} != nil {
        db = db.Where("{{.ColumnName}} {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
    }
                {{- else if eq .FieldType "time.Time" }}
    if info.{{.FieldName}} != nil {
         db = db.Where("{{.ColumnName}} {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
    }
                {{- end }}
        {{- end }}
    {{- end }}
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&{{.Abbreviation}}s).Error
	return err, {{.Abbreviation}}s, total
}
