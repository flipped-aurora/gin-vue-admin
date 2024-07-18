package ast

import (
	"go/ast"
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

func (a *PackageEnter) Rollback(file *ast.File) {
	for i := 0; i < len(file.Decls); i++ {
		v1, o1 := file.Decls[i].(*ast.GenDecl)
		if o1 {
			for j := 0; j < len(v1.Specs); j++ {
				v2, o2 := v1.Specs[j].(*ast.TypeSpec)
				if o2 {
					if v2.Name.Name != a.Type.Group() {
						continue
					}
					v3, o3 := v2.Type.(*ast.StructType)
					if o3 {
						for k := 0; k < len(v3.Fields.List); k++ {
							if len(v3.Fields.List[k].Names) >= 1 && v3.Fields.List[k].Names[0].Name == a.StructName {
								NewImport(a.ImportPath).Rollback(file)
								v3.Fields.List = append(v3.Fields.List[:k], v3.Fields.List[k+1:]...)
							}
						}
					}
				}
			}
		}
	}
}

func (a *PackageEnter) Injection(file *ast.File) {
	NewImport(a.ImportPath).Injection(file)
	for i := 0; i < len(file.Decls); i++ {
		v1, o1 := file.Decls[i].(*ast.GenDecl)
		if o1 {
			for j := 0; j < len(v1.Specs); j++ {
				v2, o2 := v1.Specs[j].(*ast.TypeSpec)
				if o2 {
					if v2.Name.Name != a.Type.Group() {
						continue
					}
					v3, o3 := v2.Type.(*ast.StructType)
					if o3 {
						var has bool
						for k := 0; k < len(v3.Fields.List); k++ {
							if len(v3.Fields.List[k].Names) == 1 && v3.Fields.List[k].Names[0].Name == a.StructName {
								has = true
								break
							}
						}
						if !has {
							field := &ast.Field{
								Names: []*ast.Ident{{Name: a.StructName}},
								Type: &ast.SelectorExpr{
									X:   &ast.Ident{Name: a.PackageName},
									Sel: &ast.Ident{Name: a.PackageStructName},
								},
							}
							v3.Fields.List = append(v3.Fields.List, field)
						}
					}
				}
			}
		}
	}
}

func (a *PackageEnter) Format(filename string, writer io.Writer, file *ast.File) error {
	if filename == "" {
		filename = a.Path
	}
	return a.Base.Format(filename, writer, file)
}
