package mcpTool

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"

	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/mark3labs/mcp-go/mcp"
	"go.uber.org/zap"
)

// 注册工具
func init() {
	RegisterTool(&GVAExecutor{})
}

// GVAExecutor GVA代码生成器
type GVAExecutor struct{}

// ExecuteRequest 执行请求结构
type ExecuteRequest struct {
	ExecutionPlan ExecutionPlan `json:"executionPlan"` // 执行计划
	Requirement   string        `json:"requirement"`   // 原始需求（可选，用于日志记录）
}

// ExecuteResponse 执行响应结构
type ExecuteResponse struct {
	Success        bool              `json:"success"`
	Message        string            `json:"message"`
	PackageID      uint              `json:"packageId,omitempty"`
	HistoryID      uint              `json:"historyId,omitempty"`
	Paths          map[string]string `json:"paths,omitempty"`
	GeneratedPaths []string          `json:"generatedPaths,omitempty"`
	NextActions    []string          `json:"nextActions,omitempty"`
}

// ExecutionPlan 执行计划结构
type ExecutionPlan struct {
	PackageName             string                            `json:"packageName"`
	PackageType             string                            `json:"packageType"` // "plugin" 或 "package"
	NeedCreatedPackage      bool                              `json:"needCreatedPackage"`
	NeedCreatedModules      bool                              `json:"needCreatedModules"`
	NeedCreatedDictionaries bool                              `json:"needCreatedDictionaries"`
	PackageInfo             *request.SysAutoCodePackageCreate `json:"packageInfo,omitempty"`
	ModulesInfo             []*request.AutoCode               `json:"modulesInfo,omitempty"`
	Paths                   map[string]string                 `json:"paths,omitempty"`
	DictionariesInfo        []*DictionaryGenerateRequest      `json:"dictionariesInfo,omitempty"`
}

