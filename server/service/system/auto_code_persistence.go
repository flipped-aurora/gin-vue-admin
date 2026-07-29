package system

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"gorm.io/gorm"
)

var errAutoCodeMenuConflict = errors.New("自动代码菜单名称冲突")

func persistAutoCodeMetadata(
	ctx context.Context,
	db *gorm.DB,
	info request.AutoCode,
	packageTemplate string,
	history request.SysAutoHistoryCreate,
) error {
	if db == nil {
		return errors.New("自动代码数据库未初始化")
	}
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := persistAutoCodeAPIs(tx, info, &history); err != nil {
			return err
		}
		if err := persistAutoCodeMenu(tx, info, packageTemplate, &history); err != nil {
			return err
		}
		if err := persistAutoCodeExportTemplate(tx, info, &history); err != nil {
			return err
		}
		entity := history.Create()
		if err := tx.Create(&entity).Error; err != nil {
			return fmt.Errorf("创建自动代码历史失败: %w", err)
		}
		return nil
	})
}

func persistAutoCodeAPIs(tx *gorm.DB, info request.AutoCode, history *request.SysAutoHistoryCreate) error {
	if !info.AutoCreateApiToSql || info.OnlyTemplate {
		return nil
	}
	for _, desired := range info.Apis() {
		var existing model.SysApi
		err := tx.Where("path = ? AND method = ?", desired.Path, desired.Method).First(&existing).Error
		switch {
		case err == nil:
			history.ApiIDs = append(history.ApiIDs, existing.ID)
		case errors.Is(err, gorm.ErrRecordNotFound):
			if err = tx.Create(&desired).Error; err != nil {
				return fmt.Errorf("创建自动代码 API %s %s 失败: %w", desired.Method, desired.Path, err)
			}
			history.ApiIDs = append(history.ApiIDs, desired.ID)
		default:
			return fmt.Errorf("查询自动代码 API %s %s 失败: %w", desired.Method, desired.Path, err)
		}
	}
	return nil
}

func persistAutoCodeMenu(
	tx *gorm.DB,
	info request.AutoCode,
	packageTemplate string,
	history *request.SysAutoHistoryCreate,
) error {
	if !info.AutoCreateMenuToSql {
		return nil
	}
	desired := info.Menu(packageTemplate)
	var existing model.SysBaseMenu
	err := tx.Where("name = ?", desired.Name).First(&existing).Error
	switch {
	case err == nil:
		if existing.Name != desired.Name || existing.Path != desired.Path || existing.Component != desired.Component {
			return fmt.Errorf(
				"%w: name=%s, 已存在 path=%s component=%s, 期望 path=%s component=%s",
				errAutoCodeMenuConflict,
				desired.Name,
				existing.Path,
				existing.Component,
				desired.Path,
				desired.Component,
			)
		}
		history.MenuID = existing.ID
		return nil
	case !errors.Is(err, gorm.ErrRecordNotFound):
		return fmt.Errorf("查询自动代码菜单 %s 失败: %w", desired.Name, err)
	}

	if info.AutoCreateBtnAuth && !info.OnlyTemplate {
		desired.MenuBtn = []model.SysBaseMenuBtn{
			{Name: "add", Desc: "新增"},
			{Name: "batchDelete", Desc: "批量删除"},
			{Name: "delete", Desc: "删除"},
			{Name: "edit", Desc: "编辑"},
			{Name: "info", Desc: "详情"},
		}
		if info.HasExcel {
			desired.MenuBtn = append(desired.MenuBtn,
				model.SysBaseMenuBtn{Name: "exportTemplate", Desc: "导出模板"},
				model.SysBaseMenuBtn{Name: "exportExcel", Desc: "导出Excel"},
				model.SysBaseMenuBtn{Name: "importExcel", Desc: "导入Excel"},
			)
		}
	}
	if err = tx.Create(&desired).Error; err != nil {
		return fmt.Errorf("创建自动代码菜单 %s 失败: %w", desired.Name, err)
	}
	history.MenuID = desired.ID
	return nil
}

func persistAutoCodeExportTemplate(tx *gorm.DB, info request.AutoCode, history *request.SysAutoHistoryCreate) error {
	if !info.HasExcel {
		return nil
	}
	fields := make(map[string]string, len(info.Fields))
	for _, field := range info.Fields {
		if field != nil && field.Excel {
			fields[field.ColumnName] = field.FieldDesc
		}
	}
	templateInfo, err := json.Marshal(fields)
	if err != nil {
		return fmt.Errorf("序列化自动代码导出模板失败: %w", err)
	}
	name := info.Package + "_" + info.StructName
	entity := model.SysExportTemplate{
		DBName:       info.BusinessDB,
		Name:         name,
		TableName:    info.TableName,
		TemplateID:   name,
		TemplateInfo: string(templateInfo),
	}
	if err = tx.Create(&entity).Error; err != nil {
		return fmt.Errorf("创建自动代码导出模板 %s 失败: %w", name, err)
	}
	history.ExportTemplateID = entity.ID
	return nil
}

func autoCodeIdentityExists(ctx context.Context, db *gorm.DB, info request.AutoCode) (bool, error) {
	if db == nil {
		return false, errors.New("自动代码数据库未初始化")
	}
	var count int64
	err := db.WithContext(ctx).Model(&model.SysAutoCodeHistory{}).
		Where(
			"business_db = ? AND (struct_name = ? OR abbreviation = ?) AND package = ? AND flag = ?",
			info.BusinessDB,
			info.StructName,
			info.Abbreviation,
			info.Package,
			0,
		).
		Count(&count).Error
	if err != nil {
		return false, fmt.Errorf("查询自动代码重复记录失败: %w", err)
	}
	return count > 0, nil
}
