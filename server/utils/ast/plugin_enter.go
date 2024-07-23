package ast

import (
	"go/ast"
	"go/token"
	"io"
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
	//回滚结构体内内容
	var structType *ast.StructType
	ast.Inspect(file, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.TypeSpec:
			if s, ok := x.Type.(*ast.StructType); ok {
				structType = s
				for i, field := range x.Type.(*ast.StructType).Fields.List {
					if len(field.Names) > 0 && field.Names[0].Name == a.StructName {
						s.Fields.List = append(s.Fields.List[:i], s.Fields.List[i+1:]...)
						return false
					}
				}
			}
		}
		return true
	})

	if len(structType.Fields.List) == 0 {
		_ = NewImport(a.ImportPath).Rollback(file)
	}

	if a.Type == TypePluginServiceEnter {
		return nil
	}

	//回滚变量内容
	ast.Inspect(file, func(n ast.Node) bool {
		genDecl, ok := n.(*ast.GenDecl)
		if ok && genDecl.Tok == token.VAR {
			for i, spec := range genDecl.Specs {
				valueSpec, vsok := spec.(*ast.ValueSpec)
				if vsok {
					for _, name := range valueSpec.Names {
						if name.Name == a.ModuleName {
							genDecl.Specs = append(genDecl.Specs[:i], genDecl.Specs[i+1:]...)
							return false
						}
					}
				}
			}
		}
		return true
	})

	return nil
}

func (a *PluginEnter) Injection(file *ast.File) error {
	_ = NewImport(a.ImportPath).Injection(file)

	has := false
	hasVar := false
	var firstStruct *ast.StructType
	var varSpec *ast.GenDecl
	//寻找是否存在结构且定位
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

	//寻找是否存在变量且定位
	ast.Inspect(file, func(n ast.Node) bool {
		genDecl, ok := n.(*ast.GenDecl)
		if ok && genDecl.Tok == token.VAR {
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
