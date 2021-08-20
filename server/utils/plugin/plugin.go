package plugin

import (
	"plugin"
	"sync"

	"github.com/gin-gonic/gin"
)

const (
	OnlyFuncName = "Plugin"
)

var ManagementPlugin = managementPlugin{mp: make(map[string]*plugin.Plugin)}

type managementPlugin struct {
	mp map[string]*plugin.Plugin
	sync.Mutex
}

func (m *managementPlugin) SetPlugin(key string, p *plugin.Plugin) {
	m.Lock()
	defer m.Unlock()
	m.mp[key] = p
}

func (m *managementPlugin) GetPlugin(key string) (p *plugin.Plugin, ok bool) {
	m.Lock()
	defer m.Unlock()
	p, ok = m.mp[key]
	return
}

// Plugin 插件模式接口化
type Plugin interface {
	// Register 注册路由
	Register(group *gin.RouterGroup)

	// RouterPath 用户返回注册路由
	RouterPath() string
}
