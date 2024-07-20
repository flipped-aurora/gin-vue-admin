package ast

import (
	"go/ast"
	"go/token"
	"io"
)

// PackageModuleEnter 模块化入口
// ModuleName := PackageName.AppName.GroupName.ServiceName
type PackageModuleEnter struct {
	Base
	Type         Type   // 类型
	Path         string // 文件路径
	ImportPath   string // 导包路径
	RelativePath string // 相对路径
	StructName   string // 结构体名称
	AppName      string // 应用名称
	GroupName    string // 分组名称
	ModuleName   string // 模块名称
	PackageName  string // 包名
	ServiceName  string // 服务名称
}

func (a *PackageModuleEnter) Parse(filename string, writer io.Writer) (file *ast.File, err error) {
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

func (a *PackageModuleEnter) Rollback(file *ast.File) error {
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
							v4, o4 := v3.Fields.List[k].Type.(*ast.Ident)
							if o4 && v4.Name == a.StructName {
								v3.Fields.List = append(v3.Fields.List[:k], v3.Fields.List[k+1:]...)
							}
						}
					}
					continue
				}
				if a.Type == TypePackageServiceModuleEnter {
					continue
				}
				v3, o3 := v1.Specs[j].(*ast.ValueSpec)
				if o3 {
					if len(v3.Names) == 1 && v3.Names[0].Name == a.ModuleName {
						v1.Specs = append(v1.Specs[:j], v1.Specs[j+1:]...)
					}
				}
				if v1.Tok == token.VAR && len(v1.Specs) == 0 {
					_ = NewImport(a.ImportPath).Rollback(file)
					if i == len(file.Decls) {
						file.Decls = append(file.Decls[:i-1])
						break
					} // 空的var(), 如果不删除则会影响的注入变量, 因为识别不到*ast.ValueSpec
					file.Decls = append(file.Decls[:i], file.Decls[i+1:]...)
				}
			}
		}
	}
	return nil
}

func (a *PackageModuleEnter) Injection(file *ast.File) error {
	_ = NewImport(a.ImportPath).Injection(file)
	var hasValue bool
	var hasVariables bool
	for i := 0; i < len(file.Decls); i++ {
		v1, o1 := file.Decls[i].(*ast.GenDecl)
		if o1 {
			if v1.Tok == token.VAR {
				hasVariables = true
			}
			for j := 0; j < len(v1.Specs); j++ {
				if a.Type == TypePackageServiceModuleEnter {
					hasValue = true
				}
				v2, o2 := v1.Specs[j].(*ast.TypeSpec)
				if o2 {
					if v2.Name.Name != a.Type.Group() {
						continue
					}
					v3, o3 := v2.Type.(*ast.StructType)
					if o3 {
						var hasStruct bool
						for k := 0; k < len(v3.Fields.List); k++ {
							v4, o4 := v3.Fields.List[k].Type.(*ast.Ident)
							if o4 && v4.Name == a.StructName {
								hasStruct = true
							}
						}
						if !hasStruct {
							field := &ast.Field{Type: &ast.Ident{Name: a.StructName}}
							v3.Fields.List = append(v3.Fields.List, field)
						}
					}
					continue
				}
				v3, o3 := v1.Specs[j].(*ast.ValueSpec)
				if o3 {
					hasVariables = true
					if len(v3.Names) == 1 && v3.Names[0].Name == a.ModuleName {
						hasValue = true
					}
				}
				if v1.Tok == token.VAR && len(v1.Specs) == 0 {
					hasVariables = false
				} // 说明是空var()
				if hasVariables && !hasValue {
					spec := &ast.ValueSpec{
						Names: []*ast.Ident{{Name: a.ModuleName}},
						Values: []ast.Expr{
							&ast.SelectorExpr{
								X: &ast.SelectorExpr{
									X: &ast.SelectorExpr{
										X:   &ast.Ident{Name: a.PackageName},
										Sel: &ast.Ident{Name: a.AppName},
									},
									Sel: &ast.Ident{Name: a.GroupName},
								},
								Sel: &ast.Ident{Name: a.ServiceName},
							},
						},
					}
					v1.Specs = append(v1.Specs, spec)
					hasValue = true
				}
			}
		}
	}
	if !hasValue && !hasVariables {
		decl := &ast.GenDecl{
			Tok: token.VAR,
			Specs: []ast.Spec{
				&ast.ValueSpec{
					Names: []*ast.Ident{{Name: a.ModuleName}},
					Values: []ast.Expr{
						&ast.SelectorExpr{
							X: &ast.SelectorExpr{
								X: &ast.SelectorExpr{
									X:   &ast.Ident{Name: a.PackageName},
									Sel: &ast.Ident{Name: a.AppName},
								},
								Sel: &ast.Ident{Name: a.GroupName},
							},
							Sel: &ast.Ident{Name: a.ServiceName},
						},
					},
				},
			},
		}
		file.Decls = append(file.Decls, decl)
	}
	return nil
}

func (a *PackageModuleEnter) Format(filename string, writer io.Writer, file *ast.File) error {
	if filename == "" {
		filename = a.Path
	}
	return a.Base.Format(filename, writer, file)
}
