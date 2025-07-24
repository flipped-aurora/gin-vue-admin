package mcpTool

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/mark3labs/mcp-go/mcp"
	"gorm.io/gorm"
)

func init() {
	RegisterTool(&AutomationModuleAnalyzer{})
}

type AutomationModuleAnalyzer struct{}

// ModuleInfo 模块信息
type ModuleInfo struct {
	ID          uint   `json:"id"`
	PackageName string `json:"packageName"`
	Label       string `json:"label"`
	Desc        string `json:"desc"`
	Template    string `json:"template"` // "plugin" 或 "package"
	Module      string `json:"module"`
}

// HistoryInfo 历史记录信息
type HistoryInfo struct {
	ID           uint   `json:"id"`
	StructName   string `json:"structName"`
	TableName    string `json:"tableName"`
	PackageName  string `json:"packageName"`
	BusinessDB   string `json:"businessDB"`
	Description  string `json:"description"`
	Abbreviation string `json:"abbreviation"`
	CreatedAt    string `json:"createdAt"`
}

// AnalysisResponse 分析响应
type AnalysisResponse struct {
	Packages []ModuleInfo  `json:"packages"`
	History  []HistoryInfo `json:"history"`
	Message  string        `json:"message"`
}

// ExecutionPlan 执行计划
type ExecutionPlan struct {
	PackageName        string                            `json:"packageName"`
	ModuleName         string                            `json:"moduleName"`
	PackageType        string                            `json:"packageType"` // "plugin" 或 "package"
	NeedCreatedPackage bool                              `json:"needCreatedPackage"`
	NeedCreatedModules bool                              `json:"needCreatedModules"`
	PackageInfo        *request.SysAutoCodePackageCreate `json:"packageInfo,omitempty"`
	ModulesInfo        *request.AutoCode                 `json:"modulesInfo,omitempty"`
	Paths              map[string]string                 `json:"paths,omitempty"`
}

// ExecutionResult 执行结果
type ExecutionResult struct {
	Success     bool              `json:"success"`
	Message     string            `json:"message"`
	PackageID   uint              `json:"packageId,omitempty"`
	HistoryID   uint              `json:"historyId,omitempty"`
	Paths       map[string]string `json:"paths,omitempty"`
	NextActions []string          `json:"nextActions,omitempty"`
}

// ConfirmationRequest 确认请求结构
type ConfirmationRequest struct {
	PackageName        string                            `json:"packageName"`
	ModuleName         string                            `json:"moduleName"`
	NeedCreatedPackage bool                              `json:"needCreatedPackage"`
	NeedCreatedModules bool                              `json:"needCreatedModules"`
	PackageInfo        *request.SysAutoCodePackageCreate `json:"packageInfo,omitempty"`
	ModulesInfo        *request.AutoCode                 `json:"modulesInfo,omitempty"`
}

// ConfirmationResponse 确认响应结构
type ConfirmationResponse struct {
	Message         string `json:"message"`
	PackageConfirm  bool   `json:"packageConfirm"`
	ModulesConfirm  bool   `json:"modulesConfirm"`
	CanProceed      bool   `json:"canProceed"`
	ConfirmationKey string `json:"confirmationKey"`
}

