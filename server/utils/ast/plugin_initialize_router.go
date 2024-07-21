package ast

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go/ast"
	"go/token"
	"io"
)

// PluginInitializeRouter 插件初始化路由
// PackageName.AppName.GroupName.FunctionName()
type PluginInitializeRouter struct {
	Base
	Type                 Type   // 类型
	Path                 string // 文件路径
	ImportPath           string // 导包路径
	ImportGlobalPath     string // 导包全局变量路径
	ImportMiddlewarePath string // 导包中间件路径
	RelativePath         string // 相对路径
	AppName              string // 应用名称
	GroupName            string // 分组名称
	PackageName          string // 包名
	FunctionName         string // 函数名
	LeftRouterGroupName  string // 左路由分组名称
	RightRouterGroupName string // 右路由分组名称
}

func (a *PluginInitializeRouter) Parse(filename string, writer io.Writer) (file *ast.File, err error) {
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

func (a *PluginInitializeRouter) Rollback(file *ast.File) error {
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
							v5, o5 := v4.X.(*ast.SelectorExpr)
							if o5 {
								v6, o6 := v5.X.(*ast.SelectorExpr)
								if o6 {
									v7, o7 := v6.X.(*ast.Ident)
									if o7 {
										if v7.Name == a.PackageName && v6.Sel.Name == a.AppName && v5.Sel.Name == a.GroupName && v4.Sel.Name == a.FunctionName {
											v1.Body.List = append(v1.Body.List[:j], v1.Body.List[j+1:]...)
											if len(v1.Body.List) == 3 {
												if a.ImportMiddlewarePath != "" {
													_ = NewImport(a.ImportMiddlewarePath).Rollback(file)
												}
												if a.ImportGlobalPath != "" {
													_ = NewImport(a.ImportGlobalPath).Rollback(file)
												}
												_ = NewImport(a.ImportPath).Rollback(file)
												v1.Body.List = nil
											}
											break
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
	return nil
}

func (a *PluginInitializeRouter) Injection(file *ast.File) error {
	_ = NewImport(a.ImportPath).Injection(file)
	for i := 0; i < len(file.Decls); i++ {
		v1, o1 := file.Decls[i].(*ast.FuncDecl)
		if o1 {
			var has bool
			for j := 0; j < len(v1.Body.List); j++ {
				v2, o2 := v1.Body.List[j].(*ast.ExprStmt)
				if o2 {
					v3, o3 := v2.X.(*ast.CallExpr)
					if o3 {
						v4, o4 := v3.Fun.(*ast.SelectorExpr)
						if o4 {
							v5, o5 := v4.X.(*ast.SelectorExpr)
							if o5 {
								v6, o6 := v5.X.(*ast.SelectorExpr)
								if o6 {
									v7, o7 := v6.X.(*ast.Ident)
									if o7 {
										if v7.Name == a.PackageName && v6.Sel.Name == a.AppName && v5.Sel.Name == a.GroupName && v4.Sel.Name == a.FunctionName {
											has = true
										}
									}
								}
							}
						}
					}
				}
			}
			if !has {
				if v1.Body == nil {
					v1.Body = new(ast.BlockStmt)
				}
				if v1.Body.List == nil {
					a.ImportMiddlewarePath = fmt.Sprintf(`"%s/middleware"`, global.GVA_CONFIG.AutoCode.Module)
					_ = NewImport(a.ImportMiddlewarePath).Injection(file)
					a.ImportGlobalPath = fmt.Sprintf(`"%s/global"`, global.GVA_CONFIG.AutoCode.Module)
					_ = NewImport(a.ImportGlobalPath).Injection(file)
					v1.Body.List = make([]ast.Stmt, 0, 3)
					public := &ast.AssignStmt{
						Lhs: []ast.Expr{
							&ast.Ident{Name: a.LeftRouterGroupName, Obj: ast.NewObj(ast.Var, a.LeftRouterGroupName)},
						},
						Tok: token.DEFINE,
						Rhs: []ast.Expr{
							&ast.CallExpr{
								Fun: &ast.SelectorExpr{
									X: &ast.CallExpr{
										Fun: &ast.SelectorExpr{
											X:   &ast.Ident{Name: "engine"},
											Sel: &ast.Ident{Name: "Group"},
										},
										Args: []ast.Expr{
											&ast.BasicLit{
												Kind:  token.STRING,
												Value: `global.GVA_CONFIG.System.RouterPrefix`,
											},
										},
									},
									Sel: &ast.Ident{Name: "Group"},
								},
								Args: []ast.Expr{
									&ast.BasicLit{
										Kind:  token.STRING,
										Value: `""`,
									},
								},
							},
						},
					}
					private := &ast.AssignStmt{
						Lhs: []ast.Expr{
							&ast.Ident{Name: a.RightRouterGroupName, Obj: ast.NewObj(ast.Var, a.RightRouterGroupName)},
						},
						Tok: token.DEFINE,
						Rhs: []ast.Expr{
							&ast.CallExpr{
								Fun: &ast.SelectorExpr{
									X: &ast.CallExpr{
										Fun: &ast.SelectorExpr{
											X:   &ast.Ident{Name: "engine"},
											Sel: &ast.Ident{Name: "Group"},
										},
										Args: []ast.Expr{
											&ast.BasicLit{
												Kind:  token.STRING,
												Value: `global.GVA_CONFIG.System.RouterPrefix`,
											},
										},
									},
									Sel: &ast.Ident{Name: "Group"},
								},
								Args: []ast.Expr{
									&ast.BasicLit{
										Kind:  token.STRING,
										Value: `""`,
									},
								},
							},
						},
					}
					useMiddleware := &ast.ExprStmt{
						X: &ast.CallExpr{
							Fun: &ast.SelectorExpr{
								X: &ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X:   &ast.Ident{Name: a.RightRouterGroupName},
										Sel: &ast.Ident{Name: "Use"},
									},
									Args: []ast.Expr{
										&ast.CallExpr{
											Fun: &ast.SelectorExpr{
												X:   &ast.Ident{Name: "middleware"},
												Sel: &ast.Ident{Name: "JWTAuth"},
											},
										},
									},
								},
								Sel: &ast.Ident{Name: "Use"},
							},
							Args: []ast.Expr{
								&ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X:   &ast.Ident{Name: "middleware"},
										Sel: &ast.Ident{Name: "CasbinHandler"},
									},
								},
							},
						},
					}
					v1.Body.List = append(v1.Body.List, public)
					v1.Body.List = append(v1.Body.List, private)
					v1.Body.List = append(v1.Body.List, useMiddleware)
				}
				body := &ast.ExprStmt{
					X: &ast.CallExpr{
						Fun: &ast.SelectorExpr{
							X: &ast.SelectorExpr{
								X: &ast.SelectorExpr{
									X:   &ast.Ident{Name: a.PackageName},
									Sel: &ast.Ident{Name: a.AppName},
								},
								Sel: &ast.Ident{Name: a.GroupName},
							},
							Sel: &ast.Ident{Name: a.FunctionName},
						},
						Args: []ast.Expr{
							&ast.Ident{Name: a.LeftRouterGroupName},
							&ast.Ident{Name: a.RightRouterGroupName},
						},
					},
				}
				v1.Body.List = append(v1.Body.List, body)
			}
		}
	}
	return nil
}

func (a *PluginInitializeRouter) Format(filename string, writer io.Writer, file *ast.File) error {
	if filename == "" {
		filename = a.Path
	}
	return a.Base.Format(filename, writer, file)
}
