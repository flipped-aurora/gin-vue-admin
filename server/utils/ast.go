package utils

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

func AddRegisterTablesAst(path string, funcName string) {
	src, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, "", src, 0)
	if err != nil {
		fmt.Println(err)
	}
	//ast.Print(fileSet, astFile)
	AddImport(astFile, "qimiao/niubi/666")
	//ast.Print(fileSet, astFile)
	FuncNode := FindFunction(astFile, "Register")
	if FuncNode != nil {
		ast.Print(fileSet, FuncNode)
	}
	AddDBVar(FuncNode.Body, "dbname")
	AddAutoMigrate(FuncNode.Body, "test2", "example", "ExaFile")
	var out []byte
	bf := bytes.NewBuffer(out)
	printer.Fprint(bf, fileSet, astFile)
	fmt.Println(bf.String())
}

func AddImport(astNode ast.Node, imp string) {
	impStr := fmt.Sprintf("\"%s\"", imp)
	ast.Inspect(astNode, func(node ast.Node) bool {
		if genDecl, ok := node.(*ast.GenDecl); ok {
			if genDecl.Tok == token.IMPORT {
				for i := range genDecl.Specs {
					if impNode, ok := genDecl.Specs[i].(*ast.ImportSpec); ok {
						if impNode.Path.Value == impStr {
							return false
						}
					}
				}
				genDecl.Specs = append(genDecl.Specs, &ast.ImportSpec{
					Path: &ast.BasicLit{
						Kind:  token.STRING,
						Value: impStr,
					},
				})
			}
		}
		return true
	})
}

func FindFunction(astNode ast.Node, FunctionName string) *ast.FuncDecl {
	var funcDeclP *ast.FuncDecl
	ast.Inspect(astNode, func(node ast.Node) bool {
		if funcDecl, ok := node.(*ast.FuncDecl); ok {
			if funcDecl.Name.String() == FunctionName {
				funcDeclP = funcDecl
				return false
			}
		}
		return true
	})
	return funcDeclP
}

func AddDBVar(astBody *ast.BlockStmt, dbName string) {
	dbStr := fmt.Sprintf("\"%s\"", dbName)
	for i := range astBody.List {
		if assignStmt, ok := astBody.List[i].(*ast.AssignStmt); ok {
			if ident, ok := assignStmt.Lhs[0].(*ast.Ident); ok {
				if ident.Name == dbName {
					return
				}
			}
		}
	}
	assignNode := &ast.AssignStmt{
		Lhs: []ast.Expr{
			&ast.Ident{
				Name: dbName,
			},
		},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X: &ast.Ident{
						Name: "global",
					},
					Sel: &ast.Ident{
						Name: "GetGlobalDBByDBName",
					},
				},
				Args: []ast.Expr{
					&ast.BasicLit{
						Kind:  token.STRING,
						Value: dbStr,
					},
				},
			},
		},
	}
	astBody.List = append([]ast.Stmt{assignNode}, astBody.List...)
}

func AddAutoMigrate(astBody *ast.BlockStmt, dbname string, pk string, model string) {
	flag := true
	ast.Inspect(astBody, func(node ast.Node) bool {
		// 首先判断需要加入的方法调用语句是否存在 不存在则直接走到下方逻辑
		switch n := node.(type) {
		case *ast.CallExpr:
			// 判断是否找到了AutoMigrate语句
			if s, ok := n.Fun.(*ast.SelectorExpr); ok {
				if x, ok := s.X.(*ast.Ident); ok {
					if s.Sel.Name == "AutoMigrate" && x.Name == dbname {
						flag = false
						if !NeedAppendModel(n, pk, model) {
							return false
						}
						// 判断已经找到了AutoMigrate语句
						n.Args = append(n.Args, &ast.CompositeLit{
							Type: &ast.SelectorExpr{
								X: &ast.Ident{
									Name: pk,
								},
								Sel: &ast.Ident{
									Name: model,
								},
							},
						})
						return false
					}
				}
			}
		}
		return true
		//然后判断 pk.model是否存在 如果存在直接跳出 如果不存在 则向已经找到的方法调用语句的node里面push一条
	})

	if flag {
		exprStmt := &ast.ExprStmt{
			X: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X: &ast.Ident{
						Name: dbname,
					},
					Sel: &ast.Ident{
						Name: "AutoMigrate",
					},
				},
				Args: []ast.Expr{
					&ast.CompositeLit{
						Type: &ast.SelectorExpr{
							X: &ast.Ident{
								Name: pk,
							},
							Sel: &ast.Ident{
								Name: model,
							},
						},
					},
				},
			}}
		astBody.List = append(astBody.List, exprStmt)
	}
}

func NeedAppendModel(callNode ast.Node, pk string, model string) bool {
	flag := true
	ast.Inspect(callNode, func(node ast.Node) bool {
		switch n := node.(type) {
		case *ast.SelectorExpr:
			if x, ok := n.X.(*ast.Ident); ok {
				if n.Sel.Name == model && x.Name == pk {
					flag = false
					return false
				}
			}
		}
		return true
	})
	return flag
}
