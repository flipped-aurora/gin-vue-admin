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
		"GenerateField":            GenerateField,
		"GenerateSearchConditions": GenerateSearchConditions,
		"GenerateSearchFormItem":   GenerateSearchFormItem,
		"GenerateTableColumn":      GenerateTableColumn,
		"GenerateFormItem":         GenerateFormItem,
		"GenerateDescriptionItem":  GenerateDescriptionItem,
		"GenerateDefaultFormValue": GenerateDefaultFormValue,
	}
}

// 渲染Model中的字段
func GenerateField(field systemReq.AutoCodeField) string {
	// 构建gorm标签
	gormTag := ``

	if field.FieldIndexType != "" {
		gormTag += field.FieldIndexType + ";"
	}

	if field.PrimaryKey {
		gormTag += "primarykey;"
	}

	if field.DefaultValue != "" {
		gormTag += fmt.Sprintf("default:%s;", field.DefaultValue)
	}

	if field.Comment != "" {
		gormTag += fmt.Sprintf("comment:%s;", field.Comment)
	}

	gormTag += "column:" + field.ColumnName + ";"

	requireTag := ` binding:"required"` + "`"

	// 根据字段类型构建不同的字段定义
	var result string
	switch field.FieldType {
	case "enum":
		result = fmt.Sprintf(`%s  string `+"`"+`json:"%s" form:"%s" gorm:"%stype:enum(%s);"`+"`",
			field.FieldName, field.FieldJson, field.FieldJson, gormTag, field.DataTypeLong)
	case "picture", "video":
		tagContent := fmt.Sprintf(`json:"%s" form:"%s" gorm:"%s"`,
			field.FieldJson, field.FieldJson, gormTag)
		if field.DataTypeLong != "" {
			tagContent += fmt.Sprintf("size:%s;", field.DataTypeLong)
		}
		result = fmt.Sprintf(`%s  string `+"`"+`%s`+"`"+``, field.FieldName, tagContent)
	case "file", "pictures", "array":
		tagContent := fmt.Sprintf(`json:"%s" form:"%s" gorm:"%s"`,
			field.FieldJson, field.FieldJson, gormTag)
		if field.DataTypeLong != "" {
			tagContent += fmt.Sprintf("size:%s;", field.DataTypeLong)
		}
		result = fmt.Sprintf(`%s  datatypes.JSON `+"`"+`%s swaggertype:"array,object"`+"`"+``,
			field.FieldName, tagContent)
	case "richtext":
		tagContent := fmt.Sprintf(`json:"%s" form:"%s" gorm:"%s"`,
			field.FieldJson, field.FieldJson, gormTag)
		if field.DataTypeLong != "" {
			tagContent += fmt.Sprintf("size:%s;", field.DataTypeLong)
		}
		result = fmt.Sprintf(`%s  *string `+"`"+`%stype:text;"`+"`"+``,
			field.FieldName, tagContent)
	case "json":
		tagContent := fmt.Sprintf(`json:"%s" form:"%s" gorm:"%s"`,
			field.FieldJson, field.FieldJson, gormTag)
		if field.DataTypeLong != "" {
			tagContent += fmt.Sprintf("size:%s;", field.DataTypeLong)
		}
		result = fmt.Sprintf(`%s  datatypes.JSON `+"`"+`%s swaggertype:"object"`+"`"+``,
			field.FieldName, tagContent)
	default:
		tagContent := fmt.Sprintf(`json:"%s" form:"%s" gorm:"%s"`,
			field.FieldJson, field.FieldJson, gormTag)
		if field.DataTypeLong != "" {
			tagContent += fmt.Sprintf("size:%s;", field.DataTypeLong)
		}
		result = fmt.Sprintf(`%s  *%s `+"`"+`%s`+"`"+``,
			field.FieldName, field.FieldType, tagContent)
	}

	if field.Require {
		result = result[0:len(result)-1] + requireTag
	}

	// 添加字段描述
	if field.FieldDesc != "" {
		result += fmt.Sprintf("  //%s", field.FieldDesc)
	}

	return result
}

