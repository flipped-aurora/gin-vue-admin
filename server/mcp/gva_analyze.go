package mcpTool

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
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
	Requirement string `json:"requirement"` // 用户需求描述（可选，仅作为调用方上下文，不影响返回快照）
}

// AnalyzeResponse 分析响应结构体
type AnalyzeResponse struct {
	ExistingPackages   []PackageInfo           `json:"existingPackages"`   // 现有包信息
	PredesignedModules []PredesignedModuleInfo `json:"predesignedModules"` // 预设计模块信息
	Dictionaries       []DictionaryPre         `json:"dictionaries"`       // 字典信息
	CleanupInfo        *CleanupInfo            `json:"cleanupInfo"`        // 清理信息（如果有）
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
		mcp.WithDescription("返回当前系统中有效的包、模块与字典清单（快照），供调用方对照 requirement 自行判断需要新建哪些包/模块/字典；同时检查并清理空包，保持系统整洁。注意：本工具只提供现状快照，不对 requirement 内容做过滤或匹配，返回结果与 requirement 文本无关。"),
		mcp.WithString("requirement",
			mcp.Description("用户需求描述，可选，仅作为调用方分析的上下文；本工具返回全量快照，不依赖也不解析其内容"),
		),
	)
}

// Handle 处理分析请求
func (g *GVAAnalyzer) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 解析请求参数：requirement 为可选上下文，工具返回全量快照，不依赖其内容
	requirementStr, _ := request.GetArguments()["requirement"].(string)

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
	return textResultWithJSON("", response)
}

