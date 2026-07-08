package system

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type SysParamsService struct{}

// CreateSysParams 创建参数记录
// Author [Mr.奇淼](https://github.com/pixelmaxQm)
func (sysParamsService *SysParamsService) CreateSysParams(ctx context.Context, sysParams *system.SysParams) (err error) {
	err = global.GVA_DB.WithContext(ctx).Create(sysParams).Error
	return err
}

// DeleteSysParams 删除参数记录
// Author [Mr.奇淼](https://github.com/pixelmaxQm)
func (sysParamsService *SysParamsService) DeleteSysParams(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.WithContext(ctx).Delete(&system.SysParams{}, "id = ?", ID).Error
	return err
}

// DeleteSysParamsByIds 批量删除参数记录
// Author [Mr.奇淼](https://github.com/pixelmaxQm)
func (sysParamsService *SysParamsService) DeleteSysParamsByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.WithContext(ctx).Delete(&[]system.SysParams{}, "id in ?", IDs).Error
	return err
}

// UpdateSysParams 更新参数记录
// Author [Mr.奇淼](https://github.com/pixelmaxQm)
func (sysParamsService *SysParamsService) UpdateSysParams(ctx context.Context, sysParams system.SysParams) (err error) {
	err = global.GVA_DB.WithContext(ctx).Model(&system.SysParams{}).Where("id = ?", sysParams.ID).Updates(&sysParams).Error
	return err
}

// GetSysParams 根据ID获取参数记录
// Author [Mr.奇淼](https://github.com/pixelmaxQm)
func (sysParamsService *SysParamsService) GetSysParams(ctx context.Context, ID string) (sysParams system.SysParams, err error) {
	err = global.GVA_DB.WithContext(ctx).Where("id = ?", ID).First(&sysParams).Error
	return
}

// GetSysParamsInfoList 分页获取参数记录
// Author [Mr.奇淼](https://github.com/pixelmaxQm)
func (sysParamsService *SysParamsService) GetSysParamsInfoList(ctx context.Context, info systemReq.SysParamsSearch) (list []system.SysParams, total int64, err error) {
	limit, offset := info.LimitOffset()
	// 创建db
	db := global.GVA_DB.WithContext(ctx).Model(&system.SysParams{})
	var sysParamss []system.SysParams
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Key != "" {
		db = db.Where("key LIKE ?", "%"+info.Key+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&sysParamss).Error
	return sysParamss, total, err
}

// GetSysParam 根据key获取参数value
// Author [Mr.奇淼](https://github.com/pixelmaxQm)
func (sysParamsService *SysParamsService) GetSysParam(ctx context.Context, key string) (param system.SysParams, err error) {
	err = global.GVA_DB.WithContext(ctx).Where(system.SysParams{Key: key}).First(&param).Error
	return
}
