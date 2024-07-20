package ast

import (
	"go/ast"
	"io"
)

type PluginInitializeV2 struct {
	Base
	Type         Type   // 类型
	Path         string // 文件路径
	PluginPath   string // 插件路径
	RelativePath string // 相对路径
	ImportPath   string // 导包路径
	StructName   string // 结构体名称
	PackageName  string // 包名
}

func (a *PluginInitializeV2) Parse(filename string, writer io.Writer) (file *ast.File, err error) {
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

func (a *PluginInitializeV2) Injection(file *ast.File) error {
	return nil
}

func (a *PluginInitializeV2) Rollback(file *ast.File) error {
	return nil
}

func (a *PluginInitializeV2) Format(filename string, writer io.Writer, file *ast.File) error {
	if filename == "" {
		filename = a.Path
	}
	return a.Base.Format(filename, writer, file)
}
