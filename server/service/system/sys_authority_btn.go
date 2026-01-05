/*
 * @Description: 权限按钮服务层 - 管理角色与菜单按钮的关联关系
 * @Author: gin-vue-admin
 * @Date: 2024-01-01
 * @功能说明:
 *   1. GetAuthorityBtn - 查询指定角色在指定菜单下已分配的按钮权限
 *   2. SetAuthorityBtn - 设置指定角色在指定菜单下的按钮权限（先删除旧权限，再添加新权限）
 *   3. CanRemoveAuthorityBtn - 检查指定按钮是否可以被删除（是否正在被某个角色使用）
 *
 * @数据模型说明:
 *   SysAuthorityBtn: 权限按钮关联表，存储角色ID、菜单ID和按钮ID的关联关系
 *   - AuthorityId: 角色ID（外键关联到角色表）
 *   - SysMenuID: 菜单ID（外键关联到菜单表）
 *   - SysBaseMenuBtnID: 菜单按钮ID（外键关联到按钮表）
 *
 * @使用场景:
 *   1. 前端页面需要根据角色动态显示/隐藏按钮
 *   2. 管理员为角色分配菜单按钮权限
 *   3. 删除按钮前检查是否被使用
 *
 * @Demo示例:
 *   // 示例1: 获取角色ID为1，菜单ID为10的已分配按钮权限
 *   req := request.SysAuthorityBtnReq{
 *       AuthorityId: 1,
 *       MenuID:      10,
 *   }
 *   res, err := AuthorityBtnServiceApp.GetAuthorityBtn(req)
 *   // res.Selected 返回 [1, 2, 3] 表示该角色在该菜单下拥有按钮1、2、3的权限
 *
 *   // 示例2: 为角色ID为1，菜单ID为10分配按钮权限（按钮ID为1、2、3）
 *   req := request.SysAuthorityBtnReq{
 *       AuthorityId: 1,
 *       MenuID:      10,
 *       Selected:    []uint{1, 2, 3},
 *   }
 *   err := AuthorityBtnServiceApp.SetAuthorityBtn(req)
 *   // 执行后，该角色在该菜单下只拥有按钮1、2、3的权限（之前的权限会被清除）
 *
 *   // 示例3: 检查按钮ID为5是否可以删除
 *   err := AuthorityBtnServiceApp.CanRemoveAuthorityBtn("5")
 *   // 如果err为nil，表示可以删除；如果err不为nil，表示按钮正在被使用，无法删除
 */
package system

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"gorm.io/gorm"
)

// AuthorityBtnService 权限按钮服务结构体
// 提供权限按钮相关的业务逻辑处理
type AuthorityBtnService struct{}

// AuthorityBtnServiceApp 权限按钮服务实例
// 全局单例，供其他模块调用
var AuthorityBtnServiceApp = new(AuthorityBtnService)

// GetAuthorityBtn 获取指定角色在指定菜单下已分配的按钮权限
//
// @功能说明:
//
//	根据角色ID和菜单ID，查询该角色在该菜单下已分配的所有按钮权限
//	返回已选中的按钮ID列表
//
// @参数说明:
//
//	req: 请求参数
//	  - AuthorityId: 角色ID（必填）
//	  - MenuID: 菜单ID（必填）
//	  - Selected: 此方法中不使用此字段
//
// @返回值说明:
//
//	res: 响应结果
//	  - Selected: 已选中的按钮ID数组，例如 [1, 2, 3] 表示拥有按钮1、2、3的权限
//	err: 错误信息，如果查询失败则返回错误
//
// @使用示例:
//
//	req := request.SysAuthorityBtnReq{
//	    AuthorityId: 1,  // 角色ID
//	    MenuID:      10, // 菜单ID
//	}
//	res, err := AuthorityBtnServiceApp.GetAuthorityBtn(req)
//	if err != nil {
//	    // 处理错误
//	}
//	// res.Selected 包含该角色在该菜单下的所有按钮权限ID
//
// @数据库操作:
//
//	SELECT * FROM sys_authority_btns
//	WHERE authority_id = ? AND sys_menu_id = ?
func (a *AuthorityBtnService) GetAuthorityBtn(req request.SysAuthorityBtnReq) (res response.SysAuthorityBtnRes, err error) {
	// 定义权限按钮数组，用于存储查询结果
	var authorityBtn []system.SysAuthorityBtn

	// 根据角色ID和菜单ID查询权限按钮关联记录
	// 查询条件: authority_id = req.AuthorityId AND sys_menu_id = req.MenuID
	err = global.GVA_DB.Find(&authorityBtn, "authority_id = ? and sys_menu_id = ?", req.AuthorityId, req.MenuID).Error
	if err != nil {
		// 如果查询出错，直接返回错误
		return
	}

	// 提取所有已选中的按钮ID
	var selected []uint
	for _, v := range authorityBtn {
		// 将每条记录的按钮ID添加到selected数组中
		selected = append(selected, v.SysBaseMenuBtnID)
	}

	// 将选中的按钮ID数组赋值给响应结果
	res.Selected = selected
	return res, err
}

