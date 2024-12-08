package ast

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

// AddImport 增加 import 方法
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

// FindFunction 查询特定function方法
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

// FindArray 查询特定数组方法
func FindArray(astNode ast.Node, identName, selectorExprName string) *ast.CompositeLit {
	var assignStmt *ast.CompositeLit
	ast.Inspect(astNode, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.AssignStmt:
			for _, expr := range node.Rhs {
				if exprType, ok := expr.(*ast.CompositeLit); ok {
					if arrayType, ok := exprType.Type.(*ast.ArrayType); ok {
						sel, ok1 := arrayType.Elt.(*ast.SelectorExpr)
						x, ok2 := sel.X.(*ast.Ident)
						if ok1 && ok2 && x.Name == identName && sel.Sel.Name == selectorExprName {
							assignStmt = exprType
							return false
						}
					}
				}
			}
		}
		return true
	})
	return assignStmt
}

func CreateMenuStructAst(menus []system.SysBaseMenu) *[]ast.Expr {
	var menuElts []ast.Expr
	for i := range menus {
		elts := []ast.Expr{ // 结构体的字段
			&ast.KeyValueExpr{
				Key:   &ast.Ident{Name: "ParentId"},
				Value: &ast.BasicLit{Kind: token.INT, Value: "0"},
			},
			&ast.KeyValueExpr{
				Key:   &ast.Ident{Name: "Path"},
				Value: &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("\"%s\"", menus[i].Path)},
			},
			&ast.KeyValueExpr{
				Key:   &ast.Ident{Name: "Name"},
				Value: &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("\"%s\"", menus[i].Name)},
			},
			&ast.KeyValueExpr{
				Key:   &ast.Ident{Name: "Hidden"},
				Value: &ast.Ident{Name: "false"},
			},
			&ast.KeyValueExpr{
				Key:   &ast.Ident{Name: "Component"},
				Value: &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("\"%s\"", menus[i].Component)},
			},
			&ast.KeyValueExpr{
				Key:   &ast.Ident{Name: "Sort"},
				Value: &ast.BasicLit{Kind: token.INT, Value: fmt.Sprintf("%d", menus[i].Sort)},
			},
			&ast.KeyValueExpr{
				Key: &ast.Ident{Name: "Meta"},
				Value: &ast.CompositeLit{
					Type: &ast.SelectorExpr{
						X:   &ast.Ident{Name: "model"},
						Sel: &ast.Ident{Name: "Meta"},
					},
					Elts: []ast.Expr{
						&ast.KeyValueExpr{
							Key:   &ast.Ident{Name: "Title"},
							Value: &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("\"%s\"", menus[i].Title)},
						},
						&ast.KeyValueExpr{
							Key:   &ast.Ident{Name: "Icon"},
							Value: &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("\"%s\"", menus[i].Icon)},
						},
					},
				},
			},
		}
		menuElts = append(menuElts, &ast.CompositeLit{
			Type: nil,
			Elts: elts,
		})
	}
	return &menuElts
}

func CreateApiStructAst(apis []system.SysApi) *[]ast.Expr {
	var apiElts []ast.Expr
	for i := range apis {
		elts := []ast.Expr{ // 结构体的字段
			&ast.KeyValueExpr{
				Key:   &ast.Ident{Name: "Path"},
				Value: &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("\"%s\"", apis[i].Path)},
			},
			&ast.KeyValueExpr{
				Key:   &ast.Ident{Name: "Description"},
				Value: &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("\"%s\"", apis[i].Description)},
			},
			&ast.KeyValueExpr{
				Key:   &ast.Ident{Name: "ApiGroup"},
				Value: &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("\"%s\"", apis[i].ApiGroup)},
			},
			&ast.KeyValueExpr{
				Key:   &ast.Ident{Name: "Method"},
				Value: &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("\"%s\"", apis[i].Method)},
			},
		}
		apiElts = append(apiElts, &ast.CompositeLit{
			Type: nil,
			Elts: elts,
		})
	}
	return &apiElts
}

// CheckImport 检查是否存在Import
func CheckImport(file *ast.File, importPath string) bool {
	for _, imp := range file.Imports {
		// Remove quotes around the import path
		path := imp.Path.Value[1 : len(imp.Path.Value)-1]

		if path == importPath {
			return true
		}
	}

	return false
}

func clearPosition(astNode ast.Node) {
	ast.Inspect(astNode, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.Ident:
			// 清除位置信息
			node.NamePos = token.NoPos
		case *ast.CallExpr:
			// 清除位置信息
			node.Lparen = token.NoPos
			node.Rparen = token.NoPos
		case *ast.BasicLit:
			// 清除位置信息
			node.ValuePos = token.NoPos
		case *ast.SelectorExpr:
			// 清除位置信息
			node.Sel.NamePos = token.NoPos
		case *ast.BinaryExpr:
			node.OpPos = token.NoPos
		case *ast.UnaryExpr:
			node.OpPos = token.NoPos
		case *ast.StarExpr:
			node.Star = token.NoPos
		}
		return true
	})
}

func CreateStmt(statement string) *ast.ExprStmt {
	expr, err := parser.ParseExpr(statement)
	if err != nil {
		log.Fatal(err)
	}
	clearPosition(expr)
	return &ast.ExprStmt{X: expr}
}

func IsBlockStmt(node ast.Node) bool {
	_, ok := node.(*ast.BlockStmt)
	return ok
}

func VariableExistsInBlock(block *ast.BlockStmt, varName string) bool {
	exists := false
	ast.Inspect(block, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.AssignStmt:
			for _, expr := range node.Lhs {
				if ident, ok := expr.(*ast.Ident); ok && ident.Name == varName {
					exists = true
					return false
				}
			}
		}
		return true
	})
	return exists
}
