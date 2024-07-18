package ast

import (
	"go/ast"
	"go/token"
	"io"
)

// PackageInitializeRouter 包初始化路由
// ModuleName := PackageName.AppName.GroupName
// ModuleName.FunctionName(RouterGroupName)
type PackageInitializeRouter struct {
	Base
	Type                 Type   // 类型
	Path                 string // 文件路径
	ImportPath           string // 导包路径
	RelativePath         string // 相对路径
	AppName              string // 应用名称
	GroupName            string // 分组名称
	ModuleName           string // 模块名称
	PackageName          string // 包名
	FunctionName         string // 函数名
	RouterGroupName      string // 路由分组名称
	LeftRouterGroupName  string // 左路由分组名称
	RightRouterGroupName string // 右路由分组名称
}

func (a *PackageInitializeRouter) Parse(filename string, writer io.Writer) (file *ast.File, err error) {
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

func (a *PackageInitializeRouter) Rollback(file *ast.File) {
	for i := 0; i < len(file.Decls); i++ {
		v1, o1 := file.Decls[i].(*ast.FuncDecl)
		if o1 {
			if v1.Name.Name != "initBizRouter" {
				continue
			}
			for j := 0; j < len(v1.Body.List); j++ {
				v2, o2 := v1.Body.List[j].(*ast.BlockStmt)
				if o2 {
					for k := 0; k < len(v2.List); k++ {
						if k == 0 {
							v3, o3 := v2.List[0].(*ast.AssignStmt)
							if o3 {
								if len(v3.Lhs) == 1 && len(v3.Rhs) == 1 {
									v4, o4 := v3.Lhs[0].(*ast.Ident)
									v5, o5 := v3.Rhs[0].(*ast.SelectorExpr)
									v6, o6 := v5.X.(*ast.SelectorExpr)
									v7, o7 := v6.X.(*ast.Ident)
									if o4 && o5 && o6 && o7 {
										if v4.Name != a.ModuleName && v7.Name != a.PackageName && v6.Sel.Name != a.AppName && v5.Sel.Name != a.GroupName {
											break
										}
									}
								}
							}
						} // 判断是否有路由组和作用域
						v3, o3 := v2.List[k].(*ast.ExprStmt)
						if o3 {
							v4, o4 := v3.X.(*ast.CallExpr)
							if o4 {
								v5, o5 := v4.Fun.(*ast.SelectorExpr)
								if o5 {
									v6, o6 := v5.X.(*ast.Ident)
									if o6 {
										if v6.Name == a.ModuleName && v5.Sel.Name == a.FunctionName {
											v2.List = append(v2.List[:k], v2.List[k+1:]...)
											length := len(v2.List)
											if length == 1 {
												v1.Body.List = append(v1.Body.List[:j], v1.Body.List[j+1:]...)
												// TODO 删除作用域之后会出现两种情况需要删除空行, 中间模块被删除和最后的模块被删除
												// if j < len(v1.Body.List) {
												// 	v2, o2 = v1.Body.List[j].(*ast.BlockStmt)
												// 	if o2 {
												// 		v2.Lbrace -= 3
												// 		// v2.Rbrace -= 2
												// 	}
												// } // 中间模块被删除
												// if j == len(v1.Body.List) {
												// 	v1.Body.Rbrace -= 10
												// } // 最后的模块被删除
												break
											} // 无调用路由初始化函数 => 删除局部变量 && 删除作用域 && 导包
											if k < length-1 {
												v3, o3 = v2.List[k].(*ast.ExprStmt)
												if o3 {
													v4, o4 = v3.X.(*ast.CallExpr)
													if o4 {
														v5, o5 = v4.Fun.(*ast.SelectorExpr)
														if o5 {
															v6, o6 = v5.X.(*ast.Ident)
															if o6 {
																v6.NamePos -= 10
															}
															v5.Sel.NamePos -= 20
														}
													}
												}
											} // 删除空行
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

func (a *PackageInitializeRouter) Injection(file *ast.File) {
	NewImport(a.ImportPath).Injection(file)
	for i := 0; i < len(file.Decls); i++ {
		v1, o1 := file.Decls[i].(*ast.FuncDecl)
		if o1 {
			if v1.Name.Name != "initBizRouter" {
				continue
			}
			var (
				hasInit   bool
				blockStmt *ast.BlockStmt
			)
			for j := 0; j < len(v1.Body.List); j++ {
				v2, o2 := v1.Body.List[j].(*ast.BlockStmt)
				if o2 {
					for k := 0; k < len(v2.List); k++ {
						if k == 0 {
							v3, o3 := v2.List[0].(*ast.AssignStmt)
							if o3 {
								if len(v3.Lhs) == 1 && len(v3.Rhs) == 1 {
									v4, o4 := v3.Lhs[0].(*ast.Ident)
									v5, o5 := v3.Rhs[0].(*ast.SelectorExpr)
									v6, o6 := v5.X.(*ast.SelectorExpr)
									v7, o7 := v6.X.(*ast.Ident)
									if o4 && o5 && o6 && o7 {
										if v4.Name == a.ModuleName && v7.Name == a.PackageName && v6.Sel.Name == a.AppName && v5.Sel.Name == a.GroupName {
											blockStmt = v2
											continue
										}
									}
								}
							}
						} // 判断是否有路由组和作用域
						v3, o3 := v2.List[k].(*ast.ExprStmt)
						if o3 {
							v4, o4 := v3.X.(*ast.CallExpr)
							if o4 {
								v5, o5 := v4.Fun.(*ast.SelectorExpr)
								if o5 {
									v6, o6 := v5.X.(*ast.Ident)
									if o6 {
										if v6.Name == a.ModuleName && v5.Sel.Name == a.FunctionName {
											hasInit = true
											continue
										} // 判断是否存在调用函数
									}
								}
							}
						}
					}
				}
			}
			if hasInit {
				continue
			}
			stmt := &ast.ExprStmt{
				X: &ast.CallExpr{
					Fun: &ast.SelectorExpr{
						X:   &ast.Ident{Name: a.ModuleName},
						Sel: &ast.Ident{Name: a.FunctionName},
					},
					Args: []ast.Expr{
						&ast.Ident{Name: a.LeftRouterGroupName},
						&ast.Ident{Name: a.RightRouterGroupName},
					},
				},
			}
			if blockStmt == nil {
				blockStmt = &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Lhs: []ast.Expr{&ast.Ident{Name: a.ModuleName, Obj: ast.NewObj(ast.Var, a.ModuleName)}},
							Tok: token.DEFINE,
							Rhs: []ast.Expr{
								&ast.SelectorExpr{
									X: &ast.SelectorExpr{
										X:   &ast.Ident{Name: a.PackageName},
										Sel: &ast.Ident{Name: a.AppName},
									},
									Sel: &ast.Ident{Name: a.GroupName},
								},
							},
						},
						stmt,
					},
				}
				v1.Body.List = append(v1.Body.List, blockStmt)
				break
			} // 创建作用域 && 创建路由组 && 调用路由初始化函数
			blockStmt.List = append(blockStmt.List, stmt)
		}
	}
}

func (a *PackageInitializeRouter) Format(filename string, writer io.Writer, file *ast.File) error {
	if filename == "" {
		filename = a.Path
	}
	return a.Base.Format(filename, writer, file)
}
