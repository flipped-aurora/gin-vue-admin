package ast

import (
	"github.com/pkg/errors"
	"go/ast"
	"io"
)

// PluginInitializeV2 插件初始化
// TODO 重构ing
type PluginInitializeV2 struct {
	Base
	Type               Type   // 类型
	Path               string // 文件路径
	PluginPath         string // 插件路径 // gva 的注册插件的位置
	RelativePath       string // 相对路径
	PluginRelativePath string // 插件相对路径
	ImportPath         string // 导包路径
	PackageName        string // 包名(interfaces/plugin)
	StructName         string // 结构体名称
	StructCamelName    string // 结构体小驼峰名称
}

func (a *PluginInitializeV2) Parse(filename string, writer io.Writer) (file *ast.File, err error) {
	if filename == "" {
		if a.RelativePath == "" {
			filename = a.Path
			a.RelativePath = a.Base.RelativePath(a.Path)
			return a.Base.Parse(filename, writer)
		}
		if a.PluginRelativePath == "" {
			a.PluginRelativePath = a.Base.RelativePath(a.PluginPath)
		}
		a.Path = a.Base.AbsolutePath(a.RelativePath)
		filename = a.Path
	}
	return a.Base.Parse(filename, writer)
}

func (a *PluginInitializeV2) Rollback(file *ast.File) error {
	return nil
}

func (a *PluginInitializeV2) Injection(file *ast.File) error {
	var achieve bool // 是否已经实现
	for i := 0; i < len(file.Decls); i++ {
		v1, o1 := file.Decls[i].(*ast.GenDecl)
		if o1 {
			for j := 0; j < len(v1.Specs); j++ {
				v2, o2 := v1.Specs[j].(*ast.ValueSpec)
				if o2 {
					for k := 0; k < len(v2.Names); k++ {
						if v2.Names[k].Obj.Kind == ast.Var && v2.Names[k].Name == "_" {
							v3, o3 := v2.Type.(*ast.SelectorExpr)
							if o3 {
								v4, o4 := v3.X.(*ast.Ident)
								if o4 {
									if (v4.Name == "interfaces" || v4.Name == "plugin") && v3.Sel.Name == "Plugin" {
										for l := 0; l < len(v2.Values); l++ {
											v5, o5 := v2.Values[l].(*ast.CallExpr)
											if o5 {
												v6, o6 := v5.Fun.(*ast.ParenExpr)
												if o6 {
													v7, o7 := v6.X.(*ast.StarExpr)
													if o7 {
														v8, o8 := v7.X.(*ast.Ident)
														if o8 {
															achieve = true
															a.StructCamelName = v8.Name
															continue
														}
													}
												}
											}
										}
									} // var _ interfaces.Plugin = (*plugin)(nil)
								}
							}
						}
						if v2.Names[k].Obj.Kind == ast.Var {
							for l := 0; l < len(v2.Values); l++ {
								v3, o3 := v2.Values[l].(*ast.CallExpr)
								if o3 {
									v4, o4 := v3.Fun.(*ast.Ident)
									if o4 {
										if v4.Name == "new" {
											for m := 0; m < len(v3.Args); m++ {
												v5, o5 := v3.Args[m].(*ast.Ident)
												if o5 {
													if v5.Name == a.StructCamelName {
														a.StructName = v2.Names[k].Name
													} // var Plugin = new(plugin)
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
	if achieve && a.StructName != "" {
		var err error
		a.PluginPath = a.Base.AbsolutePath(a.PluginRelativePath) // 相对路径转绝对路径
		file, err = a.Parse(a.PluginPath, nil)
		if err != nil {
			return err
		}
		// TODO 插入server/initialize/plugin_biz_v2.go
		for i := 0; i < len(file.Decls); i++ {

		}
	}
	if !achieve {
		return errors.New("未实现插件接口!")
	}
	return nil
}

func (a *PluginInitializeV2) Format(filename string, writer io.Writer, file *ast.File) error {
	if filename == "" {
		filename = a.Path
	}
	return a.Base.Format(filename, writer, file)
}
