package system

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gorm.io/gorm"
)

type BaseMenuService struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteBaseMenu
//@description: 删除基础路由
//@param: id float64
//@return: err error

func (baseMenuService *BaseMenuService) DeleteBaseMenu(id int) (err error) {
	err = global.GVA_DB.Preload("MenuBtn").Preload("Parameters").Where("parent_id = ?", id).First(&system.SysBaseMenu{}).Error
	if err != nil {
		var menu system.SysBaseMenu
		db := global.GVA_DB.Preload("SysAuthoritys").Where("id = ?", id).First(&menu).Delete(&menu)
		err = global.GVA_DB.Delete(&system.SysBaseMenuParameter{}, "sys_base_menu_id = ?", id).Error
		err = global.GVA_DB.Delete(&system.SysBaseMenuBtn{}, "sys_base_menu_id = ?", id).Error
		err = global.GVA_DB.Delete(&system.SysAuthorityBtn{}, "sys_menu_id = ?", id).Error
		if err != nil {
			return err
		}
		if len(menu.SysAuthoritys) > 0 {
			err = global.GVA_DB.Model(&menu).Association("SysAuthoritys").Delete(&menu.SysAuthoritys)
		} else {
			err = db.Error
			if err != nil {
				return
			}
		}
	} else {
		return errors.New("此菜单存在子菜单不可删除")
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateBaseMenu
//@description: 更新路由
//@param: menu model.SysBaseMenu
//@return: err error

func (baseMenuService *BaseMenuService) UpdateBaseMenu(menu system.SysBaseMenu) (err error) {
	var oldMenu system.SysBaseMenu
	upDateMap := make(map[string]interface{})
	upDateMap["keep_alive"] = menu.KeepAlive
	upDateMap["close_tab"] = menu.CloseTab
	upDateMap["default_menu"] = menu.DefaultMenu
	upDateMap["parent_id"] = menu.ParentId
	upDateMap["path"] = menu.Path
	upDateMap["name"] = menu.Name
	upDateMap["hidden"] = menu.Hidden
	upDateMap["component"] = menu.Component
	upDateMap["title"] = menu.Title
	upDateMap["icon"] = menu.Icon
	upDateMap["sort"] = menu.Sort

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		db := tx.Where("id = ?", menu.ID).Find(&oldMenu)
		if oldMenu.Name != menu.Name {
			if !errors.Is(tx.Where("id <> ? AND name = ?", menu.ID, menu.Name).First(&system.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
				global.GVA_LOG.Debug("存在相同name修改失败")
				return errors.New("存在相同name修改失败")
			}
		}
		txErr := tx.Unscoped().Delete(&system.SysBaseMenuParameter{}, "sys_base_menu_id = ?", menu.ID).Error
		if txErr != nil {
			global.GVA_LOG.Debug(txErr.Error())
			return txErr
		}
		txErr = tx.Unscoped().Delete(&system.SysBaseMenuBtn{}, "sys_base_menu_id = ?", menu.ID).Error
		if txErr != nil {
			global.GVA_LOG.Debug(txErr.Error())
			return txErr
		}
		if len(menu.Parameters) > 0 {
			for k := range menu.Parameters {
				menu.Parameters[k].SysBaseMenuID = menu.ID
			}
			txErr = tx.Create(&menu.Parameters).Error
			if txErr != nil {
				global.GVA_LOG.Debug(txErr.Error())
				return txErr
			}
		}

		if len(menu.MenuBtn) > 0 {
			for k := range menu.MenuBtn {
				menu.MenuBtn[k].SysBaseMenuID = menu.ID
			}
			txErr = tx.Create(&menu.MenuBtn).Error
			if txErr != nil {
				global.GVA_LOG.Debug(txErr.Error())
				return txErr
			}
		}

		txErr = db.Updates(upDateMap).Error
		if txErr != nil {
			global.GVA_LOG.Debug(txErr.Error())
			return txErr
		}
		return nil
	})
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBaseMenuById
//@description: 返回当前选中menu
//@param: id float64
//@return: menu system.SysBaseMenu, err error

func (baseMenuService *BaseMenuService) GetBaseMenuById(id int) (menu system.SysBaseMenu, err error) {
	err = global.GVA_DB.Preload("MenuBtn").Preload("Parameters").Where("id = ?", id).First(&menu).Error
	return
}
