package system

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"gorm.io/gorm"
)

type SysVersionService struct{}

// CreateSysVersion 创建版本管理记录
// Author [yourname](https://github.com/yourname)
func (sysVersionService *SysVersionService) CreateSysVersion(ctx context.Context, sysVersion *system.SysVersion) (err error) {
	err = global.GVA_DB.Create(sysVersion).Error
	return err
}

// DeleteSysVersion 删除版本管理记录
// Author [yourname](https://github.com/yourname)
func (sysVersionService *SysVersionService) DeleteSysVersion(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&system.SysVersion{}, "id = ?", ID).Error
	return err
}

// DeleteSysVersionByIds 批量删除版本管理记录
// Author [yourname](https://github.com/yourname)
func (sysVersionService *SysVersionService) DeleteSysVersionByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Where("id in ?", IDs).Delete(&system.SysVersion{}).Error
	return err
}

// GetSysVersion 根据ID获取版本管理记录
// Author [yourname](https://github.com/yourname)
func (sysVersionService *SysVersionService) GetSysVersion(ctx context.Context, ID string) (sysVersion system.SysVersion, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&sysVersion).Error
	return
}

// GetSysVersionInfoList 分页获取版本管理记录
// Author [yourname](https://github.com/yourname)
func (sysVersionService *SysVersionService) GetSysVersionInfoList(ctx context.Context, info systemReq.SysVersionSearch) (list []system.SysVersion, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&system.SysVersion{})
	var sysVersions []system.SysVersion
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.VersionName != nil && *info.VersionName != "" {
		db = db.Where("version_name LIKE ?", "%"+*info.VersionName+"%")
	}
	if info.VersionCode != nil && *info.VersionCode != "" {
		db = db.Where("version_code = ?", *info.VersionCode)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&sysVersions).Error
	return sysVersions, total, err
}
func (sysVersionService *SysVersionService) GetSysVersionPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// GetMenusByIds 根据ID列表获取菜单数据
func (sysVersionService *SysVersionService) GetMenusByIds(ctx context.Context, ids []uint) (menus []system.SysBaseMenu, err error) {
	err = global.GVA_DB.Where("id in ?", ids).Preload("Parameters").Preload("MenuBtn").Find(&menus).Error
	return
}

// GetApisByIds 根据ID列表获取API数据
func (sysVersionService *SysVersionService) GetApisByIds(ctx context.Context, ids []uint) (apis []system.SysApi, err error) {
	err = global.GVA_DB.Where("id in ?", ids).Find(&apis).Error
	return
}

// ImportMenus 导入菜单数据
func (sysVersionService *SysVersionService) ImportMenus(ctx context.Context, menus []system.SysBaseMenu) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 递归创建菜单
		return sysVersionService.createMenusRecursively(tx, menus, 0)
	})
}

// createMenusRecursively 递归创建菜单
func (sysVersionService *SysVersionService) createMenusRecursively(tx *gorm.DB, menus []system.SysBaseMenu, parentId uint) error {
	for _, menu := range menus {
		// 检查菜单是否已存在
		var existingMenu system.SysBaseMenu
		if err := tx.Where("name = ? AND path = ?", menu.Name, menu.Path).First(&existingMenu).Error; err == nil {
			// 菜单已存在，使用现有菜单ID继续处理子菜单
			if len(menu.Children) > 0 {
				if err := sysVersionService.createMenusRecursively(tx, menu.Children, existingMenu.ID); err != nil {
					return err
				}
			}
			continue
		}

		// 保存参数和按钮数据，稍后处理
		parameters := menu.Parameters
		menuBtns := menu.MenuBtn
		children := menu.Children

		// 创建新菜单（不包含关联数据）
		newMenu := system.SysBaseMenu{
			ParentId:  parentId,
			Path:      menu.Path,
			Name:      menu.Name,
			Hidden:    menu.Hidden,
			Component: menu.Component,
			Sort:      menu.Sort,
			Meta:      menu.Meta,
		}

		if err := tx.Create(&newMenu).Error; err != nil {
			return err
		}

		// 创建参数
		if len(parameters) > 0 {
			for _, param := range parameters {
				newParam := system.SysBaseMenuParameter{
					SysBaseMenuID: newMenu.ID,
					Type:          param.Type,
					Key:           param.Key,
					Value:         param.Value,
				}
				if err := tx.Create(&newParam).Error; err != nil {
					return err
				}
			}
		}

		// 创建菜单按钮
		if len(menuBtns) > 0 {
			for _, btn := range menuBtns {
				newBtn := system.SysBaseMenuBtn{
					SysBaseMenuID: newMenu.ID,
					Name:          btn.Name,
					Desc:          btn.Desc,
				}
				if err := tx.Create(&newBtn).Error; err != nil {
					return err
				}
			}
		}

		// 递归处理子菜单
		if len(children) > 0 {
			if err := sysVersionService.createMenusRecursively(tx, children, newMenu.ID); err != nil {
				return err
			}
		}
	}
	return nil
}

// ImportApis 导入API数据
func (sysVersionService *SysVersionService) ImportApis(apis []system.SysApi) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		for _, api := range apis {
			// 检查API是否已存在
			var existingApi system.SysApi
			if err := tx.Where("path = ? AND method = ?", api.Path, api.Method).First(&existingApi).Error; err == nil {
				// API已存在，跳过
				continue
			}

			// 创建新API
			newApi := system.SysApi{
				Path:        api.Path,
				Description: api.Description,
				ApiGroup:    api.ApiGroup,
				Method:      api.Method,
			}

			if err := tx.Create(&newApi).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
