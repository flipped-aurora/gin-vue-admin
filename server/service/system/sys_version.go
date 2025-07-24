package system

import (
	"context"
	"encoding/json"
	"fmt"
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
func (sysVersionService *SysVersionService) DeleteSysVersion(ctx context.Context, ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&system.SysVersion{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&system.SysVersion{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteSysVersionByIds 批量删除版本管理记录
// Author [yourname](https://github.com/yourname)
func (sysVersionService *SysVersionService) DeleteSysVersionByIds(ctx context.Context, IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&system.SysVersion{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&system.SysVersion{}).Error; err != nil {
			return err
		}
		return nil
	})
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
func (sysVersionService *SysVersionService) ImportMenus(ctx context.Context, menusData interface{}) error {
	menusBytes, err := json.Marshal(menusData)
	if err != nil {
		return fmt.Errorf("序列化菜单数据失败: %v", err)
	}

	var menus []system.SysBaseMenu
	err = json.Unmarshal(menusBytes, &menus)
	if err != nil {
		return fmt.Errorf("反序列化菜单数据失败: %v", err)
	}

	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		for _, menu := range menus {
			// 检查菜单是否已存在（根据path和name判断）
			var existingMenu system.SysBaseMenu
			err := tx.Where("path = ? AND name = ?", menu.Path, menu.Name).First(&existingMenu).Error
			if err == gorm.ErrRecordNotFound {
				// 菜单不存在，创建新菜单
				menu.ID = 0 // 重置ID，让数据库自动生成
				// 重置关联数据的ID
				for i := range menu.Parameters {
					menu.Parameters[i].ID = 0
					menu.Parameters[i].SysBaseMenuID = 0
				}
				for i := range menu.MenuBtn {
					menu.MenuBtn[i].ID = 0
					menu.MenuBtn[i].SysBaseMenuID = 0
				}
				if err := tx.Create(&menu).Error; err != nil {
					return fmt.Errorf("创建菜单失败: %v", err)
				}
			} else if err != nil {
				return fmt.Errorf("查询菜单失败: %v", err)
			}
			// 如果菜单已存在，跳过创建
		}
		return nil
	})
}

// ImportApis 导入API数据
func (sysVersionService *SysVersionService) ImportApis(ctx context.Context, apisData interface{}) error {
	apisBytes, err := json.Marshal(apisData)
	if err != nil {
		return fmt.Errorf("序列化API数据失败: %v", err)
	}

	var apis []system.SysApi
	err = json.Unmarshal(apisBytes, &apis)
	if err != nil {
		return fmt.Errorf("反序列化API数据失败: %v", err)
	}

	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		for _, api := range apis {
			// 检查API是否已存在（根据path和method判断）
			var existingApi system.SysApi
			err := tx.Where("path = ? AND method = ?", api.Path, api.Method).First(&existingApi).Error
			if err == gorm.ErrRecordNotFound {
				// API不存在，创建新API
				api.ID = 0 // 重置ID，让数据库自动生成
				if err := tx.Create(&api).Error; err != nil {
					return fmt.Errorf("创建API失败: %v", err)
				}
			} else if err != nil {
				return fmt.Errorf("查询API失败: %v", err)
			}
			// 如果API已存在，跳过创建
		}
		return nil
	})
}
