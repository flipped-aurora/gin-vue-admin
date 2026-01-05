// Package system 提供系统API管理相关的服务层功能
//
// 功能说明：
// 本包实现了API接口的CRUD操作、API同步、API分组、权限管理等核心功能。
// 主要用于管理系统中的所有API接口，包括：
// - API的创建、查询、更新、删除
// - API与路由的同步（对比内存中的路由和数据库中的API）
// - API分组管理
// - API忽略列表管理（某些API不需要权限控制）
// - 与Casbin权限系统的集成
//
// 设计模式：
// 采用服务层模式，将业务逻辑封装在ApiService结构体中，提供统一的API管理接口。
//
// 使用示例：
//
//	// 创建API
//	api := system.SysApi{
//	    Path: "/api/v1/user",
//	    Method: "POST",
//	    Description: "创建用户",
//	    ApiGroup: "用户管理",
//	}
//	err := ApiServiceApp.CreateApi(api)
//
//	// 获取API列表
//	list, total, err := ApiServiceApp.GetAPIInfoList(api, pageInfo, "id", true)
package system

import (
	"errors"
	"fmt"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"gorm.io/gorm"
)

// ApiService API服务结构体
//
// 功能说明：
// 封装了所有API管理相关的业务逻辑方法，包括CRUD操作、同步、分组等功能。
//
// 设计说明：
// 使用空结构体作为接收器，所有方法都是结构体方法，便于统一管理和扩展。
type ApiService struct{}

// ApiServiceApp API服务的全局实例
//
// 功能说明：
// 提供全局单例，方便在整个应用中访问API服务。
//
// 使用示例：
//
//	err := ApiServiceApp.CreateApi(api)
var ApiServiceApp = new(ApiService)

// CreateApi 创建新的API接口
//
// 功能说明：
// 在数据库中创建一条新的API记录。在创建前会检查是否已存在相同的API路径和请求方法组合，
// 如果已存在则返回错误，避免重复创建。
//
// 参数说明：
//   - api: 要创建的API对象，包含以下字段：
//   - Path: API路径，如 "/api/v1/user"
//   - Method: HTTP方法，如 "GET", "POST", "PUT", "DELETE"
//   - Description: API描述信息
//   - ApiGroup: API所属分组，如 "用户管理"
//
// 返回值：
//   - err: 如果创建成功返回nil，如果已存在相同API或数据库操作失败返回错误
//
// 使用示例：
//
//	api := system.SysApi{
//	    Path:        "/api/v1/user",
//	    Method:      "POST",
//	    Description: "创建用户接口",
//	    ApiGroup:    "用户管理",
//	}
//	err := ApiServiceApp.CreateApi(api)
//	if err != nil {
//	    log.Printf("创建API失败: %v", err)
//	}
//
// 注意事项：
//   - API的唯一性由Path和Method的组合决定
//   - 如果Path和Method的组合已存在，会返回"存在相同api"错误
//   - 创建成功后，该API可以被用于权限管理
//
// @author: [piexlmax](https://github.com/piexlmax)
// @function: CreateApi
// @description: 新增基础api
// @param: api model.SysApi
// @return: err error
func (apiService *ApiService) CreateApi(api system.SysApi) (err error) {
	// 检查是否已存在相同的API（路径+方法组合）
	// 使用First查询，如果找到记录说明已存在
	if !errors.Is(global.GVA_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同api")
	}
	// 创建新的API记录
	return global.GVA_DB.Create(&api).Error
}

