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
	if global.GVA_DB == nil {
		return nil
	}
	err = global.GVA_DB.WithContext(ctx).Create(sysError).Error
	return err
}

// DeleteSysError 删除错误日志记录
// Author [yourname](https://github.com/yourname)
func (sysErrorService *SysErrorService) DeleteSysError(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.WithContext(ctx).Delete(&system.SysError{}, "id = ?", ID).Error
	return err
}

// DeleteSysErrorByIds 批量删除错误日志记录
// Author [yourname](https://github.com/yourname)
func (sysErrorService *SysErrorService) DeleteSysErrorByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.WithContext(ctx).Delete(&[]system.SysError{}, "id in ?", IDs).Error
	return err
}

// UpdateSysError 更新错误日志记录
// Author [yourname](https://github.com/yourname)
func (sysErrorService *SysErrorService) UpdateSysError(ctx context.Context, sysError system.SysError) (err error) {
	err = global.GVA_DB.WithContext(ctx).Model(&system.SysError{}).Where("id = ?", sysError.ID).Updates(&sysError).Error
	return err
}

// GetSysError 根据ID获取错误日志记录
// Author [yourname](https://github.com/yourname)
func (sysErrorService *SysErrorService) GetSysError(ctx context.Context, ID string) (sysError system.SysError, err error) {
	err = global.GVA_DB.WithContext(ctx).Where("id = ?", ID).First(&sysError).Error
	return
}

// GetSysErrorInfoList 分页获取错误日志记录
// Author [yourname](https://github.com/yourname)
func (sysErrorService *SysErrorService) GetSysErrorInfoList(ctx context.Context, info systemReq.SysErrorSearch) (list []system.SysError, total int64, err error) {
	limit, offset := info.LimitOffset()
	// 创建db
	db := global.GVA_DB.WithContext(ctx).Model(&system.SysError{}).Order("created_at desc")
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

	// 异步协程生成解决方案后更新状态。
	// goroutine 存活期远超请求(LLM 调用可达数分钟),而请求 ctx 在 handler 返回后
	// 即被取消,直接用 ctx 会让协程内查询/收尾更新全部 context.Canceled、任务永远
	// 卡在"处理中"。WithoutCancel 脱离请求取消生命周期,同时保留 ctx 里的链路字段
	// (request_id/trace_id 流入本协程的 SQL 日志与出站 LLM 调用)。
	bgCtx := context.WithoutCancel(ctx)
	go func(id string) {
		// 查询当前错误信息用于生成方案
		var se system.SysError
		_ = global.GVA_DB.WithContext(bgCtx).Model(&system.SysError{}).Where("id = ?", id).First(&se).Error

		// 构造 LLM 请求参数，使用管家模式(butler)根据错误信息生成解决方案
		var form, info string
		if se.Form != nil {
			form = *se.Form
		}
		if se.Info != nil {
			info = *se.Info
		}

		llmReq := common.JSONMap{
			"mode": "solution",
			"info": info,
			"form": form,
		}

		// 调用服务层 LLMAuto，忽略错误但尽量写入方案
		var solution string
		if data, err := (&AutoCodeService{}).LLMAuto(bgCtx, llmReq); err == nil {
			solution = fmt.Sprintf("%v", data.(map[string]interface{})["text"])
			_ = global.GVA_DB.WithContext(bgCtx).Model(&system.SysError{}).Where("id = ?", id).Updates(map[string]interface{}{"status": "处理完成", "solution": solution}).Error
		} else {
			// 即使生成失败也标记为完成，避免任务卡住
			_ = global.GVA_DB.WithContext(bgCtx).Model(&system.SysError{}).Where("id = ?", id).Update("status", "处理失败").Error
		}
	}(ID)

	return nil
}
