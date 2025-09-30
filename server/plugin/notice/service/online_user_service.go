package service

import (
	"github.com/pkg/errors"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/notice/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/notice/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/notice/model/response"
	"gorm.io/gorm"
)

type OnlineUserService struct{}

// AddOnlineUser 添加在线用户
func (s *OnlineUserService) AddOnlineUser(userId uint, socketId string) error {
	// 使用事务确保操作的原子性
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 先硬删除该用户的所有在线记录（包括软删除的记录）
		if err := tx.Unscoped().Where("user_id = ?", userId).Delete(&model.OnlineUser{}).Error; err != nil {
			return err
		}

		// 创建新的在线记录
		onlineUser := model.OnlineUser{
			UserId:         userId,
			SocketId:       socketId,
			LastActiveTime: time.Now(),
			Status:         "online",
		}

		return tx.Create(&onlineUser).Error
	})
}

// RemoveOnlineUser 移除在线用户
func (s *OnlineUserService) RemoveOnlineUser(userId uint) error {
	// 使用硬删除确保用户立即从在线列表中移除
	err := global.GVA_DB.Unscoped().Where("user_id = ?", userId).Delete(&model.OnlineUser{}).Error
	if err != nil {
		return errors.Wrap(err, "移除在线用户失败")
	}
	// 广播给所有在线用户
	return ServiceGroupApp.SocketService.SendToUser(userId, "forceLogout", "您已被管理员强制下线，请重新登录")
}

// RemoveOnlineUserBySocketId 根据SocketId移除在线用户
func (s *OnlineUserService) RemoveOnlineUserBySocketId(socketId string) error {
	// 使用硬删除确保用户立即从在线列表中移除
	return global.GVA_DB.Unscoped().Where("socket_id = ?", socketId).Delete(&model.OnlineUser{}).Error
}

// UpdateUserActivity 更新用户活跃时间
func (s *OnlineUserService) UpdateUserActivity(userId uint) error {
	return global.GVA_DB.Model(&model.OnlineUser{}).
		Where("user_id = ?", userId).
		Update("last_active_time", time.Now()).Error
}

// UpdateUserActivityBySocketId 根据SocketId更新用户活跃时间
func (s *OnlineUserService) UpdateUserActivityBySocketId(socketId string) error {
	return global.GVA_DB.Model(&model.OnlineUser{}).
		Where("socket_id = ?", socketId).
		Update("last_active_time", time.Now()).Error
}

// GetOnlineUsers 获取在线用户列表
func (s *OnlineUserService) GetOnlineUsers() ([]response.OnlineUserResponse, error) {
	var results []struct {
		UserId         uint      `json:"user_id"`
		Username       string    `json:"username"`
		NickName       string    `json:"nick_name"`
		SocketId       string    `json:"socket_id"`
		LastActiveTime time.Time `json:"last_active_time"`
		Status         string    `json:"status"`
	}

	err := global.GVA_DB.Table("notice_online_users ou").
		Select("ou.user_id, u.username, u.nick_name, ou.socket_id, ou.last_active_time, ou.status").
		Joins("JOIN sys_users u ON ou.user_id = u.id").
		Where("ou.status = 'online'").
		Order("ou.last_active_time DESC").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	var onlineUsers []response.OnlineUserResponse
	for _, result := range results {
		onlineUsers = append(onlineUsers, response.OnlineUserResponse{
			UserId:         result.UserId,
			Username:       result.Username,
			NickName:       result.NickName,
			SocketId:       result.SocketId,
			LastActiveTime: result.LastActiveTime,
			Status:         result.Status,
		})
	}

	return onlineUsers, nil
}

// GetOnlineUsersWithPagination 获取在线用户列表（分页）
func (s *OnlineUserService) GetOnlineUsersWithPagination(req request.OnlineUserSearch) ([]response.OnlineUserResponse, int64, error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)

	// 构建查询条件
	db := global.GVA_DB.Table("notice_online_users ou").
		Select("ou.user_id, u.username, u.nick_name, ou.socket_id, ou.last_active_time, ou.status").
		Joins("JOIN sys_users u ON ou.user_id = u.id").
		Where("ou.status = 'online'")

	// 添加搜索条件
	if req.Username != "" {
		db = db.Where("u.username LIKE ?", "%"+req.Username+"%")
	}
	if req.NickName != "" {
		db = db.Where("u.nick_name LIKE ?", "%"+req.NickName+"%")
	}
	if req.RoleId != 0 {
		// 通过角色过滤用户
		db = db.Where("EXISTS (SELECT 1 FROM sys_user_authority sua WHERE sua.sys_user_id = u.id AND sua.sys_authority_authority_id = ?)", req.RoleId)
	}

	// 获取总数
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	var results []struct {
		UserId         uint      `json:"user_id"`
		Username       string    `json:"username"`
		NickName       string    `json:"nick_name"`
		SocketId       string    `json:"socket_id"`
		LastActiveTime time.Time `json:"last_active_time"`
		Status         string    `json:"status"`
	}

	err = db.Order("ou.last_active_time DESC").
		Limit(limit).
		Offset(offset).
		Scan(&results).Error

	if err != nil {
		return nil, 0, err
	}

	var onlineUsers []response.OnlineUserResponse
	for _, result := range results {
		// 计算在线时长
		onlineDuration := time.Since(result.LastActiveTime)
		onlineUsers = append(onlineUsers, response.OnlineUserResponse{
			UserId:         result.UserId,
			Username:       result.Username,
			NickName:       result.NickName,
			SocketId:       result.SocketId,
			LastActiveTime: result.LastActiveTime,
			Status:         result.Status,
			OnlineDuration: int64(onlineDuration.Minutes()), // 在线时长（分钟）
		})
	}

	return onlineUsers, total, nil
}