// New 创建GVA代码生成执行器工具
func (g *GVAExecutor) New() mcp.Tool {
	return mcp.NewTool("gva_execute",
		mcp.WithDescription(`**GVA代码生成执行器：直接执行代码生成，无需确认步骤**

**核心功能：**
- 根据需求分析和当前的包信息判断是否调用，如果需要调用，则根据入参描述生成json，用于直接生成代码
- 支持批量创建多个模块
- 自动创建包、模块、字典等
- 移除了确认步骤，提高执行效率

**使用场景：**
- 在gva_analyze获取了当前的包信息和字典信息之后，如果已经包含了可以使用的包和模块，那就不要调用本mcp
- 根据分析结果直接生成代码
- 适用于自动化代码生成流程

**批量创建功能：**
- 支持在单个ExecutionPlan中创建多个模块
- modulesInfo字段为数组，可包含多个模块配置
- 一次性处理多个模块的创建和字典生成

**新功能：自动字典创建**
- 当结构体字段使用了字典类型（dictType不为空）时，系统会自动检查字典是否存在
- 如果字典不存在，会自动创建对应的字典及默认的字典详情项
- 字典创建包括：字典主表记录和默认的选项值（选项1、选项2等）

**重要限制：**
- 当needCreatedModules=true时，模块创建会自动生成API和菜单，因此不应再调用api_creator和menu_creator工具
- 只有在单独创建API或菜单（不涉及模块创建）时才使用api_creator和menu_creator工具

重要：ExecutionPlan结构体格式要求（支持批量创建）：
{
  "packageName": "包名(string)",
  "packageType": "package或plugin(string)，如果用户提到了使用插件则创建plugin，如果用户没有特定说明则一律选用package。",
  "needCreatedPackage": "是否需要创建包(bool)",
  "needCreatedModules": "是否需要创建模块(bool)",
  "needCreatedDictionaries": "是否需要创建字典(bool)",
  "packageInfo": {
    "desc": "描述(string)",
    "label": "展示名(string)", 
    "template": "package或plugin(string)，如果用户提到了使用插件则创建plugin，如果用户没有特定说明则一律选用package。",
    "packageName": "包名(string)"
  },
  "modulesInfo": [{
    "package": "包名(string，必然是小写开头)",
    "tableName": "数据库表名(string，使用蛇形命名法)",
    "businessDB": "业务数据库(string)",
    "structName": "结构体名(string)",
    "packageName": "文件名称(string)",
    "description": "中文描述(string)",
    "abbreviation": "简称(string)",
    "humpPackageName": "文件名称 一般是结构体名的小驼峰(string)",
    "gvaModel": "是否使用GVA模型(bool) 固定为true 后续不需要创建ID created_at deleted_at updated_at",
    "autoMigrate": "是否自动迁移(bool)",
    "autoCreateResource": "是否创建资源(bool，默认为false)",
    "autoCreateApiToSql": "是否创建API(bool，默认为true)",
    "autoCreateMenuToSql": "是否创建菜单(bool，默认为true)",
    "autoCreateBtnAuth": "是否创建按钮权限(bool，默认为false)",
    "onlyTemplate": "是否仅模板(bool，默认为false)",
    "isTree": "是否树形结构(bool，默认为false)",
    "treeJson": "树形JSON字段(string)",
    "isAdd": "是否新增(bool) 固定为false",
    "generateWeb": "是否生成前端(bool)",
    "generateServer": "是否生成后端(bool)",
    "fields": [{
      "fieldName": "字段名(string)必须大写开头",
      "fieldDesc": "字段描述(string)",
      "fieldType": "字段类型支持：string（字符串）,richtext（富文本）,int（整型）,bool（布尔值）,float64（浮点型）,time.Time（时间）,enum（枚举）,picture（单图片，字符串）,pictures（多图片，json字符串）,video（视频，字符串）,file（文件，json字符串）,json（JSON）,array（数组）",
      "fieldJson": "JSON标签(string)",
      "dataTypeLong": "数据长度(string)",
      "comment": "注释(string)",
      "columnName": "数据库列名(string)",
      "fieldSearchType": "搜索类型:=/>/</>=/<=/NOT BETWEEN/LIKE/BETWEEN/IN/NOT IN等(string)",
      "fieldSearchHide": "是否隐藏搜索(bool)",
      "dictType": "字典类型(string)",
      "form": "表单显示(bool)",
      "table": "表格显示(bool)",
      "desc": "详情显示(bool)",
      "excel": "导入导出(bool)",
      "require": "是否必填(bool)",
      "defaultValue": "默认值(string)",
      "errorText": "错误提示(string)",
      "clearable": "是否可清空(bool)",
      "sort": "是否排序(bool)",
      "primaryKey": "是否主键(bool)",
      "dataSource": "数据源配置(object) - 用于配置字段的关联表信息，结构：{\"dbName\":\"数据库名\",\"table\":\"关联表名\",\"label\":\"显示字段\",\"value\":\"值字段\",\"association\":1或2(1=一对一,2=一对多),\"hasDeletedAt\":true/false}。\n\n**获取表名提示：**\n- 可在 server/model 和 plugin/xxx/model 目录下查看对应模块的 TableName() 接口实现获取实际表名\n- 例如：SysUser 的表名为 \"sys_users\"，ExaFileUploadAndDownload 的表名为 \"exa_file_upload_and_downloads\"\n- 插件模块示例：Info 的表名为 \"gva_announcements_info\"\n\n**获取数据库名提示：**\n- 主数据库：通常使用 \"gva\"（默认数据库标识）\n- 多数据库：可在 config.yaml 的 db-list 配置中查看可用数据库的 alias-name 字段\n- 如果用户未提及关联多数据库信息 则使用默认数据库 默认数据库的情况下 dbName此处填写为空",
      "checkDataSource": "是否检查数据源(bool) - 启用后会验证关联表的存在性",
      "fieldIndexType": "索引类型(string)"
    }]
  }, {
    "package": "包名(string)",
    "tableName": "第二个模块的表名(string)",
    "structName": "第二个模块的结构体名(string)",
    "description": "第二个模块的描述(string)",
    "...": "更多模块配置..."
  }],
	"dictionariesInfo":[{
		"dictType": "字典类型(string) - 用于标识字典的唯一性",
		"dictName": "字典名称(string) - 必须生成，字典的中文名称",
		"description": "字典描述(string) - 字典的用途说明",
		"status": "字典状态(bool) - true启用，false禁用",
		"fieldDesc": "字段描述(string) - 用于AI理解字段含义并生成合适的选项",
		"options": [{
			"label": "显示名称(string) - 用户看到的选项名",
			"value": "选项值(string) - 实际存储的值",
			"sort": "排序号(int) - 数字越小越靠前"
		}]
	}]
}

注意：
1. needCreatedPackage=true时packageInfo必需
2. needCreatedModules=true时modulesInfo必需
3. needCreatedDictionaries=true时dictionariesInfo必需
4. dictionariesInfo中的options字段可选，如果不提供将根据fieldDesc自动生成默认选项
5. 字典创建会在模块创建之前执行，确保模块字段可以正确引用字典类型
6. packageType只能是"package"或"plugin,如果用户提到了使用插件则创建plugin，如果用户没有特定说明则一律选用package。"
7. 字段类型支持：string（字符串）,richtext（富文本）,int（整型）,bool（布尔值）,float64（浮点型）,time.Time（时间）,enum（枚举）,picture（单图片，字符串）,pictures（多图片，json字符串）,video（视频，字符串）,file（文件，json字符串）,json（JSON）,array（数组）
8. 搜索类型支持：=,!=,>,>=,<,<=,NOT BETWEEN/LIKE/BETWEEN/IN/NOT IN
9. gvaModel=true时自动包含ID,CreatedAt,UpdatedAt,DeletedAt字段
10. **重要**：当gvaModel=false时，必须有一个字段的primaryKey=true，否则会导致PrimaryField为nil错误
11. **重要**：当gvaModel=true时，系统会自动设置ID字段为主键，无需手动设置primaryKey=true
12. 智能字典创建功能：当字段使用字典类型(DictType)时，系统会：
   - 自动检查字典是否存在，如果不存在则创建字典
   - 根据字典类型和字段描述智能生成默认选项，支持状态、性别、类型、等级、优先级、审批、角色、布尔值、订单、颜色、尺寸等常见场景
   - 为无法识别的字典类型提供通用默认选项
13. **模块关联配置**：当需要配置模块间的关联关系时，使用dataSource字段：
   - **dbName**: 关联的数据库名称
   - **table**: 关联的表名
   - **label**: 用于显示的字段名（如name、title等）
   - **value**: 用于存储的值字段名（通常是id）
   - **association**: 关联关系类型（1=一对一关联，2=一对多关联）一对一和一对多的前面的一是当前的实体，如果他只能关联另一个实体的一个，则选用一对一，如果他需要关联多个他的关联实体，则选用一对多。
   - **hasDeletedAt**: 关联表是否有软删除字段
   - **checkDataSource**: 设为true时会验证关联表的存在性
   - 示例：{"dbName":"","table":"sys_users","label":"username","value":"id","association":1,"hasDeletedAt":true}
14. **自动字段类型修正**：系统会自动检查和修正字段类型：
   - 当字段配置了dataSource且association=2（一对多关联）时，系统会自动将fieldType修改为'array'
   - 这确保了一对多关联数据的正确存储和处理
   - 修正操作会记录在日志中，便于开发者了解变更情况`),
		mcp.WithObject("executionPlan",
			mcp.Description("执行计划，包含包信息和模块信息"),
			mcp.Required(),
		),
		mcp.WithString("requirement",
			mcp.Description("原始需求描述（可选，用于日志记录）"),
		),
	)
}

