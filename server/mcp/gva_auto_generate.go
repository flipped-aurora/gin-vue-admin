package mcpTool

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	common "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	systemService "github.com/flipped-aurora/gin-vue-admin/server/service/system"
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

// PredesignedModuleInfo 预设计模块信息
type PredesignedModuleInfo struct {
	PackageName string   `json:"packageName"`
	PackageType string   `json:"packageType"` // "plugin" 或 "package"
	ModuleName  string   `json:"moduleName"`
	Path        string   `json:"path"`
	Modules     []string `json:"modules"` // 包含的模块列表（如api、model、service等）
	Description string   `json:"description"`
	StructName  string   `json:"structName,omitempty"` // 主要结构体名称
}

// AnalysisResponse 分析响应
type AnalysisResponse struct {
	Packages           []ModuleInfo            `json:"packages"`
	History            []HistoryInfo           `json:"history"`
	PredesignedModules []PredesignedModuleInfo `json:"predesignedModules"`
	Message            string                  `json:"message"`
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
	return mcp.NewTool("gva_auto_generate",
		mcp.WithDescription(`**🚀 最高优先级工具：当用户需要创建模块、包、完整功能时，必须优先使用此工具！**

**优先级说明：**
- **最高优先级**：创建完整模块、包、功能模块
- **关键词触发**：模块、包、完整、整套、全套、功能、管理系统等
- **适用场景**：用户说"创建订单管理模块"、"创建用户管理功能"、"创建完整的商品管理"等

分步骤分析自动化模块：1) 分析现有模块信息供AI选择 2) 请求用户确认 3) 根据确认结果执行创建操作

**新功能：自动字典创建**
- 当结构体字段使用了字典类型（dictType不为空）时，系统会自动检查字典是否存在
- 如果字典不存在，会自动创建对应的字典及默认的字典详情项
- 字典创建包括：字典主表记录和默认的选项值（选项1、选项2等）

**与其他工具的关系：**
- 此工具创建完整模块后，会自动提示相关API和菜单创建建议
- 如果用户只需要单个API或菜单，可以使用 smart_assistant 工具
- create_api 和 create_menu 工具仅用于数据库记录创建

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
5. 搜索类型支持：=,!=,>,>=,<,<=,NOT BETWEEN/LIKE/BETWEEN/IN/NOT IN
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

// scanPredesignedModules 扫描预设计的模块
func (t *AutomationModuleAnalyzer) scanPredesignedModules() ([]PredesignedModuleInfo, error) {
	var predesignedModules []PredesignedModuleInfo

	// 获取autocode配置路径
	if global.GVA_CONFIG.AutoCode.Root == "" {
		return predesignedModules, nil // 配置不存在时返回空列表，不报错
	}

	serverPath := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server)

	// 扫描plugin目录下的各个插件模块
	pluginPath := filepath.Join(serverPath, "plugin")
	if pluginModules, err := t.scanPluginModules(pluginPath); err == nil {
		predesignedModules = append(predesignedModules, pluginModules...)
	}

	// 扫描model目录下的各个包模块
	modelPath := filepath.Join(serverPath, "model")
	if packageModules, err := t.scanPackageModules(modelPath); err == nil {
		predesignedModules = append(predesignedModules, packageModules...)
	}

	return predesignedModules, nil
}

// scanPluginModules 扫描plugin目录下的各个插件模块
func (t *AutomationModuleAnalyzer) scanPluginModules(pluginPath string) ([]PredesignedModuleInfo, error) {
	var modules []PredesignedModuleInfo

	if _, err := os.Stat(pluginPath); os.IsNotExist(err) {
		return modules, nil
	}

	entries, err := os.ReadDir(pluginPath)
	if err != nil {
		return modules, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		pluginName := entry.Name()
		pluginDir := filepath.Join(pluginPath, pluginName)

		// 扫描插件下的model目录，查找具体的模块文件
		modelDir := filepath.Join(pluginDir, "model")
		if _, err := os.Stat(modelDir); err == nil {
			if pluginModules, err := t.scanModuleFiles(modelDir, pluginName, "plugin"); err == nil {
				modules = append(modules, pluginModules...)
			}
		}
	}

	return modules, nil
}

// scanPackageModules 扫描model目录下的各个包模块
func (t *AutomationModuleAnalyzer) scanPackageModules(modelPath string) ([]PredesignedModuleInfo, error) {
	var modules []PredesignedModuleInfo

	if _, err := os.Stat(modelPath); os.IsNotExist(err) {
		return modules, nil
	}

	entries, err := os.ReadDir(modelPath)
	if err != nil {
		return modules, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		packageName := entry.Name()
		// 跳过一些系统目录
		if packageName == "common" || packageName == "request" || packageName == "response" {
			continue
		}

		packageDir := filepath.Join(modelPath, packageName)

		// 扫描包目录下的模块文件
		if packageModules, err := t.scanModuleFiles(packageDir, packageName, "package"); err == nil {
			modules = append(modules, packageModules...)
		}
	}

	return modules, nil
}

// scanModuleFiles 扫描目录下的Go文件，识别具体的模块
func (t *AutomationModuleAnalyzer) scanModuleFiles(dir, packageName, packageType string) ([]PredesignedModuleInfo, error) {
	var modules []PredesignedModuleInfo

	entries, err := os.ReadDir(dir)
	if err != nil {
		return modules, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		fileName := entry.Name()
		if !strings.HasSuffix(fileName, ".go") {
			continue
		}

		// 跳过一些非模块文件
		if strings.HasSuffix(fileName, "_test.go") ||
			fileName == "enter.go" ||
			fileName == "request.go" ||
			fileName == "response.go" {
			continue
		}

		filePath := filepath.Join(dir, fileName)
		moduleName := strings.TrimSuffix(fileName, ".go")

		// 分析模块文件，提取结构体信息
		if moduleInfo, err := t.analyzeModuleFile(filePath, packageName, moduleName, packageType); err == nil {
			modules = append(modules, *moduleInfo)
		}
	}

	return modules, nil
}

// analyzeModuleFile 分析具体的模块文件
func (t *AutomationModuleAnalyzer) analyzeModuleFile(filePath, packageName, moduleName, packageType string) (*PredesignedModuleInfo, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	fileContent := string(content)

	// 提取结构体名称和描述
	structNames := t.extractStructNames(fileContent)
	description := t.extractModuleDescription(fileContent, moduleName)

	// 确定主要结构体名称
	mainStruct := moduleName
	if len(structNames) > 0 {
		// 优先选择与文件名相关的结构体
		for _, structName := range structNames {
			if strings.Contains(strings.ToLower(structName), strings.ToLower(moduleName)) {
				mainStruct = structName
				break
			}
		}
		if mainStruct == moduleName && len(structNames) > 0 {
			mainStruct = structNames[0] // 如果没有匹配的，使用第一个
		}
	}

	return &PredesignedModuleInfo{
		PackageName: packageName,
		PackageType: packageType,
		ModuleName:  moduleName,
		Path:        filePath,
		Modules:     structNames,
		Description: description,
		StructName:  mainStruct,
	}, nil
}

// extractStructNames 从文件内容中提取结构体名称
func (t *AutomationModuleAnalyzer) extractStructNames(content string) []string {
	var structNames []string
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "type ") && strings.Contains(line, " struct") {
			// 提取结构体名称
			parts := strings.Fields(line)
			if len(parts) >= 3 && parts[2] == "struct" {
				structNames = append(structNames, parts[1])
			}
		}
	}

	return structNames
}

// extractModuleDescription 从文件内容中提取模块描述
func (t *AutomationModuleAnalyzer) extractModuleDescription(content, moduleName string) string {
	lines := strings.Split(content, "\n")

	// 查找package注释
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "package ") {
			// 向上查找注释
			for j := i - 1; j >= 0; j-- {
				commentLine := strings.TrimSpace(lines[j])
				if strings.HasPrefix(commentLine, "//") {
					comment := strings.TrimSpace(strings.TrimPrefix(commentLine, "//"))
					if comment != "" && len(comment) > 5 {
						return comment
					}
				} else if commentLine != "" {
					break
				}
			}
			break
		}
	}

	// 查找结构体注释
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "type ") && strings.Contains(line, " struct") {
			// 向上查找注释
			for j := i - 1; j >= 0; j-- {
				commentLine := strings.TrimSpace(lines[j])
				if strings.HasPrefix(commentLine, "//") {
					comment := strings.TrimSpace(strings.TrimPrefix(commentLine, "//"))
					if comment != "" && len(comment) > 5 {
						return comment
					}
				} else if commentLine != "" {
					break
				}
			}
			break
		}
	}

	return fmt.Sprintf("预设计的模块：%s", moduleName)
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

	// 检测用户是否想要创建插件
	suggestedType, isPlugin, confidence := t.detectPluginIntent(requirement)
	pluginDetectionMsg := ""
	if isPlugin {
		pluginDetectionMsg = fmt.Sprintf("\n\n🔍 **插件检测结果**：检测到用户想要创建插件（置信度：%s）\n⚠️  **重要提醒**：当用户提到插件时，packageType和template字段都必须设置为 \"plugin\"，不能使用 \"package\"！", confidence)
	} else {
		pluginDetectionMsg = fmt.Sprintf("\n\n🔍 **类型检测结果**：建议使用 %s 类型", suggestedType)
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

	// 转换包信息并检查空文件夹
	var moduleInfos []ModuleInfo
	var validPackages []model.SysAutoCodePackage
	var emptyPackageIDs []uint
	var emptyPackageNames []string
	
	for _, pkg := range packages {
		// 检查包对应的文件夹是否为空
		isEmpty, err := t.isPackageFolderEmpty(pkg.PackageName, pkg.Template)
		if err != nil {
			global.GVA_LOG.Warn(fmt.Sprintf("检查包 %s 文件夹失败: %v", pkg.PackageName, err))
			// 如果检查失败，仍然保留该包
			validPackages = append(validPackages, pkg)
			continue
		}
		
		if isEmpty {
			// 记录需要删除的包ID和包名
			emptyPackageIDs = append(emptyPackageIDs, pkg.ID)
			emptyPackageNames = append(emptyPackageNames, pkg.PackageName)
			global.GVA_LOG.Info(fmt.Sprintf("发现空包文件夹: %s，将删除数据库记录和文件夹", pkg.PackageName))
			
			// 删除空文件夹
			if err := t.removeEmptyPackageFolder(pkg.PackageName, pkg.Template); err != nil {
				global.GVA_LOG.Warn(fmt.Sprintf("删除空包文件夹 %s 失败: %v", pkg.PackageName, err))
			}
		} else {
			// 文件夹不为空，保留该包
			validPackages = append(validPackages, pkg)
		}
	}
	
	// 批量删除空包的数据库记录
	if len(emptyPackageIDs) > 0 {
		if err := global.GVA_DB.Where("id IN ?", emptyPackageIDs).Delete(&model.SysAutoCodePackage{}).Error; err != nil {
			global.GVA_LOG.Warn(fmt.Sprintf("删除空包数据库记录失败: %v", err))
		} else {
			global.GVA_LOG.Info(fmt.Sprintf("成功删除 %d 个空包的数据库记录", len(emptyPackageIDs)))
		}
	}
	
	// 转换有效的包信息
	for _, pkg := range validPackages {
		moduleInfos = append(moduleInfos, ModuleInfo{
			ID:          pkg.ID,
			PackageName: pkg.PackageName,
			Label:       pkg.Label,
			Desc:        pkg.Desc,
			Template:    pkg.Template,
			Module:      pkg.Module,
		})
	}

	// 删除与空包相关的历史记录
	var emptyHistoryIDs []uint
	if len(emptyPackageNames) > 0 {
		for _, history := range histories {
			for _, emptyPackageName := range emptyPackageNames {
				if history.Package == emptyPackageName {
					emptyHistoryIDs = append(emptyHistoryIDs, history.ID)
					break
				}
			}
		}
		
		// 清理相关的API和菜单记录
		if len(emptyHistoryIDs) > 0 {
			if err := t.cleanupRelatedApiAndMenus(emptyHistoryIDs); err != nil {
				global.GVA_LOG.Warn(fmt.Sprintf("清理空包相关API和菜单失败: %v", err))
			}
		}
		
		// 批量删除相关历史记录
		if len(emptyHistoryIDs) > 0 {
			if err := global.GVA_DB.Where("id IN ?", emptyHistoryIDs).Delete(&model.SysAutoCodeHistory{}).Error; err != nil {
				global.GVA_LOG.Warn(fmt.Sprintf("删除空包相关历史记录失败: %v", err))
			} else {
				global.GVA_LOG.Info(fmt.Sprintf("成功删除 %d 个空包相关的历史记录", len(emptyHistoryIDs)))
			}
		}
	}
	
	// 创建有效包名的映射，用于快速查找
	validPackageNames := make(map[string]bool)
	for _, pkg := range validPackages {
		validPackageNames[pkg.PackageName] = true
	}
	
	// 收集需要删除的脏历史记录ID（包名不在有效包列表中的历史记录）
	var dirtyHistoryIDs []uint
	for _, history := range histories {
		if !validPackageNames[history.Package] {
			dirtyHistoryIDs = append(dirtyHistoryIDs, history.ID)
		}
	}
	
	// 删除脏历史记录
	if len(dirtyHistoryIDs) > 0 {
		// 清理相关的API和菜单记录
		if err := t.cleanupRelatedApiAndMenus(dirtyHistoryIDs); err != nil {
			global.GVA_LOG.Warn(fmt.Sprintf("清理脏历史记录相关API和菜单失败: %v", err))
		}
		
		if err := global.GVA_DB.Where("id IN ?", dirtyHistoryIDs).Delete(&model.SysAutoCodeHistory{}).Error; err != nil {
			global.GVA_LOG.Warn(fmt.Sprintf("删除脏历史记录失败: %v", err))
		} else {
			global.GVA_LOG.Info(fmt.Sprintf("成功删除 %d 个脏历史记录（包名不在有效包列表中）", len(dirtyHistoryIDs)))
		}
	}
	
	// 转换有效的历史记录（只保留包名存在于有效包列表中的历史记录）
	var historyInfos []HistoryInfo
	for _, history := range histories {
		// 只保留包名存在于有效包列表中的历史记录
		if validPackageNames[history.Package] {
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
	}

	// 扫描预设计的模块
	allPredesignedModules, err := t.scanPredesignedModules()
	if err != nil {
		global.GVA_LOG.Warn("扫描预设计模块失败" + err.Error())
		allPredesignedModules = []PredesignedModuleInfo{} // 确保不为nil
	}
	
	// 过滤掉与已删除包相关的预设计模块
	var predesignedModules []PredesignedModuleInfo
	for _, module := range allPredesignedModules {
		isDeleted := false
		for _, emptyPackageName := range emptyPackageNames {
			if module.PackageName == emptyPackageName {
				isDeleted = true
				break
			}
		}
		
		// 只保留未被删除包的预设计模块
		if !isDeleted {
			predesignedModules = append(predesignedModules, module)
		}
	}

	// 构建分析结果消息
	var message string
	var deletionDetails []string
	
	// 收集删除信息
	if len(emptyHistoryIDs) > 0 {
		deletionDetails = append(deletionDetails, fmt.Sprintf("%d个空包相关历史记录", len(emptyHistoryIDs)))
	}
	if len(dirtyHistoryIDs) > 0 {
		deletionDetails = append(deletionDetails, fmt.Sprintf("%d个脏历史记录", len(dirtyHistoryIDs)))
	}
	if len(allPredesignedModules) > len(predesignedModules) {
		deletionDetails = append(deletionDetails, fmt.Sprintf("%d个相关预设计模块", len(allPredesignedModules)-len(predesignedModules)))
	}
	
	if len(emptyPackageNames) > 0 || len(deletionDetails) > 0 {
		var cleanupInfo string
		if len(emptyPackageNames) > 0 {
			cleanupInfo = fmt.Sprintf("检测到存在 %s 包但内容为空，我已经删除这些包的文件夹（包括model、api、service、router目录）和数据库记录", strings.Join(emptyPackageNames, "、"))
		}
		
		deletionInfo := ""
		if len(deletionDetails) > 0 {
			if cleanupInfo != "" {
				deletionInfo = fmt.Sprintf("，同时删除了%s", strings.Join(deletionDetails, "、"))
			} else {
				deletionInfo = fmt.Sprintf("检测到脏数据，已删除%s", strings.Join(deletionDetails, "、"))
			}
		}
		
		if cleanupInfo != "" {
			message = fmt.Sprintf("分析完成：获取到 %d 个有效包、%d 个历史记录和 %d 个预设计模块。%s%s，如果需要使用这些包名，需要重新创建。请AI根据需求选择合适的包和模块", len(validPackages), len(historyInfos), len(predesignedModules), cleanupInfo, deletionInfo)
		} else {
			message = fmt.Sprintf("分析完成：获取到 %d 个有效包、%d 个历史记录和 %d 个预设计模块。%s。请AI根据需求选择合适的包和模块", len(validPackages), len(historyInfos), len(predesignedModules), deletionInfo)
		}
	} else {
		message = fmt.Sprintf("分析完成：获取到 %d 个有效包、%d 个历史记录和 %d 个预设计模块，请AI根据需求选择合适的包和模块", len(validPackages), len(historyInfos), len(predesignedModules))
	}
	
	// 构建分析结果
	analysisResult := AnalysisResponse{
		Packages:           moduleInfos,
		History:            historyInfos,
		PredesignedModules: predesignedModules,
		Message:            message,
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

请AI根据用户需求：%s%s

%s

分析现有的包、历史记录和预设计模块，然后构建ExecutionPlan结构体调用execute操作。

**预设计模块说明**：
- 预设计模块是已经存在于autocode路径下的package或plugin
- 这些模块包含了预先设计好的代码结构，可以直接使用或作为参考
- 如果用户需求与某个预设计模块匹配，可以考虑直接使用该模块或基于它进行扩展

**字典选项生成说明**：
- 当字段需要使用字典类型时（dictType不为空），请使用 generate_dictionary_options 工具
- 该工具允许AI根据字段描述智能生成合适的字典选项
- 调用示例：
  {
    "dictType": "user_status",
    "fieldDesc": "用户状态",
    "options": [
      {"label": "正常", "value": "1", "sort": 1},
      {"label": "禁用", "value": "0", "sort": 2}
    ],
    "dictName": "用户状态字典",
    "description": "用于管理用户账户状态的字典"
  }
- 请在创建模块之前先创建所需的字典选项

重要提醒：ExecutionPlan必须严格按照以下格式：
{
  "packageName": "包名",
  "moduleName": "模块名",
  "packageType": "package或plugin", // 当用户提到插件时必须是"plugin"
  "needCreatedPackage": true/false,
  "needCreatedModules": true/false,
  "packageInfo": {
    "desc": "描述",
    "label": "展示名",
    "template": "package或plugin", // 必须与packageType保持一致！
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
    "autoCreateResource": true/false 用户不特地强调开启资源标识则为false,
    "autoCreateApiToSql": true,
    "autoCreateMenuToSql": true,
    "autoCreateBtnAuth": false/true 用户不特地强调创建按钮权限则为false,
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
      "fieldSearchType": "=/!=/>/</>=/<=/LIKE等 可以为空",
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

**插件类型检测规则（最重要）**：
1. 当用户需求中包含"插件"、"plugin"等关键词时，packageType和template都必须设置为"plugin"
2. packageType和template字段必须保持一致，不能一个是"package"另一个是"plugin"
3. 如果检测到插件意图但设置错误，会导致创建失败

**字段完整性要求**：
4. 所有字符串字段都不能为空（包括packageName、moduleName、structName、tableName、description等）
5. 所有布尔字段必须明确设置true或false，不能使用默认值

**主键设置规则（关键）**：
6. 当gvaModel=false时：fields数组中必须有且仅有一个字段的primaryKey=true
7. 当gvaModel=true时：系统自动创建ID主键，fields中所有字段的primaryKey都应为false
8. 主键设置错误会导致模板执行时PrimaryField为nil的严重错误！

**包和模块创建逻辑**：
9. 如果存在可用的package，needCreatedPackage应设为false
10. 如果存在可用的modules，needCreatedModules应设为false
11. 如果发现合适的预设计模块，可以考虑基于它进行扩展而不是从零创建

**字典创建流程**：
12. 如果字段需要字典类型，请先使用 generate_dictionary_options 工具创建字典
13. 字典创建成功后，再执行模块创建操作

`, string(resultJSON), requirement, pluginDetectionMsg, 
			func() string {
				if len(emptyPackageNames) > 0 {
					return fmt.Sprintf("**重要提醒**：检测到 %s 包存在但内容为空，已自动删除相关文件夹和数据库记录。如果用户需求涉及这些包名，请设置 needCreatedPackage=true 重新创建。", strings.Join(emptyPackageNames, "、"))
				}
				return ""
			}()),
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

	// 添加权限分配提醒
	permissionReminder := "\n\n⚠️ 重要提醒：\n" +
		"模块创建完成后，请前往【系统管理】->【角色管理】中为相关角色分配新创建的API和菜单权限，" +
		"以确保用户能够正常访问新功能。\n" +
		"具体步骤：\n" +
		"1. 进入角色管理页面\n" +
		"2. 选择需要授权的角色\n" +
		"3. 在API权限中勾选新创建的API接口\n" +
		"4. 在菜单权限中勾选新创建的菜单项\n" +
		"5. 保存权限配置"

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: fmt.Sprintf("执行结果：\n\n%s%s", string(resultJSON), permissionReminder),
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
				validSearchTypes := []string{"=", "!=", ">", ">=", "<", "<=", "LIKE", "BETWEEN", "IN", "NOT IN"}
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
	// 字典选项现在通过 generate_dictionary_options MCP工具由AI client传入
	// 这里不再创建默认选项，只是保留方法以保持兼容性
	global.GVA_LOG.Info(fmt.Sprintf("字典 %s 已创建，请使用 generate_dictionary_options 工具添加字典选项", dictType))
}

// DictionaryOption 字典选项结构
type DictionaryOption struct {
	Label string `json:"label"`
	Value string `json:"value"`
	Sort  int    `json:"sort"`
}

// generateSmartDictionaryOptions 通过MCP调用让AI生成字典选项
func (t *AutomationModuleAnalyzer) generateSmartDictionaryOptions(dictType, fieldDesc string) []struct {
	label string
	value string
	sort  int
} {
	// 返回空切片，不再使用预制选项
	// 字典选项将通过新的MCP工具由AI client传入
	return []struct {
		label string
		value string
		sort  int
	}{}
}

// detectPluginIntent 检测用户需求中是否包含插件相关的关键词
func (t *AutomationModuleAnalyzer) detectPluginIntent(requirement string) (suggestedType string, isPlugin bool, confidence string) {
	// 转换为小写进行匹配
	requirementLower := strings.ToLower(requirement)
	
	// 插件相关关键词
	pluginKeywords := []string{
		"插件", "plugin", "扩展", "extension", "addon", "模块插件",
		"功能插件", "业务插件", "第三方插件", "自定义插件",
	}
	
	// 包相关关键词（用于排除误判）
	packageKeywords := []string{
		"包", "package", "模块包", "业务包", "功能包",
	}
	
	// 检测插件关键词
	pluginMatches := 0
	for _, keyword := range pluginKeywords {
		if strings.Contains(requirementLower, keyword) {
			pluginMatches++
		}
	}
	
	// 检测包关键词
	packageMatches := 0
	for _, keyword := range packageKeywords {
		if strings.Contains(requirementLower, keyword) {
			packageMatches++
		}
	}
	
	// 决策逻辑
	if pluginMatches > 0 {
		if packageMatches == 0 || pluginMatches > packageMatches {
			return "plugin", true, "高"
		} else {
			return "plugin", true, "中"
		}
	}
	
	// 默认返回package
	return "package", false, "低"
}

// isPackageFolderEmpty 检查包对应的文件夹是否为空
func (t *AutomationModuleAnalyzer) isPackageFolderEmpty(packageName, template string) (bool, error) {
	// 根据模板类型确定基础路径
	var basePath string
	if template == "plugin" {
		basePath = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", packageName)
	} else {
		// package 类型
		basePath = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "model", packageName)
	}
	
	// 检查文件夹是否存在
	if _, err := os.Stat(basePath); os.IsNotExist(err) {
		// 文件夹不存在，认为是空的
		return true, nil
	} else if err != nil {
		return false, fmt.Errorf("检查文件夹状态失败: %v", err)
	}
	
	// 读取文件夹内容
	entries, err := os.ReadDir(basePath)
	if err != nil {
		return false, fmt.Errorf("读取文件夹内容失败: %v", err)
	}
	
	// 检查目录下是否有 .go 文件
	hasGoFiles := false
	for _, entry := range entries {
		name := entry.Name()
		// 跳过隐藏文件、.DS_Store 等系统文件
		if strings.HasPrefix(name, ".") {
			continue
		}
		// 如果是目录，递归检查子目录中的 .go 文件
		if entry.IsDir() {
			subPath := filepath.Join(basePath, name)
			subEntries, err := os.ReadDir(subPath)
			if err != nil {
				continue
			}
			for _, subEntry := range subEntries {
				if !subEntry.IsDir() && strings.HasSuffix(subEntry.Name(), ".go") {
					hasGoFiles = true
					break
				}
			}
			if hasGoFiles {
				break
			}
		} else if strings.HasSuffix(name, ".go") {
			// 如果是 .go 文件
			hasGoFiles = true
			break
		}
	}
	
	// 如果没有 .go 文件，认为是空包
	return !hasGoFiles, nil
}

