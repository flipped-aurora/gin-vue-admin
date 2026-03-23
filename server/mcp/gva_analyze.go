package mcpTool

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/mark3labs/mcp-go/mcp"
)

// 娉ㄥ唽宸ュ叿
func init() {
	RegisterTool(&GVAAnalyzer{})
}

// GVAAnalyzer GVA鍒嗘瀽鍣?- 鐢ㄤ簬鍒嗘瀽褰撳墠鍔熻兘鏄惁闇€瑕佸垱寤虹嫭绔嬬殑package鍜宮odule
type GVAAnalyzer struct{}

// AnalyzeRequest 鍒嗘瀽璇锋眰缁撴瀯浣?
type AnalyzeRequest struct {
	Requirement string `json:"requirement" binding:"required"` // 鐢ㄦ埛闇€姹傛弿杩?
}

// AnalyzeResponse 鍒嗘瀽鍝嶅簲缁撴瀯浣?
type AnalyzeResponse struct {
	ExistingPackages   []PackageInfo           `json:"existingPackages"`   // 鐜版湁鍖呬俊鎭?
	PredesignedModules []PredesignedModuleInfo `json:"predesignedModules"` // 棰勮璁℃ā鍧椾俊鎭?
	Dictionaries       []DictionaryPre         `json:"dictionaries"`       // 瀛楀吀淇℃伅
	CleanupInfo        *CleanupInfo            `json:"cleanupInfo"`        // 娓呯悊淇℃伅锛堝鏋滄湁锛?
}

// ModuleInfo 妯″潡淇℃伅
type ModuleInfo struct {
	ModuleName  string   `json:"moduleName"`  // 妯″潡鍚嶇О
	PackageName string   `json:"packageName"` // 鍖呭悕
	Template    string   `json:"template"`    // 妯℃澘绫诲瀷
	StructName  string   `json:"structName"`  // 缁撴瀯浣撳悕绉?
	TableName   string   `json:"tableName"`   // 琛ㄥ悕
	Description string   `json:"description"` // 鎻忚堪
	FilePaths   []string `json:"filePaths"`   // 鐩稿叧鏂囦欢璺緞
}

// PackageInfo 鍖呬俊鎭?
type PackageInfo struct {
	PackageName string `json:"packageName"` // 鍖呭悕
	Template    string `json:"template"`    // 妯℃澘绫诲瀷
	Label       string `json:"label"`       // 鏍囩
	Desc        string `json:"desc"`        // 鎻忚堪
	Module      string `json:"module"`      // 妯″潡
	IsEmpty     bool   `json:"isEmpty"`     // 鏄惁涓虹┖鍖?
}

// PredesignedModuleInfo 棰勮璁℃ā鍧椾俊鎭?
type PredesignedModuleInfo struct {
	ModuleName  string   `json:"moduleName"`  // 妯″潡鍚嶇О
	PackageName string   `json:"packageName"` // 鍖呭悕
	Template    string   `json:"template"`    // 妯℃澘绫诲瀷
	FilePaths   []string `json:"filePaths"`   // 鏂囦欢璺緞鍒楄〃
	Description string   `json:"description"` // 鎻忚堪
}

// CleanupInfo 娓呯悊淇℃伅
type CleanupInfo struct {
	DeletedPackages []string `json:"deletedPackages"` // 宸插垹闄ょ殑鍖?
	DeletedModules  []string `json:"deletedModules"`  // 宸插垹闄ょ殑妯″潡
	CleanupMessage  string   `json:"cleanupMessage"`  // 娓呯悊娑堟伅
}

// New 鍒涘缓GVA鍒嗘瀽鍣ㄥ伐鍏?
func (g *GVAAnalyzer) New() mcp.Tool {
	return mcp.NewTool("gva_analyze",
		mcp.WithDescription("返回当前系统中有效的包和模块信息，并分析用户需求是否需要创建新的包、模块和字典。同时检查并清理空包，确保系统整洁。"),
		mcp.WithString("requirement",
			mcp.Description("用户需求描述，用于分析是否需要创建新的包和模块"),
			mcp.Required(),
		),
	)
}

