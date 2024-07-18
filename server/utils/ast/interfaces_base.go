package ast

import (
	"github.com/pkg/errors"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io"
)

type Base struct{}

func (a *Base) Parse(filename string, writer io.Writer) (file *ast.File, err error) {
	fileSet := token.NewFileSet()
	if writer != nil {
		file, err = parser.ParseFile(fileSet, filename, nil, parser.ParseComments)
	} else {
		file, err = parser.ParseFile(fileSet, filename, writer, parser.ParseComments)
	}
	if err != nil {
		return nil, errors.Wrapf(err, "[filepath:%s]打开/解析文件失败!", filename)
	}
	return file, nil
}

func (a *Base) Rollback(file *ast.File) {
	return
}

func (a *Base) Injection(file *ast.File) {
	return
}

func (a *Base) Format(filename string, writer io.Writer, file *ast.File) error {
	fileSet := token.NewFileSet()
	err := format.Node(writer, fileSet, file)
	if err != nil {
		return errors.Wrapf(err, "[filepath:%s]注入失败!", filename)
	}
	return nil
}
