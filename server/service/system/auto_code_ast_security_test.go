package system

import (
	"bytes"
	goast "go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"strconv"
	"testing"

	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	astUtils "github.com/flipped-aurora/gin-vue-admin/server/utils/ast"
)

func TestGeneratedStructStringsAreQuoted(t *testing.T) {
	payload := "\";\nfunc injected() {}\nvar injectedValue = \""
	tests := []struct {
		name     string
		elts     *[]goast.Expr
		typeExpr goast.Expr
	}{
		{
			name: "menu",
			elts: astUtils.CreateMenuStructAst([]model.SysBaseMenu{{
				Path: payload, Name: payload, Component: payload,
				Meta:       model.Meta{Title: payload, Icon: payload},
				Parameters: []model.SysBaseMenuParameter{{Type: payload, Key: payload, Value: payload}},
				MenuBtn:    []model.SysBaseMenuBtn{{Name: payload, Desc: payload}},
			}}),
			typeExpr: &goast.SelectorExpr{X: goast.NewIdent("model"), Sel: goast.NewIdent("SysBaseMenu")},
		},
		{
			name:     "api",
			elts:     astUtils.CreateApiStructAst([]model.SysApi{{Path: payload, Description: payload, ApiGroup: payload, Method: payload}}),
			typeExpr: &goast.SelectorExpr{X: goast.NewIdent("model"), Sel: goast.NewIdent("SysApi")},
		},
		{
			name: "dictionary",
			elts: astUtils.CreateDictionaryStructAst([]model.SysDictionary{{
				Name: payload, Type: payload, Desc: payload,
				SysDictionaryDetails: []model.SysDictionaryDetail{{Label: payload, Value: payload, Extend: payload}},
			}}),
			typeExpr: &goast.SelectorExpr{X: goast.NewIdent("model"), Sel: goast.NewIdent("SysDictionary")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := &goast.File{
				Name: goast.NewIdent("initialize"),
				Decls: []goast.Decl{&goast.GenDecl{
					Tok: token.VAR,
					Specs: []goast.Spec{&goast.ValueSpec{
						Names: []*goast.Ident{goast.NewIdent("entities")},
						Values: []goast.Expr{&goast.CompositeLit{
							Type: &goast.ArrayType{Elt: tt.typeExpr},
							Elts: *tt.elts,
						}},
					}},
				}},
			}

			quotedStrings := 0
			goast.Inspect(file, func(node goast.Node) bool {
				lit, ok := node.(*goast.BasicLit)
				if !ok || lit.Kind != token.STRING {
					return true
				}
				quotedStrings++
				if lit.Value != strconv.Quote(payload) {
					t.Errorf("string literal = %q, want %q", lit.Value, strconv.Quote(payload))
				}
				return true
			})
			if quotedStrings == 0 {
				t.Fatal("generated AST did not contain string literals")
			}

			var out bytes.Buffer
			if err := printer.Fprint(&out, token.NewFileSet(), file); err != nil {
				t.Fatalf("printer.Fprint() error = %v", err)
			}
			parsed, err := parser.ParseFile(token.NewFileSet(), tt.name+".go", out.Bytes(), 0)
			if err != nil {
				t.Fatalf("generated source is invalid: %v\n%s", err, out.String())
			}
			for _, decl := range parsed.Decls {
				if _, ok := decl.(*goast.FuncDecl); ok {
					t.Fatalf("payload escaped its string literal:\n%s", out.String())
				}
			}
		})
	}
}
