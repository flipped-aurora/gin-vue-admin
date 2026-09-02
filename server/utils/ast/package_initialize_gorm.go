package ast

import (
	"fmt"
	"go/ast"
	"go/token"
	"io"
)

// PackageInitializeGorm 包初始化gorm
type PackageInitializeGorm struct {
	Base
	Type         Type   // 类型
	Path         string // 文件路径
	ImportPath   string // 导包路径
	Business     string // 业务库 gva => gva, 不要传"gva"
	StructName   string // 结构体名称
	PackageName  string // 包名
	RelativePath string // 相对路径
	IsNew        bool   // 是否使用new关键字 true: new(PackageName.StructName) false: &PackageName.StructName{}
}

func (a *PackageInitializeGorm) Parse(filename string, writer io.Writer) (file *ast.File, err error) {
	if filename == "" {
		if a.RelativePath == "" {
			filename = a.Path
			a.RelativePath = a.Base.RelativePath(a.Path)
			return a.Base.Parse(filename, writer)
		}
		a.Path = a.Base.AbsolutePath(a.RelativePath)
		filename = a.Path
	}
	return a.Base.Parse(filename, writer)
}

func (a *PackageInitializeGorm) Rollback(file *ast.File) error {
	packageNameNum := 0
	// 寻找目标结构
	ast.Inspect(file, func(n ast.Node) bool {
		// 总调用的db变量根据business来决定
		varDB := a.Business + "Db"

		if a.Business == "" {
			varDB = "db"
		}

		callExpr, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		// 检查是不是 db.AutoMigrate() 方法
		selExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
		if !ok || selExpr.Sel.Name != "AutoMigrate" {
			return true
		}

		// 检查调用方是不是 db
		ident, ok := selExpr.X.(*ast.Ident)
		if !ok || ident.Name != varDB {
			return true
		}

		// 删除结构体参数
		for i := 0; i < len(callExpr.Args); i++ {
			if com, comok := callExpr.Args[i].(*ast.CompositeLit); comok {
				if selector, exprok := com.Type.(*ast.SelectorExpr); exprok {
					if x, identok := selector.X.(*ast.Ident); identok {
						if x.Name == a.PackageName {
							packageNameNum++
							if selector.Sel.Name == a.StructName {
								callExpr.Args = append(callExpr.Args[:i], callExpr.Args[i+1:]...)
								i--
							}
						}
					}
				}
			}
		}
		return true
	})

	if packageNameNum == 1 {
		_ = NewImport(a.ImportPath).Rollback(file)
	}
	return nil
}

func (a *PackageInitializeGorm) Injection(file *ast.File) error {
	_ = NewImport(a.ImportPath).Injection(file)
	bizModelDecl := FindFunction(file, "bizModel")
	if bizModelDecl != nil {
		a.addDbVar(bizModelDecl.Body)
	}
	// 寻找目标结构
	ast.Inspect(file, func(n ast.Node) bool {
		// 总调用的db变量根据business来决定
		varDB := a.Business + "Db"

		if a.Business == "" {
			varDB = "db"
		}

		callExpr, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		// 检查是不是 db.AutoMigrate() 方法
		selExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
		if !ok || selExpr.Sel.Name != "AutoMigrate" {
			return true
		}

		// 检查调用方是不是 db
		ident, ok := selExpr.X.(*ast.Ident)
		if !ok || ident.Name != varDB {
			return true
		}

		// 添加结构体参数
		callExpr.Args = append(callExpr.Args, &ast.CompositeLit{
			Type: &ast.SelectorExpr{
				X:   ast.NewIdent(a.PackageName),
				Sel: ast.NewIdent(a.StructName),
			},
		})
		return true
	})
	return nil
}

func (a *PackageInitializeGorm) Format(filename string, writer io.Writer, file *ast.File) error {
	if filename == "" {
		filename = a.Path
	}
	return a.Base.Format(filename, writer, file)
}

// 创建businessDB变量
func (a *PackageInitializeGorm) addDbVar(astBody *ast.BlockStmt) {
	for i := range astBody.List {
		if assignStmt, ok := astBody.List[i].(*ast.AssignStmt); ok {
			if ident, ok := assignStmt.Lhs[0].(*ast.Ident); ok {
				if (a.Business == "" && ident.Name == "db") || ident.Name == a.Business+"Db" {
					return
				}
			}
		}
	}

	// 添加 businessDb := global.GetGlobalDBByDBName("business") 变量
	assignNode := &ast.AssignStmt{
		Lhs: []ast.Expr{
			&ast.Ident{
				Name: a.Business + "Db",
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
						Value: fmt.Sprintf("\"%s\"", a.Business),
					},
				},
			},
		},
	}

	// 添加 businessDb.AutoMigrate() 方法
	autoMigrateCall := &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X: &ast.Ident{
					Name: a.Business + "Db",
				},
				Sel: &ast.Ident{
					Name: "AutoMigrate",
				},
			},
		},
	}

	returnNode := astBody.List[len(astBody.List)-1]
	astBody.List = append(astBody.List[:len(astBody.List)-1], assignNode, autoMigrateCall, returnNode)
}
