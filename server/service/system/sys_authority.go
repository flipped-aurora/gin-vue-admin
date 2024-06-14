package system

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system/internal"
	"github.com/pkg/errors"
	"strconv"

	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"gorm.io/gorm"
)

var ErrRoleExistence = errors.New("存在相同角色id")

type AuthorityService struct{}

var AuthorityServiceApp = new(AuthorityService)

// CreateAuthority 创建一个角色
//
//	@author:	[piexlmax](https://github.com/piexlmax)
//	@param:		auth model.SysAuthority
//	@return:	authority system.SysAuthority, err error
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

// Copy 复制角色和子角色

func (authorityService *AuthorityService) Copy(ctx context.Context, info systemReq.AuthorityCopy) error {
	tree, ids, err := internal.Authority.Tree(ctx)
	var old system.SysAuthority
	{
		err = global.GVA_DB.Where("authority_id = ?", info.OldAuthorityId).First(&old).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.Wrap(err, "旧角色不存在!")
		}
		internal.Authority.Children(tree, &old) // 子角色
	} // 查询旧角色及其子角色
	rule := make(map[uint][]string)
	casbin := CasbinServiceApp.Casbin()
	{
		rules := casbin.GetGroupingPolicy()
		for i := 0; i < len(rules); i++ {
			if len(rules[i]) >= 4 {
				var id uint64
				id, err = strconv.ParseUint(rules[i][1], 10, 64)
				if err != nil {
					return errors.Wrap(err, "角色id转换失败!")
				}
				_, ok := rule[uint(id)]
				if !ok {
					rule[uint(id)] = rules[i]
				}
			}
		}
	} // casbin
	create := info.Create()                                         // 复制=>主角色
	creates := make(map[uint]uint)                                  // 复制=>map[复制角色id]子角色id
	children := internal.Authority.Copy(ids, creates, &old, create) // 复制=>子角色
	var (
		rules   [][]string
		menus   []system.SysAuthorityMenu
		buttons []system.SysAuthorityBtn
	)
	{
		creates[old.AuthorityId] = create.AuthorityId        // 复制旧角色id => 复制主角色id
		rules = internal.Authority.CopyCasbin(rule, creates) // 复制角色的权限
	} // 复制角色的权限
	{
		{
			menus, err = internal.Authority.CopyMenus(ctx, creates)
			if err != nil {
				return err
			}
		} // 复制角色的菜单权限
		{
			buttons, err = internal.Authority.CopyMenuButtons(ctx, creates)
			if err != nil {
				return err
			}
		} // 复制角色的菜单按钮权限
	} // 复制角色的菜单&菜单按钮权限
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		err = tx.WithContext(ctx).Model(&system.SysAuthority{}).Create(&create).Error
		if err != nil {
			return errors.Wrap(err, "创建主角色失败!")
		}
		err = tx.WithContext(ctx).Model(&system.SysAuthority{}).Create(&children).Error
		if err != nil {
			return errors.Wrap(err, "创建子角色失败!")
		}
		err = tx.WithContext(ctx).Model(&system.SysAuthorityMenu{}).Create(&menus).Error
		if err != nil {
			return errors.Wrap(err, "创建角色菜单权限失败!")
		}
		err = tx.WithContext(ctx).Model(&system.SysAuthorityMenu{}).Create(&buttons).Error
		if err != nil {
			return errors.Wrap(err, "创建角色菜单按钮权限失败!")
		}
		var success bool
		success, err = casbin.AddPolicies(rules)
		if err != nil {
			return errors.Wrap(err, "复制角色权限失败!")
		}
		if !success {
			return errors.New("复制角色权限失败!")
		}
		return nil
	})
	if err != nil {
		var success bool
		success, err = casbin.RemoveGroupingPolicies(rules)
		if err != nil {
			return errors.Wrap(err, "删除casbin失败!")
		}
		if !success {
			return errors.New("事务回滚复制角色权限失败!")
		}
		return err
	}
	return nil
}

