package system

import (
	"context"
	"github.com/pkg/errors"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	common "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	request "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"go.uber.org/zap"
)

var AutocodeHistory = new(autocodeHistory)

type autocodeHistory struct{}

// Create 创建代码生成器历史记录
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (s *autocodeHistory) Create(ctx context.Context, info request.SysAutoHistoryCreate) error {
	create := info.Create()
	err := global.GVA_DB.WithContext(ctx).Create(&create).Error
	if err != nil {
		return errors.Wrap(err, "创建失败!")
	}
	return nil
}

// First 根据id获取代码生成器历史的数据
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (s *autocodeHistory) First(ctx context.Context, info common.GetById) (string, error) {
	var meta string
	err := global.GVA_DB.WithContext(ctx).Model(model.SysAutoCodeHistory{}).Where("id = ?", info.ID).Pluck("request", &meta).Error
	if err != nil {
		return "", errors.Wrap(err, "获取失败!")
	}
	return meta, nil
}

// Repeat 检测重复
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (s *autocodeHistory) Repeat(businessDB, structName, Package string) bool {
	var count int64
	global.GVA_DB.Model(&model.SysAutoCodeHistory{}).Where("business_db = ? and struct_name = ? and package = ? and flag = 0", businessDB, structName, Package).Count(&count)
	return count > 0
}

// RollBack 回滚
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (s *autocodeHistory) RollBack(ctx context.Context, info request.SysAutoHistoryRollBack) error {
	var entity model.SysAutoCodeHistory
	err := global.GVA_DB.Where("id = ?", info.ID).First(&entity).Error
	if err != nil {
		return err
	}
	if info.DeleteApi {
		ids := info.ApiIds(entity)
		err = ApiServiceApp.DeleteApisByIds(ids)
		if err != nil {
			global.GVA_LOG.Error("ClearTag DeleteApiByIds:", zap.Error(err))
		}
	} // 清除API表
	if info.DeleteMenu {
		err = BaseMenuServiceApp.DeleteBaseMenu(int(entity.MenuID))
		if err != nil {
			return errors.Wrap(err, "删除菜单失败!")
		}
	} // 清除菜单表
	if info.DeleteTable {
		err = s.DropTable(entity.BusinessDB, entity.Table)
		if err != nil {
			return errors.Wrap(err, "删除表失败!")
		}
	} // 删除表
	for _, value := range entity.Template {
		if !filepath.IsAbs(value) {
			continue
		}
		newPath := filepath.Join(
			global.GVA_CONFIG.AutoCode.Root,
			"rm_file",
			strconv.FormatInt(int64(time.Now().Nanosecond()), 10),
			strings.TrimPrefix(value, global.GVA_CONFIG.AutoCode.Root),
		)
		err = utils.FileMove(value, newPath)
		if err != nil {
			return errors.Wrapf(err, "[src:%s][dst:%s]文件移动失败!", value, newPath)
		}
	} // 移动文件
	for key, value := range entity.Enter {
		err = utils.AutoClearCode(key, value)
		if err != nil {
			return errors.Wrapf(err, "[key:%s][value:%s]注入代码清除失败!", key, value)
		}
	} // 清除注入代码
	err = global.GVA_DB.WithContext(ctx).Model(&model.SysAutoCodeHistory{}).Where("id = ?", info.ID).Update("flag", 1).Error
	if err != nil {
		return errors.Wrap(err, "更新失败!")
	}
	return nil
}

// Delete 删除历史数据
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (s *autocodeHistory) Delete(ctx context.Context, info common.GetById) error {
	err := global.GVA_DB.WithContext(ctx).Where("id = ?", info.Uint()).Delete(&model.SysAutoCodeHistory{}).Error
	if err != nil {
		return errors.Wrap(err, "删除失败!")
	}
	return nil
}

// GetList 获取系统历史数据
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (s *autocodeHistory) GetList(ctx context.Context, info common.PageInfo) (list []model.SysAutoCodeHistory, total int64, err error) {
	var entities []model.SysAutoCodeHistory
	db := global.GVA_DB.WithContext(ctx).Model(&model.SysAutoCodeHistory{})
	err = db.Count(&total).Error
	if err != nil {
		return nil, total, err
	}
	err = db.Scopes(info.Paginate()).Order("updated_at desc").Find(&entities).Error
	return entities, total, err
}

// DropTable 获取指定数据库和指定数据表的所有字段名,类型值等
// @author: [piexlmax](https://github.com/piexlmax)
func (s *autocodeHistory) DropTable(BusinessDb, tableName string) error {
	if BusinessDb != "" {
		return global.MustGetGlobalDBByDBName(BusinessDb).Exec("DROP TABLE " + tableName).Error
	} else {
		return global.GVA_DB.Exec("DROP TABLE " + tableName).Error
	}
}
