// Package system 提供系统服务层的业务逻辑实现
//
// 本文件实现了自动代码生成历史记录的管理功能，包括：
// - 创建代码生成历史记录
// - 查询历史记录详情
// - 检测重复生成
// - 回滚代码生成操作
// - 删除历史记录
// - 获取历史记录列表
// - 删除数据库表
//
// 功能说明：
// 自动代码生成功能是gin-vue-admin框架的核心特性之一，允许用户通过可视化界面
// 配置数据库表结构，自动生成前后端代码。本服务层负责管理这些生成操作的历史记录，
// 并提供回滚功能，方便用户在生成代码后能够撤销操作。
package system

import (
	"context"
	"encoding/json"
	"fmt"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/utils/ast"
	"github.com/pkg/errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	common "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	request "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"go.uber.org/zap"
)

// AutocodeHistory 自动代码生成历史记录服务实例
//
// 功能说明：
// 提供全局可访问的服务实例，遵循gin-vue-admin的单例模式
//
// 使用示例：
//
//	// 在其他服务或控制器中使用
//	err := system.AutocodeHistory.Create(ctx, createInfo)
var AutocodeHistory = new(autoCodeHistory)

// autoCodeHistory 自动代码生成历史记录服务结构体
//
// 功能说明：
// 封装了所有与代码生成历史记录相关的业务逻辑方法
//
// 设计模式：
// 采用空结构体实现，所有方法都是值接收者，避免不必要的内存分配
type autoCodeHistory struct{}

// Create 创建代码生成器历史记录
//
// 功能说明：
// 当用户执行自动代码生成操作时，调用此方法保存生成操作的元数据信息，
// 包括生成的表名、结构体名、包名、模板文件路径、注入代码信息等。
// 这些信息用于后续的回滚操作。
//
// 参数说明：
// - ctx: 上下文对象，用于控制请求的生命周期和超时
// - info: 创建历史记录的请求参数，包含代码生成的所有元数据
//
// 返回值：
// - error: 如果创建失败，返回包装后的错误信息
//
// 实现细节：
// 1. 通过info.Create()方法将请求参数转换为数据库模型
// 2. 使用GORM的Create方法插入数据库记录
// 3. 使用errors.Wrap包装错误，保留原始错误信息
//
// 使用示例：
//
//	// 在代码生成完成后调用
//	createInfo := request.SysAutoHistoryCreate{
//	    Table: "sys_user",
//	    StructName: "SysUser",
//	    Package: "system",
//	    // ... 其他字段
//	}
//	err := system.AutocodeHistory.Create(ctx, createInfo)
//	if err != nil {
//	    // 处理错误
//	}
//
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (s *autoCodeHistory) Create(ctx context.Context, info request.SysAutoHistoryCreate) error {
	// 将请求参数转换为数据库模型
	create := info.Create()
	// 使用上下文和GORM创建记录
	err := global.GVA_DB.WithContext(ctx).Create(&create).Error
	if err != nil {
		// 包装错误信息，便于追踪和调试
		return errors.Wrap(err, "创建失败!")
	}
	return nil
}

// First 根据ID获取代码生成器历史的请求元数据
//
// 功能说明：
// 根据历史记录ID查询该次代码生成操作的原始请求参数（JSON格式），
// 用于重新生成代码或查看历史配置。
//
// 参数说明：
// - ctx: 上下文对象
// - info: 包含历史记录ID的查询参数
//
// 返回值：
// - string: 原始请求参数的JSON字符串
// - error: 查询失败时返回错误
//
// 实现细节：
// 1. 使用GORM的Pluck方法只查询request字段，提高查询效率
// 2. request字段存储的是代码生成时的完整配置信息（JSON格式）
//
// 使用示例：
//
//	// 查询历史记录
//	var req common.GetById
//	req.ID = 1
//	meta, err := system.AutocodeHistory.First(ctx, req)
//	if err != nil {
//	    // 处理错误
//	}
//	// meta 是JSON字符串，可以反序列化后重新使用
//	var config request.SysAutoHistoryCreate
//	json.Unmarshal([]byte(meta), &config)
//
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (s *autoCodeHistory) First(ctx context.Context, info common.GetById) (string, error) {
	var meta string
	// 使用Pluck只查询request字段，避免查询整个记录
	err := global.GVA_DB.WithContext(ctx).Model(model.SysAutoCodeHistory{}).Where("id = ?", info.ID).Pluck("request", &meta).Error
	if err != nil {
		return "", errors.Wrap(err, "获取失败!")
	}
	return meta, nil
}

