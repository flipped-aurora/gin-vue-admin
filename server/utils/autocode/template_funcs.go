package autocode

import (
	"fmt"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"slices"
	"strings"
	"text/template"
)

// GetTemplateFuncMap 返回模板函数映射，用于在模板中使用
func GetTemplateFuncMap() template.FuncMap {
	return template.FuncMap{
		"RenderField":              RenderField,
		"GenerateSearchConditions": GenerateSearchConditions,
		"GenerateSearchFormItem":   GenerateSearchFormItem,
		"GenerateTableColumn":      GenerateTableColumn,
		"GenerateFormItem":         GenerateFormItem,
		"GenerateDescriptionItem":  GenerateDescriptionItem,
		"GenerateDefaultFormValue": GenerateDefaultFormValue,
	}
}

// 渲染Model中的字段
func RenderField(field systemReq.AutoCodeField) string {
	var builder strings.Builder

	var gormTag strings.Builder
	gormTag.WriteString(`gorm:"`)

	if field.FieldIndexType != "" {
		gormTag.WriteString(field.FieldIndexType + ";")
	}

	if field.PrimaryKey {
		gormTag.WriteString("primarykey;")
	}

	if field.DefaultValue != "" {
		gormTag.WriteString(fmt.Sprintf("default:%s;", field.DefaultValue))
	}

	gormTag.WriteString("column:" + field.ColumnName + ";")

	switch field.FieldType {
	case "enum":
		builder.WriteString(fmt.Sprintf("%s  string `json:\"%s\" form:\"%s\" %stype:enum(%s);comment:%s;\"",
			field.FieldName, field.FieldJson, field.FieldJson, gormTag.String(), field.DataTypeLong, field.Comment))
	case "picture", "video":
		builder.WriteString(fmt.Sprintf("%s  string `json:\"%s\" form:\"%s\" %scomment:%s;",
			field.FieldName, field.FieldJson, field.FieldJson, gormTag.String(), field.Comment))
		if field.DataTypeLong != "" {
			builder.WriteString(fmt.Sprintf("size:%s;", field.DataTypeLong))
		}
		builder.WriteString(`"`)
	case "file", "pictures", "array":
		builder.WriteString(fmt.Sprintf("%s  datatypes.JSON `json:\"%s\" form:\"%s\" %scomment:%s;",
			field.FieldName, field.FieldJson, field.FieldJson, gormTag.String(), field.Comment))
		if field.DataTypeLong != "" {
			builder.WriteString(fmt.Sprintf("size:%s;", field.DataTypeLong))
		}
		builder.WriteString(`" swaggertype:"array,object"`)
	case "richtext":
		builder.WriteString(fmt.Sprintf("%s  *string `json:\"%s\" form:\"%s\" %scomment:%s;",
			field.FieldName, field.FieldJson, field.FieldJson, gormTag.String(), field.Comment))
		if field.DataTypeLong != "" {
			builder.WriteString(fmt.Sprintf("size:%s;", field.DataTypeLong))
		}
		builder.WriteString("type:text;\"")
	case "json":
		builder.WriteString(fmt.Sprintf("%s  datatypes.JSON `json:\"%s\" form:\"%s\" %scomment:%s;",
			field.FieldName, field.FieldJson, field.FieldJson, gormTag.String(), field.Comment))
		if field.DataTypeLong != "" {
			builder.WriteString(fmt.Sprintf("size:%s;", field.DataTypeLong))
		}
		builder.WriteString(`" swaggertype:"object"`)
	default:
		builder.WriteString(fmt.Sprintf("%s  *%s `json:\"%s\" form:\"%s\" %scomment:%s;",
			field.FieldName, field.FieldType, field.FieldJson, field.FieldJson, gormTag.String(), field.Comment))
		if field.DataTypeLong != "" {
			builder.WriteString(fmt.Sprintf("size:%s;", field.DataTypeLong))
		}
		builder.WriteString(`"`)
	}

	if field.Require {
		builder.WriteString(` binding:"required"`)
	}

	builder.WriteString("`")

	if field.FieldDesc != "" {
		builder.WriteString(fmt.Sprintf("  //%s", field.FieldDesc))
	}

	return builder.String()
}

