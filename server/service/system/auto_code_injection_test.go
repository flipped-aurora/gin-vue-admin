package system

import (
	"errors"
	goast "go/ast"
	"go/parser"
	"go/token"
	"io"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	utilsAst "github.com/flipped-aurora/gin-vue-admin/server/utils/ast"
)

type failingAutoCodeAst struct {
	parseErr     error
	injectionErr error
}

func (a *failingAutoCodeAst) Parse(string, io.Writer) (*goast.File, error) {
	if a.parseErr != nil {
		return nil, a.parseErr
	}
	return &goast.File{}, nil
}

func (a *failingAutoCodeAst) Rollback(*goast.File) error { return nil }

func (a *failingAutoCodeAst) Injection(*goast.File) error { return a.injectionErr }

func (a *failingAutoCodeAst) Format(string, io.Writer, *goast.File) error { return nil }

func parseAutoCodeInjectionSource(t *testing.T, source string) *goast.File {
	t.Helper()
	file, err := parser.ParseFile(token.NewFileSet(), "test.go", source, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse source: %v", err)
	}
	return file
}

func TestPackageInitializeGormInjectionRequiresBizModel(t *testing.T) {
	file := parseAutoCodeInjectionSource(t, "package initialize\n")
	injection := &utilsAst.PackageInitializeGorm{
		ImportPath:  `"example/model"`,
		PackageName: "example",
		StructName:  "Order",
	}

	err := injection.Injection(file)
	if err == nil {
		t.Fatal("Injection() error = nil, want missing bizModel error")
	}
}

func TestPluginInitializeGormInjectionRequiresGormFunction(t *testing.T) {
	file := parseAutoCodeInjectionSource(t, "package initialize\n")
	injection := &utilsAst.PluginInitializeGorm{
		ImportPath:  `"example/model"`,
		PackageName: "model",
		StructName:  "Order",
	}

	err := injection.Injection(file)
	if err == nil {
		t.Fatal("Injection() error = nil, want missing Gorm error")
	}
}

func TestRenderAutoCodeInjectionsReturnsParseError(t *testing.T) {
	want := errors.New("parse failed")
	_, _, err := renderAutoCodeInjections(request.AutoCode{AutoMigrate: true}, map[string]utilsAst.Ast{
		"gorm_biz.go=>PackageInitializeGorm": &failingAutoCodeAst{parseErr: want},
	})
	if !errors.Is(err, want) {
		t.Fatalf("renderAutoCodeInjections() error = %v, want %v", err, want)
	}
}

func TestRenderAutoCodeInjectionsReturnsInjectionError(t *testing.T) {
	want := errors.New("injection failed")
	_, _, err := renderAutoCodeInjections(request.AutoCode{AutoMigrate: true}, map[string]utilsAst.Ast{
		"gorm_biz.go=>PackageInitializeGorm": &failingAutoCodeAst{injectionErr: want},
	})
	if !errors.Is(err, want) {
		t.Fatalf("renderAutoCodeInjections() error = %v, want %v", err, want)
	}
}
