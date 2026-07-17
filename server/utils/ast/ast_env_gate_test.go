package ast

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// TestMain 环境门禁: 本包测试依赖自动化代码生成产物(server/plugin/gva)存在,
// 且会对仓库内真实源码文件(api/service/router 的 enter.go、initialize/*_biz.go 等)做注入/回滚改写。
// 干净检出(未生成 gva 插件)时整包跳过, 避免误报失败与污染工作区;
// 开发者按原工作流生成 gva 插件后, 测试行为与之前完全一致。
func TestMain(m *testing.M) {
	fixture := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", "gva")
	if _, err := os.Stat(fixture); os.IsNotExist(err) {
		fmt.Printf("SKIP utils/ast: 测试夹具 %s 不存在(需先经代码生成器生成 gva 插件), 整包跳过\n", fixture)
		os.Exit(0)
	}
	os.Exit(m.Run())
}
