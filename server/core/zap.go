package core

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/core/internal"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

// Zap 获取 zap.Logger
// Author [SliverHorn](https://github.com/SliverHorn)
func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.GVA_CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.GVA_CONFIG.Zap.Director)
		_ = os.Mkdir(global.GVA_CONFIG.Zap.Director, os.ModePerm)
	}
	levels := global.GVA_CONFIG.Zap.Levels()
	length := len(levels)
	cores := make([]zapcore.Core, 0, length)
	for i := 0; i < length; i++ {
		core := internal.NewZapCore(levels[i])
		cores = append(cores, core)
	}
	// 通过 Hooks 捕捉 Error 及以上级别日志，写入系统错误表
	dbHook := zap.Hooks(func(entry zapcore.Entry) error {
		// 仅处理 Error 及以上级别
		if entry.Level < zapcore.ErrorLevel {
			return nil
		}
		// 避免与 GORM zap 写入互相递归：跳过由 gorm logger writer 触发的日志
		if strings.Contains(entry.Caller.File, "gorm_logger_writer.go") {
			return nil
		}
		// 避免重复记录 panic 恢复日志，panic 由 GinRecovery 单独捕捉入库
		if strings.Contains(entry.Message, "[Recovery from panic]") {
			return nil
		}

		form := "后端"
		level := entry.Level.String()
		// 尽可能携带来源与堆栈信息（使用 runtime 采集并过滤 zap 内部栈）
		info := entry.Message
		if entry.Caller.File != "" {
			info = fmt.Sprintf("%s | caller=%s:%d", info, entry.Caller.File, entry.Caller.Line)
		}
		stack := entry.Stack
		if stack != "" {
			info = fmt.Sprintf("%s | stack=%s", info, stack)
		}

		// 使用后台上下文，避免依赖 gin.Context
		ctx := context.Background()
		_ = service.ServiceGroupApp.SystemServiceGroup.SysErrorService.CreateSysError(ctx, &system.SysError{
			Form:  &form,
			Info:  &info,
			Level: level,
		})
		return nil
	})

    logger = zap.New(zapcore.NewTee(cores...), dbHook)
    // 启用 Error 及以上级别的堆栈捕捉，确保 entry.Stack 可用
    opts := []zap.Option{zap.AddStacktrace(zapcore.ErrorLevel)}
    if global.GVA_CONFIG.Zap.ShowLine {
        opts = append(opts, zap.AddCaller())
    }
    logger = logger.WithOptions(opts...)
    return logger
}
