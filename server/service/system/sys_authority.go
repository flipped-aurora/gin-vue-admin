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

var ErrRoleExistence = errors.New("存在相同角色id")

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateAuthority
//@description: 创建一个角色
//@param: auth model.SysAuthority
//@return: authority system.SysAuthority, err error

type AuthorityService struct{}

var AuthorityServiceApp = new(AuthorityService)

func (authorityService *AuthorityService) CreateAuthority(auth system.SysAuthority) (authority system.SysAuthority, err error) {

	if err = global.GVA_DB.Where("authority_id = ?", auth.AuthorityId).First(&system.SysAuthority{}).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		return auth, ErrRoleExistence
	}

	e := global.GVA_DB.Transaction(func(tx *gorm.DB) error {

		if err = tx.Create(&auth).Error; err != nil {
			return err
		}

		auth.SysBaseMenus = systemReq.DefaultMenu()
		if err = tx.Model(&auth).Association("SysBaseMenus").Replace(&auth.SysBaseMenus); err != nil {
			return err
		}
		casbinInfos := systemReq.DefaultCasbin()
		authorityId := strconv.Itoa(int(auth.AuthorityId))
		rules := [][]string{}
		for _, v := range casbinInfos {
			rules = append(rules, []string{authorityId, v.Path, v.Method})
		}
		return CasbinServiceApp.AddPolicies(tx, rules)
	})

	return auth, e
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CopyAuthority
//@description: 复制一个角色
//@param: copyInfo response.SysAuthorityCopyResponse
//@return: authority system.SysAuthority, err error

