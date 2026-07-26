package system

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/internal/testutil"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

func configurePluginSecurityTest(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	old := global.GVA_CONFIG.AutoCode
	global.GVA_CONFIG.AutoCode.Root = root
	global.GVA_CONFIG.AutoCode.Server = "server"
	global.GVA_CONFIG.AutoCode.Web = "web"
	t.Cleanup(func() { global.GVA_CONFIG.AutoCode = old })
	return root
}

func TestRemoveRejectsTraversalBeforeDeleting(t *testing.T) {
	root := configurePluginSecurityTest(t)
	sentinel := filepath.Join(root, "sentinel")
	if err := os.MkdirAll(sentinel, 0755); err != nil {
		t.Fatal(err)
	}
	marker := filepath.Join(sentinel, "marker.txt")
	if err := os.WriteFile(marker, []byte("keep"), 0600); err != nil {
		t.Fatal(err)
	}

	err := AutoCodePlugin.Remove(context.Background(), filepath.Join("..", "..", "sentinel"), "web")
	if err == nil {
		t.Fatal("Remove() error = nil, want invalid plugin name error")
	}
	if _, statErr := os.Stat(marker); statErr != nil {
		t.Fatalf("Remove() touched sentinel outside plugin root: %v", statErr)
	}
}

func TestRemoveRejectsUnknownPluginType(t *testing.T) {
	configurePluginSecurityTest(t)
	if err := AutoCodePlugin.Remove(context.Background(), "announcement", "desktop"); err == nil {
		t.Fatal("Remove() error = nil, want invalid plugin type error")
	}
}

func TestInitMenuReturnsParserError(t *testing.T) {
	root := configurePluginSecurityTest(t)
	testutil.NewMemoryDB(t, &model.SysBaseMenu{}, &model.SysBaseMenuParameter{}, &model.SysBaseMenuBtn{})
	path := filepath.Join(root, "server", "plugin", "announcement", "initialize", "menu.go")
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte("package initialize\nfunc ("), 0600); err != nil {
		t.Fatal(err)
	}

	if err := AutoCodePlugin.InitMenu(context.Background(), request.InitMenu{PlugName: "announcement"}); err == nil {
		t.Fatal("InitMenu() error = nil, want parser error")
	}
}

func TestInitAPIReturnsMissingArrayError(t *testing.T) {
	root := configurePluginSecurityTest(t)
	testutil.NewMemoryDB(t, &model.SysApi{})
	path := filepath.Join(root, "server", "plugin", "announcement", "initialize", "api.go")
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte("package initialize\nvar entities = []int{}\n"), 0600); err != nil {
		t.Fatal(err)
	}

	if err := AutoCodePlugin.InitAPI(context.Background(), request.InitApi{PlugName: "announcement"}); err == nil {
		t.Fatal("InitAPI() error = nil, want missing SysApi array error")
	}
}
