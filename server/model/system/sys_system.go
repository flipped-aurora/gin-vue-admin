package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/config"
)

// System 配置文件结构体
type System struct {
	Config config.Server `json:"config"`
}