func (authorityService *AuthorityService) CopyAuthority(adminAuthorityID uint, copyInfo response.SysAuthorityCopyResponse) (authority system.SysAuthority, err error) {
	var authorityBox system.SysAuthority
	if !errors.Is(global.GVA_DB.Where("authority_id = ?", copyInfo.Authority.AuthorityId).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return authority, ErrRoleExistence
	}
	copyInfo.Authority.Children = []system.SysAuthority{}
	menus, err := MenuServiceApp.GetMenuAuthority(&request.GetAuthorityId{AuthorityId: copyInfo.OldAuthorityId})
	if err != nil {
		return
	}
	var baseMenu []system.SysBaseMenu
	for _, v := range menus {
		intNum := v.MenuId
		v.SysBaseMenu.ID = uint(intNum)
		baseMenu = append(baseMenu, v.SysBaseMenu)
	}
	copyInfo.Authority.SysBaseMenus = baseMenu
	err = global.GVA_DB.Create(&copyInfo.Authority).Error
	if err != nil {
		return
	}

	var btns []system.SysAuthorityBtn

	err = global.GVA_DB.Find(&btns, "authority_id = ?", copyInfo.OldAuthorityId).Error
	if err != nil {
		return
	}
	if len(btns) > 0 {
		for i := range btns {
			btns[i].AuthorityId = copyInfo.Authority.AuthorityId
		}
		err = global.GVA_DB.Create(&btns).Error

		if err != nil {
			return
		}
	}
	paths := CasbinServiceApp.GetPolicyPathByAuthorityId(copyInfo.OldAuthorityId)
	err = CasbinServiceApp.UpdateCasbin(adminAuthorityID, copyInfo.Authority.AuthorityId, paths)
	if err != nil {
		_ = authorityService.DeleteAuthority(&copyInfo.Authority)
	}
	return copyInfo.Authority, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateAuthority
//@description: 更改一个角色
//@param: auth model.SysAuthority
//@return: authority system.SysAuthority, err error

func (authorityService *AuthorityService) UpdateAuthority(auth system.SysAuthority) (authority system.SysAuthority, err error) {
	var oldAuthority system.SysAuthority
	err = global.GVA_DB.Where("authority_id = ?", auth.AuthorityId).First(&oldAuthority).Error
	if err != nil {
		global.GVA_LOG.Debug(err.Error())
		return system.SysAuthority{}, errors.New("查询角色数据失败")
	}
	err = global.GVA_DB.Model(&oldAuthority).Updates(&auth).Error
	return auth, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAuthority
//@description: 删除角色
//@param: auth *model.SysAuthority
//@return: err error

func (authorityService *AuthorityService) DeleteAuthority(auth *system.SysAuthority) error {
	if errors.Is(global.GVA_DB.Debug().Preload("Users").First(&auth).Error, gorm.ErrRecordNotFound) {
		return errors.New("该角色不存在")
	}
	if len(auth.Users) != 0 {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(global.GVA_DB.Where("authority_id = ?", auth.AuthorityId).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(global.GVA_DB.Where("parent_id = ?", auth.AuthorityId).First(&system.SysAuthority{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色存在子角色不允许删除")
	}

	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var err error
		if err = tx.Preload("SysBaseMenus").Preload("DataAuthorityId").Where("authority_id = ?", auth.AuthorityId).First(auth).Unscoped().Delete(auth).Error; err != nil {
			return err
		}

		if len(auth.SysBaseMenus) > 0 {
			if err = tx.Model(auth).Association("SysBaseMenus").Delete(auth.SysBaseMenus); err != nil {
				return err
			}
			// err = db.Association("SysBaseMenus").Delete(&auth)
		}
		if len(auth.DataAuthorityId) > 0 {
			if err = tx.Model(auth).Association("DataAuthorityId").Delete(auth.DataAuthorityId); err != nil {
				return err
			}
		}

		if err = tx.Delete(&system.SysUserAuthority{}, "sys_authority_authority_id = ?", auth.AuthorityId).Error; err != nil {
			return err
		}
		if err = tx.Where("authority_id = ?", auth.AuthorityId).Delete(&[]system.SysAuthorityBtn{}).Error; err != nil {
			return err
		}

		authorityId := strconv.Itoa(int(auth.AuthorityId))

		if err = CasbinServiceApp.RemoveFilteredPolicy(tx, authorityId); err != nil {
			return err
		}

		return nil
	})
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAuthorityInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: list interface{}, total int64, err error

func (authorityService *AuthorityService) GetAuthorityInfoList(authorityID uint) (list []system.SysAuthority, err error) {
	var authority system.SysAuthority
	err = global.GVA_DB.Where("authority_id = ?", authorityID).First(&authority).Error
	if err != nil {
		return nil, err
	}
	var authorities []system.SysAuthority
	db := global.GVA_DB.Model(&system.SysAuthority{})
	if global.GVA_CONFIG.System.UseStrictAuth {
		// 当开启了严格树形结构后
		if *authority.ParentId == 0 {
			// 只有顶级角色可以修改自己的权限和以下权限
			err = db.Preload("DataAuthorityId").Where("authority_id = ?", authorityID).Find(&authorities).Error
		} else {
			// 非顶级角色只能修改以下权限
			err = db.Debug().Preload("DataAuthorityId").Where("parent_id = ?", authorityID).Find(&authorities).Error
		}
	} else {
		err = db.Preload("DataAuthorityId").Where("parent_id = ?", "0").Find(&authorities).Error
	}

	for k := range authorities {
		err = authorityService.findChildrenAuthority(&authorities[k])
	}
	return authorities, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAuthorityInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: list interface{}, total int64, err error

func (authorityService *AuthorityService) GetStructAuthorityList(authorityID uint) (list []uint, err error) {
	var auth system.SysAuthority
	_ = global.GVA_DB.First(&auth, "authority_id = ?", authorityID).Error
	var authorities []system.SysAuthority
	err = global.GVA_DB.Preload("DataAuthorityId").Where("parent_id = ?", authorityID).Find(&authorities).Error
	if len(authorities) > 0 {
		for k := range authorities {
			list = append(list, authorities[k].AuthorityId)
			childrenList, err := authorityService.GetStructAuthorityList(authorities[k].AuthorityId)
			if err == nil {
				list = append(list, childrenList...)
			}
		}
	}
	if *auth.ParentId == 0 {
		list = append(list, authorityID)
	}
	return list, err
}

func (authorityService *AuthorityService) CheckAuthorityIDAuth(authorityID, targetID uint) (err error) {
	if !global.GVA_CONFIG.System.UseStrictAuth {
		return nil
	}
	authIDS, err := authorityService.GetStructAuthorityList(authorityID)
	if err != nil {
		return err
	}
	hasAuth := false
	for _, v := range authIDS {
		if v == targetID {
			hasAuth = true
			break
		}
	}
	if !hasAuth {
		return errors.New("您提交的角色ID不合法")
	}
	return nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAuthorityInfo
//@description: 获取所有角色信息
//@param: auth model.SysAuthority
//@return: sa system.SysAuthority, err error

func (authorityService *AuthorityService) GetAuthorityInfo(auth system.SysAuthority) (sa system.SysAuthority, err error) {
	err = global.GVA_DB.Preload("DataAuthorityId").Where("authority_id = ?", auth.AuthorityId).First(&sa).Error
	return sa, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetDataAuthority
//@description: 设置角色资源权限
//@param: auth model.SysAuthority
//@return: error

func (authorityService *AuthorityService) SetDataAuthority(adminAuthorityID uint, auth system.SysAuthority) error {
	var checkIDs []uint
	checkIDs = append(checkIDs, auth.AuthorityId)
	for i := range auth.DataAuthorityId {
		checkIDs = append(checkIDs, auth.DataAuthorityId[i].AuthorityId)
	}

	for i := range checkIDs {
		err := authorityService.CheckAuthorityIDAuth(adminAuthorityID, checkIDs[i])
		if err != nil {
			return err
		}
	}

	var s system.SysAuthority
	global.GVA_DB.Preload("DataAuthorityId").First(&s, "authority_id = ?", auth.AuthorityId)
	err := global.GVA_DB.Model(&s).Association("DataAuthorityId").Replace(&auth.DataAuthorityId)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetMenuAuthority
//@description: 菜单与角色绑定
//@param: auth *model.SysAuthority
//@return: error

func (authorityService *AuthorityService) SetMenuAuthority(auth *system.SysAuthority) error {
	var s system.SysAuthority
	global.GVA_DB.Preload("SysBaseMenus").First(&s, "authority_id = ?", auth.AuthorityId)
	err := global.GVA_DB.Model(&s).Association("SysBaseMenus").Replace(&auth.SysBaseMenus)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: findChildrenAuthority
//@description: 查询子角色
//@param: authority *model.SysAuthority
//@return: err error

func (authorityService *AuthorityService) findChildrenAuthority(authority *system.SysAuthority) (err error) {
	err = global.GVA_DB.Preload("DataAuthorityId").Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = authorityService.findChildrenAuthority(&authority.Children[k])
		}
	}
	return err
}

func (authorityService *AuthorityService) GetParentAuthorityID(authorityID uint) (parentID uint, err error) {
	var authority system.SysAuthority
	err = global.GVA_DB.Where("authority_id = ?", authorityID).First(&authority).Error
	if err != nil {
		return
	}
	return *authority.ParentId, nil
}

// GetUserIdsByAuthorityId 获取拥有指定角色的所有用户ID
func (authorityService *AuthorityService) GetUserIdsByAuthorityId(authorityId uint) (userIds []uint, err error) {
	var records []system.SysUserAuthority
	err = global.GVA_DB.Where("sys_authority_authority_id = ?", authorityId).Find(&records).Error
	if err != nil {
		return nil, err
	}
	for _, r := range records {
		userIds = append(userIds, r.SysUserId)
	}
	return userIds, nil
}

// SetRoleUsers 全量覆盖某角色关联的用户列表
// 入参：角色ID + 目标用户ID列表，保存时将该角色的关联关系完全替换为传入列表
func (authorityService *AuthorityService) SetRoleUsers(authorityId uint, userIds []uint) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 1. 查出当前拥有该角色的所有用户ID
		var existingRecords []system.SysUserAuthority
		if err := tx.Where("sys_authority_authority_id = ?", authorityId).Find(&existingRecords).Error; err != nil {
			return err
		}

		currentSet := make(map[uint]struct{})
		for _, r := range existingRecords {
			currentSet[r.SysUserId] = struct{}{}
		}

		targetSet := make(map[uint]struct{})
		for _, id := range userIds {
			targetSet[id] = struct{}{}
		}

		// 2. 删除该角色所有已有的用户关联
		if err := tx.Delete(&system.SysUserAuthority{}, "sys_authority_authority_id = ?", authorityId).Error; err != nil {
			return err
		}

		// 3. 对被移除的用户：若该角色是其主角色，则将主角色切换为其剩余的其他角色
		for userId := range currentSet {
			if _, ok := targetSet[userId]; ok {
				continue // 仍在目标列表中，不处理
			}
			var user system.SysUser
			if err := tx.First(&user, "id = ?", userId).Error; err != nil {
				continue
			}
			if user.AuthorityId == authorityId {
				// 从剩余关联（已删除当前角色后）中找另一个角色作为主角色
				var another system.SysUserAuthority
				if err := tx.Where("sys_user_id = ?", userId).First(&another).Error; err != nil {
					// 没有其他角色，主角色保持不变，不做处理
					continue
				}
				if err := tx.Model(&system.SysUser{}).Where("id = ?", userId).
					Update("authority_id", another.SysAuthorityAuthorityId).Error; err != nil {
					return err
				}
			}
		}

		// 4. 批量插入新的关联记录
		if len(userIds) > 0 {
			newRecords := make([]system.SysUserAuthority, 0, len(userIds))
			for _, userId := range userIds {
				newRecords = append(newRecords, system.SysUserAuthority{
					SysUserId:               userId,
					SysAuthorityAuthorityId: authorityId,
				})
			}
			if err := tx.Create(&newRecords).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
