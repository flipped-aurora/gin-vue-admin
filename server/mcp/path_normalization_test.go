package mcpTool

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/stretchr/testify/require"
)

func TestBuildDirectoryStructureReturnsRootRelativePaths(t *testing.T) {
	original := global.GVA_CONFIG.AutoCode
	t.Cleanup(func() {
		global.GVA_CONFIG.AutoCode = original
	})

	root := filepath.Join(t.TempDir(), "repo")
	global.GVA_CONFIG.AutoCode = config.Autocode{
		Root:   root,
		Server: "server",
		Web:    "web",
		Module: "demo-module",
	}

	paths := (&GVAExecutor{}).buildDirectoryStructure(&ExecutionPlan{
		PackageName: "demo",
		PackageType: "package",
	})

	require.Equal(t, "server/api/v1/demo", paths["api"])
	require.Equal(t, "server/service/demo", paths["service"])
	require.Equal(t, "server/model/demo", paths["model"])
	require.Equal(t, "web/view/demo", paths["vue_page"])
	require.Equal(t, "web/api/demo", paths["vue_api"])
	require.NotContains(t, filepath.ToSlash(paths["api"]), filepath.ToSlash(root))
}

func TestCollectExpectedFilePathsReturnsRootRelativePaths(t *testing.T) {
	original := global.GVA_CONFIG.AutoCode
	t.Cleanup(func() {
		global.GVA_CONFIG.AutoCode = original
	})

	root := filepath.Join(t.TempDir(), "repo")
	global.GVA_CONFIG.AutoCode = config.Autocode{
		Root:   root,
		Server: "server",
		Web:    "web",
	}

	paths := (&GVAExecutor{}).collectExpectedFilePaths(&ExecutionPlan{
		PackageName:        "demo",
		PackageType:        "package",
		NeedCreatedModules: true,
		ModulesInfo: []*request.AutoCode{
			{StructName: "DemoItem"},
		},
	})

	require.Contains(t, paths, "server/api/v1/demo/demoitem.go")
	require.Contains(t, paths, "server/service/demo/demoitem.go")
	require.Contains(t, paths, "server/model/demo/demoitem.go")
	require.Contains(t, paths, "web/view/demo/demoitem.vue")
	require.NotContains(t, strings.Join(paths, "\n"), filepath.ToSlash(root))
}

func TestScanModulesInDirectoryReturnsRootRelativeFilePaths(t *testing.T) {
	original := global.GVA_CONFIG.AutoCode
	t.Cleanup(func() {
		global.GVA_CONFIG.AutoCode = original
	})

	root := t.TempDir()
	global.GVA_CONFIG.AutoCode = config.Autocode{
		Root:   root,
		Server: "server",
	}

	modelDir := filepath.Join(root, "server", "model", "demo")
	require.NoError(t, os.MkdirAll(modelDir, 0o755))
	require.NoError(t, os.WriteFile(filepath.Join(modelDir, "demo_item.go"), []byte("package demo\n"), 0o644))

	modules, err := (&GVAAnalyzer{}).scanModulesInDirectory(modelDir, "demo", "package")
	require.NoError(t, err)
	require.Len(t, modules, 1)
	require.Equal(t, []string{"server/model/demo/demo_item.go"}, modules[0].FilePaths)
}