// 格式化搜索条件语句
func GenerateSearchConditions(fields []*systemReq.AutoCodeField) string {
	var conditions []string

	for _, field := range fields {
		if field.FieldSearchType == "" {
			continue
		}

		var condition string

		if slices.Contains([]string{"enum", "pictures", "picture", "video", "json"}, field.FieldType) {
			if field.FieldType == "enum" {
				if field.FieldSearchType == "LIKE" {
					condition = fmt.Sprintf(`
    if info.%s != "" {
        db = db.Where("%s LIKE ?", "%%"+ *info.%s+"%%")
    }`,
						field.FieldName, field.ColumnName, field.FieldName)
				} else {
					condition = fmt.Sprintf(`
    if info.%s != "" {
        db = db.Where("%s %s ?", *info.%s)
    }`,
						field.FieldName, field.ColumnName, field.FieldSearchType, field.FieldName)
				}
			} else {
				condition = fmt.Sprintf(`
    if info.%s != "" {
        // 数据类型为复杂类型，请根据业务需求自行实现复杂类型的查询业务
    }`, field.FieldName)
			}

		} else if field.FieldSearchType == "BETWEEN" || field.FieldSearchType == "NOT BETWEEN" {
			condition = fmt.Sprintf(`
    if info.Start%s != nil && info.End%s != nil {
        db = db.Where("%s %s ? AND ? ", info.Start%s, info.End%s)
    }`,
				field.FieldName, field.FieldName, field.ColumnName,
				field.FieldSearchType, field.FieldName, field.FieldName)

		} else {
			nullCheck := "info." + field.FieldName + " != nil"
			if field.FieldType == "string" {
				condition = fmt.Sprintf(`
    if %s && *info.%s != "" {`, nullCheck, field.FieldName)
			} else {
				condition = fmt.Sprintf(`
    if %s {`, nullCheck)
			}

			if field.FieldSearchType == "LIKE" {
				condition += fmt.Sprintf(`
        db = db.Where("%s LIKE ?", "%%"+ *info.%s+"%%")
    }`,
					field.ColumnName, field.FieldName)
			} else {
				condition += fmt.Sprintf(`
        db = db.Where("%s %s ?", *info.%s)
    }`,
					field.ColumnName, field.FieldSearchType, field.FieldName)
			}
		}

		conditions = append(conditions, condition)
	}

	return strings.Join(conditions, "")
}

