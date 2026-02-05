package ast

import (
	"go/ast"
	"go/token"
	"io"
	"strconv"
	"strings"
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
			filename = a.PluginPath
			a.RelativePath = a.Base.RelativePath(a.PluginPath)
			return a.Base.Parse(filename, writer)
		}
		a.PluginPath = a.Base.AbsolutePath(a.RelativePath)
		filename = a.PluginPath
	}
	return a.Base.Parse(filename, writer)
}

func (a *PluginInitializeV2) Injection(file *ast.File) error {
	importPath := strings.TrimSpace(a.ImportPath)
	if importPath == "" {
		return nil
	}
	importPath = strings.Trim(importPath, "\"")
	if importPath == "" || CheckImport(file, importPath) {
		return nil
	}

	importSpec := &ast.ImportSpec{
		Name: ast.NewIdent("_"),
		Path: &ast.BasicLit{Kind: token.STRING, Value: strconv.Quote(importPath)},
	}
	var importDecl *ast.GenDecl
	for _, decl := range file.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		if genDecl.Tok == token.IMPORT {
			importDecl = genDecl
			break
		}
	}
	if importDecl == nil {
		file.Decls = append([]ast.Decl{
			&ast.GenDecl{
				Tok:   token.IMPORT,
				Specs: []ast.Spec{importSpec},
			},
		}, file.Decls...)
		return nil
	}
	importDecl.Specs = append(importDecl.Specs, importSpec)
	return nil
}

func (a *PluginInitializeV2) Rollback(file *ast.File) error {
	return nil
}

func (a *PluginInitializeV2) Format(filename string, writer io.Writer, file *ast.File) error {
	if filename == "" {
		filename = a.PluginPath
	}
	return a.Base.Format(filename, writer, file)
}
