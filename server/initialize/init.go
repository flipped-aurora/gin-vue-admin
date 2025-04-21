// 假设这是初始化逻辑的一部分

package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

// 初始化全局函数
func SetupHandlers() {
	// 注册系统重载处理函数
	utils.GlobalSystemEvents.RegisterReloadHandler(func() error {
		return Reload()
	})
}
