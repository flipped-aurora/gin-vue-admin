package system

import (
	"errors"
	"strconv"

	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"gorm.io/gorm"
)

/*
 * 角色权限管理服务
 *
 * 本文件实现了基于RBAC（基于角色的访问控制）的角色权限管理系统
 * 主要功能包括：
 * 1. 角色的创建、复制、更新、删除
 * 2. 角色与菜单的关联管理
 * 3. 角色与数据权限的关联管理
 * 4. 角色树形结构的查询和管理
 * 5. Casbin权限策略的同步管理
 *
 * 使用示例：
 *   // 创建角色
 *   auth := system.SysAuthority{
 *       AuthorityId:   888,
 *       AuthorityName: "管理员",
 *       ParentId:      &parentId,
 *   }
 *   result, err := AuthorityServiceApp.CreateAuthority(auth)
 *
 *   // 获取角色列表
 *   list, err := AuthorityServiceApp.GetAuthorityInfoList(888)
 *
 *   // 设置菜单权限
 *   auth.SysBaseMenus = menus
 *   err := AuthorityServiceApp.SetMenuAuthority(&auth)
 */

// ErrRoleExistence 角色已存在错误
var ErrRoleExistence = errors.New("存在相同角色id")

// AuthorityService 角色权限服务结构体
// 提供角色权限相关的所有业务逻辑方法
type AuthorityService struct{}

// AuthorityServiceApp 角色权限服务实例
// 全局单例，用于在项目各处调用角色权限服务方法
var AuthorityServiceApp = new(AuthorityService)

// CreateAuthority 创建角色
//
// 功能说明：
//
//	创建一个新的角色，并为其分配默认菜单权限和API权限
//	使用数据库事务确保数据一致性，如果任何步骤失败，整个操作将回滚
//
// 参数：
//
//	auth: 要创建的角色对象，必须包含 AuthorityId（角色ID）和 AuthorityName（角色名称）
//
// 返回值：
//
//	authority: 创建成功的角色对象
//	err: 错误信息，如果角色ID已存在则返回 ErrRoleExistence
//
// 执行流程：
//  1. 检查角色ID是否已存在，如果存在则返回错误
//  2. 开启数据库事务
//  3. 创建角色记录
//  4. 为角色分配默认菜单（通过 many2many 关联表）
//  5. 为角色添加默认的 Casbin 权限策略（API访问权限）
//  6. 提交事务或回滚（如果出错）
//
// 使用示例：
//
//	auth := system.SysAuthority{
//	    AuthorityId:   999,
//	    AuthorityName: "普通用户",
//	    ParentId:      &parentId, // 可选，父角色ID
//	}
//	result, err := AuthorityServiceApp.CreateAuthority(auth)
//	if err != nil {
//	    log.Printf("创建角色失败: %v", err)
//	    return
//	}
//	log.Printf("角色创建成功: %+v", result)
func (authorityService *AuthorityService) CreateAuthority(auth system.SysAuthority) (authority system.SysAuthority, err error) {
	// 检查角色ID是否已存在
	// 如果查询到记录（err不是RecordNotFound），说明角色已存在
	if err = global.GVA_DB.Where("authority_id = ?", auth.AuthorityId).First(&system.SysAuthority{}).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		return auth, ErrRoleExistence
	}

	// 使用事务确保数据一致性
	e := global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 步骤1: 创建角色记录
		if err = tx.Create(&auth).Error; err != nil {
			return err
		}

		// 步骤2: 为角色分配默认菜单权限
		// DefaultMenu() 返回系统默认菜单列表
		auth.SysBaseMenus = systemReq.DefaultMenu()
		// 使用 GORM 的 Association 功能更新 many2many 关联表
		if err = tx.Model(&auth).Association("SysBaseMenus").Replace(&auth.SysBaseMenus); err != nil {
			return err
		}

		// 步骤3: 为角色添加默认的 Casbin API权限策略
		// Casbin 是访问控制库，用于管理API级别的权限
		casbinInfos := systemReq.DefaultCasbin()
		authorityId := strconv.Itoa(int(auth.AuthorityId))
		rules := [][]string{}
		// 构建 Casbin 规则: [角色ID, API路径, HTTP方法]
		for _, v := range casbinInfos {
			rules = append(rules, []string{authorityId, v.Path, v.Method})
		}
		// 批量添加权限策略到数据库
		return CasbinServiceApp.AddPolicies(tx, rules)
	})

	return auth, e
}

