package testutil

import (
	"fmt"
	"os"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// 以下常量与 core/internal/constant.go 保持一致，但 core/internal 为 internal 包，
// 无法跨包导入，这里硬编码以保持同步（如修改请一并更新 core/internal）。
const (
	configEnv         = "GVA_CONFIG"
	configDefaultFile = "config.yaml"
	configTestFile    = "config.test.yaml"
)

// LoadConfig 自建 viper 实例读取指定 yaml 文件，并 Unmarshal 到 global.GVA_CONFIG。
//
// 与 core.Viper() 的区别：
//   - 不调用 flag.Parse（避免与 go test 的 flag 冲突）
//   - 不注册 WatchConfig（测试不需要热更新，且会持有 goroutine/文件句柄）
//
// 同时设置 gin.SetMode(gin.TestMode)，使依赖 gin 模式的逻辑进入测试路径。
// 文件不存在或解析失败时 t.Fatalf。
//
// 用法：
//
//	func TestMain(m *testing.M) {
//	    testutil.LoadConfig(t, "config.yaml")
//	    m.Run()
//	}
//
// 注意：本函数会覆盖全局 GVA_CONFIG，且不还原（测试通常只加载一次）。
// 如需还原请调用方自行保存旧值。
func LoadConfig(t testing.TB, path string) *viper.Viper {
	t.Helper()
	gin.SetMode(gin.TestMode)

	if _, err := os.Stat(path); err != nil {
		t.Fatalf("testutil: 配置文件不存在: %s (%v)", path, err)
	}

	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		t.Fatalf("testutil: ReadInConfig 失败 (%s): %v", path, err)
	}
	if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
		t.Fatalf("testutil: Unmarshal 配置失败 (%s): %v", path, err)
	}
	return v
}

// LoadDefaultTestConfig 按 TestMode 优先级加载配置到 global.GVA_CONFIG：
//
//  1. 环境变量 GVA_CONFIG 指定的路径
//  2. config.test.yaml（若存在）
//  3. config.yaml（兜底）
//
// 与 core.getConfigPath() 的优先级保持一致，但不会因为 TestMode 文件不存在
// 而 panic——直接回退到 config.yaml。
//
// 用法：
//
//	func TestMain(m *testing.M) {
//	    testutil.LoadDefaultTestConfig(nil)
//	    m.Run()
//	}
//
// t 允许传 nil（在 TestMain 场景里 testing.T 尚未创建），此时失败用 panic 表达；
// 推荐在普通测试中传入 *testing.T 以获得更友好的失败定位。
func LoadDefaultTestConfig(t testing.TB) *viper.Viper {
	gin.SetMode(gin.TestMode)

	candidates := make([]string, 0, 3)
	if env := os.Getenv(configEnv); env != "" {
		candidates = append(candidates, env)
	}
	candidates = append(candidates, configTestFile, configDefaultFile)

	for _, p := range candidates {
		if _, err := os.Stat(p); err == nil {
			return loadConfigOrFail(t, p)
		}
	}
	fail(t, fmt.Sprintf("testutil: 未找到任何配置文件: %v", candidates))
	return nil
}

// loadConfigOrFail 在 t 非 nil 时走 LoadConfig（t.Fatalf），
// 否则 panic 以适配 TestMain 场景。
func loadConfigOrFail(t testing.TB, path string) *viper.Viper {
	if t == nil {
		gin.SetMode(gin.TestMode)
		v := viper.New()
		v.SetConfigFile(path)
		v.SetConfigType("yaml")
		if err := v.ReadInConfig(); err != nil {
			panic(fmt.Sprintf("testutil: ReadInConfig 失败 (%s): %v", path, err))
		}
		if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
			panic(fmt.Sprintf("testutil: Unmarshal 配置失败 (%s): %v", path, err))
		}
		return v
	}
	return LoadConfig(t, path)
}

func fail(t testing.TB, msg string) {
	if t == nil {
		panic(msg)
	}
	t.Fatalf("%s", msg)
}
