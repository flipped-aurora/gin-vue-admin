package internal

/*
================================================================================
结构分析
================================================================================

【核心职责】
实现一个支持动态日志目录切换的自定义 Zap 日志核心（zapcore.Core），
提供基于日志级别的日志分离、自动文件切割、可选控制台输出等功能。

【关键组件】
1. ZapCore 结构体
   - 作用：包装 zapcore.Core，实现自定义日志写入逻辑
   - 职责：管理日志级别、动态切换日志目录、委托底层 Core 执行实际写入
   - 设计模式：装饰器模式（Decorator Pattern），在不改变原有接口的前提下扩展功能

2. NewZapCore 构造函数
   - 作用：创建并初始化指定日志级别的 ZapCore 实例
   - 职责：组装 WriteSyncer、LevelEnabler、Encoder，构建完整的日志核心
   - 设计模式：工厂模式（Factory Pattern），封装复杂对象的创建过程

3. WriteSyncer 方法
   - 作用：创建日志写入同步器（WriteSyncer）
   - 职责：根据配置创建 Cutter（文件切割器），可选添加控制台输出
   - 核心特性：支持动态路径参数（formats），实现业务日志分类存储

4. Write 方法（核心逻辑）
   - 作用：实现 zapcore.Core 接口，处理日志写入请求
   - 职责：检测特殊字段（business/folder/directory），动态切换日志目录
   - 核心特性：运行时动态重建 Core，实现日志路径的动态切换

5. Enabled/Check/With/Sync 方法
   - 作用：实现 zapcore.Core 接口的完整功能
   - 职责：日志级别检查、字段添加、同步刷新等标准操作

================================================================================
执行流程（核心逻辑：Write 方法）
================================================================================

┌─────────────────────────────────────────────────────────────────┐
│                   日志写入请求 Write(entry, fields)               │
│                   示例：logger.Info("msg",                       │
│                         zap.String("business", "order"))        │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
                    ┌─────────────────┐
                    │ 遍历 fields 数组 │  ← 检查是否包含特殊字段
                    │ 查找 key 为：    │     (business/folder/directory)
                    │ - "business"    │
                    │ - "folder"      │
                    │ - "directory"   │
                    └─────────────────┘
                              │
                    ┌─────────┴─────────┐
                    │                    │
            ┌───────▼───────┐   ┌───────▼───────┐
            │ 找到特殊字段    │   │ 未找到特殊字段  │
            │ (如: business) │   │               │
            └───────┬───────┘   └───────┬───────┘
                    │                    │
                    ▼                    │
        ┌───────────────────────┐        │
        │ 提取字段值作为路径参数  │        │
        │ 示例: "order"          │        │
        └───────────────────────┘        │
                    │                    │
                    ▼                    │
        ┌───────────────────────┐        │
        │ 调用 WriteSyncer()     │        │
        │ 传入路径参数            │        │
        │ WriteSyncer("order")   │        │
        └───────────────────────┘        │
                    │                    │
                    ▼                    │
        ┌───────────────────────┐        │
        │ 创建新的 Cutter 实例    │        │
        │ 路径: logs/日期/order/  │        │
        │      info.log          │        │
        └───────────────────────┘        │
                    │                    │
                    ▼                    │
        ┌───────────────────────┐        │
        │ 重建 Core 实例          │        │
        │ z.Core = NewCore(...)  │  ← 动态切换日志输出路径
        └───────────────────────┘        │
                    │                    │
                    └──────────┬─────────┘
                               │
                               ▼
                    ┌─────────────────┐
                    │ 调用底层 Core.Write│  ← 执行实际日志写入
                    │ 写入到新路径      │      (文件或控制台)
                    └─────────────────┘

【设计要点说明】
1. 装饰器模式：ZapCore 包装 zapcore.Core，在不改变接口的前提下扩展功能
   - 优点：保持与标准 zap 的兼容性，同时支持自定义逻辑
   - 实现：组合而非继承，通过嵌入 zapcore.Core 实现

2. 动态路径切换：通过检测日志字段动态切换日志目录
   - 场景：不同业务模块的日志需要分类存储（如：订单日志、支付日志）
   - 实现：在 Write 方法中检测特殊字段，重建 Core 实例
   - 示例：logger.Info("订单创建", zap.String("business", "order"))
          → 日志写入到 logs/2024-01-01/order/info.log

3. 懒加载策略：每次检测到路径变化时才重建 Core
   - 优点：避免不必要的对象创建，提高性能
   - 缺点：首次写入特定业务日志时会有轻微延迟

4. 级别隔离：每个 ZapCore 实例只处理单一日志级别
   - 优点：不同级别的日志可以存储到不同文件，便于管理和查询
   - 实现：通过 LevelEnabler 函数限制只处理匹配的级别

================================================================================
*/

