package cache

import (
	"github.com/liu-cn/ElasticCache"
	"time"
)

var (
	ProxyCache = ElasticCache.New(time.Second * 20)
)
