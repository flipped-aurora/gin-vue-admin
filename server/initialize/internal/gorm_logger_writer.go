package internal

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm/logger"
)

type Writer struct {
	config config.GeneralDB
	writer logger.Writer
}

func NewWriter(config config.GeneralDB) *Writer {
	return &Writer{config: config}
}

// Printf 格式化打印日志
func (c *Writer) Printf(message string, data ...any) {
	if c.config.LogZap {
		switch c.config.LogLevel() {
		case logger.Silent:
			global.GVA_LOG.Debug(fmt.Sprintf(message, data...))
		case logger.Error:
			global.GVA_LOG.Error(fmt.Sprintf(message, data...))
		case logger.Warn:
			global.GVA_LOG.Warn(fmt.Sprintf(message, data...))
		case logger.Info:
			global.GVA_LOG.Info(fmt.Sprintf(message, data...))
		default:
			global.GVA_LOG.Info(fmt.Sprintf(message, data...))
		}
		return
	}
}