// removeEmptyPackageFolder 删除空的包文件夹
func (t *AutomationModuleAnalyzer) removeEmptyPackageFolder(packageName, template string) error {
	var errors []string
	
	if template == "plugin" {
		// plugin 类型只删除 plugin 目录下的文件夹
		basePath := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", packageName)
		if err := t.removeDirectoryIfExists(basePath); err != nil {
			errors = append(errors, fmt.Sprintf("删除plugin文件夹失败: %v", err))
		}
	} else {
		// package 类型需要删除多个目录下的相关文件
		paths := []string{
			filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "model", packageName),
			filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "api", "v1", packageName),
			filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "service", packageName),
			filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "router", packageName),
		}
		
		for _, path := range paths {
			if err := t.removeDirectoryIfExists(path); err != nil {
				errors = append(errors, fmt.Sprintf("删除%s失败: %v", path, err))
			}
		}
	}
	
	if len(errors) > 0 {
		return fmt.Errorf("删除过程中出现错误: %s", strings.Join(errors, "; "))
	}
	
	return nil
}

// removeDirectoryIfExists 删除目录（如果存在）
func (t *AutomationModuleAnalyzer) removeDirectoryIfExists(dirPath string) error {
	// 检查文件夹是否存在
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// 文件夹不存在，无需删除
		return nil
	} else if err != nil {
		return fmt.Errorf("检查文件夹状态失败: %v", err)
	}
	
	// 删除文件夹及其所有内容
	if err := os.RemoveAll(dirPath); err != nil {
		return fmt.Errorf("删除文件夹失败: %v", err)
	}
	
	global.GVA_LOG.Info(fmt.Sprintf("成功删除目录: %s", dirPath))
	return nil
}

