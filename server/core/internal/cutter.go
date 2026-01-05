package internal

/*
================================================================================
结构分析
================================================================================

【核心职责】
实现一个基于时间轮转的日志切割器，自动按时间格式生成日志文件路径，并提供日志文件过期清理功能。

【关键组件】
1. Cutter 结构体
   - 作用：日志切割器的核心数据结构，实现 io.Writer 接口
   - 职责：管理日志文件路径生成、文件写入、并发安全控制、自动清理

2. CutterOption 函数类型
   - 作用：函数式选项模式（Functional Options Pattern）
   - 职责：提供灵活的配置方式，支持可选参数和扩展配置

3. NewCutter 构造函数
   - 作用：创建并初始化 Cutter 实例
   - 职责：组装配置参数，应用选项函数，初始化锁

4. Write 方法
   - 作用：实现 io.Writer 接口，核心写入逻辑
   - 职责：动态生成日志文件路径、确保目录存在、清理过期日志、写入数据

5. Sync 方法
   - 作用：同步文件缓冲区到磁盘
   - 职责：确保日志数据持久化，防止进程崩溃时数据丢失

6. removeNDaysFolders 函数
   - 作用：清理超过保留天数的日志目录
   - 职责：遍历日志根目录，删除过期的子目录，防止磁盘空间无限增长

================================================================================
执行流程（Write 方法核心逻辑）
================================================================================

┌─────────────────────────────────────────────────────────────────┐
│                     Write(bytes []byte)                          │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
                    ┌─────────────────┐
                    │ 加锁 (mutex.Lock) │  ← 保证并发安全，同一时刻只有一个goroutine写入
                    └─────────────────┘
                              │
                              ▼
                    ┌─────────────────┐
                    │ 构建文件路径数组 │
                    │ 1. director      │  ← 日志根目录
                    │ 2. layout格式化  │  ← 当前时间按格式转换 (如: "2006-01-02")
                    │ 3. formats...    │  ← 自定义参数 (如: "business")
                    │ 4. level.log     │  ← 日志级别文件名 (如: "info.log")
                    └─────────────────┘
                              │
                              ▼
                    ┌─────────────────┐
                    │ filepath.Join()  │  ← 拼接为完整路径
                    │ 示例:            │
                    │ logs/2006-01-02/ │
                    │ business/info.log│
                    └─────────────────┘
                              │
                              ▼
                    ┌─────────────────┐
                    │ os.MkdirAll()    │  ← 递归创建目录（如果不存在）
                    └─────────────────┘
                              │
                              ▼
                    ┌─────────────────┐
                    │ removeNDaysFolders│ ← 清理N天前的日志目录
                    │ (异步清理，防止 │    (每次写入都检查，但只会删除过期目录)
                    │  磁盘空间耗尽)   │
                    └─────────────────┘
                              │
                              ▼
                    ┌─────────────────┐
                    │ os.OpenFile()    │  ← 以追加模式打开文件
                    │ O_CREATE|O_APPEND│    (文件不存在则创建，存在则追加)
                    └─────────────────┘
                              │
                              ▼
                    ┌─────────────────┐
                    │ file.Write(bytes)│  ← 写入日志数据
                    └─────────────────┘
                              │
                              ▼
                    ┌─────────────────┐
                    │ defer: 关闭文件  │  ← 每次写入后立即关闭（懒加载策略）
                    │ defer: 释放锁   │    避免长时间持有文件句柄
                    └─────────────────┘

【设计要点说明】
1. 懒加载策略：每次 Write 时都重新打开文件，根据当前时间动态选择文件
    - 优点：自动实现时间轮转，无需额外的轮转检测逻辑
    - 缺点：频繁开关文件，性能开销略大（适合日志场景，可接受）

2. 每次写入后关闭文件：防止文件句柄泄漏，但会导致频繁IO
    - 为什么：因为路径可能变化（时间轮转），需要动态切换文件

3. 并发安全：使用 RWMutex 保证同一时刻只有一个 goroutine 执行写入
    - 为什么：避免多个 goroutine 同时操作 file 字段导致竞态

4. 路径生成算法：director + layout(time) + formats... + level.log
    - 示例：logs/ + 2006-01-02/ + business/ + info.log
    - 结果：logs/2006-01-02/business/info.log

================================================================================
*/