// SetAuthorityBtn 设置指定角色在指定菜单下的按钮权限
//
// @功能说明:
//
//	为指定角色在指定菜单下分配按钮权限
//	采用"先删除后添加"的策略，确保权限设置的原子性
//	使用数据库事务保证数据一致性
//
// @参数说明:
//
//	req: 请求参数
//	  - AuthorityId: 角色ID（必填）
//	  - MenuID: 菜单ID（必填）
//	  - Selected: 要分配的按钮ID数组（必填），例如 [1, 2, 3] 表示分配按钮1、2、3的权限
//
// @返回值说明:
//
//	err: 错误信息，如果设置失败则返回错误
//
// @使用示例:
//
//	req := request.SysAuthorityBtnReq{
//	    AuthorityId: 1,           // 角色ID
//	    MenuID:      10,          // 菜单ID
//	    Selected:    []uint{1, 2, 3}, // 要分配的按钮ID数组
//	}
//	err := AuthorityBtnServiceApp.SetAuthorityBtn(req)
//	if err != nil {
//	    // 处理错误
//	}
//	// 设置成功后，该角色在该菜单下只拥有按钮1、2、3的权限
//
// @数据库操作流程:
//  1. 开启事务
//  2. DELETE FROM sys_authority_btns WHERE authority_id = ? AND sys_menu_id = ?
//  3. INSERT INTO sys_authority_btns (authority_id, sys_menu_id, sys_base_menu_btn_id) VALUES ...
//  4. 提交事务（如果任何步骤失败则回滚）
//
// @注意事项:
//   - 此操作会先删除该角色在该菜单下的所有旧权限，再添加新权限
//   - 如果Selected为空数组，则只删除旧权限，不添加新权限
//   - 使用事务保证操作的原子性，要么全部成功，要么全部失败
func (a *AuthorityBtnService) SetAuthorityBtn(req request.SysAuthorityBtnReq) (err error) {
	// 使用数据库事务确保操作的原子性
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 定义权限按钮数组，用于批量创建
		var authorityBtn []system.SysAuthorityBtn

		// 第一步: 删除该角色在该菜单下的所有旧权限
		// 这样可以确保权限设置的准确性，避免权限残留
		err = tx.Delete(&[]system.SysAuthorityBtn{}, "authority_id = ? and sys_menu_id = ?", req.AuthorityId, req.MenuID).Error
		if err != nil {
			// 如果删除失败，返回错误，事务会自动回滚
			return err
		}

		// 第二步: 根据请求参数中的Selected数组，构建新的权限按钮关联记录
		for _, v := range req.Selected {
			authorityBtn = append(authorityBtn, system.SysAuthorityBtn{
				AuthorityId:      req.AuthorityId, // 角色ID
				SysMenuID:        req.MenuID,      // 菜单ID
				SysBaseMenuBtnID: v,               // 按钮ID
			})
		}

		// 第三步: 如果有新的权限需要添加，则批量创建
		if len(authorityBtn) > 0 {
			err = tx.Create(&authorityBtn).Error
		}

		// 检查是否有错误发生
		if err != nil {
			// 如果创建失败，返回错误，事务会自动回滚
			return err
		}

		// 所有操作成功，返回nil，事务会自动提交
		return err
	})
}

// CanRemoveAuthorityBtn 检查指定按钮是否可以被删除
//
// @功能说明:
//
//	在删除按钮前，检查该按钮是否正在被某个角色使用
//	如果按钮正在被使用，则不允许删除，避免数据不一致
//
// @参数说明:
//
//	ID: 按钮ID（字符串类型），例如 "5" 表示按钮ID为5
//
// @返回值说明:
//
//	err: 错误信息
//	  - nil: 表示按钮未被使用，可以删除
//	  - 非nil: 表示按钮正在被使用，无法删除，错误信息为"此按钮正在被使用无法删除"
//
// @使用示例:
//
//	// 示例1: 检查按钮ID为5是否可以删除
//	err := AuthorityBtnServiceApp.CanRemoveAuthorityBtn("5")
//	if err != nil {
//	    // 按钮正在被使用，无法删除
//	    log.Println(err.Error()) // 输出: "此按钮正在被使用无法删除"
//	} else {
//	    // 按钮未被使用，可以删除
//	    // 执行删除操作
//	}
//
// @数据库操作:
//
//	SELECT * FROM sys_authority_btns
//	WHERE sys_base_menu_btn_id = ?
//	LIMIT 1
//
// @业务逻辑:
//  1. 查询数据库中是否存在使用该按钮的权限记录
//  2. 如果查询结果为"记录不存在"（ErrRecordNotFound），说明按钮未被使用，可以删除
//  3. 如果查询到记录，说明按钮正在被使用，返回错误，不允许删除
//
// @注意事项:
//   - 此方法只检查按钮是否被使用，不执行实际的删除操作
//   - 删除按钮前必须调用此方法进行检查，避免破坏数据完整性
func (a *AuthorityBtnService) CanRemoveAuthorityBtn(ID string) (err error) {
	// 查询数据库中是否存在使用该按钮的权限记录
	// First方法查询第一条匹配的记录，如果不存在则返回ErrRecordNotFound错误
	fErr := global.GVA_DB.First(&system.SysAuthorityBtn{}, "sys_base_menu_btn_id = ?", ID).Error

	// 判断错误是否为"记录不存在"
	if errors.Is(fErr, gorm.ErrRecordNotFound) {
		// 记录不存在，说明按钮未被使用，可以删除
		return nil
	}

	// 如果查询到记录（或查询出错），说明按钮正在被使用，返回错误
	// 这样前端可以提示用户"此按钮正在被使用无法删除"
	return errors.New("此按钮正在被使用无法删除")
}
