package mcpTool

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"os"
	"path/filepath"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/mark3labs/mcp-go/mcp"
)

// 注册工具
func init() {
	RegisterTool(&GVAAnalyzer{})
}

// GVAAnalyzer GVA分析器 - 用于分析当前功能是否需要创建独立的package和module
type GVAAnalyzer struct{}

// AnalyzeRequest 分析请求结构体
type AnalyzeRequest struct {
	Requirement string `json:"requirement" binding:"required"` // 用户需求描述
}

// AnalyzeResponse 分析响应结构体
type AnalyzeResponse struct {
	ExistingPackages   []PackageInfo           `json:"existingPackages"`   // 现有包信息
	PredesignedModules []PredesignedModuleInfo `json:"predesignedModules"` // 预设计模块信息
	Dictionaries       []DictionaryPre         `json:"dictionaries"`       // 字典信息
	CleanupInfo        *CleanupInfo            `json:"cleanupInfo"`        // 清理信息（如果有）
}

// ModuleInfo 模块信息
type ModuleInfo struct {
	ModuleName  string   `json:"moduleName"`  // 模块名称
	PackageName string   `json:"packageName"` // 包名
	Template    string   `json:"template"`    // 模板类型
	StructName  string   `json:"structName"`  // 结构体名称
	TableName   string   `json:"tableName"`   // 表名
	Description string   `json:"description"` // 描述
	FilePaths   []string `json:"filePaths"`   // 相关文件路径
}

// PackageInfo 包信息
type PackageInfo struct {
	PackageName string `json:"packageName"` // 包名
	Template    string `json:"template"`    // 模板类型
	Label       string `json:"label"`       // 标签
	Desc        string `json:"desc"`        // 描述
	Module      string `json:"module"`      // 模块
	IsEmpty     bool   `json:"isEmpty"`     // 是否为空包
}

// PredesignedModuleInfo 预设计模块信息
type PredesignedModuleInfo struct {
	ModuleName  string   `json:"moduleName"`  // 模块名称
	PackageName string   `json:"packageName"` // 包名
	Template    string   `json:"template"`    // 模板类型
	FilePaths   []string `json:"filePaths"`   // 文件路径列表
	Description string   `json:"description"` // 描述
}

// CleanupInfo 清理信息
type CleanupInfo struct {
	DeletedPackages []string `json:"deletedPackages"` // 已删除的包
	DeletedModules  []string `json:"deletedModules"`  // 已删除的模块
	CleanupMessage  string   `json:"cleanupMessage"`  // 清理消息
}

// New 创建GVA分析器工具
func (g *GVAAnalyzer) New() mcp.Tool {
	return mcp.NewTool("gva_analyze",
		mcp.WithDescription("返回当前系统中有效的包和模块信息，并分析用户需求是否需要创建新的包、模块和字典。同时检查并清理空包，确保系统整洁。"),
		mcp.WithString("requirement",
			mcp.Description("用户需求描述，用于分析是否需要创建新的包和模块"),
			mcp.Required(),
		),
	)
}

// Handle 处理分析请求
func (g *GVAAnalyzer) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 解析请求参数
	requirementStr, ok := request.GetArguments()["requirement"].(string)
	if !ok || requirementStr == "" {
		return nil, errors.New("参数错误：requirement 必须是非空字符串")
	}

	// 创建分析请求
	analyzeReq := AnalyzeRequest{
		Requirement: requirementStr,
	}

	// 执行分析逻辑
	response, err := g.performAnalysis(ctx, analyzeReq)
	if err != nil {
		return nil, fmt.Errorf("分析失败: %v", err)
	}

	// 序列化响应
	responseJSON, err := json.Marshal(response)
	if err != nil {
		return nil, fmt.Errorf("序列化响应失败: %v", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.NewTextContent(string(responseJSON)),
		},
	}, nil
}

