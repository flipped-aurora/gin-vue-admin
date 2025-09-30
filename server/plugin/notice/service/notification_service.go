package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	systemModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/notice/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/notice/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/notice/model/response"
)

type NotificationService struct{}

// CreateNotification 创建通知
func (s *NotificationService) CreateNotification(req request.CreateNotificationRequest, senderId uint) (notification model.Notification, err error) {
	// 将 []uint 类型的 TargetIds 转换为 JSON 字符串
	targetIdsJson := ""
	if len(req.TargetIds) > 0 {
		targetIdsBytes, err := json.Marshal(req.TargetIds)
		if err != nil {
			return notification, fmt.Errorf("序列化目标用户ID失败: %v", err)
		}
		targetIdsJson = string(targetIdsBytes)
	}

	notification = model.Notification{
		Title:       req.Title,
		Content:     req.Content,
		Type:        req.Type,
		Priority:    req.Priority,
		SenderId:    senderId,
		TargetType:  req.TargetType,
		TargetIds:   targetIdsJson,
		Status:      "draft",
		PublishTime: req.PublishTime,
		ExpireTime:  req.ExpireTime,
	}

	err = global.GVA_DB.Create(&notification).Error
	return notification, err
}

// UpdateNotification 更新通知
func (s *NotificationService) UpdateNotification(req request.UpdateNotificationRequest) error {
	var notification model.Notification
	if err := global.GVA_DB.First(&notification, req.ID).Error; err != nil {
		return err
	}

	// 将 []uint 类型的 TargetIds 转换为 JSON 字符串
	targetIdsJson := ""
	if len(req.TargetIds) > 0 {
		targetIdsBytes, err := json.Marshal(req.TargetIds)
		if err != nil {
			return fmt.Errorf("序列化目标用户ID失败: %v", err)
		}
		targetIdsJson = string(targetIdsBytes)
	}

	updates := map[string]interface{}{
		"title":        req.Title,
		"content":      req.Content,
		"type":         req.Type,
		"priority":     req.Priority,
		"target_type":  req.TargetType,
		"target_ids":   targetIdsJson,
		"status":       req.Status,
		"publish_time": req.PublishTime,
		"expire_time":  req.ExpireTime,
	}

	return global.GVA_DB.Model(&notification).Updates(updates).Error
}

// DeleteNotification 删除通知
func (s *NotificationService) DeleteNotification(id uint) error {
	return global.GVA_DB.Delete(&model.Notification{}, id).Error
}