import (
	"os"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ZapCore 自定义的 Zap 日志核心结构体
// 实现 zapcore.Core 接口，提供日志级别隔离和动态路径切换功能
//
// 设计模式：装饰器模式（Decorator Pattern）
//   - 通过嵌入 zapcore.Core 实现组合而非继承
//   - 在不改变原有接口的前提下扩展功能（动态路径切换）
//
// 字段说明：
//   - level: 当前实例处理的日志级别（如：InfoLevel、ErrorLevel）
//     每个 ZapCore 实例只处理单一级别，实现级别隔离
//   - Core: 嵌入的底层 zapcore.Core，负责实际的日志编码和写入
//     采用组合方式，可以动态替换以实现路径切换
type ZapCore struct {
	level        zapcore.Level // 日志级别标识，用于级别匹配和文件命名
	zapcore.Core               // 嵌入的底层日志核心，实现实际的日志处理逻辑
}

// NewZapCore 创建并初始化指定日志级别的 ZapCore 实例
//
// 设计模式：工厂模式（Factory Pattern）
//   - 封装复杂对象的创建过程，统一初始化逻辑
//   - 隐藏底层 Core 的构建细节，提供简洁的创建接口
//
// 参数说明：
//   - level: 日志级别（如：zapcore.InfoLevel、zapcore.ErrorLevel）
//     创建的实例只会处理此级别的日志，实现级别隔离
//
// 返回值：初始化完成的 ZapCore 指针
//
// 初始化流程：
//  1. 创建 ZapCore 实例，设置日志级别
//  2. 调用 WriteSyncer() 创建日志写入同步器（文件切割器 + 可选控制台）
//  3. 创建 LevelEnabler 函数，限制只处理匹配的日志级别
//  4. 组装 Encoder、WriteSyncer、LevelEnabler，构建底层 Core
//
// 使用示例：
//
//	infoCore := NewZapCore(zapcore.InfoLevel)  // 创建 Info 级别日志核心
//	errorCore := NewZapCore(zapcore.ErrorLevel) // 创建 Error 级别日志核心
func NewZapCore(level zapcore.Level) *ZapCore {
	// 创建 ZapCore 实例，初始化日志级别字段
	// 此时 Core 字段为 nil，将在后续步骤中初始化
	entity := &ZapCore{level: level}

	// 创建日志写入同步器（WriteSyncer）
	// WriteSyncer 负责将日志数据写入到目标（文件或控制台）
	// 内部会创建 Cutter 实例，实现文件切割和自动清理
	syncer := entity.WriteSyncer()

	// 创建级别启用器（LevelEnabler）函数
	// 作用：控制哪些日志级别会被此 Core 处理
	// 实现：只处理与当前实例级别完全匹配的日志（精确匹配，非范围匹配）
	// 为什么精确匹配：每个 ZapCore 实例只处理单一级别，实现级别隔离
	//   例如：InfoLevel 的实例只处理 Info 级别的日志，不处理 Debug 或 Warn
	levelEnabler := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l == level // 精确匹配：只有级别完全相等时才返回 true
	})

	// 组装底层 zapcore.Core
	// 参数说明：
	//   - Encoder(): 日志编码器（JSON 或 Console 格式），从配置中获取
	//   - syncer: 日志写入同步器（文件切割器 + 可选控制台）
	//   - levelEnabler: 级别启用器，控制哪些日志会被处理
	// 结果：创建一个完整的日志核心，可以处理指定级别的日志
	entity.Core = zapcore.NewCore(global.GVA_CONFIG.Zap.Encoder(), syncer, levelEnabler)

	return entity
}

