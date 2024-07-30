package ast

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/pkg/errors"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
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

func (a *Base) Rollback(file *ast.File) error {
	return nil
}

func (a *Base) Injection(file *ast.File) error {
	return nil
}

func (a *Base) Format(filename string, writer io.Writer, file *ast.File) error {
	fileSet := token.NewFileSet()
	if writer == nil {
		open, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0666)
		defer open.Close()
		if err != nil {
			return errors.Wrapf(err, "[filepath:%s]打开文件失败!", filename)
		}
		writer = open
	}
	err := format.Node(writer, fileSet, file)
	if err != nil {
		return errors.Wrapf(err, "[filepath:%s]注入失败!", filename)
	}
	return nil
}

// RelativePath 绝对路径转相对路径
func (a *Base) RelativePath(filePath string) string {
	server := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server)
	hasServer := strings.Index(filePath, server)
	if hasServer != -1 {
		filePath = strings.TrimPrefix(filePath, server)
		keys := strings.Split(filePath, string(filepath.Separator))
		filePath = path.Join(keys...)
	}
	return filePath
}

// AbsolutePath 相对路径转绝对路径
func (a *Base) AbsolutePath(filePath string) string {
	server := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server)
	keys := strings.Split(filePath, "/")
	filePath = filepath.Join(keys...)
	filePath = filepath.Join(server, filePath)
	return filePath
}
