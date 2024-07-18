package ast

import (
	"go/ast"
	"go/token"
	"io"
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

func (a *Import) Rollback(file *ast.File) {
	if a.ImportPath == "" {
		return
	}
	for i := 0; i < len(file.Decls); i++ {
		v1, o1 := file.Decls[i].(*ast.GenDecl)
		if o1 {
			if v1.Tok != token.IMPORT {
				break
			}
			for j := 0; j < len(v1.Specs); j++ {
				v2, o2 := v1.Specs[j].(*ast.ImportSpec)
				if o2 && v2.Path.Value == a.ImportPath {
					v1.Specs = append(v1.Specs[:j], v1.Specs[j+1:]...)
					if len(v1.Specs) == 0 {
						file.Decls = append(file.Decls[:i], file.Decls[i+1:]...)
					} // 如果没有import声明，就删除, 如果不删除则会出现import()
					break
				}
			}
		}
	}
}

func (a *Import) Injection(file *ast.File) {
	if a.ImportPath == "" {
		return
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
				if o2 && v2.Path.Value == a.ImportPath {
					has = true
					break
				}
			}
			if !has {
				spec := &ast.ImportSpec{
					Path: &ast.BasicLit{Kind: token.STRING, Value: a.ImportPath},
				}
				v1.Specs = append(v1.Specs, spec)
				return
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
}

func (a *Import) Format(filename string, writer io.Writer, file *ast.File) error {
	return a.Base.Format(filename, writer, file)
}