import (
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Cutter 日志切割器结构体
// 实现 io.Writer 接口，用于日志的自动轮转和文件管理
// 文件路径生成规则：filepath.Join(director, layout格式化时间, formats..., level+".log")
// 示例：logs/2006-01-02/business/info.log
type Cutter struct {
	level  string // 日志级别标识 (debug, info, warn, error, dpanic, panic, fatal)
	layout string // 时间格式模板 (如: "2006-01-02" 或 "2006-01-02 15:04:05")
	// 用于生成按时间分组的日志目录，支持按天/小时轮转
	formats []string // 自定义路径参数，用于业务分类
	// 示例：[]string{"business"} 会在路径中添加 business 目录层级
	director     string // 日志文件的根目录路径
	retentionDay int    // 日志保留天数，超过此天数的日志目录会被自动清理
	// <= 0 表示不清理（永久保留）
	file *os.File // 当前打开的文件句柄
	// 注意：每次 Write 后会被关闭（懒加载策略），下次 Write 时重新打开
	mutex *sync.RWMutex // 读写互斥锁，保证并发写入的安全性
	// 使用 RWMutex 而非 Mutex，预留了未来支持读操作的扩展性
}

// CutterOption 选项函数类型
// 采用函数式选项模式（Functional Options Pattern），提供灵活的配置方式
// 优势：
//  1. 支持可选参数，无需传递 nil 或零值
//  2. 易于扩展新配置项，无需修改函数签名
//  3. 配置代码可读性强，语义明确
type CutterOption func(*Cutter)

// CutterWithLayout 设置时间格式的选项函数
// 参数 layout: 时间格式化模板，使用 Go 的时间格式 "2006-01-02 15:04:05"
// 示例：CutterWithLayout("2006-01-02") 会生成按天轮转的目录结构
func CutterWithLayout(layout string) CutterOption {
	return func(c *Cutter) {
		c.layout = layout // 设置时间格式，用于后续生成时间戳目录
	}
}

// CutterWithFormats 设置自定义路径格式参数的选项函数
// 参数 format: 可变参数，用于添加额外的路径层级
// 示例：CutterWithFormats("business", "module1") 会在路径中添加 business/module1 目录
// 用途：可以根据业务模块、服务名等维度对日志进行分类存储
func CutterWithFormats(format ...string) CutterOption {
	return func(c *Cutter) {
		if len(format) > 0 { // 只有当传入参数时才设置，避免覆盖默认值
			c.formats = format
		}
	}
}

// NewCutter 创建并初始化日志切割器实例
// 参数说明：
//   - director: 日志文件的根目录路径（必需）
//   - level: 日志级别标识，会作为文件名后缀（必需）
//   - retentionDay: 日志保留天数，<= 0 表示不清理（必需）
//   - options: 可变选项参数，用于设置 layout、formats 等可选配置
//
// 返回值: 初始化完成的 Cutter 指针
//
// 设计模式：建造者模式 + 函数式选项模式
// 示例用法：
//
//	cutter := NewCutter("logs", "info", 7,
//	    CutterWithLayout("2006-01-02"),
//	    CutterWithFormats("business"))
func NewCutter(director string, level string, retentionDay int, options ...CutterOption) *Cutter {
	// 创建 Cutter 实例并初始化必需字段
	// 注意：file 字段为 nil，采用懒加载策略，在首次 Write 时再打开文件
	rotate := &Cutter{
		level:        level,             // 设置日志级别（如：info, error）
		director:     director,          // 设置日志根目录
		retentionDay: retentionDay,      // 设置日志保留天数
		mutex:        new(sync.RWMutex), // 初始化读写锁，保证并发安全
	}

	// 应用所有选项函数，按顺序执行配置
	// 这样设计的好处：用户可以传入多个选项函数，按需定制配置
	for i := 0; i < len(options); i++ {
		options[i](rotate) // 执行选项函数，修改 Cutter 实例的字段
	}
	return rotate
}

// Write 实现 io.Writer 接口，将日志数据写入文件
//
// 核心特性：
//  1. 自动时间轮转：每次写入时根据当前时间生成文件路径，时间变化时自动切换到新文件
//  2. 懒加载策略：每次写入时打开文件，写入后立即关闭，避免长时间持有文件句柄
//  3. 自动清理：每次写入时检查并删除超过保留天数的日志目录
//  4. 并发安全：使用互斥锁保证同一时刻只有一个 goroutine 执行写入
//
// 参数：bytes - 待写入的日志数据（字节数组）
// 返回：n - 成功写入的字节数，err - 错误信息（如果发生）
//
// 执行流程：
//
//	加锁 → 构建文件路径 → 创建目录 → 清理过期日志 → 打开文件 → 写入数据 → 关闭文件 → 解锁
func (c *Cutter) Write(bytes []byte) (n int, err error) {
	// ========== 第一步：加锁保护，保证并发安全 ==========
	// 使用互斥锁确保同一时刻只有一个 goroutine 能够执行写入操作
	// 为什么需要锁：避免多个 goroutine 同时操作 file 字段和文件系统，导致竞态条件
	c.mutex.Lock()

	// defer 确保函数退出时执行清理操作（无论正常返回还是发生错误）
	defer func() {
		// 关闭文件句柄（如果已打开）
		// 为什么每次写入后都关闭：
		//   1. 懒加载策略：下次写入时可能路径已变化（时间轮转），需要打开新文件
		//   2. 避免文件句柄泄漏：确保资源及时释放
		//   3. 便于文件系统同步：关闭后确保数据刷盘
		if c.file != nil {
			_ = c.file.Close() // 忽略关闭错误，避免掩盖写入错误
			c.file = nil       // 清空句柄，防止重复关闭
		}
		// 释放锁，允许其他 goroutine 继续执行
		c.mutex.Unlock()
	}()

	// ========== 第二步：构建日志文件路径 ==========
	// 路径组成：director + layout格式化时间 + formats... + level.log
	// 示例：logs/2006-01-02/business/info.log

	length := len(c.formats) // 获取自定义格式参数的数量
	// 预分配切片容量，避免多次扩容（director + layout + formats + level.log）
	values := make([]string, 0, 3+length)

	values = append(values, c.director) // 添加根目录，如："logs"

	// 如果配置了时间格式，则添加当前时间的格式化字符串
	// 为什么每次都获取当前时间：实现自动时间轮转，时间变化时路径自动改变
	// 例如：layout="2006-01-02" 时，2024-01-01写入→logs/2024-01-01/...，2024-01-02写入→logs/2024-01-02/...
	if c.layout != "" {
		values = append(values, time.Now().Format(c.layout)) // 添加时间目录，如："2006-01-02"
	}

	// 添加所有自定义格式参数（如：业务模块名）
	for i := 0; i < length; i++ {
		values = append(values, c.formats[i]) // 添加自定义目录层级，如："business"
	}

	// 添加日志级别文件名
	values = append(values, c.level+".log") // 添加文件名，如："info.log"

	// 使用 filepath.Join 拼接路径，自动处理不同操作系统的路径分隔符
	// 结果示例：logs/2006-01-02/business/info.log
	filename := filepath.Join(values...)

	// ========== 第三步：确保目标目录存在 ==========
	// 获取文件所在的目录路径（去除文件名）
	director := filepath.Dir(filename)

	// 递归创建目录（如果不存在）
	// os.ModePerm (0777) 表示最大权限，允许读写执行
	// 为什么需要创建：时间轮转或首次运行时，目标目录可能不存在
	err = os.MkdirAll(director, os.ModePerm)
	if err != nil {
		return 0, err // 创建目录失败，直接返回错误
	}

	// ========== 第四步：清理过期日志目录 ==========
	// 每次写入时都执行清理检查（虽然频率较高，但可以防止磁盘空间无限增长）
	// removeNDaysFolders 内部会判断 retentionDay <= 0 时直接返回，开销可接受
	// 优化建议：可以添加时间戳缓存，避免每次都遍历目录（当前实现简单可靠）
	defer func() {
		removeNDaysFolders(c.director, c.retentionDay)
	}()

	// ========== 第五步：打开目标日志文件 ==========
	// 文件打开标志说明：
	//   os.O_CREATE: 文件不存在时创建
	//   os.O_APPEND: 以追加模式打开（写入时追加到文件末尾，不覆盖已有内容）
	//   os.O_WRONLY: 只写模式（提高性能，避免读取操作）
	//   0644: 文件权限（rw-r--r--），所有者可读写，其他用户只读
	c.file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err // 打开文件失败，返回错误（defer 会确保锁被释放）
	}

	// ========== 第六步：写入日志数据 ==========
	// 将数据写入文件缓冲区（系统会自动处理缓冲区刷新）
	// 返回值：n 为实际写入的字节数，err 为错误信息
	return c.file.Write(bytes)
	// 注意：函数返回后，defer 会自动执行，关闭文件并释放锁
}