// 格式化搜索条件语句
func GenerateSearchConditions(fields []systemReq.AutoCodeField) string {
	var builder strings.Builder

	for _, field := range fields {
		if field.FieldSearchType == "" {
			continue
		}

		builder.WriteString("\n")

		if slices.Contains([]string{"enum", "pictures", "picture", "video", "json"}, field.FieldType) {
			builder.WriteString(fmt.Sprintf("    if info.%s != \"\" {\n", field.FieldName))
			if field.FieldType == "enum" {
				if field.FieldSearchType == "LIKE" {
					builder.WriteString(fmt.Sprintf("        db = db.Where(\"%s LIKE ?\", \"%%\"+ *info.%s+\"%%\")\n",
						field.ColumnName, field.FieldName))
				} else {
					builder.WriteString(fmt.Sprintf("        db = db.Where(\"%s %s ?\", *info.%s)\n",
						field.ColumnName, field.FieldSearchType, field.FieldName))
				}
			} else {
				builder.WriteString("        // 数据类型为复杂类型，请根据业务需求自行实现复杂类型的查询业务\n")
			}
			builder.WriteString("    }\n")

		} else if field.FieldSearchType == "BETWEEN" || field.FieldSearchType == "NOT BETWEEN" {
			builder.WriteString(fmt.Sprintf("    if info.Start%s != nil && info.End%s != nil {\n",
				field.FieldName, field.FieldName))
			builder.WriteString(fmt.Sprintf("        db = db.Where(\"%s %s ? AND ? \", info.Start%s, info.End%s)\n",
				field.ColumnName, field.FieldSearchType, field.FieldName, field.FieldName))
			builder.WriteString("    }\n")

		} else {
			nullCheck := "info." + field.FieldName + " != nil"
			if field.FieldType == "string" {
				builder.WriteString(fmt.Sprintf("    if %s && *info.%s != \"\" {\n",
					nullCheck, field.FieldName))
			} else {
				builder.WriteString(fmt.Sprintf("    if %s {\n", nullCheck))
			}

			if field.FieldSearchType == "LIKE" {
				builder.WriteString(fmt.Sprintf("        db = db.Where(\"%s LIKE ?\", \"%%\"+ *info.%s+\"%%\")\n",
					field.ColumnName, field.FieldName))
			} else {
				builder.WriteString(fmt.Sprintf("        db = db.Where(\"%s %s ?\", *info.%s)\n",
					field.ColumnName, field.FieldSearchType, field.FieldName))
			}
			builder.WriteString("    }\n")
		}
	}

	return builder.String()
}

