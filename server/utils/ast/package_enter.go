package ast

import (
	"go/ast"
	"go/token"
	"io"
)

// PackageEnter 模块化入口
type PackageEnter struct {
	Base
	Type              Type   // 类型
	Path              string // 文件路径
	ImportPath        string // 导包路径
	StructName        string // 结构体名称
	PackageName       string // 包名
	RelativePath      string // 相对路径
	PackageStructName string // 包结构体名称
}

func (a *PackageEnter) Parse(filename string, writer io.Writer) (file *ast.File, err error) {
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

func (a *PackageEnter) Rollback(file *ast.File) error {
	// 无需回滚
	return nil
}

func (a *PackageEnter) Injection(file *ast.File) error {
	_ = NewImport(a.ImportPath).Injection(file)
	ast.Inspect(file, func(n ast.Node) bool {
		genDecl, ok := n.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			return true
		}

		for _, spec := range genDecl.Specs {
			typeSpec, specok := spec.(*ast.TypeSpec)
			if !specok || typeSpec.Name.Name != a.Type.Group() {
				continue
			}

			structType, structTypeOK := typeSpec.Type.(*ast.StructType)
			if !structTypeOK {
				continue
			}

			for _, field := range structType.Fields.List {
				if len(field.Names) == 1 && field.Names[0].Name == a.StructName {
					return true
				}
			}

			field := &ast.Field{
				Names: []*ast.Ident{{Name: a.StructName}},
				Type: &ast.SelectorExpr{
					X:   &ast.Ident{Name: a.PackageName},
					Sel: &ast.Ident{Name: a.PackageStructName},
				},
			}
			structType.Fields.List = append(structType.Fields.List, field)
			return false
		}

		return true
	})
	return nil
}

func (a *PackageEnter) Format(filename string, writer io.Writer, file *ast.File) error {
	if filename == "" {
		filename = a.Path
	}
	return a.Base.Format(filename, writer, file)
}
