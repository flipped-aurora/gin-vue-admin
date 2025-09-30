package service

import (
	"log"
	"sync"

	"github.com/doquangtan/socketio/v4"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/notice/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/notice/model/response"
)

type SocketService struct {
	server      *socketio.Socket
	connections map[string]*socketio.Socket // socketId -> socket
	userSockets map[uint]string             // userId -> socketId
	mutex       sync.RWMutex
}

// SetServer 设置Socket.IO服务器实例
func (s *SocketService) SetServer(server *socketio.Socket) {
	s.server = server
}

// AddConnection 添加连接
func (s *SocketService) AddConnection(userId uint, socketId string, socket *socketio.Socket) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.connections == nil {
		s.connections = make(map[string]*socketio.Socket)
	}
	if s.userSockets == nil {
		s.userSockets = make(map[uint]string)
	}

	// 如果用户已有连接，先断开旧连接
	if oldSocketId, exists := s.userSockets[userId]; exists {
		if oldSocket, ok := s.connections[oldSocketId]; ok {
			oldSocket.Disconnect()
			delete(s.connections, oldSocketId)
		}
	}

	s.connections[socketId] = socket
	s.userSockets[userId] = socketId

	log.Printf("用户 %d 连接成功，SocketId: %s", userId, socketId)
}

// RemoveConnection 移除连接
func (s *SocketService) RemoveConnection(socketId string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.connections, socketId)

	// 从用户映射中移除
	for userId, sid := range s.userSockets {
		if sid == socketId {
			delete(s.userSockets, userId)
			log.Printf("用户 %d 断开连接，SocketId: %s", userId, socketId)
			break
		}
	}
}

// RemoveConnectionByUserId 根据用户ID移除连接
func (s *SocketService) RemoveConnectionByUserId(userId uint) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if socketId, exists := s.userSockets[userId]; exists {
		if socket, ok := s.connections[socketId]; ok {
			socket.Disconnect()
			delete(s.connections, socketId)
		}
		delete(s.userSockets, userId)
		log.Printf("用户 %d 连接被移除", userId)
	}
}

// GetConnection 获取连接
func (s *SocketService) GetConnection(socketId string) (*socketio.Socket, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	socket, exists := s.connections[socketId]
	return socket, exists
}

// GetUserConnection 根据用户ID获取连接
func (s *SocketService) GetUserConnection(userId uint) (*socketio.Socket, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if socketId, exists := s.userSockets[userId]; exists {
		if socket, ok := s.connections[socketId]; ok {
			return socket, true
		}
	}
	return nil, false
}

// GetUserSocketId 获取用户的SocketId
func (s *SocketService) GetUserSocketId(userId uint) (string, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	socketId, exists := s.userSockets[userId]
	return socketId, exists
}

// IsUserOnline 检查用户是否在线
func (s *SocketService) IsUserOnline(userId uint) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	_, exists := s.userSockets[userId]
	return exists
}

// GetOnlineUsers 获取在线用户列表
func (s *SocketService) GetOnlineUsers() []uint {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	users := make([]uint, 0, len(s.userSockets))
	for userId := range s.userSockets {
		users = append(users, userId)
	}

	return users
}

// GetOnlineUserCount 获取在线用户数量
func (s *SocketService) GetOnlineUserCount() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return len(s.userSockets)
}

// SendToUser 向指定用户发送消息
func (s *SocketService) SendToUser(userId uint, event string, data interface{}) error {
	log.Printf("尝试向用户 %d 发送消息，事件: %s", userId, event)
	socket, exists := s.GetUserConnection(userId)
	if !exists {
		log.Printf("用户 %d 不在线，无法发送消息", userId)
		return nil // 用户不在线，不发送
	}

	log.Printf("用户 %d 在线，发送消息，事件: %s", userId, event)
	socket.Emit(event, data)
	log.Printf("消息已发送给用户 %d，事件: %s", userId, event)
	return nil
}

