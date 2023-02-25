package ast

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"log"
	"os"
	"strconv"
	"strings"
)

type Visitor struct {
	ImportCode  string
	StructName  string
	PackageName string
	GroupName   string
}

func (vi *Visitor) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.GenDecl:
		// 查找有没有import context包
		// Notice：没有考虑没有import任何包的情况
		if n.Tok == token.IMPORT && vi.ImportCode != "" {
			vi.addImport(n)
			// 不需要再遍历子树
			return nil
		}
		if n.Tok == token.TYPE && vi.StructName != "" && vi.PackageName != "" && vi.GroupName != "" {
			vi.addStruct(n)
			return nil
		}
	case *ast.FuncDecl:
		if n.Name.Name == "Routers" {
			vi.addFuncBodyVar(n)
			return nil
		}

	}
	return vi
}

func (vi *Visitor) addStruct(genDecl *ast.GenDecl) ast.Visitor {
	for i := range genDecl.Specs {
		switch n := genDecl.Specs[i].(type) {
		case *ast.TypeSpec:
			if strings.Index(n.Name.Name, "Group") > -1 {
				switch t := n.Type.(type) {
				case *ast.StructType:
					f := &ast.Field{
						Names: []*ast.Ident{
							{
								Name: vi.StructName,
								Obj: &ast.Object{
									Kind: ast.Var,
									Name: vi.StructName,
								},
							},
						},
						Type: &ast.SelectorExpr{
							X: &ast.Ident{
								Name: vi.PackageName,
							},
							Sel: &ast.Ident{
								Name: vi.GroupName,
							},
						},
					}
					t.Fields.List = append(t.Fields.List, f)
				}
			}
		}
	}
	return vi
}

func (vi *Visitor) addImport(genDecl *ast.GenDecl) ast.Visitor {
	// 是否已经import
	hasImported := false
	for _, v := range genDecl.Specs {
		importSpec := v.(*ast.ImportSpec)
		// 如果已经包含
		if importSpec.Path.Value == strconv.Quote(vi.ImportCode) {
			hasImported = true
		}
	}
	if !hasImported {
		genDecl.Specs = append(genDecl.Specs, &ast.ImportSpec{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: strconv.Quote(vi.ImportCode),
			},
		})
	}
	return vi
}

func (vi *Visitor) addFuncBodyVar(funDecl *ast.FuncDecl) ast.Visitor {
	hasVar := false
	for _, v := range funDecl.Body.List {
		switch varSpec := v.(type) {
		case *ast.AssignStmt:
			for i := range varSpec.Lhs {
				switch nn := varSpec.Lhs[i].(type) {
				case *ast.Ident:
					if nn.Name == vi.PackageName+"Router" {
						hasVar = true
					}
				}
			}
		}
	}
	if !hasVar {
		assignStmt := &ast.AssignStmt{
			Lhs: []ast.Expr{
				&ast.Ident{
					Name: vi.PackageName + "Router",
					Obj: &ast.Object{
						Kind: ast.Var,
						Name: vi.PackageName + "Router",
					},
				},
			},
			Tok: token.DEFINE,
			Rhs: []ast.Expr{
				&ast.SelectorExpr{
					X: &ast.SelectorExpr{
						X: &ast.Ident{
							Name: "router",
						},
						Sel: &ast.Ident{
							Name: "RouterGroupApp",
						},
					},
					Sel: &ast.Ident{
						Name: cases.Title(language.English).String(vi.PackageName),
					},
				},
			},
		}
		funDecl.Body.List = append(funDecl.Body.List, funDecl.Body.List[1])
		index := 1
		copy(funDecl.Body.List[index+1:], funDecl.Body.List[index:])
		funDecl.Body.List[index] = assignStmt
	}
	return vi
}

func ImportReference(filepath, importCode, structName, packageName, groupName string) error {
	fSet := token.NewFileSet()
	fParser, err := parser.ParseFile(fSet, filepath, nil, parser.ParseComments)
	if err != nil {
		return err
	}
	importCode = strings.TrimSpace(importCode)
	v := &Visitor{
		ImportCode:  importCode,
		StructName:  structName,
		PackageName: packageName,
		GroupName:   groupName,
	}
	if importCode == "" {
		ast.Print(fSet, fParser)
	}

	ast.Walk(v, fParser)

	var output []byte
	buffer := bytes.NewBuffer(output)
	err = format.Node(buffer, fSet, fParser)
	if err != nil {
		log.Fatal(err)
	}
	// 写回数据
	return os.WriteFile(filepath, buffer.Bytes(), 0o600)
}