// Handle 处理执行请求（移除确认步骤）
func (g *GVAExecutor) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	executionPlanData, ok := request.GetArguments()["executionPlan"]
	if !ok {
		return nil, errors.New("参数错误：executionPlan 必须提供")
	}

	// 解析执行计划
	planJSON, err := json.Marshal(executionPlanData)
	if err != nil {
		return nil, fmt.Errorf("解析执行计划失败: %v", err)
	}

	var plan ExecutionPlan
	err = json.Unmarshal(planJSON, &plan)
	if err != nil {
		return nil, fmt.Errorf("解析执行计划失败: %v\n\n请确保ExecutionPlan格式正确，参考工具描述中的结构体格式要求", err)
	}

	// 验证执行计划的完整性
	if err := g.validateExecutionPlan(&plan); err != nil {
		return nil, fmt.Errorf("执行计划验证失败: %v", err)
	}

	// 获取原始需求（可选）
	var originalRequirement string
	if reqData, ok := request.GetArguments()["requirement"]; ok {
		if reqStr, ok := reqData.(string); ok {
			originalRequirement = reqStr
		}
	}

	// 直接执行创建操作（无确认步骤）
	result := g.executeCreation(ctx, &plan)

	// 如果执行成功且有原始需求，提供代码复检建议
	var reviewMessage string
	if result.Success && originalRequirement != "" {
		global.GVA_LOG.Info("执行完成，返回生成的文件路径供AI进行代码复检...")

		// 构建文件路径信息供AI使用
		var pathsInfo []string
		for _, path := range result.GeneratedPaths {
			pathsInfo = append(pathsInfo, fmt.Sprintf("- %s", path))
		}

		reviewMessage = fmt.Sprintf("\n\n📁 已生成以下文件：\n%s\n\n💡 提示：可以检查生成的代码是否满足原始需求。", strings.Join(pathsInfo, "\n"))
	} else if originalRequirement == "" {
		reviewMessage = "\n\n💡 提示：如需代码复检，请提供原始需求描述。"
	}

	// 序列化响应
	response := ExecuteResponse{
		Success:        result.Success,
		Message:        result.Message,
		PackageID:      result.PackageID,
		HistoryID:      result.HistoryID,
		Paths:          result.Paths,
		GeneratedPaths: result.GeneratedPaths,
		NextActions:    result.NextActions,
	}

	responseJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("序列化结果失败: %v", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.NewTextContent(fmt.Sprintf("执行结果：\n\n%s%s", string(responseJSON), reviewMessage)),
		},
	}, nil
}