// Handle 澶勭悊鍒嗘瀽璇锋眰
func (g *GVAAnalyzer) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 瑙ｆ瀽璇锋眰鍙傛暟
	requirementStr, ok := request.GetArguments()["requirement"].(string)
	if !ok || requirementStr == "" {
		return nil, errors.New("鍙傛暟閿欒锛歳equirement 蹇呴』鏄潪绌哄瓧绗︿覆")
	}

	// 鍒涘缓鍒嗘瀽璇锋眰
	analyzeReq := AnalyzeRequest{
		Requirement: requirementStr,
	}

	// 鎵ц鍒嗘瀽閫昏緫
	response, err := g.performAnalysis(ctx, analyzeReq)
	if err != nil {
		return nil, fmt.Errorf("鍒嗘瀽澶辫触: %v", err)
	}

	// 搴忓垪鍖栧搷搴?
	responseJSON, err := json.Marshal(response)
	if err != nil {
		return nil, fmt.Errorf("搴忓垪鍖栧搷搴斿け璐? %v", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.NewTextContent(string(responseJSON)),
		},
	}, nil
}

// performAnalysis 鎵ц鍒嗘瀽閫昏緫
func (g *GVAAnalyzer) performAnalysis(ctx context.Context, req AnalyzeRequest) (*AnalyzeResponse, error) {
	_ = req

	packages, err := fetchAutoCodePackages(ctx)
	if err != nil {
		return nil, fmt.Errorf("鑾峰彇鍖呬俊鎭け璐? %v", err)
	}

	histories, err := fetchAutoCodeHistories(ctx)
	if err != nil {
		return nil, fmt.Errorf("鑾峰彇鍘嗗彶璁板綍澶辫触: %v", err)
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
			global.GVA_LOG.Warn(fmt.Sprintf("妫€鏌ュ寘 %s 鏄惁涓虹┖鏃跺嚭閿? %v", pkg.PackageName, err))
			continue
		}

		if isEmpty {
			if err := g.removeEmptyPackageFolder(pkg.PackageName, pkg.Template); err != nil {
				global.GVA_LOG.Warn(fmt.Sprintf("鍒犻櫎绌哄寘鏂囦欢澶?%s 澶辫触: %v", pkg.PackageName, err))
			} else {
				cleanupInfo.DeletedPackages = append(cleanupInfo.DeletedPackages, pkg.PackageName)
			}

			if err := deleteAutoCodePackage(ctx, pkg.ID); err != nil {
				global.GVA_LOG.Warn(fmt.Sprintf("鍒犻櫎鍖呮暟鎹簱璁板綍 %s 澶辫触: %v", pkg.PackageName, err))
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

	var dirtyHistoryIDs []uint
	for _, history := range histories {
		for _, emptyID := range emptyPackageHistoryIDs {
			if history.ID == emptyID {
				dirtyHistoryIDs = append(dirtyHistoryIDs, history.ID)
				break
			}
		}
	}

	if len(dirtyHistoryIDs) > 0 {
		deletedCount := 0
		for _, historyID := range dirtyHistoryIDs {
			if err := deleteAutoCodeHistory(ctx, historyID); err != nil {
				global.GVA_LOG.Warn(fmt.Sprintf("鍒犻櫎鑴忓巻鍙茶褰曞け璐? %v", err))
				continue
			}
			deletedCount++
		}
		if deletedCount > 0 {
			global.GVA_LOG.Info(fmt.Sprintf("鎴愬姛鍒犻櫎 %d 鏉¤剰鍘嗗彶璁板綍", deletedCount))
		}

		if err := g.cleanupRelatedApiAndMenus(dirtyHistoryIDs); err != nil {
			global.GVA_LOG.Warn(fmt.Sprintf("娓呯悊鐩稿叧API鍜岃彍鍗曡褰曞け璐? %v", err))
		}
	}

	predesignedModules, err := g.scanPredesignedModules()
	if err != nil {
		global.GVA_LOG.Warn(fmt.Sprintf("鎵弿棰勮璁℃ā鍧楀け璐? %v", err))
		predesignedModules = []PredesignedModuleInfo{}
	}

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

	dictionaries := []DictionaryPre{}
	dictEntities, err := fetchDictionaryList(ctx, "")
	if err != nil {
		global.GVA_LOG.Warn(fmt.Sprintf("鑾峰彇瀛楀吀淇℃伅澶辫触: %v", err))
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
		message.WriteString("**绯荤粺娓呯悊瀹屾垚**\n\n")
		if len(cleanupInfo.DeletedPackages) > 0 {
			message.WriteString(fmt.Sprintf("- 鍒犻櫎浜?%d 涓┖鍖? %s\n", len(cleanupInfo.DeletedPackages), strings.Join(cleanupInfo.DeletedPackages, ", ")))
		}
		if len(cleanupInfo.DeletedModules) > 0 {
			message.WriteString(fmt.Sprintf("- 鍒犻櫎浜?%d 涓浉鍏虫ā鍧? %s\n", len(cleanupInfo.DeletedModules), strings.Join(cleanupInfo.DeletedModules, ", ")))
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

// isPackageFolderEmpty 妫€鏌ュ寘鏂囦欢澶规槸鍚︿负绌?
func (g *GVAAnalyzer) isPackageFolderEmpty(packageName, template string) (bool, error) {
	// 鏍规嵁妯℃澘绫诲瀷纭畾鍩虹璺緞
	var basePath string
	if template == "plugin" {
		basePath = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", packageName)
	} else {
		basePath = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "api", "v1", packageName)
	}

	// 妫€鏌ユ枃浠跺す鏄惁瀛樺湪
	if _, err := os.Stat(basePath); os.IsNotExist(err) {
		return true, nil // 鏂囦欢澶逛笉瀛樺湪锛岃涓虹┖
	} else if err != nil {
		return false, err // 鍏朵粬閿欒
	}
	// 閫掑綊妫€鏌ユ槸鍚︽湁.go鏂囦欢
	return g.hasGoFilesRecursive(basePath)
}

// hasGoFilesRecursive 閫掑綊妫€鏌ョ洰褰曞強鍏跺瓙鐩綍涓槸鍚︽湁.go鏂囦欢
func (g *GVAAnalyzer) hasGoFilesRecursive(dirPath string) (bool, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return true, err // 璇诲彇澶辫触锛岃繑鍥炵┖
	}

	// 妫€鏌ュ綋鍓嶇洰褰曚笅鐨?go鏂囦欢
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".go") {
			return false, nil // 鎵惧埌.go鏂囦欢锛屼笉涓虹┖
		}
	}

	// 閫掑綊妫€鏌ュ瓙鐩綍
	for _, entry := range entries {
		if entry.IsDir() {
			subDirPath := filepath.Join(dirPath, entry.Name())
			isEmpty, err := g.hasGoFilesRecursive(subDirPath)
			if err != nil {
				continue // 蹇界暐瀛愮洰褰曠殑閿欒锛岀户缁鏌ュ叾浠栫洰褰?
			}
			if !isEmpty {
				return false, nil // 瀛愮洰褰曚腑鎵惧埌.go鏂囦欢锛屼笉涓虹┖
			}
		}
	}

	return true, nil // 娌℃湁鎵惧埌.go鏂囦欢锛屼负绌?
}

// removeEmptyPackageFolder 鍒犻櫎绌哄寘鏂囦欢澶?
func (g *GVAAnalyzer) removeEmptyPackageFolder(packageName, template string) error {
	var basePath string
	if template == "plugin" {
		basePath = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", packageName)
	} else {
		// 瀵逛簬package绫诲瀷锛岄渶瑕佸垹闄ゅ涓洰褰?
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

// removeDirectoryIfExists 鍒犻櫎鐩綍锛堝鏋滃瓨鍦級
func (g *GVAAnalyzer) removeDirectoryIfExists(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return nil // 鐩綍涓嶅瓨鍦紝鏃犻渶鍒犻櫎
	} else if err != nil {
		return err // 鍏朵粬閿欒
	}

	// 妫€鏌ョ洰褰曚腑鏄惁鍖呭惈go鏂囦欢
	noGoFiles, err := g.hasGoFilesRecursive(dirPath)
	if err != nil {
		return err
	}
	// hasGoFilesRecursive 杩斿洖 false 琛ㄧず鍙戠幇浜?go 鏂囦欢
	if noGoFiles {
		return os.RemoveAll(dirPath)
	}
	return nil
}

// cleanupRelatedApiAndMenus 娓呯悊鐩稿叧鐨凙PI鍜岃彍鍗曡褰?
func (g *GVAAnalyzer) cleanupRelatedApiAndMenus(historyIDs []uint) error {
	if len(historyIDs) == 0 {
		return nil
	}

	// 杩欓噷鍙互鏍规嵁闇€瑕佸疄鐜板叿浣撶殑API鍜岃彍鍗曟竻鐞嗛€昏緫
	// 鐢变簬娑夊強鍒板叿浣撶殑涓氬姟閫昏緫锛岃繖閲屽彧鍋氭棩蹇楄褰?
	global.GVA_LOG.Info(fmt.Sprintf("清理历史记录ID %v 相关的API和菜单记录", historyIDs))

	// 鍙互璋冪敤service灞傜殑鐩稿叧鏂规硶杩涜娓呯悊
	// 渚嬪锛歴ervice.ServiceGroupApp.SystemApiService.DeleteApisByIds(historyIDs)
	// 渚嬪锛歴ervice.ServiceGroupApp.MenuService.DeleteMenusByIds(historyIDs)

	return nil
}

// scanPredesignedModules 鎵弿棰勮璁℃ā鍧?
func (g *GVAAnalyzer) scanPredesignedModules() ([]PredesignedModuleInfo, error) {
	// 鑾峰彇autocode閰嶇疆璺緞
	autocodeRoot := global.GVA_CONFIG.AutoCode.Root
	if autocodeRoot == "" {
		return nil, errors.New("autocode鏍硅矾寰勬湭閰嶇疆")
	}

	var modules []PredesignedModuleInfo

	// 鎵弿plugin鐩綍
	pluginModules, err := g.scanPluginModules(filepath.Join(autocodeRoot, global.GVA_CONFIG.AutoCode.Server, "plugin"))
	if err != nil {
		global.GVA_LOG.Warn(fmt.Sprintf("鎵弿plugin妯″潡澶辫触: %v", err))
	} else {
		modules = append(modules, pluginModules...)
	}

	// 鎵弿model鐩綍
	modelModules, err := g.scanModelModules(filepath.Join(autocodeRoot, global.GVA_CONFIG.AutoCode.Server, "model"))
	if err != nil {
		global.GVA_LOG.Warn(fmt.Sprintf("鎵弿model妯″潡澶辫触: %v", err))
	} else {
		modules = append(modules, modelModules...)
	}

	return modules, nil
}

// scanPluginModules 鎵弿鎻掍欢妯″潡
func (g *GVAAnalyzer) scanPluginModules(pluginDir string) ([]PredesignedModuleInfo, error) {
	var modules []PredesignedModuleInfo

	if _, err := os.Stat(pluginDir); os.IsNotExist(err) {
		return modules, nil // 鐩綍涓嶅瓨鍦紝杩斿洖绌哄垪琛?
	}

	entries, err := os.ReadDir(pluginDir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			pluginName := entry.Name()
			pluginPath := filepath.Join(pluginDir, pluginName)

			// 鏌ユ壘model鐩綍
			modelDir := filepath.Join(pluginPath, "model")
			if _, err := os.Stat(modelDir); err == nil {
				// 鎵弿model鐩綍涓嬬殑妯″潡
				pluginModules, err := g.scanModulesInDirectory(modelDir, pluginName, "plugin")
				if err != nil {
					global.GVA_LOG.Warn(fmt.Sprintf("鎵弿鎻掍欢 %s 鐨勬ā鍧楀け璐? %v", pluginName, err))
					continue
				}
				modules = append(modules, pluginModules...)
			}
		}
	}

	return modules, nil
}

// scanModelModules 鎵弿妯″瀷妯″潡
func (g *GVAAnalyzer) scanModelModules(modelDir string) ([]PredesignedModuleInfo, error) {
	var modules []PredesignedModuleInfo

	if _, err := os.Stat(modelDir); os.IsNotExist(err) {
		return modules, nil // 鐩綍涓嶅瓨鍦紝杩斿洖绌哄垪琛?
	}

	entries, err := os.ReadDir(modelDir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			packageName := entry.Name()
			packagePath := filepath.Join(modelDir, packageName)

			// 鎵弿鍖呯洰褰曚笅鐨勬ā鍧?
			packageModules, err := g.scanModulesInDirectory(packagePath, packageName, "package")
			if err != nil {
				global.GVA_LOG.Warn(fmt.Sprintf("鎵弿鍖?%s 鐨勬ā鍧楀け璐? %v", packageName, err))
				continue
			}
			modules = append(modules, packageModules...)
		}
	}

	return modules, nil
}

// scanModulesInDirectory 鎵弿鐩綍涓殑妯″潡
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
				Description: fmt.Sprintf("%s妯″潡涓殑%s", packageName, moduleName),
			}
			modules = append(modules, module)
		}
	}

	return modules, nil
}
