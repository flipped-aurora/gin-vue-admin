package ast

import (
	"github.com/pkg/errors"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
)

// PluginInitializeRouter 插件初始化路由
// PackageName.AppName.GroupName.FunctionName()
type PluginInitializeRouter struct {
	Type                 Type   // 类型
	Path                 string // 文件路径
	ImportPath           string // 导包路径
	AppName              string // 应用名称
	GroupName            string // 分组名称
	PackageName          string // 包名
	FunctionName         string // 函数名
	LeftRouterGroupName  string // 左路由分组名称
	RightRouterGroupName string // 右路由分组名称
}

func (a *PluginInitializeRouter) Rollback() error {
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
							v5, o5 := v4.X.(*ast.SelectorExpr)
							if o5 {
								v6, o6 := v5.X.(*ast.SelectorExpr)
								if o6 {
									v7, o7 := v6.X.(*ast.Ident)
									if o7 {
										if v7.Name == a.PackageName && v6.Sel.Name == a.AppName && v5.Sel.Name == a.GroupName && v4.Sel.Name == a.FunctionName {
											v1.Body.List = append(v1.Body.List[:j], v1.Body.List[j+1:]...)
											if len(v1.Body.List) >= 2 {
												err = NewImport(file, a.ImportPath).Rollback()
												if err != nil {
													return err
												}
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

func (a *PluginInitializeRouter) Injection() error {
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
					v1.Body.List = make([]ast.Stmt, 0, 3)
					public := &ast.AssignStmt{
						Lhs: []ast.Expr{
							&ast.Ident{Name: a.LeftRouterGroupName, Obj: ast.NewObj(ast.Var, a.LeftRouterGroupName)},
						},
						Tok: token.DEFINE,
						Rhs: []ast.Expr{
							&ast.CallExpr{
								Fun: &ast.SelectorExpr{
									X:   &ast.Ident{Name: "engine"},
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
									X:   &ast.Ident{Name: "engine"},
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
					v1.Body.List = append(v1.Body.List, public)
					v1.Body.List = append(v1.Body.List, private)
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