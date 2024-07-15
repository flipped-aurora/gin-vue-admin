package ast

import (
	"github.com/pkg/errors"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
)

type PluginGenModel struct {
	Type        Type   // 类型
	Path        string // 文件路径
	ImportPath  string // 导包路径
	PackageName string // 包名
	StructName  string // 结构体名称
	IsNew       bool   // 是否使用new关键字
}

func (a *PluginGenModel) Rollback() error {
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, a.Path, nil, parser.ParseComments)
	if err != nil {
		return errors.Wrapf(err, "[filepath:%s]打开文件失败!", a.Path)
	}
	for i := 0; i < len(file.Decls); i++ {
		v1, o1 := file.Decls[i].(*ast.FuncDecl)
		if o1 {
			for j := 0; j < len(v1.Body.List); j++ {
				v2, o2 := v1.Body.List[j].(*ast.ExprStmt)
				if o2 {
					v3, o3 := v2.X.(*ast.CallExpr)
					if o3 {
						v4, o4 := v3.Fun.(*ast.SelectorExpr)
						if o4 {
							if v4.Sel.Name != "ApplyBasic" {
								continue
							}
							for k := 0; k < len(v3.Args); k++ {
								v5, o5 := v3.Args[k].(*ast.CallExpr)
								if o5 {
									v6, o6 := v5.Fun.(*ast.Ident)
									if o6 {
										if v6.Name != "new" {
											continue
										}
										for l := 0; l < len(v5.Args); l++ {
											v7, o7 := v5.Args[l].(*ast.SelectorExpr)
											if o7 {
												v8, o8 := v7.X.(*ast.Ident)
												if o8 {
													if v8.Name == a.PackageName && v7.Sel.Name == a.StructName {
														v3.Args = append(v3.Args[:k], v3.Args[k+1:]...)
														continue
													}
												}
											}
										}
									}
								}
								if k >= len(v3.Args) {
									break
								}
								v6, o6 := v3.Args[k].(*ast.CompositeLit)
								if o6 {
									v7, o7 := v6.Type.(*ast.SelectorExpr)
									if o7 {
										v8, o8 := v7.X.(*ast.Ident)
										if o8 {
											if v8.Name == a.PackageName && v7.Sel.Name == a.StructName {
												v3.Args = append(v3.Args[:k], v3.Args[k+1:]...)
												continue
											}
										}
									}
								}
							}
							if len(v3.Args) == 0 {
								err = NewImport(file, a.ImportPath).Rollback()
								if err != nil {
									return err
								}
							}
						}
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

func (a *PluginGenModel) Injection() error {
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
		v1, o1 := file.Decls[i].(*ast.FuncDecl)
		if o1 {
			for j := 0; j < len(v1.Body.List); j++ {
				v2, o2 := v1.Body.List[j].(*ast.ExprStmt)
				if o2 {
					v3, o3 := v2.X.(*ast.CallExpr)
					if o3 {
						v4, o4 := v3.Fun.(*ast.SelectorExpr)
						if o4 {
							if v4.Sel.Name != "ApplyBasic" {
								continue
							}
							var has bool
							for k := 0; k < len(v3.Args); k++ {
								v5, o5 := v3.Args[k].(*ast.CallExpr)
								if o5 {
									v6, o6 := v5.Fun.(*ast.Ident)
									if o6 {
										if v6.Name != "new" {
											continue
										}
										for l := 0; l < len(v5.Args); l++ {
											v7, o7 := v5.Args[l].(*ast.SelectorExpr)
											if o7 {
												v8, o8 := v7.X.(*ast.Ident)
												if o8 {
													if v8.Name == a.PackageName && v7.Sel.Name == a.StructName {
														has = true
														break
													}
												}
											}
										}
									}
								}
								v6, o6 := v3.Args[k].(*ast.CompositeLit)
								if o6 {
									v7, o7 := v6.Type.(*ast.SelectorExpr)
									if o7 {
										v8, o8 := v7.X.(*ast.Ident)
										if o8 {
											if v8.Name == a.PackageName && v7.Sel.Name == a.StructName {
												has = true
												break
											}
										}
									}
								}
							}
							if !has {
								if a.IsNew {
									arg := &ast.CallExpr{
										Fun: &ast.Ident{Name: "\n\t\tnew"},
										Args: []ast.Expr{
											&ast.SelectorExpr{
												X:   &ast.Ident{Name: a.PackageName},
												Sel: &ast.Ident{Name: a.StructName},
											},
										},
									}
									v3.Args = append(v3.Args, arg)
									v3.Args = append(v3.Args, &ast.BasicLit{
										Kind:  token.STRING,
										Value: "\n",
									})
									break
								}
								arg := &ast.CompositeLit{
									Type: &ast.SelectorExpr{
										X:   &ast.Ident{Name: a.PackageName},
										Sel: &ast.Ident{Name: a.StructName},
									},
								}
								v3.Args = append(v3.Args, arg)
							}
						}
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