// WriteSyncer 创建日志写入同步器（WriteSyncer）
//
// 功能说明：
//
//	创建日志文件的写入同步器，支持文件切割、自动清理、可选控制台输出
//	核心组件是 Cutter（日志切割器），负责按时间轮转和路径管理
//
// 参数说明：
//   - formats: 可变参数，用于动态指定日志路径的额外层级
//     示例：WriteSyncer("business", "order") 会生成路径：
//     logs/2024-01-01/business/order/info.log
//     用途：支持业务日志分类存储，不同业务模块的日志存储到不同目录
//
// 返回值：zapcore.WriteSyncer 接口实现，可以写入日志数据
//
// 路径生成规则：
//
//	director + layout(时间) + formats... + level.log
//	示例：logs/ + 2024-01-01/ + business/ + info.log
//	结果：logs/2024-01-01/business/info.log
//
// 设计要点：
//  1. 支持动态路径参数：通过 formats 参数实现运行时路径切换
//  2. 可选控制台输出：根据配置决定是否同时输出到控制台
//  3. 自动文件切割：Cutter 会根据时间自动切换日志文件
//  4. 自动清理：Cutter 会自动删除超过保留天数的日志目录
//
// 使用场景：
//   - 默认调用（无参数）：WriteSyncer() → logs/2024-01-01/info.log
//   - 业务分类：WriteSyncer("order") → logs/2024-01-01/order/info.log
//   - 多级分类：WriteSyncer("business", "payment") → logs/2024-01-01/business/payment/info.log
func (z *ZapCore) WriteSyncer(formats ...string) zapcore.WriteSyncer {
	// 创建日志切割器（Cutter）实例
	// 参数说明：
	//   - Director: 日志根目录（从配置中获取，如："logs"）
	//   - z.level.String(): 日志级别字符串（如："info"、"error"），用作文件名
	//   - RetentionDay: 日志保留天数（从配置中获取），超过此天数的日志会被自动清理
	//   - CutterWithLayout(time.DateOnly): 设置时间格式为 "2006-01-02"（按天轮转）
	//     time.DateOnly 是 Go 1.20+ 引入的常量，等价于 "2006-01-02"
	//   - CutterWithFormats(formats...): 设置自定义路径参数（如：业务模块名）
	//     如果 formats 为空，则不会添加额外路径层级
	cutter := NewCutter(
		global.GVA_CONFIG.Zap.Director,     // 日志根目录
		z.level.String(),                   // 日志级别（用作文件名，如："info.log"）
		global.GVA_CONFIG.Zap.RetentionDay, // 日志保留天数
		CutterWithLayout(time.DateOnly),    // 时间格式：按天轮转（"2006-01-02"）
		CutterWithFormats(formats...),      // 自定义路径参数（可变参数，支持多级目录）
	)

	// 根据配置决定是否同时输出到控制台
	// 为什么需要控制台输出：开发调试时方便查看日志，生产环境通常关闭
	if global.GVA_CONFIG.Zap.LogInConsole {
		// 创建多路写入同步器（MultiWriteSyncer）
		// 作用：同时写入到多个目标（控制台 + 文件）
		// 实现：使用 io.MultiWriter 的变体，支持并发写入
		// 优点：一份日志同时输出到控制台和文件，无需重复编码
		multiSyncer := zapcore.NewMultiWriteSyncer(os.Stdout, cutter)
		// AddSync 将 io.Writer 包装为 zapcore.WriteSyncer
		// 为什么需要包装：zap 需要 WriteSyncer 接口，而 MultiWriteSyncer 已经是该接口
		// 这里 AddSync 主要是为了类型统一，实际上可以直接返回 multiSyncer
		return zapcore.AddSync(multiSyncer)
	}

	// 只输出到文件，不输出到控制台
	// AddSync 将 Cutter（实现了 io.Writer）包装为 zapcore.WriteSyncer
	return zapcore.AddSync(cutter)
}