// CopyAuthority 复制角色
//
// 功能说明：
//
//	基于现有角色创建一个新角色，复制原角色的所有权限配置
//	包括：菜单权限、按钮权限、API权限（Casbin策略）
//	如果复制过程中出现错误，会自动清理已创建的角色
//
// 参数：
//
//	adminAuthorityID: 执行操作的管理员角色ID，用于权限验证
//	copyInfo: 复制信息，包含新角色对象和源角色ID
//	  - Authority: 新角色对象（必须包含新的 AuthorityId）
//	  - OldAuthorityId: 要复制的源角色ID
//
// 返回值：
//
//	authority: 复制成功的新角色对象
//	err: 错误信息
//
// 执行流程：
//  1. 检查新角色ID是否已存在
//  2. 获取源角色的菜单权限
//  3. 创建新角色并关联菜单权限
//  4. 复制源角色的按钮权限
//  5. 复制源角色的API权限（Casbin策略）
//  6. 如果任何步骤失败，删除已创建的角色
//
// 使用示例：
//
//	copyInfo := response.SysAuthorityCopyResponse{
//	    Authority: system.SysAuthority{
//	        AuthorityId:   1001,
//	        AuthorityName: "新角色（复制自888）",
//	    },
//	    OldAuthorityId: 888, // 源角色ID
//	}
//	adminID := uint(999) // 管理员角色ID
//	newAuth, err := AuthorityServiceApp.CopyAuthority(adminID, copyInfo)
//	if err != nil {
//	    log.Printf("复制角色失败: %v", err)
//	    return
//	}
//	log.Printf("角色复制成功: %+v", newAuth)
func (authorityService *AuthorityService) CopyAuthority(adminAuthorityID uint, copyInfo response.SysAuthorityCopyResponse) (authority system.SysAuthority, err error) {
	// 步骤1: 检查新角色ID是否已存在
	var authorityBox system.SysAuthority
	if !errors.Is(global.GVA_DB.Where("authority_id = ?", copyInfo.Authority.AuthorityId).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return authority, ErrRoleExistence
	}

	// 清空子角色列表，避免复制时携带子角色
	copyInfo.Authority.Children = []system.SysAuthority{}

	// 步骤2: 获取源角色的菜单权限
	menus, err := MenuServiceApp.GetMenuAuthority(&request.GetAuthorityId{AuthorityId: copyInfo.OldAuthorityId})
	if err != nil {
		return
	}

	// 转换菜单格式，提取基础菜单信息
	var baseMenu []system.SysBaseMenu
	for _, v := range menus {
		intNum := v.MenuId
		v.SysBaseMenu.ID = uint(intNum)
		baseMenu = append(baseMenu, v.SysBaseMenu)
	}
	copyInfo.Authority.SysBaseMenus = baseMenu

	// 步骤3: 创建新角色（会自动关联菜单权限）
	err = global.GVA_DB.Create(&copyInfo.Authority).Error
	if err != nil {
		return
	}

	// 步骤4: 复制源角色的按钮权限
	var btns []system.SysAuthorityBtn
	err = global.GVA_DB.Find(&btns, "authority_id = ?", copyInfo.OldAuthorityId).Error
	if err != nil {
		return
	}
	if len(btns) > 0 {
		// 更新按钮权限的角色ID为新角色ID
		for i := range btns {
			btns[i].AuthorityId = copyInfo.Authority.AuthorityId
		}
		// 批量创建按钮权限记录
		err = global.GVA_DB.Create(&btns).Error
		if err != nil {
			return
		}
	}

	// 步骤5: 复制源角色的API权限（Casbin策略）
	paths := CasbinServiceApp.GetPolicyPathByAuthorityId(copyInfo.OldAuthorityId)
	err = CasbinServiceApp.UpdateCasbin(adminAuthorityID, copyInfo.Authority.AuthorityId, paths)
	if err != nil {
		// 如果API权限复制失败，删除已创建的角色（回滚操作）
		_ = authorityService.DeleteAuthority(&copyInfo.Authority)
	}
	return copyInfo.Authority, err
}

