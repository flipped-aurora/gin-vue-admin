package initialize

import (
	_ "github.com/gzpz/golf-sales-system/server/source/example"
	_ "github.com/gzpz/golf-sales-system/server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