// CopyAuthority 复制一个角色
//
//	@author:	[piexlmax](https://github.com/piexlmax)
//	@param:		copyInfo response.SysAuthorityCopyResponse
//	@return:	authority system.SysAuthority, err error
func (authorityService *AuthorityService) CopyAuthority(copyInfo response.SysAuthorityCopyResponse) (authority system.SysAuthority, err error) {
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
	err = CasbinServiceApp.UpdateCasbin(copyInfo.Authority.AuthorityId, paths)
	if err != nil {
		_ = authorityService.DeleteAuthority(&copyInfo.Authority)
	}
	return copyInfo.Authority, err
}

// UpdateAuthority 更改一个角色
//
//	@author:	[piexlmax](https://github.com/piexlmax)
//	@param:		auth model.SysAuthority
//	@return:	authority system.SysAuthority, err error
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

// DeleteAuthority 删除角色
//
//	@author:	[piexlmax](https://github.com/piexlmax)
//	@param:		auth *model.SysAuthority
//	@return:	err error
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

// GetAuthorityInfoList 分页获取数据
//
//	@author:	[piexlmax](https://github.com/piexlmax)
//	@param:		info request.PageInfo
//	@return:	list interface{}, total int64, err error
func (authorityService *AuthorityService) GetAuthorityInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysAuthority{})
	if err = db.Where("parent_id = ?", "0").Count(&total).Error; total == 0 || err != nil {
		return
	}
	var authority []system.SysAuthority
	err = db.Limit(limit).Offset(offset).Preload("DataAuthorityId").Where("parent_id = ?", "0").Find(&authority).Error
	for k := range authority {
		err = authorityService.findChildrenAuthority(&authority[k])
	}
	return authority, total, err
}

//	@author:		[piexlmax](https://github.com/piexlmax)
//	@function:		GetAuthorityInfo
//	@description:	获取所有角色信息
//	@param:			auth model.SysAuthority
//	@return:		sa system.SysAuthority, err error

func (authorityService *AuthorityService) GetAuthorityInfo(auth system.SysAuthority) (sa system.SysAuthority, err error) {
	err = global.GVA_DB.Preload("DataAuthorityId").Where("authority_id = ?", auth.AuthorityId).First(&sa).Error
	return sa, err
}

// SetDataAuthority 设置角色资源权限
//
//	@author:	[piexlmax](https://github.com/piexlmax)
//	@param:		auth model.SysAuthority
//	@return:	error
func (authorityService *AuthorityService) SetDataAuthority(auth system.SysAuthority) error {
	var s system.SysAuthority
	global.GVA_DB.Preload("DataAuthorityId").First(&s, "authority_id = ?", auth.AuthorityId)
	err := global.GVA_DB.Model(&s).Association("DataAuthorityId").Replace(&auth.DataAuthorityId)
	return err
}

// SetMenuAuthority 菜单与角色绑定
//
//	@author:	[piexlmax](https://github.com/piexlmax)
//	@param:		auth *model.SysAuthority
//	@return:	error
func (authorityService *AuthorityService) SetMenuAuthority(auth *system.SysAuthority) error {
	var s system.SysAuthority
	global.GVA_DB.Preload("SysBaseMenus").First(&s, "authority_id = ?", auth.AuthorityId)
	err := global.GVA_DB.Model(&s).Association("SysBaseMenus").Replace(&auth.SysBaseMenus)
	return err
}

// findChildrenAuthority 查询子角色
//
//	@author:	[piexlmax](https://github.com/piexlmax)
//	@param:		authority *model.SysAuthority
//	@return:	err error
func (authorityService *AuthorityService) findChildrenAuthority(authority *system.SysAuthority) (err error) {
	err = global.GVA_DB.Preload("DataAuthorityId").Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = authorityService.findChildrenAuthority(&authority.Children[k])
		}
	}
	return err
}