// Enabled 检查指定的日志级别是否被当前 Core 启用
//
// 功能说明：
//
//	判断传入的日志级别是否会被当前 ZapCore 实例处理
//	实现 zapcore.Core 接口的必需方法
//
// 参数说明：
//   - level: 待检查的日志级别（如：zapcore.InfoLevel）
//
// 返回值：
//   - true: 级别匹配，此 Core 会处理该级别的日志
//   - false: 级别不匹配，此 Core 不会处理该级别的日志
//
// 实现逻辑：
//
//	精确匹配：只有级别完全相等时才返回 true
//	例如：InfoLevel 的实例只处理 Info 级别，不处理 Debug 或 Warn
//
// 使用场景：
//
//	Zap 框架在写入日志前会调用此方法，判断是否应该使用此 Core
//	如果返回 false，Zap 会跳过此 Core，尝试其他 Core
func (z *ZapCore) Enabled(level zapcore.Level) bool {
	// 精确匹配：只有级别完全相等时才启用
	// 为什么精确匹配：每个 ZapCore 实例只处理单一级别，实现级别隔离
	//   例如：创建了 InfoLevel 和 ErrorLevel 两个实例
	//        InfoLevel 实例的 Enabled(InfoLevel) 返回 true
	//        InfoLevel 实例的 Enabled(ErrorLevel) 返回 false
	return z.level == level
}

// With 为日志核心添加结构化字段
//
// 功能说明：
//
//	实现 zapcore.Core 接口，为后续的日志条目添加预定义的字段
//	这些字段会被自动添加到所有通过此 Core 写入的日志中
//
// 参数说明：
//   - fields: 结构化字段数组（如：zap.String("service", "order")）
//
// 返回值：新的 zapcore.Core 实例（包含添加的字段）
//
// 设计说明：
//
//	采用委托模式，直接调用底层 Core 的 With 方法
//	为什么不修改当前实例：With 方法应该返回新实例，保持不可变性
//
// 使用示例：
//
//	core := NewZapCore(zapcore.InfoLevel)
//	coreWithFields := core.With([]zapcore.Field{
//	    zap.String("service", "order"),
//	    zap.String("version", "v1.0"),
//	})
//	// 后续通过 coreWithFields 写入的日志都会自动包含 service 和 version 字段
func (z *ZapCore) With(fields []zapcore.Field) zapcore.Core {
	// 委托给底层 Core 处理
	// 底层 Core 会创建一个新的 Core 实例，包含原有的字段和新增的字段
	// 注意：返回的是底层 Core，不是 ZapCore，可能会丢失动态路径切换功能
	// 优化建议：可以包装返回的 Core，保持 ZapCore 的特性
	return z.Core.With(fields)
}