// 格式化前端搜索条件
func GenerateSearchFormItem(field systemReq.AutoCodeField) string {
	var builder strings.Builder

	// Start form item
	builder.WriteString(fmt.Sprintf("<el-form-item label=\"%s\" prop=\"%s\">\n", field.FieldDesc, field.FieldJson))

	// Generate different input types based on field properties
	if field.FieldType == "bool" {
		builder.WriteString(fmt.Sprintf("  <el-select v-model=\"searchInfo.%s\" clearable placeholder=\"请选择\">\n", field.FieldJson))
		builder.WriteString("    <el-option key=\"true\" label=\"是\" value=\"true\"></el-option>\n")
		builder.WriteString("    <el-option key=\"false\" label=\"否\" value=\"false\"></el-option>\n")
		builder.WriteString("  </el-select>\n")
	} else if field.DictType != "" {
		multipleAttr := ""
		if field.FieldType == "array" {
			multipleAttr = "multiple "
		}
		builder.WriteString(fmt.Sprintf("  <el-select %sv-model=\"searchInfo.%s\" clearable placeholder=\"请选择\" @clear=\"()=>{searchInfo.%s=undefined}\">\n",
			multipleAttr, field.FieldJson, field.FieldJson))
		builder.WriteString(fmt.Sprintf("    <el-option v-for=\"(item,key) in %sOptions\" :key=\"key\" :label=\"item.label\" :value=\"item.value\" />\n",
			field.DictType))
		builder.WriteString("  </el-select>\n")
	} else if field.CheckDataSource {
		multipleAttr := ""
		if field.DataSource.Association == 2 {
			multipleAttr = "multiple "
		}
		builder.WriteString(fmt.Sprintf("  <el-select %sv-model=\"searchInfo.%s\" placeholder=\"请选择%s\" :clearable=\"%v\">\n",
			multipleAttr, field.FieldJson, field.FieldDesc, field.Clearable))
		builder.WriteString(fmt.Sprintf("    <el-option v-for=\"(item,key) in dataSource.%s\" :key=\"key\" :label=\"item.label\" :value=\"item.value\" />\n",
			field.FieldJson))
		builder.WriteString("  </el-select>\n")
	} else if field.FieldType == "float64" || field.FieldType == "int" {
		if field.FieldSearchType == "BETWEEN" || field.FieldSearchType == "NOT BETWEEN" {
			builder.WriteString(fmt.Sprintf("  <el-input v-model.number=\"searchInfo.start%s\" placeholder=\"最小值\" />\n", field.FieldName))
			builder.WriteString("  —\n")
			builder.WriteString(fmt.Sprintf("  <el-input v-model.number=\"searchInfo.end%s\" placeholder=\"最大值\" />\n", field.FieldName))
		} else {
			builder.WriteString(fmt.Sprintf("  <el-input v-model.number=\"searchInfo.%s\" placeholder=\"搜索条件\" />\n", field.FieldJson))
		}
	} else if field.FieldType == "time.Time" {
		if field.FieldSearchType == "BETWEEN" || field.FieldSearchType == "NOT BETWEEN" {
			builder.WriteString("  <template #label>\n")
			builder.WriteString("    <span>\n")
			builder.WriteString(fmt.Sprintf("      %s\n", field.FieldDesc))
			builder.WriteString("      <el-tooltip content=\"搜索范围是开始日期（包含）至结束日期（不包含）\">\n")
			builder.WriteString("        <el-icon><QuestionFilled /></el-icon>\n")
			builder.WriteString("      </el-tooltip>\n")
			builder.WriteString("    </span>\n")
			builder.WriteString("  </template>\n")
			builder.WriteString(fmt.Sprintf("  <el-date-picker v-model=\"searchInfo.start%s\" type=\"datetime\" placeholder=\"开始日期\" "+
				":disabled-date=\"time=> searchInfo.end%s ? time.getTime() > searchInfo.end%s.getTime() : false\"></el-date-picker>\n",
				field.FieldName, field.FieldName, field.FieldName))
			builder.WriteString("  —\n")
			builder.WriteString(fmt.Sprintf("  <el-date-picker v-model=\"searchInfo.end%s\" type=\"datetime\" placeholder=\"结束日期\" "+
				":disabled-date=\"time=> searchInfo.start%s ? time.getTime() < searchInfo.start%s.getTime() : false\"></el-date-picker>\n",
				field.FieldName, field.FieldName, field.FieldName))
		} else {
			builder.WriteString(fmt.Sprintf("  <el-date-picker v-model=\"searchInfo.%s\" type=\"datetime\" placeholder=\"搜索条件\"></el-date-picker>\n",
				field.FieldJson))
		}
	} else {
		builder.WriteString(fmt.Sprintf("  <el-input v-model=\"searchInfo.%s\" placeholder=\"搜索条件\" />\n", field.FieldJson))
	}

	// Close form item
	builder.WriteString("</el-form-item>")

	return builder.String()
}

