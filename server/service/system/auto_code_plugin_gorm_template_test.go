package system

import (
	"context"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	utilsAst "github.com/flipped-aurora/gin-vue-admin/server/utils/ast"
)

func TestPluginInitializeGormInjectionCarriesBusinessDB(t *testing.T) {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller() failed")
	}
	repoRoot := filepath.Clean(filepath.Join(filepath.Dir(currentFile), "..", "..", ".."))

	oldRoot := global.GVA_CONFIG.AutoCode.Root
	oldServer := global.GVA_CONFIG.AutoCode.Server
	oldModule := global.GVA_CONFIG.AutoCode.Module
	global.GVA_CONFIG.AutoCode.Root = repoRoot
	global.GVA_CONFIG.AutoCode.Server = "server"
	global.GVA_CONFIG.AutoCode.Module = "github.com/flipped-aurora/gin-vue-admin/server"
	defer func() {
		global.GVA_CONFIG.AutoCode.Root = oldRoot
		global.GVA_CONFIG.AutoCode.Server = oldServer
		global.GVA_CONFIG.AutoCode.Module = oldModule
	}()

	info := request.AutoCode{
		Package:         "demoPlugin",
		PackageName:     "demo",
		HumpPackageName: "demo",
		StructName:      "Demo",
		Abbreviation:    "demo",
		BusinessDB:      "bizdb",
		GenerateServer:  true,
	}
	entity := model.SysAutoCodePackage{
		Template:    "plugin",
		PackageName: info.Package,
	}

	_, asts, _, err := AutoCodePackage.templates(context.Background(), entity, info, false)
	if err != nil {
		t.Fatalf("templates() error = %v", err)
	}

	var pluginInitializeGorm *utilsAst.PluginInitializeGorm
	for _, injection := range asts {
		if candidate, ok := injection.(*utilsAst.PluginInitializeGorm); ok {
			pluginInitializeGorm = candidate
			break
		}
	}
	if pluginInitializeGorm == nil {
		t.Fatal("expected plugin initialize gorm injection")
	}

	businessField := reflect.ValueOf(pluginInitializeGorm).Elem().FieldByName("Business")
	if !businessField.IsValid() || businessField.String() != info.BusinessDB {
		t.Fatalf("expected PluginInitializeGorm.Business = %q, got %v", info.BusinessDB, businessField)
	}
}
