package cache

import (
	"sync"
	"time"
)

// Memory struct contains *memcache.Client
type Memory struct {
	sync.Mutex

	data map[string]*data
}

type data struct {
	Data    interface{}
	Expired time.Time
}

// NewMemory create new memcache
func NewMemory() *Memory {
	return &Memory{
		data: map[string]*data{},
	}
}

// Get return cached value
func (mem *Memory) Get(key string) interface{} {
	if ret, ok := mem.data[key]; ok {
		if ret.Expired.Before(time.Now()) {
			mem.deleteKey(key)
			return nil
		}
		return ret.Data
	}
	return nil
}

// IsExist check value exists in memcache.
func (mem *Memory) IsExist(key string) bool {
	if ret, ok := mem.data[key]; ok {
		if ret.Expired.Before(time.Now()) {
			mem.deleteKey(key)
			return false
		}
		return true
	}
	return false
}

// Set cached value with key and expire time.
func (mem *Memory) Set(key string, val interface{}, timeout time.Duration) (err error) {
	mem.Lock()
	defer mem.Unlock()

	mem.data[key] = &data{
		Data:    val,
		Expired: time.Now().Add(timeout),
	}
	return nil
}

// Delete delete value in memcache.
func (mem *Memory) Delete(key string) error {
	mem.deleteKey(key)
	return nil
}

// deleteKey
func (mem *Memory) deleteKey(key string) {
	mem.Lock()
	defer mem.Unlock()
	delete(mem.data, key)
}