// UpdateAuthority 更新角色信息
//
// 功能说明：
//
//	更新角色的基本信息（如角色名称、父角色ID等）
//	注意：此方法只更新角色的基本信息，不更新菜单权限和API权限
//	如需更新权限，请使用 SetMenuAuthority 和 SetDataAuthority 方法
//
// 参数：
//
//	auth: 包含要更新字段的角色对象，必须包含 AuthorityId（角色ID）
//	      只更新非零值字段，零值字段不会被更新
//
// 返回值：
//
//	authority: 更新后的角色对象
//	err: 错误信息，如果角色不存在则返回错误
//
// 使用示例：
//
//	auth := system.SysAuthority{
//	    AuthorityId:   888,
//	    AuthorityName: "更新后的角色名",
//	    ParentId:      &newParentId,
//	}
//	result, err := AuthorityServiceApp.UpdateAuthority(auth)
//	if err != nil {
//	    log.Printf("更新角色失败: %v", err)
//	    return
//	}
//	log.Printf("角色更新成功: %+v", result)
func (authorityService *AuthorityService) UpdateAuthority(auth system.SysAuthority) (authority system.SysAuthority, err error) {
	// 步骤1: 查询要更新的角色是否存在
	var oldAuthority system.SysAuthority
	err = global.GVA_DB.Where("authority_id = ?", auth.AuthorityId).First(&oldAuthority).Error
	if err != nil {
		global.GVA_LOG.Debug(err.Error())
		return system.SysAuthority{}, errors.New("查询角色数据失败")
	}

	// 步骤2: 使用 Updates 方法更新角色信息
	// Updates 方法只会更新非零值字段，零值字段不会被更新
	err = global.GVA_DB.Model(&oldAuthority).Updates(&auth).Error
	return auth, err
}

// DeleteAuthority 删除角色
//
// 功能说明：
//
//	删除指定的角色及其所有关联数据
//	删除前会进行安全检查，确保角色可以被安全删除
//	使用数据库事务确保删除操作的原子性
//
// 参数：
//
//	auth: 要删除的角色对象指针，必须包含 AuthorityId（角色ID）
//
// 返回值：
//
//	err: 错误信息
//
// 安全检查：
//  1. 检查角色是否存在
//  2. 检查是否有用户正在使用此角色（通过 Users 关联）
//  3. 检查是否有用户直接关联此角色（通过 SysUser 表）
//  4. 检查是否存在子角色（不允许删除有子角色的角色）
//
// 删除内容：
//  1. 角色主记录（软删除）
//  2. 角色与菜单的关联关系
//  3. 角色的数据权限关联关系
//  4. 用户与角色的关联关系（sys_user_authority 表）
//  5. 角色的按钮权限（SysAuthorityBtn）
//  6. 角色的 Casbin API权限策略
//
// 使用示例：
//
//	auth := &system.SysAuthority{AuthorityId: 888}
//	err := AuthorityServiceApp.DeleteAuthority(auth)
//	if err != nil {
//	    log.Printf("删除角色失败: %v", err)
//	    return
//	}
//	log.Println("角色删除成功")
func (authorityService *AuthorityService) DeleteAuthority(auth *system.SysAuthority) error {
	// 安全检查1: 检查角色是否存在，并预加载用户关联
	if errors.Is(global.GVA_DB.Debug().Preload("Users").First(&auth).Error, gorm.ErrRecordNotFound) {
		return errors.New("该角色不存在")
	}

	// 安全检查2: 检查是否有用户通过 many2many 关联使用此角色
	if len(auth.Users) != 0 {
		return errors.New("此角色有用户正在使用禁止删除")
	}

	// 安全检查3: 检查是否有用户直接关联此角色
	if !errors.Is(global.GVA_DB.Where("authority_id = ?", auth.AuthorityId).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用禁止删除")
	}

	// 安全检查4: 检查是否存在子角色（不允许删除有子角色的角色）
	if !errors.Is(global.GVA_DB.Where("parent_id = ?", auth.AuthorityId).First(&system.SysAuthority{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色存在子角色不允许删除")
	}

	// 使用事务确保删除操作的原子性
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var err error

		// 步骤1: 预加载关联数据并软删除角色主记录
		// Unscoped().Delete() 执行软删除（设置 DeletedAt 字段）
		if err = tx.Preload("SysBaseMenus").Preload("DataAuthorityId").Where("authority_id = ?", auth.AuthorityId).First(auth).Unscoped().Delete(auth).Error; err != nil {
			return err
		}

		// 步骤2: 删除角色与菜单的关联关系（many2many 中间表）
		if len(auth.SysBaseMenus) > 0 {
			if err = tx.Model(auth).Association("SysBaseMenus").Delete(auth.SysBaseMenus); err != nil {
				return err
			}
		}

		// 步骤3: 删除角色的数据权限关联关系（many2many 中间表）
		if len(auth.DataAuthorityId) > 0 {
			if err = tx.Model(auth).Association("DataAuthorityId").Delete(auth.DataAuthorityId); err != nil {
				return err
			}
		}

		// 步骤4: 删除用户与角色的关联关系（sys_user_authority 表）
		if err = tx.Delete(&system.SysUserAuthority{}, "sys_authority_authority_id = ?", auth.AuthorityId).Error; err != nil {
			return err
		}

		// 步骤5: 删除角色的按钮权限
		if err = tx.Where("authority_id = ?", auth.AuthorityId).Delete(&[]system.SysAuthorityBtn{}).Error; err != nil {
			return err
		}

		// 步骤6: 删除角色的 Casbin API权限策略
		authorityId := strconv.Itoa(int(auth.AuthorityId))
		if err = CasbinServiceApp.RemoveFilteredPolicy(tx, authorityId); err != nil {
			return err
		}

		return nil
	})
}

