package utils

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

func RegisterApis( apis ...system.SysApi) {
	err := global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		for _, api := range apis {
			err := tx.Model(system.SysApi{}).Where("path = ? AND method = ? AND api_group = ? ", api.Path, api.Method, api.ApiGroup).FirstOrCreate(&api).Error
			if err != nil {
				zap.L().Error("注册API失败", zap.Error(err), zap.String("api", api.Path), zap.String("method", api.Method), zap.String("apiGroup", api.ApiGroup))
				return err
			}
		}
		return nil
	})
	if err != nil {
		zap.L().Error("注册API失败", zap.Error(err))
	}
}

func RegisterMenus( menus ...system.SysBaseMenu) {
	parentMenu := menus[0]
	otherMenus := menus[1:]
	err := global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(system.SysBaseMenu{}).Where("name = ? ", parentMenu.Name).FirstOrCreate(&parentMenu).Error
		if err != nil {
			zap.L().Error("注册菜单失败", zap.Error(err))
			return errors.Wrap(err, "注册菜单失败")
		}
		pid := parentMenu.ID
		for i := range otherMenus {
			otherMenus[i].ParentId = pid
			err = tx.Model(system.SysBaseMenu{}).Where("name = ? ", otherMenus[i].Name).FirstOrCreate(&otherMenus[i]).Error
			if err != nil {
				zap.L().Error("注册菜单失败", zap.Error(err))
				return errors.Wrap(err, "注册菜单失败")
			}
		}

		return nil
	})
	if err != nil {
		zap.L().Error("注册菜单失败", zap.Error(err))
	}

}
