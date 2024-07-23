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
	cutAutoMigrateFuncVisitor := &cutAutoMigrateFunc{
		pkgInitGorm: a,
	}
	ast.Walk(cutAutoMigrateFuncVisitor, file)

	if cutAutoMigrateFuncVisitor.PackageNameNum == 1 {
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

	addAutoMigrateVisitor := &addAutoMigrateFunc{
		pkgInitGorm: a,
	}

	ast.Walk(addAutoMigrateVisitor, file)
	return nil
}

func (a *PackageInitializeGorm) Format(filename string, writer io.Writer, file *ast.File) error {
	if filename == "" {
		filename = a.Path
	}
	return a.Base.Format(filename, writer, file)
}

type addAutoMigrateFunc struct {
	pkgInitGorm *PackageInitializeGorm
}

func (v *addAutoMigrateFunc) Visit(n ast.Node) ast.Visitor {
	// 总调用的db变量根据business来决定
	varDB := v.pkgInitGorm.Business + "Db"

	if v.pkgInitGorm.Business == "" {
		varDB = "db"
	}

	callExpr, ok := n.(*ast.CallExpr)
	if !ok {
		return v
	}

	// 检查是不是 db.AutoMigrate() 方法
	selExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
	if !ok || selExpr.Sel.Name != "AutoMigrate" {
		return v
	}

	// 检查调用方是不是 db
	ident, ok := selExpr.X.(*ast.Ident)
	if !ok || ident.Name != varDB {
		return v
	}

	// 添加结构体参数
	callExpr.Args = append(callExpr.Args, &ast.CompositeLit{
		Type: &ast.SelectorExpr{
			X:   ast.NewIdent(v.pkgInitGorm.PackageName),
			Sel: ast.NewIdent(v.pkgInitGorm.StructName),
		},
	})
	return v
}

type cutAutoMigrateFunc struct {
	pkgInitGorm    *PackageInitializeGorm
	PackageNameNum int
}

func (v *cutAutoMigrateFunc) Visit(n ast.Node) ast.Visitor {
	// 总调用的db变量根据business来决定
	varDB := v.pkgInitGorm.Business + "Db"

	if v.pkgInitGorm.Business == "" {
		varDB = "db"
	}

	callExpr, ok := n.(*ast.CallExpr)
	if !ok {
		return v
	}

	// 检查是不是 db.AutoMigrate() 方法
	selExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
	if !ok || selExpr.Sel.Name != "AutoMigrate" {
		return v
	}

	// 检查调用方是不是 db
	ident, ok := selExpr.X.(*ast.Ident)
	if !ok || ident.Name != varDB {
		return v
	}

	// 删除结构体参数
	for i := range callExpr.Args {
		if com, comok := callExpr.Args[i].(*ast.CompositeLit); comok {
			if selector, exprok := com.Type.(*ast.SelectorExpr); exprok {
				if x, identok := selector.X.(*ast.Ident); identok {
					if x.Name == v.pkgInitGorm.PackageName {
						v.PackageNameNum++
						if selector.Sel.Name == v.pkgInitGorm.StructName {
							callExpr.Args = append(callExpr.Args[:i], callExpr.Args[i+1:]...)
							i--
						}
					}
				}
			}
		}
	}
	return v
}

// 创建businessDB变量
func (a *PackageInitializeGorm) addDbVar(astBody *ast.BlockStmt) {
	for i := range astBody.List {
		if assignStmt, ok := astBody.List[i].(*ast.AssignStmt); ok {
			if ident, ok := assignStmt.Lhs[0].(*ast.Ident); ok {
				if ident.Name == "db" || ident.Name == a.Business+"Db" {
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
