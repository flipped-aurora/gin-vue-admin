package ast

import (
	"go/ast"
	"go/token"
)

type Import struct {
	file       *ast.File
	importPath string // 导包路径
}

func NewImport(file *ast.File, importPath string) *Import {
	return &Import{file: file, importPath: importPath}
}

func (a *Import) Rollback() error {
	if a.importPath == "" {
		return nil
	}
	for i := 0; i < len(a.file.Decls); i++ {
		v1, o1 := a.file.Decls[i].(*ast.GenDecl)
		if o1 {
			if v1.Tok != token.IMPORT {
				break
			}
			for j := 0; j < len(v1.Specs); j++ {
				v2, o2 := v1.Specs[j].(*ast.ImportSpec)
				if o2 && v2.Path.Value == a.importPath {
					v1.Specs = append(v1.Specs[:j], v1.Specs[j+1:]...)
					if len(v1.Specs) == 0 {
						a.file.Decls = append(a.file.Decls[:i], a.file.Decls[i+1:]...)
					}
					break
				}
			}
		}
	}
	return nil
}

func (a *Import) Injection() error {
	if a.importPath == "" {
		return nil
	}
	var has bool
	for i := 0; i < len(a.file.Decls); i++ {
		v1, o1 := a.file.Decls[i].(*ast.GenDecl)
		if o1 {
			if v1.Tok != token.IMPORT {
				break
			}
			for j := 0; j < len(v1.Specs); j++ {
				v2, o2 := v1.Specs[j].(*ast.ImportSpec)
				if o2 && v2.Path.Value == a.importPath {
					has = true
					break
				}
			}
			if !has {
				spec := &ast.ImportSpec{
					Path: &ast.BasicLit{Kind: token.STRING, Value: a.importPath},
				}
				v1.Specs = append(v1.Specs, spec)
				return nil
			}
		}
	}
	if !has {
		decls := a.file.Decls
		a.file.Decls = make([]ast.Decl, 0, len(a.file.Decls)+1)
		decl := &ast.GenDecl{
			Tok: token.IMPORT,
			Specs: []ast.Spec{
				&ast.ImportSpec{
					Path: &ast.BasicLit{Kind: token.STRING, Value: a.importPath},
				},
			},
		}
		a.file.Decls = append(a.file.Decls, decl)
		a.file.Decls = append(a.file.Decls, decls...)
	}
	return nil
}
