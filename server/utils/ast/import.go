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
		if v1.Tok != token.IMPORT {
			break
		}
		if o1 {
			for j := 0; j < len(v1.Specs); j++ {
				v2, o2 := v1.Specs[j].(*ast.ImportSpec)
				if o2 && v2.Path.Value == a.importPath {
					v1.Specs = append(v1.Specs[:j], v1.Specs[j+1:]...)
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
	found := false
	for i := 0; i < len(a.file.Decls); i++ {
		v1, o1 := a.file.Decls[i].(*ast.GenDecl)
		if v1.Tok != token.IMPORT {
			break
		}
		if o1 {
			for j := 0; j < len(v1.Specs); j++ {
				v2, o2 := v1.Specs[j].(*ast.ImportSpec)
				if o2 && v2.Path.Value == a.importPath {
					found = true
					break
				}
			}
			if !found {
				v1.Specs = append(v1.Specs, &ast.ImportSpec{
					Path: &ast.BasicLit{Kind: token.STRING, Value: a.importPath},
				})
				break
			}
		}
	}
	return nil
}