// GetApiGroups 获取所有API分组信息
//
// 功能说明：
// 从数据库中查询所有API，提取出所有的API分组，并建立路径前缀到分组的映射关系。
// 主要用于前端展示API分组列表，以及根据路径快速查找对应的分组。
//
// 返回值：
//   - groups: API分组名称数组，如 ["用户管理", "角色管理", "菜单管理"]
//   - groupApiMap: 路径前缀到分组的映射表，key为路径的第一段（如"api"），value为分组名
//     例如：map["api"] = "用户管理"
//   - err: 查询失败时返回错误
//
// 使用示例：
//
//	groups, groupMap, err := ApiServiceApp.GetApiGroups()
//	if err != nil {
//	    log.Printf("获取API分组失败: %v", err)
//	    return
//	}
//	// 打印所有分组
//	for _, group := range groups {
//	    fmt.Printf("分组: %s\n", group)
//	}
//	// 根据路径前缀查找分组
//	if group, ok := groupMap["api"]; ok {
//	    fmt.Printf("路径前缀'api'对应的分组: %s\n", group)
//	}
//
// 实现逻辑：
//  1. 查询数据库中所有API记录
//  2. 遍历所有API，提取唯一的ApiGroup值
//  3. 解析API路径，取第一段作为key，建立路径前缀到分组的映射
//  4. 返回分组列表和映射表
//
// 注意事项：
//   - 如果路径格式不规范（如不以"/"开头），pathArr[1]可能越界
//   - 分组名称可能为空字符串，需要前端过滤
func (apiService *ApiService) GetApiGroups() (groups []string, groupApiMap map[string]string, err error) {
	var apis []system.SysApi
	// 查询所有API记录
	err = global.GVA_DB.Find(&apis).Error
	if err != nil {
		return
	}
	// 初始化映射表
	groupApiMap = make(map[string]string, 0)

	// 遍历所有API，提取分组信息
	for i := range apis {
		// 解析API路径，按"/"分割
		// 例如："/api/v1/user" -> ["", "api", "v1", "user"]
		pathArr := strings.Split(apis[i].Path, "/")

		// 检查当前API的分组是否已经存在于groups数组中
		newGroup := true
		for i2 := range groups {
			if groups[i2] == apis[i].ApiGroup {
				newGroup = false
				break
			}
		}

		// 如果是新分组，添加到groups数组
		if newGroup {
			groups = append(groups, apis[i].ApiGroup)
		}

		// 建立路径第一段到分组的映射
		// 例如："/api/v1/user" -> pathArr[1] = "api"
		if len(pathArr) > 1 {
			groupApiMap[pathArr[1]] = apis[i].ApiGroup
		}
	}
	return
}

