package mcpTool

/*
================================================================================
performAnalysis 函数详细说明
================================================================================

【功能概述】
performAnalysis 是 GVA 分析器的核心函数，用于全面分析系统当前状态，包括：
1. 扫描并清理空包（没有实际代码文件的包）
2. 扫描预设计模块（已存在但未注册的模块）
3. 获取系统字典信息
4. 返回完整的系统分析结果

【执行流程】
函数按照以下步骤执行：

步骤1: 获取数据库中的包信息
  - 从 sys_auto_code_packages 表获取所有包记录
  - 这些包可能是 package 类型或 plugin 类型

步骤2: 获取历史记录
  - 从 sys_auto_code_histories 表获取所有代码生成历史记录
  - 用于关联包和模块的关系

步骤3: 检查并清理空包
  - 遍历每个包，检查其文件夹是否为空（没有 .go 文件）
  - 如果为空包：
    * 删除空包文件夹（plugin 类型删除单个目录，package 类型删除多个目录）
    * 删除数据库中的包记录
    * 收集相关的历史记录ID和模块名称

步骤4: 清理脏历史记录
  - 删除与空包相关的历史记录
  - 清理相关的API和菜单记录（目前只记录日志，具体清理逻辑待实现）

步骤5: 扫描预设计模块
  - 扫描 plugin 目录下的模块（在 plugin/{name}/model/ 下查找 .go 文件）
  - 扫描 model 目录下的模块（在 model/{package}/ 下查找 .go 文件）
  - 这些模块是已存在但可能未在数据库中注册的模块

步骤6: 过滤已删除包的模块
  - 从预设计模块列表中移除属于已删除包的模块

步骤7: 构建清理消息
  - 如果有清理操作，生成清理信息消息

步骤8: 转换包信息
  - 将数据库模型转换为响应结构体

步骤9: 获取字典信息
  - 从 sys_dictionaries 表获取所有未删除的字典

步骤10: 构建并返回响应
  - 组装所有信息到 AnalyzeResponse 结构体

【Demo 示例】

示例1: 正常分析场景 - 发现并清理空包
----------------------------------------
场景描述:
  系统中有3个包，其中一个是空包（没有实际代码文件），需要清理。

输入:
  requirement: "我需要创建一个用户管理模块"

执行前状态:
  数据库中有3个包记录：
  - system (package类型) - 有实际代码文件
  - example (package类型) - 有实际代码文件
  - empty_pkg (package类型) - 空包，没有.go文件

  文件系统状态：
  - server/api/v1/empty_pkg/ 存在但为空（没有.go文件）
  - server/model/empty_pkg/ 存在但为空
  - server/router/empty_pkg/ 存在但为空
  - server/service/empty_pkg/ 存在但为空

执行过程:
  1. 查询数据库，获取3个包记录
  2. 查询历史记录，发现 empty_pkg 有1条历史记录（关联模块：empty_module）
  3. 检查 empty_pkg 包：
     - 调用 isPackageFolderEmpty("empty_pkg", "package")
     - 检查 server/api/v1/empty_pkg/ 目录
     - 递归检查所有子目录，未发现任何.go文件
     - 判定为空包
  4. 执行清理操作：
     - 删除 server/api/v1/empty_pkg/ 目录
     - 删除 server/model/empty_pkg/ 目录
     - 删除 server/router/empty_pkg/ 目录
     - 删除 server/service/empty_pkg/ 目录
     - 从数据库删除 empty_pkg 包记录
     - 收集相关历史记录ID和模块名称
  5. 删除脏历史记录：
     - 删除 empty_pkg 相关的历史记录
     - 记录清理日志
  6. 扫描预设计模块：
     - 扫描 plugin 目录，发现2个模块
     - 扫描 model 目录，发现3个模块
  7. 过滤模块：移除属于 empty_pkg 的模块
  8. 获取字典信息：查询到10个字典

输出结果:
  {
    "existingPackages": [
      {
        "packageName": "system",
        "template": "package",
        "label": "系统管理",
        "desc": "系统相关功能模块",
        "module": "system",
        "isEmpty": false
      },
      {
        "packageName": "example",
        "template": "package",
        "label": "示例模块",
        "desc": "示例功能演示",
        "module": "example",
        "isEmpty": false
      }
    ],
    "predesignedModules": [
      {
        "moduleName": "sys_user",
        "packageName": "system",
        "template": "package",
        "filePaths": ["/root/server/model/system/sys_user.go"],
        "description": "system模块中的sys_user"
      },
      {
        "moduleName": "exa_customer",
        "packageName": "example",
        "template": "package",
        "filePaths": ["/root/server/model/example/exa_customer.go"],
        "description": "example模块中的exa_customer"
      }
    ],
    "dictionaries": [
      {
        "id": 1,
        "name": "性别",
        "type": "gender",
        ...
      }
    ],
    "cleanupInfo": {
      "deletedPackages": ["empty_pkg"],
      "deletedModules": ["empty_module"],
      "cleanupMessage": "**系统清理完成**\n\n- 删除了 1 个空包: empty_pkg\n- 删除了 1 个相关模块: empty_module\n\n"
    }
  }

示例2: 无空包场景 - 正常分析
----------------------------------------
场景描述:
  系统中所有包都有实际代码，无需清理。

输入:
  requirement: "查看当前系统状态"

执行前状态:
  数据库中有2个包记录：
  - system (package类型) - 有实际代码
  - example (package类型) - 有实际代码

执行过程:
  1. 查询数据库，获取2个包记录
  2. 检查每个包：
     - system: 发现 server/api/v1/system/sys_user.go 等文件，不为空
     - example: 发现 server/api/v1/example/exa_customer.go 等文件，不为空
  3. 无需清理操作
  4. 扫描预设计模块，发现5个模块
  5. 获取字典信息，发现8个字典

输出结果:
  {
    "existingPackages": [
      {
        "packageName": "system",
        "template": "package",
        "label": "系统管理",
        "desc": "系统相关功能",
        "module": "system",
        "isEmpty": false
      },
      {
        "packageName": "example",
        "template": "package",
        "label": "示例",
        "desc": "示例功能",
        "module": "example",
        "isEmpty": false
      }
    ],
    "predesignedModules": [
      {
        "moduleName": "sys_user",
        "packageName": "system",
        "template": "package",
        "filePaths": ["/root/server/model/system/sys_user.go"],
        "description": "system模块中的sys_user"
      }
    ],
    "dictionaries": [...],
    "cleanupInfo": null
  }

示例3: Plugin 类型包清理
----------------------------------------
场景描述:
  发现一个 plugin 类型的空包，需要清理。

输入:
  requirement: "分析插件模块"

执行前状态:
  数据库中有1个 plugin 类型包：
  - test_plugin (plugin类型) - 空包

  文件系统状态：
  - server/plugin/test_plugin/ 目录存在但为空

执行过程:
  1. 查询数据库，获取 test_plugin 包记录
  2. 检查 test_plugin 包：
     - 调用 isPackageFolderEmpty("test_plugin", "plugin")
     - 检查 server/plugin/test_plugin/ 目录
     - 递归检查，未发现任何.go文件
     - 判定为空包
  3. 执行清理操作：
     - 删除整个 server/plugin/test_plugin/ 目录（plugin类型只需删除一个目录）
     - 从数据库删除 test_plugin 包记录
  4. 扫描预设计模块（无）
  5. 获取字典信息

输出结果:
  {
    "existingPackages": [],
    "predesignedModules": [],
    "dictionaries": [...],
    "cleanupInfo": {
      "deletedPackages": ["test_plugin"],
      "deletedModules": [],
      "cleanupMessage": "**系统清理完成**\n\n- 删除了 1 个空包: test_plugin\n\n"
    }
  }

示例4: 递归检查空包 - 深层目录结构
----------------------------------------
场景描述:
  包目录存在，有子目录结构，但所有目录都没有.go文件。

执行前状态:
  数据库中有1个包：
  - test_pkg (package类型)

  文件系统状态：
  - server/api/v1/test_pkg/
    - subdir1/
    - subdir2/
      - nested/
  （所有目录都只有空文件夹，没有任何.go文件）

执行过程:
  1. 检查 test_pkg 包：
     - 检查 server/api/v1/test_pkg/ 目录
     - 递归检查 subdir1/，未发现.go文件
     - 递归检查 subdir2/，未发现.go文件
     - 递归检查 subdir2/nested/，未发现.go文件
     - 判定为空包（hasGoFilesRecursive 返回 true）
  2. 执行清理操作：
     - 删除所有相关目录
     - 删除数据库记录

输出结果:
  {
    "existingPackages": [],
    "predesignedModules": [...],
    "dictionaries": [...],
    "cleanupInfo": {
      "deletedPackages": ["test_pkg"],
      "deletedModules": [],
      "cleanupMessage": "**系统清理完成**\n\n- 删除了 1 个空包: test_pkg\n\n"
    }
  }

示例5: 混合场景 - 多个空包和有效包
----------------------------------------
场景描述:
  系统中有多个包，部分为空包，需要清理。

执行前状态:
  数据库中有5个包：
  - system (package) - 有效
  - example (package) - 有效
  - empty1 (package) - 空包
  - empty2 (package) - 空包
  - test_plugin (plugin) - 空包

执行过程:
  1. 检查所有包，发现3个空包
  2. 逐个清理空包：
     - empty1: 删除4个目录 + 数据库记录
     - empty2: 删除4个目录 + 数据库记录
     - test_plugin: 删除1个目录 + 数据库记录
  3. 清理相关历史记录
  4. 扫描预设计模块
  5. 过滤已删除包的模块

输出结果:
  {
    "existingPackages": [
      {
        "packageName": "system",
        "template": "package",
        ...
      },
      {
        "packageName": "example",
        "template": "package",
        ...
      }
    ],
    "predesignedModules": [...],
    "dictionaries": [...],
    "cleanupInfo": {
      "deletedPackages": ["empty1", "empty2", "test_plugin"],
      "deletedModules": ["module1", "module2"],
      "cleanupMessage": "**系统清理完成**\n\n- 删除了 3 个空包: empty1, empty2, test_plugin\n- 删除了 2 个相关模块: module1, module2\n\n"
    }
  }

示例6: 错误处理场景 - 部分操作失败
----------------------------------------
场景描述:
  清理过程中部分操作失败，但函数继续执行。

执行过程:
  1. 检查 empty_pkg 包，判定为空包
  2. 尝试删除文件夹：
     - 删除 server/api/v1/empty_pkg/ 成功
     - 删除 server/model/empty_pkg/ 失败（权限问题）
     - 记录警告日志，继续执行
  3. 尝试删除数据库记录：成功
  4. 继续执行后续步骤
  5. 最终返回结果（包含部分清理信息）

说明:
  即使部分清理操作失败，函数也会继续执行并返回结果。
  失败的操作会记录在日志中，但不会中断整个流程。

【注意事项】
1. 函数会实际删除文件和数据库记录，请谨慎使用
2. 清理操作失败不会中断整个流程，只会记录警告日志
3. 预设计模块扫描失败不会影响主流程，会返回空列表
4. 字典信息获取失败也不会影响主流程
5. Package 类型的包会删除多个目录（api/v1, model, router, service）
6. Plugin 类型的包只删除单个目录（plugin/{name}）

【相关函数说明】
- isPackageFolderEmpty: 检查包文件夹是否为空（递归检查 .go 文件）
- removeEmptyPackageFolder: 删除空包文件夹（根据类型删除不同路径）
- scanPredesignedModules: 扫描预设计模块（plugin 和 model 目录）
- cleanupRelatedApiAndMenus: 清理相关API和菜单（目前只记录日志）

================================================================================
*/

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"

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