// GenerateTableColumn generates HTML for table column based on field properties
func GenerateTableColumn(field systemReq.AutoCodeField) string {
	var builder strings.Builder

	// Add sortable attribute if needed
	sortAttr := ""
	if field.Sort {
		sortAttr = " sortable"
	}

	// Handle different field types
	if field.CheckDataSource {
		builder.WriteString(fmt.Sprintf("<el-table-column%s align=\"left\" label=\"%s\" prop=\"%s\" width=\"120\">\n",
			sortAttr, field.FieldDesc, field.FieldJson))
		builder.WriteString("    <template #default=\"scope\">\n")

		if field.DataSource.Association == 2 {
			builder.WriteString(fmt.Sprintf("        <el-tag v-for=\"(item,key) in filterDataSource(dataSource.%s,scope.row.%s)\" :key=\"key\">\n",
				field.FieldJson, field.FieldJson))
			builder.WriteString("             {{ item }}\n")
			builder.WriteString("        </el-tag>\n")
		} else {
			builder.WriteString(fmt.Sprintf("        <span>{{ filterDataSource(dataSource.%s,scope.row.%s) }}</span>\n",
				field.FieldJson, field.FieldJson))
		}

		builder.WriteString("    </template>\n")
		builder.WriteString("</el-table-column>")
	} else if field.DictType != "" {
		builder.WriteString(fmt.Sprintf("<el-table-column%s align=\"left\" label=\"%s\" prop=\"%s\" width=\"120\">\n",
			sortAttr, field.FieldDesc, field.FieldJson))
		builder.WriteString("    <template #default=\"scope\">\n")

		if field.FieldType == "array" {
			builder.WriteString(fmt.Sprintf("    <el-tag class=\"mr-1\" v-for=\"item in scope.row.%s\" :key=\"item\"> {{ filterDict(item,%sOptions) }}</el-tag>\n",
				field.FieldJson, field.DictType))
		} else {
			builder.WriteString(fmt.Sprintf("    {{ filterDict(scope.row.%s,%sOptions) }}\n",
				field.FieldJson, field.DictType))
		}

		builder.WriteString("    </template>\n")
		builder.WriteString("</el-table-column>")
	} else if field.FieldType == "bool" {
		builder.WriteString(fmt.Sprintf("<el-table-column%s align=\"left\" label=\"%s\" prop=\"%s\" width=\"120\">\n",
			sortAttr, field.FieldDesc, field.FieldJson))
		builder.WriteString(fmt.Sprintf("    <template #default=\"scope\">{{ formatBoolean(scope.row.%s) }}</template>\n", field.FieldJson))
		builder.WriteString("</el-table-column>")
	} else if field.FieldType == "time.Time" {
		builder.WriteString(fmt.Sprintf("<el-table-column%s align=\"left\" label=\"%s\" prop=\"%s\" width=\"180\">\n",
			sortAttr, field.FieldDesc, field.FieldJson))
		builder.WriteString(fmt.Sprintf("   <template #default=\"scope\">{{ formatDate(scope.row.%s) }}</template>\n", field.FieldJson))
		builder.WriteString("</el-table-column>")
	} else if field.FieldType == "picture" {
		builder.WriteString(fmt.Sprintf("<el-table-column label=\"%s\" prop=\"%s\" width=\"200\">\n", field.FieldDesc, field.FieldJson))
		builder.WriteString("    <template #default=\"scope\">\n")
		builder.WriteString(fmt.Sprintf("      <el-image preview-teleported style=\"width: 100px; height: 100px\" :src=\"getUrl(scope.row.%s)\" fit=\"cover\"/>\n", field.FieldJson))
		builder.WriteString("    </template>\n")
		builder.WriteString("</el-table-column>")
	} else if field.FieldType == "pictures" {
		builder.WriteString(fmt.Sprintf("<el-table-column label=\"%s\" prop=\"%s\" width=\"200\">\n", field.FieldDesc, field.FieldJson))
		builder.WriteString("   <template #default=\"scope\">\n")
		builder.WriteString("      <div class=\"multiple-img-box\">\n")
		builder.WriteString(fmt.Sprintf("         <el-image preview-teleported v-for=\"(item,index) in scope.row.%s\" :key=\"index\" style=\"width: 80px; height: 80px\" :src=\"getUrl(item)\" fit=\"cover\"/>\n", field.FieldJson))
		builder.WriteString("     </div>\n")
		builder.WriteString("   </template>\n")
		builder.WriteString("</el-table-column>")
	} else if field.FieldType == "video" {
		builder.WriteString(fmt.Sprintf("<el-table-column label=\"%s\" prop=\"%s\" width=\"200\">\n", field.FieldDesc, field.FieldJson))
		builder.WriteString("   <template #default=\"scope\">\n")
		builder.WriteString("    <video\n")
		builder.WriteString("       style=\"width: 100px; height: 100px\"\n")
		builder.WriteString("       muted\n")
		builder.WriteString("       preload=\"metadata\"\n")
		builder.WriteString("       >\n")
		builder.WriteString(fmt.Sprintf("         <source :src=\"getUrl(scope.row.%s) + '#t=1'\">\n", field.FieldJson))
		builder.WriteString("       </video>\n")
		builder.WriteString("   </template>\n")
		builder.WriteString("</el-table-column>")
	} else if field.FieldType == "richtext" {
		builder.WriteString(fmt.Sprintf("<el-table-column label=\"%s\" prop=\"%s\" width=\"200\">\n", field.FieldDesc, field.FieldJson))
		builder.WriteString("   <template #default=\"scope\">\n")
		builder.WriteString("      [富文本内容]\n")
		builder.WriteString("   </template>\n")
		builder.WriteString("</el-table-column>")
	} else if field.FieldType == "file" {
		builder.WriteString(fmt.Sprintf("<el-table-column label=\"%s\" prop=\"%s\" width=\"200\">\n", field.FieldDesc, field.FieldJson))
		builder.WriteString("    <template #default=\"scope\">\n")
		builder.WriteString("         <div class=\"file-list\">\n")
		builder.WriteString(fmt.Sprintf("           <el-tag v-for=\"file in scope.row.%s\" :key=\"file.uid\" @click=\"onDownloadFile(file.url)\">{{ file.name }}</el-tag>\n", field.FieldJson))
		builder.WriteString("         </div>\n")
		builder.WriteString("    </template>\n")
		builder.WriteString("</el-table-column>")
	} else if field.FieldType == "json" {
		builder.WriteString(fmt.Sprintf("<el-table-column label=\"%s\" prop=\"%s\" width=\"200\">\n", field.FieldDesc, field.FieldJson))
		builder.WriteString("    <template #default=\"scope\">\n")
		builder.WriteString("        [JSON]\n")
		builder.WriteString("    </template>\n")
		builder.WriteString("</el-table-column>")
	} else if field.FieldType == "array" {
		builder.WriteString(fmt.Sprintf("<el-table-column label=\"%s\" prop=\"%s\" width=\"200\">\n", field.FieldDesc, field.FieldJson))
		builder.WriteString("    <template #default=\"scope\">\n")
		builder.WriteString(fmt.Sprintf("       <ArrayCtrl v-model=\"scope.row.%s\"/>\n", field.FieldJson))
		builder.WriteString("    </template>\n")
		builder.WriteString("</el-table-column>")
	} else {
		builder.WriteString(fmt.Sprintf("<el-table-column%s align=\"left\" label=\"%s\" prop=\"%s\" width=\"120\" />\n",
			sortAttr, field.FieldDesc, field.FieldJson))
	}

	return builder.String()
}