// performAnalysis 执行分析逻辑
func (g *GVAAnalyzer) performAnalysis(ctx context.Context, req AnalyzeRequest) (*AnalyzeResponse, error) {
	_ = req

	packages, err := fetchAutoCodePackages(ctx)
	if err != nil {
		return nil, fmt.Errorf("获取包信息失败: %v", err)
	}

	histories, err := fetchAutoCodeHistories(ctx)
	if err != nil {
		return nil, fmt.Errorf("获取历史记录失败: %v", err)
	}

	cleanupInfo := &CleanupInfo{
		DeletedPackages: []string{},
		DeletedModules:  []string{},
	}

	validPackages := make([]PackageInfo, 0, len(packages))
	var emptyPackageHistoryIDs []uint

	for _, pkg := range packages {
		isEmpty, err := g.isPackageFolderEmpty(pkg.PackageName, pkg.Template)
		if err != nil {
			logger.WithCtx(ctx).Mod("mcp").Err(err).Warn(fmt.Sprintf("检查包 %s 是否为空时出错", pkg.PackageName))
			continue
		}

		if isEmpty {
			if err := g.removeEmptyPackageFolder(pkg.PackageName, pkg.Template); err != nil {
				logger.WithCtx(ctx).Mod("mcp").Err(err).Warn(fmt.Sprintf("删除空包文件夹 %s 失败", pkg.PackageName))
			} else {
				cleanupInfo.DeletedPackages = append(cleanupInfo.DeletedPackages, pkg.PackageName)
			}

			if err := deleteAutoCodePackage(ctx, pkg.ID); err != nil {
				logger.WithCtx(ctx).Mod("mcp").Err(err).Warn(fmt.Sprintf("删除包数据库记录 %s 失败", pkg.PackageName))
			}

			for _, history := range histories {
				if history.Package == pkg.PackageName {
					emptyPackageHistoryIDs = append(emptyPackageHistoryIDs, history.ID)
					cleanupInfo.DeletedModules = append(cleanupInfo.DeletedModules, history.StructName)
				}
			}
			continue
		}

		validPackages = append(validPackages, PackageInfo{
			PackageName: pkg.PackageName,
			Template:    pkg.Template,
			Label:       pkg.Label,
			Desc:        pkg.Desc,
			Module:      pkg.Module,
			IsEmpty:     false,
		})
	}

	emptyHistoryIDSet := make(map[uint]struct{}, len(emptyPackageHistoryIDs))
	for _, emptyID := range emptyPackageHistoryIDs {
		emptyHistoryIDSet[emptyID] = struct{}{}
	}

	var dirtyHistoryIDs []uint
	for _, history := range histories {
		if _, ok := emptyHistoryIDSet[history.ID]; ok {
			dirtyHistoryIDs = append(dirtyHistoryIDs, history.ID)
		}
	}

	if len(dirtyHistoryIDs) > 0 {
		deletedCount := 0
		for _, historyID := range dirtyHistoryIDs {
			if err := deleteAutoCodeHistory(ctx, historyID); err != nil {
				logger.WithCtx(ctx).Mod("mcp").Err(err).Warn("删除脏历史记录失败")
				continue
			}
			deletedCount++
		}
		if deletedCount > 0 {
			logger.WithCtx(ctx).Mod("mcp").Info(fmt.Sprintf("成功删除 %d 条脏历史记录", deletedCount))
		}

		if err := g.cleanupRelatedApiAndMenus(ctx, dirtyHistoryIDs); err != nil {
			logger.WithCtx(ctx).Mod("mcp").Err(err).Warn("清理相关API和菜单记录失败")
		}
	}

	predesignedModules, err := g.scanPredesignedModules(ctx)
	if err != nil {
		logger.WithCtx(ctx).Mod("mcp").Err(err).Warn("扫描预设计模块失败")
		predesignedModules = []PredesignedModuleInfo{}
	}

	filteredModules := []PredesignedModuleInfo{}
	for _, module := range predesignedModules {
		if !slices.Contains(cleanupInfo.DeletedPackages, module.PackageName) {
			filteredModules = append(filteredModules, module)
		}
	}

	dictionaries := []DictionaryPre{}
	dictEntities, err := fetchDictionaryList(ctx, "")
	if err != nil {
		logger.WithCtx(ctx).Mod("mcp").Err(err).Warn("获取字典信息失败")
	} else {
		for _, dictionary := range dictEntities {
			dictionaries = append(dictionaries, DictionaryPre{
				Type: dictionary.Type,
				Desc: dictionary.Desc,
			})
		}
	}

	var cleanupResult *CleanupInfo
	if len(cleanupInfo.DeletedPackages) > 0 || len(cleanupInfo.DeletedModules) > 0 {
		var message strings.Builder
		message.WriteString("**系统清理完成**\n\n")
		if len(cleanupInfo.DeletedPackages) > 0 {
			message.WriteString(fmt.Sprintf("- 删除了 %d 个空包: %s\n", len(cleanupInfo.DeletedPackages), strings.Join(cleanupInfo.DeletedPackages, ", ")))
		}
		if len(cleanupInfo.DeletedModules) > 0 {
			message.WriteString(fmt.Sprintf("- 删除了 %d 个相关模块: %s\n", len(cleanupInfo.DeletedModules), strings.Join(cleanupInfo.DeletedModules, ", ")))
		}
		cleanupInfo.CleanupMessage = message.String()
		cleanupResult = cleanupInfo
	}

	response := &AnalyzeResponse{
		ExistingPackages:   validPackages,
		PredesignedModules: filteredModules,
		Dictionaries:       dictionaries,
		CleanupInfo:        cleanupResult,
	}

	return response, nil
}

// isPackageFolderEmpty 检查包文件夹是否为空
func (g *GVAAnalyzer) isPackageFolderEmpty(packageName, template string) (bool, error) {
	root := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server)

	// package 模板的代码散落在 api/v1、model、router、service 四处(与 removeEmptyPackageFolder 删除的目录一致);
	// 必须四处全无 .go 才算空,否则仅凭 api/v1 空就删掉 DB 包记录会遗留 service/model 下的孤儿代码
	var dirs []string
	if template == "plugin" {
		dirs = []string{filepath.Join(root, "plugin", packageName)}
	} else {
		dirs = []string{
			filepath.Join(root, "api", "v1", packageName),
			filepath.Join(root, "model", packageName),
			filepath.Join(root, "router", packageName),
			filepath.Join(root, "service", packageName),
		}
	}

	for _, dir := range dirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			continue // 该目录不存在,视作此处为空,继续看其它目录
		} else if err != nil {
			return false, err
		}
		noGoFiles, err := g.dirHasNoGoFiles(dir)
		if err != nil {
			return false, err
		}
		if !noGoFiles {
			return false, nil // 任一目录仍有 .go,整包不算空
		}
	}
	return true, nil
}

