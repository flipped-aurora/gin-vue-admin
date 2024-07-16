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

var (
	PluginVersionV1 = `"github.com/flipped-aurora/gin-vue-admin/server/utils/plugin"`
	PluginVersionV2 = `"github.com/flipped-aurora/gin-vue-admin/server/utils/plugin/v2"`
)

// PluginInitialize 插件初始化
type PluginInitialize struct {
	Type        Type   // 类型
	Path        string // 文件路径
	PluginPath  string // 插件路径
	ImportPath  string // 导包路径
	StructName  string // 结构体名称
	PackageName string // 包名
	PreviewPath string // 预览路径
	Version     string // 版本 v1/v2
}

func NewPluginInitialize(astType Type, path string, pluginPath string, importPath string) *PluginInitialize {
	return &PluginInitialize{Type: astType, Path: path, PluginPath: pluginPath, ImportPath: importPath}
}

func (a *PluginInitialize) Rollback() error {
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, a.Path, nil, parser.ParseComments)
	if err != nil {
		return errors.Wrapf(err, "[filepath:%s]打开文件失败!", a.PluginPath)
	}
	switch a.Version {
	case "bizPluginV1":
		a.Path = filepath.Join(a.Path, "plugin_biz_v1.go")
	case "bizPluginV2":
		a.Path = filepath.Join(a.Path, "plugin_biz_v2.go")
	}
	for i := 0; i < len(file.Decls); i++ {
		v1, o1 := file.Decls[i].(*ast.FuncDecl)
		if o1 {
			if v1.Name.Name == a.Version {
				var (
					isV1 bool
					isV2 bool
				)
				for j := 0; j < len(v1.Body.List); j++ {
					v2, o2 := v1.Body.List[j].(*ast.ExprStmt)
					if o2 {
						v3, o3 := v2.X.(*ast.CallExpr)
						if o3 {
							v4, o4 := v3.Fun.(*ast.Ident)
							if o4 {
								if v4.Name == "PluginInitV1" {
									isV1 = true
								}
								if v4.Name == "PluginInitV2" {
									isV2 = true
								}
								for k := 0; k < len(v3.Args); k++ {
									v5, o5 := v3.Args[k].(*ast.Ident)
									if o5 {
										if isV2 {
											if v5.Name != "engine" {
												return errors.Errorf("[filepath:%s][function:%s][InitFunction:%s]存在错误注册插件,请手动修正后再次重试插件初始化!", a.Path, a.Version, v4.Name)
											}
										}
										if isV1 {
											if v5.Name != "private" && v5.Name != "public" {
												return errors.Errorf("[filepath:%s][function:%s][InitFunction:%s]存在错误注册插件,请手动修正后再次重试插件初始化!", a.Path, a.Version, v4.Name)
											}
										}
									}
									v6, o6 := v3.Args[k].(*ast.SelectorExpr)
									if o6 {
										v7, o7 := v6.X.(*ast.Ident)
										if o7 {
											if v7.Name == a.PackageName && v6.Sel.Name == a.StructName {
												v1.Body.List = append(v1.Body.List[:j], v1.Body.List[j+1:]...)
												err = NewImport(file, a.ImportPath).Rollback()
												if err != nil {
													return err
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

func (a *PluginInitialize) Injection() error {
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, a.PluginPath, nil, parser.ParseComments)
	if err != nil {
		return errors.Wrapf(err, "[filepath:%s]打开文件失败!", a.PluginPath)
	}
	var isPlugin bool
	for i := 0; i < len(file.Decls); i++ {
		v1, o1 := file.Decls[i].(*ast.GenDecl)
		if o1 {
			for j := 0; j < len(v1.Specs); j++ {
				v2, o2 := v1.Specs[j].(*ast.ValueSpec)
				if o2 {
					for k := 0; k < len(v2.Names); k++ {
						if v2.Names[k].Name == "_" && v2.Names[k].Obj.Kind == ast.Var {
							v3, o3 := v2.Type.(*ast.SelectorExpr)
							if o3 {
								v4, o4 := v3.X.(*ast.Ident)
								if o4 {
									if (v4.Name == "interfaces" || v4.Name == "plugin") && v3.Sel.Name == "Plugin" {
										for l := 0; l < len(v2.Values); l++ {
											v5, o5 := v2.Values[l].(*ast.CallExpr)
											if o5 {
												for m := 0; m < len(v5.Args); m++ {
													v6, o6 := v5.Args[m].(*ast.Ident)
													if o6 {
														if v6.Name == "nil" {
															isPlugin = true
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
					if isPlugin {
						for k := 0; k < len(v2.Names); k++ {
							if v2.Names[k].Name != "_" && v2.Names[k].Obj.Kind == ast.Var {
								a.StructName = v2.Names[k].Name
								a.PackageName = file.Name.Name
							}
						}
					}
				}
				v3, o3 := v1.Specs[j].(*ast.ImportSpec)
				if o3 {
					if v3.Path.Value == PluginVersionV1 {
						a.Version = "bizPluginV1"
					}
					if v3.Path.Value == PluginVersionV2 {
						a.Version = "bizPluginV2"
					}
				}
			}
		}
	}
	if !isPlugin {
		return errors.Errorf("[filepath:%s]此插件不符合插件规范命名无法自动注册,请手动注册插件!", a.PluginPath)
	}
	switch a.Version {
	case "bizPluginV1":
		a.Path = filepath.Join(a.Path, "plugin_biz_v1.go")
	case "bizPluginV2":
		a.Path = filepath.Join(a.Path, "plugin_biz_v2.go")
	default:
		return errors.Errorf("[filepath:%s]插件版本不存在,请手动注册插件初始化!", a.PluginPath)
	}
	fileSet = token.NewFileSet()
	file, err = parser.ParseFile(fileSet, a.Path, nil, parser.ParseComments)
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
			if v1.Name.Name == a.Version {
				var (
					isV1 bool
					isV2 bool
				)
				var has bool
				for j := 0; j < len(v1.Body.List); j++ {
					v2, o2 := v1.Body.List[j].(*ast.ExprStmt)
					if o2 {
						v3, o3 := v2.X.(*ast.CallExpr)
						if o3 {
							v4, o4 := v3.Fun.(*ast.Ident)
							if o4 {
								if v4.Name == "PluginInitV1" {
									isV1 = true
								}
								if v4.Name == "PluginInitV2" {
									isV2 = true
								}
								for k := 0; k < len(v3.Args); k++ {
									v5, o5 := v3.Args[k].(*ast.Ident)
									if o5 {
										if isV2 {
											if v5.Name != "engine" {
												return errors.Errorf("[filepath:%s][function:%s][InitFunction:%s]存在错误注册插件,请手动修正后再次重试插件初始化!", a.Path, a.Version, v4.Name)
											}
										}
										if isV1 {
											if v5.Name != "private" && v5.Name != "public" {
												return errors.Errorf("[filepath:%s][function:%s][InitFunction:%s]存在错误注册插件,请手动修正后再次重试插件初始化!", a.Path, a.Version, v4.Name)
											}
										}
									}
									v6, o6 := v3.Args[k].(*ast.SelectorExpr)
									if o6 {
										v7, o7 := v6.X.(*ast.Ident)
										if o7 {
											if v7.Name == a.PackageName && v6.Sel.Name == a.StructName {
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
					var functionName string
					switch {
					case isV1:
						functionName = "PluginInitV1"
						// TODO v1版本插件注册注入
					case isV2:
						functionName = "PluginInitV2"
						body := &ast.ExprStmt{
							X: &ast.CallExpr{
								Fun: &ast.Ident{
									Name: functionName,
								},
								Args: []ast.Expr{
									&ast.Ident{
										Name: "engine",
									},
									&ast.SelectorExpr{
										X: &ast.Ident{
											Name: a.PackageName,
										},
										Sel: &ast.Ident{
											Name: a.StructName,
										},
									},
								},
							},
						}
						v1.Body.List = append(v1.Body.List, body)
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