// Repeat 检测代码生成是否重复
//
// 功能说明：
// 在生成代码之前，检查是否已经存在相同配置的未回滚历史记录。
// 用于防止重复生成相同的代码结构。
//
// 参数说明：
// - businessDB: 业务数据库名称
// - structName: 结构体名称（如：SysUser）
// - abbreviation: 结构体缩写（如：user）
// - Package: 包名（如：system）
//
// 返回值：
// - bool: true表示存在重复记录，false表示不存在重复
//
// 检测逻辑：
// 1. 检查相同业务数据库、相同包名下的记录
// 2. 结构体名或缩写名相同
// 3. flag=0表示未回滚的记录（flag=1表示已回滚）
//
// 使用示例：
//
//	// 在生成代码前检测
//	isRepeat := system.AutocodeHistory.Repeat("default", "SysUser", "user", "system")
//	if isRepeat {
//	    return errors.New("该结构体已存在，请先回滚后再生成")
//	}
//
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (s *autoCodeHistory) Repeat(businessDB, structName, abbreviation, Package string) bool {
	var count int64
	// 查询条件：
	// - business_db: 业务数据库名称必须相同
	// - struct_name OR abbreviation: 结构体名或缩写名相同
	// - package: 包名必须相同
	// - flag = 0: 只检查未回滚的记录
	global.GVA_DB.Model(&model.SysAutoCodeHistory{}).Where("business_db = ? and (struct_name = ? OR abbreviation = ?) and package = ? and flag = ?", businessDB, structName, abbreviation, Package, 0).Count(&count).Debug()
	// count > 0 表示存在重复记录
	return count > 0
}

