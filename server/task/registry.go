// server/task/registry.go
package task

import (
	"context"
	"encoding/json"
	"sort"
	"sync"
)

// TaskFunc 命名任务函数签名: ctx 由统一 Runner 注入(带 datascope.WithSystem 与超时),
// params 为面板配置的自由 JSON(无参任务忽略即可), 任务函数应自行解析并响应 ctx 取消。
type TaskFunc func(ctx context.Context, params json.RawMessage) error

// TaskMeta 已注册任务的展示元信息(供面板下拉选择)
type TaskMeta struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type registeredTask struct {
	meta TaskMeta
	fn   TaskFunc
}

// 放本体 task 包: initialize 与 service 双方引用, 规避反向依赖/import cycle
var (
	registryMu sync.RWMutex
	registry   = make(map[string]registeredTask)
)

// Register 注册命名任务; 同名重复注册以最后一次为准(系统 reload 幂等)。
func Register(name, description string, fn TaskFunc) {
	registryMu.Lock()
	defer registryMu.Unlock()
	registry[name] = registeredTask{meta: TaskMeta{Name: name, Description: description}, fn: fn}
}

// Get 按名取任务函数
func Get(name string) (TaskFunc, bool) {
	registryMu.RLock()
	defer registryMu.RUnlock()
	t, ok := registry[name]
	if !ok {
		return nil, false
	}
	return t.fn, true
}

// List 返回全部已注册任务元信息, 按名称稳定排序
func List() []TaskMeta {
	registryMu.RLock()
	defer registryMu.RUnlock()
	metas := make([]TaskMeta, 0, len(registry))
	for _, t := range registry {
		metas = append(metas, t.meta)
	}
	sort.Slice(metas, func(i, j int) bool { return metas[i].Name < metas[j].Name })
	return metas
}