// GetAuthorityInfoList 获取角色信息列表（树形结构）
//
// 功能说明：
//
//	根据当前角色ID获取其可管理的角色列表，返回树形结构
//	支持严格权限模式和普通模式两种查询方式
//
// 参数：
//
//	authorityID: 当前操作的角色ID，用于确定可查看的角色范围
//
// 返回值：
//
//	list: 角色列表（树形结构），每个角色包含其子角色
//	err: 错误信息
//
// 权限模式说明：
//  1. 严格权限模式（UseStrictAuth = true）：
//     - 顶级角色（ParentId = 0）：只能查看和管理自己
//     - 非顶级角色：只能查看和管理自己的子角色
//  2. 普通模式（UseStrictAuth = false）：
//     - 返回所有顶级角色（ParentId = 0）及其子角色树
//
// 使用示例：
//
//	// 获取角色ID为888的角色可管理的角色列表
//	list, err := AuthorityServiceApp.GetAuthorityInfoList(888)
//	if err != nil {
//	    log.Printf("获取角色列表失败: %v", err)
//	    return
//	}
//	// 遍历角色树
//	for _, auth := range list {
//	    log.Printf("角色: %s (ID: %d)", auth.AuthorityName, auth.AuthorityId)
//	    // 递归处理子角色...
//	}
func (authorityService *AuthorityService) GetAuthorityInfoList(authorityID uint) (list []system.SysAuthority, err error) {
	// 步骤1: 查询当前角色信息
	var authority system.SysAuthority
	err = global.GVA_DB.Where("authority_id = ?", authorityID).First(&authority).Error
	if err != nil {
		return nil, err
	}

	var authorities []system.SysAuthority
	db := global.GVA_DB.Model(&system.SysAuthority{})

	// 步骤2: 根据权限模式查询角色列表
	if global.GVA_CONFIG.System.UseStrictAuth {
		// 严格权限模式：根据角色层级限制查询范围
		if *authority.ParentId == 0 {
			// 顶级角色：只能查看和管理自己
			err = db.Preload("DataAuthorityId").Where("authority_id = ?", authorityID).Find(&authorities).Error
		} else {
			// 非顶级角色：只能查看和管理自己的子角色
			err = db.Debug().Preload("DataAuthorityId").Where("parent_id = ?", authorityID).Find(&authorities).Error
		}
	} else {
		// 普通模式：返回所有顶级角色
		err = db.Preload("DataAuthorityId").Where("parent_id = ?", "0").Find(&authorities).Error
	}

	// 步骤3: 递归加载每个角色的子角色树
	for k := range authorities {
		err = authorityService.findChildrenAuthority(&authorities[k])
	}
	return authorities, err
}