// validateExecutionPlan 验证执行计划的完整性
func (g *GVAExecutor) validateExecutionPlan(plan *ExecutionPlan) error {
	// 验证基本字段
	if plan.PackageName == "" {
		return errors.New("packageName 不能为空")
	}
	if plan.PackageType != "package" && plan.PackageType != "plugin" {
		return errors.New("packageType 必须是 'package' 或 'plugin'")
	}

	// 验证packageType和template字段的一致性
	if plan.NeedCreatedPackage && plan.PackageInfo != nil {
		if plan.PackageType != plan.PackageInfo.Template {
			return errors.New("packageType 和 packageInfo.template 必须保持一致")
		}
	}

	// 验证包信息
	if plan.NeedCreatedPackage {
		if plan.PackageInfo == nil {
			return errors.New("当 needCreatedPackage=true 时，packageInfo 不能为空")
		}
		if plan.PackageInfo.PackageName == "" {
			return errors.New("packageInfo.packageName 不能为空")
		}
		if plan.PackageInfo.Template != "package" && plan.PackageInfo.Template != "plugin" {
			return errors.New("packageInfo.template 必须是 'package' 或 'plugin'")
		}
		if plan.PackageInfo.Label == "" {
			return errors.New("packageInfo.label 不能为空")
		}
		if plan.PackageInfo.Desc == "" {
			return errors.New("packageInfo.desc 不能为空")
		}
	}

	// 验证模块信息（批量验证）
	if plan.NeedCreatedModules {
		if len(plan.ModulesInfo) == 0 {
			return errors.New("当 needCreatedModules=true 时，modulesInfo 不能为空")
		}

		// 遍历验证每个模块
		for moduleIndex, moduleInfo := range plan.ModulesInfo {
			if moduleInfo.Package == "" {
				return fmt.Errorf("模块 %d 的 package 不能为空", moduleIndex+1)
			}
			if moduleInfo.StructName == "" {
				return fmt.Errorf("模块 %d 的 structName 不能为空", moduleIndex+1)
			}
			if moduleInfo.TableName == "" {
				return fmt.Errorf("模块 %d 的 tableName 不能为空", moduleIndex+1)
			}
			if moduleInfo.Description == "" {
				return fmt.Errorf("模块 %d 的 description 不能为空", moduleIndex+1)
			}
			if moduleInfo.Abbreviation == "" {
				return fmt.Errorf("模块 %d 的 abbreviation 不能为空", moduleIndex+1)
			}
			if moduleInfo.PackageName == "" {
				return fmt.Errorf("模块 %d 的 packageName 不能为空", moduleIndex+1)
			}
			if moduleInfo.HumpPackageName == "" {
				return fmt.Errorf("模块 %d 的 humpPackageName 不能为空", moduleIndex+1)
			}

			// 验证字段信息
			if len(moduleInfo.Fields) == 0 {
				return fmt.Errorf("模块 %d 的 fields 不能为空，至少需要一个字段", moduleIndex+1)
			}

			for i, field := range moduleInfo.Fields {
				if field.FieldName == "" {
					return fmt.Errorf("模块 %d 字段 %d 的 fieldName 不能为空", moduleIndex+1, i+1)
				}

				// 确保字段名首字母大写
				if len(field.FieldName) > 0 {
					firstChar := string(field.FieldName[0])
					if firstChar >= "a" && firstChar <= "z" {
						moduleInfo.Fields[i].FieldName = strings.ToUpper(firstChar) + field.FieldName[1:]
					}
				}
				if field.FieldDesc == "" {
					return fmt.Errorf("模块 %d 字段 %d 的 fieldDesc 不能为空", moduleIndex+1, i+1)
				}
				if field.FieldType == "" {
					return fmt.Errorf("模块 %d 字段 %d 的 fieldType 不能为空", moduleIndex+1, i+1)
				}
				if field.FieldJson == "" {
					return fmt.Errorf("模块 %d 字段 %d 的 fieldJson 不能为空", moduleIndex+1, i+1)
				}
				if field.ColumnName == "" {
					return fmt.Errorf("模块 %d 字段 %d 的 columnName 不能为空", moduleIndex+1, i+1)
				}

				// 验证字段类型
				validFieldTypes := []string{"string", "int", "int64", "float64", "bool", "time.Time", "enum", "picture", "video", "file", "pictures", "array", "richtext", "json"}
				validType := false
				for _, validFieldType := range validFieldTypes {
					if field.FieldType == validFieldType {
						validType = true
						break
					}
				}
				if !validType {
					return fmt.Errorf("模块 %d 字段 %d 的 fieldType '%s' 不支持，支持的类型：%v", moduleIndex+1, i+1, field.FieldType, validFieldTypes)
				}

				// 验证搜索类型（如果设置了）
				if field.FieldSearchType != "" {
					validSearchTypes := []string{"=", "!=", ">", ">=", "<", "<=", "LIKE", "BETWEEN", "IN", "NOT IN"}
					validSearchType := false
					for _, validType := range validSearchTypes {
						if field.FieldSearchType == validType {
							validSearchType = true
							break
						}
					}
					if !validSearchType {
						return fmt.Errorf("模块 %d 字段 %d 的 fieldSearchType '%s' 不支持，支持的类型：%v", moduleIndex+1, i+1, field.FieldSearchType, validSearchTypes)
					}
				}

				// 验证 dataSource 字段配置
				if field.DataSource != nil {
					associationValue := field.DataSource.Association
					// 当 association 为 2（一对多关联）时，强制修改 fieldType 为 array
					if associationValue == 2 {
						if field.FieldType != "array" {
							global.GVA_LOG.Info(fmt.Sprintf("模块 %d 字段 %d：检测到一对多关联(association=2)，自动将 fieldType 从 '%s' 修改为 'array'", moduleIndex+1, i+1, field.FieldType))
							moduleInfo.Fields[i].FieldType = "array"
						}
					}

					// 验证 association 值的有效性
					if associationValue != 1 && associationValue != 2 {
						return fmt.Errorf("模块 %d 字段 %d 的 dataSource.association 必须是 1（一对一）或 2（一对多）", moduleIndex+1, i+1)
					}
				}
			}

			// 验证主键设置
			if !moduleInfo.GvaModel {
				// 当不使用GVA模型时，必须有且仅有一个字段设置为主键
				primaryKeyCount := 0
				for _, field := range moduleInfo.Fields {
					if field.PrimaryKey {
						primaryKeyCount++
					}
				}
				if primaryKeyCount == 0 {
					return fmt.Errorf("模块 %d：当 gvaModel=false 时，必须有一个字段的 primaryKey=true", moduleIndex+1)
				}
				if primaryKeyCount > 1 {
					return fmt.Errorf("模块 %d：当 gvaModel=false 时，只能有一个字段的 primaryKey=true", moduleIndex+1)
				}
			} else {
				// 当使用GVA模型时，所有字段的primaryKey都应该为false
				for i, field := range moduleInfo.Fields {
					if field.PrimaryKey {
						return fmt.Errorf("模块 %d：当 gvaModel=true 时，字段 %d 的 primaryKey 应该为 false，系统会自动创建ID主键", moduleIndex+1, i+1)
					}
				}
			}
		}
	}

	return nil
}