// 格式化前端搜索条件
func GenerateSearchFormItem(field systemReq.AutoCodeField) string {
	// 开始构建表单项
	result := fmt.Sprintf(`<el-form-item label="%s" prop="%s">
`, field.FieldDesc, field.FieldJson)

	// 根据字段属性生成不同的输入类型
	if field.FieldType == "bool" {
		result += fmt.Sprintf(`  <el-select v-model="searchInfo.%s" clearable placeholder="请选择">
`, field.FieldJson)
		result += `    <el-option key="true" label="是" value="true"></el-option>
`
		result += `    <el-option key="false" label="否" value="false"></el-option>
`
		result += `  </el-select>
`
	} else if field.DictType != "" {
		multipleAttr := ""
		if field.FieldType == "array" {
			multipleAttr = "multiple "
		}
		result += fmt.Sprintf(`  <el-select %sv-model="searchInfo.%s" clearable filterable placeholder="请选择" @clear="()=>{searchInfo.%s=undefined}">
`,
			multipleAttr, field.FieldJson, field.FieldJson)
		result += fmt.Sprintf(`    <el-option v-for="(item,key) in %sOptions" :key="key" :label="item.label" :value="item.value" />
`,
			field.DictType)
		result += `  </el-select>
`
	} else if field.CheckDataSource {
		multipleAttr := ""
		if field.DataSource.Association == 2 {
			multipleAttr = "multiple "
		}
		result += fmt.Sprintf(`  <el-select %sv-model="searchInfo.%s" filterable placeholder="请选择%s" :clearable="%v">
`,
			multipleAttr, field.FieldJson, field.FieldDesc, field.Clearable)
		result += fmt.Sprintf(`    <el-option v-for="(item,key) in dataSource.%s" :key="key" :label="item.label" :value="item.value" />
`,
			field.FieldJson)
		result += `  </el-select>
`
	} else if field.FieldType == "float64" || field.FieldType == "int" {
		if field.FieldSearchType == "BETWEEN" || field.FieldSearchType == "NOT BETWEEN" {
			result += fmt.Sprintf(`  <el-input v-model.number="searchInfo.start%s" placeholder="最小值" />
`, field.FieldName)
			result += `  —
`
			result += fmt.Sprintf(`  <el-input v-model.number="searchInfo.end%s" placeholder="最大值" />
`, field.FieldName)
		} else {
			result += fmt.Sprintf(`  <el-input v-model.number="searchInfo.%s" placeholder="搜索条件" />
`, field.FieldJson)
		}
	} else if field.FieldType == "time.Time" {
		if field.FieldSearchType == "BETWEEN" || field.FieldSearchType == "NOT BETWEEN" {
			result += `  <template #label>
`
			result += `    <span>
`
			result += fmt.Sprintf(`      %s
`, field.FieldDesc)
			result += `      <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
`
			result += `        <el-icon><QuestionFilled /></el-icon>
`
			result += `      </el-tooltip>
`
			result += `    </span>
`
			result += `  </template>
`
			result += fmt.Sprintf(`  <el-date-picker v-model="searchInfo.start%s" type="datetime" placeholder="开始日期" `+
				`:disabled-date="time=> searchInfo.end%s ? time.getTime() > searchInfo.end%s.getTime() : false"></el-date-picker>
`,
				field.FieldName, field.FieldName, field.FieldName)
			result += `  —
`
			result += fmt.Sprintf(`  <el-date-picker v-model="searchInfo.end%s" type="datetime" placeholder="结束日期" `+
				`:disabled-date="time=> searchInfo.start%s ? time.getTime() < searchInfo.start%s.getTime() : false"></el-date-picker>
`,
				field.FieldName, field.FieldName, field.FieldName)
		} else {
			result += fmt.Sprintf(`  <el-date-picker v-model="searchInfo.%s" type="datetime" placeholder="搜索条件"></el-date-picker>
`,
				field.FieldJson)
		}
	} else {
		result += fmt.Sprintf(`  <el-input v-model="searchInfo.%s" placeholder="搜索条件" />
`, field.FieldJson)
	}

	// 关闭表单项
	result += `</el-form-item>`

	return result
}

