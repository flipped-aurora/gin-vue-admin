package ast

import (
	"go/ast"
	"go/token"
	"io"
	"strings"
)

type Import struct {
	Base
	ImportPath string // 导包路径
}

func NewImport(importPath string) *Import {
	return &Import{ImportPath: importPath}
}

func (a *Import) Parse(filename string, writer io.Writer) (file *ast.File, err error) {
	return a.Base.Parse(filename, writer)
}

func (a *Import) Rollback(file *ast.File) error {
	if a.ImportPath == "" {
		return nil
	}
	for i := 0; i < len(file.Decls); i++ {
		v1, o1 := file.Decls[i].(*ast.GenDecl)
		if o1 {
			if v1.Tok != token.IMPORT {
				break
			}
			for j := 0; j < len(v1.Specs); j++ {
				v2, o2 := v1.Specs[j].(*ast.ImportSpec)
				if o2 && strings.HasSuffix(a.ImportPath, v2.Path.Value) {
					v1.Specs = append(v1.Specs[:j], v1.Specs[j+1:]...)
					if len(v1.Specs) == 0 {
						file.Decls = append(file.Decls[:i], file.Decls[i+1:]...)
					} // 如果没有import声明，就删除, 如果不删除则会出现import()
					break
				}
			}
		}
	}
	return nil
}

func (a *Import) Injection(file *ast.File) error {
	if a.ImportPath == "" {
		return nil
	}
	var has bool
	for i := 0; i < len(file.Decls); i++ {
		v1, o1 := file.Decls[i].(*ast.GenDecl)
		if o1 {
			if v1.Tok != token.IMPORT {
				break
			}
			for j := 0; j < len(v1.Specs); j++ {
				v2, o2 := v1.Specs[j].(*ast.ImportSpec)
				if o2 && strings.HasSuffix(a.ImportPath, v2.Path.Value) {
					has = true
					break
				}
			}
			if !has {
				spec := &ast.ImportSpec{
					Path: &ast.BasicLit{Kind: token.STRING, Value: a.ImportPath},
				}
				v1.Specs = append(v1.Specs, spec)
				return nil
			}
		}
	}
	if !has {
		decls := file.Decls
		file.Decls = make([]ast.Decl, 0, len(file.Decls)+1)
		decl := &ast.GenDecl{
			Tok: token.IMPORT,
			Specs: []ast.Spec{
				&ast.ImportSpec{
					Path: &ast.BasicLit{Kind: token.STRING, Value: a.ImportPath},
				},
			},
		}
		file.Decls = append(file.Decls, decl)
		file.Decls = append(file.Decls, decls...)
	} // 如果没有import声明，就创建一个, 主要要放在第一个
	return nil
}

func (a *Import) Format(filename string, writer io.Writer, file *ast.File) error {
	return a.Base.Format(filename, writer, file)
}
