package system

import (
	"github.com/gzpz/golf-sales-system/server/config"
)

// 配置文件结构体
type System struct {
	Config config.Server `json:"config"`
}