// SyncApi 同步API接口（对比内存路由和数据库API）
//
// 功能说明：
// 对比内存中注册的路由（global.GVA_ROUTERS）和数据库中的API记录，找出差异：
// 1. 新增的API：内存中存在但数据库中不存在的路由
// 2. 待删除的API：数据库中存在但内存中不存在的路由（可能已删除或修改）
// 3. 忽略的API：在忽略列表中的API（不需要权限控制）
//
// 使用场景：
// - 系统启动后，需要将新添加的路由同步到数据库
// - 代码中删除了某些路由，需要从数据库中清理
// - 管理员可以通过前端界面查看差异，选择性地同步
//
// 返回值：
//   - newApis: 需要新增到数据库的API列表（内存中有，数据库中没有）
//   - deleteApis: 需要从数据库删除的API列表（数据库中有，内存中没有）
//   - ignoreApis: 忽略列表中的API（转换为SysApi格式返回，用于前端展示）
//   - err: 查询失败时返回错误
//
// 使用示例：
//
//	newApis, deleteApis, ignoreApis, err := ApiServiceApp.SyncApi()
//	if err != nil {
//	    log.Printf("同步API失败: %v", err)
//	    return
//	}
//	fmt.Printf("需要新增的API数量: %d\n", len(newApis))
//	fmt.Printf("需要删除的API数量: %d\n", len(deleteApis))
//	fmt.Printf("忽略的API数量: %d\n", len(ignoreApis))
//
// 实现逻辑：
//  1. 查询数据库中的所有API记录
//  2. 查询数据库中的所有忽略API记录
//  3. 遍历内存中的路由（global.GVA_ROUTERS），过滤掉忽略的API，得到有效的路由列表
//  4. 对比有效路由和数据库API，找出新增的API（路由有但数据库没有）
//  5. 对比数据库API和有效路由，找出待删除的API（数据库有但路由没有）
//
// 注意事项：
//   - 此方法只返回差异，不会实际修改数据库，需要调用EnterSyncApi来执行同步
//   - 忽略的API不会出现在newApis或deleteApis中
//   - 新增的API默认Description和ApiGroup为空，需要后续手动补充
func (apiService *ApiService) SyncApi() (newApis, deleteApis, ignoreApis []system.SysApi, err error) {
	// 初始化返回数组
	newApis = make([]system.SysApi, 0)
	deleteApis = make([]system.SysApi, 0)
	ignoreApis = make([]system.SysApi, 0)

	// 步骤1: 查询数据库中所有API记录
	var apis []system.SysApi
	err = global.GVA_DB.Find(&apis).Error
	if err != nil {
		return
	}

	// 步骤2: 查询数据库中所有忽略的API记录
	var ignores []system.SysIgnoreApi
	err = global.GVA_DB.Find(&ignores).Error
	if err != nil {
		return
	}

	// 步骤3: 将忽略的API转换为SysApi格式，用于返回给前端展示
	for i := range ignores {
		ignoreApis = append(ignoreApis, system.SysApi{
			Path:        ignores[i].Path,
			Description: "",
			ApiGroup:    "",
			Method:      ignores[i].Method,
		})
	}

	// 步骤4: 从内存路由中提取有效的API（排除忽略的API）
	var cacheApis []system.SysApi
	for i := range global.GVA_ROUTERS {
		// 检查当前路由是否在忽略列表中
		ignoresFlag := false
		for j := range ignores {
			if ignores[j].Path == global.GVA_ROUTERS[i].Path && ignores[j].Method == global.GVA_ROUTERS[i].Method {
				ignoresFlag = true
				break
			}
		}
		// 如果不在忽略列表中，添加到有效API列表
		if !ignoresFlag {
			cacheApis = append(cacheApis, system.SysApi{
				Path:   global.GVA_ROUTERS[i].Path,
				Method: global.GVA_ROUTERS[i].Method,
			})
		}
	}

	// 步骤5: 找出需要新增的API（内存中有，数据库中没有）
	for i := range cacheApis {
		var flag bool
		// 检查当前内存API是否已存在于数据库中
		for j := range apis {
			if cacheApis[i].Path == apis[j].Path && cacheApis[i].Method == apis[j].Method {
				flag = true
				break
			}
		}
		// 如果不存在于数据库，添加到新增列表
		if !flag {
			newApis = append(newApis, system.SysApi{
				Path:        cacheApis[i].Path,
				Description: "",
				ApiGroup:    "",
				Method:      cacheApis[i].Method,
			})
		}
	}

	// 步骤6: 找出需要删除的API（数据库中有，内存中没有）
	for i := range apis {
		var flag bool
		// 检查当前数据库API是否还存在于内存路由中
		for j := range cacheApis {
			if cacheApis[j].Path == apis[i].Path && cacheApis[j].Method == apis[i].Method {
				flag = true
				break
			}
		}
		// 如果不存在于内存路由，添加到删除列表
		if !flag {
			deleteApis = append(deleteApis, apis[i])
		}
	}
	return
}

