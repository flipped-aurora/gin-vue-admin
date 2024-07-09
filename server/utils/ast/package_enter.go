package ast

import (
	"encoding/json"
	"github.com/pkg/errors"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
)

// PackageEnter .
// 以 server/api/v1/enter.go的SystemApiGroup为例
// enter := NewPackageEnter("server/api/v1/enter.go", "github.com/flipped-aurora/gin-vue-admin/server/api/v1/system", "SystemApiGroup", "system", "ApiGroup")
type PackageEnter struct {
	Type              Type   // 类型
	Path              string // 文件路径
	ImportPath        string // 导包路径
	StructName        string // 结构体名称
	PackageName       string // 包名
	PackageStructName string // 包结构体名称
}

func NewPackageEnter(path, importPath, structName, packageName, packageStructName string) Ast {
	return &PackageEnter{
		Path:              path,
		ImportPath:        importPath,
		StructName:        structName,
		PackageName:       packageName,
		PackageStructName: packageStructName,
	}
}

func NewPackageApiEnterField(data string) Ast {
	packageApiEnter := &PackageEnter{}
	err := json.Unmarshal([]byte(data), packageApiEnter)
	if err != nil {
		return nil
	}
	return packageApiEnter
}

func (a *PackageEnter) Rollback() error {
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, a.Path, nil, parser.ParseComments)
	if err != nil {
		return errors.Wrapf(err, "[filepath:%s]打开文件失败!", a.Path)
	}
	err = NewImport(file, a.ImportPath).Rollback()
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
						for k := 0; k < len(v3.Fields.List); k++ {
							for l := 0; l < len(v3.Fields.List[k].Names); l++ {
								if v3.Fields.List[k].Names[l].Name == a.StructName {
									v3.Fields.List = append(v3.Fields.List[:k], v3.Fields.List[k+1:]...)
									goto Next
								}
							}
						}
					}
				}
			}
		}
	}
Next:
	create, err := os.Create(a.Path)
	if err != nil {
		return errors.Wrapf(err, "[filepath:%s]打开文件失败!", a.Path)
	}
	defer func() {
		_ = create.Close()
	}()
	err = format.Node(create, fileSet, file)
	if err != nil {
		return errors.Wrapf(err, "[filepath:%s]回滚失败!", a.Path)
	}
	return nil
}

func (a *PackageEnter) Injection() error {
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
						for k := 0; k < len(v3.Fields.List); k++ {
							for l := 0; l < len(v3.Fields.List[k].Names); l++ {
								if v3.Fields.List[k].Names[l].Name == a.StructName {
									goto Next
								}
							}
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
Next:
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