/*
================================================================================
performAnalysis 函数详细说明
================================================================================

【函数概述】
performAnalysis 是 GVA 分析器的核心函数，负责全面分析系统当前状态。
该函数会扫描系统中的包、模块和字典信息，同时自动清理空包（没有实际代码文件的包），
确保系统整洁，并返回完整的分析结果供后续使用。

【函数签名】
func (g *GVAAnalyzer) performAnalysis(ctx context.Context, req AnalyzeRequest) (*AnalyzeResponse, error)

【参数说明】
- ctx: 上下文对象，用于控制函数执行的生命周期和取消操作
- req: 分析请求，包含用户的需求描述（虽然当前实现中未直接使用，但保留用于未来扩展）

【返回值说明】
- *AnalyzeResponse: 分析结果，包含：
  * ExistingPackages: 现有有效包列表（已过滤空包）
  * PredesignedModules: 预设计模块列表（已存在但可能未注册的模块）
  * Dictionaries: 系统字典列表
  * CleanupInfo: 清理信息（如果有清理操作）
- error: 如果执行过程中出现严重错误则返回

【执行流程详解】

步骤1: 获取数据库中的包信息
----------------------------------------
从 sys_auto_code_packages 表查询所有包记录。
这些包可能是两种类型：
- "package" 类型：标准包，会在多个目录下创建文件（api/v1, model, router, service）
- "plugin" 类型：插件包，只在一个目录下创建文件（plugin/{name}）

代码示例：
  var packages []model.Sy /s/AutoCodePackage
  global.GVA_DB.Find(&packages)÷

步骤2: 获取历史记录
----------------------------------------
从 sys_auto_code_histories 表查询所有代码生成历史记录。
这些历史记录用于：
- 关联包和模块的关系
- 识别哪些模块属于哪个包
- 清理空包时，同时清理相关的历史记录

代码示例：
  var histories []model.SysAutoCodeHistory
  global.GVA_DB.Find(&histories)

步骤3: 检查并清理空包（核心逻辑）
----------------------------------------
遍历每个包，执行以下检查：

3.1 检查包文件夹是否为空
  - 调用 isPackageFolderEmpty() 函数
  - 对于 "plugin" 类型：检查 server/plugin/{packageName}/ 目录
  - 对于 "package" 类型：检查 server/api/v1/{packageName}/ 目录
  - 递归检查目录及其子目录中是否有 .go 文件
  - 如果文件夹不存在或没有任何 .go 文件，则判定为空包

3.2 如果发现空包，执行清理操作：
  a) 删除空包文件夹
     - Plugin 类型：删除 server/plugin/{packageName}/ 整个目录
     - Package 类型：删除以下4个目录：
       * server/api/v1/{packageName}/
       * server/model/{packageName}/
       * server/router/{packageName}/
       * server/service/{packageName}/

  b) 删除数据库中的包记录
     - 从 sys_auto_code_packages 表中删除该包记录

  c) 收集相关的历史记录ID和模块名称
     - 遍历所有历史记录，找出属于该包的历史记录
     - 收集这些历史记录的ID和对应的模块名称（StructName）

3.3 如果包不为空，则加入有效包列表

步骤4: 清理脏历史记录
----------------------------------------
4.1 识别脏历史记录
  - 遍历所有历史记录
  - 找出属于已删除空包的历史记录ID
  - 这些历史记录被称为"脏历史记录"

4.2 删除脏历史记录
  - 从 sys_auto_code_histories 表中批量删除这些历史记录
  - 使用 SQL: DELETE FROM sys_auto_code_histories WHERE id IN (dirtyHistoryIDs)

4.3 清理相关的API和菜单记录
  - 调用 cleanupRelatedApiAndMenus() 函数
  - 目前只记录日志，具体清理逻辑待实现
  - 未来可以调用 service 层的方法清理相关的 API 和菜单

步骤5: 扫描预设计模块
----------------------------------------
扫描文件系统中已存在但可能未在数据库中注册的模块。

5.1 扫描 Plugin 模块
  - 扫描 server/plugin/ 目录下的所有插件
  - 对每个插件，查找其 model/ 目录
  - 扫描 model/ 目录下的所有 .go 文件
  - 每个 .go 文件对应一个模块

5.2 扫描 Model 模块
  - 扫描 server/model/ 目录下的所有包目录
  - 对每个包目录，扫描其中的所有 .go 文件
  - 每个 .go 文件对应一个模块

5.3 构建预设计模块信息
  - 模块名称：文件名（去掉 .go 后缀）
  - 包名：所属的包或插件名称
  - 模板类型："plugin" 或 "package"
  - 文件路径：.go 文件的完整路径
  - 描述：自动生成的描述信息

步骤6: 过滤已删除包的模块
----------------------------------------
从预设计模块列表中移除属于已删除包的模块。
这样可以确保返回的模块列表只包含有效的模块。

步骤7: 构建清理消息
----------------------------------------
如果有清理操作（删除了空包或模块），则构建清理消息：
- 显示删除了多少个空包
- 显示删除了多少个相关模块
- 消息格式为 Markdown 格式，便于显示

步骤8: 转换包信息
----------------------------------------
将数据库模型（model.SysAutoCodePackage）转换为响应结构体（PackageInfo）。
所有返回的包都标记为 IsEmpty: false（因为空包已经被过滤掉）。

步骤9: 获取字典信息
----------------------------------------
从 sys_dictionaries 表查询所有未删除的字典记录。
查询条件：deleted_at IS NULL（软删除机制）

步骤10: 构建并返回响应
----------------------------------------
将所有收集到的信息组装到 AnalyzeResponse 结构体中：
- ExistingPackages: 有效包列表
- PredesignedModules: 过滤后的预设计模块列表
- Dictionaries: 字典列表
- CleanupInfo: 清理信息（如果有清理操作）

【错误处理策略】
1. 数据库查询失败：返回错误，中断执行
2. 检查包是否为空失败：记录警告日志，跳过该包，继续处理其他包
3. 删除空包文件夹失败：记录警告日志，继续执行
4. 删除数据库记录失败：记录警告日志，继续执行
5. 扫描预设计模块失败：记录警告日志，返回空列表，不影响主流程
6. 获取字典信息失败：记录警告日志，返回空列表，不影响主流程

这种错误处理策略确保了即使部分操作失败，函数也能尽可能完成主要任务。

【性能考虑】
1. 数据库查询：使用 Find() 一次性查询所有记录，避免多次查询
2. 文件系统操作：递归检查目录，但会在找到第一个 .go 文件时立即返回
3. 批量删除：使用 IN 子句批量删除历史记录，而不是逐条删除

【使用场景】
1. 系统初始化时，清理历史遗留的空包
2. 代码生成前，分析当前系统状态
3. 系统维护时，检查并清理无效的包和模块
4. 获取系统概览信息，用于展示给用户

【注意事项】
⚠️ 重要：此函数会实际删除文件和数据库记录，请谨慎使用！
1. 删除操作不可逆，建议在执行前备份
2. 空包的判定标准是：目录不存在或没有任何 .go 文件
3. Package 类型和 Plugin 类型的包处理方式不同
4. 清理操作失败不会中断整个流程，只会记录警告日志
5. 预设计模块扫描失败不会影响主流程

================================================================================
*/
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
