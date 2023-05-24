package cache

import (
	"testing"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/stretchr/testify/assert"
)

func TestMemcache(t *testing.T) {
	mem := NewMemcache("127.0.0.1:11211")
	var err error
	timeoutDuration := 10 * time.Second
	if err = mem.Set("username", "silenceper", timeoutDuration); err != nil {
		t.Error("set Error", err)
	}

	if !mem.IsExist("username") {
		t.Error("IsExist Error")
	}
	exists := mem.IsExist("unknown-key")
	assert.Equal(t, false, exists)

	name := mem.Get("username").(string)
	if name != "" {
		if name != "silenceper" {
			t.Error("get Error")
		}
	}
	data := mem.Get("unknown-key")
	assert.Nil(t, data)

	if err = mem.Delete("username"); err != nil {
		t.Errorf("delete Error , err=%v", err)
	}

	err = mem.Delete("unknown-key")
	assert.Equal(t, memcache.ErrCacheMiss, err)
}
