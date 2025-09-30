package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
	"strconv"
	"time"

	"github.com/doquangtan/socketio/v4"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	noticeService "github.com/flipped-aurora/gin-vue-admin/server/plugin/notice/service"
	"github.com/gin-gonic/gin"
)

type SocketRouter struct{}

var (
	onlineUserService = noticeService.ServiceGroupApp.OnlineUserService
	socketService     = noticeService.ServiceGroupApp.SocketService
)

// InitSocketRouter 初始化Socket.IO路由
func (r *SocketRouter) InitSocketRouter(Router *gin.RouterGroup) {
	// 创建Socket.IO服务器
	io := socketio.New()
	var claims *request.CustomClaims
	// 认证中间件
	io.OnAuthentication(func(params map[string]string) bool {
		// 从参数获取token
		token, ok := params["token"]
		if !ok || token == "" {
			return false
		}

		j := utils.NewJWT()
		// parseToken 解析token包含的信息
		claimsP, err := j.ParseToken(token)
		if err != nil {
			zap.L().Error("JWT解析失败: " + err.Error())
			return false
		}
		claims = claimsP

		return true
	})

	// 连接事件
	io.OnConnection(func(socket *socketio.Socket) {
		if claims == nil {
			global.GVA_LOG.Error("用户认证失败，无法建立连接")
			socket.Disconnect()
			return
		}
		userInfo := claims
		userID := userInfo.BaseClaims.ID
		socketID := socket.Id

		// 加入用户房间
		userRoom := "user_" + strconv.Itoa(int(userID))
		socket.Join(userRoom)
		// 添加到在线用户列表
		err := onlineUserService.AddOnlineUser(userID, socketID)
		if err != nil {
			global.GVA_LOG.Error("添加在线用户失败: " + err.Error())
		}
		// 添加到Socket连接管理
		socketService.AddConnection(userID, socketID, socket)

		// 记录用户连接信息
		global.GVA_LOG.Info("用户连接详情: UserID=" + strconv.Itoa(int(userID)) +
			", SocketID=" + socketID)

		// 发送用户上线通知
		socket.Emit("user_online", map[string]interface{}{
			"userId":    userID,
			"socketId":  socketID,
			"timestamp": time.Now(),
			"message":   "连接成功",
		})

		global.GVA_LOG.Info("用户连接成功: UserID=" + strconv.Itoa(int(userID)) + ", SocketID=" + socketID)

		// 监听心跳事件
		socket.On("heartbeat", func(event *socketio.EventPayload) {
			// 更新用户活跃时间
			err := onlineUserService.UpdateUserActivity(userID)
			if err != nil {
				global.GVA_LOG.Error("更新用户活跃时间失败: " + err.Error())
			}

			// 回复心跳
			socket.Emit("heartbeat_reply", map[string]interface{}{
				"timestamp": time.Now(),
				"status":    "ok",
			})
		})

		// 监听通知已读事件
		socket.On("mark_notification_read", func(event *socketio.EventPayload) {
			if event.Data != nil && len(event.Data) > 0 {
				if dataMap, ok := event.Data[0].(map[string]interface{}); ok {
					if notificationIds, exists := dataMap["notificationIds"]; exists {
						if ids, ok := notificationIds.([]interface{}); ok {
							var uintIds []uint
							for _, id := range ids {
								if idFloat, ok := id.(float64); ok {
									uintIds = append(uintIds, uint(idFloat))
								}
							}

							// 标记通知为已读
							notificationService := noticeService.ServiceGroupApp.NotificationService
							err := notificationService.MarkNotificationsAsRead(userID, uintIds)
							if err != nil {
								global.GVA_LOG.Error("标记通知已读失败: " + err.Error())
								socket.Emit("error", map[string]interface{}{
									"message": "标记通知已读失败",
									"error":   err.Error(),
								})
								return
							}

							// 发送成功响应
							socket.Emit("notification_read_success", map[string]interface{}{
								"notificationIds": uintIds,
								"timestamp":       time.Now(),
							})
						}
					}
				}
			}
		})

		// 监听获取通知统计事件
		socket.On("get_notification_stats", func(event *socketio.EventPayload) {
			notificationService := noticeService.ServiceGroupApp.NotificationService
			stats, err := notificationService.GetNotificationStats(userID)
			if err != nil {
				global.GVA_LOG.Error("获取通知统计失败: " + err.Error())
				socket.Emit("error", map[string]interface{}{
					"message": "获取通知统计失败",
					"error":   err.Error(),
				})
				return
			}

			socket.Emit("notification_stats", stats)
		})

		// 监听断开连接事件
		socket.On("disconnect", func(event *socketio.EventPayload) {
			// 从在线用户列表移除
			err := onlineUserService.RemoveOnlineUser(userID)
			if err != nil {
				global.GVA_LOG.Error("移除在线用户失败: " + err.Error())
			}

			// 从Socket连接管理移除
			socketService.RemoveConnectionByUserId(userID)

			global.GVA_LOG.Info("用户断开连接: UserID=" + strconv.Itoa(int(userID)) + ", SocketID=" + socketID)
		})
	})

	// 注册Socket.IO路由
	Router.GET("/socket.io/*any", gin.WrapH(io.HttpHandler()))
	Router.POST("/socket.io/*any", gin.WrapH(io.HttpHandler()))
}
