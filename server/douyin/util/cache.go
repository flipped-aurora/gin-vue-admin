package util

import (
	"time"

	gocache "github.com/patrickmn/go-cache"
)

// MemCache memory cache.
type MemCache struct {
	c *gocache.Cache
}

// NewMemCache new memory cache.
func NewMemCache() *MemCache {
	c := gocache.New(2*time.Hour, 10*time.Minute)
	return &MemCache{
		c,
	}
}

// Get get value.
func (cache *MemCache) Get(key string) interface{} {
	value, isExist := cache.c.Get(key)
	if isExist {
		return value
	}
	return nil
}

// Set set value.
func (cache *MemCache) Set(key string, val interface{}, timeout time.Duration) error {
	cache.c.Set(key, val, timeout)
	return nil
}

// IsExist key is exist.
func (cache *MemCache) IsExist(key string) bool {
	_, isExist := cache.c.Get(key)
	if isExist {
		return true
	}
	return false
}

// Delete delete key.
func (cache *MemCache) Delete(key string) error {
	cache.c.Delete(key)
	return nil
}