// GetOnlineUserByUserId 根据用户ID获取在线用户信息
func (s *OnlineUserService) GetOnlineUserByUserId(userId uint) (model.OnlineUser, error) {
	var onlineUser model.OnlineUser
	err := global.GVA_DB.Where("user_id = ? AND status = 'online'", userId).First(&onlineUser).Error
	return onlineUser, err
}

// GetOnlineUserBySocketId 根据SocketId获取在线用户信息
func (s *OnlineUserService) GetOnlineUserBySocketId(socketId string) (model.OnlineUser, error) {
	var onlineUser model.OnlineUser
	err := global.GVA_DB.Where("socket_id = ? AND status = 'online'", socketId).First(&onlineUser).Error
	return onlineUser, err
}

// GetUserSocketId 获取用户的SocketId
func (s *OnlineUserService) GetUserSocketId(userId uint) (string, error) {
	var onlineUser model.OnlineUser
	err := global.GVA_DB.Select("socket_id").Where("user_id = ? AND status = 'online'", userId).First(&onlineUser).Error
	if err != nil {
		return "", err
	}
	return onlineUser.SocketId, nil
}

// GetUsersByRoleIds 根据角色ID获取在线用户列表
func (s *OnlineUserService) GetUsersByRoleIds(roleIds []uint) ([]uint, error) {
	var userIds []uint

	err := global.GVA_DB.Table("notice_online_users ou").
		Select("DISTINCT ou.user_id").
		Joins("JOIN sys_user_authority sua ON ou.user_id = sua.sys_user_id").
		Where("sua.sys_authority_id IN ? AND ou.status = 'online'", roleIds).
		Pluck("user_id", &userIds).Error

	return userIds, err
}

// IsUserOnline 检查用户是否在线
func (s *OnlineUserService) IsUserOnline(userId uint) bool {
	var count int64
	global.GVA_DB.Model(&model.OnlineUser{}).Where("user_id = ? AND status = 'online'", userId).Count(&count)
	return count > 0
}

// GetOnlineUserCount 获取在线用户数量
func (s *OnlineUserService) GetOnlineUserCount() (int64, error) {
	var count int64
	err := global.GVA_DB.Model(&model.OnlineUser{}).Where("status = 'online'").Count(&count).Error
	return count, err
}

// CleanOfflineUsers 清理离线用户（超过指定时间未活跃的用户）
func (s *OnlineUserService) CleanOfflineUsers(timeout time.Duration) error {
	cutoffTime := time.Now().Add(-timeout)
	// 使用硬删除确保用户立即从在线列表中移除
	return global.GVA_DB.Unscoped().Where("last_active_time < ?", cutoffTime).Delete(&model.OnlineUser{}).Error
}

// GetOnlineUserStats 获取在线用户统计数据
func (s *OnlineUserService) GetOnlineUserStats() (response.OnlineUserStatsResponse, error) {
	var stats response.OnlineUserStatsResponse

	// 获取当前在线总数
	err := global.GVA_DB.Model(&model.OnlineUser{}).Where("status = 'online'").Count(&stats.TotalOnline).Error
	if err != nil {
		return stats, err
	}

	// 获取今日登录人数（基于最后活跃时间在今天的用户）
	today := time.Now().Format("2006-01-02")
	todayStart := today + " 00:00:00"
	todayEnd := today + " 23:59:59"

	err = global.GVA_DB.Model(&model.OnlineUser{}).
		Where("last_active_time >= ? AND last_active_time <= ?", todayStart, todayEnd).
		Count(&stats.TodayLogin).Error
	if err != nil {
		return stats, err
	}

	// 峰值在线人数（这里简化处理，可以根据实际需求存储历史峰值数据）
	// 暂时使用当前在线数作为峰值
	stats.PeakOnline = stats.TotalOnline

	// 平均在线人数（这里简化处理，实际应该基于历史数据计算）
	// 暂时使用当前在线数的80%作为平均值
	stats.AverageOnline = int64(float64(stats.TotalOnline) * 0.8)

	return stats, nil
}