// performAnalysis 执行分析逻辑
func (g *GVAAnalyzer) performAnalysis(ctx context.Context, req AnalyzeRequest) (*AnalyzeResponse, error) {
	// 1. 获取数据库中的包信息
	var packages []model.SysAutoCodePackage
	if err := global.GVA_DB.Find(&packages).Error; err != nil {
		return nil, fmt.Errorf("获取包信息失败: %v", err)
	}

	// 2. 获取历史记录
	var histories []model.SysAutoCodeHistory
	if err := global.GVA_DB.Find(&histories).Error; err != nil {
		return nil, fmt.Errorf("获取历史记录失败: %v", err)
	}

	// 3. 检查空包并进行清理
	cleanupInfo := &CleanupInfo{
		DeletedPackages: []string{},
		DeletedModules:  []string{},
	}

	var validPackages []model.SysAutoCodePackage
	var emptyPackageHistoryIDs []uint

	for _, pkg := range packages {
		isEmpty, err := g.isPackageFolderEmpty(pkg.PackageName, pkg.Template)
		if err != nil {
			global.GVA_LOG.Warn(fmt.Sprintf("检查包 %s 是否为空时出错: %v", pkg.PackageName, err))
			continue
		}

		if isEmpty {
			// 删除空包文件夹
			if err := g.removeEmptyPackageFolder(pkg.PackageName, pkg.Template); err != nil {
				global.GVA_LOG.Warn(fmt.Sprintf("删除空包文件夹 %s 失败: %v", pkg.PackageName, err))
			} else {
				cleanupInfo.DeletedPackages = append(cleanupInfo.DeletedPackages, pkg.PackageName)
			}

			// 删除数据库记录
			if err := global.GVA_DB.Delete(&pkg).Error; err != nil {
				global.GVA_LOG.Warn(fmt.Sprintf("删除包数据库记录 %s 失败: %v", pkg.PackageName, err))
			}

			// 收集相关的历史记录ID
			for _, history := range histories {
				if history.Package == pkg.PackageName {
					emptyPackageHistoryIDs = append(emptyPackageHistoryIDs, history.ID)
					cleanupInfo.DeletedModules = append(cleanupInfo.DeletedModules, history.StructName)
				}
			}
		} else {
			validPackages = append(validPackages, pkg)
		}
	}

	// 5. 清理空包相关的历史记录和脏历史记录
	var dirtyHistoryIDs []uint
	for _, history := range histories {
		// 检查是否为空包相关的历史记录
		for _, emptyID := range emptyPackageHistoryIDs {
			if history.ID == emptyID {
				dirtyHistoryIDs = append(dirtyHistoryIDs, history.ID)
				break
			}
		}
	}

	// 删除脏历史记录
	if len(dirtyHistoryIDs) > 0 {
		if err := global.GVA_DB.Delete(&model.SysAutoCodeHistory{}, "id IN ?", dirtyHistoryIDs).Error; err != nil {
			global.GVA_LOG.Warn(fmt.Sprintf("删除脏历史记录失败: %v", err))
		} else {
			global.GVA_LOG.Info(fmt.Sprintf("成功删除 %d 条脏历史记录", len(dirtyHistoryIDs)))
		}

		// 清理相关的API和菜单记录
		if err := g.cleanupRelatedApiAndMenus(dirtyHistoryIDs); err != nil {
			global.GVA_LOG.Warn(fmt.Sprintf("清理相关API和菜单记录失败: %v", err))
		}
	}

	// 6. 扫描预设计模块
	predesignedModules, err := g.scanPredesignedModules()
	if err != nil {
		global.GVA_LOG.Warn(fmt.Sprintf("扫描预设计模块失败: %v", err))
		predesignedModules = []PredesignedModuleInfo{} // 设置为空列表，不影响主流程
	}

	// 7. 过滤掉与已删除包相关的模块
	filteredModules := []PredesignedModuleInfo{}
	for _, module := range predesignedModules {
		isDeleted := false
		for _, deletedPkg := range cleanupInfo.DeletedPackages {
			if module.PackageName == deletedPkg {
				isDeleted = true
				break
			}
		}
		if !isDeleted {
			filteredModules = append(filteredModules, module)
		}
	}

	// 8. 构建分析结果消息
	var analysisMessage strings.Builder
	if len(cleanupInfo.DeletedPackages) > 0 || len(cleanupInfo.DeletedModules) > 0 {
		analysisMessage.WriteString("**系统清理完成**\n\n")
		if len(cleanupInfo.DeletedPackages) > 0 {
			analysisMessage.WriteString(fmt.Sprintf("- 删除了 %d 个空包: %s\n", len(cleanupInfo.DeletedPackages), strings.Join(cleanupInfo.DeletedPackages, ", ")))
		}
		if len(cleanupInfo.DeletedModules) > 0 {
			analysisMessage.WriteString(fmt.Sprintf("- 删除了 %d 个相关模块: %s\n", len(cleanupInfo.DeletedModules), strings.Join(cleanupInfo.DeletedModules, ", ")))
		}
		analysisMessage.WriteString("\n")
		cleanupInfo.CleanupMessage = analysisMessage.String()
	}

	analysisMessage.WriteString(" **分析结果**\n\n")
	analysisMessage.WriteString(fmt.Sprintf("- **现有包数量**: %d\n", len(validPackages)))
	analysisMessage.WriteString(fmt.Sprintf("- **预设计模块数量**: %d\n\n", len(filteredModules)))

	// 9. 转换包信息
	existingPackages := make([]PackageInfo, len(validPackages))
	for i, pkg := range validPackages {
		existingPackages[i] = PackageInfo{
			PackageName: pkg.PackageName,
			Template:    pkg.Template,
			Label:       pkg.Label,
			Desc:        pkg.Desc,
			Module:      pkg.Module,
			IsEmpty:     false, // 已经过滤掉空包
		}
	}

	dictionaries := []DictionaryPre{} // 这里可以根据需要填充字典信息
	err = global.GVA_DB.Table("sys_dictionaries").Find(&dictionaries, "deleted_at is null").Error
	if err != nil {
		global.GVA_LOG.Warn(fmt.Sprintf("获取字典信息失败: %v", err))
		dictionaries = []DictionaryPre{} // 设置为空列表，不影响主流程
	}

	// 10. 构建响应
	response := &AnalyzeResponse{
		ExistingPackages:   existingPackages,
		PredesignedModules: filteredModules,
		Dictionaries:       dictionaries,
	}

	return response, nil
}