// GetStructAuthorityList 获取角色结构列表（扁平化的角色ID列表）
//
// 功能说明：
//
//	递归获取指定角色及其所有子角色的ID列表，返回扁平化的角色ID数组
//	用于权限验证和批量操作
//
// 参数：
//
//	authorityID: 起始角色ID，从此角色开始向下递归查找所有子角色
//
// 返回值：
//
//	list: 角色ID列表，包含起始角色及其所有子角色的ID
//	err: 错误信息
//
// 递归逻辑：
//  1. 查询当前角色的直接子角色
//  2. 对每个子角色递归调用此方法
//  3. 如果是顶级角色（ParentId = 0），将自身ID也加入列表
//
// 使用示例：
//
//	// 获取角色888及其所有子角色的ID列表
//	roleIDs, err := AuthorityServiceApp.GetStructAuthorityList(888)
//	if err != nil {
//	    log.Printf("获取角色结构失败: %v", err)
//	    return
//	}
//	// roleIDs = [888, 889, 890, 891, ...]  // 包含所有子角色ID
//	log.Printf("角色结构: %v", roleIDs)
func (authorityService *AuthorityService) GetStructAuthorityList(authorityID uint) (list []uint, err error) {
	// 查询当前角色信息
	var auth system.SysAuthority
	_ = global.GVA_DB.First(&auth, "authority_id = ?", authorityID).Error

	// 查询当前角色的直接子角色
	var authorities []system.SysAuthority
	err = global.GVA_DB.Preload("DataAuthorityId").Where("parent_id = ?", authorityID).Find(&authorities).Error

	// 递归处理每个子角色
	if len(authorities) > 0 {
		for k := range authorities {
			// 将子角色ID加入列表
			list = append(list, authorities[k].AuthorityId)
			// 递归获取子角色的子角色
			childrenList, err := authorityService.GetStructAuthorityList(authorities[k].AuthorityId)
			if err == nil {
				list = append(list, childrenList...)
			}
		}
	}

	// 如果是顶级角色，将自身ID也加入列表
	if *auth.ParentId == 0 {
		list = append(list, authorityID)
	}
	return list, err
}

// CheckAuthorityIDAuth 检查角色权限
//
// 功能说明：
//
//	验证当前角色是否有权限操作目标角色
//	在严格权限模式下，只能操作自己及其子角色
//
// 参数：
//
//	authorityID: 当前操作的角色ID
//	targetID: 目标角色ID（要操作的角色）
//
// 返回值：
//
//	err: 错误信息，如果没有权限则返回错误
//
// 权限规则：
//  1. 如果未开启严格权限模式，直接返回成功（不检查）
//  2. 如果开启严格权限模式：
//     - 只能操作自己及其所有子角色
//     - 不能操作父角色或兄弟角色
//
// 使用示例：
//
//	// 检查角色888是否有权限操作角色889
//	err := AuthorityServiceApp.CheckAuthorityIDAuth(888, 889)
//	if err != nil {
//	    log.Printf("权限检查失败: %v", err)
//	    return
//	}
//	log.Println("权限检查通过")
func (authorityService *AuthorityService) CheckAuthorityIDAuth(authorityID, targetID uint) (err error) {
	// 如果未开启严格权限模式，直接通过检查
	if !global.GVA_CONFIG.System.UseStrictAuth {
		return nil
	}

	// 获取当前角色及其所有子角色的ID列表
	authIDS, err := authorityService.GetStructAuthorityList(authorityID)
	if err != nil {
		return err
	}

	// 检查目标角色ID是否在可操作的角色列表中
	hasAuth := false
	for _, v := range authIDS {
		if v == targetID {
			hasAuth = true
			break
		}
	}

	// 如果没有权限，返回错误
	if !hasAuth {
		return errors.New("您提交的角色ID不合法")
	}
	return nil
}

