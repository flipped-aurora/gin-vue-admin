import { io } from 'socket.io-client'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'
import { ElNotification } from 'element-plus'
import { useRouter } from 'vue-router'
const router = useRouter()

class SocketManager {
  constructor() {
    this.socket = null
    this.connected = false
    this.reconnectAttempts = 0
    this.maxReconnectAttempts = 5
    this.reconnectInterval = 3000
    this.heartbeatInterval = null
    this.callbacks = new Map()
  }

  // 连接Socket.IO服务器
  connect() {
    if (this.socket && this.connected) {
      return Promise.resolve()
    }

    const userStore = useUserStore()
    const token = userStore.token

    if (!token) {
      console.warn('No token found, cannot connect to socket')
      return Promise.reject(new Error('No token'))
    }

    return new Promise((resolve, reject) => {
      this.socket = io('http://127.0.0.1:8888/', {
        auth: {
          token: token
        },
        transports: ['websocket', 'polling'],
        timeout: 10000,
        forceNew: true
      })

      // 连接成功
      this.socket.on('connect', () => {
        console.log('Socket connected:', this.socket.id)
        this.connected = true
        this.reconnectAttempts = 0
        this.startHeartbeat()
        resolve()
      })

      // 连接失败
      this.socket.on('connect_error', (error) => {
        console.error('Socket connection error:', error)
        this.connected = false
        reject(error)
      })

      // 断开连接
      this.socket.on('disconnect', (reason) => {
        console.log('Socket disconnected:', reason)
        this.connected = false
        this.stopHeartbeat()

        // 如果是服务器主动断开，尝试重连
        if (reason === 'io server disconnect') {
          this.reconnect()
        }
      })

      // 认证失败
      this.socket.on('auth_error', (error) => {
        console.error('Socket auth error:', error)
        ElMessage.error('Socket认证失败，请重新登录')
        this.disconnect()
      })

      // 新通知
      this.socket.on('new_notification', (data) => {
        console.log('Received new notification:', data)
        this.emit('new_notification', data)
      })

      // 通知统计更新
      this.socket.on('notification_stats', (data) => {
        console.log('Notification stats updated:', data)
        this.emit('notification_stats', data)
      })

      // 用户被踢下线
      this.socket.on('forceLogout', (data) => {
        console.log('forceLogout:', data)
        this.emit('forceLogout', data)
      })
    })
  }

  // 断开连接
  disconnect() {
    if (this.socket) {
      this.stopHeartbeat()
      this.socket.disconnect()
      this.socket = null
      this.connected = false
    }
  }

  // 重连
  reconnect() {
    if (this.reconnectAttempts >= this.maxReconnectAttempts) {
      console.error('Max reconnect attempts reached')
      return
    }

    this.reconnectAttempts++
    console.log(
      `Attempting to reconnect (${this.reconnectAttempts}/${this.maxReconnectAttempts})`
    )

    setTimeout(() => {
      this.connect().catch(() => {
        this.reconnect()
      })
    }, this.reconnectInterval)
  }

  // 开始心跳
  startHeartbeat() {
    this.heartbeatInterval = setInterval(() => {
      if (this.socket && this.connected) {
        this.socket.emit('heartbeat', { timestamp: Date.now() })
      }
    }, 30000) // 30秒心跳
  }

  // 停止心跳
  stopHeartbeat() {
    if (this.heartbeatInterval) {
      clearInterval(this.heartbeatInterval)
      this.heartbeatInterval = null
    }
  }

  // 加入房间
  joinRoom(roomId) {
    if (this.socket && this.connected) {
      this.socket.emit('join_room', { roomId })
    }
  }

  // 离开房间
  leaveRoom(roomId) {
    if (this.socket && this.connected) {
      this.socket.emit('leave_room', { roomId })
    }
  }

  // 标记通知为已读
  markNotificationAsRead(notificationId) {
    if (this.socket && this.connected) {
      this.socket.emit('notification_read', { notificationId })
    }
  }

  // 获取通知统计
  getNotificationStats() {
    if (this.socket && this.connected) {
      this.socket.emit('get_notification_stats')
    }
  }

  // 监听事件
  on(event, callback) {
    if (!this.callbacks.has(event)) {
      this.callbacks.set(event, [])
    }
    this.callbacks.get(event).push(callback)
  }

  // 移除事件监听
  off(event, callback) {
    if (this.callbacks.has(event)) {
      const callbacks = this.callbacks.get(event)
      const index = callbacks.indexOf(callback)
      if (index > -1) {
        callbacks.splice(index, 1)
      }
    }
  }

  // 触发事件
  emit(event, data) {
    if (this.callbacks.has(event)) {
      this.callbacks.get(event).forEach((callback) => {
        try {
          callback(data)
        } catch (error) {
          console.error('Error in socket callback:', error)
        }
      })
    }
  }

  // 获取连接状态
  isConnected() {
    return this.connected
  }

  // 获取Socket ID
  getSocketId() {
    return this.socket ? this.socket.id : null
  }
}

// 创建全局Socket管理器实例
export const socketManager = new SocketManager()

export const initSocketConnection = (userStore) => {
  if (userStore.token && userStore.userInfo.ID) {
    // 监听新通知事件
    socketManager.on('new_notification', (notification) => {
      ElNotification({
        title: '新通知',
        message: notification.title,
        type: 'info',
        duration: 5000,
        onClick: () => {
          // 点击通知跳转到通知详情或通知列表
          router.push({ name: 'NotificationManagement' })
        }
      })
    })

    // 监听通知统计更新
    socketManager.on('notificationStatsUpdate', (stats) => {
      // 可以在这里更新全局的未读通知数量
      console.log('通知统计更新:', stats)
    })

    // 监听被踢下线事件
    socketManager.on('forceLogout', (data) => {
      console.log('收到被踢下线通知:', data)
      ElNotification({
        title: '系统提示',
        message: data.message || '您已被管理员强制下线',
        type: 'warning',
        duration: 0
      })

      // 清除用户信息并跳转到登录页
      setTimeout(() => {
        userStore.LoginOut()
        router.push({ name: 'Login' })
      }, 2000)
    })

    // 连接到服务器
    socketManager.connect()
  }
}