// IgnoreApi 添加或移除API到忽略列表
//
// 功能说明：
// 管理API忽略列表，某些API不需要进行权限控制（如登录接口、健康检查接口等）。
// 当Flag为true时，将API添加到忽略列表；当Flag为false时，从忽略列表中移除。
//
// 参数说明：
//   - ignoreApi: 忽略API对象，包含以下字段：
//   - Path: API路径
//   - Method: HTTP方法
//   - Flag: true表示添加到忽略列表，false表示从忽略列表移除
//
// 返回值：
//   - err: 操作失败时返回错误
//
// 使用示例：
//
//	// 添加API到忽略列表（不需要权限控制）
//	ignoreApi := system.SysIgnoreApi{
//	    Path:   "/api/base/login",
//	    Method: "POST",
//	    Flag:   true,
//	}
//	err := ApiServiceApp.IgnoreApi(ignoreApi)
//
//	// 从忽略列表移除API（需要权限控制）
//	ignoreApi.Flag = false
//	err = ApiServiceApp.IgnoreApi(ignoreApi)
//
// 使用场景：
//   - 登录、注册等公开接口不需要权限控制
//   - 健康检查、Swagger文档等接口不需要权限控制
//   - 某些内部接口可以临时忽略权限检查
//
// 注意事项：
//   - 使用Unscoped().Delete()确保真正删除记录（包括软删除的记录）
//   - 忽略的API在SyncApi时会被过滤，不会出现在同步列表中
func (apiService *ApiService) IgnoreApi(ignoreApi system.SysIgnoreApi) (err error) {
	// Flag为true表示添加到忽略列表
	if ignoreApi.Flag {
		return global.GVA_DB.Create(&ignoreApi).Error
	}
	// Flag为false表示从忽略列表移除
	// 使用Unscoped()确保真正删除（包括软删除的记录）
	return global.GVA_DB.Unscoped().Delete(&ignoreApi, "path = ? AND method = ?", ignoreApi.Path, ignoreApi.Method).Error
}

// EnterSyncApi 执行API同步操作（事务方式）
//
// 功能说明：
// 在事务中执行API同步操作，包括：
// 1. 批量创建新增的API
// 2. 批量删除待删除的API，并清理相关的Casbin权限规则
//
// 参数说明：
//   - syncApis: 同步API对象，包含：
//   - NewApis: 需要新增的API列表
//   - DeleteApis: 需要删除的API列表
//
// 返回值：
//   - err: 同步失败时返回错误（事务会回滚）
//
// 使用示例：
//
//	syncApis := systemRes.SysSyncApis{
//	    NewApis: []system.SysApi{
//	        {Path: "/api/v1/new", Method: "GET", Description: "新接口", ApiGroup: "测试"},
//	    },
//	    DeleteApis: []system.SysApi{
//	        {Path: "/api/v1/old", Method: "POST"},
//	    },
//	}
//	err := ApiServiceApp.EnterSyncApi(syncApis)
//	if err != nil {
//	    log.Printf("同步API失败: %v", err)
//	}
//
// 实现逻辑：
//  1. 开启数据库事务
//  2. 如果有新增API，批量创建
//  3. 遍历待删除的API，先清理Casbin权限规则，再删除API记录
//  4. 如果任何步骤失败，事务回滚
//
// 注意事项：
//   - 使用事务确保数据一致性，要么全部成功，要么全部回滚
//   - 删除API时会同时清理Casbin权限规则，避免产生孤立数据
//   - ClearCasbin的第一个参数1表示v1（subject），第二个和第三个参数是path和method
func (apiService *ApiService) EnterSyncApi(syncApis systemRes.SysSyncApis) (err error) {
	// 使用事务确保数据一致性
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var txErr error

		// 步骤1: 批量创建新增的API
		if len(syncApis.NewApis) > 0 {
			txErr = tx.Create(&syncApis.NewApis).Error
			if txErr != nil {
				return txErr
			}
		}

		// 步骤2: 批量删除待删除的API
		for i := range syncApis.DeleteApis {
			// 先清理Casbin权限规则（避免产生孤立数据）
			CasbinServiceApp.ClearCasbin(1, syncApis.DeleteApis[i].Path, syncApis.DeleteApis[i].Method)
			// 再删除API记录
			txErr = tx.Delete(&system.SysApi{}, "path = ? AND method = ?", syncApis.DeleteApis[i].Path, syncApis.DeleteApis[i].Method).Error
			if txErr != nil {
				return txErr
			}
		}
		return nil
	})
}

