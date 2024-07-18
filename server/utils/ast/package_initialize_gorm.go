package ast

import (
	"go/ast"
	"go/token"
	"io"
)

// PackageInitializeGorm 包初始化gorm
type PackageInitializeGorm struct {
	Base
	Type         Type   // 类型
	Path         string // 文件路径
	ImportPath   string // 导包路径
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
				v2, o2 := v1.Body.List[j].(*ast.AssignStmt)
				if o2 {
					if v2.Tok != token.DEFINE {
						break
					}
					for k := 0; k < len(v2.Rhs); k++ {
						v3, o3 := v2.Rhs[k].(*ast.CallExpr)
						if o3 {
							v4, o4 := v3.Fun.(*ast.SelectorExpr)
							if o4 {
								if v4.Sel.Name == "AutoMigrate" {
									var needImport bool
									for l := 0; l < len(v3.Args); l++ {
										if a.IsNew {
											v5, o5 := v3.Args[l].(*ast.CallExpr)
											if o5 {
												for m := 0; m < len(v5.Args); m++ {
													v6, o6 := v5.Args[m].(*ast.SelectorExpr)
													if o6 {
														v7, o7 := v6.X.(*ast.Ident)
														if o7 {
															if v7.Name == a.PackageName {
																needImport = true
															}
															if v7.Name == a.PackageName && v6.Sel.Name == a.StructName {
																v3.Args = append(v3.Args[:l], v3.Args[l+1:]...)
																break
															}
														}
													}
												}
											}
											continue
										}
										v5, o5 := v3.Args[l].(*ast.UnaryExpr)
										if o5 {
											if v5.Op != token.AND {
												continue
											}
											v6, o6 := v5.X.(*ast.CompositeLit)
											if o6 {
												v7, o7 := v6.Type.(*ast.SelectorExpr)
												if o7 {
													v8, o8 := v7.X.(*ast.Ident)
													if o8 {
														if v8.Name == a.PackageName {
															needImport = true
														}
														if v8.Name == a.PackageName && v7.Sel.Name == a.StructName {
															v3.Args = append(v3.Args[:l], v3.Args[l+1:]...)
															break
														}
													}
												}
											}
										}
									}
									if !needImport {
										NewImport(a.ImportPath).Rollback(file)
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
			for j := 0; j < len(v1.Body.List); j++ {
				v2, o2 := v1.Body.List[j].(*ast.AssignStmt)
				if o2 {
					if v2.Tok != token.DEFINE {
						break
					}
					for k := 0; k < len(v2.Rhs); k++ {
						v3, o3 := v2.Rhs[k].(*ast.CallExpr)
						if o3 {
							v4, o4 := v3.Fun.(*ast.SelectorExpr)
							if o4 {
								if v4.Sel.Name == "AutoMigrate" {
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
										v5, o5 := v3.Args[l].(*ast.UnaryExpr)
										if o5 {
											if v5.Op != token.AND {
												continue
											}
											v6, o6 := v5.X.(*ast.CompositeLit)
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
										arg := &ast.UnaryExpr{
											Op: token.AND,
											X: &ast.CompositeLit{
												Type: &ast.SelectorExpr{
													X:   &ast.Ident{Name: a.PackageName},
													Sel: &ast.Ident{Name: a.StructName},
												},
											},
										}
										basicLit := &ast.BasicLit{
											Kind:  token.STRING,
											Value: "\n\t\t",
										}
										v3.Args = append(v3.Args, basicLit)
										v3.Args = append(v3.Args, arg)
										basicLit = &ast.BasicLit{
											Kind:  token.STRING,
											Value: "\n",
										}
										v3.Args = append(v3.Args, basicLit)
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

func (a *PackageInitializeGorm) Format(filename string, writer io.Writer, file *ast.File) error {
	if filename == "" {
		filename = a.Path
	}
	return a.Base.Format(filename, writer, file)
}
