package testutil

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/gin-gonic/gin"
)

// chdirToServer 把测试工作目录切到 server/，使相对路径 config.yaml 可用。
// LoadDefaultTestConfig 与 core.getConfigPath() 的设计假设运行目录是 server/，
// 但 go test 默认在包目录（server/internal/testutil/）运行，需要切换。
// 还原在 t.Cleanup 中完成。
func chdirToServer(t *testing.T) {
	t.Helper()
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Getwd: %v", err)
	}
	// 包路径 server/internal/testutil，向上两级即 server/
	target, err := filepath.Abs(filepath.Join(wd, "..", ".."))
	if err != nil {
		t.Fatalf("Abs: %v", err)
	}
	if err := os.Chdir(target); err != nil {
		t.Fatalf("Chdir %s: %v", target, err)
	}
	t.Cleanup(func() { _ = os.Chdir(wd) })
}

// TestLoadConfig_LoadsConfigYAML 用仓库自带的 server/config.yaml 验证加载。
func TestLoadConfig_LoadsConfigYAML(t *testing.T) {
	old := global.GVA_CONFIG
	t.Cleanup(func() {
		global.GVA_CONFIG = old
		gin.SetMode("") // 还原 gin 模式
	})

	chdirToServer(t)
	v := LoadConfig(t, "config.yaml")
	if v == nil {
		t.Fatal("LoadConfig 返回 nil")
	}
	// 更可靠的断言：DbType 是非空字符串（config.yaml 必定含此字段）
	if global.GVA_CONFIG.System.DbType == "" {
		t.Fatal("Unmarshal 未生效：System.DbType 为空")
	}
	// JWT 应有 SigningKey
	if global.GVA_CONFIG.JWT.SigningKey == "" {
		t.Fatal("Unmarshal 未生效：JWT.SigningKey 为空")
	}
	if gin.Mode() != gin.TestMode {
		t.Fatalf("gin 模式未设为 TestMode, 实际=%s", gin.Mode())
	}
}

func TestLoadDefaultTestConfig_FallsBackToConfigYAML(t *testing.T) {
	old := global.GVA_CONFIG
	t.Cleanup(func() {
		global.GVA_CONFIG = old
		gin.SetMode("")
	})

	chdirToServer(t)
	v := LoadDefaultTestConfig(t)
	if v == nil {
		t.Fatal("LoadDefaultTestConfig 返回 nil")
	}
	if global.GVA_CONFIG.System.DbType == "" {
		t.Fatal("LoadDefaultTestConfig 未成功加载（DbType 为空）")
	}
}

func TestLoadDefaultTestConfig_PanicWhenNoneExist(t *testing.T) {
	// 用一个 nil 的 fakeT，fail(nil, ...) 应 panic。
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("fail(nil,...) 未 panic")
		}
	}()
	fail(nil, "expected panic")
}
