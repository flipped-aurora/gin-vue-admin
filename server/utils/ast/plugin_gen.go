package ast

import (
	"go/ast"
	"go/token"
	"io"
)

type PluginGen struct {
	Base
	Type         Type   // 类型
	Path         string // 文件路径
	ImportPath   string // 导包路径
	RelativePath string // 相对路径
	StructName   string // 结构体名称
	PackageName  string // 包名
	IsNew        bool   // 是否使用new关键字
}

func (a *PluginGen) Parse(filename string, writer io.Writer) (file *ast.File, err error) {
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
func (a *PluginGen) Rollback(file *ast.File) error {
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
								_ = NewImport(a.ImportPath).Rollback(file)
							}
						}
					}
				}
			}
		}
	}
	return nil
}

func (a *PluginGen) Injection(file *ast.File) error {
	_ = NewImport(a.ImportPath).Injection(file)
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
	return nil
}

func (a *PluginGen) Format(filename string, writer io.Writer, file *ast.File) error {
	if filename == "" {
		filename = a.Path
	}
	return a.Base.Format(filename, writer, file)
}
