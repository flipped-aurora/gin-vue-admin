package internal

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	astutil "github.com/flipped-aurora/gin-vue-admin/server/utils/ast"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/stacktrace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"time"
)

type ZapCore struct {
	level  zapcore.Level
	fields []zapcore.Field // With 附加的常驻字段(如 node/app_id/env),Write 时并入输出
	zapcore.Core
}

func NewZapCore(level zapcore.Level) *ZapCore {
	entity := &ZapCore{level: level}
	syncer := entity.WriteSyncer()
	levelEnabler := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l == level
	})
	entity.Core = zapcore.NewCore(global.GVA_CONFIG.Zap.Encoder(), syncer, levelEnabler)
	return entity
}

// WriteSyncer 返回只写文件的 syncer（不含控制台）。
// formats 用于 business/folder/directory 子目录路由。
func (z *ZapCore) WriteSyncer(formats ...string) zapcore.WriteSyncer {
	cutter := NewCutter(
		global.GVA_CONFIG.Zap.Director,
		z.level.String(),
		global.GVA_CONFIG.Zap.RetentionDay,
		CutterWithLayout(time.DateOnly),
		CutterWithFormats(formats...),
	)
	return zapcore.AddSync(cutter)
}

// WriteSyncerWithConsole 返回文件+控制台的 syncer。
// 仅当本次日志允许进控制台时使用。
func (z *ZapCore) WriteSyncerWithConsole(formats ...string) zapcore.WriteSyncer {
	cutter := NewCutter(
		global.GVA_CONFIG.Zap.Director,
		z.level.String(),
		global.GVA_CONFIG.Zap.RetentionDay,
		CutterWithLayout(time.DateOnly),
		CutterWithFormats(formats...),
	)
	multiSyncer := zapcore.NewMultiWriteSyncer(os.Stdout, cutter)
	return zapcore.AddSync(multiSyncer)
}

// shouldShowConsole 判断本条日志是否应进控制台：
// 全局开关打开，且该模块不在 file-only-modules 清单里。
func shouldShowConsole(mod string) bool {
	return global.GVA_CONFIG.Zap.LogInConsole && !global.GVA_CONFIG.Zap.IsFileOnly(mod)
}

func (z *ZapCore) Enabled(level zapcore.Level) bool {
	return z.level == level
}

// With 必须保留 *ZapCore 包装:zap 的 logger.With 会用返回值替换原 core,
// 若直接返回内层 core.With(...),Write 的错误入库/控制台路由/business 子目录
// 全部被旁路。附加字段暂存于 fields,Write 时并入输出。
func (z *ZapCore) With(fields []zapcore.Field) zapcore.Core {
	combined := make([]zapcore.Field, 0, len(z.fields)+len(fields))
	combined = append(combined, z.fields...)
	combined = append(combined, fields...)
	return &ZapCore{level: z.level, Core: z.Core, fields: combined}
}

func (z *ZapCore) Check(entry zapcore.Entry, check *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if z.Enabled(entry.Level) {
		return check.AddCore(entry, z)
	}
	return check
}

func (z *ZapCore) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	// 读出本条日志的 mod（logger.WithCtx(ctx).Mod("xxx") 附带的字段）
	var mod string
	var subdir string
	for i := 0; i < len(fields); i++ {
		if fields[i].Key == logger.FieldMod && mod == "" {
			mod = fields[i].String
		}
		if subdir == "" && (fields[i].Key == "business" || fields[i].Key == "folder" || fields[i].Key == "directory") {
			subdir = fields[i].String
		}
	}

	// 决定本次写入使用的 syncer：是否进控制台（按模块）、是否进子目录（business 路由）
	var syncer zapcore.WriteSyncer
	switch {
	case shouldShowConsole(mod) && subdir != "":
		syncer = z.WriteSyncerWithConsole(subdir)
	case shouldShowConsole(mod):
		syncer = z.WriteSyncerWithConsole()
	case subdir != "":
		syncer = z.WriteSyncer(subdir)
	default:
		syncer = z.WriteSyncer()
	}
	// 用局部 core 写入:不改写 z.Core(共享状态,并发 Write 下会互相踩踏),
	// 并把 With 暂存的常驻字段(node/app_id/env)并入本条输出
	core := zapcore.NewCore(global.GVA_CONFIG.Zap.Encoder(), syncer, z.level)
	if len(z.fields) > 0 {
		fields = append(append(make([]zapcore.Field, 0, len(z.fields)+len(fields)), z.fields...), fields...)
	}

	// 先写入原日志目标
	err := core.Write(entry, fields)

	// 捕捉 Error 及以上级别日志并入库，且可提取 zap.Error(err) 的错误内容
	if entry.Level >= zapcore.ErrorLevel {
		// 避免与 GORM zap 写入互相递归：跳过由 gorm logger writer 触发的日志
		if strings.Contains(entry.Caller.File, "gorm_logger_writer.go") {
			return err
		}

		form := "后端"
		level := entry.Level.String()
		// 生成基础信息
		info := entry.Message

		// 提取 zap.Error(err) 内容与 request_id/trace_id
		var errStr string
		var reqID string
		var traceID string
		for i := 0; i < len(fields); i++ {
			f := fields[i]
			if f.Key == logger.FieldRequestID && f.String != "" {
				reqID = f.String
			}
			if f.Key == logger.FieldTraceID && f.String != "" {
				traceID = f.String
			}
			if errStr == "" && (f.Type == zapcore.ErrorType || f.Key == "error" || f.Key == "err") {
				if f.Interface != nil {
					errStr = fmt.Sprintf("%v", f.Interface)
				} else if f.String != "" {
					errStr = f.String
				}
			}
		}
		if errStr != "" {
			info = fmt.Sprintf("%s | 错误: %s", info, errStr)
		}

		// 附加来源与堆栈信息
		if entry.Caller.File != "" {
			info = fmt.Sprintf("%s \n 源文件:%s:%d", info, entry.Caller.File, entry.Caller.Line)
		}
		stack := entry.Stack
		if stack != "" {
			info = fmt.Sprintf("%s \n 调用栈：%s", info, stack)
			// 解析最终业务调用方，并提取其方法源码
			if frame, ok := stacktrace.FindFinalCaller(stack); ok {
				fnName, fnSrc, sLine, eLine, exErr := astutil.ExtractFuncSourceByPosition(frame.File, frame.Line)
				if exErr == nil {
					info = fmt.Sprintf("%s \n 最终调用方法:%s:%d (%s lines %d-%d)\n----- 产生日志的方法代码如下 -----\n%s", info, frame.File, frame.Line, fnName, sLine, eLine, fnSrc)
				} else {
					info = fmt.Sprintf("%s \n 最终调用方法:%s:%d (%s) | extract_err=%v", info, frame.File, frame.Line, fnName, exErr)
				}
			}
		}

		// 使用后台上下文，避免依赖 gin.Context
		ctx := context.Background()
		_ = service.ServiceGroupApp.SystemServiceGroup.SysErrorService.CreateSysError(ctx, &system.SysError{
			Form:      &form,
			Info:      &info,
			Level:     level,
			RequestID: reqID,
			TraceID:   traceID,
		})
	}
	return err
}

func (z *ZapCore) Sync() error {
	return z.Core.Sync()
}
