package utils

import (
	"strings"
	"text/template"
)

// GetTemplateFuncMap 返回模板函数映射，用于在模板中使用
func GetTemplateFuncMap() template.FuncMap {
	return template.FuncMap{
		"formatGormTag":          FormatGormTag,
		"formatFieldTag":         FormatFieldTag,
		"formatSearchCondition":  FormatSearchCondition,
		"formatBetweenCondition": FormatBetweenCondition,
		"formatLikeCondition":    FormatLikeCondition,
		"formatModelField":       FormatModelField,
		"formatRequestField":     FormatRequestField,
		"indent":                 Indent,
	}
}

// FormatGormTag 格式化GORM标签
func FormatGormTag(fieldIndexType, primaryKey, defaultValue, columnName, comment, dataTypeLong string) string {
	var tags []string

	if fieldIndexType != "" {
		tags = append(tags, fieldIndexType)
	}
	if primaryKey == "true" {
		tags = append(tags, "primarykey")
	}
	if defaultValue != "" {
		tags = append(tags, "default:"+defaultValue)
	}
	tags = append(tags, "column:"+columnName)
	if comment != "" {
		tags = append(tags, "comment:"+comment)
	}
	if dataTypeLong != "" {
		tags = append(tags, "size:"+dataTypeLong)
	}

	return strings.Join(tags, ";")
}

// FormatFieldTag 格式化字段标签
func FormatFieldTag(fieldJson string, required bool) string {
	result := `json:"` + fieldJson + `" form:"` + fieldJson + `"`
	if required {
		result += ` binding:"required"`
	}
	return result
}

// FormatSearchCondition 格式化搜索条件
func FormatSearchCondition(columnName, fieldSearchType, fieldName, fieldType string) string {
	var condition string
	if fieldType == "string" {
		condition = `if info.` + fieldName + ` != nil && *info.` + fieldName + ` != "" {`
	} else {
		condition = `if info.` + fieldName + ` != nil {`
	}

	var whereClause string
	if fieldSearchType == "LIKE" {
		whereClause = `db = db.Where("` + columnName + ` ` + fieldSearchType + ` ?","%"+*info.` + fieldName + `+"%")`
	} else {
		whereClause = `db = db.Where("` + columnName + ` ` + fieldSearchType + ` ?",*info.` + fieldName + `)`
	}

	return condition + "\n    " + whereClause + "\n}"
}

// FormatBetweenCondition 格式化BETWEEN条件
func FormatBetweenCondition(columnName, fieldSearchType, fieldName string) string {
	return `if info.Start` + fieldName + ` != nil && info.End` + fieldName + ` != nil {` + "\n" +
		`    db = db.Where("` + columnName + ` ` + fieldSearchType + ` ? AND ? ",info.Start` + fieldName + `,info.End` + fieldName + `)` + "\n" +
		`}`
}

// FormatLikeCondition 格式化LIKE条件
func FormatLikeCondition(columnName, fieldName string) string {
	return `db = db.Where("` + columnName + ` LIKE ?","%"+ *info.` + fieldName + `+"%")`
}

// FormatModelField 格式化模型字段
func FormatModelField(fieldName, fieldType, fieldJson, gormTag string, fieldDesc string) string {
	result := fieldName + "  " + fieldType + " `" + fieldJson + " " + gormTag + "`"
	if fieldDesc != "" {
		result += " //" + fieldDesc
	}
	return result
}

// FormatRequestField 格式化请求字段
func FormatRequestField(fieldName, fieldType, fieldJson string) string {
	return fieldName + "  " + fieldType + " `json:\"" + fieldJson + "\" form:\"" + fieldJson + "\" `"
}

// Indent 缩进文本
func Indent(text string, indent string) string {
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		if line != "" {
			lines[i] = indent + line
		}
	}
	return strings.Join(lines, "\n")
}
