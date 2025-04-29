package utils

import (
	"sync"
)

// SystemEvents 定义系统级事件处理
type SystemEvents struct {
	reloadHandlers []func() error
	mu             sync.RWMutex
}

// 全局事件管理器
var GlobalSystemEvents = &SystemEvents{}

// RegisterReloadHandler 注册系统重载处理函数
func (e *SystemEvents) RegisterReloadHandler(handler func() error) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.reloadHandlers = append(e.reloadHandlers, handler)
}

// TriggerReload 触发所有注册的重载处理函数
func (e *SystemEvents) TriggerReload() error {
	e.mu.RLock()
	defer e.mu.RUnlock()
	
	for _, handler := range e.reloadHandlers {
		if err := handler(); err != nil {
			return err
		}
	}
	return nil
}