// dirHasNoGoFiles 递归检查目录及其子目录：无 .go 文件（含读取失败）时返回 true，发现 .go 文件时返回 false
func (g *GVAAnalyzer) dirHasNoGoFiles(dirPath string) (bool, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return true, err // 读取失败，视作无 .go 文件
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
			noGoFiles, err := g.dirHasNoGoFiles(subDirPath)
			if err != nil {
				continue // 忽略子目录的错误，继续检查其他目录
			}
			if !noGoFiles {
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

	// 检查目录中是否包含go文件
	noGoFiles, err := g.dirHasNoGoFiles(dirPath)
	if err != nil {
		return err
	}
	if noGoFiles {
		return os.RemoveAll(dirPath)
	}
	return nil
}

// cleanupRelatedApiAndMenus 清理相关的API和菜单记录
func (g *GVAAnalyzer) cleanupRelatedApiAndMenus(ctx context.Context, historyIDs []uint) error {
	if len(historyIDs) == 0 {
		return nil
	}

	// 这里可以根据需要实现具体的API和菜单清理逻辑
	// 由于涉及到具体的业务逻辑，这里只做日志记录
	logger.WithCtx(ctx).Mod("mcp").Info(fmt.Sprintf("清理历史记录ID %v 相关的API和菜单记录", historyIDs))

	// 可以调用service层的相关方法进行清理
	// 例如：service.ServiceGroupApp.SystemApiService.DeleteApisByIds(historyIDs)
	// 例如：service.ServiceGroupApp.MenuService.DeleteMenusByIds(historyIDs)

	return nil
}

// scanPredesignedModules 扫描预设计模块
func (g *GVAAnalyzer) scanPredesignedModules(ctx context.Context) ([]PredesignedModuleInfo, error) {
	// 获取autocode配置路径
	autocodeRoot := global.GVA_CONFIG.AutoCode.Root
	if autocodeRoot == "" {
		return nil, errors.New("autocode根路径未配置")
	}

	var modules []PredesignedModuleInfo

	// 扫描plugin目录
	pluginModules, err := g.scanPluginModules(ctx, filepath.Join(autocodeRoot, global.GVA_CONFIG.AutoCode.Server, "plugin"))
	if err != nil {
		logger.WithCtx(ctx).Mod("mcp").Err(err).Warn("扫描plugin模块失败")
	} else {
		modules = append(modules, pluginModules...)
	}

	// 扫描model目录
	modelModules, err := g.scanModelModules(ctx, filepath.Join(autocodeRoot, global.GVA_CONFIG.AutoCode.Server, "model"))
	if err != nil {
		logger.WithCtx(ctx).Mod("mcp").Err(err).Warn("扫描model模块失败")
	} else {
		modules = append(modules, modelModules...)
	}

	return modules, nil
}

// scanPluginModules 扫描插件模块
func (g *GVAAnalyzer) scanPluginModules(ctx context.Context, pluginDir string) ([]PredesignedModuleInfo, error) {
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
					logger.WithCtx(ctx).Mod("mcp").Err(err).Warn(fmt.Sprintf("扫描插件 %s 的模块失败", pluginName))
					continue
				}
				modules = append(modules, pluginModules...)
			}
		}
	}

	return modules, nil
}

// scanModelModules 扫描模型模块
func (g *GVAAnalyzer) scanModelModules(ctx context.Context, modelDir string) ([]PredesignedModuleInfo, error) {
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
				logger.WithCtx(ctx).Mod("mcp").Err(err).Warn(fmt.Sprintf("扫描包 %s 的模块失败", packageName))
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
			filePath := outputPath(filepath.Join(dir, entry.Name()))

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
