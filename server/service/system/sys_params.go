package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type SysParamsService struct{}

// CreateSysParams 创建参数记录
// Author [Mr.奇淼](https://github.com/pixelmaxQm)
func (sysParamsService *SysParamsService) CreateSysParams(sysParams *system.SysParams) (err error) {
	err = global.GVA_DB.Create(sysParams).Error
	return err
}

// DeleteSysParams 删除参数记录
// Author [Mr.奇淼](https://github.com/pixelmaxQm)
func (sysParamsService *SysParamsService) DeleteSysParams(ID string) (err error) {
	err = global.GVA_DB.Delete(&system.SysParams{}, "id = ?", ID).Error
	return err
}

// DeleteSysParamsByIds 批量删除参数记录
// Author [Mr.奇淼](https://github.com/pixelmaxQm)
func (sysParamsService *SysParamsService) DeleteSysParamsByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]system.SysParams{}, "id in ?", IDs).Error
	return err
}

// UpdateSysParams 更新参数记录
// Author [Mr.奇淼](https://github.com/pixelmaxQm)
func (sysParamsService *SysParamsService) UpdateSysParams(sysParams system.SysParams) (err error) {
	err = global.GVA_DB.Model(&system.SysParams{}).Where("id = ?", sysParams.ID).Updates(&sysParams).Error
	return err
}

// GetSysParams 根据ID获取参数记录
// Author [Mr.奇淼](https://github.com/pixelmaxQm)
func (sysParamsService *SysParamsService) GetSysParams(ID string) (sysParams system.SysParams, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&sysParams).Error
	return
}

// GetSysParamsInfoList 分页获取参数记录
// Author [Mr.奇淼](https://github.com/pixelmaxQm)
func (sysParamsService *SysParamsService) GetSysParamsInfoList(info systemReq.SysParamsSearch) (list []system.SysParams, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&system.SysParams{})
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
func (sysParamsService *SysParamsService) GetSysParam(key string) (param system.SysParams, err error) {
	err = global.GVA_DB.Where(system.SysParams{Key: key}).First(&param).Error
	return
}
