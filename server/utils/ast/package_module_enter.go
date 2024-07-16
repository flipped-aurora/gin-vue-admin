package ast

import (
	"github.com/pkg/errors"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

// PackageModuleEnter 模块化入口
// ModuleName := PackageName.AppName.GroupName.ServiceName
type PackageModuleEnter struct {
	Type         Type   // 类型
	Path         string // 文件路径
	ImportPath   string // 导包路径
	StructName   string // 结构体名称
	AppName      string // 应用名称
	GroupName    string // 分组名称
	ModuleName   string // 模块名称
	PackageName  string // 包名
	PreviewPath  string // 预览文件路径
	ServiceName  string // 服务名称
	TemplatePath string // 模板路径
}

func (a *PackageModuleEnter) Rollback() error {
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, a.Path, nil, parser.ParseComments)
	if err != nil {
		return errors.Wrapf(err, "[filepath:%s]打开文件失败!", a.Path)
	}
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
				if len(v1.Specs) == 0 {
					err = NewImport(file, a.ImportPath).Rollback()
					if err != nil {
						return err
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

func (a *PackageModuleEnter) Injection() error {
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
				if a.Type == TypePackageServiceModuleEnter {
					continue
				}
				v3, o3 := v1.Specs[j].(*ast.ValueSpec)
				if o3 {
					var hasValue bool
					if len(v3.Names) == 1 && v3.Names[0].Name == a.ModuleName {
						hasValue = true
					}
					if !hasValue {
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
					}
				}
			}
		}
	}
	var create *os.File
	path := a.Path
	if a.PreviewPath != "" {
		path = a.PreviewPath
	}
	err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return errors.Wrapf(err, "[filepath:%s]创建文件夹失败!", path)
	}
	create, err = os.Create(path)
	if err != nil {
		return errors.Wrapf(err, "[filepath:%s]创建文件失败!", path)
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
