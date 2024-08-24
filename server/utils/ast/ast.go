package ast

import (
	"fmt"
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

// 检查是否存在Import
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
