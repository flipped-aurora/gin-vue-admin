package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Create{{.StructName}}
//@description: 创建{{.StructName}}记录
//@param: {{.Abbreviation}} model.{{.StructName}}
//@return: err error

func Create{{.StructName}}({{.Abbreviation}} model.{{.StructName}}) (err error) {
	err = global.GVA_DB.Create(&{{.Abbreviation}}).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Delete{{.StructName}}
//@description: 删除{{.StructName}}记录
//@param: {{.Abbreviation}} model.{{.StructName}}
//@return: err error

func Delete{{.StructName}}({{.Abbreviation}} model.{{.StructName}}) (err error) {
	err = global.GVA_DB.Delete({{.Abbreviation}}).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Delete{{.StructName}}ByIds
//@description: 批量删除{{.StructName}}记录
//@param: ids request.IdsReq
//@return: err error

func Delete{{.StructName}}ByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.{{.StructName}}{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Update{{.StructName}}
//@description: 更新{{.StructName}}记录
//@param: {{.Abbreviation}} *model.{{.StructName}}
//@return: err error

func Update{{.StructName}}({{.Abbreviation}} *model.{{.StructName}}) (err error) {
	err = global.GVA_DB.Save({{.Abbreviation}}).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Get{{.StructName}}
//@description: 根据id获取{{.StructName}}记录
//@param: id uint
//@return: err error, {{.Abbreviation}} model.{{.StructName}}

func Get{{.StructName}}(id uint) (err error, {{.Abbreviation}} model.{{.StructName}}) {
	err = global.GVA_DB.Where("id = ?", id).First(&{{.Abbreviation}}).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Get{{.StructName}}InfoList
//@description: 分页获取{{.StructName}}记录
//@param: info request.{{.StructName}}Search
//@return: err error, list interface{}, total int64

func Get{{.StructName}}InfoList(info request.{{.StructName}}Search) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.{{.StructName}}{})
    var {{.Abbreviation}}s []model.{{.StructName}}
    // 如果有条件搜索 下方会自动创建搜索语句
        {{- range .Fields}}
            {{- if .FieldSearchType}}
                {{- if eq .FieldType "string" }}
    if info.{{.FieldName}} != "" {
        db = db.Where("`{{.ColumnName}}` {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+ {{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
    }
                {{- else if eq .FieldType "bool" }}
    if info.{{.FieldName}} != nil {
        db = db.Where("`{{.ColumnName}}` {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
    }
                {{- else if eq .FieldType "int" }}
    if info.{{.FieldName}} != 0 {
        db = db.Where("`{{.ColumnName}}` {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
    }
                {{- else if eq .FieldType "float64" }}
    if info.{{.FieldName}} != 0 {
        db = db.Where("`{{.ColumnName}}` {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
    }
                {{- else if eq .FieldType "time.Time" }}
    if !info.{{.FieldName}}.IsZero() {
         db = db.Where("`{{.ColumnName}}` {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
    }
                {{- end }}
        {{- end }}
    {{- end }}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&{{.Abbreviation}}s).Error
	return err, {{.Abbreviation}}s, total
}