func GenerateFormItem(field systemReq.AutoCodeField) string {
	var builder strings.Builder

	// Start form item
	builder.WriteString(fmt.Sprintf("<el-form-item label=\"%s:\" prop=\"%s\">\n", field.FieldDesc, field.FieldJson))

	// Handle different field types
	if field.CheckDataSource {
		multipleAttr := ""
		if field.DataSource.Association == 2 {
			multipleAttr = " multiple"
		}
		builder.WriteString(fmt.Sprintf("    <el-select%s v-model=\"formData.%s\" placeholder=\"请选择%s\" style=\"width:100%%\" :clearable=\"%v\">\n",
			multipleAttr, field.FieldJson, field.FieldDesc, field.Clearable))
		builder.WriteString(fmt.Sprintf("        <el-option v-for=\"(item,key) in dataSource.%s\" :key=\"key\" :label=\"item.label\" :value=\"item.value\" />\n",
			field.FieldJson))
		builder.WriteString("    </el-select>\n")
	} else {
		switch field.FieldType {
		case "bool":
			builder.WriteString(fmt.Sprintf("    <el-switch v-model=\"formData.%s\" active-color=\"#13ce66\" inactive-color=\"#ff4949\" active-text=\"是\" inactive-text=\"否\" clearable ></el-switch>\n",
				field.FieldJson))

		case "string":
			if field.DictType != "" {
				builder.WriteString(fmt.Sprintf("    <el-select v-model=\"formData.%s\" placeholder=\"请选择%s\" style=\"width:100%%\" :clearable=\"%v\">\n",
					field.FieldJson, field.FieldDesc, field.Clearable))
				builder.WriteString(fmt.Sprintf("        <el-option v-for=\"(item,key) in %sOptions\" :key=\"key\" :label=\"item.label\" :value=\"item.value\" />\n",
					field.DictType))
				builder.WriteString("    </el-select>\n")
			} else {
				builder.WriteString(fmt.Sprintf("    <el-input v-model=\"formData.%s\" :clearable=\"%v\" placeholder=\"请输入%s\" />\n",
					field.FieldJson, field.Clearable, field.FieldDesc))
			}

		case "richtext":
			builder.WriteString(fmt.Sprintf("    <RichEdit v-model=\"formData.%s\"/>\n", field.FieldJson))

		case "json":
			builder.WriteString(fmt.Sprintf("    // 此字段为json结构，可以前端自行控制展示和数据绑定模式 需绑定json的key为 formData.%s 后端会按照json的类型进行存取\n", field.FieldJson))
			builder.WriteString(fmt.Sprintf("    {{ formData.%s }}\n", field.FieldJson))

		case "array":
			if field.DictType != "" {
				builder.WriteString(fmt.Sprintf("    <el-select multiple v-model=\"formData.%s\" placeholder=\"请选择%s\" style=\"width:100%%\" :clearable=\"%v\">\n",
					field.FieldJson, field.FieldDesc, field.Clearable))
				builder.WriteString(fmt.Sprintf("        <el-option v-for=\"(item,key) in %sOptions\" :key=\"key\" :label=\"item.label\" :value=\"item.value\" />\n",
					field.DictType))
				builder.WriteString("    </el-select>\n")
			} else {
				builder.WriteString(fmt.Sprintf("    <ArrayCtrl v-model=\"formData.%s\" editable/>\n", field.FieldJson))
			}

		case "int":
			builder.WriteString(fmt.Sprintf("    <el-input v-model.number=\"formData.%s\" :clearable=\"%v\" placeholder=\"请输入%s\" />\n",
				field.FieldJson, field.Clearable, field.FieldDesc))

		case "time.Time":
			builder.WriteString(fmt.Sprintf("    <el-date-picker v-model=\"formData.%s\" type=\"date\" style=\"width:100%%\" placeholder=\"选择日期\" :clearable=\"%v\" />\n",
				field.FieldJson, field.Clearable))

		case "float64":
			builder.WriteString(fmt.Sprintf("    <el-input-number v-model=\"formData.%s\" style=\"width:100%%\" :precision=\"2\" :clearable=\"%v\" />\n",
				field.FieldJson, field.Clearable))

		case "enum":
			builder.WriteString(fmt.Sprintf("    <el-select v-model=\"formData.%s\" placeholder=\"请选择%s\" style=\"width:100%%\" :clearable=\"%v\">\n",
				field.FieldJson, field.FieldDesc, field.Clearable))
			builder.WriteString(fmt.Sprintf("       <el-option v-for=\"item in [%s]\" :key=\"item\" :label=\"item\" :value=\"item\" />\n",
				field.DataTypeLong))
			builder.WriteString("    </el-select>\n")

		case "picture":
			builder.WriteString(fmt.Sprintf("    <SelectImage\n     v-model=\"formData.%s\"\n     file-type=\"image\"\n    />\n", field.FieldJson))

		case "pictures":
			builder.WriteString(fmt.Sprintf("    <SelectImage\n     multiple\n     v-model=\"formData.%s\"\n     file-type=\"image\"\n     />\n", field.FieldJson))

		case "video":
			builder.WriteString(fmt.Sprintf("    <SelectImage\n    v-model=\"formData.%s\"\n    file-type=\"video\"\n    />\n", field.FieldJson))

		case "file":
			builder.WriteString(fmt.Sprintf("    <SelectFile v-model=\"formData.%s\" />\n", field.FieldJson))
		}
	}

	// Close form item
	builder.WriteString("</el-form-item>")

	return builder.String()
}

