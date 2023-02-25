package ast

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

func ImportForAutoEnter(path string, funcName string, code string) {
	src, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, "", src, 0)
	ast.Inspect(astFile, func(node ast.Node) bool {
		if typeSpec, ok := node.(*ast.TypeSpec); ok {
			if typeSpec.Name.Name == funcName {
				if st, ok := typeSpec.Type.(*ast.StructType); ok {
					for i := range st.Fields.List {
						if t, ok := st.Fields.List[i].Type.(*ast.Ident); ok {
							if t.Name == code {
								return false
							}
						}
					}
					sn := &ast.Field{
						Type: &ast.Ident{Name: code},
					}
					st.Fields.List = append(st.Fields.List, sn)
				}
			}
		}
		return true
	})
	var out []byte
	bf := bytes.NewBuffer(out)
	err = printer.Fprint(bf, fileSet, astFile)
	if err != nil {
		return
	}
	_ = os.WriteFile(path, bf.Bytes(), 0666)
}
