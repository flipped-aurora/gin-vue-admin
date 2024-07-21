package ast

import (
	"go/ast"
	"io"
)

type Ast interface {
	// Parse 解析文件/代码
	Parse(filename string, writer io.Writer) (file *ast.File, err error)
	// Rollback 回滚
	Rollback(file *ast.File) error
	// Injection 注入
	Injection(file *ast.File) error
	// Format 格式化输出
	Format(filename string, writer io.Writer, file *ast.File) error
}