// SendToUsers 向多个用户发送消息
func (s *SocketService) SendToUsers(userIds []uint, event string, data interface{}) {
	for _, userId := range userIds {
		if err := s.SendToUser(userId, event, data); err != nil {
			log.Printf("向用户 %d 发送消息失败: %v", userId, err)
		}
	}
}

// BroadcastToAll 向所有在线用户广播消息
func (s *SocketService) BroadcastToAll(event string, data interface{}) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	log.Printf("开始广播消息，事件: %s，在线用户数: %d", event, len(s.userSockets))
	log.Printf("当前在线用户列表: %v", s.userSockets)

	for userId := range s.userSockets {
		log.Printf("向用户 %d 发送消息，事件: %s", userId, event)
		if err := s.SendToUser(userId, event, data); err != nil {
			log.Printf("向用户 %d 广播消息失败: %v", userId, err)
		} else {
			log.Printf("向用户 %d 广播消息成功", userId)
		}
	}
}

// SendNotificationToUser 向用户发送通知
func (s *SocketService) SendNotificationToUser(userId uint, notification *model.Notification) error {
	notificationData := response.NotificationResponse{
		ID:          notification.ID,
		Title:       notification.Title,
		Content:     notification.Content,
		Type:        notification.Type,
		Priority:    notification.Priority,
		SenderId:    notification.SenderId,
		TargetType:  notification.TargetType,
		Status:      notification.Status,
		PublishTime: notification.PublishTime,
		ExpireTime:  notification.ExpireTime,
		CreatedAt:   notification.CreatedAt,
		UpdatedAt:   notification.UpdatedAt,
	}

	return s.SendToUser(userId, "new_notification", notificationData)
}

// SendNotificationToUsers 向多个用户发送通知
func (s *SocketService) SendNotificationToUsers(userIds []uint, notification *model.Notification) {
	notificationData := response.NotificationResponse{
		ID:          notification.ID,
		Title:       notification.Title,
		Content:     notification.Content,
		Type:        notification.Type,
		Priority:    notification.Priority,
		SenderId:    notification.SenderId,
		TargetType:  notification.TargetType,
		Status:      notification.Status,
		PublishTime: notification.PublishTime,
		ExpireTime:  notification.ExpireTime,
		CreatedAt:   notification.CreatedAt,
		UpdatedAt:   notification.UpdatedAt,
	}

	for _, userId := range userIds {
		if err := s.SendToUser(userId, "new_notification", notificationData); err != nil {
			log.Printf("向用户 %d 发送通知失败: %v", userId, err)
		}
	}
}

// BroadcastNotification 广播通知给所有在线用户
func (s *SocketService) BroadcastNotification(notification *model.Notification) {
	notificationData := response.NotificationResponse{
		ID:          notification.ID,
		Title:       notification.Title,
		Content:     notification.Content,
		Type:        notification.Type,
		Priority:    notification.Priority,
		SenderId:    notification.SenderId,
		TargetType:  notification.TargetType,
		Status:      notification.Status,
		PublishTime: notification.PublishTime,
		ExpireTime:  notification.ExpireTime,
		CreatedAt:   notification.CreatedAt,
		UpdatedAt:   notification.UpdatedAt,
	}

	s.BroadcastToAll("new_notification", notificationData)
}

// SendReadStatusUpdate 发送已读状态更新
func (s *SocketService) SendReadStatusUpdate(userId uint, notificationId uint) error {
	data := map[string]interface{}{
		"notification_id": notificationId,
		"read_status":     true,
	}

	return s.SendToUser(userId, "read_status_update", data)
}

// GetConnectionInfo 获取连接信息
func (s *SocketService) GetConnectionInfo() map[string]interface{} {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return map[string]interface{}{
		"total_connections": len(s.connections),
		"online_users":      len(s.userSockets),
		"user_sockets":      s.userSockets,
	}
}