// executeCreation 执行创建操作
func (g *GVAExecutor) executeCreation(ctx context.Context, plan *ExecutionPlan) *ExecuteResponse {
	result := &ExecuteResponse{
		Success:        false,
		Paths:          make(map[string]string),
		GeneratedPaths: []string{}, // 初始化生成文件路径列表
	}

	// 无论如何都先构建目录结构信息，确保paths始终返回
	result.Paths = g.buildDirectoryStructure(plan)

	// 记录预期生成的文件路径
	result.GeneratedPaths = g.collectExpectedFilePaths(plan)

	if !plan.NeedCreatedModules {
		result.Success = true
		result.Message += "已列出当前功能所涉及的目录结构信息; 请在paths中查看; 并且在对应指定文件中实现相关的业务逻辑; "
		return result
	}

	// 创建包（如果需要）
	if plan.NeedCreatedPackage && plan.PackageInfo != nil {
		packageService := service.ServiceGroupApp.SystemServiceGroup.AutoCodePackage
		err := packageService.Create(ctx, plan.PackageInfo)
		if err != nil {
			result.Message = fmt.Sprintf("创建包失败: %v", err)
			// 即使创建包失败，也要返回paths信息
			return result
		}
		result.Message += "包创建成功; "
	}

	// 创建指定字典（如果需要）
	if plan.NeedCreatedDictionaries && len(plan.DictionariesInfo) > 0 {
		dictResult := g.createDictionariesFromInfo(ctx, plan.DictionariesInfo)
		result.Message += dictResult
	}

	// 批量创建字典和模块（如果需要）
	if plan.NeedCreatedModules && len(plan.ModulesInfo) > 0 {
		templateService := service.ServiceGroupApp.SystemServiceGroup.AutoCodeTemplate

		// 遍历所有模块进行创建
		for _, moduleInfo := range plan.ModulesInfo {

			// 创建模块
			err := moduleInfo.Pretreatment()
			if err != nil {
				result.Message += fmt.Sprintf("模块 %s 信息预处理失败: %v; ", moduleInfo.StructName, err)
				continue // 继续处理下一个模块
			}

			err = templateService.Create(ctx, *moduleInfo)
			if err != nil {
				result.Message += fmt.Sprintf("创建模块 %s 失败: %v; ", moduleInfo.StructName, err)
				continue // 继续处理下一个模块
			}
			result.Message += fmt.Sprintf("模块 %s 创建成功; ", moduleInfo.StructName)
		}

		result.Message += fmt.Sprintf("批量创建完成，共处理 %d 个模块; ", len(plan.ModulesInfo))

		// 添加重要提醒：不要使用其他MCP工具
		result.Message += "\n\n⚠️ 重要提醒：\n"
		result.Message += "模块创建已完成，API和菜单已自动生成。请不要再调用以下MCP工具：\n"
		result.Message += "- api_creator：API权限已在模块创建时自动生成\n"
		result.Message += "- menu_creator：前端菜单已在模块创建时自动生成\n"
		result.Message += "如需修改API或菜单，请直接在系统管理界面中进行配置。\n"
	}

	result.Message += "已构建目录结构信息; "
	result.Success = true

	if result.Message == "" {
		result.Message = "执行计划完成"
	}

	return result
}