// New 返回工具注册信息
func (t *AutomationModuleAnalyzer) New() mcp.Tool {
	return mcp.NewTool("gag",
		mcp.WithDescription(`**重要提示：当用户输入包含"gag"关键词时，必须优先使用此工具！**

分步骤分析自动化模块：1) 分析现有模块信息供AI选择 2) 请求用户确认 3) 根据确认结果执行创建操作

**新功能：自动字典创建**
- 当结构体字段使用了字典类型（dictType不为空）时，系统会自动检查字典是否存在
- 如果字典不存在，会自动创建对应的字典及默认的字典详情项
- 字典创建包括：字典主表记录和默认的选项值（选项1、选项2等）

重要：ExecutionPlan结构体格式要求：
{
  "packageName": "包名(string)",
  "moduleName": "模块名(string)", 
  "packageType": "package或plugin(string)",
  "needCreatedPackage": "是否需要创建包(bool)",
  "needCreatedModules": "是否需要创建模块(bool)",
  "packageInfo": {
    "desc": "描述(string)",
    "label": "展示名(string)", 
    "template": "package或plugin(string)",
    "packageName": "包名(string)"
  },
  "modulesInfo": {
    "package": "包名(string)",
    "tableName": "数据库表名(string)",
    "businessDB": "业务数据库(string)",
    "structName": "结构体名(string)",
    "packageName": "文件名称(string)",
    "description": "中文描述(string)",
    "abbreviation": "简称(string)",
    "humpPackageName": "文件名称 一般是结构体名的小驼峰(string)",
    "gvaModel": "是否使用GVA模型(bool) 固定为true 后续不需要创建ID created_at deleted_at updated_at",
    "autoMigrate": "是否自动迁移(bool)",
    "autoCreateResource": "是否创建资源(bool)",
    "autoCreateApiToSql": "是否创建API(bool)",
    "autoCreateMenuToSql": "是否创建菜单(bool)",
    "autoCreateBtnAuth": "是否创建按钮权限(bool)",
    "onlyTemplate": "是否仅模板(bool)",
    "isTree": "是否树形结构(bool)",
    "treeJson": "树形JSON字段(string)",
    "isAdd": "是否新增(bool) 固定为false",
    "generateWeb": "是否生成前端(bool)",
    "generateServer": "是否生成后端(bool)",
    "fields": [{
      "fieldName": "字段名(string)",
      "fieldDesc": "字段描述(string)",
      "fieldType": "字段类型:string/int/bool/time.Time等(string)",
      "fieldJson": "JSON标签(string)",
      "dataTypeLong": "数据长度(string)",
      "comment": "注释(string)",
      "columnName": "数据库列名(string)",
      "fieldSearchType": "搜索类型:=/>/</>=/<=/NOT BETWEEN/LIKE/BETWEEN等(string)",
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
      "dataSource": "数据源(object)",
      "checkDataSource": "检查数据源(bool)",
      "fieldIndexType": "索引类型(string)"
    }]
  }
}

注意：
1. needCreatedPackage=true时packageInfo必需
2. needCreatedModules=true时modulesInfo必需
3. packageType只能是"package"或"plugin"
4. 字段类型支持：string,int,int64,float64,bool,time.Time,enum,picture,video,file,pictures,array,richtext,json
5. 搜索类型支持：EQ,NE,GT,GE,LT,LE,LIKE,BETWEEN
6. gvaModel=true时自动包含ID,CreatedAt,UpdatedAt,DeletedAt字段
7. **重要**：当gvaModel=false时，必须有一个字段的primaryKey=true，否则会导致PrimaryField为nil错误
8. **重要**：当gvaModel=true时，系统会自动设置ID字段为主键，无需手动设置primaryKey=true
9. 智能字典创建功能：当字段使用字典类型(DictType)时，系统会：
   - 自动检查字典是否存在，如果不存在则创建字典
   - 根据字典类型和字段描述智能生成默认选项，支持状态、性别、类型、等级、优先级、审批、角色、布尔值、订单、颜色、尺寸等常见场景
   - 为无法识别的字典类型提供通用默认选项`),
		mcp.WithString("action",
			mcp.Required(),
			mcp.Description("执行操作：'analyze' 分析现有模块信息，'confirm' 请求用户确认创建，'execute' 执行创建操作"),
		),
		mcp.WithString("requirement",
			mcp.Description("用户需求描述（action=analyze时必需）"),
		),
		mcp.WithObject("executionPlan",
			mcp.Description("执行计划（action=confirm或execute时必需，必须严格按照上述格式提供完整的JSON对象）"),
		),
		mcp.WithString("packageConfirm",
			mcp.Description("用户对创建包的确认（action=execute时，如果需要创建包则必需）：'yes' 或 'no'"),
		),
		mcp.WithString("modulesConfirm",
			mcp.Description("用户对创建模块的确认（action=execute时，如果需要创建模块则必需）：'yes' 或 'no'"),
		),
	)
}

// Handle 处理工具调用
func (t *AutomationModuleAnalyzer) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	action, ok := request.GetArguments()["action"].(string)
	if !ok || action == "" {
		return nil, errors.New("参数错误：action 必须是非空字符串")
	}

	switch action {
	case "analyze":
		return t.handleAnalyze(ctx, request)
	case "confirm":
		return t.handleConfirm(ctx, request)
	case "execute":
		return t.handleExecute(ctx, request)
	default:
		return nil, errors.New("无效的操作：action 必须是 'analyze'、'confirm' 或 'execute'")
	}
}