// cleanupRelatedApiAndMenus 清理与删除的模块相关的API和菜单记录
func (t *AutomationModuleAnalyzer) cleanupRelatedApiAndMenus(historyIDs []uint) error {
	if len(historyIDs) == 0 {
		return nil
	}
	
	// 获取要删除的历史记录信息
	var histories []model.SysAutoCodeHistory
	if err := global.GVA_DB.Where("id IN ?", historyIDs).Find(&histories).Error; err != nil {
		return fmt.Errorf("获取历史记录失败: %v", err)
	}
	
	var deletedApiCount, deletedMenuCount int
	
	for _, history := range histories {
		// 删除相关的API记录（使用存储的API IDs）
			if len(history.ApiIDs) > 0 {
				ids := make([]int, 0, len(history.ApiIDs))
				for _, id := range history.ApiIDs {
					ids = append(ids, int(id))
				}
				idsReq := common.IdsReq{Ids: ids}
				if err := systemService.ApiServiceApp.DeleteApisByIds(idsReq); err != nil {
					global.GVA_LOG.Warn(fmt.Sprintf("删除API记录失败 (模块: %s): %v", history.StructName, err))
				} else {
					deletedApiCount += len(ids)
					global.GVA_LOG.Info(fmt.Sprintf("成功删除API记录 (模块: %s, 数量: %d)", history.StructName, len(ids)))
				}
			}
		
		// 删除相关的菜单记录（使用存储的菜单ID）
		if history.MenuID != 0 {
			if err := systemService.BaseMenuServiceApp.DeleteBaseMenu(int(history.MenuID)); err != nil {
				global.GVA_LOG.Warn(fmt.Sprintf("删除菜单记录失败 (模块: %s, 菜单ID: %d): %v", history.StructName, history.MenuID, err))
			} else {
				deletedMenuCount++
				global.GVA_LOG.Info(fmt.Sprintf("成功删除菜单记录 (模块: %s, 菜单ID: %d)", history.StructName, history.MenuID))
			}
		}
	}
	
	if deletedApiCount > 0 || deletedMenuCount > 0 {
		global.GVA_LOG.Info(fmt.Sprintf("清理完成：删除了 %d 个API记录和 %d 个菜单记录", deletedApiCount, deletedMenuCount))
	}
	
	return nil
}