// buildDirectoryStructure 构建目录结构信息
func (g *GVAExecutor) buildDirectoryStructure(plan *ExecutionPlan) map[string]string {
	paths := make(map[string]string)

	// 获取配置信息
	autoCodeConfig := global.GVA_CONFIG.AutoCode

	// 构建基础路径
	rootPath := autoCodeConfig.Root
	serverPath := autoCodeConfig.Server
	webPath := autoCodeConfig.Web
	moduleName := autoCodeConfig.Module

	// 如果计划中有包名，使用计划中的包名，否则使用默认
	packageName := "example"
	if plan.PackageName != "" {
		packageName = plan.PackageName
	}

	// 如果计划中有模块信息，获取第一个模块的结构名作为默认值
	structName := "ExampleStruct"
	if len(plan.ModulesInfo) > 0 && plan.ModulesInfo[0].StructName != "" {
		structName = plan.ModulesInfo[0].StructName
	}

	// 根据包类型构建不同的路径结构
	packageType := plan.PackageType
	if packageType == "" {
		packageType = "package" // 默认为package模式
	}

	// 构建服务端路径
	if serverPath != "" {
		serverBasePath := fmt.Sprintf("%s/%s", rootPath, serverPath)

		if packageType == "plugin" {
			// Plugin 模式：所有文件都在 /plugin/packageName/ 目录下
			plugingBasePath := fmt.Sprintf("%s/plugin/%s", serverBasePath, packageName)

			// API 路径
			paths["api"] = fmt.Sprintf("%s/api", plugingBasePath)

			// Service 路径
			paths["service"] = fmt.Sprintf("%s/service", plugingBasePath)

			// Model 路径
			paths["model"] = fmt.Sprintf("%s/model", plugingBasePath)

			// Router 路径
			paths["router"] = fmt.Sprintf("%s/router", plugingBasePath)

			// Request 路径
			paths["request"] = fmt.Sprintf("%s/model/request", plugingBasePath)

			// Response 路径
			paths["response"] = fmt.Sprintf("%s/model/response", plugingBasePath)

			// Plugin 特有文件
			paths["plugin_main"] = fmt.Sprintf("%s/main.go", plugingBasePath)
			paths["plugin_config"] = fmt.Sprintf("%s/plugin.go", plugingBasePath)
			paths["plugin_initialize"] = fmt.Sprintf("%s/initialize", plugingBasePath)
		} else {
			// Package 模式：传统的目录结构
			// API 路径
			paths["api"] = fmt.Sprintf("%s/api/v1/%s", serverBasePath, packageName)

			// Service 路径
			paths["service"] = fmt.Sprintf("%s/service/%s", serverBasePath, packageName)

			// Model 路径
			paths["model"] = fmt.Sprintf("%s/model/%s", serverBasePath, packageName)

			// Router 路径
			paths["router"] = fmt.Sprintf("%s/router/%s", serverBasePath, packageName)

			// Request 路径
			paths["request"] = fmt.Sprintf("%s/model/%s/request", serverBasePath, packageName)

			// Response 路径
			paths["response"] = fmt.Sprintf("%s/model/%s/response", serverBasePath, packageName)
		}
	}

	// 构建前端路径（两种模式相同）
	if webPath != "" {
		webBasePath := fmt.Sprintf("%s/%s", rootPath, webPath)

		if packageType == "plugin" {
			// Plugin 模式：前端文件也在 /plugin/packageName/ 目录下
			pluginWebBasePath := fmt.Sprintf("%s/plugin/%s", webBasePath, packageName)

			// Vue 页面路径
			paths["vue_page"] = fmt.Sprintf("%s/view", pluginWebBasePath)

			// API 路径
			paths["vue_api"] = fmt.Sprintf("%s/api", pluginWebBasePath)
		} else {
			// Package 模式：传统的目录结构
			// Vue 页面路径
			paths["vue_page"] = fmt.Sprintf("%s/view/%s", webBasePath, packageName)

			// API 路径
			paths["vue_api"] = fmt.Sprintf("%s/api/%s", webBasePath, packageName)
		}
	}

	// 添加模块信息
	paths["module"] = moduleName
	paths["package_name"] = packageName
	paths["package_type"] = packageType
	paths["struct_name"] = structName
	paths["root_path"] = rootPath
	paths["server_path"] = serverPath
	paths["web_path"] = webPath

	return paths
}