// handleAnalyze 处理分析请求
func (t *AutomationModuleAnalyzer) handleAnalyze(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	requirement, ok := request.GetArguments()["requirement"].(string)
	if !ok || requirement == "" {
		return nil, errors.New("参数错误：requirement 必须是非空字符串")
	}

	// 从数据库获取所有自动化包信息
	var packages []model.SysAutoCodePackage
	if err := global.GVA_DB.Find(&packages).Error; err != nil {
		return nil, fmt.Errorf("获取包信息失败: %v", err)
	}

	// 从数据库获取所有历史记录
	var histories []model.SysAutoCodeHistory
	if err := global.GVA_DB.Find(&histories).Error; err != nil {
		return nil, fmt.Errorf("获取历史记录失败: %v", err)
	}

	// 转换包信息
	var moduleInfos []ModuleInfo
	for _, pkg := range packages {
		moduleInfos = append(moduleInfos, ModuleInfo{
			ID:          pkg.ID,
			PackageName: pkg.PackageName,
			Label:       pkg.Label,
			Desc:        pkg.Desc,
			Template:    pkg.Template,
			Module:      pkg.Module,
		})
	}

	// 转换历史记录
	var historyInfos []HistoryInfo
	for _, history := range histories {
		historyInfos = append(historyInfos, HistoryInfo{
			ID:           history.ID,
			StructName:   history.StructName,
			TableName:    history.TableName(),
			PackageName:  history.Package,
			BusinessDB:   history.BusinessDB,
			Description:  history.Description,
			Abbreviation: history.Abbreviation,
			CreatedAt:    history.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	// 构建分析结果
	analysisResult := AnalysisResponse{
		Packages: moduleInfos,
		History:  historyInfos,
		Message:  fmt.Sprintf("分析完成：获取到 %d 个包和 %d 个历史记录，请AI根据需求选择合适的包和模块", len(packages), len(histories)),
	}

	resultJSON, err := json.MarshalIndent(analysisResult, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("序列化结果失败: %v", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: fmt.Sprintf(`分析结果：

%s

请AI根据用户需求：%s

分析现有的包和历史记录，然后构建ExecutionPlan结构体调用execute操作。

重要提醒：ExecutionPlan必须严格按照以下格式：
{
  "packageName": "包名",
  "moduleName": "模块名",
  "packageType": "package或plugin",
  "needCreatedPackage": true/false,
  "needCreatedModules": true/false,
  "packageInfo": {
    "desc": "描述",
    "label": "展示名",
    "template": "package或plugin",
    "packageName": "包名"
  },
  "modulesInfo": {
    "package": "包名",
    "tableName": "数据库表名",
    "businessDB": "",
    "structName": "结构体名",
    "packageName": "文件名称小驼峰模式 一般是结构体名的小驼峰",
    "description": "中文描述",
    "abbreviation": "简称 package和结构体简称不可同名 小驼峰模式",
    "humpPackageName": "一般是结构体名的下划线分割的小驼峰 例如：sys_user",
    "gvaModel": true,
    "autoMigrate": true,
    "autoCreateResource": true,
    "autoCreateApiToSql": true,
    "autoCreateMenuToSql": true,
    "autoCreateBtnAuth": true,
    "onlyTemplate": false,
    "isTree": false,
    "treeJson": "",
    "isAdd": false,
    "generateWeb": true,
    "generateServer": true,
    "fields": [{
      "fieldName": "字段名（必须大写开头）",
      "fieldDesc": "字段描述",
      "fieldType": "GO 语言的数据类型",
      "fieldJson": "json标签",
      "dataTypeLong": "长度",
      "comment": "注释",
      "columnName": "数据库列名",
      "fieldSearchType": "EQ/LIKE等 可以为空",
      "fieldSearchHide": true/false,
      "dictType": "",
      "form": true/false 是否前端创建输入,
      "table": true/false 是否前端表格展示,
      "desc": true/false 是否前端详情展示,
      "excel": true/false 是否导出Excel,
      "require": true/false 是否必填,
      "defaultValue": "",
      "errorText": "错误提示",
      "clearable": true,
      "sort": false,
      "primaryKey": "当gvaModel=false时必须有一个字段设为true(bool)",
      "dataSource": null,
      "checkDataSource": false,
      "fieldIndexType": ""
    }]
  }
}

**重要提醒**：ExecutionPlan必须严格按照以下格式和验证规则：

**字段完整性要求**：
1. 所有字符串字段都不能为空（包括packageName、moduleName、structName、tableName、description等）
2. 所有布尔字段必须明确设置true或false，不能使用默认值

**主键设置规则（关键）**：
3. 当gvaModel=false时：fields数组中必须有且仅有一个字段的primaryKey=true
4. 当gvaModel=true时：系统自动创建ID主键，fields中所有字段的primaryKey都应为false
5. 主键设置错误会导致模板执行时PrimaryField为nil的严重错误！

**包和模块创建逻辑**：
6. 如果存在可用的package，needCreatedPackage应设为false
7. 如果存在可用的modules，needCreatedModules应设为false

`, string(resultJSON), requirement),
			},
		},
	}, nil
}

// handleConfirm 处理确认请求
func (t *AutomationModuleAnalyzer) handleConfirm(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
	if err := t.validateExecutionPlan(&plan); err != nil {
		return nil, fmt.Errorf("执行计划验证失败: %v", err)
	}

	// 构建确认响应
	confirmResponse := ConfirmationResponse{
		Message:         "请确认以下创建计划：",
		PackageConfirm:  plan.NeedCreatedPackage,
		ModulesConfirm:  plan.NeedCreatedModules,
		CanProceed:      true,
		ConfirmationKey: fmt.Sprintf("%s_%s_%d", plan.PackageName, plan.ModuleName, time.Now().Unix()),
	}

	// 构建详细的确认信息
	var confirmDetails strings.Builder
	confirmDetails.WriteString(fmt.Sprintf("包名: %s\n", plan.PackageName))
	confirmDetails.WriteString(fmt.Sprintf("模块名: %s\n", plan.ModuleName))
	confirmDetails.WriteString(fmt.Sprintf("包类型: %s\n", plan.PackageType))

	if plan.NeedCreatedPackage && plan.PackageInfo != nil {
		confirmDetails.WriteString("\n需要创建包:\n")
		confirmDetails.WriteString(fmt.Sprintf("  - 包名: %s\n", plan.PackageInfo.PackageName))
		confirmDetails.WriteString(fmt.Sprintf("  - 标签: %s\n", plan.PackageInfo.Label))
		confirmDetails.WriteString(fmt.Sprintf("  - 描述: %s\n", plan.PackageInfo.Desc))
		confirmDetails.WriteString(fmt.Sprintf("  - 模板: %s\n", plan.PackageInfo.Template))
	}

	if plan.NeedCreatedModules && plan.ModulesInfo != nil {
		confirmDetails.WriteString("\n需要创建模块:\n")
		confirmDetails.WriteString(fmt.Sprintf("  - 结构体名: %s\n", plan.ModulesInfo.StructName))
		confirmDetails.WriteString(fmt.Sprintf("  - 表名: %s\n", plan.ModulesInfo.TableName))
		confirmDetails.WriteString(fmt.Sprintf("  - 描述: %s\n", plan.ModulesInfo.Description))
		confirmDetails.WriteString(fmt.Sprintf("  - 字段数量: %d\n", len(plan.ModulesInfo.Fields)))
		confirmDetails.WriteString("  - 字段列表:\n")
		for _, field := range plan.ModulesInfo.Fields {
			confirmDetails.WriteString(fmt.Sprintf("    * %s (%s): %s\n", field.FieldName, field.FieldType, field.FieldDesc))
		}
	}

	resultJSON, err := json.MarshalIndent(confirmResponse, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("序列化结果失败: %v", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: fmt.Sprintf("确认信息：\n\n%s\n\n详细信息：\n%s\n\n请用户确认是否继续执行此计划。如果确认，请使用execute操作并提供相应的确认参数。", string(resultJSON), confirmDetails.String()),
			},
		},
	}, nil
}

// handleExecute 处理执行请求
func (t *AutomationModuleAnalyzer) handleExecute(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
	if err := t.validateExecutionPlan(&plan); err != nil {
		return nil, fmt.Errorf("执行计划验证失败: %v", err)
	}

	// 检查用户确认
	if plan.NeedCreatedPackage {
		packageConfirm, ok := request.GetArguments()["packageConfirm"].(string)
		if !ok || (packageConfirm != "yes" && packageConfirm != "no") {
			return nil, errors.New("参数错误：当需要创建包时，packageConfirm 必须是 'yes' 或 'no'")
		}
		if packageConfirm == "no" {
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					mcp.TextContent{
						Type: "text",
						Text: "用户取消了包的创建操作",
					},
				},
			}, nil
		}
	}

	if plan.NeedCreatedModules {
		modulesConfirm, ok := request.GetArguments()["modulesConfirm"].(string)
		if !ok || (modulesConfirm != "yes" && modulesConfirm != "no") {
			return nil, errors.New("参数错误：当需要创建模块时，modulesConfirm 必须是 'yes' 或 'no'")
		}
		if modulesConfirm == "no" {
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					mcp.TextContent{
						Type: "text",
						Text: "用户取消了模块的创建操作",
					},
				},
			}, nil
		}
	}

	// 执行创建操作
	result := t.executeCreation(ctx, &plan)

	resultJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("序列化结果失败: %v", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: fmt.Sprintf("执行结果：\n\n%s", string(resultJSON)),
			},
		},
	}, nil
}