// GetNotificationList 获取通知列表
func (s *NotificationService) GetNotificationList(req request.NotificationSearch) (list []response.NotificationResponse, total int64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	db := global.GVA_DB.Model(&model.Notification{})

	// 构建查询条件
	if req.Title != "" {
		db = db.Where("title LIKE ?", "%"+req.Title+"%")
	}
	if req.Type != "" {
		db = db.Where("type = ?", req.Type)
	}
	if req.Priority != "" {
		db = db.Where("priority = ?", req.Priority)
	}
	if req.Status != "" {
		db = db.Where("status = ?", req.Status)
	}
	if req.SenderId != 0 {
		db = db.Where("sender_id = ?", req.SenderId)
	}
	if req.TargetType != "" {
		db = db.Where("target_type = ?", req.TargetType)
	}
	if req.StartTime != nil {
		db = db.Where("created_at >= ?", req.StartTime)
	}
	if req.EndTime != nil {
		db = db.Where("created_at <= ?", req.EndTime)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	var notifications []model.Notification
	err = db.Limit(limit).Offset(offset).Order("created_at DESC").Find(&notifications).Error
	if err != nil {
		return
	}

	// 获取发送者信息
	var senderIds []uint
	for _, notification := range notifications {
		senderIds = append(senderIds, notification.SenderId)
	}

	var users []systemModel.SysUser
	userMap := make(map[uint]string)
	if len(senderIds) > 0 {
		global.GVA_DB.Where("id IN ?", senderIds).Find(&users)
		for _, user := range users {
			userMap[user.ID] = user.Username
		}
	}

	// 构建响应数据
	for _, notification := range notifications {
		senderName := userMap[notification.SenderId]
		list = append(list, response.NotificationResponse{
			ID:          notification.ID,
			Title:       notification.Title,
			Content:     notification.Content,
			Type:        notification.Type,
			Priority:    notification.Priority,
			SenderId:    notification.SenderId,
			SenderName:  senderName,
			TargetType:  notification.TargetType,
			TargetIds:   notification.TargetIds,
			Status:      notification.Status,
			PublishTime: notification.PublishTime,
			ExpireTime:  notification.ExpireTime,
			CreatedAt:   notification.CreatedAt,
			UpdatedAt:   notification.UpdatedAt,
		})
	}

	return list, total, err
}

// GetNotificationById 根据ID获取通知
func (s *NotificationService) GetNotificationById(id uint) (notification model.Notification, err error) {
	err = global.GVA_DB.First(&notification, id).Error
	return
}

// PublishNotification 发布通知
func (s *NotificationService) PublishNotification(id uint) error {
	// 获取通知详情
	var notification model.Notification
	if err := global.GVA_DB.First(&notification, id).Error; err != nil {
		return err
	}

	// 检查通知是否已经发布
	if notification.Status == "published" {
		return fmt.Errorf("通知已经发布")
	}

	// 更新通知状态为已发布
	now := time.Now()
	if err := global.GVA_DB.Model(&model.Notification{}).Where("id = ?", id).Updates(map[string]interface{}{
		"status":       "published",
		"publish_time": &now,
	}).Error; err != nil {
		return err
	}

	// 重新获取更新后的通知信息
	if err := global.GVA_DB.First(&notification, id).Error; err != nil {
		return err
	}

	// 根据目标类型创建用户读取记录
	var targetUserIds []uint
	var targetIds []uint
	if notification.TargetIds != "" {
		json.Unmarshal([]byte(notification.TargetIds), &targetIds)
	}

	switch notification.TargetType {
	case "all":
		// 获取所有用户ID
		var users []systemModel.SysUser
		global.GVA_DB.Select("id").Find(&users)
		for _, user := range users {
			targetUserIds = append(targetUserIds, user.ID)
		}
	case "users":
		targetUserIds = targetIds
	case "roles":
		// 根据角色获取用户ID
		if len(targetIds) > 0 {
			var users []systemModel.SysUser
			global.GVA_DB.Joins("JOIN sys_user_authority ON sys_users.id = sys_user_authority.sys_user_id").
				Where("sys_user_authority.sys_authority_id IN ?", targetIds).
				Select("DISTINCT sys_users.id").
				Find(&users)
			for _, user := range users {
				targetUserIds = append(targetUserIds, user.ID)
			}
		}
	}

	// 批量创建用户读取记录
	var userReads []model.UserRead
	for _, userId := range targetUserIds {
		userReads = append(userReads, model.UserRead{
			UserId:         userId,
			NotificationId: notification.ID,
			IsRead:         false,
		})
	}

	if len(userReads) > 0 {
		if err := global.GVA_DB.CreateInBatches(userReads, 100).Error; err != nil {
			return fmt.Errorf("创建用户读取记录失败: %v", err)
		}
	}

	// 通过Socket实时推送通知给在线用户
	socketService := ServiceGroupApp.SocketService
	switch notification.TargetType {
	case "all":
		// 广播给所有在线用户
		socketService.BroadcastNotification(&notification)
	default:
		// 发送给指定的在线用户
		if len(targetUserIds) > 0 {
			socketService.SendNotificationToUsers(targetUserIds, &notification)
		}
	}
	return nil
}

// SendNotification 发送通知
func (s *NotificationService) SendNotification(req request.SendNotificationRequest, senderId uint) (response.SendNotificationResponse, error) {
	var result response.SendNotificationResponse

	// 创建通知记录
	targetIds, _ := json.Marshal(req.TargetIds)
	notification := model.Notification{
		Title:       req.Title,
		Content:     req.Content,
		Type:        req.Type,
		Priority:    req.Priority,
		SenderId:    senderId,
		TargetType:  req.TargetType,
		TargetIds:   string(targetIds),
		Status:      "published",
		PublishTime: &time.Time{},
	}

	now := time.Now()
	notification.PublishTime = &now

	if err := global.GVA_DB.Create(&notification).Error; err != nil {
		return result, err
	}

	result.NotificationId = notification.ID

	// 根据目标类型创建用户读取记录
	var targetUserIds []uint
	switch req.TargetType {
	case "all":
		// 获取所有用户ID
		var users []systemModel.SysUser
		global.GVA_DB.Select("id").Find(&users)
		for _, user := range users {
			targetUserIds = append(targetUserIds, user.ID)
		}
	case "user":
		targetUserIds = req.TargetIds
	case "role":
		// 根据角色获取用户ID
		if len(req.RoleIds) > 0 {
			var users []systemModel.SysUser
			global.GVA_DB.Joins("JOIN sys_user_authority ON sys_users.id = sys_user_authority.sys_user_id").
				Where("sys_user_authority.sys_authority_id IN ?", req.RoleIds).
				Select("DISTINCT sys_users.id").
				Find(&users)
			for _, user := range users {
				targetUserIds = append(targetUserIds, user.ID)
			}
		}
	}

	// 批量创建用户读取记录
	var userReads []model.UserRead
	for _, userId := range targetUserIds {
		userReads = append(userReads, model.UserRead{
			UserId:         userId,
			NotificationId: notification.ID,
			IsRead:         false,
		})
	}

	if len(userReads) > 0 {
		if err := global.GVA_DB.CreateInBatches(userReads, 100).Error; err != nil {
			result.FailCount = len(userReads)
			result.Message = "创建用户读取记录失败: " + err.Error()
			return result, err
		}
		result.SuccessCount = len(userReads)
	}

	result.Message = fmt.Sprintf("通知发送成功，共发送给 %d 个用户", result.SuccessCount)
	return result, nil
}

// GetUserNotifications 获取用户通知列表
func (s *NotificationService) GetUserNotifications(userId uint, req request.UserNotificationSearch) (list []response.UserNotificationResponse, total int64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)

	db := global.GVA_DB.Table("notice_user_reads ur").
		Select("n.id, n.title, n.content, n.type, n.priority, n.sender_id, n.publish_time, n.expire_time, n.created_at, ur.is_read, ur.read_time").
		Joins("JOIN notice_notifications n ON ur.notification_id = n.id").
		Where("ur.user_id = ? AND n.status = 'published'", userId)

	// 构建查询条件
	if req.IsRead != nil {
		db = db.Where("ur.is_read = ?", *req.IsRead)
	}
	if req.Type != "" {
		db = db.Where("n.type = ?", req.Type)
	}
	if req.Priority != "" {
		db = db.Where("n.priority = ?", req.Priority)
	}
	if req.StartTime != nil {
		db = db.Where("n.created_at >= ?", req.StartTime)
	}
	if req.EndTime != nil {
		db = db.Where("n.created_at <= ?", req.EndTime)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	var results []struct {
		ID          uint       `json:"id"`
		Title       string     `json:"title"`
		Content     string     `json:"content"`
		Type        string     `json:"type"`
		Priority    string     `json:"priority"`
		SenderId    uint       `json:"sender_id"`
		PublishTime *time.Time `json:"publish_time"`
		ExpireTime  *time.Time `json:"expire_time"`
		CreatedAt   time.Time  `json:"created_at"`
		IsRead      bool       `json:"is_read"`
		ReadTime    *time.Time `json:"read_time"`
	}

	err = db.Limit(limit).Offset(offset).Order("n.created_at DESC").Scan(&results).Error
	if err != nil {
		return
	}

	// 获取发送者信息
	var senderIds []uint
	for _, result := range results {
		senderIds = append(senderIds, result.SenderId)
	}

	var users []systemModel.SysUser
	userMap := make(map[uint]string)
	if len(senderIds) > 0 {
		global.GVA_DB.Where("id IN ?", senderIds).Find(&users)
		for _, user := range users {
			userMap[user.ID] = user.Username
		}
	}

	// 构建响应数据
	for _, result := range results {
		senderName := userMap[result.SenderId]
		list = append(list, response.UserNotificationResponse{
			ID:          result.ID,
			Title:       result.Title,
			Content:     result.Content,
			Type:        result.Type,
			Priority:    result.Priority,
			SenderId:    result.SenderId,
			SenderName:  senderName,
			IsRead:      result.IsRead,
			ReadTime:    result.ReadTime,
			PublishTime: result.PublishTime,
			ExpireTime:  result.ExpireTime,
			CreatedAt:   result.CreatedAt,
		})
	}

	return list, total, err
}

// MarkNotificationsAsRead 标记通知为已读
func (s *NotificationService) MarkNotificationsAsRead(userId uint, notificationIds []uint) error {
	now := time.Now()
	return global.GVA_DB.Model(&model.UserRead{}).
		Where("user_id = ? AND notification_id IN ?", userId, notificationIds).
		Updates(map[string]interface{}{
			"is_read":   true,
			"read_time": &now,
		}).Error
}

// GetNotificationStats 获取用户通知统计
func (s *NotificationService) GetNotificationStats(userId uint) (response.NotificationStatsResponse, error) {
	var stats response.NotificationStatsResponse

	// 总通知数
	err := global.GVA_DB.Table("notice_user_reads ur").
		Joins("JOIN notice_notifications n ON ur.notification_id = n.id").
		Where("ur.user_id = ? AND n.status = 'published'", userId).
		Count(&stats.TotalCount).Error
	if err != nil {
		return stats, err
	}

	// 未读通知数
	err = global.GVA_DB.Table("notice_user_reads ur").
		Joins("JOIN notice_notifications n ON ur.notification_id = n.id").
		Where("ur.user_id = ? AND ur.is_read = false AND n.status = 'published'", userId).
		Count(&stats.UnreadCount).Error
	if err != nil {
		return stats, err
	}

	stats.ReadCount = stats.TotalCount - stats.UnreadCount
	return stats, nil
}

// DeleteUserNotification 删除用户通知记录
// 如果通知只针对单个用户，则删除通知本身；否则只删除用户读取记录
func (s *NotificationService) DeleteUserNotification(userId uint, notificationId uint) error {
	// 首先检查通知是否存在
	var notification model.Notification
	err := global.GVA_DB.First(&notification, notificationId).Error
	if err != nil {
		return fmt.Errorf("通知不存在")
	}

	// 检查通知的接收者数量
	var receiverCount int64
	err = global.GVA_DB.Model(&model.UserRead{}).
		Where("notification_id = ?", notificationId).
		Count(&receiverCount).Error
	if err != nil {
		return err
	}

	// 如果只有一个接收者，删除通知本身
	if receiverCount == 1 {
		// 开启事务
		tx := global.GVA_DB.Begin()
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()

		// 删除用户读取记录
		if err := tx.Where("notification_id = ? AND user_id = ?", notificationId, userId).Delete(&model.UserRead{}).Error; err != nil {
			tx.Rollback()
			return err
		}

		// 删除通知本身
		if err := tx.Delete(&notification).Error; err != nil {
			tx.Rollback()
			return err
		}

		return tx.Commit().Error
	} else {
		// 如果有多个接收者，只删除当前用户的读取记录
		return global.GVA_DB.Where("notification_id = ? AND user_id = ?", notificationId, userId).Delete(&model.UserRead{}).Error
	}
}