// collectExpectedFilePaths 收集预期生成的文件路径
func (g *GVAExecutor) collectExpectedFilePaths(plan *ExecutionPlan) []string {
	var paths []string

	// 获取目录结构
	dirPaths := g.buildDirectoryStructure(plan)

	// 如果需要创建模块，添加预期的文件路径
	if plan.NeedCreatedModules && len(plan.ModulesInfo) > 0 {
		for _, moduleInfo := range plan.ModulesInfo {
			structName := moduleInfo.StructName

			// 后端文件
			if apiPath, ok := dirPaths["api"]; ok {
				paths = append(paths, fmt.Sprintf("%s/%s.go", apiPath, strings.ToLower(structName)))
			}
			if servicePath, ok := dirPaths["service"]; ok {
				paths = append(paths, fmt.Sprintf("%s/%s.go", servicePath, strings.ToLower(structName)))
			}
			if modelPath, ok := dirPaths["model"]; ok {
				paths = append(paths, fmt.Sprintf("%s/%s.go", modelPath, strings.ToLower(structName)))
			}
			if routerPath, ok := dirPaths["router"]; ok {
				paths = append(paths, fmt.Sprintf("%s/%s.go", routerPath, strings.ToLower(structName)))
			}
			if requestPath, ok := dirPaths["request"]; ok {
				paths = append(paths, fmt.Sprintf("%s/%s.go", requestPath, strings.ToLower(structName)))
			}
			if responsePath, ok := dirPaths["response"]; ok {
				paths = append(paths, fmt.Sprintf("%s/%s.go", responsePath, strings.ToLower(structName)))
			}

			// 前端文件
			if vuePage, ok := dirPaths["vue_page"]; ok {
				paths = append(paths, fmt.Sprintf("%s/%s.vue", vuePage, strings.ToLower(structName)))
			}
			if vueApi, ok := dirPaths["vue_api"]; ok {
				paths = append(paths, fmt.Sprintf("%s/%s.js", vueApi, strings.ToLower(structName)))
			}
		}
	}

	return paths
}