// GenerateTableColumn generates HTML for table column based on field properties
func GenerateTableColumn(field systemReq.AutoCodeField) string {
	// Add sortable attribute if needed
	sortAttr := ""
	if field.Sort {
		sortAttr = " sortable"
	}

	// Handle different field types
	if field.CheckDataSource {
		result := fmt.Sprintf(`<el-table-column%s align="left" label="%s" prop="%s" width="120">
`,
			sortAttr, field.FieldDesc, field.FieldJson)
		result += `    <template #default="scope">
`

		if field.DataSource.Association == 2 {
			result += fmt.Sprintf(`        <el-tag v-for="(item,key) in filterDataSource(dataSource.%s,scope.row.%s)" :key="key">
`,
				field.FieldJson, field.FieldJson)
			result += `             {{ item }}
`
			result += `        </el-tag>
`
		} else {
			result += fmt.Sprintf(`        <span>{{ filterDataSource(dataSource.%s,scope.row.%s) }}</span>
`,
				field.FieldJson, field.FieldJson)
		}

		result += `    </template>
`
		result += `</el-table-column>`
		return result
	} else if field.DictType != "" {
		result := fmt.Sprintf(`<el-table-column%s align="left" label="%s" prop="%s" width="120">
`,
			sortAttr, field.FieldDesc, field.FieldJson)
		result += `    <template #default="scope">
`

		if field.FieldType == "array" {
			result += fmt.Sprintf(`    <el-tag class="mr-1" v-for="item in scope.row.%s" :key="item"> {{ filterDict(item,%sOptions) }}</el-tag>
`,
				field.FieldJson, field.DictType)
		} else {
			result += fmt.Sprintf(`    {{ filterDict(scope.row.%s,%sOptions) }}
`,
				field.FieldJson, field.DictType)
		}

		result += `    </template>
`
		result += `</el-table-column>`
		return result
	} else if field.FieldType == "bool" {
		result := fmt.Sprintf(`<el-table-column%s align="left" label="%s" prop="%s" width="120">
`,
			sortAttr, field.FieldDesc, field.FieldJson)
		result += fmt.Sprintf(`    <template #default="scope">{{ formatBoolean(scope.row.%s) }}</template>
`, field.FieldJson)
		result += `</el-table-column>`
		return result
	} else if field.FieldType == "time.Time" {
		result := fmt.Sprintf(`<el-table-column%s align="left" label="%s" prop="%s" width="180">
`,
			sortAttr, field.FieldDesc, field.FieldJson)
		result += fmt.Sprintf(`   <template #default="scope">{{ formatDate(scope.row.%s) }}</template>
`, field.FieldJson)
		result += `</el-table-column>`
		return result
	} else if field.FieldType == "picture" {
		result := fmt.Sprintf(`<el-table-column label="%s" prop="%s" width="200">
`, field.FieldDesc, field.FieldJson)
		result += `    <template #default="scope">
`
		result += fmt.Sprintf(`      <el-image preview-teleported style="width: 100px; height: 100px" :src="getUrl(scope.row.%s)" fit="cover"/>
`, field.FieldJson)
		result += `    </template>
`
		result += `</el-table-column>`
		return result
	} else if field.FieldType == "pictures" {
		result := fmt.Sprintf(`<el-table-column label="%s" prop="%s" width="200">
`, field.FieldDesc, field.FieldJson)
		result += `   <template #default="scope">
`
		result += `      <div class="multiple-img-box">
`
		result += fmt.Sprintf(`         <el-image preview-teleported v-for="(item,index) in scope.row.%s" :key="index" style="width: 80px; height: 80px" :src="getUrl(item)" fit="cover"/>
`, field.FieldJson)
		result += `     </div>
`
		result += `   </template>
`
		result += `</el-table-column>`
		return result
	} else if field.FieldType == "video" {
		result := fmt.Sprintf(`<el-table-column label="%s" prop="%s" width="200">
`, field.FieldDesc, field.FieldJson)
		result += `   <template #default="scope">
`
		result += `    <video
`
		result += `       style="width: 100px; height: 100px"
`
		result += `       muted
`
		result += `       preload="metadata"
`
		result += `       >
`
		result += fmt.Sprintf(`         <source :src="getUrl(scope.row.%s) + '#t=1'">
`, field.FieldJson)
		result += `       </video>
`
		result += `   </template>
`
		result += `</el-table-column>`
		return result
	} else if field.FieldType == "richtext" {
		result := fmt.Sprintf(`<el-table-column label="%s" prop="%s" width="200">
`, field.FieldDesc, field.FieldJson)
		result += `   <template #default="scope">
`
		result += `      [富文本内容]
`
		result += `   </template>
`
		result += `</el-table-column>`
		return result
	} else if field.FieldType == "file" {
		result := fmt.Sprintf(`<el-table-column label="%s" prop="%s" width="200">
`, field.FieldDesc, field.FieldJson)
		result += `    <template #default="scope">
`
		result += `         <div class="file-list">
`
		result += fmt.Sprintf(`           <el-tag v-for="file in scope.row.%s" :key="file.uid" @click="onDownloadFile(file.url)">{{ file.name }}</el-tag>
`, field.FieldJson)
		result += `         </div>
`
		result += `    </template>
`
		result += `</el-table-column>`
		return result
	} else if field.FieldType == "json" {
		result := fmt.Sprintf(`<el-table-column label="%s" prop="%s" width="200">
`, field.FieldDesc, field.FieldJson)
		result += `    <template #default="scope">
`
		result += `        [JSON]
`
		result += `    </template>
`
		result += `</el-table-column>`
		return result
	} else if field.FieldType == "array" {
		result := fmt.Sprintf(`<el-table-column label="%s" prop="%s" width="200">
`, field.FieldDesc, field.FieldJson)
		result += `    <template #default="scope">
`
		result += fmt.Sprintf(`       <ArrayCtrl v-model="scope.row.%s"/>
`, field.FieldJson)
		result += `    </template>
`
		result += `</el-table-column>`
		return result
	} else {
		return fmt.Sprintf(`<el-table-column%s align="left" label="%s" prop="%s" width="120" />
`,
			sortAttr, field.FieldDesc, field.FieldJson)
	}
}

