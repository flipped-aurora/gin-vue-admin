package {{.Package}}

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/{{.Package}}"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    {{.Package}}Req "github.com/flipped-aurora/gin-vue-admin/server/model/{{.Package}}/request"
    {{- if .AutoCreateResource }}
    "gorm.io/gorm"
    {{- end}}
)

type {{.StructName}}Service struct {
}

{{- $db := "" }}
{{- if eq .BusinessDB "" }}
 {{- $db = "global.GVA_DB" }}
{{- else}}
 {{- $db =  printf "global.MustGetGlobalDBByDBName(\"%s\")" .BusinessDB   }}
{{- end}}

// Create{{.StructName}} 创建{{.StructName}}记录
// Author [piexlmax](https://github.com/piexlmax)
func ({{.Abbreviation}}Service *{{.StructName}}Service) Create{{.StructName}}({{.Abbreviation}} *{{.Package}}.{{.StructName}}) (err error) {
	err = {{$db}}.Create({{.Abbreviation}}).Error
	return err
}

// Delete{{.StructName}} 删除{{.StructName}}记录
// Author [piexlmax](https://github.com/piexlmax)
func ({{.Abbreviation}}Service *{{.StructName}}Service)Delete{{.StructName}}({{.Abbreviation}} {{.Package}}.{{.StructName}}) (err error) {
	{{- if .AutoCreateResource }}
	err = {{$db}}.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&{{.Package}}.{{.StructName}}{}).Where("id = ?", {{.Abbreviation}}.ID).Update("deleted_by", {{.Abbreviation}}.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&{{.Abbreviation}}).Error; err != nil {
              return err
        }
        return nil
	})
    {{- else }}
	err = {{$db}}.Delete(&{{.Abbreviation}}).Error
	{{- end }}
	return err
}

// Delete{{.StructName}}ByIds 批量删除{{.StructName}}记录
// Author [piexlmax](https://github.com/piexlmax)
func ({{.Abbreviation}}Service *{{.StructName}}Service)Delete{{.StructName}}ByIds(ids request.IdsReq{{- if .AutoCreateResource }},deleted_by uint{{- end}}) (err error) {
	{{- if .AutoCreateResource }}
	err = {{$db}}.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&{{.Package}}.{{.StructName}}{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&{{.Package}}.{{.StructName}}{}).Error; err != nil {
            return err
        }
        return nil
    })
    {{- else}}
	err = {{$db}}.Delete(&[]{{.Package}}.{{.StructName}}{},"id in ?",ids.Ids).Error
    {{- end}}
	return err
}

// Update{{.StructName}} 更新{{.StructName}}记录
// Author [piexlmax](https://github.com/piexlmax)
func ({{.Abbreviation}}Service *{{.StructName}}Service)Update{{.StructName}}({{.Abbreviation}} {{.Package}}.{{.StructName}}) (err error) {
	err = {{$db}}.Save(&{{.Abbreviation}}).Error
	return err
}

// Get{{.StructName}} 根据id获取{{.StructName}}记录
// Author [piexlmax](https://github.com/piexlmax)
func ({{.Abbreviation}}Service *{{.StructName}}Service)Get{{.StructName}}(id uint) ({{.Abbreviation}} {{.Package}}.{{.StructName}}, err error) {
	err = {{$db}}.Where("id = ?", id).First(&{{.Abbreviation}}).Error
	return
}

// Get{{.StructName}}InfoList 分页获取{{.StructName}}记录
// Author [piexlmax](https://github.com/piexlmax)
func ({{.Abbreviation}}Service *{{.StructName}}Service)Get{{.StructName}}InfoList(info {{.Package}}Req.{{.StructName}}Search) (list []{{.Package}}.{{.StructName}}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := {{$db}}.Model(&{{.Package}}.{{.StructName}}{})
    var {{.Abbreviation}}s []{{.Package}}.{{.StructName}}
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
        {{- range .Fields}}
            {{- if .FieldSearchType}}
                {{- if eq .FieldType "string" }}
    if info.{{.FieldName}} != "" {
        db = db.Where("{{.ColumnName}} {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+ {{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
    }
    {{- else if eq .FieldSearchType "BETWEEN" "NOT BETWEEN"}}
        if info.Start{{.FieldName}} != nil && info.End{{.FieldName}} != nil {
            db = db.Where("{{.ColumnName}} {{.FieldSearchType}} ? AND ? ",info.Start{{.FieldName}},info.End{{.FieldName}})
        }
    {{- else}}
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
    {{- if .NeedSort}}
        var OrderStr string
        orderMap := make(map[string]bool)
       {{- range .Fields}}
            {{- if .Sort}}
         	orderMap["{{.FieldJson}}"] = true
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

	err = db.Limit(limit).Offset(offset).Find(&{{.Abbreviation}}s).Error
	return  {{.Abbreviation}}s, total, err
}