// checkDictionaryExists 检查字典是否存在
func (g *GVAExecutor) checkDictionaryExists(dictType string) (bool, error) {
	dictionaryService := service.ServiceGroupApp.SystemServiceGroup.DictionaryService
	_, err := dictionaryService.GetSysDictionary(dictType, 0, nil)
	if err != nil {
		// 如果是记录不存在的错误，返回false
		if strings.Contains(err.Error(), "record not found") {
			return false, nil
		}
		// 其他错误返回错误信息
		return false, err
	}
	return true, nil
}

// createDictionariesFromInfo 根据 DictionariesInfo 创建字典
func (g *GVAExecutor) createDictionariesFromInfo(ctx context.Context, dictionariesInfo []*DictionaryGenerateRequest) string {
	var messages []string
	dictionaryService := service.ServiceGroupApp.SystemServiceGroup.DictionaryService
	dictionaryDetailService := service.ServiceGroupApp.SystemServiceGroup.DictionaryDetailService

	messages = append(messages, fmt.Sprintf("开始创建 %d 个指定字典: ", len(dictionariesInfo)))

	for _, dictInfo := range dictionariesInfo {
		// 检查字典是否存在
		exists, err := g.checkDictionaryExists(dictInfo.DictType)
		if err != nil {
			messages = append(messages, fmt.Sprintf("检查字典 %s 时出错: %v; ", dictInfo.DictType, err))
			continue
		}

		if !exists {
			// 字典不存在，创建字典
			dictionary := model.SysDictionary{
				Name:   dictInfo.DictName,
				Type:   dictInfo.DictType,
				Status: utils.Pointer(true),
				Desc:   dictInfo.Description,
			}

			err = dictionaryService.CreateSysDictionary(dictionary)
			if err != nil {
				messages = append(messages, fmt.Sprintf("创建字典 %s 失败: %v; ", dictInfo.DictType, err))
				continue
			}

			messages = append(messages, fmt.Sprintf("成功创建字典 %s (%s); ", dictInfo.DictType, dictInfo.DictName))

			// 获取刚创建的字典ID
			var createdDict model.SysDictionary
			err = global.GVA_DB.Where("type = ?", dictInfo.DictType).First(&createdDict).Error
			if err != nil {
				messages = append(messages, fmt.Sprintf("获取创建的字典失败: %v; ", err))
				continue
			}

			// 创建字典选项
			if len(dictInfo.Options) > 0 {
				successCount := 0
				for _, option := range dictInfo.Options {
					dictionaryDetail := model.SysDictionaryDetail{
						Label:           option.Label,
						Value:           option.Value,
						Status:          &[]bool{true}[0], // 默认启用
						Sort:            option.Sort,
						SysDictionaryID: int(createdDict.ID),
					}

					err = dictionaryDetailService.CreateSysDictionaryDetail(dictionaryDetail)
					if err != nil {
						global.GVA_LOG.Warn("创建字典详情项失败", zap.Error(err))
					} else {
						successCount++
					}
				}
				messages = append(messages, fmt.Sprintf("创建了 %d 个字典选项; ", successCount))
			}
		} else {
			messages = append(messages, fmt.Sprintf("字典 %s 已存在，跳过创建; ", dictInfo.DictType))
		}
	}

	return strings.Join(messages, "")
}