// DeleteApi 删除单个API接口
//
// 功能说明：
// 根据API的ID删除指定的API记录，同时清理相关的Casbin权限规则。
// 删除前会先查询API是否存在，避免删除不存在的记录。
//
// 参数说明：
//   - api: 要删除的API对象，只需要ID字段有效
//
// 返回值：
//   - err: 删除失败时返回错误
//
// 使用示例：
//
//	api := system.SysApi{ID: 123}
//	err := ApiServiceApp.DeleteApi(api)
//	if err != nil {
//	    if errors.Is(err, gorm.ErrRecordNotFound) {
//	        log.Println("API不存在")
//	    } else {
//	        log.Printf("删除API失败: %v", err)
//	    }
//	}
//
// 实现逻辑：
//  1. 根据ID查询API记录
//  2. 如果记录不存在，返回错误
//  3. 删除API记录（软删除）
//  4. 清理Casbin权限规则
//
// 注意事项：
//   - 删除API时会同时清理Casbin权限规则，确保权限数据一致性
//   - 使用软删除，记录不会真正从数据库删除，只是标记为已删除
//   - 如果API不存在，会返回gorm.ErrRecordNotFound错误
//
// @author: [piexlmax](https://github.com/piexlmax)
// @function: DeleteApi
// @description: 删除基础api
// @param: api model.SysApi
// @return: err error
func (apiService *ApiService) DeleteApi(api system.SysApi) (err error) {
	var entity system.SysApi
	// 步骤1: 根据ID查询API记录
	err = global.GVA_DB.First(&entity, "id = ?", api.ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// API记录不存在，返回错误
		return err
	}

	// 步骤2: 删除API记录（软删除）
	err = global.GVA_DB.Delete(&entity).Error
	if err != nil {
		return err
	}

	// 步骤3: 清理Casbin权限规则（避免产生孤立数据）
	CasbinServiceApp.ClearCasbin(1, entity.Path, entity.Method)
	return nil
}

// GetAPIInfoList 分页查询API列表（支持多条件筛选和排序）
//
// 功能说明：
// 根据查询条件分页获取API列表，支持按路径、描述、方法、分组进行模糊或精确查询，
// 支持按指定字段排序（升序或降序）。
//
// 参数说明：
//   - api: 查询条件对象，各字段作为筛选条件：
//   - Path: 路径模糊查询（LIKE查询）
//   - Description: 描述模糊查询（LIKE查询）
//   - Method: 方法精确查询（=查询）
//   - ApiGroup: 分组精确查询（=查询）
//   - info: 分页信息
//   - Page: 页码，从1开始
//   - PageSize: 每页大小
//   - order: 排序字段，可选值：id, path, api_group, description, method
//   - desc: 是否降序，true为降序，false为升序
//
// 返回值：
//   - list: API列表（[]system.SysApi）
//   - total: 符合条件的总记录数（用于前端分页计算）
//   - err: 查询失败时返回错误
//
// 使用示例：
//
//	// 查询路径包含"user"的API，按id降序，第1页，每页20条
//	api := system.SysApi{Path: "user"}
//	pageInfo := request.PageInfo{Page: 1, PageSize: 20}
//	list, total, err := ApiServiceApp.GetAPIInfoList(api, pageInfo, "id", true)
//	if err != nil {
//	    log.Printf("查询失败: %v", err)
//	    return
//	}
//	apiList := list.([]system.SysApi)
//	fmt.Printf("共找到 %d 条记录，当前页: %d 条\n", total, len(apiList))
//
//	// 查询POST方法的API，按路径升序
//	api = system.SysApi{Method: "POST"}
//	list, total, err = ApiServiceApp.GetAPIInfoList(api, pageInfo, "path", false)
//
// 实现逻辑：
//  1. 构建基础查询（Model）
//  2. 根据条件添加WHERE子句（Path、Description、Method、ApiGroup）
//  3. 统计总记录数（Count）
//  4. 应用分页（Limit和Offset）
//  5. 应用排序（Order）
//  6. 执行查询（Find）
//
// 注意事项：
//   - 排序字段必须是指定的几个字段之一，否则返回错误
//   - 默认排序为"id desc"（按ID降序）
//   - Path和Description使用LIKE模糊查询，Method和ApiGroup使用精确查询
//   - 空字符串条件会被忽略（不添加WHERE子句）
//
// @author: [piexlmax](https://github.com/piexlmax)
// @function: GetAPIInfoList
// @description: 分页获取数据,
// @param: api model.SysApi, info request.PageInfo, order string, desc bool
// @return: list interface{}, total int64, err error
func (apiService *ApiService) GetAPIInfoList(api system.SysApi, info request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	// 计算分页参数
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	// 构建查询对象
	db := global.GVA_DB.Model(&system.SysApi{})
	var apiList []system.SysApi

	// 添加查询条件：路径模糊查询
	if api.Path != "" {
		db = db.Where("path LIKE ?", "%"+api.Path+"%")
	}

	// 添加查询条件：描述模糊查询
	if api.Description != "" {
		db = db.Where("description LIKE ?", "%"+api.Description+"%")
	}

	// 添加查询条件：方法精确查询
	if api.Method != "" {
		db = db.Where("method = ?", api.Method)
	}

	// 添加查询条件：分组精确查询
	if api.ApiGroup != "" {
		db = db.Where("api_group = ?", api.ApiGroup)
	}

	// 统计总记录数（用于前端分页计算）
	err = db.Count(&total).Error
	if err != nil {
		return apiList, total, err
	}

	// 应用分页
	db = db.Limit(limit).Offset(offset)

	// 构建排序字符串
	OrderStr := "id desc" // 默认按ID降序
	if order != "" {
		// 定义允许的排序字段（防止SQL注入）
		orderMap := make(map[string]bool, 5)
		orderMap["id"] = true
		orderMap["path"] = true
		orderMap["api_group"] = true
		orderMap["description"] = true
		orderMap["method"] = true

		// 验证排序字段是否合法
		if !orderMap[order] {
			err = fmt.Errorf("非法的排序字段: %v", order)
			return apiList, total, err
		}

		OrderStr = order
		if desc {
			OrderStr = order + " desc"
		}
	}

	// 执行查询
	err = db.Order(OrderStr).Find(&apiList).Error
	return apiList, total, err
}

