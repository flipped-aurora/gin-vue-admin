package ast

import (
	"go/ast"
	"go/token"
	"io"
	"log"
)

// PluginEnter 插件化入口
// ModuleName := PackageName.GroupName.ServiceName
type PluginEnter struct {
	Base
	Type            Type   // 类型
	Path            string // 文件路径
	ImportPath      string // 导包路径
	RelativePath    string // 相对路径
	StructName      string // 结构体名称
	StructCamelName string // 结构体小驼峰名称
	ModuleName      string // 模块名称
	GroupName       string // 分组名称
	PackageName     string // 包名
	ServiceName     string // 服务名称
}

func (a *PluginEnter) Parse(filename string, writer io.Writer) (file *ast.File, err error) {
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

func (a *PluginEnter) Rollback(file *ast.File) error {
	for i := 0; i < len(file.Decls); i++ {
		v1, o1 := file.Decls[i].(*ast.GenDecl)
		if o1 {
			for j := 0; j < len(v1.Specs); j++ {
				v2, o2 := v1.Specs[j].(*ast.ValueSpec)
				if o2 {
					for k := 0; k < len(v2.Values); k++ {
						v3, o3 := v2.Values[k].(*ast.CallExpr)
						if o3 {
							for l := 0; l < len(v3.Args); l++ {
								v4, o4 := v3.Args[l].(*ast.Ident)
								if o4 {
									v5, o5 := v4.Obj.Decl.(*ast.TypeSpec)
									if o5 {
										if v5.Name.Name != a.Type.Group() {
											break
										}
										v6, o6 := v5.Type.(*ast.StructType)
										if o6 {
											for m := 0; m < len(v6.Fields.List); m++ {
												v7, o7 := v6.Fields.List[m].Type.(*ast.Ident)
												if len(v6.Fields.List[m].Names) >= 1 && v6.Fields.List[m].Names[0].Name == a.StructName && o7 && v7.Name == a.StructCamelName {
													v6.Fields.List = append(v6.Fields.List[:m], v6.Fields.List[m+1:]...)
													break
												}
											}
										}
									}
								}
							}
						}
					}
					if a.Type == TypePluginServiceEnter {
						continue
					}
					if len(v2.Names) >= 1 && v2.Names[0].Name == a.ModuleName {
						v1.Specs = append(v1.Specs[:j], v1.Specs[j+1:]...)
						if len(v1.Specs) <= 1 {
							_ = NewImport(a.ImportPath).Rollback(file)
						}
					}
				}
			}
		}
	}
	return nil
}

func (a *PluginEnter) Injection(file *ast.File) error {
	_ = NewImport(a.ImportPath).Injection(file)

	has := false
	hasVar := false
	var firstStruct *ast.StructType
	var varSpec *ast.GenDecl
	ast.Inspect(file, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.TypeSpec:
			if s, ok := x.Type.(*ast.StructType); ok {
				firstStruct = s
				for _, field := range x.Type.(*ast.StructType).Fields.List {
					if len(field.Names) > 0 && field.Names[0].Name == a.StructName {
						has = true
						return false
					}
				}
			}
		}
		return true
	})

	if !has {
		field := &ast.Field{
			Names: []*ast.Ident{{Name: a.StructName}},
			Type:  &ast.Ident{Name: a.StructCamelName},
		}
		firstStruct.Fields.List = append(firstStruct.Fields.List, field)
	}

	if a.Type == TypePluginServiceEnter {
		return nil
	}

	ast.Inspect(file, func(n ast.Node) bool {
		genDecl, ok := n.(*ast.GenDecl)
		if ok && genDecl.Tok == token.VAR {
			log.Println("Found var declaration block")
			for _, spec := range genDecl.Specs {
				valueSpec, vsok := spec.(*ast.ValueSpec)
				if vsok {
					varSpec = genDecl
					for _, name := range valueSpec.Names {
						if name.Name == a.ModuleName {
							hasVar = true
							return false
						}
					}
				}
			}
		}
		return true
	})

	if !hasVar {
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
		varSpec.Specs = append(varSpec.Specs, spec)
	}

	return nil
}

func (a *PluginEnter) Format(filename string, writer io.Writer, file *ast.File) error {
	if filename == "" {
		filename = a.Path
	}
	return a.Base.Format(filename, writer, file)
}
