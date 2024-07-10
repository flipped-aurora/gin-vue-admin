package ast

import (
	"github.com/pkg/errors"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
)

// PluginEnter 插件化入口
// ModuleName := PackageName.GroupName.ServiceName
type PluginEnter struct {
	Type            Type   // 类型
	Path            string // 文件路径
	ImportPath      string // 导包路径
	StructName      string // 结构体名称
	StructCamelName string // 结构体小驼峰名称
	ModuleName      string // 模块名称
	GroupName       string // 分组名称
	PackageName     string // 包名
	ServiceName     string // 服务名称
}

func (a *PluginEnter) Rollback() error {
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, a.Path, nil, parser.ParseComments)
	if err != nil {
		return errors.Wrapf(err, "[filepath:%s]打开文件失败!", a.Path)
	}

	for i := 0; i < len(file.Decls); i++ {
		v1, o1 := file.Decls[i].(*ast.GenDecl)
		if o1 {
			for j := 0; j < len(v1.Specs); j++ {
				v2, o2 := v1.Specs[j].(*ast.ValueSpec)
				if o2 {
					for k := 0; k < len(v2.Values); k++ {
						v3, o3 := v2.Values[k].(*ast.Ident)
						if o3 {
							if v3.Obj.Name != a.Type.Group() {
								break
							}
							v4, o4 := v3.Obj.Decl.(*ast.TypeSpec)
							if o4 {
								if v4.Name.Name != a.Type.Group() {
									break
								}
								v5, o5 := v4.Type.(*ast.StructType)
								if o5 {
									for l := 0; l < len(v5.Fields.List); l++ {
										v6, o6 := v5.Fields.List[l].Type.(*ast.Ident)
										if len(v5.Fields.List[l].Names) >= 1 && o6 && v6.Name == a.StructCamelName && v5.Fields.List[l].Names[0].Name == a.StructName {
											v5.Fields.List = append(v5.Fields.List[:l], v5.Fields.List[l+1:]...)
										}
									}
								}
							}
						}
					}
					if len(v2.Names) >= 1 && v2.Names[0].Name == a.ModuleName {
						if len(v1.Specs) <= 1 {
							err = NewImport(file, a.ImportPath).Rollback()
							if err != nil {
								return err
							}
						}
						v1.Specs = append(v1.Specs[:j], v1.Specs[j+1:]...)
					}
				}
			}
		}
	}
	create, err := os.Create(a.Path)
	if err != nil {
		return errors.Wrapf(err, "[filepath:%s]打开文件失败!", a.Path)
	}
	defer func() {
		_ = create.Close()
	}()
	err = format.Node(create, fileSet, file)
	if err != nil {
		return errors.Wrapf(err, "[filepath:%s]注入失败!", a.Path)
	}
	return nil
}

func (a *PluginEnter) Injection() error {
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, a.Path, nil, parser.ParseComments)
	if err != nil {
		return errors.Wrapf(err, "[filepath:%s]打开文件失败!", a.Path)
	}
	err = NewImport(file, a.ImportPath).Injection()
	if err != nil {
		return err
	}
	for i := 0; i < len(file.Decls); i++ {
		v1, o1 := file.Decls[i].(*ast.GenDecl)
		if o1 {
			for j := 0; j < len(v1.Specs); j++ {
				v2, o2 := v1.Specs[j].(*ast.ValueSpec)
				if o2 {
					for k := 0; k < len(v2.Values); k++ {
						v3, o3 := v2.Values[k].(*ast.Ident)
						if o3 {
							if v3.Obj.Name != a.Type.Group() {
								break
							}
							v4, o4 := v3.Obj.Decl.(*ast.TypeSpec)
							if o4 {
								if v4.Name.Name != a.Type.Group() {
									break
								}
								v5, o5 := v4.Type.(*ast.StructType)
								if o5 {
									var has bool
									for l := 0; l < len(v5.Fields.List); l++ {
										v6, o6 := v5.Fields.List[l].Type.(*ast.Ident)
										if len(v5.Fields.List[l].Names) >= 1 && o6 && v6.Name == a.StructCamelName && v5.Fields.List[l].Names[0].Name == a.StructName {
											has = true
										}
									}
									if !has {
										field := &ast.Field{
											Names: []*ast.Ident{{Name: a.StructName}},
											Type:  &ast.Ident{Name: a.StructCamelName},
										}
										v5.Fields.List = append(v5.Fields.List, field)
									}
								}
							}
						}
					}
					var has bool
					if len(v2.Names) >= 1 && v2.Names[0].Name == a.ModuleName {
						has = true
					}
					if !has {
						spec := &ast.ValueSpec{
							Names: []*ast.Ident{{Name: a.ModuleName}},
							Values: []ast.Expr{
								&ast.SelectorExpr{
									X: &ast.SelectorExpr{
										X:   &ast.Ident{Name: a.PackageName},
										Sel: &ast.Ident{Name: a.GroupName},
									},
									Sel: &ast.Ident{Name: a.ServiceName},
								},
							},
						}
						v1.Specs = append(v1.Specs, spec)
					}
				}
			}
		}
	}
	create, err := os.Create(a.Path)
	if err != nil {
		return errors.Wrapf(err, "[filepath:%s]打开文件失败!", a.Path)
	}
	defer func() {
		_ = create.Close()
	}()
	err = format.Node(create, fileSet, file)
	if err != nil {
		return errors.Wrapf(err, "[filepath:%s]注入失败!", a.Path)
	}
	return nil
}