// Sync 同步文件缓冲区到磁盘
//
// 作用：强制将文件缓冲区中的数据刷新到磁盘，确保数据持久化
// 使用场景：
//  1. 程序即将退出时调用，防止未刷盘的数据丢失
//  2. 关键日志写入后立即调用，确保重要数据已落盘
//  3. 定时同步，降低因系统崩溃导致的数据丢失风险
//
// 返回值：错误信息（如果同步失败）
// 注意：由于 Write 方法采用懒加载策略（每次写入后关闭文件），
//
//	此方法在 Write 调用后可能 file 为 nil，此时返回 nil 表示无需同步
func (c *Cutter) Sync() error {
	// 加锁保护，防止在同步过程中 file 被其他 goroutine 修改或关闭
	c.mutex.Lock()
	defer c.mutex.Unlock() // defer 确保锁一定会被释放

	// 检查文件句柄是否存在（可能因懒加载策略已被关闭）
	if c.file != nil {
		// 调用系统调用 fsync，强制将缓冲区数据写入磁盘
		// 性能开销：fsync 是阻塞调用，会等待磁盘IO完成，耗时较长
		return c.file.Sync()
	}

	// 文件未打开，无需同步（可能已被关闭或尚未首次写入）
	return nil
}

// removeNDaysFolders 清理超过保留天数的日志目录
//
// 功能说明：
//
//	遍历指定根目录下的所有子目录，删除修改时间早于 (当前时间 - days) 的目录
//	此函数会在每次 Write 时被调用，及时清理过期日志，防止磁盘空间无限增长
//
// 参数说明：
//   - dir: 日志根目录路径（清理操作的起始目录）
//   - days: 日志保留天数，<= 0 时直接返回（不执行清理）
//
// 清理策略：
//  1. 只删除目录，不删除单个文件（因为日志文件存储在目录中，删除目录即删除所有文件）
//  2. 保留根目录本身（path != dir 的判断）
//  3. 基于目录的修改时间（ModTime）判断是否过期
//
// 算法流程：
//
//	遍历目录树 → 检查每个子目录的修改时间 → 与截止时间比较 → 删除过期目录
//
// 注意事项：
//   - 此操作会递归删除整个目录及其所有内容，请谨慎使用
//   - 如果目录正在被其他进程使用，删除可能失败
//   - 每次 Write 都执行清理可能带来性能开销，但可接受（可优化为定时清理）
//
// 使用示例：
//
//	示例1：清理7天前的日志目录（保留最近7天的日志）
//		err := removeNDaysFolders("logs", 7)
//		if err != nil {
//			log.Printf("清理日志目录失败: %v", err)
//		}
//		// 结果：删除 logs/ 目录下所有修改时间早于7天前的子目录
//		// 例如：logs/2024-01-01/、logs/2024-01-02/ 等过期目录会被删除
//
//	示例2：清理30天前的日志目录（长期保留策略）
//		err := removeNDaysFolders("/var/log/myapp", 30)
//		if err != nil {
//			log.Printf("清理日志失败: %v", err)
//		}
//		// 结果：只保留最近30天的日志目录，删除更早的目录
//
//	示例3：禁用自动清理功能（永久保留所有日志）
//		err := removeNDaysFolders("logs", 0)  // days <= 0 时直接返回，不执行清理
//		// 或者
//		err := removeNDaysFolders("logs", -1) // 负数也会被忽略
//		// 结果：函数直接返回 nil，不会删除任何目录
//
//	示例4：带错误处理的完整示例
//		logDir := "logs"
//		retentionDays := 7
//		if err := removeNDaysFolders(logDir, retentionDays); err != nil {
//			// 处理可能的错误：权限不足、目录不存在、文件被占用等
//			log.Errorf("清理 %d 天前的日志目录失败 (目录: %s): %v", retentionDays, logDir, err)
//			// 根据业务需求决定是否继续执行或返回错误
//			return err
//		}
//		log.Infof("成功清理 %d 天前的日志目录: %s", retentionDays, logDir)
//
//	示例5：在实际日志写入场景中的使用（由 Cutter.Write 内部调用）
//		cutter := NewCutter("logs", "info", 7) // retentionDay=7 表示保留7天
//		cutter.Write([]byte("log message"))     // Write 内部会调用 removeNDaysFolders("logs", 7)
//		// 结果：每次写入日志时，都会自动清理7天前的日志目录
func removeNDaysFolders(dir string, days int) error {
	// 如果保留天数 <= 0，表示不清理日志，直接返回
	// 这样设计的好处：允许用户通过设置 retentionDay=0 来禁用自动清理功能
	if days <= 0 {
		return nil
	}

	// 计算截止时间：当前时间减去保留天数
	// AddDate(年, 月, 日)，负数表示向前推时间
	// 示例：当前 2024-01-10，days=7，cutoff=2024-01-03
	cutoff := time.Now().AddDate(0, 0, -days)

	// 使用 filepath.Walk 递归遍历目录树
	// 第一个参数：起始目录路径
	// 第二个参数：遍历回调函数，对每个文件和目录都会调用一次
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		// 处理遍历过程中可能出现的错误（如权限不足、文件被删除等）
		if err != nil {
			return err // 返回错误，Walk 会停止遍历
		}

		// 判断是否应该删除此目录：
		//   1. info.IsDir()：必须是目录（不删除单个文件）
		//   2. info.ModTime().Before(cutoff)：修改时间早于截止时间（已过期）
		//   3. path != dir：不是根目录本身（保护根目录不被删除）
		if info.IsDir() && info.ModTime().Before(cutoff) && path != dir {
			// 递归删除整个目录及其所有内容
			// RemoveAll 会删除目录、子目录和所有文件
			err = os.RemoveAll(path)
			if err != nil {
				return err // 删除失败时返回错误（Walk 会停止遍历）
			}
		}

		// 继续遍历下一个文件/目录
		return nil
	})
}