// RollBack 回滚代码生成操作
//
// 功能说明：
// 这是代码生成回滚的核心方法，用于撤销一次代码生成操作。
// 回滚操作包括：
// 1. 删除导出的模板（如果存在）
// 2. 删除生成的API接口（可选）
// 3. 删除生成的菜单（可选）
// 4. 删除数据库表（可选）
// 5. 回滚注入到现有文件的代码
// 6. 移动或删除生成的文件
// 7. 标记历史记录为已回滚状态
//
// 参数说明：
// - ctx: 上下文对象
// - info: 回滚配置参数，包含ID和各项删除选项
//
// 返回值：
// - error: 回滚过程中的任何错误
//
// 回滚流程：
// 1. 查询历史记录
// 2. 删除导出模板（如果存在）
// 3. 根据配置删除API、菜单、数据表
// 4. 处理模板文件路径映射
// 5. 回滚注入的代码
// 6. 移动生成的文件到临时目录
// 7. 更新历史记录状态为已回滚
//
// 使用示例：
//
//	rollbackInfo := request.SysAutoHistoryRollBack{
//	    ID: 1,
//	    DeleteApi: true,    // 删除API
//	    DeleteMenu: true,   // 删除菜单
//	    DeleteTable: false, // 保留数据表
//	}
//	err := system.AutocodeHistory.RollBack(ctx, rollbackInfo)
//	if err != nil {
//	    // 处理错误
//	}
//
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (s *autoCodeHistory) RollBack(ctx context.Context, info request.SysAutoHistoryRollBack) error {
	// 步骤1: 查询历史记录
	var history model.SysAutoCodeHistory
	err := global.GVA_DB.Where("id = ?", info.ID).First(&history).Error
	if err != nil {
		return err
	}

	// 步骤2: 删除导出模板（如果存在）
	// ExportTemplateID不为0表示生成了导出模板
	if history.ExportTemplateID != 0 {
		err = global.GVA_DB.Delete(&model.SysExportTemplate{}, "id = ?", history.ExportTemplateID).Error
		if err != nil {
			return err
		}
	}

	// 步骤3: 根据配置删除API接口
	// DeleteApi为true时，删除本次生成的所有API接口
	if info.DeleteApi {
		ids := info.ApiIds(history)
		err = ApiServiceApp.DeleteApisByIds(ids)
		if err != nil {
			// API删除失败只记录日志，不中断回滚流程
			global.GVA_LOG.Error("ClearTag DeleteApiByIds:", zap.Error(err))
		}
	} // 清除API表

	// 步骤4: 根据配置删除菜单
	// DeleteMenu为true时，删除本次生成的菜单项
	if info.DeleteMenu {
		err = BaseMenuServiceApp.DeleteBaseMenu(int(history.MenuID))
		if err != nil {
			return errors.Wrap(err, "删除菜单失败!")
		}
	} // 清除菜单表

	// 步骤5: 根据配置删除数据表
	// DeleteTable为true时，删除本次生成的数据表
	if info.DeleteTable {
		err = s.DropTable(history.BusinessDB, history.Table)
		if err != nil {
			return errors.Wrap(err, "删除表失败!")
		}
	} // 删除表
	// 步骤6: 处理模板文件路径映射
	// 将相对路径转换为绝对路径，便于后续文件操作
	templates := make(map[string]string, len(history.Templates))
	for key, template := range history.Templates {
		// 处理key（文件路径）：移除server根路径前缀
		{
			server := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server)
			keys := strings.Split(key, "/")
			key = filepath.Join(keys...)
			key = strings.TrimPrefix(key, server)
		} // key处理完成

		// 处理value（模板路径）：根据文件扩展名确定完整路径
		{
			web := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.WebRoot())
			server := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server)
			slices := strings.Split(template, "/")
			template = filepath.Join(slices...)
			ext := path.Ext(template)
			// 根据文件扩展名确定是前端文件还是后端文件
			switch ext {
			case ".js", ".vue":
				// 前端文件：.js和.vue文件放在web目录
				template = filepath.Join(web, template)
			case ".go":
				// 后端文件：.go文件放在server目录
				template = filepath.Join(server, template)
			}
		} // value处理完成
		templates[key] = template
	}
	history.Templates = templates
	// 步骤7: 回滚注入到现有文件的代码
	// 代码生成时会在现有文件中注入代码（如路由注册、服务注册等），
	// 回滚时需要将这些注入的代码移除
	for key, value := range history.Injections {
		var injection ast.Ast
		// 根据注入类型反序列化为对应的AST结构体
		switch key {
		case ast.TypePackageApiEnter, ast.TypePackageRouterEnter, ast.TypePackageServiceEnter:
			// 包级别的入口注入，暂不处理（空case）

		case ast.TypePackageApiModuleEnter, ast.TypePackageRouterModuleEnter, ast.TypePackageServiceModuleEnter:
			// 包模块级别的入口注入
			var entity ast.PackageModuleEnter
			_ = json.Unmarshal([]byte(value), &entity)
			injection = &entity
		case ast.TypePackageInitializeGorm:
			// 包级别的GORM初始化注入
			var entity ast.PackageInitializeGorm
			_ = json.Unmarshal([]byte(value), &entity)
			injection = &entity
		case ast.TypePackageInitializeRouter:
			// 包级别的路由初始化注入
			var entity ast.PackageInitializeRouter
			_ = json.Unmarshal([]byte(value), &entity)
			injection = &entity
		case ast.TypePluginGen:
			// 插件生成注入
			var entity ast.PluginGen
			_ = json.Unmarshal([]byte(value), &entity)
			injection = &entity
		case ast.TypePluginApiEnter, ast.TypePluginRouterEnter, ast.TypePluginServiceEnter:
			// 插件入口注入
			var entity ast.PluginEnter
			_ = json.Unmarshal([]byte(value), &entity)
			injection = &entity
		case ast.TypePluginInitializeGorm:
			// 插件GORM初始化注入
			var entity ast.PluginInitializeGorm
			_ = json.Unmarshal([]byte(value), &entity)
			injection = &entity
		case ast.TypePluginInitializeRouter:
			// 插件路由初始化注入
			var entity ast.PluginInitializeRouter
			_ = json.Unmarshal([]byte(value), &entity)
			injection = &entity
		}
		// 如果无法识别注入类型，跳过
		if injection == nil {
			continue
		}
		// 解析目标文件
		file, _ := injection.Parse("", nil)
		if file != nil {
			// 执行回滚操作：移除注入的代码
			_ = injection.Rollback(file)
			// 格式化文件并写回
			err = injection.Format("", nil, file)
			if err != nil {
				return err
			}
			fmt.Printf("[filepath:%s]回滚注入代码成功!\n", key)
		}
	} // 清除注入代码完成
	// 步骤8: 移动生成的文件到临时目录
	// 不直接删除文件，而是移动到rm_file目录，便于恢复
	// 使用纳秒时间戳创建唯一的临时目录名，避免冲突
	removeBasePath := filepath.Join(global.GVA_CONFIG.AutoCode.Root, "rm_file", strconv.FormatInt(int64(time.Now().Nanosecond()), 10))
	for _, value := range history.Templates {
		// 只处理绝对路径的文件
		if !filepath.IsAbs(value) {
			continue
		}
		// 计算目标路径：保持相对路径结构
		removePath := filepath.Join(removeBasePath, strings.TrimPrefix(value, global.GVA_CONFIG.AutoCode.Root))
		// 移动文件到临时目录
		err = utils.FileMove(value, removePath)
		if err != nil {
			return errors.Wrapf(err, "[src:%s][dst:%s]文件移动失败!", value, removePath)
		}
	} // 移动文件完成

	// 步骤9: 更新历史记录状态为已回滚
	// flag=1表示该历史记录已被回滚
	err = global.GVA_DB.WithContext(ctx).Model(&model.SysAutoCodeHistory{}).Where("id = ?", info.ID).Update("flag", 1).Error
	if err != nil {
		return errors.Wrap(err, "更新失败!")
	}
	return nil
}