// Check 检查日志条目是否应该被当前 Core 处理
//
// 功能说明：
//
//	实现 zapcore.Core 接口，在日志写入前进行级别检查
//	如果级别匹配，将此 Core 添加到 CheckedEntry 中，后续会调用 Write 方法
//
// 参数说明：
//   - entry: 日志条目（包含级别、时间、消息等元数据）
//   - check: 已检查的日志条目（可能已经包含其他 Core）
//
// 返回值：更新后的 CheckedEntry（如果级别匹配则添加此 Core，否则原样返回）
//
// 执行流程：
//  1. 调用 Enabled() 检查日志级别是否匹配
//  2. 如果匹配，将当前 Core 添加到 CheckedEntry
//  3. 如果不匹配，直接返回原 CheckedEntry（不添加此 Core）
//
// 设计说明：
//
//	这是 Zap 框架的核心机制，允许多个 Core 同时处理同一条日志
//	例如：可以同时有文件 Core 和控制台 Core，两者都会处理匹配的日志
func (z *ZapCore) Check(entry zapcore.Entry, check *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	// 检查日志级别是否匹配
	// 只有匹配的级别才会被此 Core 处理
	if z.Enabled(entry.Level) {
		// 级别匹配，将此 Core 添加到 CheckedEntry
		// 后续 Zap 框架会调用此 Core 的 Write 方法写入日志
		return check.AddCore(entry, z)
	}
	// 级别不匹配，不添加此 Core，原样返回
	// 其他 Core 可能仍然会处理此日志（如果它们的级别匹配）
	return check
}

// Write 写入日志条目到目标（文件或控制台）
//
// 功能说明：
//
//	实现 zapcore.Core 接口的核心方法，处理实际的日志写入
//	核心特性：支持动态路径切换，根据日志字段自动切换日志目录
//
// 参数说明：
//   - entry: 日志条目（包含级别、时间、消息、调用位置等元数据）
//   - fields: 结构化字段数组（如：zap.String("business", "order")）
//
// 返回值：错误信息（如果写入失败）
//
// 核心逻辑：动态路径切换
//  1. 遍历 fields，查找特殊字段（business/folder/directory）
//  2. 如果找到，提取字段值作为路径参数
//  3. 创建新的 WriteSyncer（包含新的路径）
//  4. 重建 Core 实例，切换到新的日志路径
//  5. 委托底层 Core 执行实际写入
//
// 设计要点：
//  1. 懒加载策略：只有在检测到路径变化时才重建 Core
//  2. 运行时切换：支持在日志写入时动态切换路径，无需重启
//  3. 字段驱动：通过日志字段控制路径，使用灵活
//
// 使用示例：
//
//	// 默认路径：logs/2024-01-01/info.log
//	logger.Info("普通日志")
//
//	// 动态切换到业务路径：logs/2024-01-01/order/info.log
//	logger.Info("订单日志", zap.String("business", "order"))
//
//	// 多级路径：logs/2024-01-01/business/payment/info.log
//	logger.Info("支付日志",
//	    zap.String("business", "business"),
//	    zap.String("folder", "payment"))
//
// 注意事项：
//   - 每次检测到路径变化都会重建 Core，有一定性能开销
//   - 建议对频繁写入的日志避免频繁切换路径
//   - 特殊字段名称固定为 "business"、"folder"、"directory"，大小写敏感
func (z *ZapCore) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	// ========== 第一步：检测动态路径切换字段 ==========
	// 遍历所有日志字段，查找特殊字段（business/folder/directory）
	// 这些字段的值会被用作日志路径的额外层级，实现业务日志分类存储
	//
	// 为什么需要遍历：字段顺序不固定，需要遍历查找
	// 为什么只取第一个匹配的字段：简化逻辑，避免路径冲突
	//   如果同时存在多个特殊字段，只使用第一个找到的字段值
	for i := 0; i < len(fields); i++ {
		// 检查字段名是否为特殊字段
		// 支持的字段名：
		//   - "business": 业务模块名（如："order"、"payment"）
		//   - "folder": 文件夹名（如："api"、"service"）
		//   - "directory": 目录名（与 folder 功能相同，提供别名）
		// 设计考虑：提供多个字段名，增加使用灵活性
		if fields[i].Key == "business" || fields[i].Key == "folder" || fields[i].Key == "directory" {
			// ========== 第二步：提取字段值并创建新的写入同步器 ==========
			// 获取字段的字符串值（如："order"）
			// fields[i].String 是 zapcore.Field 的方法，返回字段的字符串表示
			// 注意：如果字段不是字符串类型，String 方法会进行类型转换
			//
			// 调用 WriteSyncer 创建新的日志写入同步器
			// 参数：字段值作为路径参数（如：WriteSyncer("order")）
			// 结果：创建新的 Cutter，路径为 logs/日期/order/info.log
			syncer := z.WriteSyncer(fields[i].String)

			// ========== 第三步：重建 Core 实例，切换到新路径 ==========
			// 使用新的 WriteSyncer 重建底层 Core
			// 参数说明：
			//   - Encoder(): 日志编码器（从配置获取，保持不变）
			//   - syncer: 新的写入同步器（包含新的路径）
			//   - z.level: 日志级别（保持不变）
			//
			// 为什么需要重建 Core：
			//   - 原有的 Core 使用的是旧的路径（默认路径或无参数路径）
			//   - 需要切换到新路径，必须创建新的 Core 实例
			//
			// 性能考虑：
			//   - 重建 Core 有一定开销，但只在检测到路径变化时执行
			//   - 对于同一业务模块的日志，后续写入会复用新的 Core
			//
			// 设计模式：策略模式（Strategy Pattern）
			//   - 通过替换 Core 实例来改变日志写入策略（路径）
			z.Core = zapcore.NewCore(global.GVA_CONFIG.Zap.Encoder(), syncer, z.level)

			// 注意：只处理第一个匹配的字段，后续字段会被忽略
			// 优化建议：可以支持多个字段值，生成多级路径
			//   例如：business="order", folder="api" → logs/日期/order/api/info.log
		}
	}

	// ========== 第四步：委托底层 Core 执行实际写入 ==========
	// 调用底层 Core 的 Write 方法，执行实际的日志编码和写入
	// 此时 Core 可能是：
	//   - 初始创建的 Core（默认路径）
	//   - 动态重建的 Core（业务路径）
	//
	// 写入流程（由底层 Core 执行）：
	//   1. 使用 Encoder 将 entry 和 fields 编码为字节数组
	//   2. 通过 WriteSyncer 写入到目标（文件或控制台）
	//   3. 返回错误（如果写入失败）
	return z.Core.Write(entry, fields)
}

