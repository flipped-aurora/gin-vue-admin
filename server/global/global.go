// Package global 提供全局变量和单例对象
//
// 功能说明：
// 本包定义了整个应用生命周期中需要全局访问的对象，包括：
// - 数据库连接（GORM）
// - Redis客户端
// - MongoDB客户端
// - 配置对象
// - 日志对象
// - 定时任务管理器
// - 路由信息
// 等全局单例对象。
//
// 设计模式：
// 采用全局变量的方式提供单例对象，方便在整个应用中访问。
// 所有全局变量在main.go的initializeSystem()函数中初始化。
package global

import (
	"fmt"
	"sync"

	"github.com/mark3labs/mcp-go/server"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/qmgo"

	"github.com/flipped-aurora/gin-vue-admin/server/utils/timer"
	"github.com/songzhibin97/gkit/cache/local_cache"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/config"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	// GVA_DB 默认数据库连接（GORM）
	// 功能：主数据库连接，用于大部分业务数据操作
	// 初始化：在initialize.Gorm()中初始化
	GVA_DB *gorm.DB

	// GVA_DBList 多数据库连接映射
	// 功能：支持多数据库配置，key为数据库名称，value为对应的GORM连接
	// 使用场景：多租户系统、分库分表等场景
	// 初始化：在initialize.DBList()中初始化
	GVA_DBList map[string]*gorm.DB

	// GVA_REDIS 默认Redis客户端
	// 功能：用于缓存、会话存储、分布式锁等
	// 初始化：在initialize.Redis()中初始化
	GVA_REDIS redis.UniversalClient

	// GVA_REDISList 多Redis连接映射
	// 功能：支持多个Redis实例，key为Redis名称，value为对应的客户端
	// 使用场景：读写分离、不同业务使用不同Redis实例
	// 初始化：在initialize.RedisList()中初始化
	GVA_REDISList map[string]redis.UniversalClient

	// GVA_MONGO MongoDB客户端
	// 功能：用于MongoDB数据库操作
	// 初始化：在initialize.Mongo.Initialization()中初始化
	GVA_MONGO *qmgo.QmgoClient

	// GVA_CONFIG 服务器配置对象
	// 功能：存储从config.yaml读取的所有配置信息
	// 初始化：在core.Viper()中读取配置后填充
	GVA_CONFIG config.Server

	// GVA_VP Viper配置管理器
	// 功能：用于读取和管理配置文件（config.yaml）
	// 初始化：在core.Viper()中初始化
	GVA_VP *viper.Viper

	// GVA_LOG Zap日志记录器
	// 功能：用于记录应用日志，支持不同日志级别
	// 初始化：在core.Zap()中初始化
	GVA_LOG *zap.Logger

	// GVA_Timer 定时任务管理器
	// 功能：用于管理定时任务，如数据清理、定时统计等
	// 初始化：在initialize.Timer()中初始化
	GVA_Timer timer.Timer = timer.NewTimerTask()

	// GVA_Concurrency_Control 并发控制组
	// 功能：使用singleflight模式防止缓存击穿，相同请求只执行一次
	// 使用场景：高并发场景下的缓存查询、数据库查询等
	GVA_Concurrency_Control = &singleflight.Group{}

	// GVA_ROUTERS 路由信息
	// 功能：存储所有注册的路由信息，用于路由查询、调试等
	// 初始化：在initialize.Routers()中初始化
	GVA_ROUTERS gin.RoutesInfo

	// GVA_ACTIVE_DBNAME 当前激活的数据库名称
	// 功能：在多数据库场景下，标识当前使用的数据库
	GVA_ACTIVE_DBNAME *string

	// GVA_MCP_SERVER MCP服务器实例
	// 功能：Model Context Protocol服务器，用于AI辅助开发
	// 初始化：在MCP相关初始化中创建
	GVA_MCP_SERVER *server.MCPServer

	// BlackCache 本地缓存（黑名单缓存）
	// 功能：用于存储JWT黑名单、IP黑名单等，避免频繁查询数据库
	// 使用场景：JWT token失效后加入黑名单，防止重复使用
	BlackCache local_cache.Cache

	// lock 读写锁
	// 功能：保护GVA_DBList等共享资源的并发访问
	lock sync.RWMutex
)

// GetGlobalDBByDBName 通过名称获取数据库连接
//
// 功能说明：
// 从GVA_DBList中根据数据库名称获取对应的GORM数据库连接。
// 如果数据库不存在，返回nil（不会panic）。
//
// 参数说明：
// - dbname: 数据库名称，对应config.yaml中配置的数据库名称
//
// 返回值：
// - *gorm.DB: 数据库连接，如果不存在则返回nil
//
// 使用示例：
//
//	// 获取名为"business_db"的数据库连接
//	db := global.GetGlobalDBByDBName("business_db")
//	if db != nil {
//	    db.Find(&users)
//	}
//
// 注意事项：
// - 使用读锁保护并发访问
// - 如果数据库不存在，返回nil，调用方需要检查
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return GVA_DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取数据库连接，如果不存在则panic
//
// 功能说明：
// 从GVA_DBList中根据数据库名称获取对应的GORM数据库连接。
// 与GetGlobalDBByDBName不同，如果数据库不存在会panic，用于确保数据库必须存在的情况。
//
// 参数说明：
// - dbname: 数据库名称
//
// 返回值：
// - *gorm.DB: 数据库连接，保证不为nil
//
// 使用示例：
//
//	// 确保数据库存在，否则panic
//	db := global.MustGetGlobalDBByDBName("business_db")
//	db.Find(&users) // 不需要检查nil
//
// 注意事项：
// - 使用读锁保护并发访问
// - 如果数据库不存在或为nil，会panic
// - 适用于确定数据库一定存在的场景
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := GVA_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}

// GetRedis 通过名称获取Redis客户端
//
// 功能说明：
// 从GVA_REDISList中根据Redis名称获取对应的Redis客户端。
// 如果Redis不存在，会panic，确保Redis必须存在。
//
// 参数说明：
// - name: Redis名称，对应config.yaml中配置的Redis名称
//
// 返回值：
// - redis.UniversalClient: Redis客户端，保证不为nil
//
// 使用示例：
//
//	// 获取名为"cache"的Redis客户端
//	redis := global.GetRedis("cache")
//	redis.Set(ctx, "key", "value", time.Hour)
//
// 注意事项：
// - 如果Redis不存在或为nil，会panic并显示错误信息
// - 适用于确定Redis一定存在的场景
func GetRedis(name string) redis.UniversalClient {
	redis, ok := GVA_REDISList[name]
	if !ok || redis == nil {
		panic(fmt.Sprintf("redis `%s` no init", name))
	}
	return redis
}
