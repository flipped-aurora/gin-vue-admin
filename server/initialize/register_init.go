package initialize

import (
	_ "kirer.cn/server/source/example"
	_ "kirer.cn/server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