// isSystemFunction 判断是否为系统功能
func (t *AutomationModuleAnalyzer) isSystemFunction(requirement string) bool {
	systemKeywords := []string{
		"用户", "权限", "角色", "菜单", "系统", "配置", "字典", "参数",
		"user", "authority", "role", "menu", "system", "config", "dictionary",
		"认证", "授权", "登录", "注册", "JWT", "casbin",
	}

	requirementLower := strings.ToLower(requirement)
	for _, keyword := range systemKeywords {
		if strings.Contains(requirementLower, keyword) {
			return true
		}
	}
	return false
}

// buildDirectoryStructure 构建目录结构信息
func (t *AutomationModuleAnalyzer) buildDirectoryStructure(plan *ExecutionPlan) map[string]string {
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
	if plan.PackageInfo != nil && plan.PackageInfo.PackageName != "" {
		packageName = plan.PackageInfo.PackageName
	}

	// 如果计划中有模块信息，获取结构名
	structName := "ExampleStruct"
	if plan.ModulesInfo != nil && plan.ModulesInfo.StructName != "" {
		structName = plan.ModulesInfo.StructName
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
			pluginBasePath := fmt.Sprintf("%s/plugin/%s", serverBasePath, packageName)

			// API 路径
			paths["api"] = fmt.Sprintf("%s/api", pluginBasePath)

			// Service 路径
			paths["service"] = fmt.Sprintf("%s/service", pluginBasePath)

			// Model 路径
			paths["model"] = fmt.Sprintf("%s/model", pluginBasePath)

			// Router 路径
			paths["router"] = fmt.Sprintf("%s/router", pluginBasePath)

			// Request 路径
			paths["request"] = fmt.Sprintf("%s/model/request", pluginBasePath)

			// Response 路径
			paths["response"] = fmt.Sprintf("%s/model/response", pluginBasePath)

			// Plugin 特有文件
			paths["plugin_main"] = fmt.Sprintf("%s/main.go", pluginBasePath)
			paths["plugin_config"] = fmt.Sprintf("%s/plugin.go", pluginBasePath)
			paths["plugin_initialize"] = fmt.Sprintf("%s/initialize", pluginBasePath)
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

		// Vue 页面路径
		paths["vue_page"] = fmt.Sprintf("%s/view/%s", webBasePath, packageName)

		// API 路径
		paths["vue_api"] = fmt.Sprintf("%s/api/%s", webBasePath, packageName)
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

// validateExecutionPlan 验证执行计划的完整性
func (t *AutomationModuleAnalyzer) validateExecutionPlan(plan *ExecutionPlan) error {
	// 验证基本字段
	if plan.PackageName == "" {
		return errors.New("packageName 不能为空")
	}
	if plan.ModuleName == "" {
		return errors.New("moduleName 不能为空")
	}
	if plan.PackageType != "package" && plan.PackageType != "plugin" {
		return errors.New("packageType 必须是 'package' 或 'plugin'")
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

	// 验证模块信息
	if plan.NeedCreatedModules {
		if plan.ModulesInfo == nil {
			return errors.New("当 needCreatedModules=true 时，modulesInfo 不能为空")
		}
		if plan.ModulesInfo.Package == "" {
			return errors.New("modulesInfo.package 不能为空")
		}
		if plan.ModulesInfo.StructName == "" {
			return errors.New("modulesInfo.structName 不能为空")
		}
		if plan.ModulesInfo.TableName == "" {
			return errors.New("modulesInfo.tableName 不能为空")
		}
		if plan.ModulesInfo.Description == "" {
			return errors.New("modulesInfo.description 不能为空")
		}
		if plan.ModulesInfo.Abbreviation == "" {
			return errors.New("modulesInfo.abbreviation 不能为空")
		}
		if plan.ModulesInfo.PackageName == "" {
			return errors.New("modulesInfo.packageName 不能为空")
		}
		if plan.ModulesInfo.HumpPackageName == "" {
			return errors.New("modulesInfo.humpPackageName 不能为空")
		}

		// 验证字段信息
		if len(plan.ModulesInfo.Fields) == 0 {
			return errors.New("modulesInfo.fields 不能为空，至少需要一个字段")
		}

		for i, field := range plan.ModulesInfo.Fields {
			if field.FieldName == "" {
				return fmt.Errorf("字段 %d 的 fieldName 不能为空", i+1)
			}
			if field.FieldDesc == "" {
				return fmt.Errorf("字段 %d 的 fieldDesc 不能为空", i+1)
			}
			if field.FieldType == "" {
				return fmt.Errorf("字段 %d 的 fieldType 不能为空", i+1)
			}
			if field.FieldJson == "" {
				return fmt.Errorf("字段 %d 的 fieldJson 不能为空", i+1)
			}
			if field.ColumnName == "" {
				return fmt.Errorf("字段 %d 的 columnName 不能为空", i+1)
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
				return fmt.Errorf("字段 %d 的 fieldType '%s' 不支持，支持的类型：%v", i+1, field.FieldType, validFieldTypes)
			}

			// 验证搜索类型（如果设置了）
			if field.FieldSearchType != "" {
				validSearchTypes := []string{"EQ", "NE", "GT", "GE", "LT", "LE", "LIKE", "BETWEEN"}
				validSearchType := false
				for _, validType := range validSearchTypes {
					if field.FieldSearchType == validType {
						validSearchType = true
						break
					}
				}
				if !validSearchType {
					return fmt.Errorf("字段 %d 的 fieldSearchType '%s' 不支持，支持的类型：%v", i+1, field.FieldSearchType, validSearchTypes)
				}
			}
		}

		// 验证主键设置
		if !plan.ModulesInfo.GvaModel {
			// 当不使用GVA模型时，必须有且仅有一个字段设置为主键
			primaryKeyCount := 0
			for _, field := range plan.ModulesInfo.Fields {
				if field.PrimaryKey {
					primaryKeyCount++
				}
			}
			if primaryKeyCount == 0 {
				return errors.New("当 gvaModel=false 时，必须有一个字段的 primaryKey=true")
			}
			if primaryKeyCount > 1 {
				return errors.New("当 gvaModel=false 时，只能有一个字段的 primaryKey=true")
			}
		} else {
			// 当使用GVA模型时，所有字段的primaryKey都应该为false
			for i, field := range plan.ModulesInfo.Fields {
				if field.PrimaryKey {
					return fmt.Errorf("当 gvaModel=true 时，字段 %d 的 primaryKey 应该为 false，系统会自动创建ID主键", i+1)
				}
			}
		}
	}

	return nil
}

// executeCreation 执行创建操作
func (t *AutomationModuleAnalyzer) executeCreation(ctx context.Context, plan *ExecutionPlan) *ExecutionResult {
	result := &ExecutionResult{
		Success: false,
		Paths:   make(map[string]string),
	}

	// 无论如何都先构建目录结构信息，确保paths始终返回
	result.Paths = t.buildDirectoryStructure(plan)

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

	// 创建字典（如果需要）
	if plan.NeedCreatedModules && plan.ModulesInfo != nil {
		dictResult := t.createRequiredDictionaries(ctx, plan.ModulesInfo)
		result.Message += dictResult
	}

	// 创建模块（如果需要）
	if plan.NeedCreatedModules && plan.ModulesInfo != nil {
		templateService := service.ServiceGroupApp.SystemServiceGroup.AutoCodeTemplate

		err := plan.ModulesInfo.Pretreatment()
		if err != nil {
			result.Message += fmt.Sprintf("模块信息预处理失败: %v", err)
			// 即使预处理失败，也要返回paths信息
			return result
		}

		err = templateService.Create(ctx, *plan.ModulesInfo)
		if err != nil {
			result.Message += fmt.Sprintf("创建模块失败: %v", err)
			// 即使创建模块失败，也要返回paths信息
			return result
		}
		result.Message += "模块创建成功; "
	}

	result.Message += "已构建目录结构信息; "
	result.Success = true

	if result.Message == "" {
		result.Message = "执行计划完成"
	}

	return result
}

// createRequiredDictionaries 创建所需的字典
func (t *AutomationModuleAnalyzer) createRequiredDictionaries(ctx context.Context, modulesInfo *request.AutoCode) string {
	var messages []string
	dictionaryService := service.ServiceGroupApp.SystemServiceGroup.DictionaryService

	// 遍历所有字段，查找使用字典的字段
	for _, field := range modulesInfo.Fields {
		if field.DictType != "" {
			// 检查字典是否存在
			exists, err := t.checkDictionaryExists(field.DictType)
			if err != nil {
				messages = append(messages, fmt.Sprintf("检查字典 %s 时出错: %v; ", field.DictType, err))
				continue
			}

			if !exists {
				// 字典不存在，创建字典
				dictionary := model.SysDictionary{
					Name:   t.generateDictionaryName(field.DictType, field.FieldDesc),
					Type:   field.DictType,
					Status: &[]bool{true}[0], // 默认启用
					Desc:   fmt.Sprintf("自动生成的字典，用于字段: %s (%s)", field.FieldName, field.FieldDesc),
				}

				err = dictionaryService.CreateSysDictionary(dictionary)
				if err != nil {
					messages = append(messages, fmt.Sprintf("创建字典 %s 失败: %v; ", field.DictType, err))
				} else {
					messages = append(messages, fmt.Sprintf("成功创建字典 %s (%s); ", field.DictType, dictionary.Name))

					// 创建默认的字典详情项
					t.createDefaultDictionaryDetails(ctx, field.DictType, field.FieldDesc)
				}
			} else {
				messages = append(messages, fmt.Sprintf("字典 %s 已存在，跳过创建; ", field.DictType))
			}
		}
	}

	if len(messages) == 0 {
		return "未发现需要创建的字典; "
	}

	return strings.Join(messages, "")
}

// checkDictionaryExists 检查字典是否存在
func (t *AutomationModuleAnalyzer) checkDictionaryExists(dictType string) (bool, error) {
	var dictionary model.SysDictionary
	err := global.GVA_DB.Where("type = ?", dictType).First(&dictionary).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil // 字典不存在
		}
		return false, err // 其他错误
	}
	return true, nil // 字典存在
}

// generateDictionaryName 生成字典名称
func (t *AutomationModuleAnalyzer) generateDictionaryName(dictType, fieldDesc string) string {
	if fieldDesc != "" {
		return fmt.Sprintf("%s字典", fieldDesc)
	}
	return fmt.Sprintf("%s字典", dictType)
}

// createDefaultDictionaryDetails 创建默认的字典详情项
func (t *AutomationModuleAnalyzer) createDefaultDictionaryDetails(ctx context.Context, dictType, fieldDesc string) {
	dictionaryDetailService := service.ServiceGroupApp.SystemServiceGroup.DictionaryDetailService

	// 获取刚创建的字典ID
	var dictionary model.SysDictionary
	err := global.GVA_DB.Where("type = ?", dictType).First(&dictionary).Error
	if err != nil {
		return
	}

	// 根据字典类型和字段描述智能生成默认选项
	defaultDetails := t.generateSmartDictionaryOptions(dictType, fieldDesc)

	for _, detail := range defaultDetails {
		dictionaryDetail := model.SysDictionaryDetail{
			Label:           detail.label,
			Value:           detail.value,
			Status:          &[]bool{true}[0], // 默认启用
			Sort:            detail.sort,
			SysDictionaryID: int(dictionary.ID),
		}

		// 忽略创建详情项的错误，因为这只是默认值
		dictionaryDetailService.CreateSysDictionaryDetail(dictionaryDetail)
	}
}

// generateSmartDictionaryOptions 根据字典类型和字段描述智能生成默认选项
func (t *AutomationModuleAnalyzer) generateSmartDictionaryOptions(dictType, fieldDesc string) []struct {
	label string
	value string
	sort  int
} {
	// 转换为小写进行匹配
	lowerDictType := strings.ToLower(dictType)
	lowerFieldDesc := strings.ToLower(fieldDesc)
	combinedText := lowerDictType + " " + lowerFieldDesc

	// 根据字典类型和字段描述的关键词生成相应的选项
	switch {
	case strings.Contains(combinedText, "status") || strings.Contains(combinedText, "状态"):
		return []struct {
			label string
			value string
			sort  int
		}{
			{"启用", "1", 1},
			{"禁用", "0", 2},
		}
	case strings.Contains(combinedText, "gender") || strings.Contains(combinedText, "性别"):
		return []struct {
			label string
			value string
			sort  int
		}{
			{"男", "1", 1},
			{"女", "2", 2},
			{"未知", "0", 3},
		}
	case strings.Contains(combinedText, "type") || strings.Contains(combinedText, "类型"):
		return []struct {
			label string
			value string
			sort  int
		}{
			{"类型一", "1", 1},
			{"类型二", "2", 2},
			{"类型三", "3", 3},
		}
	case strings.Contains(combinedText, "level") || strings.Contains(combinedText, "等级") || strings.Contains(combinedText, "级别"):
		return []struct {
			label string
			value string
			sort  int
		}{
			{"初级", "1", 1},
			{"中级", "2", 2},
			{"高级", "3", 3},
		}
	case strings.Contains(combinedText, "priority") || strings.Contains(combinedText, "优先级"):
		return []struct {
			label string
			value string
			sort  int
		}{
			{"低", "1", 1},
			{"中", "2", 2},
			{"高", "3", 3},
			{"紧急", "4", 4},
		}
	case strings.Contains(combinedText, "approve") || strings.Contains(combinedText, "审批") || strings.Contains(combinedText, "审核"):
		return []struct {
			label string
			value string
			sort  int
		}{
			{"待审核", "0", 1},
			{"已通过", "1", 2},
			{"已拒绝", "2", 3},
		}
	case strings.Contains(combinedText, "role") || strings.Contains(combinedText, "角色"):
		return []struct {
			label string
			value string
			sort  int
		}{
			{"管理员", "admin", 1},
			{"用户", "user", 2},
			{"访客", "guest", 3},
		}
	case strings.Contains(combinedText, "bool") || strings.Contains(combinedText, "boolean") || strings.Contains(combinedText, "是否"):
		return []struct {
			label string
			value string
			sort  int
		}{
			{"是", "true", 1},
			{"否", "false", 2},
		}
	case strings.Contains(combinedText, "order") || strings.Contains(combinedText, "订单"):
		return []struct {
			label string
			value string
			sort  int
		}{
			{"待付款", "1", 1},
			{"已付款", "2", 2},
			{"已发货", "3", 3},
			{"已完成", "4", 4},
			{"已取消", "0", 5},
		}
	case strings.Contains(combinedText, "color") || strings.Contains(combinedText, "颜色"):
		return []struct {
			label string
			value string
			sort  int
		}{
			{"红色", "red", 1},
			{"绿色", "green", 2},
			{"蓝色", "blue", 3},
		}
	case strings.Contains(combinedText, "size") || strings.Contains(combinedText, "尺寸") || strings.Contains(combinedText, "大小"):
		return []struct {
			label string
			value string
			sort  int
		}{
			{"小", "S", 1},
			{"中", "M", 2},
			{"大", "L", 3},
			{"特大", "XL", 4},
		}
	default:
		// 默认选项，使用通用的选项
		return []struct {
			label string
			value string
			sort  int
		}{
			{"选项一", "1", 1},
			{"选项二", "2", 2},
			{"选项三", "3", 3},
		}
	}
}