func GenerateDescriptionItem(field systemReq.AutoCodeField) string {
	var builder strings.Builder

	// Start description item
	builder.WriteString(fmt.Sprintf("<el-descriptions-item label=\"%s\">\n", field.FieldDesc))

	if field.CheckDataSource {
		builder.WriteString("    <template #default=\"scope\">\n")
		if field.DataSource.Association == 2 {
			builder.WriteString(fmt.Sprintf("        <el-tag v-for=\"(item,key) in filterDataSource(dataSource.%s,detailFrom.%s)\" :key=\"key\">\n",
				field.FieldJson, field.FieldJson))
			builder.WriteString("             {{ item }}\n")
			builder.WriteString("        </el-tag>\n")
		} else {
			builder.WriteString(fmt.Sprintf("        <span>{{ filterDataSource(dataSource.%s,detailFrom.%s) }}</span>\n",
				field.FieldJson, field.FieldJson))
		}
		builder.WriteString("    </template>\n")
	} else if field.FieldType != "picture" && field.FieldType != "pictures" &&
		field.FieldType != "file" && field.FieldType != "array" &&
		field.FieldType != "richtext" {
		builder.WriteString(fmt.Sprintf("    {{ detailFrom.%s }}\n", field.FieldJson))
	} else {
		switch field.FieldType {
		case "picture":
			builder.WriteString(fmt.Sprintf("    <el-image style=\"width: 50px; height: 50px\" :preview-src-list=\"returnArrImg(detailFrom.%s)\" :src=\"getUrl(detailFrom.%s)\" fit=\"cover\" />\n",
				field.FieldJson, field.FieldJson))
		case "array":
			builder.WriteString(fmt.Sprintf("    <ArrayCtrl v-model=\"detailFrom.%s\"/>\n", field.FieldJson))
		case "pictures":
			builder.WriteString(fmt.Sprintf("    <el-image style=\"width: 50px; height: 50px; margin-right: 10px\" :preview-src-list=\"returnArrImg(detailFrom.%s)\" :initial-index=\"index\" v-for=\"(item,index) in detailFrom.%s\" :key=\"index\" :src=\"getUrl(item)\" fit=\"cover\" />\n",
				field.FieldJson, field.FieldJson))
		case "richtext":
			builder.WriteString(fmt.Sprintf("    <RichView v-model=\"detailFrom.%s\" />\n", field.FieldJson))
		case "file":
			builder.WriteString(fmt.Sprintf("    <div class=\"fileBtn\" v-for=\"(item,index) in detailFrom.%s\" :key=\"index\">\n", field.FieldJson))
			builder.WriteString("        <el-button type=\"primary\" text bg @click=\"onDownloadFile(item.url)\">\n")
			builder.WriteString("          <el-icon style=\"margin-right: 5px\"><Download /></el-icon>\n")
			builder.WriteString("          {{ item.name }}\n")
			builder.WriteString("        </el-button>\n")
			builder.WriteString("    </div>\n")
		}
	}

	// Close description item
	builder.WriteString("</el-descriptions-item>")

	return builder.String()
}

func GenerateDefaultFormValue(field systemReq.AutoCodeField) string {
	var defaultValue string

	switch field.FieldType {
	case "bool":
		defaultValue = "false"
	case "string", "richtext":
		defaultValue = "''"
	case "int":
		if field.DataSource != nil { // Check if data source exists
			defaultValue = "undefined"
		} else {
			defaultValue = "0"
		}
	case "time.Time":
		defaultValue = "new Date()"
	case "float64":
		defaultValue = "0"
	case "picture", "video":
		defaultValue = "\"\""
	case "pictures", "file", "array":
		defaultValue = "[]"
	case "json":
		defaultValue = "{}"
	default:
		defaultValue = "null"
	}

	return fmt.Sprintf("%s: %s,", field.FieldJson, defaultValue)
}
