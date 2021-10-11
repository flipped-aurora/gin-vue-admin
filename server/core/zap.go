package core

import (
	"github.com/flipped-aurora/gin-vue-admin/server/core/internal"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Zap() (logger *zap.Logger) {
	cores := internal.Zap.GetZapCore()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.GVA_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
