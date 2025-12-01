package system

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type SysErrorService struct{}

// CreateSysError 创建错误日志记录
// Author [yourname](https://github.com/yourname)
func (sysErrorService *SysErrorService) CreateSysError(ctx context.Context, sysError *system.SysError) (err error) {
	err = global.GVA_DB.Create(sysError).Error
	return err
}

// DeleteSysError 删除错误日志记录
// Author [yourname](https://github.com/yourname)
func (sysErrorService *SysErrorService) DeleteSysError(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&system.SysError{}, "id = ?", ID).Error
	return err
}

// DeleteSysErrorByIds 批量删除错误日志记录
// Author [yourname](https://github.com/yourname)
func (sysErrorService *SysErrorService) DeleteSysErrorByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]system.SysError{}, "id in ?", IDs).Error
	return err
}

// UpdateSysError 更新错误日志记录
// Author [yourname](https://github.com/yourname)
func (sysErrorService *SysErrorService) UpdateSysError(ctx context.Context, sysError system.SysError) (err error) {
	err = global.GVA_DB.Model(&system.SysError{}).Where("id = ?", sysError.ID).Updates(&sysError).Error
	return err
}

// GetSysError 根据ID获取错误日志记录
// Author [yourname](https://github.com/yourname)
func (sysErrorService *SysErrorService) GetSysError(ctx context.Context, ID string) (sysError system.SysError, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&sysError).Error
	return
}

// GetSysErrorInfoList 分页获取错误日志记录
// Author [yourname](https://github.com/yourname)
func (sysErrorService *SysErrorService) GetSysErrorInfoList(ctx context.Context, info systemReq.SysErrorSearch) (list []system.SysError, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&system.SysError{}).Order("created_at desc")
	var sysErrors []system.SysError
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

	err = db.Find(&sysErrors).Error
	return sysErrors, total, err
}

// GetSysErrorSolution 异步处理错误
// Author [yourname](https://github.com/yourname)
func (sysErrorService *SysErrorService) GetSysErrorSolution(ctx context.Context, ID string) (err error) {
	// 立即更新为处理中
	err = global.GVA_DB.WithContext(ctx).Model(&system.SysError{}).Where("id = ?", ID).Update("status", "处理中").Error
	if err != nil {
		return err
	}

	// 异步协程在一分钟后更新为处理完成
	go func(id string) {
		// 查询当前错误信息用于生成方案
		var se system.SysError
		_ = global.GVA_DB.Model(&system.SysError{}).Where("id = ?", id).First(&se).Error

		// 构造 LLM 请求参数，使用管家模式(butler)根据错误信息生成解决方案
		var form, info string
		if se.Form != nil {
			form = *se.Form
		}
		if se.Info != nil {
			info = *se.Info
		}

		llmReq := common.JSONMap{
			"mode":    "solution",
			"command": "solution",
			"info":    info,
			"form":    form,
		}

		// 调用服务层 LLMAuto，忽略错误但尽量写入方案
		var solution string
		if data, err := (&AutoCodeService{}).LLMAuto(context.Background(), llmReq); err == nil {
			solution = fmt.Sprintf("%v", data)
			_ = global.GVA_DB.Model(&system.SysError{}).Where("id = ?", id).Updates(map[string]interface{}{"status": "处理完成", "solution": solution}).Error
		} else {
			// 即使生成失败也标记为完成，避免任务卡住
			_ = global.GVA_DB.Model(&system.SysError{}).Where("id = ?", id).Update("status", "处理失败").Error
		}
	}(ID)

	return nil
}
