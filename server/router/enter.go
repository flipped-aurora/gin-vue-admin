package router

import (
	"github.com/gzpz/golf-sales-system/server/router/example"
	"github.com/gzpz/golf-sales-system/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