// GetAllApis 获取所有API（根据权限ID过滤）
//
// 功能说明：
// 获取所有API列表，如果启用了严格权限模式（UseStrictAuth），则只返回当前角色有权限的API。
// 如果未启用严格权限模式，或者角色没有父角色，则返回所有API。
//
// 参数说明：
//   - authorityID: 角色/权限ID，用于过滤API
//
// 返回值：
//   - apis: API列表，如果启用严格权限则只返回有权限的API
//   - err: 查询失败时返回错误
//
// 使用示例：
//
//	// 获取角色ID为1的所有API
//	apis, err := ApiServiceApp.GetAllApis(1)
//	if err != nil {
//	    log.Printf("获取API失败: %v", err)
//	    return
//	}
//	fmt.Printf("角色1共有 %d 个API权限\n", len(apis))
//
// 实现逻辑：
//  1. 获取角色的父角色ID
//  2. 查询所有API（按ID降序）
//  3. 如果未启用严格权限或没有父角色，直接返回所有API
//  4. 如果启用严格权限，从Casbin获取角色的权限路径列表
//  5. 过滤API，只返回角色有权限的API
//
// 使用场景：
//   - 前端权限管理界面，显示角色可用的API列表
//   - 角色权限配置时，只显示该角色有权限的API
//   - 严格权限模式下，确保用户只能看到自己有权限的API
//
// 注意事项：
//   - 严格权限模式由配置项global.GVA_CONFIG.System.UseStrictAuth控制
//   - 如果角色没有父角色（parentAuthorityID == 0），返回所有API
//   - 权限检查基于Casbin策略，需要确保Casbin规则已正确配置
//
// @author: [piexlmax](https://github.com/piexlmax)
// @function: GetAllApis
// @description: 获取所有的api
// @return:  apis []model.SysApi, err error
func (apiService *ApiService) GetAllApis(authorityID uint) (apis []system.SysApi, err error) {
	// 步骤1: 获取角色的父角色ID
	parentAuthorityID, err := AuthorityServiceApp.GetParentAuthorityID(authorityID)
	if err != nil {
		return nil, err
	}

	// 步骤2: 查询所有API（按ID降序）
	err = global.GVA_DB.Order("id desc").Find(&apis).Error

	// 步骤3: 如果未启用严格权限或没有父角色，直接返回所有API
	if parentAuthorityID == 0 || !global.GVA_CONFIG.System.UseStrictAuth {
		return
	}

	// 步骤4: 启用严格权限模式，从Casbin获取角色的权限路径列表
	paths := CasbinServiceApp.GetPolicyPathByAuthorityId(authorityID)

	// 步骤5: 过滤API，只返回角色有权限的API
	// 遍历所有API，检查是否在角色的权限路径列表中
	var authApis []system.SysApi
	for i := range apis {
		for j := range paths {
			// 如果API的路径和方法都在权限列表中，添加到结果
			if paths[j].Path == apis[i].Path && paths[j].Method == apis[i].Method {
				authApis = append(authApis, apis[i])
				break // 找到匹配项后跳出内层循环
			}
		}
	}
	return authApis, err
}