// GetAuthorityInfo 获取单个角色的详细信息
//
// 功能说明：
//
//	根据角色ID查询角色的完整信息，包括数据权限关联
//
// 参数：
//
//	auth: 角色对象，必须包含 AuthorityId（角色ID）
//
// 返回值：
//
//	sa: 查询到的角色对象，包含数据权限关联信息
//	err: 错误信息，如果角色不存在则返回错误
//
// 使用示例：
//
//	auth := system.SysAuthority{AuthorityId: 888}
//	result, err := AuthorityServiceApp.GetAuthorityInfo(auth)
//	if err != nil {
//	    log.Printf("获取角色信息失败: %v", err)
//	    return
//	}
//	log.Printf("角色信息: %+v", result)
//	log.Printf("数据权限: %+v", result.DataAuthorityId)
func (authorityService *AuthorityService) GetAuthorityInfo(auth system.SysAuthority) (sa system.SysAuthority, err error) {
	// 查询角色信息，并预加载数据权限关联
	err = global.GVA_DB.Preload("DataAuthorityId").Where("authority_id = ?", auth.AuthorityId).First(&sa).Error
	return sa, err
}

// SetDataAuthority 设置角色的数据权限
//
// 功能说明：
//
//	为指定角色设置数据权限范围，数据权限决定了角色可以访问哪些数据
//	数据权限通过关联其他角色来实现，表示可以访问这些角色管辖的数据
//
// 参数：
//
//	adminAuthorityID: 执行操作的管理员角色ID，用于权限验证
//	auth: 角色对象，必须包含：
//	  - AuthorityId: 要设置权限的角色ID
//	  - DataAuthorityId: 数据权限角色列表（可以访问这些角色的数据）
//
// 返回值：
//
//	error: 错误信息
//
// 权限验证：
//  1. 验证管理员是否有权限操作目标角色
//  2. 验证管理员是否有权限操作所有数据权限角色
//
// 使用示例：
//
//	auth := system.SysAuthority{
//	    AuthorityId: 888,
//	    DataAuthorityId: []*system.SysAuthority{
//	        {AuthorityId: 889}, // 可以访问角色889的数据
//	        {AuthorityId: 890}, // 可以访问角色890的数据
//	    },
//	}
//	adminID := uint(999)
//	err := AuthorityServiceApp.SetDataAuthority(adminID, auth)
//	if err != nil {
//	    log.Printf("设置数据权限失败: %v", err)
//	    return
//	}
//	log.Println("数据权限设置成功")
func (authorityService *AuthorityService) SetDataAuthority(adminAuthorityID uint, auth system.SysAuthority) error {
	// 步骤1: 收集所有需要权限验证的角色ID
	var checkIDs []uint
	checkIDs = append(checkIDs, auth.AuthorityId) // 目标角色ID
	// 添加所有数据权限角色ID
	for i := range auth.DataAuthorityId {
		checkIDs = append(checkIDs, auth.DataAuthorityId[i].AuthorityId)
	}

	// 步骤2: 验证管理员是否有权限操作所有相关角色
	for i := range checkIDs {
		err := authorityService.CheckAuthorityIDAuth(adminAuthorityID, checkIDs[i])
		if err != nil {
			return err
		}
	}

	// 步骤3: 查询目标角色并更新数据权限关联
	var s system.SysAuthority
	global.GVA_DB.Preload("DataAuthorityId").First(&s, "authority_id = ?", auth.AuthorityId)
	// 使用 Replace 方法替换数据权限关联（many2many 中间表）
	err := global.GVA_DB.Model(&s).Association("DataAuthorityId").Replace(&auth.DataAuthorityId)
	return err
}