// isPackageFolderEmpty 检查包文件夹是否为空
func (g *GVAAnalyzer) isPackageFolderEmpty(packageName, template string) (bool, error) {
	// 根据模板类型确定基础路径
	var basePath string
	if template == "plugin" {
		basePath = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", packageName)
	} else {
		basePath = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "api", "v1", packageName)
	}

	// 检查文件夹是否存在
	if _, err := os.Stat(basePath); os.IsNotExist(err) {
		return true, nil // 文件夹不存在，视为空
	} else if err != nil {
		return false, err // 其他错误
	}
	// 递归检查是否有.go文件
	return g.hasGoFilesRecursive(basePath)
}

// hasGoFilesRecursive 递归检查目录及其子目录中是否有.go文件
func (g *GVAAnalyzer) hasGoFilesRecursive(dirPath string) (bool, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return true, err // 读取失败，返回空
	}

	// 检查当前目录下的.go文件
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".go") {
			return false, nil // 找到.go文件，不为空
		}
	}

	// 递归检查子目录
	for _, entry := range entries {
		if entry.IsDir() {
			subDirPath := filepath.Join(dirPath, entry.Name())
			isEmpty, err := g.hasGoFilesRecursive(subDirPath)
			if err != nil {
				continue // 忽略子目录的错误，继续检查其他目录
			}
			if !isEmpty {
				return false, nil // 子目录中找到.go文件，不为空
			}
		}
	}

	return true, nil // 没有找到.go文件，为空
}

// removeEmptyPackageFolder 删除空包文件夹
func (g *GVAAnalyzer) removeEmptyPackageFolder(packageName, template string) error {
	var basePath string
	if template == "plugin" {
		basePath = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", packageName)
	} else {
		// 对于package类型，需要删除多个目录
		paths := []string{
			filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "api", "v1", packageName),
			filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "model", packageName),
			filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "router", packageName),
			filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "service", packageName),
		}
		for _, path := range paths {
			if err := g.removeDirectoryIfExists(path); err != nil {
				return err
			}
		}
		return nil
	}

	return g.removeDirectoryIfExists(basePath)
}

// removeDirectoryIfExists 删除目录（如果存在）
func (g *GVAAnalyzer) removeDirectoryIfExists(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return nil // 目录不存在，无需删除
	} else if err != nil {
		return err // 其他错误
	}

	return os.RemoveAll(dirPath)
}

