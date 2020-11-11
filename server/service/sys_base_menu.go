package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gorm.io/gorm"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteBaseMenu
//@description: 删除基础路由
//@param: id float64
//@return: err error

func DeleteBaseMenu(id float64) (err error) {
	err = global.GVA_DB.Preload("Parameters").Where("parent_id = ?", id).First(&model.SysBaseMenu{}).Error
	if err != nil {
		var menu model.SysBaseMenu
		db := global.GVA_DB.Preload("SysAuthoritys").Where("id = ?", id).First(&menu).Delete(&menu)
		err = global.GVA_DB.Delete(&model.SysBaseMenuParameter{}, "sys_base_menu_id = ?", id).Error
		if len(menu.SysAuthoritys) > 0 {
			err = global.GVA_DB.Model(&menu).Association("SysAuthoritys").Delete(&menu.SysAuthoritys)
		} else {
			err = db.Error
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
//@return:err error

func UpdateBaseMenu(menu model.SysBaseMenu) (err error) {
	var oldMenu model.SysBaseMenu
	upDateMap := make(map[string]interface{})
	upDateMap["keep_alive"] = menu.KeepAlive
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
			if !errors.Is(tx.Where("id <> ? AND name = ?", menu.ID, menu.Name).First(&model.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
				global.GVA_LOG.Debug("存在相同name修改失败")
				return errors.New("存在相同name修改失败")
			}
		}
		txErr := tx.Unscoped().Delete(&model.SysBaseMenuParameter{}, "sys_base_menu_id = ?", menu.ID).Error
		if txErr != nil {
			global.GVA_LOG.Debug(txErr.Error())
			return txErr
		}
		if len(menu.Parameters) > 0 {
			for k, _ := range menu.Parameters {
				menu.Parameters[k].SysBaseMenuID = menu.ID
			}
			txErr = tx.Create(&menu.Parameters).Error
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
//@return: err error, menu model.SysBaseMenu

func GetBaseMenuById(id float64) (err error, menu model.SysBaseMenu) {
	err = global.GVA_DB.Preload("Parameters").Where("id = ?", id).First(&menu).Error
	return
}