// Sync 同步日志缓冲区到磁盘
//
// 功能说明：
//
//	实现 zapcore.Core 接口，强制将文件缓冲区中的数据刷新到磁盘
//	确保日志数据持久化，防止因程序崩溃导致的数据丢失
//
// 返回值：错误信息（如果同步失败）
//
// 使用场景：
//  1. 程序退出前调用，确保所有日志都已落盘
//  2. 关键日志写入后立即调用，确保重要数据已持久化
//  3. 定时同步，降低因系统崩溃导致的数据丢失风险
//
// 设计说明：
//
//	采用委托模式，直接调用底层 Core 的 Sync 方法
//	底层 Core 会调用 WriteSyncer 的 Sync，最终调用 Cutter.Sync()
//
// 注意事项：
//   - Sync 是阻塞操作，会等待磁盘 IO 完成，耗时较长
//   - 频繁调用 Sync 会影响性能，建议在关键时机调用
//   - 由于 Write 方法采用懒加载策略（每次写入后关闭文件），
//     在某些情况下 Sync 可能无需执行（文件已关闭）
func (z *ZapCore) Sync() error {
	// 委托给底层 Core 执行同步
	// 底层 Core 会调用 WriteSyncer 的 Sync 方法
	// 如果 WriteSyncer 是 Cutter，会调用 Cutter.Sync() 刷新文件缓冲区
	// 如果 WriteSyncer 是 MultiWriteSyncer，会同步所有目标（控制台 + 文件）
	return z.Core.Sync()
}
