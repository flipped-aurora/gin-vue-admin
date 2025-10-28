package system

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type SysErrorService struct{}

// CreateSysError 创建错误日志记录
// Author [yourname](https://github.com/yourname)
func (sysErrprService *SysErrorService) CreateSysError(ctx context.Context, sysErrpr *system.SysError) (err error) {
	err = global.GVA_DB.Create(sysErrpr).Error
	return err
}

// DeleteSysError 删除错误日志记录
// Author [yourname](https://github.com/yourname)
func (sysErrprService *SysErrorService) DeleteSysError(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&system.SysError{}, "id = ?", ID).Error
	return err
}

// DeleteSysErrorByIds 批量删除错误日志记录
// Author [yourname](https://github.com/yourname)
func (sysErrprService *SysErrorService) DeleteSysErrorByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]system.SysError{}, "id in ?", IDs).Error
	return err
}

// UpdateSysError 更新错误日志记录
// Author [yourname](https://github.com/yourname)
func (sysErrprService *SysErrorService) UpdateSysError(ctx context.Context, sysErrpr system.SysError) (err error) {
	err = global.GVA_DB.Model(&system.SysError{}).Where("id = ?", sysErrpr.ID).Updates(&sysErrpr).Error
	return err
}

// GetSysError 根据ID获取错误日志记录
// Author [yourname](https://github.com/yourname)
func (sysErrprService *SysErrorService) GetSysError(ctx context.Context, ID string) (sysErrpr system.SysError, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&sysErrpr).Error
	return
}

// GetSysErrorInfoList 分页获取错误日志记录
// Author [yourname](https://github.com/yourname)
func (sysErrprService *SysErrorService) GetSysErrorInfoList(ctx context.Context, info systemReq.SysErrorSearch) (list []system.SysError, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&system.SysError{}).Order("created desc")
	var sysErrprs []system.SysError
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.Form != nil && *info.Form != "" {
		db = db.Where("form = ?", *info.Form)
	}
	if info.Info != nil && *info.Info != "" {
		db = db.Where("info LIKE ?", "%"+*info.Info+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&sysErrprs).Error
	return sysErrprs, total, err
}
