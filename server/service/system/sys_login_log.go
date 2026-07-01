package system

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type LoginLogService struct{}

var LoginLogServiceApp = new(LoginLogService)

func (loginLogService *LoginLogService) CreateLoginLog(ctx context.Context, loginLog system.SysLoginLog) (err error) {
	// 系统未初始化(GVA_DB==nil)时 /base/login 仍可被未认证请求触达
	// 此处静默跳过登录日志 避免对 nil *gorm.DB 解引用 panic(与 CreateSysError 守卫一致)
	if global.GVA_DB == nil {
		return nil
	}
	err = global.GVA_DB.WithContext(ctx).Create(&loginLog).Error
	return err
}

func (loginLogService *LoginLogService) DeleteLoginLogByIds(ctx context.Context, ids request.IdsReq) (err error) {
	err = global.GVA_DB.WithContext(ctx).Delete(&[]system.SysLoginLog{}, "id in (?)", ids.Ids).Error
	return err
}

func (loginLogService *LoginLogService) DeleteLoginLog(ctx context.Context, loginLog system.SysLoginLog) (err error) {
	err = global.GVA_DB.WithContext(ctx).Delete(&loginLog).Error
	return err
}

func (loginLogService *LoginLogService) GetLoginLog(ctx context.Context, id uint) (loginLog system.SysLoginLog, err error) {
	err = global.GVA_DB.WithContext(ctx).Where("id = ?", id).First(&loginLog).Error
	return
}

func (loginLogService *LoginLogService) GetLoginLogInfoList(ctx context.Context, info systemReq.SysLoginLogSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.WithContext(ctx).Model(&system.SysLoginLog{})
	var loginLogs []system.SysLoginLog
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Username != "" {
		db = db.Where("username LIKE ?", "%"+info.Username+"%")
	}
	if info.Status != false {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("id desc").Preload("User").Find(&loginLogs).Error
	return loginLogs, total, err
}
