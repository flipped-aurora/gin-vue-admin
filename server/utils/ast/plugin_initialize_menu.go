package ast

import (
	"github.com/pkg/errors"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
)

type PluginInitializeMenu struct {
	Type          Type   // 类型
	Path          string // 文件路径
	MenuSort      string // 菜单排序
	MenuPath      string // 菜单路径
	MenuName      string // 菜单名称
	MenuMetaIcon  string // 菜单图标
	MenuMetaTitle string // 菜单标题
	MenuComponent string // 菜单组件
}

func (a *PluginInitializeMenu) Rollback() error {
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, a.Path, nil, parser.ParseComments)
	if err != nil {
		return errors.Wrapf(err, "[filepath:%s]打开文件失败!", a.Path)
	}
	create, err := os.Create(a.Path)
	if err != nil {
		return errors.Wrapf(err, "[filepath:%s]打开文件失败!", a.Path)
	}
	for i := 0; i < len(file.Decls); i++ {
		v1, o1 := file.Decls[i].(*ast.FuncDecl)
		if o1 {
			if v1.Name.Name != "Menu" {
				continue
			}
			for j := 0; j < len(v1.Body.List); j++ {
				v2, o2 := v1.Body.List[j].(*ast.AssignStmt)
				if o2 {
					if len(v2.Lhs) >= 1 {
						v3, o3 := v2.Lhs[0].(*ast.Ident)
						if o3 {
							if v3.Name != "entities" {
								break
							}
						}
					}
					for k := 0; k < len(v2.Rhs); k++ {
						v3, o3 := v2.Rhs[k].(*ast.CompositeLit)
						if o3 {
							v4, o4 := v3.Type.(*ast.ArrayType)
							if o4 {
								v5, o5 := v4.Elt.(*ast.SelectorExpr)
								if o5 {
									v6, o6 := v5.X.(*ast.Ident)
									if o6 {
										if v6.Name != "model" {
											break
										}
										if v5.Sel.Name != "SysBaseMenu" {
											break
										}
									}
								}
							}
							for l := 0; l < len(v3.Elts); l++ {
								v5, o5 := v3.Elts[l].(*ast.CompositeLit)
								if o5 {
									for m := 0; m < len(v5.Elts); m++ {
										v6, o6 := v5.Elts[m].(*ast.KeyValueExpr)
										if o6 {
											v7, o7 := v6.Key.(*ast.Ident)
											if o7 {
												if v7.Name == "Path" {
													v8, o8 := v6.Value.(*ast.BasicLit)
													if o8 {
														if v8.Value == a.MenuPath {
															v3.Elts = append(v3.Elts[:l], v3.Elts[l+1:]...)
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
				}
			}
		}
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

func (a *PluginInitializeMenu) Injection() error {
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, a.Path, nil, parser.ParseComments)
	if err != nil {
		return errors.Wrapf(err, "[filepath:%s]打开文件失败!", a.Path)
	}
	for i := 0; i < len(file.Decls); i++ {
		v1, o1 := file.Decls[i].(*ast.FuncDecl)
		if o1 {
			if v1.Name.Name != "Menu" {
				continue
			}
			for j := 0; j < len(v1.Body.List); j++ {
				v2, o2 := v1.Body.List[j].(*ast.AssignStmt)
				if o2 {
					if len(v2.Lhs) >= 1 {
						v3, o3 := v2.Lhs[0].(*ast.Ident)
						if o3 {
							if v3.Name != "entities" {
								break
							}
						}
					}
					for k := 0; k < len(v2.Rhs); k++ {
						v3, o3 := v2.Rhs[k].(*ast.CompositeLit)
						if o3 {
							v4, o4 := v3.Type.(*ast.ArrayType)
							if o4 {
								v5, o5 := v4.Elt.(*ast.SelectorExpr)
								if o5 {
									v6, o6 := v5.X.(*ast.Ident)
									if o6 {
										if v6.Name != "model" && v5.Sel.Name != "SysBaseMenu" {
											break
										}
									}
								}
							}
							var has bool
							for l := 0; l < len(v3.Elts); l++ {
								v5, o5 := v3.Elts[l].(*ast.CompositeLit)
								if o5 {
									for m := 0; m < len(v5.Elts); m++ {
										v6, o6 := v5.Elts[m].(*ast.KeyValueExpr)
										if o6 {
											v7, o7 := v6.Key.(*ast.Ident)
											if o7 {
												if v7.Name == "Path" {
													v8, o8 := v6.Value.(*ast.BasicLit)
													if o8 {
														if v8.Value == a.MenuPath {
															has = true
															continue
														}
													}
												}
											}
										}
									}
								}
							}
							if !has {
								// TODO 需要加回车和两个tab缩进
								elt := &ast.CompositeLit{
									Elts: []ast.Expr{
										&ast.KeyValueExpr{Key: ast.NewIdent("Path"), Value: &ast.BasicLit{Kind: token.STRING, Value: a.MenuPath}},
										&ast.KeyValueExpr{Key: ast.NewIdent("Name"), Value: &ast.BasicLit{Kind: token.STRING, Value: a.MenuName}},
										&ast.KeyValueExpr{Key: ast.NewIdent("Component"), Value: &ast.BasicLit{Kind: token.STRING, Value: a.MenuComponent}},
										&ast.KeyValueExpr{Key: ast.NewIdent("Sort"), Value: &ast.BasicLit{Kind: token.STRING, Value: a.MenuSort}},
										&ast.KeyValueExpr{Key: ast.NewIdent("Meta"), Value: &ast.CompositeLit{
											Type: &ast.SelectorExpr{
												X:   ast.NewIdent("model"),
												Sel: ast.NewIdent("Meta"),
											},
											Elts: []ast.Expr{
												&ast.KeyValueExpr{Key: ast.NewIdent("Icon"), Value: &ast.BasicLit{Kind: token.STRING, Value: a.MenuMetaIcon}},
												&ast.KeyValueExpr{Key: ast.NewIdent("Title"), Value: &ast.BasicLit{Kind: token.STRING, Value: a.MenuMetaTitle}},
											},
										},
										},
									},
								}
								v3.Elts = append(v3.Elts, elt)
								if len(v3.Elts) == 1 {
									// TODO 需要加回车
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
