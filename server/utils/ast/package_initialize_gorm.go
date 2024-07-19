package ast

import (
	"fmt"
	"go/ast"
	"go/token"
	"io"
	"strings"
)

// PackageInitializeGorm 包初始化gorm
type PackageInitializeGorm struct {
	Base
	Type         Type   // 类型
	Path         string // 文件路径
	ImportPath   string // 导包路径
	Business     string // 业务库 gva => gva, 不要传"gva"
	StructName   string // 结构体名称
	PackageName  string // 包名
	RelativePath string // 相对路径
	IsNew        bool   // 是否使用new关键字 true: new(PackageName.StructName) false: &PackageName.StructName{}
}

func (a *PackageInitializeGorm) Parse(filename string, writer io.Writer) (file *ast.File, err error) {
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

func (a *PackageInitializeGorm) Rollback(file *ast.File) {
	for i := 0; i < len(file.Decls); i++ {
		v1, o1 := file.Decls[i].(*ast.FuncDecl)
		if o1 {
			if v1.Name.Name != "bizModel" {
				continue
			}
			for j := 0; j < len(v1.Body.List); j++ {
				_, ok := v1.Body.List[j].(*ast.IfStmt)
				if ok {
					continue
				} // if err != nil { return err }
				v2, o2 := v1.Body.List[j].(*ast.AssignStmt)
				if o2 {
					if v2.Tok != token.DEFINE && v2.Tok != token.ASSIGN {
						break
					}
					for k := 0; k < len(v2.Rhs); k++ {
						v3, o3 := v2.Rhs[k].(*ast.CallExpr)
						if o3 {
							v4, o4 := v3.Fun.(*ast.SelectorExpr)
							if o4 {
								v5, o5 := v4.X.(*ast.CallExpr)
								if o5 {
									v6, o6 := v5.Fun.(*ast.SelectorExpr)
									if o6 {
										v7, o7 := v6.X.(*ast.Ident)
										if o7 {
											if (v7.Name == "global" && v6.Sel.Name == "GVA_DB" && v4.Sel.Name == "AutoMigrate" || v7.Name == "global") && (v6.Sel.Name == "MustGetGlobalDBByDBName" && v4.Sel.Name == "AutoMigrate") {
												if a.Business != "" {
													if len(v5.Args) == 1 {
														v8, o8 := v5.Args[0].(*ast.BasicLit)
														if o8 {
															name := strings.Trim(v8.Value, "\"")
															if name != a.Business {
																break
															}
														}
													}
												}
												{
													for l := 0; l < len(v3.Args); l++ {
														if a.IsNew {
															v8, o8 := v3.Args[l].(*ast.CallExpr)
															if o8 {
																for m := 0; m < len(v8.Args); m++ {
																	v9, o9 := v8.Args[m].(*ast.SelectorExpr)
																	if o9 {
																		v10, o10 := v9.X.(*ast.Ident)
																		if o10 {
																			if v10.Name == a.PackageName && v9.Sel.Name == a.StructName {
																				v3.Args = append(v3.Args[:l], v3.Args[l+1:]...)
																			}
																		}
																	}
																}
															}
															continue
														}
														v8, o8 := v3.Args[l].(*ast.UnaryExpr)
														if o8 {
															if v8.Op != token.AND {
																continue
															}
															v9, o9 := v8.X.(*ast.CompositeLit)
															if o9 {
																v10, o10 := v9.Type.(*ast.SelectorExpr)
																if o10 {
																	v11, o11 := v10.X.(*ast.Ident)
																	if o11 {
																		if v11.Name == a.PackageName && v10.Sel.Name == a.StructName {
																			v3.Args = append(v3.Args[:l], v3.Args[l+1:]...)
																		}
																	}
																}
															}
														}
													}
												} // 判断有没有注册结构体
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func (a *PackageInitializeGorm) Injection(file *ast.File) {
	NewImport(a.ImportPath).Injection(file)
	for i := 0; i < len(file.Decls); i++ {
		v1, o1 := file.Decls[i].(*ast.FuncDecl)
		if o1 {
			if v1.Name.Name != "bizModel" {
				continue
			}
			var hasStruct bool
			var structCallExpr *ast.CallExpr
			var business *ast.CallExpr
			for j := 0; j < len(v1.Body.List); j++ {
				_, ok := v1.Body.List[j].(*ast.IfStmt)
				if ok {
					continue
				} // if err != nil { return err }
				v2, o2 := v1.Body.List[j].(*ast.AssignStmt)
				if o2 {
					if v2.Tok != token.DEFINE && v2.Tok != token.ASSIGN {
						break
					}
					for k := 0; k < len(v2.Rhs); k++ {
						v3, o3 := v2.Rhs[k].(*ast.CallExpr)
						if o3 {
							v4, o4 := v3.Fun.(*ast.SelectorExpr)
							if o4 {
								v5, o5 := v4.X.(*ast.CallExpr)
								if o5 {
									v6, o6 := v5.Fun.(*ast.SelectorExpr)
									if o6 {
										v7, o7 := v6.X.(*ast.Ident)
										if o7 {
											if (v7.Name == "global" && v6.Sel.Name == "GVA_DB" && v4.Sel.Name == "AutoMigrate") || (v7.Name == "global" && v6.Sel.Name == "MustGetGlobalDBByDBName" && v4.Sel.Name == "AutoMigrate") {
												if a.Business != "" {
													if len(v5.Args) == 1 {
														v8, o8 := v5.Args[0].(*ast.BasicLit)
														if o8 {
															name := strings.Trim(v8.Value, "\"")
															if name != a.Business {
																break
															}
															business = v3
														}
													}
												}
											}
											{
												for l := 0; l < len(v3.Args); l++ {
													if a.IsNew {
														v8, o8 := v3.Args[l].(*ast.CallExpr)
														if o8 {
															for m := 0; m < len(v8.Args); m++ {
																v9, o9 := v8.Args[m].(*ast.SelectorExpr)
																if o9 {
																	v10, o10 := v9.X.(*ast.Ident)
																	if o10 {
																		if v10.Name == a.PackageName && v9.Sel.Name == a.StructName {
																			hasStruct = true
																		}
																	}
																}
															}
														}
														continue
													}
													v8, o8 := v3.Args[l].(*ast.UnaryExpr)
													if o8 {
														if v8.Op != token.AND {
															continue
														}
														v9, o9 := v8.X.(*ast.CompositeLit)
														if o9 {
															v10, o10 := v9.Type.(*ast.SelectorExpr)
															if o10 {
																v11, o11 := v10.X.(*ast.Ident)
																if o11 {
																	if v11.Name == a.PackageName && v10.Sel.Name == a.StructName {
																		hasStruct = true
																	}
																}
															}
														}
													}
												}
												if !hasStruct {
													structCallExpr = v3
												}
											} // 判断有没有注册结构体
										}
									}
								}
							}
						}
					}
				}
			}
			basicLit := &ast.BasicLit{Kind: token.STRING, Value: "\n"}
			if !hasStruct {
				var expr ast.Expr
				if a.IsNew {
					expr = &ast.CallExpr{
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
				} else {
					expr = &ast.UnaryExpr{
						Op: token.AND,
						X: &ast.CompositeLit{
							Type: &ast.SelectorExpr{
								X:   &ast.Ident{Name: a.PackageName},
								Sel: &ast.Ident{Name: a.StructName},
							},
						},
					}
				}
				if a.Business != "" {
					if business != nil {
						business.Args = append(business.Args, expr)
						business.Args = append(business.Args, basicLit)
						break
					} // 业务库
					ifStmt := &ast.IfStmt{
						Cond: &ast.BinaryExpr{
							X:  &ast.Ident{Name: "err"},
							Op: token.NEQ,
							Y:  ast.NewIdent("nil"),
						},
						Body: &ast.BlockStmt{
							List: []ast.Stmt{
								&ast.ReturnStmt{
									Results: []ast.Expr{
										ast.NewIdent("err"),
									},
								},
							},
						},
					} // if err != nil { return err }
					businessAssignStmt := &ast.AssignStmt{
						Lhs: []ast.Expr{ast.NewIdent("err")},
						Tok: token.DEFINE,
						Rhs: []ast.Expr{
							&ast.CallExpr{
								Fun: &ast.SelectorExpr{
									X: &ast.CallExpr{
										Fun: &ast.SelectorExpr{
											X:   &ast.Ident{Name: "global"},
											Sel: &ast.Ident{Name: "MustGetGlobalDBByDBName"},
										},
										Args: []ast.Expr{
											&ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%s"`, a.Business)},
										},
									},
								},
								Args: []ast.Expr{expr},
							},
						},
					}
					v1.Body.List = append(v1.Body.List, businessAssignStmt)
					v1.Body.List = append(v1.Body.List, ifStmt)
					break
				} // 有business
				if structCallExpr != nil {
					structCallExpr.Args = append(structCallExpr.Args, expr)
				}
			}
		}
	}
}

func (a *PackageInitializeGorm) Format(filename string, writer io.Writer, file *ast.File) error {
	if filename == "" {
		filename = a.Path
	}
	return a.Base.Format(filename, writer, file)
}
