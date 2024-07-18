package system

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go/ast"
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

var AutocodeHistory = new(autoCodeHistory)

type autoCodeHistory struct{}

// Create 创建代码生成器历史记录
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (s *autoCodeHistory) Create(ctx context.Context, info request.SysAutoHistoryCreate) error {
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
func (s *autoCodeHistory) First(ctx context.Context, info common.GetById) (string, error) {
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
func (s *autoCodeHistory) Repeat(businessDB, structName, Package string) bool {
	var count int64
	global.GVA_DB.Model(&model.SysAutoCodeHistory{}).Where("business_db = ? and struct_name = ? and package = ? and flag = 0", businessDB, structName, Package).Count(&count)
	return count > 0
}

// RollBack 回滚
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (s *autoCodeHistory) RollBack(ctx context.Context, info request.SysAutoHistoryRollBack) error {
	var entity model.SysAutoCodeHistory
	err := global.GVA_DB.Where("id = ?", info.ID).First(&entity).Error
	if err != nil {
		return err
	}
	entity.AfterFirst()
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
	for _, value := range entity.Templates {
		if !filepath.IsAbs(value) {
			continue
		}
		newPath := filepath.Join(global.GVA_CONFIG.AutoCode.Root, "rm_file", strconv.FormatInt(int64(time.Now().Nanosecond()), 10), strings.TrimPrefix(value, global.GVA_CONFIG.AutoCode.Root))
		err = utils.FileMove(value, newPath)
		if err != nil {
			return errors.Wrapf(err, "[src:%s][dst:%s]文件移动失败!", value, newPath)
		}
	} // 移动文件
	for key, value := range entity.Injections {
		var file *ast.File
		file, err = value.Parse("", nil)
		if err != nil {
			return err
		}
		value.Rollback(file)
		err = value.Format("", nil, file)
		if err != nil {
			return err
		}
		fmt.Printf("[filepath:%s]回滚注入代码成功!\n", key)
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
func (s *autoCodeHistory) Delete(ctx context.Context, info common.GetById) error {
	err := global.GVA_DB.WithContext(ctx).Where("id = ?", info.Uint()).Delete(&model.SysAutoCodeHistory{}).Error
	if err != nil {
		return errors.Wrap(err, "删除失败!")
	}
	return nil
}

// GetList 获取系统历史数据
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (s *autoCodeHistory) GetList(ctx context.Context, info common.PageInfo) (list []model.SysAutoCodeHistory, total int64, err error) {
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
func (s *autoCodeHistory) DropTable(BusinessDb, tableName string) error {
	if BusinessDb != "" {
		return global.MustGetGlobalDBByDBName(BusinessDb).Exec("DROP TABLE " + tableName).Error
	} else {
		return global.GVA_DB.Exec("DROP TABLE " + tableName).Error
	}
}
