package ast

import (
	"go/ast"
	"io"
)

type PluginInitializeGorm struct {
	Base
	Type         Type   // 类型
	Path         string // 文件路径
	ImportPath   string // 导包路径
	RelativePath string // 相对路径
	StructName   string // 结构体名称
	PackageName  string // 包名
	IsNew        bool   // 是否使用new关键字 true: new(PackageName.StructName) false: &PackageName.StructName{}
}

func (a *PluginInitializeGorm) Parse(filename string, writer io.Writer) (file *ast.File, err error) {
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

func (a *PluginInitializeGorm) Rollback(file *ast.File) error {
	var needRollBackImport bool
	ast.Inspect(file, func(n ast.Node) bool {
		callExpr, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		selExpr, seok := callExpr.Fun.(*ast.SelectorExpr)
		if !seok || selExpr.Sel.Name != "AutoMigrate" {
			return true
		}
		if len(callExpr.Args) <= 1 {
			needRollBackImport = true
		}
		// 删除指定的参数
		for i, arg := range callExpr.Args {
			compLit, cok := arg.(*ast.CompositeLit)
			if !cok {
				continue
			}

			cselExpr, sok := compLit.Type.(*ast.SelectorExpr)
			if !sok {
				continue
			}

			ident, idok := cselExpr.X.(*ast.Ident)
			if idok && ident.Name == a.PackageName && cselExpr.Sel.Name == a.StructName {
				// 删除参数
				callExpr.Args = append(callExpr.Args[:i], callExpr.Args[i+1:]...)
				break
			}
		}

		return true
	})

	if needRollBackImport {
		_ = NewImport(a.ImportPath).Rollback(file)
	}

	return nil
}

func (a *PluginInitializeGorm) Injection(file *ast.File) error {
	_ = NewImport(a.ImportPath).Injection(file)
	var call *ast.CallExpr
	ast.Inspect(file, func(n ast.Node) bool {
		callExpr, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		selExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
		if ok && selExpr.Sel.Name == "AutoMigrate" {
			call = callExpr
			return false
		}

		return true
	})

	arg := &ast.CompositeLit{
		Type: &ast.SelectorExpr{
			X:   &ast.Ident{Name: a.PackageName},
			Sel: &ast.Ident{Name: a.StructName},
		},
	}

	call.Args = append(call.Args, arg)
	return nil
}

func (a *PluginInitializeGorm) Format(filename string, writer io.Writer, file *ast.File) error {
	if filename == "" {
		filename = a.Path
	}
	return a.Base.Format(filename, writer, file)
}
