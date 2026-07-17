package system

import (
	"context"
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"gorm.io/gorm"
)

type SysPositionService struct{}

var SysPositionServiceApp = new(SysPositionService)

// CreateSysPosition 创建岗位
func (s *SysPositionService) CreateSysPosition(ctx context.Context, pos *system.SysPosition) (err error) {
	return global.GVA_DB.WithContext(ctx).Create(pos).Error
}

// UpdateSysPosition 更新岗位
func (s *SysPositionService) UpdateSysPosition(ctx context.Context, pos *system.SysPosition) (err error) {
	if pos.ID == 0 {
		return errors.New("岗位ID不能为空")
	}
	return global.GVA_DB.WithContext(ctx).Model(&system.SysPosition{}).Where("id = ?", pos.ID).Updates(map[string]interface{}{
		"name":   pos.Name,
		"code":   pos.Code,
		"sort":   pos.Sort,
		"status": pos.Status,
		"remark": pos.Remark,
	}).Error
}

// DeleteSysPosition 删除岗位, 已有用户绑定时禁止删除
func (s *SysPositionService) DeleteSysPosition(ctx context.Context, id uint) (err error) {
	if id == 0 {
		return errors.New("岗位ID不能为空")
	}
	var joinCount int64
	if err = global.GVA_DB.WithContext(ctx).Model(&system.SysUserPosition{}).Where("sys_position_id = ?", id).Count(&joinCount).Error; err != nil {
		return err
	}
	if joinCount > 0 {
		return errors.New("该岗位下存在用户,不允许删除")
	}
	return global.GVA_DB.WithContext(ctx).Delete(&system.SysPosition{}, id).Error
}

// GetSysPosition 获取单个岗位
func (s *SysPositionService) GetSysPosition(ctx context.Context, id uint) (pos system.SysPosition, err error) {
	err = global.GVA_DB.WithContext(ctx).First(&pos, id).Error
	return
}

// GetSysPositionList 分页获取岗位列表
func (s *SysPositionService) GetSysPositionList(ctx context.Context, info systemReq.SysPositionSearch) (list []system.SysPosition, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.WithContext(ctx).Model(&system.SysPosition{})
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Code != "" {
		db = db.Where("code LIKE ?", "%"+info.Code+"%")
	}
	if info.Status != nil {
		db = db.Where("status = ?", *info.Status)
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}
	err = db.Order("sort").Limit(limit).Offset(offset).Find(&list).Error
	return list, total, err
}

// GetPositionUserIds 获取岗位下的用户ID(反向查询)
func (s *SysPositionService) GetPositionUserIds(ctx context.Context, positionId uint) (ids []uint, err error) {
	err = global.GVA_DB.WithContext(ctx).Model(&system.SysUserPosition{}).
		Where("sys_position_id = ?", positionId).Pluck("sys_user_id", &ids).Error
	return
}

// SetPositionUsers 全量覆盖岗位的用户成员(反向分配)
func (s *SysPositionService) SetPositionUsers(ctx context.Context, positionId uint, userIds []uint) error {
	if positionId == 0 {
		return errors.New("岗位ID不能为空")
	}
	return global.GVA_DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&[]system.SysUserPosition{}, "sys_position_id = ?", positionId).Error; err != nil {
			return err
		}
		if len(userIds) > 0 {
			records := make([]system.SysUserPosition, 0, len(userIds))
			for _, uid := range userIds {
				records = append(records, system.SysUserPosition{SysUserId: uid, SysPositionId: positionId})
			}
			if err := tx.Create(&records).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
