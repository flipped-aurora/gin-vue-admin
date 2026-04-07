package ast

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestPluginInitializeGormInjectionUsesBusinessDB(t *testing.T) {
	const source = `package initialize

import (
	"context"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func Gorm(ctx context.Context) {
	err := global.GVA_DB.WithContext(ctx).AutoMigrate()
	if err != nil {
		err = errors.Wrap(err, "注册表失败!")
		zap.L().Error(fmt.Sprintf("%+v", err))
	}
}
`

	dir := t.TempDir()
	path := filepath.Join(dir, "gorm.go")
	if err := os.WriteFile(path, []byte(source), 0o666); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	injection := &PluginInitializeGorm{
		Type:        TypePluginInitializeGorm,
		Path:        path,
		ImportPath:  `"github.com/flipped-aurora/gin-vue-admin/server/plugin/demo/model"`,
		StructName:  "Demo",
		PackageName: "model",
		IsNew:       true,
	}

	businessField := reflect.ValueOf(injection).Elem().FieldByName("Business")
	if !businessField.IsValid() {
		t.Fatal("expected PluginInitializeGorm.Business field")
	}
	businessField.SetString("bizdb")

	file, err := injection.Parse(path, nil)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}
	if err := injection.Injection(file); err != nil {
		t.Fatalf("Injection() error = %v", err)
	}
	if err := injection.Format(path, nil, file); err != nil {
		t.Fatalf("Format() error = %v", err)
	}

	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}

	got := string(content)
	if !strings.Contains(got, "global.GVA_DB.WithContext(ctx).AutoMigrate()") {
		t.Fatalf("expected default gorm block to remain, got:\n%s", got)
	}
	if !strings.Contains(got, `global.MustGetGlobalDBByDBName("bizdb").WithContext(ctx).AutoMigrate(`) {
		t.Fatalf("expected gorm injection to use business db, got:\n%s", got)
	}
	if !strings.Contains(got, "model.Demo{}") {
		t.Fatalf("expected model injection, got:\n%s", got)
	}
}
