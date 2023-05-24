package cache

import (
	"encoding/json"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

// Memcache struct contains *memcache.Client
type Memcache struct {
	conn *memcache.Client
}

// NewMemcache create new memcache
func NewMemcache(server ...string) *Memcache {
	mc := memcache.New(server...)
	return &Memcache{mc}
}

// Get return cached value
func (mem *Memcache) Get(key string) interface{} {
	var err error
	var item *memcache.Item
	if item, err = mem.conn.Get(key); err != nil {
		return nil
	}
	var result interface{}
	if err = json.Unmarshal(item.Value, &result); err != nil {
		return nil
	}
	return result
}

// IsExist check value exists in memcache.
func (mem *Memcache) IsExist(key string) bool {
	if _, err := mem.conn.Get(key); err != nil {
		return false
	}
	return true
}

// Set cached value with key and expire time.
func (mem *Memcache) Set(key string, val interface{}, timeout time.Duration) (err error) {
	var data []byte
	if data, err = json.Marshal(val); err != nil {
		return err
	}

	item := &memcache.Item{Key: key, Value: data, Expiration: int32(timeout / time.Second)}
	return mem.conn.Set(item)
}

// Delete delete value in memcache.
func (mem *Memcache) Delete(key string) error {
	return mem.conn.Delete(key)
}