// Delete 删除历史记录
//
// 功能说明：
// 从数据库中永久删除指定的代码生成历史记录。
// 注意：此操作只删除数据库记录，不会删除已生成的文件或回滚代码。
// 如果需要撤销代码生成，应该使用RollBack方法。
//
// 参数说明：
// - ctx: 上下文对象
// - info: 包含要删除的历史记录ID
//
// 返回值：
// - error: 删除失败时返回错误
//
// 使用示例：
//
//	var req common.GetById
//	req.ID = 1
//	err := system.AutocodeHistory.Delete(ctx, req)
//	if err != nil {
//	    // 处理错误
//	}
//
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (s *autoCodeHistory) Delete(ctx context.Context, info common.GetById) error {
	// 使用GORM的Delete方法软删除或硬删除记录（取决于模型定义）
	err := global.GVA_DB.WithContext(ctx).Where("id = ?", info.Uint()).Delete(&model.SysAutoCodeHistory{}).Error
	if err != nil {
		return errors.Wrap(err, "删除失败!")
	}
	return nil
}

// GetList 获取代码生成历史记录列表（分页）
//
// 功能说明：
// 分页查询所有代码生成历史记录，支持分页参数，按更新时间倒序排列。
//
// 参数说明：
// - ctx: 上下文对象
// - info: 分页参数，包含页码、每页大小等
//
// 返回值：
// - list: 历史记录列表
// - total: 总记录数（用于前端分页计算）
// - err: 查询错误
//
// 实现细节：
// 1. 先查询总记录数
// 2. 应用分页作用域
// 3. 按更新时间倒序排列（最新的在前）
//
// 使用示例：
//
//	var pageInfo common.PageInfo
//	pageInfo.Page = 1
//	pageInfo.PageSize = 20
//	list, total, err := system.AutocodeHistory.GetList(ctx, pageInfo)
//	if err != nil {
//	    // 处理错误
//	}
//	// 返回给前端：{list: [...], total: 100}
//
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (s *autoCodeHistory) GetList(ctx context.Context, info common.PageInfo) (list []model.SysAutoCodeHistory, total int64, err error) {
	var entities []model.SysAutoCodeHistory
	db := global.GVA_DB.WithContext(ctx).Model(&model.SysAutoCodeHistory{})
	// 先查询总记录数
	err = db.Count(&total).Error
	if err != nil {
		return nil, total, err
	}
	// 应用分页并排序：按更新时间倒序
	err = db.Scopes(info.Paginate()).Order("updated_at desc").Find(&entities).Error
	return entities, total, err
}

// DropTable 删除指定的数据表
//
// 功能说明：
// 删除指定数据库中的指定数据表。用于代码生成回滚时删除自动创建的数据表。
//
// 参数说明：
// - BusinessDb: 业务数据库名称，如果为空字符串则使用默认数据库
// - tableName: 要删除的表名
//
// 返回值：
// - error: 删除失败时返回错误
//
// 实现细节：
// 1. 如果指定了业务数据库，使用对应的数据库连接
// 2. 如果未指定，使用默认的全局数据库连接
// 3. 直接执行SQL的DROP TABLE语句
//
// 注意事项：
// - 此操作不可逆，删除前请确认
// - 如果表不存在，可能会返回错误
// - 建议在事务中执行，确保数据一致性
//
// 使用示例：
//
//	// 删除默认数据库中的表
//	err := system.AutocodeHistory.DropTable("", "sys_user")
//
//	// 删除指定业务数据库中的表
//	err := system.AutocodeHistory.DropTable("business_db", "sys_order")
//
// @author: [piexlmax](https://github.com/piexlmax)
func (s *autoCodeHistory) DropTable(BusinessDb, tableName string) error {
	if BusinessDb != "" {
		// 使用指定的业务数据库连接
		return global.MustGetGlobalDBByDBName(BusinessDb).Exec("DROP TABLE " + tableName).Error
	} else {
		// 使用默认的全局数据库连接
		return global.GVA_DB.Exec("DROP TABLE " + tableName).Error
	}
}
