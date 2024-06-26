package ast

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"strings"
)

func AppendNodeToList(stmts []ast.Stmt, stmt ast.Stmt, index int) []ast.Stmt {
	return append(stmts[:index], append([]ast.Stmt{stmt}, stmts[index:]...)...)
}

func AddRouterCode(path, funcName, pk, model string) {
	src, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, "", src, parser.ParseComments)

	if err != nil {
		fmt.Println(err)
	}

	FuncNode := FindFunction(astFile, funcName)

	pkName := strings.ToUpper(pk[:1]) + pk[1:]
	routerName := fmt.Sprintf("%sRouter", pk)
	modelName := fmt.Sprintf("Init%sRouter", model)
	var bloctPre *ast.BlockStmt
	for i := len(FuncNode.Body.List) - 1; i >= 0; i-- {
		if block, ok := FuncNode.Body.List[i].(*ast.BlockStmt); ok {
			bloctPre = block
		}
	}
	ast.Print(fileSet, FuncNode)
	if ok, b := needAppendRouter(FuncNode, pk); ok {
		routerNode :=
			&ast.BlockStmt{
				List: []ast.Stmt{
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							&ast.Ident{Name: routerName},
						},
						Tok: token.DEFINE,
						Rhs: []ast.Expr{
							&ast.SelectorExpr{
								X: &ast.SelectorExpr{
									X:   &ast.Ident{Name: "router"},
									Sel: &ast.Ident{Name: "RouterGroupApp"},
								},
								Sel: &ast.Ident{Name: pkName},
							},
						},
					},
				},
			}

		FuncNode.Body.List = AppendNodeToList(FuncNode.Body.List, routerNode, len(FuncNode.Body.List)-1)
		bloctPre = routerNode
	} else {
		bloctPre = b
	}

	if needAppendInit(FuncNode, routerName, modelName) {
		bloctPre.List = append(bloctPre.List,
			&ast.ExprStmt{
				X: &ast.CallExpr{
					Fun: &ast.SelectorExpr{
						X:   &ast.Ident{Name: routerName},
						Sel: &ast.Ident{Name: modelName},
					},
					Args: []ast.Expr{
						&ast.Ident{
							Name: "privateGroup",
						},
						&ast.Ident{
							Name: "publicGroup",
						},
					},
				},
			})
	}
	var out []byte
	bf := bytes.NewBuffer(out)
	printer.Fprint(bf, fileSet, astFile)
	os.WriteFile(path, bf.Bytes(), 0666)
}

func needAppendRouter(funcNode ast.Node, pk string) (bool, *ast.BlockStmt) {
	flag := true
	var block *ast.BlockStmt
	ast.Inspect(funcNode, func(node ast.Node) bool {
		switch n := node.(type) {
		case *ast.BlockStmt:
			for i := range n.List {
				if assignNode, ok := n.List[i].(*ast.AssignStmt); ok {
					if identNode, ok := assignNode.Lhs[0].(*ast.Ident); ok {
						if identNode.Name == fmt.Sprintf("%sRouter", pk) {
							flag = false
							block = n
							return false
						}
					}
				}
			}

		}
		return true
	})
	return flag, block
}

func needAppendInit(funcNode ast.Node, routerName string, modelName string) bool {
	flag := true
	ast.Inspect(funcNode, func(node ast.Node) bool {
		switch n := funcNode.(type) {
		case *ast.CallExpr:
			if selectNode, ok := n.Fun.(*ast.SelectorExpr); ok {
				x, xok := selectNode.X.(*ast.Ident)
				if xok && x.Name == routerName && selectNode.Sel.Name == modelName {
					flag = false
					return false
				}
			}
		}
		return true
	})
	return flag
}
