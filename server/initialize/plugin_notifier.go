package initialize

import (
	"sync"

	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/gin-gonic/gin"
)

func init() {
	// 注入数据库就绪回调到 service 层
	system.SetDBReadyCallback(func() {
		GetDBReadyNotifier().NotifyDBReady()
	})
}

var (
	dbReadyNotifier     *DBReadyNotifier
	dbReadyNotifierOnce sync.Once
)

// DBReadyNotifier 数据库就绪通知器
type DBReadyNotifier struct {
	mu           sync.RWMutex
	subscribers  []func()
	dbReady      bool
	pluginParams *PluginParams // 保存插件初始化所需的参数
}

// PluginParams 插件初始化参数
type PluginParams struct {
	PrivateGroup *gin.RouterGroup
	PublicRouter *gin.RouterGroup
	Engine       *gin.Engine
}

// GetDBReadyNotifier 获取全局通知器实例（单例）
func GetDBReadyNotifier() *DBReadyNotifier {
	dbReadyNotifierOnce.Do(func() {
		dbReadyNotifier = &DBReadyNotifier{
			subscribers: make([]func(), 0),
			dbReady:     false,
		}
	})
	return dbReadyNotifier
}

// Subscribe 订阅数据库就绪事件
// 如果数据库已经就绪，立即执行回调
// 否则将回调加入订阅列表，等待通知
func (n *DBReadyNotifier) Subscribe(callback func()) {
	n.mu.Lock()
	defer n.mu.Unlock()

	if n.dbReady {
		// 数据库已就绪，立即执行
		callback()
	} else {
		// 数据库未就绪，加入订阅列表
		n.subscribers = append(n.subscribers, callback)
	}
}

// SetPluginParams 设置插件初始化参数
func (n *DBReadyNotifier) SetPluginParams(params *PluginParams) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.pluginParams = params
}

// GetPluginParams 获取插件初始化参数
func (n *DBReadyNotifier) GetPluginParams() *PluginParams {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.pluginParams
}

// NotifyDBReady 通知数据库已就绪
// 会触发所有订阅的回调函数
func (n *DBReadyNotifier) NotifyDBReady() {
	n.mu.Lock()

	if n.dbReady {
		// 已经通知过了，避免重复通知
		n.mu.Unlock()
		return
	}

	n.dbReady = true

	// 复制订阅列表，避免在执行回调时持有锁
	callbacks := make([]func(), len(n.subscribers))
	copy(callbacks, n.subscribers)

	// 清空订阅列表
	n.subscribers = nil

	// 释放锁后再执行回调，避免死锁
	n.mu.Unlock()

	// 执行所有订阅的回调
	for _, callback := range callbacks {
		callback()
	}
}

// IsDBReady 检查数据库是否已就绪
func (n *DBReadyNotifier) IsDBReady() bool {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.dbReady
}
