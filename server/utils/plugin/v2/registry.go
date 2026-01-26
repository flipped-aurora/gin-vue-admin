package plugin

import "sync"

var (
	registryMu sync.RWMutex
	registry   []Plugin
)

// Register records a plugin for auto initialization.
func Register(p Plugin) {
	if p == nil {
		return
	}
	registryMu.Lock()
	registry = append(registry, p)
	registryMu.Unlock()
}

// Registered returns a snapshot of all registered plugins.
func Registered() []Plugin {
	registryMu.RLock()
	defer registryMu.RUnlock()
	out := make([]Plugin, len(registry))
	copy(out, registry)
	return out
}
