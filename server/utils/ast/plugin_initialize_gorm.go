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

type PluginInitializeGorm struct {
	Type        Type   // 类型
	Path        string // 文件路径
	ImportPath  string // 导包路径
	StructName  string // 结构体名称
	PackageName string // 包名
	PreviewPath string // 预览路径
	IsNew       bool   // 是否使用new关键字 true: new(PackageName.StructName) false: &PackageName.StructName{}
}

func (a *PluginInitializeGorm) Rollback() error {
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, a.Path, nil, parser.ParseComments)
	if err != nil {
		return errors.Wrapf(err, "[filepath:%s]打开文件失败!", a.Path)
	}

	for i := 0; i < len(file.Decls); i++ {
		v1, o1 := file.Decls[i].(*ast.FuncDecl)
		if o1 {
			if v1.Name.Name != "Gorm" {
				continue
			}
			for j := 0; j < len(v1.Body.List); j++ {
				v2, o2 := v1.Body.List[j].(*ast.AssignStmt)
				if o2 {
					for k := 0; k < len(v2.Rhs); k++ {
						v3, o3 := v2.Rhs[k].(*ast.CallExpr)
						if o3 {
							v4, o4 := v3.Fun.(*ast.SelectorExpr)
							if o4 {
								if v4.Sel.Name != "AutoMigrate" {
									continue
								}
								length := len(v3.Args)
								var removeStruct bool
								var hasImport bool
								for l := 0; l < length; l++ {
									if a.IsNew {
										v5, o5 := v3.Args[l].(*ast.CallExpr)
										if o5 {
											for m := 0; m < len(v5.Args); m++ {
												v6, o6 := v5.Args[m].(*ast.SelectorExpr)
												if o6 {
													v7, o7 := v6.X.(*ast.Ident)
													if o7 {
														if v7.Name == a.PackageName && v6.Sel.Name == a.StructName {
															v3.Args = append(v3.Args[:l], v3.Args[l+1:]...)
															length--
															removeStruct = true
															continue
														} // TODO 优化 空行未删除
													}
												}
											}
										}
										continue
									}
									if l >= length {
										break
									} // 防止越界
									v6, o6 := v3.Args[l].(*ast.UnaryExpr)
									if o6 {
										v7, o7 := v6.X.(*ast.CompositeLit)
										if o7 {
											v8, o8 := v7.Type.(*ast.SelectorExpr)
											if o8 {
												v9, o9 := v8.X.(*ast.Ident)
												if o9 {
													if v6.Op == token.AND && v9.Name == a.PackageName && v8.Sel.Name == a.StructName {
														v3.Args = append(v3.Args[:l], v3.Args[l+1:]...)
														length--
														removeStruct = true
														continue
													}
												}
											}
										}
									}
								}
								for l := 0; l < length; l++ {
									v5, o5 := v3.Args[l].(*ast.CallExpr)
									if o5 {
										for m := 0; m < len(v5.Args); m++ {
											v6, o6 := v5.Args[m].(*ast.SelectorExpr)
											if o6 {
												v7, o7 := v6.X.(*ast.Ident)
												if o7 {
													if removeStruct {
														if v7.Name == a.PackageName {
															hasImport = true
														}
													}
												}
											}
										} // 判断new关键字的package是否有其他结构体使用
									}
									v6, o6 := v3.Args[l].(*ast.UnaryExpr)
									if o6 {
										v7, o7 := v6.X.(*ast.CompositeLit)
										if o7 {
											v8, o8 := v7.Type.(*ast.SelectorExpr)
											if o8 {
												v9, o9 := v8.X.(*ast.Ident)
												if o9 {
													if v9.Name == a.PackageName {
														hasImport = true
													}
												}
											}
										}
									} // 判断&关键字的package是否有其他结构体使用
								}
								if removeStruct && !hasImport {
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

func (a *PluginInitializeGorm) Injection() error {
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
			if v1.Name.Name != "Gorm" {
				continue
			}
			for j := 0; j < len(v1.Body.List); j++ {
				v2, o2 := v1.Body.List[j].(*ast.AssignStmt)
				if o2 {
					for k := 0; k < len(v2.Rhs); k++ {
						v3, o3 := v2.Rhs[k].(*ast.CallExpr)
						if o3 {
							v4, o4 := v3.Fun.(*ast.SelectorExpr)
							if o4 {
								if v4.Sel.Name != "AutoMigrate" {
									continue
								}
								var has bool

								for l := 0; l < len(v3.Args); l++ {
									if a.IsNew {
										v5, o5 := v3.Args[l].(*ast.CallExpr)
										if o5 {
											for m := 0; m < len(v5.Args); m++ {
												v6, o6 := v5.Args[m].(*ast.SelectorExpr)
												if o6 {
													v7, o7 := v6.X.(*ast.Ident)
													if o7 {
														if v7.Name == a.PackageName && v6.Sel.Name == a.StructName {
															has = true
															break
														}
													}
												}
											}
										}
										continue
									}
									v6, o6 := v3.Args[l].(*ast.UnaryExpr)
									if o6 {
										v7, o7 := v6.X.(*ast.CompositeLit)
										if o7 {
											v8, o8 := v7.Type.(*ast.SelectorExpr)
											if o8 {
												v9, o9 := v8.X.(*ast.Ident)
												if o9 {
													if v6.Op == token.AND && v9.Name == a.PackageName && v8.Sel.Name == a.StructName {
														has = true
														break
													}
												}
											}
										}
									}
								}
								if !has {
									if a.IsNew {
										arg := &ast.CallExpr{
											Fun: &ast.Ident{
												Name: "\n\t\tnew",
											},
											Args: []ast.Expr{
												&ast.SelectorExpr{
													X:   &ast.Ident{Name: a.PackageName},
													Sel: &ast.Ident{Name: a.StructName},
												},
											},
										}
										v3.Args = append(v3.Args, arg)
										v3.Args = append(v3.Args, ast.NewIdent("\n"))
									} else { // TODO 程序可用,但是格式并没有优雅回车
										arg := &ast.UnaryExpr{
											Op: token.AND,
											X: &ast.CompositeLit{
												Type: &ast.SelectorExpr{
													X:   &ast.Ident{Name: a.PackageName},
													Sel: &ast.Ident{Name: a.StructName},
												},
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