// SetMenuAuthority 设置角色的菜单权限
//
// 功能说明：
//
//	为指定角色设置可访问的菜单列表
//	菜单权限决定了角色在系统中可以看到哪些菜单项
//
// 参数：
//
//	auth: 角色对象指针，必须包含：
//	  - AuthorityId: 要设置权限的角色ID
//	  - SysBaseMenus: 菜单列表（角色可以访问的菜单）
//
// 返回值：
//
//	error: 错误信息
//
// 注意事项：
//   - 使用 Replace 方法会完全替换原有的菜单关联
//   - 如果传入空菜单列表，会清空角色的所有菜单权限
//   - 菜单权限通过 many2many 关联表（sys_authority_menus）存储
//
// 使用示例：
//
//	auth := &system.SysAuthority{
//	    AuthorityId: 888,
//	    SysBaseMenus: []system.SysBaseMenu{
//	        {ID: 1}, // 菜单ID为1
//	        {ID: 2}, // 菜单ID为2
//	        {ID: 3}, // 菜单ID为3
//	    },
//	}
//	err := AuthorityServiceApp.SetMenuAuthority(auth)
//	if err != nil {
//	    log.Printf("设置菜单权限失败: %v", err)
//	    return
//	}
//	log.Println("菜单权限设置成功")
func (authorityService *AuthorityService) SetMenuAuthority(auth *system.SysAuthority) error {
	// 查询目标角色并预加载现有菜单权限
	var s system.SysAuthority
	global.GVA_DB.Preload("SysBaseMenus").First(&s, "authority_id = ?", auth.AuthorityId)
	// 使用 Replace 方法替换菜单关联（many2many 中间表）
	err := global.GVA_DB.Model(&s).Association("SysBaseMenus").Replace(&auth.SysBaseMenus)
	return err
}

// findChildrenAuthority 递归查询子角色（私有方法）
//
// 功能说明：
//
//	递归查询指定角色的所有子角色，构建完整的角色树结构
//	这是一个辅助方法，用于构建角色的树形结构
//
// 参数：
//
//	authority: 角色对象指针，查询结果会填充到其 Children 字段中
//
// 返回值：
//
//	err: 错误信息
//
// 递归逻辑：
//  1. 查询当前角色的直接子角色（parent_id = authority.AuthorityId）
//  2. 预加载子角色的数据权限关联
//  3. 对每个子角色递归调用此方法，继续查询子角色的子角色
//  4. 直到没有子角色为止
//
// 使用示例：
//
//	// 通常在 GetAuthorityInfoList 方法中调用
//	auth := &system.SysAuthority{AuthorityId: 888}
//	err := authorityService.findChildrenAuthority(auth)
//	if err != nil {
//	    log.Printf("查询子角色失败: %v", err)
//	    return
//	}
//	// auth.Children 现在包含所有子角色及其子角色
func (authorityService *AuthorityService) findChildrenAuthority(authority *system.SysAuthority) (err error) {
	// 查询当前角色的直接子角色，并预加载数据权限关联
	err = global.GVA_DB.Preload("DataAuthorityId").Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
	if err != nil {
		return err
	}

	// 如果有子角色，递归查询每个子角色的子角色
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = authorityService.findChildrenAuthority(&authority.Children[k])
			if err != nil {
				return err
			}
		}
	}
	return err
}

// GetParentAuthorityID 获取角色的父角色ID
//
// 功能说明：
//
//	根据角色ID查询其父角色ID
//	用于角色层级关系的查询和验证
//
// 参数：
//
//	authorityID: 角色ID
//
// 返回值：
//
//	parentID: 父角色ID，如果是顶级角色则返回0
//	err: 错误信息，如果角色不存在则返回错误
//
// 使用示例：
//
//	parentID, err := AuthorityServiceApp.GetParentAuthorityID(888)
//	if err != nil {
//	    log.Printf("获取父角色ID失败: %v", err)
//	    return
//	}
//	if parentID == 0 {
//	    log.Println("这是顶级角色")
//	} else {
//	    log.Printf("父角色ID: %d", parentID)
//	}
func (authorityService *AuthorityService) GetParentAuthorityID(authorityID uint) (parentID uint, err error) {
	// 查询角色信息
	var authority system.SysAuthority
	err = global.GVA_DB.Where("authority_id = ?", authorityID).First(&authority).Error
	if err != nil {
		return
	}
	// 返回父角色ID（解引用指针）
	return *authority.ParentId, nil
}