func GenerateFormItem(field systemReq.AutoCodeField) string {
	// 开始构建表单项
	result := fmt.Sprintf(`<el-form-item label="%s:" prop="%s">
`, field.FieldDesc, field.FieldJson)

	// 处理不同字段类型
	if field.CheckDataSource {
		multipleAttr := ""
		if field.DataSource.Association == 2 {
			multipleAttr = " multiple"
		}
		result += fmt.Sprintf(`    <el-select%s v-model="formData.%s" placeholder="请选择%s" filterable style="width:100%%" :clearable="%v">
`,
			multipleAttr, field.FieldJson, field.FieldDesc, field.Clearable)
		result += fmt.Sprintf(`        <el-option v-for="(item,key) in dataSource.%s" :key="key" :label="item.label" :value="item.value" />
`,
			field.FieldJson)
		result += `    </el-select>
`
	} else {
		switch field.FieldType {
		case "bool":
			result += fmt.Sprintf(`    <el-switch v-model="formData.%s" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
`,
				field.FieldJson)

		case "string":
			if field.DictType != "" {
				result += fmt.Sprintf(`    <el-select v-model="formData.%s" placeholder="请选择%s" style="width:100%%" filterable :clearable="%v">
`,
					field.FieldJson, field.FieldDesc, field.Clearable)
				result += fmt.Sprintf(`        <el-option v-for="(item,key) in %sOptions" :key="key" :label="item.label" :value="item.value" />
`,
					field.DictType)
				result += `    </el-select>
`
			} else {
				result += fmt.Sprintf(`    <el-input v-model="formData.%s" :clearable="%v" placeholder="请输入%s" />
`,
					field.FieldJson, field.Clearable, field.FieldDesc)
			}

		case "richtext":
			result += fmt.Sprintf(`    <RichEdit v-model="formData.%s"/>
`, field.FieldJson)

		case "json":
			result += fmt.Sprintf(`    // 此字段为json结构，可以前端自行控制展示和数据绑定模式 需绑定json的key为 formData.%s 后端会按照json的类型进行存取
`, field.FieldJson)
			result += fmt.Sprintf(`    {{ formData.%s }}
`, field.FieldJson)

		case "array":
			if field.DictType != "" {
				result += fmt.Sprintf(`    <el-select multiple v-model="formData.%s" placeholder="请选择%s" filterable style="width:100%%" :clearable="%v">
`,
					field.FieldJson, field.FieldDesc, field.Clearable)
				result += fmt.Sprintf(`        <el-option v-for="(item,key) in %sOptions" :key="key" :label="item.label" :value="item.value" />
`,
					field.DictType)
				result += `    </el-select>
`
			} else {
				result += fmt.Sprintf(`    <ArrayCtrl v-model="formData.%s" editable/>
`, field.FieldJson)
			}

		case "int":
			result += fmt.Sprintf(`    <el-input v-model.number="formData.%s" :clearable="%v" placeholder="请输入%s" />
`,
				field.FieldJson, field.Clearable, field.FieldDesc)

		case "time.Time":
			result += fmt.Sprintf(`    <el-date-picker v-model="formData.%s" type="date" style="width:100%%" placeholder="选择日期" :clearable="%v" />
`,
				field.FieldJson, field.Clearable)

		case "float64":
			result += fmt.Sprintf(`    <el-input-number v-model="formData.%s" style="width:100%%" :precision="2" :clearable="%v" />
`,
				field.FieldJson, field.Clearable)

		case "enum":
			result += fmt.Sprintf(`    <el-select v-model="formData.%s" placeholder="请选择%s" style="width:100%%" filterable :clearable="%v">
`,
				field.FieldJson, field.FieldDesc, field.Clearable)
			result += fmt.Sprintf(`       <el-option v-for="item in [%s]" :key="item" :label="item" :value="item" />
`,
				field.DataTypeLong)
			result += `    </el-select>
`

		case "picture":
			result += fmt.Sprintf(`    <SelectImage
     v-model="formData.%s"
     file-type="image"
    />
`, field.FieldJson)

		case "pictures":
			result += fmt.Sprintf(`    <SelectImage
     multiple
     v-model="formData.%s"
     file-type="image"
     />
`, field.FieldJson)

		case "video":
			result += fmt.Sprintf(`    <SelectImage
    v-model="formData.%s"
    file-type="video"
    />
`, field.FieldJson)

		case "file":
			result += fmt.Sprintf(`    <SelectFile v-model="formData.%s" />
`, field.FieldJson)
		}
	}

	// 关闭表单项
	result += `</el-form-item>`

	return result
}

