package ast

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io"
)

type PluginInitializeGorm struct {
	Base
	Type         Type   // 类型
	Path         string // 文件路径
	ImportPath   string // 导包路径
	RelativePath string // 相对路径
	Business     string // 业务库
	StructName   string // 结构体名称
	PackageName  string // 包名
	IsNew        bool   // 是否使用 new 关键字 true: new(PackageName.StructName) false: &PackageName.StructName{}
}

func (a *PluginInitializeGorm) Parse(filename string, writer io.Writer) (file *ast.File, err error) {
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

func (a *PluginInitializeGorm) Rollback(file *ast.File) error {
	var needRollBackImport bool
	ast.Inspect(file, func(n ast.Node) bool {
		callExpr, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		selExpr, seok := callExpr.Fun.(*ast.SelectorExpr)
		if !seok || selExpr.Sel.Name != "AutoMigrate" {
			return true
		}
		if len(callExpr.Args) <= 1 {
			needRollBackImport = true
		}
		// 删除指定参数
		for i, arg := range callExpr.Args {
			compLit, cok := arg.(*ast.CompositeLit)
			if !cok {
				continue
			}

			cselExpr, sok := compLit.Type.(*ast.SelectorExpr)
			if !sok {
				continue
			}

			ident, idok := cselExpr.X.(*ast.Ident)
			if idok && ident.Name == a.PackageName && cselExpr.Sel.Name == a.StructName {
				// 删除参数
				callExpr.Args = append(callExpr.Args[:i], callExpr.Args[i+1:]...)
				break
			}
		}

		return true
	})

	if needRollBackImport {
		_ = NewImport(a.ImportPath).Rollback(file)
	}

	return nil
}

func (a *PluginInitializeGorm) Injection(file *ast.File) error {
	_ = NewImport(a.ImportPath).Injection(file)

	var targetCall *ast.CallExpr
	ast.Inspect(file, func(n ast.Node) bool {
		callExpr, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		selExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
		if !ok || selExpr.Sel.Name != "AutoMigrate" {
			return true
		}

		if a.isTargetAutoMigrateCall(callExpr) {
			targetCall = callExpr
			return false
		}

		return true
	})

	if targetCall == nil {
		targetCall = a.appendAutoMigrateBlock(file)
	}
	if targetCall == nil {
		return nil
	}

	if a.hasModelArg(targetCall) {
		return nil
	}

	targetCall.Args = append(targetCall.Args, &ast.CompositeLit{
		Type: &ast.SelectorExpr{
			X:   &ast.Ident{Name: a.PackageName},
			Sel: &ast.Ident{Name: a.StructName},
		},
	})
	return nil
}

func (a *PluginInitializeGorm) Format(filename string, writer io.Writer, file *ast.File) error {
	if filename == "" {
		filename = a.Path
	}
	return a.Base.Format(filename, writer, file)
}

func (a *PluginInitializeGorm) isTargetAutoMigrateCall(callExpr *ast.CallExpr) bool {
	selExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
	if !ok || selExpr.Sel.Name != "AutoMigrate" {
		return false
	}
	return exprString(selExpr.X) == exprString(a.autoMigrateReceiverExpr())
}

func (a *PluginInitializeGorm) appendAutoMigrateBlock(file *ast.File) *ast.CallExpr {
	gormFunc := FindFunction(file, "Gorm")
	if gormFunc == nil || gormFunc.Body == nil {
		return nil
	}

	src := fmt.Sprintf(`package placeholder
func Gorm() {
	if err = %s.AutoMigrate(); err != nil {
		err = errors.Wrap(err, "注册表失败!")
		zap.L().Error(fmt.Sprintf("%%+v", err))
	}
}
`, exprString(a.autoMigrateReceiverExpr()))

	parsed, err := parser.ParseFile(token.NewFileSet(), "", src, 0)
	if err != nil || len(parsed.Decls) == 0 {
		return nil
	}

	stmt := parsed.Decls[0].(*ast.FuncDecl).Body.List[0].(*ast.IfStmt)
	clearPosition(stmt)
	gormFunc.Body.List = append(gormFunc.Body.List, stmt)

	assignStmt := stmt.Init.(*ast.AssignStmt)
	callExpr := assignStmt.Rhs[0].(*ast.CallExpr)
	return callExpr
}

func (a *PluginInitializeGorm) autoMigrateReceiverExpr() ast.Expr {
	return &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   a.dbExpr(),
			Sel: &ast.Ident{Name: "WithContext"},
		},
		Args: []ast.Expr{&ast.Ident{Name: "ctx"}},
	}
}

func (a *PluginInitializeGorm) dbExpr() ast.Expr {
	if a.Business == "" {
		return &ast.SelectorExpr{
			X:   &ast.Ident{Name: "global"},
			Sel: &ast.Ident{Name: "GVA_DB"},
		}
	}
	return &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "global"},
			Sel: &ast.Ident{Name: "MustGetGlobalDBByDBName"},
		},
		Args: []ast.Expr{
			&ast.BasicLit{
				Kind:  token.STRING,
				Value: fmt.Sprintf("\"%s\"", a.Business),
			},
		},
	}
}

func (a *PluginInitializeGorm) hasModelArg(callExpr *ast.CallExpr) bool {
	for _, arg := range callExpr.Args {
		compositeLit, ok := arg.(*ast.CompositeLit)
		if !ok {
			continue
		}
		selectorExpr, ok := compositeLit.Type.(*ast.SelectorExpr)
		if !ok {
			continue
		}
		packageIdent, ok := selectorExpr.X.(*ast.Ident)
		if ok && packageIdent.Name == a.PackageName && selectorExpr.Sel.Name == a.StructName {
			return true
		}
	}
	return false
}

func exprString(expr ast.Expr) string {
	var buffer bytes.Buffer
	_ = format.Node(&buffer, token.NewFileSet(), expr)
	return buffer.String()
}