// GetApiById 根据ID获取单个API详情
//
// 功能说明：
// 根据API的ID查询单个API记录的详细信息。
//
// 参数说明：
//   - id: API的主键ID
//
// 返回值：
//   - api: API对象，包含路径、方法、描述、分组等信息
//   - err: 查询失败时返回错误，如果API不存在返回gorm.ErrRecordNotFound
//
// 使用示例：
//
//	api, err := ApiServiceApp.GetApiById(123)
//	if err != nil {
//	    if errors.Is(err, gorm.ErrRecordNotFound) {
//	        log.Println("API不存在")
//	    } else {
//	        log.Printf("查询失败: %v", err)
//	    }
//	    return
//	}
//	fmt.Printf("API路径: %s, 方法: %s\n", api.Path, api.Method)
//
// 注意事项：
//   - 如果API不存在，会返回gorm.ErrRecordNotFound错误
//   - 使用First方法，如果记录不存在会返回错误
//
// @author: [piexlmax](https://github.com/piexlmax)
// @function: GetApiById
// @description: 根据id获取api
// @param: id float64
// @return: api model.SysApi, err error
func (apiService *ApiService) GetApiById(id int) (api system.SysApi, err error) {
	// 根据ID查询API记录
	err = global.GVA_DB.First(&api, "id = ?", id).Error
	return
}

// UpdateApi 更新API接口信息
//
// 功能说明：
// 根据API的ID更新API信息。如果更新了路径或方法，会检查是否与其他API冲突，
// 并同步更新Casbin权限规则中的路径和方法。
//
// 参数说明：
//   - api: 要更新的API对象，必须包含ID字段，其他字段为要更新的值
//
// 返回值：
//   - err: 更新失败时返回错误
//
// 使用示例：
//
//	api := system.SysApi{
//	    ID:          123,
//	    Path:        "/api/v1/user/update",
//	    Method:      "PUT",
//	    Description: "更新用户信息",
//	    ApiGroup:    "用户管理",
//	}
//	err := ApiServiceApp.UpdateApi(api)
//	if err != nil {
//	    log.Printf("更新API失败: %v", err)
//	}
//
// 实现逻辑：
//  1. 查询原API记录
//  2. 如果路径或方法发生变化，检查是否与其他API冲突
//  3. 如果路径或方法变化，更新Casbin权限规则
//  4. 保存更新后的API信息
//
// 注意事项：
//   - 更新路径或方法时，会检查是否与其他API的路径+方法组合冲突
//   - 如果路径或方法变化，会同步更新Casbin权限规则，确保权限数据一致性
//   - 如果API不存在，会返回gorm.ErrRecordNotFound错误
//   - 如果新路径+方法组合与其他API冲突，返回"存在相同api路径"错误
//
// @author: [piexlmax](https://github.com/piexlmax)
// @function: UpdateApi
// @description: 根据id更新api
// @param: api model.SysApi
// @return: err error
func (apiService *ApiService) UpdateApi(api system.SysApi) (err error) {
	var oldA system.SysApi
	// 步骤1: 查询原API记录
	err = global.GVA_DB.First(&oldA, "id = ?", api.ID).Error

	// 步骤2: 如果路径或方法发生变化，检查是否与其他API冲突
	if oldA.Path != api.Path || oldA.Method != api.Method {
		var duplicateApi system.SysApi
		// 查询是否存在相同的路径+方法组合
		if ferr := global.GVA_DB.First(&duplicateApi, "path = ? AND method = ?", api.Path, api.Method).Error; ferr != nil {
			// 如果查询出错，且不是"记录不存在"错误，返回错误
			if !errors.Is(ferr, gorm.ErrRecordNotFound) {
				return ferr
			}
			// 如果是"记录不存在"错误，说明没有冲突，可以继续
		} else {
			// 如果找到了相同路径+方法的记录，检查是否是同一条记录
			if duplicateApi.ID != api.ID {
				// 是其他API的记录，存在冲突
				return errors.New("存在相同api路径")
			}
			// 是同一条记录，可以继续
		}
	}

	// 如果查询原记录失败，返回错误
	if err != nil {
		return err
	}

	// 步骤3: 如果路径或方法变化，更新Casbin权限规则
	err = CasbinServiceApp.UpdateCasbinApi(oldA.Path, api.Path, oldA.Method, api.Method)
	if err != nil {
		return err
	}

	// 步骤4: 保存更新后的API信息
	return global.GVA_DB.Save(&api).Error
}