func GenerateDescriptionItem(field systemReq.AutoCodeField) string {
	// 开始构建描述项
	result := fmt.Sprintf(`<el-descriptions-item label="%s">
`, field.FieldDesc)

	if field.CheckDataSource {
		result += `    <template #default="scope">
`
		if field.DataSource.Association == 2 {
			result += fmt.Sprintf(`        <el-tag v-for="(item,key) in filterDataSource(dataSource.%s,detailFrom.%s)" :key="key">
`,
				field.FieldJson, field.FieldJson)
			result += `             {{ item }}
`
			result += `        </el-tag>
`
		} else {
			result += fmt.Sprintf(`        <span>{{ filterDataSource(dataSource.%s,detailFrom.%s) }}</span>
`,
				field.FieldJson, field.FieldJson)
		}
		result += `    </template>
`
	} else if field.FieldType != "picture" && field.FieldType != "pictures" &&
		field.FieldType != "file" && field.FieldType != "array" &&
		field.FieldType != "richtext" {
		result += fmt.Sprintf(`    {{ detailFrom.%s }}
`, field.FieldJson)
	} else {
		switch field.FieldType {
		case "picture":
			result += fmt.Sprintf(`    <el-image style="width: 50px; height: 50px" :preview-src-list="returnArrImg(detailFrom.%s)" :src="getUrl(detailFrom.%s)" fit="cover" />
`,
				field.FieldJson, field.FieldJson)
		case "array":
			result += fmt.Sprintf(`    <ArrayCtrl v-model="detailFrom.%s"/>
`, field.FieldJson)
		case "pictures":
			result += fmt.Sprintf(`    <el-image style="width: 50px; height: 50px; margin-right: 10px" :preview-src-list="returnArrImg(detailFrom.%s)" :initial-index="index" v-for="(item,index) in detailFrom.%s" :key="index" :src="getUrl(item)" fit="cover" />
`,
				field.FieldJson, field.FieldJson)
		case "richtext":
			result += fmt.Sprintf(`    <RichView v-model="detailFrom.%s" />
`, field.FieldJson)
		case "file":
			result += fmt.Sprintf(`    <div class="fileBtn" v-for="(item,index) in detailFrom.%s" :key="index">
`, field.FieldJson)
			result += `        <el-button type="primary" text bg @click="onDownloadFile(item.url)">
`
			result += `          <el-icon style="margin-right: 5px"><Download /></el-icon>
`
			result += `          {{ item.name }}
`
			result += `        </el-button>
`
			result += `    </div>
`
		}
	}

	// 关闭描述项
	result += `</el-descriptions-item>`

	return result
}

func GenerateDefaultFormValue(field systemReq.AutoCodeField) string {
	// 根据字段类型确定默认值
	var defaultValue string

	switch field.FieldType {
	case "bool":
		defaultValue = "false"
	case "string", "richtext":
		defaultValue = "''"
	case "int":
		if field.DataSource != nil { // 检查数据源是否存在
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

	// 返回格式化后的默认值字符串
	return fmt.Sprintf(`%s: %s,`, field.FieldJson, defaultValue)
}