// cleanupRelatedApiAndMenus 清理相关的API和菜单记录
func (g *GVAAnalyzer) cleanupRelatedApiAndMenus(historyIDs []uint) error {
	if len(historyIDs) == 0 {
		return nil
	}

	// 这里可以根据需要实现具体的API和菜单清理逻辑
	// 由于涉及到具体的业务逻辑，这里只做日志记录
	global.GVA_LOG.Info(fmt.Sprintf("清理历史记录ID %v 相关的API和菜单记录", historyIDs))

	// 可以调用service层的相关方法进行清理
	// 例如：service.ServiceGroupApp.SystemApiService.DeleteApisByIds(historyIDs)
	// 例如：service.ServiceGroupApp.MenuService.DeleteMenusByIds(historyIDs)

	return nil
}

// scanPredesignedModules 扫描预设计模块
func (g *GVAAnalyzer) scanPredesignedModules() ([]PredesignedModuleInfo, error) {
	// 获取autocode配置路径
	autocodeRoot := global.GVA_CONFIG.AutoCode.Root
	if autocodeRoot == "" {
		return nil, errors.New("autocode根路径未配置")
	}

	var modules []PredesignedModuleInfo

	// 扫描plugin目录
	pluginModules, err := g.scanPluginModules(filepath.Join(autocodeRoot, global.GVA_CONFIG.AutoCode.Server, "plugin"))
	if err != nil {
		global.GVA_LOG.Warn(fmt.Sprintf("扫描plugin模块失败: %v", err))
	} else {
		modules = append(modules, pluginModules...)
	}

	// 扫描model目录
	modelModules, err := g.scanModelModules(filepath.Join(autocodeRoot, global.GVA_CONFIG.AutoCode.Server, "model"))
	if err != nil {
		global.GVA_LOG.Warn(fmt.Sprintf("扫描model模块失败: %v", err))
	} else {
		modules = append(modules, modelModules...)
	}

	return modules, nil
}

// scanPluginModules 扫描插件模块
func (g *GVAAnalyzer) scanPluginModules(pluginDir string) ([]PredesignedModuleInfo, error) {
	var modules []PredesignedModuleInfo

	if _, err := os.Stat(pluginDir); os.IsNotExist(err) {
		return modules, nil // 目录不存在，返回空列表
	}

	entries, err := os.ReadDir(pluginDir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			pluginName := entry.Name()
			pluginPath := filepath.Join(pluginDir, pluginName)

			// 查找model目录
			modelDir := filepath.Join(pluginPath, "model")
			if _, err := os.Stat(modelDir); err == nil {
				// 扫描model目录下的模块
				pluginModules, err := g.scanModulesInDirectory(modelDir, pluginName, "plugin")
				if err != nil {
					global.GVA_LOG.Warn(fmt.Sprintf("扫描插件 %s 的模块失败: %v", pluginName, err))
					continue
				}
				modules = append(modules, pluginModules...)
			}
		}
	}

	return modules, nil
}

// scanModelModules 扫描模型模块
func (g *GVAAnalyzer) scanModelModules(modelDir string) ([]PredesignedModuleInfo, error) {
	var modules []PredesignedModuleInfo

	if _, err := os.Stat(modelDir); os.IsNotExist(err) {
		return modules, nil // 目录不存在，返回空列表
	}

	entries, err := os.ReadDir(modelDir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			packageName := entry.Name()
			packagePath := filepath.Join(modelDir, packageName)

			// 扫描包目录下的模块
			packageModules, err := g.scanModulesInDirectory(packagePath, packageName, "package")
			if err != nil {
				global.GVA_LOG.Warn(fmt.Sprintf("扫描包 %s 的模块失败: %v", packageName, err))
				continue
			}
			modules = append(modules, packageModules...)
		}
	}

	return modules, nil
}

// scanModulesInDirectory 扫描目录中的模块
func (g *GVAAnalyzer) scanModulesInDirectory(dir, packageName, template string) ([]PredesignedModuleInfo, error) {
	var modules []PredesignedModuleInfo

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".go") {
			moduleName := strings.TrimSuffix(entry.Name(), ".go")
			filePath := filepath.Join(dir, entry.Name())

			module := PredesignedModuleInfo{
				ModuleName:  moduleName,
				PackageName: packageName,
				Template:    template,
				FilePaths:   []string{filePath},
				Description: fmt.Sprintf("%s模块中的%s", packageName, moduleName),
			}
			modules = append(modules, module)
		}
	}

	return modules, nil
}