// DeleteApisByIds 批量删除API（事务方式）
//
// 功能说明：
// 根据ID数组批量删除API记录，同时清理相关的Casbin权限规则。
// 使用事务确保数据一致性，要么全部成功，要么全部回滚。
//
// 参数说明：
//   - ids: ID数组请求对象，包含要删除的API ID列表
//
// 返回值：
//   - err: 删除失败时返回错误（事务会回滚）
//
// 使用示例：
//
//	ids := request.IdsReq{Ids: []int{1, 2, 3, 4, 5}}
//	err := ApiServiceApp.DeleteApisByIds(ids)
//	if err != nil {
//	    log.Printf("批量删除API失败: %v", err)
//	}
//
// 实现逻辑：
//  1. 开启数据库事务
//  2. 根据ID数组查询要删除的API记录（用于后续清理Casbin）
//  3. 批量删除API记录
//  4. 遍历删除的API，清理Casbin权限规则
//  5. 如果任何步骤失败，事务回滚
//
// 使用场景：
//   - 前端批量选择删除API
//   - 管理员批量清理不需要的API
//   - 系统维护时批量删除废弃的API
//
// 注意事项：
//   - 使用事务确保数据一致性
//   - 删除API时会同时清理Casbin权限规则，避免产生孤立数据
//   - 如果某个ID不存在，不会报错，只是跳过该记录
//   - ClearCasbin在事务外执行，因为Casbin可能使用不同的存储
//
// @author: [piexlmax](https://github.com/piexlmax)
// @function: DeleteApisByIds
// @description: 删除选中API
// @param: apis []model.SysApi
// @return: err error
func (apiService *ApiService) DeleteApisByIds(ids request.IdsReq) (err error) {
	// 使用事务确保数据一致性
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var apis []system.SysApi

		// 步骤1: 根据ID数组查询要删除的API记录（用于后续清理Casbin）
		err = tx.Find(&apis, "id in ?", ids.Ids).Error
		if err != nil {
			return err
		}

		// 步骤2: 批量删除API记录
		err = tx.Delete(&[]system.SysApi{}, "id in ?", ids.Ids).Error
		if err != nil {
			return err
		}

		// 步骤3: 遍历删除的API，清理Casbin权限规则
		// 注意：ClearCasbin在事务外执行，因为Casbin可能使用不同的存储（如内存或Redis）
		for _, sysApi := range apis {
			CasbinServiceApp.ClearCasbin(1, sysApi.Path, sysApi.Method)
		}
		return err
	})
}
